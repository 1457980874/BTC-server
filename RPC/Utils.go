package RPC

import (
	"ConnectBTCrpc/RPC/entity"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

//准备rpc连接的json数据
func PrepareJSON(method string,params ...interface{})[]byte{
	request := entity.RPCrequest{
		ID:      time.Now().Unix(),
		Method:  method,
		JsonRpc: "",
		Params:  params,
	}
	ReqBytes,err:=json.Marshal(&request)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return ReqBytes
}

//发起rpc请求
func DoPost(url string, body io.Reader){
	/*client:=&http.Client{}
	request,err:=http.NewRequest("Post",url,body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}*/

}

//go标椎库rpc请求
func ConnectRPCserver(){
	/*rpc,err:=jsonrpc.Dial("Post",RPCURL)
	if err != nil {
		log.Fatal(err)
	}
	rpc.Call()*/
}

//Base64编码工具
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
func Base64Decode(str string ) string {
	result,_:=base64.StdEncoding.DecodeString(str)
	return string(result)
}
