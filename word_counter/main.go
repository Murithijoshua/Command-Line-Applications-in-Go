package main

import (
	"bufio"
	// printing text
	"fmt"
	//flag library
	"flag"
	"io"
	"os"
)

func main() {
	byt := flag.Bool("b", false, "count bytes")
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *byt))
}

func count(r io.Reader, countlines bool, bits bool) int {
	// read text
	scanner := bufio.NewScanner(r)
	//fmt.Println(scanner)
	if bits {
		scanner.Split(bufio.ScanBytes)
	}
	if !countlines {
		scanner.Split(bufio.ScanWords)
	}
	//split the text

	wordCounter := 0
	// count words scanned
	for scanner.Scan() {
		wordCounter++
	}
	return wordCounter
}
