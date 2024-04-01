package main

import (
	"encoding/json"
	"io"
	"log"
	"net/smtp"
	"os"

	"github.com/google/uuid"
)

type Mailer struct {
	SmtpServer string
	Port       string
	Sender     string
	Password   string
}

func (m *Mailer) SendEmail(to string, message string) {
	// 创建邮件
	msg := []byte("To: " + to + "\r\n" +
		"Subject: Golang Mail Service \r\n" +
		"\r\n" +
		message + "\r\n")

	// 发送邮件
	auth := smtp.PlainAuth("", m.Sender, m.Password, m.SmtpServer)
	err := smtp.SendMail(m.SmtpServer+":"+m.Port, auth, m.Sender, []string{to}, msg)
	if err != nil {
		log.Fatal(err)
	}
}

type MailServiceConfig struct {
	Token  string
	Mailer Mailer
}

func ReadConfig(path string) (*MailServiceConfig, error) {
	var mail_service_config = new(MailServiceConfig)
	file, open_err := os.Open(path)
	if os.IsNotExist(open_err) {
		mail_service_config.Token = uuid.New().String()
		file, create_err := os.Create(path)
		if create_err != nil {
			return nil, create_err
		}
		defer file.Close()
		bytes, marshal_err := json.Marshal(mail_service_config)
		if marshal_err != nil {
			return nil, marshal_err
		}
		_, write_err := file.Write(bytes)
		if write_err != nil {
			return nil, write_err
		}
	} else if open_err != nil {
		return nil, open_err
	} else {
		defer file.Close()
		bytes, read_err := io.ReadAll(file)
		if read_err != nil {
			return nil, read_err
		}
		unmarshal_err := json.Unmarshal(bytes, mail_service_config)
		if unmarshal_err != nil {
			return nil, unmarshal_err
		}
	}
	return mail_service_config, nil
}

func main() {

}
