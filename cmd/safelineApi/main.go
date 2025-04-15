package main

import (
	"SafelineAPI/internal/app/config"
	"SafelineAPI/internal/app/logger"
	"SafelineAPI/internal/app/safeLineApi"
	"SafelineAPI/pkg/moudle"
	"SafelineAPI/pkg/services"
	"SafelineAPI/pkg/utils"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	logger.LogInit()
	var conf config.Config
	if len(os.Args) == 1 || os.Args[1][0] == '-' {
		conf.Command()
	} else {
		conf.Read(os.Args[1])
	}

	err := os.MkdirAll(conf.SavePath, os.ModePerm)
	if err != nil {
		logger.Error.Printf("创建文件夹 %s%s%s 时发生错误: %s%s%s", logger.Cyan, conf.SavePath, logger.Reset, logger.Red, err.Error(), logger.Reset)
	}

	urlStr := conf.Url()
	certNodes := utils.GetList(urlStr, conf.ApiToken.String())
	certUpsert := moudle.CheckNodes(certNodes, conf.Days)
	p, err := moudle.ChooseDNSProvider(conf.DNSProviderConfig)
	if err != nil {
		logger.Error.Println(fmt.Sprintf("验证 DNS 服务提供商时发生错误: %s%s%s", logger.Red, err, logger.Reset))
		os.Exit(0)
	}
	log.Printf("本次需要更新证书数量有 %d 个", len(certUpsert))
	var failedApply [][]string
	var successApply [][]string
	for _, cert := range certUpsert {
		if services.ApplyCert(cert.Domains, conf.ApplyCert.Email, conf.SavePath, p) {
			failedApply = append(failedApply, cert.Domains)
			continue
		}
		var upsertReq safeLineApi.UpsertReq
		upsertReq.Create(cert.Domains, conf.ApplyCert.Email, conf.SavePath, cert.Id, cert.Type)
		body := bytes.NewReader(upsertReq.Marshal())
		result := utils.Upsert(urlStr, conf.ApiToken.String(), body)
		if result.Msg != "" {
			logger.Error.Printf("域名 %s%s%s 证书更新失败: %s%s%s", logger.Cyan, cert.Domains, logger.Reset, logger.Red, result.Msg, logger.Reset)
			failedApply = append(failedApply, cert.Domains)
			continue
		}
		successApply = append(successApply, cert.Domains)
		logger.Success.Printf("域名 %s%s%s 证书更新成功！", logger.Cyan, cert.Domains, logger.Reset)
	}
	if len(successApply) != 0 {
		log.Printf("本次成功更新的域名证书如下: %s%s%s", logger.Cyan, successApply, logger.Reset)
	}

	if len(failedApply) != 0 {
		log.Printf("未成功更新的域名证书如下: %s%s%s", logger.Cyan, failedApply, logger.Reset)
	}
	log.Printf("本次任务执行完成")
	_ = os.RemoveAll(conf.SavePath)
}
