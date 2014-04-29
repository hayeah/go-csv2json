package main

import (
  "os"
  // "io"
  "fmt"
  "flag"
  // "strconv"
  "encoding/csv"
  "encoding/json"
)

var separator string

func init() {
  flag.StringVar(&separator,"s",",","Field separator")
  // separator = strconv.Unqoute()
}

type json_object map[string]string
func main() {
  flag.Parse()

  var input *os.File
  var err error

  path := flag.Arg(0)
  if(path == "") {
    input = os.Stdin
  } else {
    input, err = os.Open(path)
    defer input.Close()
    if(err != nil) {
      fmt.Println(err)
      os.Exit(1)
    }
  }

  output := os.Stdout

  reader := csv.NewReader(input)

  // read header
  header, err := reader.Read()
  if(err != nil) {
    fmt.Println(err)
    os.Exit(1)
  }

  // read body
  records, err := reader.ReadAll()
  if(err != nil) {
    fmt.Println(err)
    os.Exit(1)
  }

  jsons := make([]interface{},len(records))

  for ri, record := range records {
    json := make(map[string]string)
    for fi, value := range record {
      key := header[fi]
      json[key] = value
    }
    jsons[ri] = json
  }

  encoder := json.NewEncoder(output)
  encoder.Encode(jsons)
}
