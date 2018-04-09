# BTC Trade Robot for HuoBi
### Fill your accessKey,secretKey,email password in PersonalData.go 
### Install Dependency:
```bash
go get -v github.com/Snooowgh/GoEx
go get -v github.com/robfig/cron
```

### How to implement?
#### achieve 5 functions for your strategy in "strategy" folder
```Go
type Strategy interface {
	Flag() bool            
	Handle_data()
	Initialize(hb huobi.HuoBi_V2,acc goex.Account)
	RefreshRate() (int,string)
	EndHandle()
}
```

#### Example:
```Go

# Enable this strategy
func (this *MyStrategy) Flag() bool {
	return true
}

# Execute when start
func (this *MyStrategy) Initialize(hb huobi.HuoBi_V2,acc goex.Account)  {
	# initialize your params
}

# Define the time duration to execute "Handle_data" function
func (this *MyStrategy) RefreshRate() (int,string){
	# Execute every 5 seconds
	return 5,"s"
}

# Put Your idea and innovation here
func (this *MyStrategy) Handle_data() {
	
}

# Execute when the strategy exits
func (this *MyStrategy) EndHandle() {
	
}
```

### Welcome to Star&Fork