package parser

import (
	"io/ioutil"
	"log"
)

type Parser struct {
}

func (p *Parser) Parse(path string, v interface{}) interface{} {
	byteData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var all string = string(byteData)
	return nil
}
