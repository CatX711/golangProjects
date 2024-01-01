// copied from web lol


package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// Transaction represents a basic transaction
type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
}

// Block represents a block in the blockchain
type Block struct {
	Index        int
	Timestamp    string
	Transactions []*Transaction
	PrevHash     string
	Hash         string
	Nonce        int
}

// CalculateHash generates the hash for the block
func (b *Block) CalculateHash() {
	sha := sha256.New()
	sha.Write([]byte(strconv.Itoa(b.Index) + b.Timestamp + b.PrevHash + strconv.Itoa(b.Nonce)))
	for _, tx := range b.Transactions {
		sha.Write([]byte(tx.Sender + tx.Receiver + strconv.FormatFloat(tx.Amount, 'f', -1, 64)))
	}
	hash := sha.Sum(nil)
	b.Hash = hex.EncodeToString(hash)
}

// NewBlock creates a new block in the blockchain
func NewBlock(index int, transactions []*Transaction, prevHash string) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PrevHash:     prevHash,
		Nonce:        0,
	}
	block.CalculateHash()
	return block
}

// Blockchain represents the blockchain
type Blockchain struct {
	Chain []*Block
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := NewBlock(prevBlock.Index+1, transactions, prevBlock.Hash)
	bc.Chain = append(bc.Chain, newBlock)
}

func main() {
	genesisBlock := NewBlock(0, []*Transaction{}, "")
	blockchain := Blockchain{Chain: []*Block{genesisBlock}}

	for {
		fmt.Println("\nEnter transaction details (or 'exit' to quit)")
		fmt.Print("Sender: ")
		var sender string
		fmt.Scanln(&sender)

		if sender == "exit" {
			break
		}

		fmt.Print("Receiver: ")
		var receiver string
		fmt.Scanln(&receiver)

		fmt.Print("Amount: ")
		var amount float64
		fmt.Scanln(&amount)

		transaction := &Transaction{
			Sender:   sender,
			Receiver: receiver,
			Amount:   amount,
		}

		transactions := []*Transaction{transaction}
		blockchain.AddBlock(transactions)

		fmt.Println("Transaction added to the blockchain!")
	}

	// Print the blockchain
	for _, block := range blockchain.Chain {
		fmt.Printf("\nIndex: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("Transactions:")
		for _, tx := range block.Transactions {
			fmt.Printf("\tSender: %s, Receiver: %s, Amount: %f\n", tx.Sender, tx.Receiver, tx.Amount)
		}
	}
}
