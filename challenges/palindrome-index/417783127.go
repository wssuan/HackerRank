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
 * Complete the 'palindromeIndex' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func palindromeIndex(s string) int32 {
    boundary := len(s)
    sleft := boundary / 2
    sright := sleft + 1
    if boundary % 2 == 0 {
        sleft -= 1
    }
    // additional check: if s is already palindrome, return -1
    left, right := sleft, sright - 1
    for left >= 0 {
        if s[left] != s[right] {
            break
        }
        left -= 1
        right += 1
    }
    if left < 0 {
        return -1
    }
    found := -1
    // assume the redundancy is in the left part
    left, right = sleft, sright
    for right < boundary {
        if s[left] != s[right] {
            if found != -1 {
                break
            }
            found = left
        } else {
            right += 1
        }
        left -= 1
    }
    if found == -1 {
        return 0    // left must be 0
    } else if left < 0 {
        // redundancy is found
        return int32(found)
    }
    // multiple mismatch found, the redundancy is not in the left part
    // check the right part
    found = -1
    left, right = sleft - 1, sright - 1
    for left >= 0 {
        if s[left] != s[right] {
            if found != -1 {
                break
            }
            found = right
        } else {
            left -= 1
        }
        right += 1
    }
    if found == -1 {
        return int32(right) // right must be boundary - 1
    } else if right == boundary {
        // redundancy is found
        return int32(found)
    }
    return -1
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        s := readLine(reader)

        result := palindromeIndex(s)

        fmt.Fprintf(writer, "%d\n", result)
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
