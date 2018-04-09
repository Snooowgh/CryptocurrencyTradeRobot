package strategy

import (
	"github.com/Snooowgh/CryptocurrencyTradeRobot/strategyframe"
	"github.com/Snooowgh/GoEx"
	"github.com/Snooowgh/GoEx/huobi"
)
//高抛低吸策略   杠杆交易  提前借好Pair   设定最低最高价  自动网格交易
//使用了点卡并且交易网格足够大  因此不考虑手续费

type SHBLStrategy struct {
	Name         string
	hb2          huobi.HuoBi_V2
	Pair		goex.CurrencyPair
	Volume,
	CurrentVolume,
	High,
	Low			float64
}

func init() {
	Strategy := &SHBLStrategy{Name: "高抛低吸",Pair:goex.ETH_USDT,Volume:3,High:420,Low:360}
	strategyframe.Regist(Strategy.Name, Strategy)
}

func (this *SHBLStrategy) Flag() bool {
	return false
}

func (this *SHBLStrategy) Initialize(hb huobi.HuoBi_V2, acc goex.Account) {
	this.hb2 = hb
}

func (this *SHBLStrategy) Handle_data() {
	//buy,sell := this.hb2.GetRealtimePrice(this.Pair)

}
func (this *SHBLStrategy) RefreshRate() (int, string) {
	return 0, ""
}

func (this *SHBLStrategy) EndHandle() {

}
