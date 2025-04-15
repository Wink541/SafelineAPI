package safeLineApi

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type UpsertReq struct {
	Acme struct {
		Domains []string `json:"domains"`
		Email   string   `json:"email"`
	} `json:"acme"`
	Id     int `json:"id"`
	Manual struct {
		Crt string `json:"crt"`
		Key string `json:"key"`
	} `json:"manual"`
	Type int `json:"type"`
}

func (upsertReq *UpsertReq) Create(domains []string, email, dir string, id, Type int) {
	certificate, _ := os.ReadFile(filepath.Join(dir, domains[0]+".crt"))
	privateKey, _ := os.ReadFile(filepath.Join(dir, domains[0]+".key"))
	upsertReq.Acme.Domains = domains
	upsertReq.Acme.Email = email
	upsertReq.Manual.Crt = string(certificate)
	upsertReq.Manual.Key = string(privateKey)
	upsertReq.Id = id
	upsertReq.Type = Type
}

func (upsertReq *UpsertReq) Marshal() []byte {
	data, _ := json.Marshal(upsertReq)
	return data
}

type UpsertResp struct {
	Data int         `json:"data"`
	Err  interface{} `json:"err"`
	Msg  string      `json:"msg"`
}

func (upsertResp *UpsertResp) Unmarshal(data []byte) {
	_ = json.Unmarshal(data, &upsertResp)
}

type ListResp struct {
	Data struct {
		Nodes `json:"nodes"`
		Total int `json:"total"`
	} `json:"data"`
	Err string `json:"err"`
	Msg string `json:"msg"`
}

type Nodes []struct {
	Id            int       `json:"id"`
	Domains       []string  `json:"domains"`
	Issuer        string    `json:"issuer"`
	SelfSignature bool      `json:"self_signature"`
	Trusted       bool      `json:"trusted"`
	Revoked       bool      `json:"revoked"`
	Expired       bool      `json:"expired"`
	Type          int       `json:"type"`
	AcmeMessage   string    `json:"acme_message"`
	ValidBefore   time.Time `json:"valid_before"`
	RelatedSites  []string  `json:"related_sites"`
}

func (listResp *ListResp) Unmarshal(data []byte) {
	_ = json.Unmarshal(data, &listResp)
}
