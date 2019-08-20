/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package httpproxy

import (
	"crypto/tls"
	"net/http"
	"time"
)

type Proxy interface {
	Send(req *http.Request) (resp *http.Response, err error)
}

type HttpProxy struct {
	client  *http.Client
	req     []*http.Request
	elapsed map[string]time.Duration
}

//Elapsed get all request elapsed time
func (p *HttpProxy) Elapsed() map[string]time.Duration {
	return p.elapsed
}

//NewHttpProxy
func NewHttpProxy() *HttpProxy {
	return &HttpProxy{
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       1 * time.Second,
		},
		req:     nil,
		elapsed: make(map[string]time.Duration),
	}
}

//Send
func (p *HttpProxy) Send(req *http.Request) (resp *http.Response, err error) {
	defer func(s time.Time, e error) {
		if err == nil {
			p.elapsed[req.URL.String()] = time.Now().Sub(s)
		}
	}(time.Now(), err)
	// add to list
	p.req = append(p.req, req)

	// httpproxy http request
	return p.client.Do(req)
}
