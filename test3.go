package main

import (
	mini "github.com/sasbury/mini"
	"fmt"
	"log"
)

func main() {

	config, err := mini.LoadConfiguration("./cwr.ini")
	if err != nil {
		log.Fatal(err)
	}
	first := config.String("first", "")
	fmt.Println("first:",first )
	second := config.String("second", "")
	fmt.Println("second:", second)
	third := config.String("third", "")
	fmt.Println("third:",third )
	fourth := config.String("forth", "")
	fmt.Println("fourth:", fourth)
	fifth := config.Integer("fifth", 0)
	fmt.Println("fifth:",fifth )
	sixt := config.Float("sixt", 0)
	fmt.Println("sixt:", sixt)
	seventh := config.Boolean("seventh", true)
	fmt.Println("seventh:",seventh )
	eighth := config.Boolean("eighth", true)
	fmt.Println("eighth:",eighth )

}