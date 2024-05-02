package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe("127.0.0.1:80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("enter root handler\n")
	h := w.Header()
	h.Set("VERSION", os.Getenv("VERSION"))
	// fmt.Printf("set response header, key:VERSION, value:%s\n", os.Getenv("VERSION"))
	for k, v := range r.Header {
		h.Set(k, fmt.Sprintf("%s", v))
		// fmt.Printf("set response header, key:%s, value:%s\n", k, v)
	}
	log.Printf("clientIp:%s, responseStatus:200\n", getClientIp(r))
	io.WriteString(w, "ok\n")
}

func getClientIp(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	var arr []string
	if ip == "" {
		arr = strings.Split(r.RemoteAddr, ":")
	} else {
		arr = strings.Split(ip, ",")
	}
	if len(arr) > 0 {
		ip = arr[0]
	}
	return ip
}
