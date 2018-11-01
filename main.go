package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	db, _ := LoadDatabase("GBOHNXALOLI6FXK7ECKKOMNG5FIZ3F25RZDLWRNHOUNHCELACSBYYIM5", "SDTSDIEUASWONDSVPMVTUK22UGJC3BGOUBO3IWMU5ZXSKWFJAFUUMIAT")
	go db.c()

	fmt.Println(`
	Hey folks, this is Gravity, a super hacky decentralized database built on Stellar. It's very crappy and needs a refractor.
	Definitely DO NOT use this for production. If you feel like helping out, checkout https://github.com/sbsends/gravity.
	Also, maybe follow @sbsends on twitter. He just got a twitter and is self-conscious about how few followers he has. 
	Commands:

	put <key> <value>
	get <key>
	
	when you see something like 'syncing: <key>.0' just wait a second and then query. 
	`)

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		arr := strings.Split(text, " ")
		if arr[0] == "put" && len(arr) == 3 {
			fmt.Println("this takes a lil while (~5s)")
			err := db.Put(arr[1], arr[2])
			if err != nil {
				fmt.Println("something broke probably")
			}
		} else if arr[0] == "get" && len(arr) == 2 {
			x, err := db.Get(arr[1])
			if err != nil {
				fmt.Println("something broke probably")
			}
			fmt.Println(string(x.([]byte)))
		} else {
			fmt.Println("something broke probably")
		}
	}

	// db.Put(s, "test")

	// x, _ := db.Get(s)
	// ss := x.([]byte)
	// fmt.Println(string(ss))

	wg.Wait()
}
