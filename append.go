package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    arg1 := os.Args[1]
    arg2 := os.Args[2]

    file1, err := os.Open(arg1)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file1.Close()

    file2, err := os.Open(arg2)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file2.Close()

    scanner1 := bufio.NewScanner(file1)

    for scanner1.Scan() {
        scanner2 := bufio.NewScanner(file2)
        for scanner2.Scan() {
            line2 := scanner2.Text()
            words := strings.Split(line2, " ")
            if len(words) > 0 {
                firstWord := words[0]
                restOfLine := strings.Join(words[1:], "")
                fmt.Println(strings.TrimSpace(scanner1.Text()) + firstWord + strings.TrimSpace(restOfLine))
            }
        }
        file2.Seek(0, 0)
    }
}

