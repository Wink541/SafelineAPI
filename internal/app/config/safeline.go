package config

import (
	"SafelineAPI/internal/app/logger"
	"SafelineAPI/pkg/utils"
)

type SafeLine struct {
	Host     `json:"Host"`
	ApiToken `json:"ApiToken"`
}
type ApiToken string

func (apiToken ApiToken) GetApiToken() (string, string) {
	return "X-SLCE-API-TOKEN", apiToken.String()
}

func (apiToken ApiToken) String() string {
	return string(apiToken)
}

func (apiToken ApiToken) Verify() bool {
	if apiToken.String() == "" {
		logger.Warning.Printf("未设置 %sSafeLine API Token%s : 请检查配置文件中的 %sSafeLine.ApiToken%s 参数", logger.Cyan, logger.Reset, logger.Yellow, logger.Reset)
		return true
	}
	return false
}

func (safeLine SafeLine) Verify() bool {
	a := safeLine.ApiToken.Verify()
	b := safeLine.Host.Verify()
	if a || b {
		return true
	}

	NoLoginAuthTokenResp, NoLoginStatusCode, NoLoginErr := utils.AuthSafeLine(*safeLine.Host.Url())
	if NoLoginErr != nil {
		logger.Error.Printf("请求服务端时发生错误: %s%s%s", logger.Red, NoLoginErr.Error(), logger.Reset)
		return true
	}
	LoginAuthTokenResp, LoginStatusCode, LoginErr := utils.VerifyAuthToken(*safeLine.Host.Url(), safeLine.ApiToken.String())
	if LoginErr != nil {
		logger.Error.Printf("验证 %sSafeLine API Token%s 时发生错误: %s%s%s", logger.Cyan, logger.Reset, logger.Red, LoginErr.Error(), logger.Reset)
		return true
	}

	if !(NoLoginAuthTokenResp.Err == "login-required" && NoLoginStatusCode == 401) {
		logger.Warning.Printf("服务端接口 %s/open/auth/token%s 请求有误: 请检查配置文件中的 %sSafeLine.Host%s 参数", logger.Cyan, logger.Reset, logger.Yellow, logger.Reset)
		return true
	}

	if LoginAuthTokenResp.Err == "login-required" && LoginStatusCode == 401 {
		logger.Warning.Printf("%sSafeLine API Token%s 有误: 请检查后重试", logger.Cyan, logger.Reset)
		return true
	}

	logger.Success.Printf("%sSafeLine%s 相关配置检验完成!", logger.Cyan, logger.Reset)
	return false
}

func (apiToken ApiToken) VerifyCommand() bool {
	if apiToken.String() == "" {
		logger.Warning.Printf("未设置 %sSafeLine API Token%s : 请检查命令中的 %s-t%s 参数", logger.Cyan, logger.Reset, logger.Yellow, logger.Reset)
		return true
	}
	return false
}

func (safeLine SafeLine) VerifyCommand() bool {
	a := safeLine.ApiToken.VerifyCommand()
	b := safeLine.Host.VerifyCommand()
	if a || b {
		return true
	}

	NoLoginAuthTokenResp, NoLoginStatusCode, NoLoginErr := utils.AuthSafeLine(*safeLine.Host.Url())
	if NoLoginErr != nil {
		logger.Error.Printf("请求服务端时发生错误: %s%s%s", logger.Red, NoLoginErr.Error(), logger.Reset)
		return true
	}
	LoginAuthTokenResp, LoginStatusCode, LoginErr := utils.VerifyAuthToken(*safeLine.Host.Url(), safeLine.ApiToken.String())
	if LoginErr != nil {
		logger.Error.Printf("验证 %sSafeLine API Token%s 时发生错误: %s%s%s", logger.Cyan, logger.Reset, logger.Red, LoginErr.Error(), logger.Reset)
		return true
	}

	if !(NoLoginAuthTokenResp.Err == "login-required" && NoLoginStatusCode == 401) {
		logger.Warning.Printf("服务端接口 %s/open/auth/token%s 请求有误: 请检查命令中的 %s-h%s 参数", logger.Cyan, logger.Reset, logger.Yellow, logger.Reset)
		return true
	}

	if LoginAuthTokenResp.Err == "login-required" && LoginStatusCode == 401 {
		logger.Warning.Printf("%sSafeLine API Token%s 有误: 请检查后重试", logger.Cyan, logger.Reset)
		return true
	}

	logger.Success.Printf("%sSafeLine%s 相关配置检验完成!", logger.Cyan, logger.Reset)
	return false
}
