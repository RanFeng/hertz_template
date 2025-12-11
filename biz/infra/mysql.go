package infra

import (
	"context"
	"fmt"
	"hertz/demo/biz/conf"
	"hertz/demo/biz/consts"

	"github.com/RanFeng/ierror"
	"github.com/RanFeng/ilog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MysqlDB *gorm.DB

func InitMysql() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.MustGet[string]("mysql_user"),
		conf.MustGet[string]("mysql_pass"),
		conf.MustGet[string]("mysql_addr"),
		conf.MustGet[string]("mysql_db_name"),
	)
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}

// TransactionTemplate 事务运行模版
func Transaction(ctx context.Context, db *gorm.DB, fc func(tx *gorm.DB) error) (err error) {
	// 标记panic
	panicked := true
	// 开启事务
	tx := db.WithContext(ctx).Debug().Begin()
	if tx.Error != nil {
		// 事务开启异常，直接返回错误
		return tx.Error
	}
	// 后置检查
	defer func() {
		// 如果出现panic或返回err，都需要回滚事务
		if panicked || err != nil {
			rollBackErr := tx.Rollback().Error
			if rollBackErr != nil {
				ilog.EventError(ctx, rollBackErr, "rollback_error")
				err = ierror.NewIError(consts.MySqlError, "rollback_error:"+rollBackErr.Error())
			}
		}
	}()
	// 执行业务方法
	err = fc(tx)
	// 方法执行无err，提交事务
	if err == nil {
		commitError := tx.Commit().Error
		// 标记无panic
		panicked = false
		if commitError == nil {
			// 事务提交成功，返回nil
			return nil
		}
		// 事务提交失败，返回err
		return err
	}
	// 方法执行有err，标记无panic，返回err，defer里进行回滚
	panicked = false
	return err
}
