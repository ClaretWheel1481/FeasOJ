package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DbName     string `xml:"dbname"`
	DbUser     string `xml:"dbuser"`
	DbPassword string `xml:"dbpassword"`
	DbAddress  string `xml:"dbaddress"`
	DbPort     string `xml:"dbport"`
}

func writeConfig(dbName, dbUser, dbPassword, dbAddress string) error {
	// 创建配置文件
	config := Config{
		DbName:     dbName,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		DbAddress:  dbAddress,
	}
	configXml, err := xml.Marshal(config)
	if err != nil {
		return err
	}
	err = os.WriteFile("config.xml", configXml, 0644)
	return err
}

// 连接数据库、创建表
func initSql() bool {
	//判断是否有config.xml文件，没有则输入
	if _, err := os.Stat("config.xml"); os.IsNotExist(err) {
		fmt.Println("[FeasOJ]请输入数据库连接信息：")
		fmt.Print("数据库名：")
		var dbName string
		fmt.Scanln(&dbName)
		fmt.Print("数据库用户名：")
		var dbUser string
		fmt.Scanln(&dbUser)
		fmt.Print("数据库密码：")
		var dbPassword string
		fmt.Scanln(&dbPassword)
		fmt.Print("数据库地址（加上端口）：")
		var dbAddress string
		fmt.Scanln(&dbAddress)
		fmt.Println("[FeasOJ]正在保存数据库连接信息...")
		// 将连接信息作为xml保存到目录
		err := writeConfig(dbName, dbUser, dbPassword, dbAddress)
		if err != nil {
			fmt.Println("[FeasOJ]保存失败，请检查权限。")
			return false
		}
	}

	// 读取config.xml文件并赋值
	configFile, err := os.Open("config.xml")
	if err != nil {
		return false
	}
	defer configFile.Close()
	var config Config
	err = xml.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return false
	}
	dbName := config.DbName
	dbUser := config.DbUser
	dbPassword := config.DbPassword
	dbAddress := config.DbAddress

	fmt.Println("[FeasOJ]连接数据库中...")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbAddress + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return false
	}
	fmt.Println("[FeasOJ]数据库连接成功。")
	fmt.Println("[FeasOJ]创建数据表中...")
	db.AutoMigrate(&User{}, &Problem{})
	fmt.Println("[FeasOJ]创建数据表成功。")
	return true
}
