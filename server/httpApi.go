package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "wosimple/db"
)

type ResData struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	resData *ResData = new(ResData)
)

func Count(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.Form.Get("param")
	param_map := make(map[string]string)
	err := json.Unmarshal([]byte(param), &param_map)
	if err != nil {
		log.Println("param 转型错误:", err)
	}
	Count := param_map["Count"]
	res, _ := db.ExitsKey(Count)
	if res == 0 {
		db.SetRedisString(Count, 1)
	} else {
		db.Incr(Count)
	}
	count, _ := db.GetRedisString(Count)
	resData.Data = count
	resData.renderJson(&w)
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.Form.Get("param")
	fmt.Println(param)
	resData.Data = "v0.1"
	resData.Code = "0"
	resData.Msg = "操作成功"
	resData.renderJson(&w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.Form.Get("param")
	fmt.Println(param)
	resData.Data = "{\"sessionId\":\"123456\",\"userId\":\"aa\",\"userName\":\"zhangsan\"}"
	resData.Code = "0"
	resData.Msg = "登陆成功"
	resData.renderJson(&w)
}

func (d *ResData) renderJson(w *http.ResponseWriter) {
	b, err := json.Marshal(d)
	if err != nil {
		log.Println(err)
	}
	(*w).Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	(*w).Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	(*w).Header().Set("content-type", "application/json")             //返回数据格式是json
	(*w).Write(b)
}
