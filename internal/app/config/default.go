package config

func (config *Config) Default() {
	a := Config{
		SafeLine: SafeLine{
			Host: Host{
				HostName: "192.168.1.4",
				Port:     "1443",
			},
			ApiToken: "xxx",
		},
		ApplyCert: ApplyCert{
			Days:     30,
			Email:    "xxx",
			SavePath: "/tmp/ssl",
			DNSProviderConfig: DNSProviderConfig{
				DNSProvider: "xxx",
				TencentCloud: TencentCloud{
					SecretID:  "xxx",
					SecretKey: "xxx",
				},
				AliCloud: AliCloud{
					AccessKeyId:     "xxx",
					AccessKeySecret: "xxx",
					RAMRole:         "xxx(可选)",
					STSToken:        "xxx(可选)",
				},
				HuaweiCloud: HuaweiCloud{
					AccessKeyId:     "xxx",
					Region:          "xxx",
					SecretAccessKey: "xxx",
				},
				WestCN: WestCN{
					Username: "xxx",
					Password: "xxx",
				},
				RainYun: RainYun{
					ApiKey: "xxx",
				},
			},
		},
	}
	a.Write("./config.json")
}
