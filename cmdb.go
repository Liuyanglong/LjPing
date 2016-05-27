package main

import (
	"encoding/json"
	"errors"
	//"fmt"
	"github.com/astaxie/beego/httplib"
	js "github.com/bitly/go-simplejson"
    "strconv"
	"io/ioutil"
    "time"
)

var cmdbUrl string = "http://machine-manage.lianjia.com/lianjia/machine/api"
var name string = "********"
var passwd string = "********"
var baseid int = 1
var queryAuth string = ""

type MyData struct {
	Version string      `json:"version"`
	Method  string      `json:"method"`
	Auth    string      `json:"auth"`
	Params  interface{} `json:"params"`
}

func getServerIPList() ( []string,error ) {
    result := make( []string,0)
    serverList,err := getServerIdFromBaseTag()    
    if err != nil {
        return result,err    
    }
    auth ,err := getAuth()
    if err != nil {
        return result,err     
    }
    
    method := "machine.getmsg"
    params := make([]map[string]string,0)
    for _,sid := range serverList {
        smsg := make(map[string]string)
        smsg["id"] = sid    
        params = append(params,smsg)
    }
    res,cerr := CurlSendCmdb(method,params,auth)
    if cerr != nil {
        return result,cerr    
    }
    
    serverMsgList,merr := res.Get("result").Map()
    if merr != nil {
        return result,merr    
    }

    for _,sm := range serverMsgList {
        smMap,ok := sm.(map[string]interface{})
        if !ok {
            continue    
        }
        ip := smMap["ip"].(string)    
        result = append(result,ip)
    }

    return result,nil
}

func getServerIdFromBaseTag() ([]string,error) {
    result := make( []string,0)
    auth,err := getAuth()
    if err != nil {
        return result,err    
    }
    
    method := "tag.getmachines"
    params := []int{baseid}
    res,cerr := CurlSendCmdb(method,params,auth)
    if cerr != nil {
        return result,cerr    
    }
    return res.Get("result").Get(strconv.Itoa(baseid)).StringArray()
}

func getAuth() (string,error) {
    if queryAuth != "" {
        return queryAuth,nil    
    }
    method := "user.login"
    auth := ""
    params := make(map[string]string)
    params["username"] = name
    params["password"] = passwd
    res, err := CurlSendCmdb(method,params,auth)
    if err != nil {
        return "",err
    }

    queryAuth,err = res.Get("result").Get("auth").String()
    return queryAuth,err
}

func CurlSendCmdb(method string, params interface{}, auth string) (*js.Json, error) {
	values := make(map[string]interface{})
	values["version"] = "1.0"
	values["method"] = method
	values["params"] = params
	values["auth"] = auth
	b, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	req := httplib.Post(cmdbUrl)
	req.Body(b)
	req.SetTimeout(60*time.Second, 60*time.Second)
	resp, err1 := req.Response()
	if err1 != nil {
		return nil, err1
	}

	status := resp.StatusCode
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err2
	}

	jsondata, err3 := js.NewJson(body)
	if err3 != nil {
		return nil, err3
	}

	if status != 200 {
		return nil, errors.New("response is not 200!")
	}

    return jsondata,nil
}
