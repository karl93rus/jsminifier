package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "math/rand"
)

func main() {
  r := bufio.NewScanner(os.Stdin)
  var res string
  varMap := make(map[string]string)

  for r.Scan() {
    line := r.Text()
    line = strings.Trim(line, " ")
    if strings.Index(line, "//") >= 0 {
      lineSplitted := strings.Split(line, "//")
      line = lineSplitted[0]
    }
    if strings.Index(line, "var ") >= 0 {
      lineSplitted := strings.Split(line, " ")
      varName := lineSplitted[1]
      rb := []string{"x0_", "$G_z", "n_$ij", "zt_X$o"}
      varMap[varName] = rb[rand.Intn(len(rb))] + varName + rb[rand.Intn(len(rb))]
    }
    res += line
  }

  for k, v := range varMap {
    res = strings.Replace(res, k, v, -1)
  }

  f, err := os.Create("./script.js")
  defer f.Close()
  if err != nil {
    fmt.Println("Error...")
  }

  f.WriteString(res)
}
