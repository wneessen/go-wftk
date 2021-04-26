package kratos

import (
	"fmt"
	"net/http"
	"time"
)

type KratosParms struct {
	KratosHost   string
	KratosPort   uint32
	Schema       string
	PublicPrefix string
	AdminPrefix  string
	HttpTimeOut  int
}

type kratosConfig struct {
	kratosPublicAddress string
	kratosAdminAddress  string
	sessionToken        string
	httpClient          *http.Client
}

type KratosClientObj struct {
	config *kratosConfig
}

// Create a new vault object
func NewKratosClient(p *KratosParms) *KratosClientObj {
	// Defaults
	KratosSchema := "https"
	KratosHost := "localhost"
	var KratosPort uint32 = 443
	KratosPublicPrefix := "/"
	KratosAdminPrefix := "/"
	KratosClientTimeout := http.DefaultClient.Timeout

	// Override defaults
	if p.Schema != "" {
		KratosSchema = p.Schema
	}
	if p.KratosHost != "" {
		KratosHost = p.KratosHost
	}
	if p.KratosPort > 0 {
		KratosPort = p.KratosPort
	}
	if p.PublicPrefix != "" {
		KratosPublicPrefix = p.PublicPrefix
	}
	if p.AdminPrefix != "" {
		KratosAdminPrefix = p.AdminPrefix
	}
	if p.HttpTimeOut > 0 {
		KratosClientTimeout = time.Second * time.Duration(p.HttpTimeOut)
	}

	KratosPublicAddress := fmt.Sprintf("%v://%v:%v%v",
		KratosSchema, KratosHost, KratosPort, KratosPublicPrefix)
	KratosAdminAddress := fmt.Sprintf("%v://%v:%v%v",
		KratosSchema, KratosHost, KratosPort, KratosAdminPrefix)

	// Create KratosClientObj
	k := KratosClientObj{
		config: &kratosConfig{
			kratosPublicAddress: KratosPublicAddress,
			kratosAdminAddress:  KratosAdminAddress,
			sessionToken:        "",
			httpClient: &http.Client{
				Timeout: KratosClientTimeout,
			},
		},
	}

	return &k
}

// Get the current admin base URL
func (k *KratosClientObj) GetAdminUrl() string {
	return k.config.kratosAdminAddress
}

// Get the current public base URL
func (k *KratosClientObj) GetPublicUrl() string {
	return k.config.kratosPublicAddress
}

// Get the current session token
func (k *KratosClientObj) GetSessionToken() string {
	return k.config.sessionToken
}
