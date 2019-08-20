/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package main

import (
	"github.com/tkstorm/go-design/structural/proxy/httpproxy"
	"log"
	"net/http"
	"sync"
)

func main() {
	reqUrl := []string{
		"https://tkstorm.com",
		"https://qq.com",
		"https://toutiao.com",
		"https://baidu.com",
		"https://google.com",
		"https://sina.com",
		"https://jinshan.com",
		"https://alibaba.com",
	}

	proxy := httpproxy.NewHttpProxy()
	send := func(url string, wg *sync.WaitGroup) {
		defer wg.Done()
		req, _ := http.NewRequest("GET", url, nil)
		resp, err := proxy.Send(req)
		if err != nil {
			log.Println(err)
			return
		}
		resp.Body.Close()
		//fmt.Printf("%v: %d %s\n", url, resp.StatusCode, resp.Proto)
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(reqUrl))
	for _, url := range reqUrl {
		go send(url, wg)
	}
	wg.Wait()

	// show consume time
	for k, t := range proxy.Elapsed() {
		log.Printf("Request OK: %s, elapse %s\n", k, t.String())
	}
}
