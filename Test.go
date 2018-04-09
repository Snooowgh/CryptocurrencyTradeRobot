package main

import (
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/tools"
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/PersonalData"
	"fmt"
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/MetaData"
	//. "github.com/Snooowgh/GoEx"
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/TradeTools"
	//"strconv"
	//"time"
)


//获取accountID函数
func main() {
	InitializeSymbols(*Hb2)
	InitializeOrderTools(*Hb2)
	InitializeEmail(Email , Pwd , ToEmail)
	id,_ := Hb2.GetAccountId()
	fmt.Println(id)
	//
	//o,_:=Order_to_test(GNX_BTC,0.9)
	//if o!=nil{
	//	fmt.Println(o.Type,o.Price,o.Amount,o.Pair,o.GetFees())
	//}
	//fmt.Println(GetPrice_MA_N_D(BTC_USDT,10,"1min"))
	//klines,_ := Hb2.GetKlineHistoryRecords(BTC_USDT,"1day","5")
	//for i,k:=range klines{
	//	fmt.Printf("%d %v\n",i,k)
	//}
	//Hb2.LimitMarginBuy("0.01","1",ETH_USDT)
	//o,_ := Hb2.MarketMarginSell("0.1","",ETH_USDT)
	//fmt.Printf("%v\n",o)
	//time.Sleep(time.Second*2)
	//m,_ := Hb2.GetOneOrder(strconv.Itoa(o.OrderID),ETH_USDT)
	//m,_ := Hb2.GetOneOrder("3051250327",ETH_USDT)
	//m,_:= Hb2.GetMarginBalance(ETH_USDT)
	//fmt.Printf("%v\n",m)
}