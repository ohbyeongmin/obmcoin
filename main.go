package main

import (
	"github.com/ohbyeongmin/obmcoin/explorer"
	"github.com/ohbyeongmin/obmcoin/rest"
)



func main(){
	go explorer.Start(3000)
	rest.Start(4000)
}

