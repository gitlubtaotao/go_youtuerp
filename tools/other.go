//其他一些帮助方法
package tools

import (
	"net"
	"net/http"
	"strings"
)

type IOtherHelper interface {
	GetIPAddress() (string, error)
}

type OtherHelper struct {
}

//获取访问的iP真实地址
func (o OtherHelper) GetIPAddress(r *http.Request) (string, error) {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" {
			return ip, nil
		}
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip, nil
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip, nil
	}
	return "", nil
}
