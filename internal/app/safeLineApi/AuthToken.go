package safeLineApi

import "encoding/json"

type AuthTokenResp struct {
	Data string `json:"data"`
	Err  string `json:"err"`
	Msg  string `json:"msg"`
}

func (authResp *AuthTokenResp) Unmarshal(data []byte) {
	_ = json.Unmarshal(data, authResp)
}
