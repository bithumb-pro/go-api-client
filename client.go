package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"crypto/sha256"
	"encoding/hex"
	"crypto/hmac"
	"encoding/json"
	"strconv"
	"time"
)

var baseUrl = "https://global-openapi.bithumb.pro/openapi/v1"

var serverTimeUrl = baseUrl + "/serverTime"

var spotUrl = baseUrl + "/spot"
var configUrl = spotUrl + "/config"
var tickerUrl = spotUrl + "/ticker?symbol=ALL"
var placeOrderUrl = spotUrl + "/placeOrder"
var cancelUrl = spotUrl + "/cancelOrder"
var assetsUrl = spotUrl + "/assetList"

type Client struct {
	apiKey    string
	secretKey string
}

func (c *Client) getApi() Api {
	var api Api
	api = c
	return api
}

func (c *Client) toJson(obj interface{}) string {
	jsonBytes, _ := json.Marshal(&obj)
	return string(jsonBytes)
}

func (c *Client) getSha256HashCode(preSign string) string {
	h := hmac.New(sha256.New, []byte(c.secretKey))
	h.Write([]byte(preSign))
	hashCode := hex.EncodeToString(h.Sum(nil))
	return hashCode
}

func (c *Client) struct2map(params interface{}) map[string]string {
	t := reflect.TypeOf(params)
	v := reflect.ValueOf(params)
	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Tag.Get("json")] = v.Field(i).String()
	}
	return data
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Client) get(url string, params interface{}, result interface{}) {
	err := get(url, &result)
	handleErr(err)
}

func (c *Client) post(url string, params interface{}, result interface{}) {
	preMap := c.struct2map(params)
	preMap["apiKey"] = c.apiKey
	preMap["timestamp"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	preMap["signature"] = c.sign(preMap)
	err := post(url, preMap, result)
	handleErr(err)
}

func (c *Client) sign(preMap map[string]string) string {
	var keys []string
	for k := range preMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var preSign string
	for _, k := range keys {
		preSign += k + "=" + preMap[k] + "&"
	}
	preSign = strings.TrimSuffix(preSign, "&")
	fmt.Println("prepare signature string >======= ", preSign)
	signature := c.getSha256HashCode(preSign)
	fmt.Println("signature string >====== ", signature)
	return signature
}

func (c *Client) Time() int64 {
	var r serverTime
	c.get(serverTimeUrl, nil, &r)
	return (&r).Data
}

func (c*Client) Config() *configResp {
	var r configResp
	c.get(configUrl, nil, &r)
	return &r
}

func (c *Client) Ticker() *tickerResp {
	var r tickerResp
	c.get(tickerUrl, nil, &r)
	return &r
}

/* -------------------------------------------------- */

func (c *Client) PlaceOrder(orderReq orderReq) *orderResp {
	var r orderResp
	c.post(placeOrderUrl, orderReq, &r)
	return &r
}

func (c *Client) CancelOrder(symbol string, orderId string) *baseResp {
	var r baseResp
	p := struct {
		Symbol  string `json:"symbol"`
		OrderId string `json:"orderId"`
	}{
		symbol, orderId,
	}
	c.post(cancelUrl, p, &r)
	return &r
}

func (c *Client) Assets(coinType string) *assetsResp{
	var r assetsResp
	p := struct {
		CoinType  string `json:"coinType"`
		AssetType string `json:"assetType"`
	}{
		coinType, "spot",
	}
	c.post(assetsUrl, p, &r)
	return &r
}


type Api interface {
	Time() int64
	Ticker() *tickerResp
	Config() *configResp

	PlaceOrder(orderReq) *orderResp
	CancelOrder(symbol string, orderId string) *baseResp
	Assets(coinType string) *assetsResp
}