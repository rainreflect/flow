package cmd

import (
	"fmt"
	"github.com/rainreflect/flow/internal/blockchain"
)

func main() {
	bn := blockchain.OpenGate()

	bn.AddBlock("this is some text")
	bn.AddBlock("and it's really usefull")

	for _, b := range bn.Blocks {
		fmt.Printf("PrevHash %x\n", b.PrevHash)
		fmt.Printf("Data in block %s\n", b.Data)
		fmt.Printf("Hash %x\n", b.Hash)
	}
}
