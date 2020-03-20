package main

import (
	"fmt"
	"strconv"
	"time"
)

var c = &Client{
	apiKey:    "xxx",
	secretKey: "xxx",
}

func testNormal(api Api)  {
	fmt.Println(api.Time())

}

func testPlaceOrder(api Api) {
	orderReq := orderReq{
		Symbol:   "ETH-USDT",
		Type:     "limit",
		Side:     "buy",
		Price:    "110.01",
		Quantity: "0.1",
		MsgNo:    strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
	}
	fmt.Println(c.toJson(api.PlaceOrder(orderReq)))
}

func testCancelOrder(api Api) {
	fmt.Println(c.toJson(api.CancelOrder("ETH-USDT", "xxx")))
}

func testAssets(api Api) {
	fmt.Println(c.toJson(api.Assets("")))
}

func main() {

	api := c.getApi()

	testNormal(api)

	//testPlaceOrder(api)

	//testCancelOrder(api)

	//testAssets(api)


}
