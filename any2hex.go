package main

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	var b64 = flag.Bool("b", false, "encode base64 to hex")
	var hex = flag.Bool("h", false, "encode hex to base64")
	var decimal = flag.Bool("d", false, "encode decimal to base64")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options] <data>\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, "\t anything to hex\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "no data to process\n")
		flag.Usage()
		return
	}

	var inputs = flag.Args()
	switch {
	case *b64:
		base64ToHex(inputs)
	case *hex:
		hexToBase64(inputs)
	case *decimal:
		decimalToHex(inputs)
	}
}

func base64ToHex(inputs []string) {
	for _, input := range inputs {
		data, err := base64.StdEncoding.DecodeString(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not decode base64 value: %s: %s\n", err, input)
			continue
		}
		fmt.Fprintln(os.Stdout, strings.ToUpper(hex.EncodeToString(data)))
	}
}

func hexToBase64(inputs []string) {
	for _, input := range inputs {
		data, err := hex.DecodeString(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not decode hex value: %s: %s\n", err, input)
			continue
		}
		fmt.Fprintln(os.Stdout, base64.StdEncoding.EncodeToString(data))
	}
}

func decimalToHex(inputs []string) {
	for _, input := range inputs {
		data, err := strconv.ParseInt(input, 10, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not parse decimal value: %s: %s\n", err, input)
			continue
		}
		fmt.Fprintf(os.Stdout, "%X\n", data)
	}
}
