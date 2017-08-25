package main

import (
    "log"
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)

  fmt.Println("Enter a number enging with m to create a file of size m megabytes")
  fmt.Println("Enter a number enging with g to create a file of size g gigabytes")

  fmt.Print("Enter number ending with m or g: ")
  scanner.Scan()
  input := scanner.Text()
  input = strings.ToLower(input)

  var current_num int64

  if input[len(input) - 1] == 'm' {
    start_text := input[:len(input) - 1]
    current_num = get_int(start_text)
  } else if input[len(input) - 1] == 'g' {
    start_text := input[:len(input) - 1]
    current_num = get_int(start_text) * 1024
  } else {
    fmt.Println("Your number didn't match the format. Try entering 10m as an example.")
    return
  }

  fmt.Println("If you are sure you want to create this file click enter, else control-c.")
  scanner.Scan()
  scanner.Text()

  create_space_file(current_num)
  fmt.Println("File Created!")
}

func create_space_file(input_size int64) {
  size := int64(input_size * 1024 * 1024)
  fd, err := os.Create("output")
  if err != nil { log.Fatal("Failed to create output") }
  _, err = fd.Seek(size-1, 0)
  if err != nil { log.Fatal("Failed to seek") }
  _, err = fd.Write([]byte{0})
  if err != nil { log.Fatal("Write failed") }
  err = fd.Close()
  if err != nil { log.Fatal("Failed to close file")}
}

func get_int(num string) int64 {
  i, err := strconv.ParseInt(num, 10, 64)
  if err != nil { log.Fatal("Number conversion fail") }
  return int64(i)
}
