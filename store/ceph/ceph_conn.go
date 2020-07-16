package ceph

import (
	"gopkg.in/amz.v1/aws"
	"gopkg.in/amz.v1/s3"
)
var cephConn *s3.S3

func GetCephConnection() *s3.S3 {
	if cephConn != nil {
		return  cephConn
	}
	//1. 初始化ceph的一些信息

	auth := aws.Auth{
		AccessKey: "test",
		SecretKey: "secret",
	}

	curRegion := aws.Region{
		Name:                 "default",
		EC2Endpoint:          "EC2Endpoint",
		S3Endpoint:           "S3Endpoint",
		S3BucketEndpoint:     "",
		S3LocationConstraint: false,
		S3LowercaseBucket:    false,
		Sign:                 aws.SignV2,
	}

	// 2. 创建S3类型的连接
	return s3.New(auth, curRegion)
}

// GetCephBucket : 获取指定的bucket对象
func GetCephBucket(bucket string) *s3.Bucket  {
	conn := GetCephConnection()
	return conn.Bucket(bucket)
}

func PutObject(bucket string, path string, data []byte) error  {
	return GetCephBucket(bucket).Put(path, data, "octet-stream", s3.PublicRead)
}