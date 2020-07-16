package cos

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func NewClient() *cos.Client {
	// 将 examplebucket-1250000000 和 COS_REGION修改为真实的信息
	u, _ := url.Parse("https://examplebucket-1250000000.cos.COS_REGION.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	// 1.永久密钥
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "COS_SECRETID",
			SecretKey: "COS_SECRETKEY",
		},
	})

	_, err := client.Bucket.Put(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	}

	return client
}

func QueryBucketList()  {
	c := NewClient()

	s, _, err := c.Service.Get(context.Background())
	if err != nil {
		panic(err)
	}

	for _, b := range s.Buckets {
		fmt.Printf("%#v\n", b)
	}
}

func UploadFile()  {
	c := NewClient()
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	name := "test/objectPut.go"
	// 1.通过字符串上传对象
	f := strings.NewReader("test")

	_, err := c.Object.Put(context.Background(), name, f, nil)
	if err != nil {
		panic(err)
	}
	// 2.通过本地文件上传对象
	_, err = c.Object.PutFromFile(context.Background(), name, "../test", nil)
	if err != nil {
		panic(err)
	}
}

func QueryFileList()  {
	c := NewClient()

	opt := &cos.BucketGetOptions{
		Prefix:  "test",
		MaxKeys: 3,
	}
	v, _, err := c.Bucket.Get(context.Background(), opt)
	if err != nil {
		panic(err)
	}

	for _, c := range v.Contents {
		fmt.Printf("%s, %d\n", c.Key, c.Size)
	}
}

func DownloadFile()  {
	c := NewClient()
	// 1.通过响应体获取对象
	name := "test/objectPut.go"
	resp, err := c.Object.Get(context.Background(), name, nil)
	if err != nil {
		panic(err)
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s\n", string(bs))
	// 2.获取对象到本地文件
	_, err = c.Object.GetToFile(context.Background(), name, "exampleobject", nil)
	if err != nil {
		panic(err)
	}
}

func DelFile()  {
	c := NewClient()
	name := "test/objectPut.go"
	_, err := c.Object.Delete(context.Background(), name)
	if err != nil {
		panic(err)
	}
}