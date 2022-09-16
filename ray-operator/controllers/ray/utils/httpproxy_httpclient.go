package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	DefaultHttpProxyPort int32  = 8000
	HealthCheckPath      string = "/-/healthz"
)

type RayHttpProxyClientInterface interface {
	InitClient()
	CheckHealth() error
	SetHostIp(hostIp string, port int32)
}

// GetRayHttpProxyClientFunc Used for unit tests.
var GetRayHttpProxyClientFunc = GetRayHttpProxyClient

func GetRayHttpProxyClient() RayHttpProxyClientInterface {
	return &RayHttpProxyClient{}
}

type RayHttpProxyClient struct {
	client       http.Client
	httpProxyURL string
}

func (r *RayHttpProxyClient) InitClient() {
	r.client = http.Client{
		Timeout: 20 * time.Millisecond,
	}
}

func (r *RayHttpProxyClient) SetHostIp(hostIp string, port int32) {
	// If $port is equal to -1, use DefaultHttpProxyPort.
	// if port == -1 {
	// 	port = DefaultHttpProxyPort
	// }
	r.httpProxyURL = fmt.Sprint("http://", hostIp, ":", DefaultHttpProxyPort)
}

func (r *RayHttpProxyClient) CheckHealth() error {
	req, err := http.NewRequest("GET", r.httpProxyURL+HealthCheckPath, nil)
	if err != nil {
		return err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("RayHttpProxyClient CheckHealth fail: %s %s", resp.Status, string(body))
	}

	return nil
}
