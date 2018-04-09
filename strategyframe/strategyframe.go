package strategyframe

import (
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/tools"
	"fmt"
	"github.com/Snooowgh/GoEx/huobi"
	"github.com/robfig/cron"
	"github.com/Snooowgh/GoEx"
)

var Plugins map[string]Strategy

func init() {
	Plugins = make(map[string]Strategy)
}

type Strategy interface {
	Flag() bool
	Handle_data()
	Initialize(hb huobi.HuoBi_V2,acc goex.Account)
	RefreshRate() (int,string)
	EndHandle()
}


func Start(hb huobi.HuoBi_V2,acc goex.Account) {
	c := cron.New()
	Notification_Content := ""
	for name, plugin := range Plugins {
		if plugin.Flag() {
			plugin.Initialize(hb,acc)
			handle := plugin.Handle_data
			//刷新频率
			num,rateChar:=plugin.RefreshRate()
			spec := ""
			if  rateChar=="s"{
				spec = fmt.Sprintf("*/%d * * * * ?",num)
			}else if rateChar=="min" {
				spec = fmt.Sprintf("0 */%d * * * ?",num)
			}else if rateChar=="d"{
				//每天num点执行一次
				spec = fmt.Sprintf("0 0 %d * * ?",num)
			}else if num==0 && rateChar == ""{
				go func() {
					for   {
						handle()
					}
				}()
			}
			if spec!="" {
				c.AddFunc(spec, handle)
			}
			Notification_Content+=fmt.Sprintf("启动策略 %s\n", name)
		} else {
			Notification_Content+=fmt.Sprintf("未启用的策略 %s\n", name)
		}
	}
	SendInfoToMyEmail(Notification_Content,"交易脚本启动通知")
	c.Start()
	select {}
}

func Regist(name string, plugin Strategy) {
	Plugins[name] = plugin
}