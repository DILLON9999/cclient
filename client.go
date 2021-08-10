package cclient

import (
	http "github.com/useflyent/fhttp"

	"time"

	utls "github.com/refraction-networking/utls"
)

func NewClient(clientHello utls.ClientHelloID, redirectOption string, timeout int32, proxyUrl ...string) (http.Client, error) {
	
	var transport http.RoundTripper = &http.Transport{
        	DisableKeepAlives: true,
    	}

	if redirectOption == "False" {
		if len(proxyUrl) > 0 && len(proxyUrl[0]) > 0 {
			return http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
				Transport: transport,
				Timeout:   time.Duration(timeout) * time.Second,
			}, nil
		} else {
			return http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
				Transport: transport,
				Timeout:   time.Duration(timeout) * time.Second,
			}, nil
		}
	} else {
		if len(proxyUrl) > 0 && len(proxyUrl[0]) > 0 {
			return http.Client{
				Transport: transport,
				Timeout:   time.Duration(timeout) * time.Second,
			}, nil
		} else {
			return http.Client{
				Transport: transport,
				Timeout:   time.Duration(timeout) * time.Second,
			}, nil
		}
	}
}
