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
    kvp := make(map[string]string)
  for scanner.Scan() {
    line := scanner.Text())
    objs = strings.split(line," ")
    kvp[objs[0]]= objs[1]
  }
  
  
s := reflect.ValueOf(&obj).Elem()
 if s.Kind() == reflect.Struct {
   for k, v := range kvp {
      // exported field
        f := s.FieldByName(k)
        if f.IsValid() {
            // A Value can be changed only if it is 
            // addressable and was not obtained by 
            // the use of unexported struct fields.
            if f.CanSet() {
                // change value of N
                if f.Kind() == reflect.Int {
                    x := int64(v)
                    if !f.OverflowInt(x) {
                        f.SetInt(x)
                    }
                }else if f.Kind() == reflect.Float64 {
                    x := float64(v)
                    if !f.OverflowFloat(x) {
                        f.SetFloat64(x)
                    }
                } else if f.Kind() ==reflect.String {
                  f.SetString(v)
                }
            }
        }
    }
       
    }
  
  return obj
}
