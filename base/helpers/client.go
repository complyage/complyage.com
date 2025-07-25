package helpers

import (
	"net"
	"net/http"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| GetClientIP
//||------------------------------------------------------------------------------------------------||

func GetClientIP(r *http.Request) string {
	return "2.19.144.0" // For testing purposes, return a fixed IP
	//||------------------------------------------------------------------------------------------------||
	//|| Try Forwarded Address
	//||------------------------------------------------------------------------------------------------||
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		parts := strings.Split(ip, ",")
		ip = strings.TrimSpace(parts[0])
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Use Header
	//||------------------------------------------------------------------------------------------------||
	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Get the Direct Request Address
	//||------------------------------------------------------------------------------------------------||
	if ip == "" {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		} else {
			ip = host
		}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Localhost
	//||------------------------------------------------------------------------------------------------||
	if ip == "::1" {
		return "127.0.0.1"
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Return
	//||------------------------------------------------------------------------------------------------||
	return ip
}
