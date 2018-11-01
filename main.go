package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go c()

	s := "hey"
	db, _ := LoadDatabase("GBOHNXALOLI6FXK7ECKKOMNG5FIZ3F25RZDLWRNHOUNHCELACSBYYIM5", "SDTSDIEUASWONDSVPMVTUK22UGJC3BGOUBO3IWMU5ZXSKWFJAFUUMIAT")

	db.Put(s, "test")

	x, _ := db.Get(s)
	ss := x.([]byte)
	fmt.Println(string(ss))
	wg.Wait()
}
