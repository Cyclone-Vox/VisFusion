package main

import (
	"VisFusion/CfgLoad"
	"fmt"
)

func init()  {

}



//this place SetUp Director & ReversProxy Goroutines And Connection of data.
//And SetUp Watcher to Configuration File
//Also,it Sets Cfg into Goroutines in this function.
func main()  {

	c:=CfgLoad.LoadCfg()
	fmt.Printf("%+v",c.DevCfgMap["flow"])
	//ctx, cancel := context.WithCancel(context.Background())
	//Director.Director(ctx,"8899")
	//ReversProxy.ProxySetUp(ctx,)

}
