package main

import (
	"ast/lexer"
	"encoding/json"
	"fmt"
)

func parseLex(code string) {
	le, err := lexer.Parse(code)
	if err != nil {
		panic(err)
	}
	fmt.Println("tokens: ", le.Tokens)
	if err := le.CreateList(); err != nil {
		panic(err)
	}
	for e := le.List.Front(); e != nil; e = e.Next() {
		fmt.Println("list: ", e.Value)
	}
	if err := le.ConstructAst(); err != nil {
		panic(err)
	}
	data, _ := json.Marshal(le.Node)
	fmt.Println("ast: ", string(data))
	result, err := le.Interpreter()
	if err != nil {
		panic(err)
	}
	fmt.Println("result: ", result)
}

func main() {
	parseLex("1 + 2 - 3 + (4 * (5 -2) ) - 1")
}
