package blockchain

import (
	"sync"

	"github.com/ohbyeongmin/obmcoin/db"
	"github.com/ohbyeongmin/obmcoin/utils"
)

const (
	defaultDifficulty 	int = 2
	difficultyInterval 	int = 5
	blockInterval 		int = 2
	allowedRange		int = 2
)

// blockchain 의 구조체
type blockchain struct {
	// 가장 최근 만들어진 블록의 해쉬 값이 저장 된다.
	NewestHash string `json:"newestHash"`
	// 지금 까지 만든 블록의 Height 값이 저장 된다.
	Height int 		  `json:"height"`	
	CurrentDifficulty int 	`json:"currentDifficulty"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte){
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveCheckpoint(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func (b *blockchain) recalculateDifficulty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	lastRecalculatedBlock := allBlocks[difficultyInterval - 1]
	actualTime := (newestBlock.Timestamp/60) - (lastRecalculatedBlock.Timestamp/60)
	expectedTime := difficultyInterval * blockInterval
	if actualTime <= (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime >= (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	} else {
		return b.CurrentDifficulty
	}
}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height % difficultyInterval == 0 {
		return b.recalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}


func Blockchain() *blockchain {
	if b == nil {
		// 최초 한번만 실행 됨
		once.Do(func() {
			// 빈 블록체인으로 값을 담음
			b = &blockchain{
				Height: 0,
			}
			// search for checkpoint on the db
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis")
			} else {
				// restore b from bytes
				b.restore(checkpoint)
			}	
		})
	}
	return b
}



