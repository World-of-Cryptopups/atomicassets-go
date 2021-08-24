package main

import (
	"fmt"
	"log"

	"github.com/World-of-Cryptopups/atomicassets-go"
)

func main() {
	GetAssetsByIDStats()
	GetAssetByIDLogs()
	GetAssetByID()
	GetAssets()
}

func GetAssetsByIDStats() {
	a := atomicassets.New()
	q, err := a.GetAssetIDStats("1099559587700")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)
}

func GetAssetByIDLogs() {
	a := atomicassets.New()

	q, err := a.GetAssetIDLogs("1099559587700", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)
}

func GetAssetByID() {
	a := atomicassets.New()

	q, err := a.GetAssetID("1099559587700")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q.Data.Name + " " + q.Data.Owner)
}

func GetAssets() {
	a := atomicassets.New()

	q, err := a.GetAssets(&atomicassets.GetAssetsQuery{
		CollectionName: "cryptopuppie",
		Owner:          "5g2vm.wam",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q.Data[0].AssetID)
	fmt.Println(q.Data[0].Name)
}
