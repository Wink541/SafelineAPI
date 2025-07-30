package config

import (
	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/providers/dns/alidns"
	"github.com/go-acme/lego/v4/providers/dns/dode"
	"github.com/go-acme/lego/v4/providers/dns/huaweicloud"
	"github.com/go-acme/lego/v4/providers/dns/rainyun"
	"github.com/go-acme/lego/v4/providers/dns/tencentcloud"
	"github.com/go-acme/lego/v4/providers/dns/westcn"
)

func (tencent TencentCloud) Provider() (challenge.Provider, error) {
	cfg := tencentcloud.NewDefaultConfig()
	cfg.SecretID = tencent.SecretID
	cfg.SecretKey = tencent.SecretKey
	p, err := tencentcloud.NewDNSProviderConfig(cfg)
	return p, err
}

func (ali AliCloud) Provider() (challenge.Provider, error) {
	cfg := alidns.NewDefaultConfig()
	cfg.SecurityToken = ali.STSToken
	cfg.SecretKey = ali.AccessKeySecret
	cfg.RAMRole = ali.RAMRole
	cfg.APIKey = ali.AccessKeyId
	p, err := alidns.NewDNSProviderConfig(cfg)
	return p, err
}

func (huawei HuaweiCloud) Provider() (challenge.Provider, error) {
	cfg := huaweicloud.NewDefaultConfig()
	cfg.Region = huawei.Region
	cfg.AccessKeyID = huawei.AccessKeyId
	cfg.SecretAccessKey = huawei.SecretAccessKey
	p, err := huaweicloud.NewDNSProviderConfig(cfg)
	return p, err
}

func (west WestCN) Provider() (challenge.Provider, error) {
	cfg := westcn.NewDefaultConfig()
	cfg.Username = west.Username
	cfg.Password = west.Password
	p, err := westcn.NewDNSProviderConfig(cfg)
	return p, err
}

func (rain RainYun) Provider() (challenge.Provider, error) {
	cfg := rainyun.NewDefaultConfig()
	cfg.APIKey = rain.ApiKey
	p, err := rainyun.NewDNSProviderConfig(cfg)
	return p, err
}

func (Dode Dode) Provider() (challenge.Provider, error) {
	cfg := dode.NewDefaultConfig()
	cfg.Token = Dode.Token
	p, err := dode.NewDNSProviderConfig(cfg)
	return p, err
}
