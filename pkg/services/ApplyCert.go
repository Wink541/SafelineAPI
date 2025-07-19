package services

import (
	"SafelineAPI/internal/app/logger"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	"os"
	"path/filepath"
)

type MyUser struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

func (u *MyUser) GetEmail() string {
	return u.Email
}
func (u *MyUser) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *MyUser) GetPrivateKey() crypto.PrivateKey {
	return u.key
}

func ApplyCert(domains []string, email, dir string, provider challenge.Provider) bool {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		logger.Error.Printf("申请 %s%s%s 证书时发生错误: %s%s%s", logger.Cyan, domains, logger.Reset, logger.Red, err, logger.Reset)
		return true
	}
	myUser := MyUser{
		Email: email,
		key:   privateKey,
	}
	config := lego.NewConfig(&myUser)
	config.Certificate.KeyType = certcrypto.RSA2048
	client, err := lego.NewClient(config)
	if err != nil {
		logger.Error.Printf("申请 %s%s%s 证书时发生错误: %s%s%s", logger.Cyan, domains, logger.Reset, logger.Red, err, logger.Reset)
		return true
	}
	err = client.Challenge.SetDNS01Provider(provider)
	if err != nil {
		logger.Error.Printf("申请 %s%s%s 证书时发生错误: %s%s%s", logger.Cyan, domains, logger.Reset, logger.Red, err, logger.Reset)
		return true
	}

	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		logger.Error.Printf("申请 %s%s%s 证书时发生错误: %s%s%s", logger.Cyan, domains, logger.Reset, logger.Red, err, logger.Reset)
		return true
	}
	myUser.Registration = reg
	request := certificate.ObtainRequest{
		Domains: domains,
		Bundle:  true,
	}
	certificates, err := client.Certificate.Obtain(request)
	if err != nil {
		logger.Error.Printf("申请 %s%s%s 证书时发生错误: %s%s%s", logger.Cyan, domains, logger.Reset, logger.Red, err, logger.Reset)
		return true
	}
	err = os.WriteFile(filepath.Join(dir, domains[0]+".crt"), certificates.Certificate, os.ModePerm)
	if err != nil {
		logger.Error.Printf("保存 %s%s%s 证书时发生错误: %s%s%s", logger.Cyan, domains, logger.Reset, logger.Red, err, logger.Reset)
		return true
	}
	err = os.WriteFile(filepath.Join(dir, domains[0]+".key"), certificates.PrivateKey, os.ModePerm)
	if err != nil {
		logger.Error.Printf("保存 %s%s%s 证书密钥时发生错误: %s%s%s", logger.Cyan, domains, logger.Reset, logger.Red, err, logger.Reset)
		return true
	}
	return false
}
