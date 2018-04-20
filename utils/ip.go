package utils

import (
	"net"
	"net/http"
	"strings"
	"bytes"
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

func CheckIpIsValid(ip string) bool {
	trial := net.ParseIP(ip)
	if trial.To4() == nil {
		return false
	}

	return false
}

// 判断ip 是否在ip 段内
func CheckIsInRange(ip, beginIp, endIp string) bool {

	trial := net.ParseIP(ip)
	beginIpTrial := net.ParseIP(beginIp)

	if beginIpTrial.To4() == nil {
		return false
	}

	endIpTrial := net.ParseIP(endIp)
	if endIpTrial.To4() == nil {
		return false
	}

	if bytes.Compare(trial, beginIpTrial) >= 0 && bytes.Compare(trial, endIpTrial) <= 0 {
		return true
	}

	return false
}

// 判断是否本机IP
func IsLocalIP(ip string) (bool, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return false, err
	}
	for i := range addrs {
		intf, _, err := net.ParseCIDR(addrs[i].String())
		if err != nil {
			return false, err
		}
		if net.ParseIP(ip).Equal(intf) {
			return true, nil
		}
	}
	return false, nil
}
