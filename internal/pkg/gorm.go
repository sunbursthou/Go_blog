package pkg

import (
	"fmt"
	"log"
	"time"

	// "github.com/go-sql-driver/mysql"
	"github.com/sunbursthou/Go_blog/configs"
	"github.com/sunbursthou/Go_blog/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitMySQL() *gorm.DB {
	mysqlCfg := configs.Cfg.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlCfg.Username,
		mysqlCfg.Password,
		mysqlCfg.Host,
		mysqlCfg.Port,
		mysqlCfg.Dbname,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// gorm日志模式：error
		Logger: logger.Default.LogMode(getLogMode(configs.Cfg.Mysql.LogMode)),
		// 禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled

		},
	})
	if err != nil {
		log.Fatalln("Mysql连接失败，请检查参数是否正确！", err)
	}
	log.Println("Mysql 连接成功！")

	db.AutoMigrate(
		&model.Article{},
		&model.ArticleToTag{},
		&model.BlackList{},
		&model.Category{},
		&model.Collections{},
		&model.ArticleComment{},
		&model.ChildComment{},
		&model.Message{},
		&model.Notice{},
		&model.Tag{},
		&model.User{},
		&model.Announcement{},
		&model.Siteinfo{},
	)
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Second * 10)
	return db

}

// 根据字符串获取Mysql对应 LogLevel
func getLogMode(str string) logger.LogLevel {
	switch str {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}
