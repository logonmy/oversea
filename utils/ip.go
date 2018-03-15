package utils

import (
	"net"
	"net/http"
	"strings"
)

// IpAddress 获取当前IP地址
func GetIpAddress() (string) {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return "0.0.0.0"
	}
	for _, address := range addresses {
		if ip, ok := address.(*net.IPNet); ok &&
			!ip.IP.IsLoopback() { // 检查ip地址判断是否回环地址
			if ip.IP.To4() != nil {
				ipAddress := ip.IP.String()
				return ipAddress
			}
		}
	}
	return "0.0.0.0"
}

// RemoteIpAddress 获取请求IP地址
func GetRemoteIpAddress(req *http.Request) string {
	ip := req.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip, _, _ = net.SplitHostPort(req.RemoteAddr)
	} else if ips := strings.Split(ip, ","); len(ips) != 0 {
		ip = ips[0]
	}
	return strings.TrimSpace(ip)
}
