package main

import (
	"fmt"
	"os"
)

func main() {
	fileBytes, err := os.ReadFile("./dump.sql")
	if err != nil {
		fmt.Printf("error reading file: %s", err)
		os.Exit(1)
	}

	contents := string(fileBytes)

	tables := ParseDumpFile(contents)
	fmt.Printf("%+v", tables["public.posts"])
}
