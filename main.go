package main

import "prefect/services/parser"

func main() {
	for {
		parser.SysDataParser()
	}
}