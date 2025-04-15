package utils

import (
	"SafelineAPI/internal/app/logger"
	"SafelineAPI/internal/app/safeLineApi"
	"io"
	"net/http"
	"os"
)

func GetList(url *safeLineApi.URL, token string) safeLineApi.Nodes {
	header := http.Header{
		"X-SLCE-API-TOKEN": []string{token},
	}
	resp, err := Request(LIST, url.SSLCertUrl(), nil, header)
	if err != nil {
		logger.Error.Printf("请求接口 %s/api/open/cert%s 时发生错误: %s%s%s", logger.Cyan, logger.Reset, logger.Red, err, logger.Reset)
		os.Exit(0)
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	var listResp safeLineApi.ListResp
	listResp.Unmarshal(data)
	return listResp.Data.Nodes
}
