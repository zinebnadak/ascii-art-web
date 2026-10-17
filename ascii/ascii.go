package main

import (
    "fmt"
    "os"
    "strings"
)
// function CALCULATES THE STARTING POINT
func Forumla(TheRune rune) int {
    return (int(TheRune)-32)*9 + 1
}
// function PRINTS EACH ROW OF THE LETTERS ITEM BY ITEM THEN LETTER BY LETTER FROM THE LIST
func printRow(words []string, LinesToChooseFrom []string) {
    for _, word := range words { // each ITEM in the list
        if word == "" { // if the item is empty, print a blank line and continue to next item
            fmt.Println()
            continue
        }
        for row := 0; row < 8; row++ { // 8 rows
            line := ""                    // start the row empty
            for _, letter := range word { // each LETTER in the item
                line += LinesToChooseFrom[Forumla(letter)+row] // where the letter starts & how far down we are,
            }
            fmt.Println(line)
        }
    }
}
// MAIN FUNCTION
func main() {
    // CHECK ARGS
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run . \"your text\"")
        return
    }
    // READ FILE
    contentFromSample, err := os.ReadFile("standard.txt") // always need the name in double quotes
    // ERROR HANDLING
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    LinesToChooseFrom := strings.Split(string(contentFromSample), "\n")
    // TAKE USERS INPUT AND START DRAWING
    input := os.Args[1]
    if input == "" {
        return
    }
    // VALIDATION BEFORE PRINTING ANYTHING
    // VALIDATE CHARACTERS PASSED
    for _, letter := range input {
        if letter < 32 || letter > 126 {
            fmt.Println("Error: unsupported character")
            return
        }
    }
    word := strings.Split(input, "\\n")
    // EDGE CASE: IF CHARACHTERS ARE ONLY NEWLINES PRINT ONE LINE FEWER
    if strings.ReplaceAll(input, "\\n", "") == "" {
        for i := 0; i < len(word)-1; i++ {
            fmt.Println()
        }
        return
    }
    printRow(word, LinesToChooseFrom)
}