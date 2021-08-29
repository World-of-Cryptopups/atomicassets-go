package main

import (
	"fmt"
	"log"

	"github.com/World-of-Cryptopups/atomicassets-go"
)

func main() {
	GetTemplateByIDStats()
	GetTemplateByIDLogs()
	GetTemplateID()
	GetTemplates()
}

func GetTemplateByIDStats() {
	a := atomicassets.New()
	q, err := a.GetTemplateIDStats("cryptopuppie", 219081)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)
}

func GetTemplateByIDLogs() {
	a := atomicassets.New()

	q, err := a.GetTemplateIDLogs("cryptopuppie", 219081, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)
}

func GetTemplateID() {
	a := atomicassets.New()

	q, err := a.GetTemplateID("cryptopuppie", 219081)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)
}

func GetTemplates() {
	a := atomicassets.New()

	q, err := a.GetTemplates(&atomicassets.GetTemplatesQuery{
		CollectionName: "cryptopuppie",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q.Data[0].Collection.Author)
}
