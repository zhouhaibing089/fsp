package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var cert string
var key string

var proxy string

func init() {
	flag.StringVar(&cert, "cert", "", "the path to certificate")
	flag.StringVar(&key, "key", "", "the path to key file")
	flag.StringVar(&proxy, "proxy", "", "the ip address of proxy server")
}

func main() {
	flag.Parse()

	target, err := url.Parse("http://" + proxy)
	if err != nil {
		log.Fatalf("failed to parse proxy %q: %s", proxy, err)
	}

	upstream := httputil.NewSingleHostReverseProxy(target)
	http.ListenAndServeTLS(":443", cert, key, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		upstream.ServeHTTP(w, r)
	}))
}
