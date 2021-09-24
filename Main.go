package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

func main() {

	var path string
	flag.StringVar(&path, "file", "", "账户配置")

	flag.Parse()
	if len(path)==0 {
		fmt.Printf("没有配置账户文件！")
		os.Exit(1)
	}
	byte, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("打开用户信息失败%s", err.Error())
	}
	fmt.Printf("==================%s================\r\n",time.Now())
	fmt.Printf("续期任务开始拉~")
	
	var config Config
	yaml.Unmarshal(byte, &config)
	for item := range config.Accounts {
		LoginEuserv(config.Accounts[item])
	}
	fmt.Printf("发送邮件通知")
	if config.SmtpSSL{
		SendMailTls(config.SmtpUserName,config.SmtpPassWord,config.SmtpServer,"deji续期",Log,config.ContactsMail)
	}else{
		SendMail(config.SmtpUserName,config.SmtpPassWord,config.SmtpServer,"deji续期",Log,config.ContactsMail)
	}


	fmt.Printf("\r\n所有账号全部续期完成")
}
//func  SenMail(user string,password string)  {
//	auth := sasl.NewPlainClient("", user, password)
//	to := []string{"xhuan_blog@yeah.net"}
//	msg := strings.NewReader("To: xhuan_blog@yeah.net\r\n" +
//		"Subject: Euserv续期提醒!\r\n" +
//		"\r\n" +Log+"\r\n")
//	er2:= smtp.SendMail("smtp.yeah.net:25", auth, user, to, msg)
//	if er2 != nil {
//		fmt.Print(er2)
//	}
//}
