package config

type DNSProviderConfig struct {
	DNSProvider  string `json:"DNSProvider"`
	TencentCloud `json:"TencentCloud,omitempty"`
	AliCloud     `json:"AliCloud,omitempty"`
	HuaweiCloud  `json:"HuaweiCloud,omitempty"`
	WestCN       `json:"WestCN,omitempty"`
	RainYun      `json:"RainYun,omitempty"`
	Dode         `json:"Dode,omitempty"`
}

type TencentCloud struct {
	SecretID  string `json:"SecretId,omitempty"`
	SecretKey string `json:"SecretKey,omitempty"`
}

type AliCloud struct {
	AccessKeyId     string `json:"AccessKeyId,omitempty"`
	AccessKeySecret string `json:"AccessKeySecret,omitempty"`
	RAMRole         string `json:"RAMRole,omitempty"`
	STSToken        string `json:"STSToken,omitempty"`
}

type HuaweiCloud struct {
	AccessKeyId     string `json:"AccessKeyId,omitempty"`
	Region          string `json:"Region,omitempty"`
	SecretAccessKey string `json:"SecretAccessKey,omitempty"`
}

type WestCN struct {
	Username string `json:"Username,omitempty"`
	Password string `json:"Password,omitempty"`
}

type RainYun struct {
	ApiKey string `json:"ApiKey,omitempty"`
}

type Dode struct {
	Token string `json:"Token,omitempty"`
}
