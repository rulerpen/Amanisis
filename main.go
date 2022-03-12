package main

import (
	"Amanisis/register"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go register.Server(wg.Done)
	fmt.Printf("123")
	wg.Wait()

}
