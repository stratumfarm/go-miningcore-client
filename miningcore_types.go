package miningcore

type Meta struct {
	PageCount           int64    `json:"pageCount"`
	Success             bool     `json:"success"`
	ResponseMessageType int64    `json:"responseMessageType"`
	ResponseMessageId   string   `json:"responseMessageId"`
	ResponseMessageArgs []string `json:"responseMessageArgs"`
}

type PoolInfo struct {
	ID                      string                          `json:"id"`
	Coin                    *ApiCoinConfig                  `json:"coin"`
	Ports                   map[string]PoolEndpoint         `json:"ports"`
	PaymentProcessing       *ApiPoolPaymentProcessingConfig `json:"paymentProcessing"`
	ShareBasedBanning       *PoolShareBasedBanningConfig    `json:"shareBasedBanning"`
	ClientConnectionTimeout int32                           `json:"clientConnectionTimeout"`
	JobRebroadcastTimeout   int32                           `json:"jobRebroadcastTimeout"`
	BlockRefreshInterval    int32                           `json:"blockRefreshInterval"`
	PoolFeePercent          float64                         `json:"poolFeePercent"`
	Address                 string                          `json:"address"`
	AddressInfoLink         string                          `json:"addressInfoLink"`
	PoolStats               *PoolStats                      `json:"poolStats"`
	NetworkStats            *BlockchainStats                `json:"networkStats"`
	TopMiners               []*MinerPerformanceStats        `json:"topMiners"`
	TotalPaid               float64                         `json:"totalPaid"`
	TotalBlocks             int32                           `json:"totalBlocks"`
	LastPoolBlockTime       string                          `json:"lastPoolBlockTime"`
	APIEndpoint             string                          `json:"apiEndpoint"`
}

type ApiCoinConfig struct {
	Type          string `json:"type"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	Website       string `json:"website"`
	Family        string `json:"family"`
	Algorithm     string `json:"algorithm"`
	Twitter       string `json:"twitter"`
	Discord       string `json:"discord"`
	Telegram      string `json:"telegram"`
	CanonicalName string `json:"canonicalName"`
}

type PoolEndpoint struct {
	ListenAddress    string                  `json:"listenAddress"`
	Name             string                  `json:"name"`
	Difficulty       float64                 `json:"difficulty"`
	TCPProxyProtocol *TCPProxyProtocolConfig `json:"tcpProxyProtocol"`
	VarDiff          *VarDiffConfig          `json:"varDiff"`
	TLS              bool                    `json:"tls"`
	TLSAuto          bool                    `json:"tlsAuto"`
	TLSPfxFile       string                  `json:"tlsPfxFile"`
	TLSPfxPassword   string                  `json:"tlsPfxPassword"`
}

type TCPProxyProtocolConfig struct {
	Enable         bool     `json:"enable"`
	Mandatory      bool     `json:"mandatory"`
	ProxyAddresses []string `json:"proxyAddresses"`
}

type VarDiffConfig struct {
	MinDiff         float64 `json:"minDiff"`
	MaxDiff         float64 `json:"maxDiff"`
	MaxDelta        float64 `json:"maxDelta"`
	TargetTime      float64 `json:"targetTime"`
	RetargetTime    float64 `json:"retargetTime"`
	VariancePercent float64 `json:"variancePercent"`
}

type ApiPoolPaymentProcessingConfig struct {
	Enabled        bool                   `json:"enabled"`
	MinimumPayment float64                `json:"minimumPayment"`
	PayoutScheme   string                 `json:"payoutScheme"`
	Extra          map[string]interface{} `json:"extra"`
}

type PoolShareBasedBanningConfig struct {
	Enabeld         bool    `json:"enabled"`
	CheckThresghold int32   `json:"checkThreshold"`
	InvalidPercent  float64 `json:"invalidPercent"`
	Time            int32   `json:"time"`
}

type PoolStats struct {
	LastPoolBlockTime string `json:"lastPoolBlockTime"`
	ConnectedMiners   int32  `json:"connectedMiners"`
	PoolHashrate      int64  `json:"poolHashrate"`
	SharesPerSecond   int32  `json:"sharesPerSecond"`
}

type BlockchainStats struct {
	NetworkType          string  `json:"networkType"`
	NetworkHashrate      float64 `json:"networkHashrate"`
	NetworkDifficulty    float64 `json:"networkDifficulty"`
	NextNetworkTarget    string  `json:"nextNetworkTarget"`
	NextNetworkBits      string  `json:"nextNetworkBits"`
	LastNetworkBlockTime string  `json:"lastNetworkBlockTime"`
	BlockHeight          int64   `json:"blockHeight"`
	ConnectedPeers       int32   `json:"connectedPeers"`
	RewardType           string  `json:"rewardType"`
}

type MinerPerformanceStats struct {
	Miner           string  `json:"miner"`
	Hashrate        float64 `json:"hashrate"`
	SharesPerSecond float64 `json:"sharesPerSecond"`
}

type Block struct {
	PoolID                      string  `json:"poolId"`
	BlockHeight                 int64   `json:"blockHeight"`
	NetworkDifficulty           float64 `json:"networkDifficulty"`
	Status                      string  `json:"status"`
	Type                        string  `json:"type"`
	ConfirmationProgress        float64 `json:"confirmationProgress"`
	Effort                      float64 `json:"effort"`
	TransactionConfirmationData string  `json:"transactionConfirmationData"`
	Reward                      float64 `json:"reward"`
	InfoLink                    string  `json:"infoLink"`
	Hash                        string  `json:"hash"`
	Miner                       string  `json:"miner"`
	Source                      string  `json:"source"`
	Created                     string  `json:"created"`
}

type BlocksRes struct {
	Meta
	Result []*Block `json:"result"`
}

type Payment struct {
	Coin                        string  `json:"coin,omitempty"`
	Address                     string  `json:"address,omitempty"`
	AddressInfoLink             string  `json:"addressInfoLink,omitempty"`
	Amount                      float64 `json:"amount,omitempty"`
	TransactionConfirmationData string  `json:"transactionConfirmationData,omitempty"`
	TransactionInfoLink         string  `json:"transactionInfoLink,omitempty"`
	Created                     string  `json:"created,omitempty"`
}

type MinerStats struct {
	PendingShares      int64          `json:"pendingShares"`
	PendingBalance     float64        `json:"pendingBalance"`
	TotalPaid          float64        `json:"totalPaid"`
	TodayPaid          float64        `json:"todayPaid"`
	LastPayment        string         `json:"lastPayment"`
	LastPaymentLink    string         `json:"lastPaymentLink"`
	Performance        *WorkerStats   `json:"performance"`
	PerformanceSamples []*WorkerStats `json:"performanceSamples"`
}

type WorkerStats struct {
	Created string                             `json:"created"`
	Workers map[string]*WorkerPerformanceStats `json:"workers"`
}

type WorkerPerformanceStats struct {
	Hashrate         int64   `json:"hashrate"`
	ReportedHashrate int64   `json:"reportedHashrate"`
	SharesPerSecond  float64 `json:"sharesPerSecond"`
}

type DailyEarning struct {
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
}

type BalanceChange struct {
	PoolId  string  `json:"poolId"`
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
	Usage   string  `json:"usage"`
	Created string  `json:"created"`
}

type PoolPerformance struct {
	PoolHashrate         float64 `json:"poolHashrate"`
	ConnectedMiners      int32   `json:"connectedMiners"`
	ValidSharesPerSecond int32   `json:"validSharesPerSecond"`
	NetworkHashrate      float64 `json:"networkHashrate"`
	NetworkDifficulty    float64 `json:"networkDifficulty"`
	Created              string  `json:"created"`
}

type MinerSettings struct {
	PaymentThreshold float64 `json:"paymentThreshold"`
}

type MinerSettingsUpdateReq struct {
	IPAddress string         `json:"ipAddress"`
	Settings  *MinerSettings `json:"settings"`
}

type MinerSettingsUpdateRes struct {
	Success             bool           `json:"success"`
	ResponseMessageType int64          `json:"responseMessageType"`
	ResponseMessageId   string         `json:"responseMessageId"`
	ResponseMessageArgs []string       `json:"responseMessageArgs"`
	Result              *MinerSettings `json:"result"`
}
