package blockchain

type Blockchain struct {
	Blocks []*Block
}

func (b *Blockchain) AddBlock(data string) {
	//b.blocks = append(b.blocks, block)
	//TODO: add head on block to blchn struct because of this...
	prevBlock := b.Blocks[len(b.Blocks)-1]
	nBlock := CreateBlock(data, prevBlock.Hash)
	b.Blocks = append(b.Blocks, nBlock)
}

func Genesis() *Block {
	return CreateBlock("G", []byte{})
}

func OpenGate() *Blockchain {
	return &Blockchain{
		Blocks: []*Block{Genesis()},
	}
}
