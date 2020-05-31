package controller

type ChainDb struct {
	ChannelName   string   `json:"channelName"`
	ChainCodeName string   `json:"chainCodeName"`
	FunctionName  string   `json:"functionName"`
	Data          []string `json:"data"`
}

type AssetChaincode struct {
	ChannelName   string `json:"channelName"`
	ChainCodeName string `json:"chainCodeName"`
	FunctionName  string `json:"functionName"`
	Assets        Asset  `json:"assets"`
}
type Asset struct {
	AssetId      string `json:"Asset_id"`
	CargoName    string `json:"cargo_name"`
	CargoPrice   string `json:"cargo_price"`
	CargoAmount  string `json:"cargo_amount"`
	ContractId   string `json:"contract_id"`
	ProviderId   string `json:"provider_id"`
	ProviderName string `json:"provider_name"`
	CreateTime   string `json:"create_time"`
	CreateUser   string `json:"create_user"`
	UserDetail   string `json:"user_detail"`
	CargoAddress string `json:"cargo_address"`
}
