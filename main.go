package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "math/rand"
  "flag"
  "io"
)

func main() {

  var (
    path string
    src  string
    help bool
  )

  flag.StringVar(&path, "path", "./", "Specify output path")
  flag.StringVar(&src, "src", "", "Specify input file")
  flag.BoolVar(&help, "h", false, "Usage:")

  flag.Parse()

  var reader io.Reader

  if help == true {
    fmt.Println("USAGE:\t jsminifier -src \"[INPUT FILE]\" -p \"[OUTPUT FILE]\"\n")
    fmt.Println("\t FLAGS: -src    : Input file\n\t      \t-path   : Output file")
    return
  }

  if src == "" {
    a, err := os.Stdin.Stat()
    if err != nil {
      fmt.Println("Error")
    }
    if ((a.Mode() & os.ModeNamedPipe) == 0) {
      fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
      fmt.Println("SRC not defined. Use -h for help of use pipline input.")
      fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -\n")
      return
    } else {
      fmt.Println("Input source: STDIN")
      reader = os.Stdin
    }
  } else {
    fl, err := os.Open(src)
    if err != nil {
      fmt.Println("Error opening file...")
      return
    }
    defer fl.Close()
    reader = fl
  }

  var output string
  if path == "./" {
    output = path + "minified.js"
  } else {
    output = path
  }

  var res string
  varMap := make(map[string]string)
  r := bufio.NewScanner(reader)

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

  f, err := os.Create(output)
  defer f.Close()
  if err != nil {
    fmt.Println("Error...")
  }

  f.WriteString(res)
}
