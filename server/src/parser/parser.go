package parser

import (
  "log"
  "io/ioutil"
)

type Parser struct {
    path string
}

func (p *Parser) Parse(path string, obj interface{}) interface{} {
    byteData, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }
    var all string = string(byteData)
    return nil
}
