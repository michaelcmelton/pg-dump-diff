package main

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"strings"
)

type Table struct {
	Name    string
	Hash    string
	Columns map[string]*Column
	Raw     string
}

type Column struct {
	Hash string
	Raw  string
	Name string
}

func ParseDumpFile(fileContents string) map[string]*Table {
	createTable := regexp.MustCompile(`(?i)CREATE\s+TABLE\s+(\w+\.\w+)\s*\((?:[^()]*|\((?:[^()]*|\([^()]*\))*\))*\);?`)
	matches := createTable.FindAllStringSubmatch(fileContents, -1)

	tables := make(map[string]*Table)
	for _, tableMatch := range matches {
		table := &Table{}
		table.Name = tableMatch[1]
		table.Raw = tableMatch[0]
		table.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(table.Raw)))
		table.Columns = parseColumns(tableMatch[0])
		tables[table.Name] = table
	}

	return tables
}

func parseColumns(tableDefinition string) map[string]*Column {
	tableContents := regexp.MustCompile(`\(\s*([\s\S]*?)\s*\);`)
	columns := make(map[string]*Column)
	match := tableContents.FindStringSubmatch(tableDefinition)
	columnStrings := strings.Split(match[0], ",")
	fmt.Printf("%v", columnStrings)
	return columns
}
