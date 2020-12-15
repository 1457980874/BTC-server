package main

import (
	"ConnectBTCrpc/RPC"
	"fmt"
)

func main() {
	count := RPC.GetBlockCount()
	fmt.Println("当前节点的区块数为：", count)
	bestBlockHash := RPC.GetBestBlockHash()
	fmt.Println("当前节点最新区块的hash是：", bestBlockHash)
}
