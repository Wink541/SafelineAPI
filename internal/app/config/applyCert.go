package config

import (
	"SafelineAPI/internal/app/logger"
	"log"
)

type ApplyCert struct {
	Days              int    `json:"Days"`
	Email             string `json:"Email"`
	SavePath          string `json:"SavePath"`
	DNSProviderConfig `json:"DNSProviderConfig"`
}

func (applyCert *ApplyCert) GetDays() int {
	return applyCert.Days
}

func (applyCert *ApplyCert) GetEmail() string {
	return applyCert.Email
}

func (applyCert *ApplyCert) Verify() bool {
	var flag = false
	if applyCert.Days == 0 {
		applyCert.Days = 30
	}

	if applyCert.SavePath == "" {
		applyCert.SavePath = "/tmp/ssl"
	}

	if applyCert.DNSProvider == "" {
		logger.Warning.Printf("未设置 %sDNS服务提供商%s: 请检查配置文件中的 %sApplyCert.DNSProviderConfig.DNSProvider%s 参数", logger.Cyan, logger.Reset, logger.Yellow, logger.Reset)
		flag = true
	}
	if applyCert.Email == "" {
		logger.Warning.Printf("未设置 %s证书申请邮箱%s: 请检查配置文件中的 %sApplyCert.Email%s 参数", logger.Cyan, logger.Reset, logger.Yellow, logger.Reset)
		flag = true
	}
	if !flag {
		log.Printf("%sApplyCert%s 相关配置检验完成!", logger.Cyan, logger.Reset)
	}
	return flag
}

func (applyCert *ApplyCert) VerifyCommand() bool {
	var flag = false
	if applyCert.Days == 0 {
		applyCert.Days = 30
	}

	if applyCert.SavePath == "" {
		applyCert.SavePath = "/tmp/ssl"
	}

	if applyCert.DNSProvider == "" {
		logger.Warning.Printf("未设置 %sDNS服务提供商%s: 请检查命令中的 %s-D%s 参数", logger.Cyan, logger.Reset, logger.Yellow, logger.Reset)
		flag = true
	}
	if applyCert.Email == "" {
		logger.Warning.Printf("未设置 %s证书申请邮箱%s: 请检查命令中的 %s-e%s 参数", logger.Cyan, logger.Reset, logger.Yellow, logger.Reset)
		flag = true
	}
	if !flag {
		log.Printf("%sApplyCert%s 相关配置检验完成!", logger.Cyan, logger.Reset)
	}
	return flag

}
