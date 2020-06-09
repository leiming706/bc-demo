package main

import (
	"fmt"
	"github.com/leiming706/bc-demo/blockchain"
)

func main() {
	b := blockchain.NewBlock("","Gensis block")
	fmt.Println(b)
}
