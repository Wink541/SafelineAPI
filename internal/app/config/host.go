package config

import (
	"SafelineAPI/internal/app/logger"
	"SafelineAPI/internal/app/safeLineApi"
	"fmt"
	"net/url"
)

type Host struct {
	HostName string `json:"HostName"`
	Port     string `json:"Port"`
}

func (host Host) String() string {
	if host.Port == "" {
		return host.HostName
	}
	return host.HostName + ":" + host.Port
}

func (host Host) Verify() bool {
	if host.HostName == "" {
		logger.Warning.Printf("未设置 %s主机名称%s: 请检查配置文件中的 %sSafeLine.Host.HostName%s 参数", logger.Cyan, logger.Reset, logger.Yellow, logger.Reset)
		return true
	}
	return false
}

func (host Host) VerifyCommand() bool {
	if host.HostName == "" {
		logger.Warning.Printf("未设置 %s主机名称%s: 请检查命令中的 %s-h%s 参数", logger.Cyan, logger.Reset, logger.Yellow, logger.Reset)
		return true
	}
	return false
}

func (host Host) Url() *safeLineApi.URL {
	var u *url.URL
	if host.Port == "" {
		u, _ = url.Parse(fmt.Sprintf("https://%s", host.HostName))
	} else {
		u, _ = url.Parse(fmt.Sprintf("https://%s:%s", host.HostName, host.Port))
	}
	return (*safeLineApi.URL)(u)
}
