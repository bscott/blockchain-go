package miner

import "github.com/hoshsadiq/tescoin/block"

const CURRENT_BLOCK_REWARD = 12.5

func Mine(blockchain *block.Blockchain, blockRewardReceiver string) *block.Block {
    lastBlock := blockchain.GetLastBlock()

    tx := block.NewTransaction("__mined", blockRewardReceiver, CURRENT_BLOCK_REWARD)
    blockchain.AddTransaction(tx)

    prevHash := lastBlock.GetHash()
    nonce := lastBlock.ProofOfWork()
    return blockchain.NewBlock(nonce, prevHash)
}
