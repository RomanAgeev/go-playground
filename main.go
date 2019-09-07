package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "sample name is expected\n")
		os.Exit(1)
	}

	sampleName := args[0]

	sample, ok := repository[sampleName]
	if !ok {
		fmt.Fprintf(os.Stderr, "\"%v\" sample is unknown\n", sampleName)
		os.Exit(1)
	}

	fmt.Printf("%v: enter parameters...\n", sampleName)

	params := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()

		param := scanner.Text()
		if param == "" {
			break
		}

		params = append(params, param)
	}

	result, err := sample(params)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, res := range result {
		fmt.Println(res)
	}
}
