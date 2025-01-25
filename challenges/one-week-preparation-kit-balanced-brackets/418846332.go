package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "slices"
)

/*
 * Complete the 'isBalanced' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isBalanced(s string) string {
    opening := []rune{'(', '[', '{'}
    closing := []rune{')', ']', '}'}
    stack := make([]rune, 0, len(s))
    for _, c := range s {
        open := slices.Index(opening, c)
        if open >= 0 {
            stack = append(stack, c)
            continue
        }
        close := slices.Index(closing, c)
        if close >= 0 {
            if len(stack) == 0 {
                return "NO"
            }
            pair := stack[len(stack) - 1]
            if slices.Index(opening, pair) != close {
                return "NO"
            }
            stack = stack[:len(stack) - 1]
        }
    }
    if len(stack) == 0 {
        return "YES"
    } else {
        return "NO"
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        s := readLine(reader)

        result := isBalanced(s)

        fmt.Fprintf(writer, "%s\n", result)
    }

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
