package main

import (
	"ConnectBTCrpc/RPC"
	"ConnectBTCrpc/RPC/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"time"
)

func main() {
	fmt.Println("go")

	request := entity.RPCrequest{
		ID:      time.Now().Unix(),
		Method:  RPC.GETBESTBLOCKHASH,
		JsonRpc: RPC.RPCVERSION,
	}
	reqBytes, err := json.Marshal(&request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("json数据", string(reqBytes))

	client := &http.Client{}
	Httprequest, err := http.NewRequest("POST", RPC.RPCURL, bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//请求头设置
	Httprequest.Header.Add("Encoding", "UTF-8")
	Httprequest.Header.Add("Content-Type", "application/json")
	Httprequest.Header.Add("Authorization", "Basic "+RPC.Base64Encode(RPC.RPCUSER+":"+RPC.RPCPASSWORD))

	respone, err := client.Do(Httprequest)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("请求结果：", respone)
	code := respone.StatusCode
	if code == 200 {
		fmt.Println("请求成功")
	} else {
		fmt.Println("请求失败！", code)
	}

	RPCresultBytes, err := ioutil.ReadAll(respone.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(RPCresultBytes))
	//RPCresult := make(map[string]interface{})
	var RPCdata entity.Data
	json.Unmarshal(RPCresultBytes, &RPCdata)
	fmt.Println(RPCdata.Result)
}
