package strategy

import (
	"fmt"
	"github.com/Snooowgh/CryptocurrencyTradeRobot/strategyframe"
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/MetaData"
	"github.com/Snooowgh/GoEx"
	"github.com/Snooowgh/GoEx/huobi"
	"strconv"
)

type MarketMakerStrategy struct {
	Name string
	Capital_base float64
	OrderPool map[goex.CurrencyPair][]goex.Order
	hb2 huobi.HuoBi_V2
}


func init() {
	pool := make(map[goex.CurrencyPair][]goex.Order)
	Strategy := &MarketMakerStrategy{Name:"做市商交易策略",Capital_base:0.5,OrderPool:pool}
	strategyframe.Regist(Strategy.Name, Strategy)
}


func (this *MarketMakerStrategy) Flag() bool {
	return false
}

func (this *MarketMakerStrategy) Initialize(hb huobi.HuoBi_V2,acc goex.Account)  {
	this.hb2 = hb
	fmt.Print("初始化")
}

func (this *MarketMakerStrategy) Handle_data() {
	for _,pair := range AllPairs{
		if len(this.OrderPool[pair])==0 {
		//	无订单
			b,s := this.hb2.GetRealtimePrice(pair)
			p,a:=GetSymbolPrecision(pair)
			var minPre float64
			minPre = 1.0/float64(10^p)
			b+=minPre
			s-=minPre
			if CheckFees(b,s) {
				MinAmount := 1.0/float64(10^a)
				
				this.hb2.LimitSell(TransFloat(MinAmount,a),TransFloat(s,p),pair)
				this.hb2.LimitBuy(TransFloat(MinAmount,a),TransFloat(b,p),pair)
			}
		}else{
		//	有订单
			for _,order:=range this.OrderPool[pair] {
				o,err:=this.hb2.GetOneOrder(strconv.Itoa(order.OrderID),pair)
				if err!=nil{
					continue
				}else{
					if o.Status==goex.ORDER_FINISH{
						this.OrderPool[pair] = removeOrder(this.OrderPool[pair],*o)
					}else {
					//	有未完成的订单
					}
				}
			}
		}
	}
}
func (this *MarketMakerStrategy) RefreshRate() (int,string){
	return 0,""
}

func (this *MarketMakerStrategy) EndHandle() {


}

func removeOrder(slice []goex.Order, elems goex.Order) []goex.Order {
	for i := range slice {
		if slice[i].OrderID== elems.OrderID {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return  slice
}

func CheckFees(buy float64,sell float64)bool{
	if buy>sell{
		return false
	}else{
		a := buy*Fee_Percentage+sell*Fee_Percentage
		b := sell-buy
		if a > b {
			return false
		}else {
			return true
		}
	}
}
