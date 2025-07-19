package utils

import (
	"SafelineAPI/internal/app/safeLineApi"
	"io"
	"net/http"
)

func AuthSafeLine(url safeLineApi.URL) (safeLineApi.AuthTokenResp, int, error) {
	resp, err := Request(safeLineApi.GetTOKEN, url.AuthTokenUrl(), nil, nil)
	if err != nil {
		return safeLineApi.AuthTokenResp{}, 0, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	var authTokenResp safeLineApi.AuthTokenResp
	authTokenResp.Unmarshal(data)
	return authTokenResp, resp.StatusCode, nil
}

func VerifyAuthToken(url safeLineApi.URL, token string) (safeLineApi.AuthTokenResp, int, error) {
	header := http.Header{
		"X-SLCE-API-TOKEN": []string{token},
	}

	resp, err := Request(safeLineApi.GetTOKEN, url.AuthTokenUrl(), nil, header)
	if err != nil {
		return safeLineApi.AuthTokenResp{}, 0, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	var authTokenResp safeLineApi.AuthTokenResp
	authTokenResp.Unmarshal(data)
	return authTokenResp, resp.StatusCode, nil
}
