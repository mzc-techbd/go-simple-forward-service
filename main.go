package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

var (
	ForwardingServiceUrl string
	Port                 string
)

func main() {
	log.Printf("[Forward demo service up]")

	ForwardingServiceUrl = os.Getenv("FORWARDING_SERVICE_URL")
	if ForwardingServiceUrl == "" {
		panic("please set FORWARDING_SERVICE_URL environment")
	}

	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}

	log.Printf("FORWARDING_SERVICE_URL env is %s", ForwardingServiceUrl)

	target, _ := url.Parse(ForwardingServiceUrl)
	reverseProxy := httputil.NewSingleHostReverseProxy(target)

	reverseProxy.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		r.URL.Host = target.Host
		r.URL.Scheme = target.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = target.Host

		reverseProxy.ServeHTTP(w, r)
	})

	http.ListenAndServe(fmt.Sprintf(":%s", Port), nil)
}
