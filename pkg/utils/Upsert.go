package utils

import (
	"SafelineAPI/internal/app/logger"
	"SafelineAPI/internal/app/safeLineApi"
	"io"
	"net/http"
)

func Upsert(url *safeLineApi.URL, token string, body io.Reader) safeLineApi.UpsertResp {
	header := http.Header{
		"X-SLCE-API-TOKEN": []string{token},
		"Content-Type":     []string{"application/json"},
	}
	resp, err := Request(UPSERT, url.SSLCertUrl(), body, header)
	if err != nil {
		logger.Error.Printf("更新证书时发生错误: %s%s%s", logger.Red, err, logger.Reset)
		return safeLineApi.UpsertResp{}
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	var upsertResp safeLineApi.UpsertResp
	upsertResp.Unmarshal(data)
	return upsertResp
}
