package main

import (
	"flag"
	"log"
	"os"
)

var (
	inputFile  = flag.String("i", "", "Fully qualified path of packages to analyse")
	outputFile = flag.String("o", ".", "Protobuf output file.")
)

func main() {
	flag.Parse()

	if *inputFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := ensureOutputFile(*outputFile)
	if err != nil {
		log.Fatalf("ensureOutputFile error: %s", err)
	}
	*outputFile = file

	err = buildFile2proto(*inputFile, *outputFile)

	if err != nil {
		log.Fatalf("error buildFile2proto: %s", err)
	}
	log.Printf("output file written to %v\n", *outputFile)
}
