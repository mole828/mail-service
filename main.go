package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Mailer struct {
	SmtpServer string
	Port       string
	Sender     string
	Password   string
}

func (m *Mailer) SendEmail(recipient string, message string) error {
	// 创建邮件
	msg := []byte("To: " + recipient + "\r\n" +
		"Subject: Go Mail service\r\n" +
		"\r\n" +
		message + "\r\n")

	// 发送邮件
	auth := smtp.PlainAuth("", m.Sender, m.Password, m.SmtpServer)
	err := smtp.SendMail(m.SmtpServer+":"+m.Port, auth, m.Sender, []string{recipient}, msg)
	return err
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

type Email struct {
	To      string
	Message string
	Token   string
}

func main() {
	app := gin.New()
	config, _ := ReadConfig("./test/config.json")
	app.POST("/send_email", func(c *gin.Context) {
		var email Email
		if err := c.ShouldBindJSON(&email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if email.Token != config.Token {
			c.JSON(http.StatusForbidden, gin.H{"error": "wrong token"})
			return
		}

		err := config.Mailer.SendEmail(email.To, email.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Email sent!"})
	})
	app.Run(":8080")
}
