package main

import (
	"Amanisis/register"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

func main(){

	var wg sync.WaitGroup
	wg.Add(1)
	go register.Server(wg.Done)
	wg.Wait()

}
