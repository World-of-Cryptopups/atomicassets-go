package main

import (
	"fmt"
	"log"

	"github.com/World-of-Cryptopups/atomicassets-go"
)

func main() {
	GetCollectionByNameStats()
	GetCollectionByNameLogs()
	GetCollectionByName()
	GetCollections()
}

func GetCollectionByNameStats() {
	a := atomicassets.New()
	q, err := a.GetCollectionNameStats("cryptopuppie")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)
}

func GetCollectionByNameLogs() {
	a := atomicassets.New()

	q, err := a.GetCollectionNameLogs("cryptopuppie", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)
}

func GetCollectionByName() {
	a := atomicassets.New()

	q, err := a.GetCollectionName("cryptopuppie")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)
}

func GetCollections() {
	a := atomicassets.New()

	q, err := a.GetCollections(&atomicassets.GetCollectionsQuery{
		Author: "v.dr4.wam",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q.Data[0].AuthorizedAccounts)
	fmt.Println(q.Data[0].Name)
}
