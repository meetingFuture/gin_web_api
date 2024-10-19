package rsa

import (
	"io/ioutil"
	"path/filepath"

	"gin_web_api/apikey"
	"gin_web_api/pkg/setting"
	"github.com/dgrijalva/jwt-go"
)

func Setup() {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(readKey(setting.ServerSetting.RsaPrivateKey))
	if err != nil {
		panic(err)
	}
	apikey.RsaPrivateKey = privateKey

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(readKey(setting.ServerSetting.RsaPublicKey))
	if err != nil {
		panic(err)
	}
	apikey.RsaPublicKey = publicKey
}

func readKey(fileName string) []byte {
	filename := fileName
	// get the abs
	// which will try to find the 'filename' from current workind dir too.
	pem, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}

	// read the raw contents of the file
	data, err := ioutil.ReadFile(pem)
	if err != nil {
		panic(err)
	}

	return data
}
