package utils

import (
	"fmt"
	"net/http"
	"time"
)

func GetFakeRayHttpProxyClient() RayHttpProxyClientInterface {
	return &FakeRayHttpProxyClient{}
}

type FakeRayHttpProxyClient struct {
	client       http.Client
	httpProxyURL string
}

func (r *FakeRayHttpProxyClient) InitClient() {
	r.client = http.Client{
		Timeout: 20 * time.Millisecond,
	}
}

func (r *FakeRayHttpProxyClient) SetHostIp(hostIp string, port int32) {
	// If $port is equal to -1, use DefaultHttpProxyPort.
	if port == -1 {
		port = DefaultHttpProxyPort
	}
	r.httpProxyURL = fmt.Sprint("http://", hostIp, ":", port)
}

func (r *FakeRayHttpProxyClient) CheckHealth() error {
	// TODO: test check return error cases.
	// Always return successful.
	return nil
}
