package config

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Config struct {
	Address  string
	User     string
	Pwd      string
	DataBase string
}

var (
	db     *sql.DB
	config Config
)

func init() {
	err := loadConfig()
	if err != nil {
		log.Fatal("can't parseConfig,err:", err)
		return
	}
	err = loadMysql()
	if err != nil {
		log.Fatal("can't open the sql,err:", err)
		return
	}
}

func loadConfig() error {
	p := "../pkg/config/config.toml"
	_, err := toml.DecodeFile(p, &config)
	if err != nil {
		p = "../config/config.toml"
		_, err = toml.DecodeFile(p, &config)
		if err != nil {
			return err
		}
	}
	log.Printf("解析配置config成功！%v\n", config)
	return nil
}

func loadMysql() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Pwd, config.Address, config.DataBase)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	log.Println("数据库连接成功！")
	return nil
}

func GetDb() *sql.DB {
	return db
}
