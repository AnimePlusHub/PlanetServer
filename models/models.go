package models

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// type Model struct {
// 	ID         int       `gorm:"primary_key" json:"id"`
// 	CreatedOn  time.Time `json:"created_on"`
// 	ModifiedOn time.Time `json:"modified_on"`
// }

func init() {
	var dbName, dbUser, dbPwd, host string
	viper.SetConfigName("mysql")
	viper.AddConfigPath("../../../conf")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic(fmt.Errorf("读取配置文件失败: %s", err))
	}
	// 读取mysql配置文件
	dbName = viper.GetString("mysql.DBNAME")
	// dbType = viper.GetString("database.TYPE")
	dbUser = viper.GetString("mysql.USER")
	dbPwd = viper.GetString("mysql.PWD")
	host = viper.GetString("mysql.HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPwd, host, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix:   "",
		SingularTable: true, //使用单数表名，启用该选项后，`User`表将是 `user`
		NameReplacer:  nil,
		NoLowerCase:   false,
	},
	})

	if err != nil {
		log.Println(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}
