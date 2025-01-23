package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'minimumBribes' function below.
 *
 * The function accepts INTEGER_ARRAY q as parameter.
 */

func minimumBribes(q []int32) {
    count := 0
    n := int32(len(q))
    for n > 2 {
        if q[n - 3] == n {
            count += 2
            q[n - 3], q[n - 2], q[n - 1] = q[n - 2], q[n - 1], n
        } else if q[n - 2] == n {
            count += 1
            q[n - 2], q[n - 1] = q[n - 1], n
        } else if q[n - 1] != n {
            fmt.Println("Too chaotic")
            return
        }
        q = q[:n - 1]
        n -= 1
    }
    if n == 2 && q[0] == n {
        count += 1
    }
    fmt.Printf("%d\n", count)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n := int32(nTemp)

        qTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        var q []int32

        for i := 0; i < int(n); i++ {
            qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
            checkError(err)
            qItem := int32(qItemTemp)
            q = append(q, qItem)
        }

        minimumBribes(q)
    }
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
