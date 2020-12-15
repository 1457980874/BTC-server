package RPC

import (
	"bytes"
	"fmt"
)

func GetBlockCount() interface{} {
	RPCjson := PrepareJSON("getblockcount")
	data, err := DoPost(RPCURL, bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return data.Result
}

func GetBestBlockHash() interface{} {
	RPCjson := PrepareJSON("getbestblockhash")
	data, err := DoPost(RPCURL, bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return data.Result
}
