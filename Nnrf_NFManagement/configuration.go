/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFManagement

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/pexip/go-openssl"
)

type Configuration struct {
	url           string
	basePath      string
	host          string
	defaultHeader map[string]string
	userAgent     string
	httpClient    *http.Client
	tlsCtx        *openssl.Ctx
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		basePath:      "https://example.com/nnrf-nfm/v1",
		url:           "{apiRoot}/nnrf-nfm/v1",
		defaultHeader: make(map[string]string),
		userAgent:     "OpenAPI-Generator/1.0.0/go",
	}

	var err error
	cfg.tlsCtx, err = openssl.NewCtxFromFiles("cert/nrf.pem", "cert/nrf.key")
	if err != nil {
		fmt.Println("could not set openssl ctx")
		return nil
	}

	cfg.tlsCtx.SetVerify(openssl.VerifyNone, nil)

	//if err := cfg.tlsCtx.SetNextProtos([]string{"http/1.1"}); err != nil {
	//	fmt.Println("Failed to set Next Protos (ALPN)")
	//	return nil
	//}

	// Custom dial function to use OpenSSL for TLS connections
	dialTLS := func(network, addr string) (net.Conn, error) {
		fmt.Println(addr)
		fmt.Println("\n")
		fmt.Println(network)
		fmt.Println("\n")
		cfg.tlsCtx.SetVerify(openssl.VerifyNone, nil)
		conn, err := openssl.Dial(network, addr, cfg.tlsCtx, 0)
		if err != nil {
			fmt.Printf("Failed to establish TLS connection: %v\n", err)
			return nil, nil
		}
		return conn, nil
	}

	// Create a custom transport using the custom dial function
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // Skip certificate verification
		},
		TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
		DialTLS:      dialTLS,
	}

	cfg.httpClient = &http.Client{Transport: tr}

	return cfg
}

func (c *Configuration) SetBasePath(apiRoot string) {
	url := c.url

	// Replace apiRoot
	url = strings.Replace(url, "{"+"apiRoot"+"}", apiRoot, -1)

	c.basePath = url
}

func (c *Configuration) SetBasePathNoGroup(basePath string) {
	c.basePath = basePath
}

func (c *Configuration) BasePath() string {
	return c.basePath
}

func (c *Configuration) Host() string {
	return c.host
}

func (c *Configuration) SetHost(host string) {
	c.host = host
}

func (c *Configuration) UserAgent() string {
	return c.userAgent
}

func (c *Configuration) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Configuration) DefaultHeader() map[string]string {
	return c.defaultHeader
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.defaultHeader[key] = value
}

func (c *Configuration) HTTPClient() *http.Client {
	return c.httpClient
}
