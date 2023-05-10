package main

import (
"bufio"
"flag"
"fmt"
"os"
"path/filepath"
"regexp"
"strings"
)

func main() {
wordlistPath := flag.String("l", "", "wordlist path")
outputPath := flag.String("o", "", "output path")
silent := flag.Bool("s", false, "silent mode")
flag.Usage = func() {
fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n\n", os.Args[0])
fmt.Fprintln(os.Stderr, "Options:")
flag.PrintDefaults()
}
flag.Parse()

if *wordlistPath == "" {
fmt.Println("Please provide a wordlist using the -l flag.")
os.Exit(1)
}

file, err := os.Open(*wordlistPath)
if err != nil {
fmt.Printf("Error opening file: %v\n", err)
os.Exit(1)
}
defer file.Close()

scanner := bufio.NewScanner(file)
ipRegex := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{,3}`)
numericRegex := regexp.MustCompile(`^\d+$`)
englishRegex := regexp.MustCompile(`[a-zA-Z]`)

subdomains := make(map[string]bool)

for scanner.Scan() {
line := scanner.Text()
parts := strings.Split(line, ".")

if len(parts) > 2 {
for _, part := range parts[:len(parts)-2] {
if !ipRegex.MatchString(part) {
subparts := strings.Split(part, "-")
for _, subpart := range subparts {
if len(subpart) > 2 && subpart[0] != '0' && !numericRegex.MatchString(subpart) && englishRegex.MatchString(subpart) {
subdomains[subpart] = true
}
}
}
}
lastPart := parts[len(parts)-3]
if len(lastPart) > 2 && lastPart[0] != '0' && !numericRegex.MatchString(lastPart) && englishRegex.MatchString(lastPart) {
subdomains[lastPart] = true
}
}
}

if err := scanner.Err(); err != nil {
fmt.Printf("Error reading file: %v\n", err)
os.Exit(1)
}

if *outputPath != "" {
outputFile, err := os.OpenFile(*outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
fmt.Printf("Error opening output file: %v\n", err)
os.Exit(1)
}
defer outputFile.Close()

writer := bufio.NewWriter(outputFile)
for subdomain := range subdomains {
fmt.Fprintln(writer, subdomain)
}
writer.Flush()

if !*silent {
fmt.Printf("Output saved to %s\n", *outputPath)
}
} else {
for subdomain := range subdomains {
if !*silent {
fmt.Println(subdomain)
}
}
}

if *silent {
sublistPath := filepath.Join(os.Getenv("HOME"), "database", "sublist.txt")
sublistFile, err := os.OpenFile(sublistPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
fmt.Printf("Error opening sublist file: %v\n", err)
os.Exit(1)
}
defer sublistFile.Close()

writer := bufio.NewWriter(sublistFile)
for subdomain := range subdomains {
fmt.Fprintln(writer, subdomain)
}
writer.Flush()

if _, err := os.Stat(sublistPath); os.IsNotExist(err) {
fmt.Printf("Sublist file created at %s\n", sublistPath)
} else {
// Remove duplicates from sublist.txt
sublist, err := os.ReadFile(sublistPath)
if err != nil {
fmt.Printf("Error reading sublist file: %v\n", err)
os.Exit(1)
}
sublistLines := strings.Split(string(sublist), "\n")
sublistMap := make(map[string]bool)
for _, line := range sublistLines {
if line != "" {
sublistMap[line] = true
}
}
sublistFile.Truncate(0)
sublistFile.Seek(0, 0)
writer := bufio.NewWriter(sublistFile)
for subdomain := range sublistMap {
fmt.Fprintln(writer, subdomain)
}
writer.Flush()
if !*silent {
fmt.Printf("Output appended to %s\n", sublistPath)
}
}
}
}

