//其他一些帮助方法
package tools

import (
	"net"
	"net/http"
	"reflect"
	"strings"
)

var (
	MaxDepth = 32
)

type IOtherHelper interface {
	GetIPAddress() (string, error)
	MapMerge(dst, src map[string]interface{}) map[string]interface{}
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

func (o OtherHelper) MapMerge(dst, src map[string]interface{}) map[string]interface{} {
	return o.merge(dst, src, 0)
}

func (o OtherHelper) merge(dst, src map[string]interface{}, depth int) map[string]interface{} {
	if depth > MaxDepth {
		panic("too deep!")
	}
	for key, srcVal := range src {
		if dstVal, ok := dst[key]; ok {
			srcMap, srcMapOk := o.mapify(srcVal)
			dstMap, dstMapOk := o.mapify(dstVal)
			if srcMapOk && dstMapOk {
				srcVal = o.merge(dstMap, srcMap, depth+1)
			}
		}
		dst[key] = srcVal
	}
	return dst
}

func (o OtherHelper) mapify(i interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, k := range value.MapKeys() {
			m[k.String()] = value.MapIndex(k).Interface()
		}
		return m, true
	}
	return map[string]interface{}{}, false
}
