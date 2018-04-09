package strategy

import (
	"fmt"
	"github.com/Snooowgh/CryptocurrencyTradeRobot/strategyframe"
	"github.com/Snooowgh/GoEx"
	"github.com/Snooowgh/GoEx/huobi"
)

type GoingStrategy struct {
	Name string
	Capital_base float64
}


func init() {
	Strategy := &GoingStrategy{Name:"24点追高交易策略",Capital_base:0.1}
	strategyframe.Regist(Strategy.Name, Strategy)
}


func (this *GoingStrategy) Flag() bool {
	return false
}

func (this *GoingStrategy) Initialize(hb huobi.HuoBi_V2,acc goex.Account)  {
	fmt.Print("初始化")
}

func (this *GoingStrategy) Handle_data() {

	fmt.Println("我是",this.Name)


}
func (this *GoingStrategy) RefreshRate() (int,string){
	return 2,"d"
}

func (this *GoingStrategy) EndHandle() {


}