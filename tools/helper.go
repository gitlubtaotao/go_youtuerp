package tools

import "youtuerp/conf"

type IHelper interface {
}
type Helper struct {
}

func (h Helper) AssetsPublic(fileName string) string {
	assetHost := conf.Configuration.AssetsHost
	return assetHost + fileName
}
