package main

import (
	"filestore_server/config"
	handler2 "filestore_server/handler"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// 静态资源处理
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	fmt.Println(os.Args[0])
	http.Handle("/static/", http.FileServer(http.Dir(filepath.Join(pwd, "../../"))))
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	http.HandleFunc("/file/upload", handler2.HTTPInterceptor(handler2.UploadHandler))
	http.HandleFunc("/file/upload/suc", handler2.HTTPInterceptor(handler2.UploadSucHandler))
	http.HandleFunc("/file/meta", handler2.HTTPInterceptor(handler2.GetFileMetaHandler))
	http.HandleFunc("/file/download", handler2.HTTPInterceptor(handler2.DownloadHandler))
	http.HandleFunc("/file/update", handler2.HTTPInterceptor(handler2.FileMetaUpdateHandler))
	http.HandleFunc("/file/delete", handler2.HTTPInterceptor(handler2.FileDeleteHandler))
	http.HandleFunc("/file/query", handler2.HTTPInterceptor(handler2.FileQueryHandler))

	// 秒传接口
	http.HandleFunc("/file/fastupload", handler2.HTTPInterceptor(
		handler2.TryFastUploadHandler))

	// 分块上传接口
	http.HandleFunc("/file/mpupload/init",
		handler2.HTTPInterceptor(handler2.InitialMultipartUploadHandler))
	http.HandleFunc("/file/mpupload/uppart",
		handler2.HTTPInterceptor(handler2.UploadPartHandler))
	http.HandleFunc("/file/mpupload/complete",
		handler2.HTTPInterceptor(handler2.CompleteUploadHandler))

	// 用户相关接口
	http.HandleFunc("/", handler2.SignInHandler)
	http.HandleFunc("/user/signup", handler2.SignUpHandler)
	http.HandleFunc("/user/signin", handler2.SignInHandler)
	http.HandleFunc("/user/info", handler2.HTTPInterceptor(handler2.UserInfoHandler))

	fmt.Printf("上传服务启动中，开始监听监听[%s]...\n", config.UploadServiceHost)
	err := http.ListenAndServe(config.UploadServiceHost, nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
