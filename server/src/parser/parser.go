package parser

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type Parser struct {
    path string
}

func (p *Parser) Parse(path string, obj interface{}) interface{} {
  file, err := os.Open(path)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
 defer file.Close()

  reader := bufio.NewReader(file)
  scanner := bufio.NewScanner(reader)
    map := make(map[string]string)
  for scanner.Scan() {
    line := scanner.Text())
    objs = strings.split(line," ")
    map[objs[0]]= objs[1]
  }
  return obj
}
