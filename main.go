package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

var (
	prefix = flag.String("prefix", "", "The prefix for the new CSV files")
	split  = flag.Int("split", 0, "The number of files to split the CSV into")
)

func main() {
	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "USAGE: csvsplit [options] --prefix <prefix name of new CSV files> --split <number of files to split into> <filename.csv>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *split < 1 {
		fmt.Fprintln(os.Stderr, "--split must be greater than 1")
		flag.Usage()
	}

	if *prefix == "" {
		fmt.Fprintln(os.Stderr, "--prefix must be defined as a name (e.g novemeber-2016)")
		flag.Usage()
	}

	file := os.Args[1]

	if len(flag.Args()) != 1 {
		fmt.Println("no file name given")
		return
	}

	original, err := os.OpenFile(flag.Args()[0], 0, 0666)
	if err != nil {
		fmt.Printf("ERROR: Unable to locate the CSV file named: %v", file)
		return
	}

	fmt.Printf("INFO: Attempting to split the CSV file into %v files\n", *split)

	lines, err := csv.NewReader(original).ReadAll()
	if err != nil {
		fmt.Println("ERROR: Could not open the CSV reader...")
		return
	}

	// get the total
	total := len(lines)
	fmt.Printf("INFO: Total number of lines in original file: %v\n", total)

	per := total / *split
	fmt.Printf("INFO: There will be around %v lines per file\n", per)

	// return
}
