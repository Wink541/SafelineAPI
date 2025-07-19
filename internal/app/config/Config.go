package config

import (
	"SafelineAPI/internal/app/logger"
	"flag"
	"github.com/go-jose/go-jose/v4/json"
	"os"
)

type Config struct {
	SafeLine  `json:"SafeLine"`
	ApplyCert `json:"ApplyCert"`
}

func (config *Config) Read(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		logger.Error.Printf("配置文件读取失败: %s%s%s", logger.Red, err, logger.Reset)
		os.Exit(0)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		logger.Error.Printf("配置文件读取失败: %s%s%s", logger.Red, err, logger.Reset)
		os.Exit(0)
	}
	config.Verify()
}
func (config *Config) Write(path string) {
	data, _ := json.MarshalIndent(config, "", "	")
	_ = os.WriteFile(path, data, 0644)
}

func (config *Config) Command() {
	var hostname, port, apiToken, save, email *string
	var days *int
	var DNSProvider *string

	hostname = flag.String("h", "172.22.222.4", "-h <hostname>")
	port = flag.String("p", "9443", "-p <port>")
	apiToken = flag.String("t", "", "-t <apiToken>")
	days = flag.Int("d", 30, "-t <days>")
	save = flag.String("s", "/tmp/ssl", "-s <save file>")
	email = flag.String("e", "", "-e <email>")
	DNSProvider = flag.String("D", "", "-D <DNS Provider> (e.g., TencentCloud, AliCloud, HuaweiCloud, WestCN, RainYun)")
	kvp := flag.String("kv", "", "-kv <key=value>,<key=value>...")
	flag.Parse()

	var KVP = make(KVPair)
	if *kvp != "" {
		KVP.Set(*kvp)
	}

	config.SafeLine = SafeLine{
		Host: Host{
			HostName: *hostname,
			Port:     *port,
		},
		ApiToken: ApiToken(*apiToken),
	}
	config.ApplyCert = ApplyCert{
		Days:     *days,
		SavePath: *save,
		Email:    *email,
		DNSProviderConfig: DNSProviderConfig{
			DNSProvider: *DNSProvider,
			TencentCloud: TencentCloud{
				SecretID:  KVP["SecretID"],
				SecretKey: KVP["SecretKey"],
			},
			AliCloud: AliCloud{
				AccessKeyId:     KVP["AccessKeyId"],
				AccessKeySecret: KVP["AccessKeySecret"],
				RAMRole:         KVP["RAMRole"],
				STSToken:        KVP["STSToken"],
			},
			HuaweiCloud: HuaweiCloud{
				AccessKeyId:     KVP["AccessKeyId"],
				Region:          KVP["Region"],
				SecretAccessKey: KVP["SecretAccessKey"],
			},
			WestCN: WestCN{
				Username: KVP["Username"],
				Password: KVP["Password"],
			},
			RainYun: RainYun{
				ApiKey: KVP["ApiKey"],
			},
		},
	}
	config.VerifyCommand()
}
