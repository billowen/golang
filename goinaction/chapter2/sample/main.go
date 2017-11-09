package main

import (
	"log"
	"os"
	_ "github.com/billowen/golang/goinaction/chapter2/sample/matchers"
	"github.com/billowen/golang/goinaction/chapter2/sample/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}