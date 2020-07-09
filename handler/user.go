package handler

import (
	dblayer "../db"
	"../util"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const passwd_salt = "&%sdf"

// SignUpHandler处理用户注册请求
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回上传html页面
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username")
		passwd := r.Form.Get("passwd")

		if len(username) < 3 || len(passwd) < 5 {
			w.Write([]byte("invalid parameter"))
			return
		}

		enc_passwd := util.Sha1([]byte(passwd + passwd_salt))
		suc := dblayer.UserSignUp(username, enc_passwd)

		if suc {
			w.Write([]byte("SUCCESS"))
		} else {
			w.Write([]byte("FAILED"))
		}
	}
}

// GenToken : 生成token
func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

// Signin : 处理用户登录请求
func SignInHandler(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("passwd")

	enc_passwd := util.Sha1([]byte(passwd + passwd_salt))

	// 1. 校验用户名及密码
	dbResp := dblayer.UserSignin(username, enc_passwd)
	if !dbResp {
		w.Write([]byte("FAILED"))
	}

}
