package main

import (
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/PersonalData"
	"github.com/Snooowgh/CryptocurrencyTradeRobot/strategyframe"
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/MetaData"
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/TradeTools"
	. "github.com/Snooowgh/CryptocurrencyTradeRobot/tools"
	_ "github.com/Snooowgh/CryptocurrencyTradeRobot/strategy"
	"os"
	"os/signal"
)


func main() {
	InitializeSymbols(*Hb2)
	InitializeOrderTools(*Hb2)
	InitializeBill()
	InitializeEmail(Email , Pwd , ToEmail)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func(plugins map[string]strategyframe.Strategy) {
		for sig := range c {
			// handle
			if sig != nil {
				for _, plugin := range plugins {
					if plugin.Flag(){
						plugin.EndHandle()
					}
				}
				os.Exit(0)
				return
			}
		}
	}(strategyframe.Plugins)

	acc, err := Hb2.GetAccount()
	if err != nil {
		panic(err)
	}
	strategyframe.Start(*Hb2,*acc)
}