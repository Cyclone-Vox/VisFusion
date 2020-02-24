package main

import (
	"VisFusion/Director"
	"VisFusion/ReversProxy"
	"context"
)

func init()  {

}



//this place SetUp Director & ReversProxy Goroutines And Connection of data.
//And SetUp Watcher to Configuration File
//Also,it Sets Cfg into Goroutines in this function.
func main()  {
	ctx, cancel := context.WithCancel(context.Background())
	Director.Director(ctx,"8899")
	ReversProxy.ProxySetUp(ctx,)

}
