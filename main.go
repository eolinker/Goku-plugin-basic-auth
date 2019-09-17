package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	goku_plugin "github.com/eolinker/goku-plugin"
)

type basicAuthNode struct {
	UserName       string `json:"userName"`
	Password       string `json:"password"`
	HideCredential bool   `json:"hideCredential"`
	Remark         string `json:"remark"`
}

type basicAuthConf []basicAuthNode

var builder = new(gokuBaseAuthPluginFactory)

func Builder() goku_plugin.PluginFactory {
	return builder
}

type gokuBaseAuthPluginFactory struct {
}

func (f *gokuBaseAuthPluginFactory) Create(config string, clusterName string, updateTag string, strategyId string, apiId int) (*goku_plugin.PluginObj, error) {
	var conf basicAuthConf

	if err := json.Unmarshal([]byte(config), &conf); err != nil {
		//解析配置信息失败
		return nil, fmt.Errorf("[basic_auth] Parsing plugin config error:%s", err.Error())
	}
	p := &gokuBasicAuth{
		conf: conf,
	}
	return &goku_plugin.PluginObj{
		Access: p,
	}, nil
}

type gokuBasicAuth struct {
	conf basicAuthConf
}

func (p *gokuBasicAuth) Access(ctx goku_plugin.ContextAccess) (isContinue bool, e error) {
	if ctx.Request().GetHeader("Authorization") == "" && ctx.Request().GetHeader("Proxy-Authorization") == "" {
		//缺basicAuth认证信息
		ctx.AddHeader("WWW-Authenticate", "Basic Restricted")
		ctx.SetStatus(401, "401")
		ctx.SetBody([]byte("[basic_auth] Can not find the header named Authorization or named Proxy-Authorization."))
		return false, nil
	}

	headerName := "Proxy-Authorization"
	givenUserName, givenPassword, err := retrieveCredentials(ctx.Request().GetHeader(headerName))
	if err != nil {
		headerName = "Authorization"
		givenUserName, givenPassword, err = retrieveCredentials(ctx.Request().GetHeader(headerName))
		if err != nil {
			//解析认证信息失败
			ctx.SetStatus(403, "403")

			ctx.SetBody([]byte(err.Error()))
			return false, err
		}
	}
	for _, node := range p.conf {
		if givenUserName == node.UserName && givenPassword == node.Password {
			if node.HideCredential {
				ctx.Proxy().DelHeader(headerName)
			}
			return true, nil
		}
	}

	//认证失败
	ctx.SetStatus(403, "403")
	ctx.SetBody([]byte("[basic_auth] Invalid basic authentication"))
	return false, nil
}

// 获取basicAuth认证信息
func retrieveCredentials(authInfo string) (string, string, error) {

	if authInfo != "" {
		const basic = "basic"
		l := len(basic)

		if len(authInfo) > l+1 && strings.ToLower(authInfo[:l]) == basic {
			b, err := base64.StdEncoding.DecodeString(authInfo[l+1:])
			if err != nil {
				return "", "", err
			}
			cred := string(b)
			for i := 0; i < len(cred); i++ {
				if cred[i] == ':' {
					return cred[:i], cred[i+1:], nil
				}
			}
			return "", "", errors.New("[basic_auth] header has unrecognized format")
		} else {
			return "", "", errors.New("[basic_auth] header has unrecognized format")
		}
	} else {
		return "", "", errors.New("[basic_auth] authorization required")
	}
}
