package MetaData

import (
	"github.com/Snooowgh/GoEx"
	"time"
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/strategyframe"
)

type Bill struct {
	Type string
	Amount,
	TotalQuate,
	Price float64
	Pair goex.CurrencyPair
	Time time.Time
}

func (bill *Bill) GetFees() float64{
	return bill.Amount*Fee_Percentage
}

var AllBills map[Strategy]map[goex.CurrencyPair]([]Bill)

func InitializeBill(){
	AllBills = make(map[Strategy]map[goex.CurrencyPair]([]Bill))
}
func AppendBill(bill Bill,pair goex.CurrencyPair,s Strategy){
	pair_bill := AllBills[s][pair]
	if len(pair_bill)==0{
		pair_bill = make([]Bill,0)
	}
	pair_bill = append(pair_bill,bill)
	AllBills[s][pair] = pair_bill
}

func GetBills(pair goex.CurrencyPair,s Strategy)[]Bill{
	return AllBills[s][pair]
}

func CalcProfitOfStrategy(pair goex.CurrencyPair,s Strategy)float64{
	return CalcBillsProfit(GetBills(pair,s))
}

func CalcBillsProfit(StrategyBills []Bill)float64{
	if len(StrategyBills)==0{
		return 0.0
	}
	s := StrategyBills[0]
	e := StrategyBills[len(StrategyBills)-1]
	var fees float64 = 0

	for _,s := range StrategyBills{
		fees+=s.GetFees()
	}

	if s.TotalQuate!=0 {
		return (e.TotalQuate-s.TotalQuate-fees)/s.TotalQuate
	}else {
		return 0.0
	}
}