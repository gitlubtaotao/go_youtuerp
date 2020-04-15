//页面帮助方法
package tools

type IViewHelper interface {
}
type ViewHelper struct {
}

func (h ViewHelper) AssetsPublic(fileName string) string {
	//assetHost := conf.Configuration.AssetsHost
	//return assetHost + fileName
	return ""
}

