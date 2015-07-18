package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	if os.Args[1] == "iso3166" {
		result, err := scrapeISO3166()
		checkError(err)

		out, err := json.Marshal(result)
		checkError(err)

		fmt.Fprintln(os.Stdout, string(out))
	} else if os.Args[1] == "iso639" {
		result, err := scrapeISO639()
		checkError(err)

		out, err := json.Marshal(result)
		checkError(err)

		fmt.Fprintln(os.Stdout, string(out))
	} else {
		fmt.Fprintln(os.Stderr, "unknown command")
		os.Exit(1)
	}
}
