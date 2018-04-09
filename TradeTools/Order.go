package TradeTools

import (
	. "github.com/Snooowgh/GoEx"
	. "github.com/Snooowgh/GoEx/huobi"
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/MetaData"
	"time"
	"fmt"
)

var hb2 HuoBi_V2

var DEBUG bool = false

func InitializeOrderTools(hb HuoBi_V2){
	hb2 = hb
}


func GetReferencePrice(pair CurrencyPair)float64{
	b,s := hb2.GetRealtimePrice(pair)
	return (b+s)/2
}

//Control the position of target pair to a certain rate (Market Vale Rate)
//Little amount
func order_to(pair CurrencyPair,position float64,isTest bool)(*Bill,error){
	acc,err:=hb2.GetAccount()
	if err!=nil {
		panic(err)
	}
	cur_posT,cur_posQ := GetCurrentPosition(*acc,pair)
	ticker,err := hb2.GetTicker(pair)
	if err!=nil {
		return nil,err
	}
	buy:=ticker.Buy
	sell:=ticker.Sell
	total_quote := cur_posT*buy+cur_posQ
	cur_position := cur_posT*buy/total_quote

	//fmt.Println(position," ",cur_position," ",total_quote,pair.CurrencyB,time.Now())

	if(IsWaveIn(0.01,cur_position,position)){
		return nil,nil
	}else{
		_,a := GetSymbolPrecision(pair)
		amount := Abs((cur_position-position)*total_quote)
		if pair.CurrencyB.Symbol==BTC.Symbol {
			if amount<MinBTC_MarketSell {
				return nil,nil
			}
		}else if pair.CurrencyB.Symbol==USDT.Symbol{
			if amount<MinUSDT_MarketBuy{
				return nil,nil
			}
		}
		if cur_position>position {
		//	sell
			amount = amount/sell
			if PreciseFloat(amount,a)!=0 {
				times := amount/5
				if times:=int(times);times>1 {
					if !isTest {
						for times--; times>0;  {
							_,err:=hb2.MarketSell(TransFloat(amount,a),"",pair)
							fmt.Println("sell amount:",TransFloat(amount,a))
							if err!=nil {
								fmt.Println(err)
							}
						}
					}
				}else {
					if !isTest {
						_,err:=hb2.MarketSell(TransFloat(amount, a), "", pair)
						fmt.Println("sell amount:",TransFloat(amount,a))
						if err!=nil {
							fmt.Println(err)
						}

					}
				}
				return &Bill{"sell",amount,total_quote,buy,pair,time.Now()},nil
			}
		}else{
		//	buy
			if PreciseFloat(amount,a)!=0 {
				times := amount/5
				if times:=int(times);times>1 {
					if !isTest {
						for times--; times>0;  {
							_,err:=hb2.MarketBuy(TransFloat(amount,a),"",pair)
							fmt.Println("buy amount:",TransFloat(amount,a))
							if err!=nil {
								fmt.Println(err)
							}
						}
					}
				}else {
					if !isTest {
						_,err:=hb2.MarketBuy(TransFloat(amount,a),"",pair)
						fmt.Println("buy amount:",TransFloat(amount,a))
						if err!=nil {
							fmt.Println(err)
						}
					}
				}
				return &Bill{"buy",amount,total_quote,sell,pair,time.Now()},nil
			}
		}
	}
	return nil,nil
}
//type Bill struct {
//	Type string
//	Amount,
//	TotalQuate,
//	Price float64
//	Pair goex.CurrencyPair
//	Time time.Time
//}

func Order_to_test(pair CurrencyPair,position float64)(*Bill,error) {
	return order_to(pair,position,true)
}

func Order_to(pair CurrencyPair,position float64)(*Bill,error) {
	if !DEBUG {
		return order_to(pair, position, false)
	}else{
		return order_to(pair,position,true)
	}
}


func IsWaveIn(wave ,quote,target float64,) bool{
	if target>quote*(1-wave) && target<quote*(1+wave) {
		return true
	}else{
		return false
	}
}





func GetCurrentPosition(acc Account,pair CurrencyPair) (float64,float64) {
	targetAmountA := acc.SubAccounts[pair.CurrencyA].Amount
	targetAmountB := acc.SubAccounts[pair.CurrencyB].Amount
	return targetAmountA,targetAmountB
}


