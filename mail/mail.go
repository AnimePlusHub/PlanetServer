package mail

import (
	"crypto/tls"
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

var (
	userName, password, host, subject, body string
)

func init() {
	viper.SetConfigName("mail")
	viper.AddConfigPath("./conf") // ./conf or ../../conf
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic(fmt.Errorf("读取邮件配置文件失败: %s", err))
	}
	userName = viper.GetString("mail.UserName")
	password = viper.GetString("mail.Password")
	host = viper.GetString("mail.Host")
	subject = viper.GetString("mail.Subject")
	body = viper.GetString("mail.Body")
}

func SendEmail(toEmail string) {

	m := gomail.NewMessage()
	m.SetHeader(`From`, userName)
	m.SetHeader(`To`, toEmail)
	m.SetHeader(`Subject`, subject)
	m.SetBody("text/html", body) // 发送html格式邮件，发送的内容
	// if cc != "" {
	// 	m.SetHeader("Cc", cc)  // 抄送
	// }

	d := gomail.NewDialer(host, 25, userName, password)
	// 修改TLSconfig
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("邮件发送失败：", err)
		return
	}
	fmt.Println("邮件发送成功！")
}
