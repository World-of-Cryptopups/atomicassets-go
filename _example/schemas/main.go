package main

import (
	"fmt"
	"log"

	"github.com/World-of-Cryptopups/atomicassets-go"
)

func main() {
	GetSchemaByNameStats()
	GetSchemaByNameLogs()
	GetSchemaByName()
	GetSchemas()
}

func GetSchemaByNameStats() {
	a := atomicassets.New()
	q, err := a.GetSchemaNameStats("cryptopuppie", "pupitems")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)
}

func GetSchemaByNameLogs() {
	a := atomicassets.New()

	q, err := a.GetSchemaNameLogs("cryptopuppie", "pupitems", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)
}

func GetSchemaByName() {
	a := atomicassets.New()

	q, err := a.GetSchemaName("cryptopuppie", "pupitems")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)
}

func GetSchemas() {
	a := atomicassets.New()

	q, err := a.GetSchemas(&atomicassets.GetSchemasQuery{
		CollectionName: "cryptopuppie",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q.Data[0].Collection.Author)
	fmt.Println(q.Data[0].Collection.Name)
}
