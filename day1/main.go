package main
import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
)

func main() {
  inputFile := "input.txt"

  file, err := os.Open(inputFile)
  if err != nil { 
    fmt.Printf("%s", err)
    os.Exit(1) 
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  
  for i, line := range lines {
    fmt.Println(i, line)
  }
}
