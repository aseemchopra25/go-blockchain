package blockchain

type Block struct {
	PrevHash []byte
	Data     []byte
	Hash     []byte
	Nonce    int
}

type Blockchain struct {
	Blocks []*Block
}

func CreateBlock(prevHash []byte, data string) *Block {
	newBlock := &Block{prevHash, []byte(data), []byte{}, 0}
	pow := NewProofOfWork(newBlock)
	nonce, hash := pow.Run()
	newBlock.Hash = hash
	newBlock.Nonce = nonce
	return newBlock
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(prevBlock.Hash, data)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func InitializeBlockchain() *Blockchain {

	chain := &Blockchain{}
	newBlock := CreateBlock([]byte{}, "Genesis")
	chain.Blocks = append(chain.Blocks, newBlock)
	return chain
}
