package blockchain

import "github.com/boltdb/bolt"

const BLOCKCHAIN =""

/*
 *区块链结构体的定义，代表的是一条区块链
 */
type BlockChain struct {
	LastHash []byte//表示区块链中最新区块的哈希，用于查找最新的区块内容
	BoltDb  *bolt.DB//区块链中操作区块数据文件的数据库操作对象
}

func NewBlockChain()BlockChain{
	//创世区块
	genesis:=CreateGenesisBlock()
	db,err:=bolt.Open("blockchain.db",0600,nil)
	if err != nil {
		panic(err.Error())
	}
	bc:=BlockChain{
		LastHash: genesis.Hash,
		BoltDb: db,
	}
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucket([]byte(""))
	})

	return bc
}
func (bc BlockChain)SaveData(data []byte){
	//1.从文件中读取最新的区块
	db:=bc.BoltDb
	db.View(func(tx *bolt.Tx) error {

		return nil
	})
	//新建一个区块
	newBlock:=NewBlock(,,data)
}
