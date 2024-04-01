package main

import "testing"

func TestReadConfig(t *testing.T) {
	t.Log(ReadConfig("./config.json"))
}
