package main

import (
	"github.com/ohbyeongmin/obmcoin/cli"
	"github.com/ohbyeongmin/obmcoin/db"
)


func main(){
	defer db.Close()
	cli.Start()
}

