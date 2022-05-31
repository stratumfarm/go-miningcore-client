package miningcore

type WebsocketMsg string

const (
	WsBlockFound            WebsocketMsg = "blockfound"
	WsNewChainHeight        WebsocketMsg = "newchainheight"
	WsPayment               WebsocketMsg = "payment"
	WsBlockUnlockedProgress WebsocketMsg = "blockunlockedprogress"
	WsHashrateUpdated       WebsocketMsg = "hashrateupdated"
)

type RawMessage struct {
	Type string `json:"type"`
}

type BlockMessage struct {
	PoolID      string `json:"poolId"`
	BlockHeight uint64 `json:"blockHeight"`
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
}

type BlockFoundMessage struct {
	BlockMessage
	Miner             string `json:"miner"`
	MinerExplorerLink string `json:"minerExplorerLink"`
	Source            string `json:"source"`
}

type ChainHeightMessage struct {
	BlockMessage
}

type PaymentMessage struct{}

type BlockUnlockedMessage struct {
	BlockMessage
	BlockType         string  `json:"blockType"`
	BlockHash         string  `json:"blockHash"`
	Reward            float64 `json:"reward"`
	Effort            float64 `json:"effort"`
	Miner             string  `json:"miner"`
	ExplorerLink      string  `json:"explorerLink"`
	MinerExplorerLink string  `json:"minerExplorerLink"`
}

type BlockUnlockProgressMessage struct {
	BlockMessage
	Progress float64 `json:"progress"`
	Effort   float64 `json:"effort"`
}

type HashRateUpdateMessage struct {
	PoolID   string  `json:"poolId"`
	Hashrate float64 `json:"hashrate"`
	Miner    string  `json:"miner"`
	Worker   string  `json:"worker"`
}
