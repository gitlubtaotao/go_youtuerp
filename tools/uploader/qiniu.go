package uploader

import (
	"context"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"mime/multipart"
	"time"
	"youtuerp/conf"
)

type IQiNiuUploader interface {
	FilerUpload(file multipart.File, fileHeader *multipart.FileHeader) (url string, key string, err error)
	PublicReadUrl(key string) string
	PrivateReadURL(key string) string
	DeleteFile(key string) error
}

//自定义返回值结构体
type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

type QiNiuUploader struct {
	Bucket        string
	FileServer    string
	Zone          *storage.Zone
	UseHTTPS      bool
	UseCdnDomains bool
	Expires       uint64
}

func NewQiNiuUploaderDefault() QiNiuUploader {
	zone := &storage.ZoneHuanan
	config := conf.Configuration
	return NewQiNiuUploader(config.QiniuBucket,
		config.QiniuFileServer, zone, false,
		false, 7200)
}
func NewQiNiuUploader(bucket string, fileServer string, zone *storage.Zone, useHttps bool,
	useCdnDomains bool, expires uint64) QiNiuUploader {
	return QiNiuUploader{
		Bucket:        bucket,
		FileServer:    fileServer,
		Zone:          zone,
		UseHTTPS:      useHttps,
		UseCdnDomains: useCdnDomains,
		Expires:       expires,
	}
}

func (d *QiNiuUploader) Upload(file multipart.File, fileHeader *multipart.FileHeader) (string, string, error) {
	return d.FilerUpload(file, fileHeader)
}

func (d *QiNiuUploader) FilerUpload(file multipart.File, fileHeader *multipart.FileHeader) (url string, key string, err error) {
	var (
		ret  MyPutRet
		size int64
	)
	if statInterface, ok := file.(Stat); ok {
		fileInfo, _ := statInterface.Stat()
		size = fileInfo.Size()
	}
	if sizeInterface, ok := file.(Size); ok {
		size = sizeInterface.Size()
	}
	putPolicy := storage.PutPolicy{
		Scope:   d.Bucket,
		Expires: d.Expires,
	}
	mac := d.newMac()
	token := putPolicy.UploadToken(mac)
	cfg := d.storageConfig()
	uploader := storage.NewFormUploader(&cfg)
	putExtra := storage.PutExtra{}
	err = uploader.PutWithoutKey(context.Background(), &ret, token, file, size, &putExtra)
	if err != nil {
		return
	}
	
	return d.FileServer + ret.Key, ret.Key, nil
}

/*
对于私有空间，首先需要按照公开空间的文件访问方式构建对应的公开空间访问链接，然后再对这个链接进行私有授权签名。
*/
func (d *QiNiuUploader) PrivateReadURL(key string) string {
	mac := d.newMac()
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	privateAccessURL := storage.MakePrivateURL(mac, d.FileServer, key, deadline)
	return privateAccessURL
}



func (d *QiNiuUploader) PublicReadUrl(key string) string {
	publicAccessURL := storage.MakePublicURL(d.FileServer, key)
	return publicAccessURL
}

//删除文件
func (d *QiNiuUploader) DeleteFile(key string) error {
	return d.BucketManager().Delete(d.Bucket, key)
}

func (d *QiNiuUploader) BucketManager() *storage.BucketManager {
	mac := d.newMac()
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS:      d.UseHTTPS,
		Zone:          d.Zone,
		UseCdnDomains: d.UseCdnDomains,
	}
	return storage.NewBucketManager(mac, &cfg)
}

func (d *QiNiuUploader) newMac() (mac *qbox.Mac) {
	return qbox.NewMac(conf.Configuration.QiniuAccessKey, conf.Configuration.QiniuSecretKey)
}

func (d *QiNiuUploader) storageConfig() storage.Config {
	return storage.Config{
		Zone:          d.Zone,
		UseHTTPS:      d.UseHTTPS,
		UseCdnDomains: d.UseCdnDomains,
	}
}
