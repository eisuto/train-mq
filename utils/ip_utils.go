package utils

import "net/http"

func GetClientIp(r *http.Request) string {
	clientIP := r.RemoteAddr
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		clientIP = forwarded
	}
	return clientIP
}
