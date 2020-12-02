package main
import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
)

func main() {
  inputUrl := "https://adventofcode.com/2020/day/1/input"

  resp, err := http.Get(inputUrl)
  if err != nil {
    fmt.Printf("%s", err)
    os.Exit(1)
  }
  defer resp.Body.Close()
  contents,err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Printf("%s", err)
    os.Exit(1)
  }
  fmt.Printf("%s\n", string(contents))
}
