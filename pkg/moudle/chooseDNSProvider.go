package moudle

import (
	"SafelineAPI/internal/app/config"
	"errors"
	"github.com/go-acme/lego/v4/challenge"
)

func ChooseDNSProvider(config config.DNSProviderConfig) (challenge.Provider, error) {
	if config.DNSProvider == "TencentCloud" {
		return config.TencentCloud.Provider()
	} else if config.DNSProvider == "AliCloud" {
		return config.AliCloud.Provider()
	} else if config.DNSProvider == "HuaweiCloud" {
		return config.HuaweiCloud.Provider()
	} else if config.DNSProvider == "WestCN" {
		return config.WestCN.Provider()
	} else if config.DNSProvider == "RainYun" {
		return config.RainYun.Provider()
	}
	return nil, errors.New("未正确设置 DNS 服务提供商")
}
