package main

import "testing"

func TestReadConfig(t *testing.T) {
	config, err := ReadConfig("./test/config.json")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", config)
}

func TestSend(t *testing.T) {
	config, err := ReadConfig("./test/config.json")
	if err != nil {
		t.Error(err)
	}
	config.Mailer.SendEmail("mrx800@163.com", "smtp test")
}
