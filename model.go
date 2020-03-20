package main

import (
	"encoding/json"
)

type orderReq struct {
	Symbol string `json:"symbol"`
	Type string `json:"type"`
	Side string `json:"side"`
	Price string `json:"price"`
	Quantity string `json:"quantity"`
	MsgNo string `json:"msgNo"`
}



/* --------------------------------------------  */

type baseResp struct {
	Code      string
	Msg       string
	Timestamp int64
	Data interface{}
}

type serverTime struct {
	baseResp
	Data int64
}

type coinConfig struct {
	Name           string
	FullName       string
	DepositStatus  string
	WithDrawStatus string
	MinWithDraw    json.Number
	WithDrawFee    json.Number
	TakerFeeRate   json.Number
	MakerFeeRate   json.Number
}

type spotConfig struct {
	Symbol   string
	Accuracy []string
}

type contractConfig struct {
	Symbol       string
	MakerFeeRate string
	TakerFeeRate string
}

type configResp struct {
	baseResp
	Data struct {
		CoinConfig     []coinConfig
		SpotConfig     []spotConfig
		ContractConfig []contractConfig
	}
}

type tickerResp struct {
	baseResp
	Data []struct {
		P   string
		C   string
		S   string
		V   string
		H   string
		L   string
		Ver string
	}
}

type orderResp struct {
	baseResp
	Data struct{
		Symbol string
		OrderId string
	}
}

type assetsResp struct {
	baseResp
	Data []struct{
		CoinType string
		Count string
		Frozen string
		Type string
		BtcQuantity string
	}
}
