package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

func mkConfig() {
	fmt.Println("To make config.ini file...")
	cfg := ini.Empty()
	cfg.Section("SMTP").Key("server").SetValue("smtp.nlsn.media")
	cfg.Section("SMTP").Key("port").SetValue("25")
	cfg.SaveTo("config.ini")
}

func main() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
		mkConfig()
		os.Exit(1)
	}

	fmt.Println("==============================================")
	fmt.Println("SMTP server:", cfg.Section("SMTP").Key("server").String())
	fmt.Println("SMTP port:", cfg.Section("SMTP").Key("port").String())
	fmt.Printf("%+v\n", cfg)
	fmt.Printf("%+v\n", err)
}
