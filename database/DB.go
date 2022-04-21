package database

import (
	"context"
	"database/sql"
	"time"

	// register mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/sirupsen/logrus"

	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/sunliang711/goutils/mongodb"
	umysql "github.com/sunliang711/goutils/mysql"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/spf13/viper"
)

var (
	MysqlConn    *sql.DB
	MysqlORMConn *gorm.DB
	MongoConn    *mongo.Client
)

func InitDB(tables []interface{}) {
	logrus.Infof("database init()...")
	if viper.GetBool("mysql.enable") {
		dsn := viper.GetString("mysql.dsn")
		initMysql(dsn)
	} else {
		logrus.Info("mysql is disabled.")
	}

	if viper.GetBool("mongodb.enable") {
		dsn := viper.GetString("mongodb.url")
		initMongo(dsn)
	} else {
		logrus.Info("mongodb is disabled.")
	}

	migration(tables)
}

func migration(tables []interface{}) {
	if viper.GetBool("mysql.enable") && viper.GetBool("mysql.orm") {
		for idx, table := range tables {
			result := MysqlORMConn.AutoMigrate(table)
			if result.Error != nil {
				logrus.Fatalf("failed to migrate table[%d]: %+v %v", idx, table, result.Error)
			}
		}
		logrus.Infof("migration done")
	}
}

// initMysql open mysql with dsn
func initMysql(dsn string) {
	logrus.Infof("try to connect to mysql: '%v'", dsn)
	var err error
	if viper.GetBool("mysql.orm") {
		MysqlORMConn, err = gorm.Open("mysql", dsn)
		if viper.GetBool("mysql.orm_logger") {
			MysqlORMConn.LogMode(true)
		}
		logrus.Infof("use gorm driver...")
		if err != nil {
			panic(err)
		}
		db := MysqlORMConn.DB()
		if db != nil {
			db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
			db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
			db.SetConnMaxLifetime(time.Hour)
		}

	} else {
		MysqlConn, err = umysql.New(dsn, viper.GetInt("mysql.max_idle_conns"), viper.GetInt("mysql.max_open_conns"))
		if err != nil {
			panic(err)
		}
	}
	logrus.Infof("connected to mysql")
}

// CloseMysql close mysql connection
func closeMysql() {
	if viper.GetBool("mysql.enable") {
		logrus.Infoln("close mysql")
		if viper.GetBool("mysql.orm") {
			MysqlORMConn.Close()
		} else {
			MysqlConn.Close()
		}
	}
}

// InitMongo opens a mongodb connection
func initMongo(url string) {
	logrus.Infof("try to connect to mongodb: '%v'", url)
	var err error
	MongoConn, err = mongodb.New(url, 5)
	if err != nil {
		panic(err)
	}
	logrus.Infof("connected to mongodb")
}

func closeMongo() {
	if viper.GetBool("mongodb.enable") {
		logrus.Infof("close mongodb")
		MongoConn.Disconnect(context.Background())
	}
}

func Release() {
	closeMongo()
	closeMysql()
}
