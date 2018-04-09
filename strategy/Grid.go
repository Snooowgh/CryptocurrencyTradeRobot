package strategy

import (
	"github.com/Snooowgh/CryptocurrencyTradeRobot/strategyframe"
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/MetaData"
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/TradeTools"
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/tools"
	"fmt"
	"github.com/Snooowgh/GoEx/huobi"
	"github.com/Snooowgh/GoEx"
	"github.com/robfig/cron"
	"os"
)

//改进方案   设定价格波动百分率 自动画网格   固定价格或动态价格

type gridStrategy struct {
	Name string
	Capital_base float64
}

type GridParams struct{
	TargetPriceRate,
	TargetPosition float64
}

var AllBills map[goex.CurrencyPair]([]Bill)
var CoinPool []goex.CurrencyPair

func init() {
	Strategy := &gridStrategy{Name:"网格交易策略",Capital_base:0.5}
	strategyframe.Regist(Strategy.Name, Strategy)
}


func (this *gridStrategy) Flag() bool {
	return false
}

func (this *gridStrategy) Initialize(hb huobi.HuoBi_V2,acc goex.Account)  {
	// Coin_Pool
	//
	AllBills = make(map[goex.CurrencyPair]([]Bill))
	CoinPool = []goex.CurrencyPair{goex.BTC_USDT}
	for _,coin := range CoinPool{
		AllBills[coin] = make([]Bill,0)
	}

	c := cron.New()
	var numOfbill map[goex.CurrencyPair]int = make(map[goex.CurrencyPair]int)
	c.AddFunc("0 */10 * * * ?", func() {
		content := ""
		for _,pair := range CoinPool{
			if numOfbill[pair]<len(AllBills[pair]){
				numOfbill[pair]=len(AllBills[pair])
				content += fmt.Sprintf("%v\n",AllBills[pair][len(AllBills[pair])-1])
				content+=fmt.Sprintf("%s 收益率:%f",pair.String(),CalcBillsProfit(AllBills[pair]))
			}
		}
		if content!=""{
			SendInfoToMyEmail(content,"网格交易成交通知")
		}
	})
	c.AddFunc("0 */50 * * * ?", func() {
		this.EndHandle()
		content := ""
		for _,pair := range CoinPool{
			base_price_Map[pair] = GetPrice_MA_N_D(pair,10,"30min")
			content+=fmt.Sprintf(pair.String()+":%f\n",base_price_Map[pair])
		}
		SendInfoToMyEmail(content,"MA平均价格更新")
	})
	c.Start()

}


var buy4,buy3,buy2,buy1,sell4,sell3,sell2,sell1 float64= 0.954,0.964,0.974,0.984,1.034,1.024,1.016,1.011

//buy4,buy3,buy2,buy1,sell4,sell3,sell2,sell1 = 0.68,0.76,0.84,0.92,1.6,1.45,1.3,1.15

var base_price_Map map[goex.CurrencyPair]float64 = make(map[goex.CurrencyPair]float64)

func (this *gridStrategy) Handle_data() {

	var Capital_To_Coin float64 = float64(len(CoinPool))

	for _,pair := range CoinPool{

		if CalcBillsProfit(AllBills[pair]) < -0.1{
			Make_Order(pair,0)
			SendInfoToMyEmail("",pair.String()+"止损通知  程序退出")
			os.Exit(0)
		} else if CalcBillsProfit(AllBills[pair]) >0.2{
			Make_Order(pair,0)
			SendInfoToMyEmail("",pair.String()+"止盈通知  程序退出")
			os.Exit(0)
		}

		reprice := GetReferencePrice(pair)

		//fmt.Println(reprice/GetPrice_MA_N_D(pair,5,"1day"),reprice/GetPrice_MA_N_D(pair,15,"1day"))
		//if reprice/GetPrice_MA_N_D(pair,5,"1day")>1.5 || reprice/GetPrice_MA_N_D(pair,15,"1day")>1.5{
		//	continue
		//}
		if base_price_Map[pair]==0.0{
			base_price_Map[pair] = GetPrice_MA_N_D(pair,10,"30min")
		}

		base_price := base_price_Map[pair]
		if reprice/base_price<buy4{
			Make_Order(pair,this.Capital_base*1/Capital_To_Coin)
		}else if reprice/base_price<buy3{
			Make_Order(pair,this.Capital_base*0.9/Capital_To_Coin)
		}else if reprice/base_price<buy2{
			Make_Order(pair,this.Capital_base*0.7/Capital_To_Coin)
		}else if reprice/base_price<buy1{
			Make_Order(pair,this.Capital_base*0.4/Capital_To_Coin)
		}else if reprice/base_price>sell4{
			Make_Order(pair,this.Capital_base*0/Capital_To_Coin)
		}else if reprice/base_price>sell3{
			Make_Order(pair,this.Capital_base*0.2/Capital_To_Coin)
		}else if reprice/base_price>sell2{
			Make_Order(pair,this.Capital_base*0.4/Capital_To_Coin)
		}else if reprice/base_price>sell1{
			Make_Order(pair,this.Capital_base*0.6/Capital_To_Coin)
		}
	}
}

func (this *gridStrategy) RefreshRate() (int,string){
	return 5,"s"
}

func (this *gridStrategy) EndHandle() {
	content := ""
	for k,v := range AllBills{
		content += fmt.Sprint("交易对",k.String(),"交易记录\n")
		content += fmt.Sprint("交易类型\t价格\t交易量\t成交时间\n")
		for _,bill := range v{
			content+= fmt.Sprint("%s\t%f\t%f\t%s\n",bill.Type,bill.Price,bill.Amount,bill.Time.Format("2006-01-02 15:04:05"))
		}
	}
	SendInfoToMyEmail(content,"策略交易通知")
}

//Order_to(hb,stock, 0)  仓位控制函数  目标持仓降为0  百分比


//函数缺陷   只针对一个交易对  仓位 USDT针对所有交易对
//
//处理bills
func Make_Order(pair goex.CurrencyPair,position float64){
	bill,err := Order_to(pair,position)
	if err != nil {
		SendInfoToMyEmail("下单错误 130行 错误内容："+err.Error(),"程序出错通知")
	}else if bill!=nil{
		AllBills[pair] = append(AllBills[pair],*bill)
	}
}