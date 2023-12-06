package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KarelKubat/jdcal"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: go run <this-thing.go>) ZONE-NAME")
	}
	for _, e := range jdcal.ZonesByName(os.Args[1]) {
		fmt.Println(e)
	}
}
