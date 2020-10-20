package blockchain

import "time"

type Block struct {
	Height		int64//区块的高度
	TimeStamp	int64//区块产生的时间戳
	PrevHash	[]byte//前一个区块的hash
	Data		[]byte//数据字段
	Hash		[]byte//当前区块的hash值
	Version		string//版本号
}

func NewBlock(height int64, prevHash []byte, data []byte)Block  {

	block :=Block{
		Height:			height+1,
		TimeStamp:		time.Now().Unix(),
		PrevHash:		prevHash,
		Data:			data,
		Version:		"0x01",
	}
	//block.Hash =
	return block
}

func CreatGenesisBlock() Block {
	blockGenesis:=NewBlock(0, []byte{0,0,0,0,0,0,0,0,0,0,0,},nil)
	return blockGenesis
}
