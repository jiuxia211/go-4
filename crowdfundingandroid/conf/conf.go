package conf

import (
	"fmt"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	DbUser     string
	DbPassWord string
	DbHost     string
	DbName     string
	DbPort     string
	Path       string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	AccessKey   string
	SerectKey   string
	Bucket      string
	QiniuServer string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误")
	}
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbName = file.Section("mysql").Key("DbName").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	// fmt.Println(AppMode, HttpPort, DbUser, DbPassWord, DbHost, DbName)
	//"用户名:密码@tcp(ip:port)/dbName?charset=utf8mb4&parseTime=True&loc=Local"
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
	// fmt.Println(ValidEmail, SmtpHost, SmtpEmail, SmtpPass)

	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SerectKey = file.Section("qiniu").Key("SerectKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()
	// fmt.Println(AccessKey, SerectKey, Bucket, QiniuServer)
	Path = strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort,
		")/", DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")

}
