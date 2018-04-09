package TradeTools

import(
	. "github.com/Snooowgh/GoEx"
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/MetaData"
	"sort"
	"strconv"
)


func GetPrice_MA_N_D(pair CurrencyPair,n int,period string)float64{
	klines,err:=hb2.GetKlineHistoryRecords(pair,period,strconv.Itoa(n))
	if err!=nil{
		panic(err)
	}
	var res float64 = 0.0
	for _,kline:=range klines{
		res+=(kline.High+kline.Low)/2
	}
	p,_ := GetSymbolPrecision(pair)
	return PreciseFloat(res/float64(n),p)
}
//n天内最大回撤
func Max_drawDown(pair CurrencyPair,n int,period string) float64{
	klines,_:=hb2.GetKlineHistoryRecords(pair,period,string(n))
	var drawDown []float64
	for _,kline:=range klines{
		drawDown = append(drawDown,(kline.High-kline.Low)/kline.Low)
	}
	sort.Float64s(drawDown)
	return drawDown[len(drawDown)-1]
}