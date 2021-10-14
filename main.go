package main

import (
	"flag"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	inputFile = flag.String("i", "", "Fully qualified path of packages to analyse")
	output    = flag.String("o", ".", "Protobuf output file.")
)

func main() {
	flag.Parse()

	if *inputFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	outputFile, err := ensureOutputFile(*output)
	if err != nil {
		log.Fatalf("ensureOutputFile error: %s", err)
	}

	err = buildFile2proto(*inputFile, outputFile)
	if err != nil {
		log.Fatalf("error buildFile2proto: %s", err)
	}
	log.Printf("output file written to %v\n", outputFile)
}
