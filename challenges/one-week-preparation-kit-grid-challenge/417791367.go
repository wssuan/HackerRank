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
 * Complete the 'gridChallenge' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func gridChallenge(grid []string) string {
    latest := []byte(grid[0])
    slices.SortFunc(latest, func(a, b byte) int { return int(a) - int(b) })
    for _, s := range grid[1:] {
        curr := []byte(s)
        slices.SortFunc(curr, func(a, b byte) int { return int(a) - int(b) })
        for i := range curr {
            if latest[i] > curr[i] {
                return "NO"
            }
        }
        latest = curr
    }
    return "YES"
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
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n := int32(nTemp)

        var grid []string

        for i := 0; i < int(n); i++ {
            gridItem := readLine(reader)
            grid = append(grid, gridItem)
        }

        result := gridChallenge(grid)

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
