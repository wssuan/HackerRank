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
 * Complete the 'bigSorting' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY unsorted as parameter.
 */

func bigSorting(unsorted []string) []string {
    slices.SortFunc(unsorted, func(a, b string) int {
        if len(a) != len(b) {
            return len(a) - len(b)
        }
        if a<b {
            return -1
        } else {
            return 1
        }
    })
    return unsorted
    // Write your code here

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    var unsorted []string

    for i := 0; i < int(n); i++ {
        unsortedItem := readLine(reader)
        unsorted = append(unsorted, unsortedItem)
    }

    result := bigSorting(unsorted)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%s", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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
