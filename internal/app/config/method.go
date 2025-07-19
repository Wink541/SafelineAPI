package config

import (
	"log"
	"os"
)

func (config *Config) Verify() {
	a := config.SafeLine.Verify()
	b := config.ApplyCert.Verify()
	if a || b {
		log.Printf("配置检查完毕,请检查相关配置后重新运行！")
		os.Exit(0)
	}
	log.Printf("配置检查完毕,即将开始更新证书！")
}

func (config *Config) VerifyCommand() {
	a := config.SafeLine.VerifyCommand()
	b := config.ApplyCert.VerifyCommand()
	if a || b {
		log.Printf("配置检查完毕,请检查相关配置后重新运行！")
		os.Exit(0)
	}
	log.Printf("配置检查完毕,即将开始更新证书！")
}
