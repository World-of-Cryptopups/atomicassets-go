package main

import (
	"fmt"
	"log"

	"github.com/World-of-Cryptopups/atomicassets-go"
)

func main() {
	a := atomicassets.New()

	q, err := a.GetAccountCollection("5g2vm.wam", "cryptopuppie")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q.Data)

}
