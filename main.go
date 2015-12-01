package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var rds []io.Reader
	for _, f := range os.Args[1:] {
		file, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}
		rds = append(rds, file)
	}
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		rds = append(rds, os.Stdin)
	}

	reader := io.MultiReader(rds...)
	empty := struct{}{}
	midMap := make(map[string]struct{})
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		midMap[scanner.Text()] = empty
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(midMap))
}
