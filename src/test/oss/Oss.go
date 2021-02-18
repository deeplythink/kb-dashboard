package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)
func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func main() {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId := "LTAI4G4gZHLSmq44tFg4wkc1"
	accessKeySecret := "v7gI4Wf1LyRVzQYW6wuZHltk4YVFy8"
	bucketName := "kb-sources"
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 创建存储空间。
	//err = client.CreateBucket(bucketName)
	//if err != nil {
	//	log.Fatal(err)
	//}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}

	//上传文件
	//err = bucket.PutObjectFromFile("Oss.go", "Oss.go")
	//if err != nil {
	//	handleError(err)
	//}

	//下载文件
	err = bucket.GetObjectToFile("Oss.go", "test.txt")
	if err != nil {
		handleError(err)
	}
}
