package model

import (
	"fmt"
	"time"
	"v1/pkg"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

// 基本模型
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewDBEngine() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			pkg.Ds.Username, pkg.Ds.Password, pkg.Ds.DBName), // DSN data source name
		DefaultStringSize:         256,  // string 类型字段的默认长度
		DisableDatetimePrecision:  true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: true, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	})
	if err != nil {
		pkg.Log.Panic("connect database error", zap.String("error", err.Error()))
	}
	db.AutoMigrate(&User{})
	DB = db
}
