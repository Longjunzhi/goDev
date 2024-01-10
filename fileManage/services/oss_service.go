package services

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"time"
)

func OssUpload(localFileName string, fileName string) (OssFileName string, err error) {
	endpoint := os.Getenv("OSS_END_POINT")
	accessKeyID := os.Getenv("OSS_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("OSS_ACCESS_KEY_SECRET")
	ossBucket := os.Getenv("OSS_BUCKET")
	OssFileName = ""

	// yourBucketName填写存储空间名称。
	// yourObjectName填写Object完整路径，完整路径不包含Bucket名称。
	// yourLocalFileName填写本地文件的完整路径。
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
		return OssFileName, err
	}

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret, oss.SetCredentialsProvider(&provider))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
		return OssFileName, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(ossBucket)
	if err != nil {
		fmt.Println(err)
		return OssFileName, err
	}
	// 上传文件。
	fileName = time.Now().Format("2006-01-02") + "/" + fileName
	err = bucket.PutObjectFromFile(fileName, localFileName)
	if err != nil {
		fmt.Println(err)
		return OssFileName, err
	}
	OssFileName = endpoint[:8] + ossBucket + "." + endpoint[8:] + "/" + fileName
	return OssFileName, nil
}
