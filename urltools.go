package main

import (
	"fmt"
	"os"
	"flag"
	"net/url"
)

const ENCODE = "encode"
const DECODE = "decode"

func main() {
	encodeCmd := flag.NewFlagSet(ENCODE, flag.ExitOnError)
	decodeCmd := flag.NewFlagSet(DECODE, flag.ExitOnError)

	if len(os.Args) < 3 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	switch os.Args[1] {
	case ENCODE:
		encodeCmd.Parse(os.Args[2:])
	case DECODE:
		decodeCmd.Parse(os.Args[2:])
	default:
		fmt.Println("Invalid command")
		flag.PrintDefaults()
	}

	inputString := os.Args[2]

	if encodeCmd.Parsed() {
		result := url.QueryEscape(inputString)
		fmt.Println(result)
	}

	if decodeCmd.Parsed() {
		result, err := url.QueryUnescape(inputString)
		if err != nil {
			fmt.Printf("Could not decode %s, reason: %q", inputString, err)
			os.Exit(1)
		}
		fmt.Println(result)
	}
}
