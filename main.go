package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/google/uuid"
)

type MailServiceConfig struct {
	Token string
}

func ReadConfig(path string) *MailServiceConfig {
	var mail_service_config = new(MailServiceConfig)
	file, open_err := os.Open("/etc/mail/config.json")
	if os.IsNotExist(open_err) {
		mail_service_config.Token = uuid.New().String()
		file, _ := os.Create(path)
		var bytes, _ = json.Marshal(mail_service_config)
		file.Write(bytes)
	}
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	json.Unmarshal(bytes, mail_service_config)
	return mail_service_config
}

func main() {

}
