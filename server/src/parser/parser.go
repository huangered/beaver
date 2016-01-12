package parser

import (
  "encoding/json"
  "io/ioutils"
)

type Parser struct {
}
// parse json obj to go structure obj
func (p *Parser) Parse(path string, v interface{}) interface{} {
  byt, err := ioutils.ReadAll(path)
  if err != nil {
    panic(err)
  }
  if err := json.Unmarshal(byt, v); err != nil {
        panic(err)
    }
    return obj
}
