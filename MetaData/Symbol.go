package MetaData

import (
	"github.com/Snooowgh/GoEx"
	"github.com/Snooowgh/GoEx/huobi"
	"strings"
	"strconv"
)

//base-currency	true	string	基础币种
//quote-currency	true	string	计价币种
//price-precision	true	string	价格精度位数（0为个位）
//amount-precision	true	string	数量精度位数（0为个位）
//symbol-partition	true	string	交易区	main主区，innovation创新区，bifurcation分叉区

func init(){

}

type Symbol struct {
	Base_Currency,
	Quote_Currency string
	Price_prec int
	Amount_prec int
	Partition string
}

var Symbols map[string]Symbol

var AllPairs []goex.CurrencyPair

func InitializeSymbols(v2 huobi.HuoBi_V2){
	a,_:=v2.GetSymbols("")
	Symbols = make(map[string]Symbol)
	AllPairs = make([]goex.CurrencyPair,0)
	for _,j := range a{
		j:=j.(map[string]interface{})
		base_cur := strings.ToUpper(j["base-currency"].(string))
		quote_cur := strings.ToUpper(j["quote-currency"].(string))
		Symbols[base_cur+"_"+quote_cur] = Symbol{base_cur,quote_cur,int(j["price-precision"].(float64)),int(j["amount-precision"].(float64)),j["symbol-partition"].(string)}
		AllPairs = append(AllPairs,goex.CurrencyPair{goex.Currency{base_cur,""},goex.Currency{quote_cur,""}})
	}
}

func GetSymbolPrecision(pair goex.CurrencyPair)(int,int){
	if Symbols==nil {
		panic("Symbols not initialized!")
	}
	symbol := Symbols[pair.ToSymbol("_")]
	return symbol.Price_prec,symbol.Amount_prec
}


func TransFloat(a float64,prec int) string {
	return strconv.FormatFloat(a, 'f', prec, 64)
}

func PreciseFloat(a float64,prec int) float64 {
	f,_:=strconv.ParseFloat(strconv.FormatFloat(a, 'f', prec, 64),64)
	return f
}

func Abs(a float64)float64{
	if a>0 {
		return a
	}else{
		return -a
	}
}