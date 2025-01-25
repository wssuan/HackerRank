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
 * Complete the 'morganAndString' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING a
 *  2. STRING b
 */

type Stack struct {
    Data    []byte
    Dup     int
    Next    byte
}

func morganAndString(a string, b string) string {
    var pickStack func(a *Stack, b *Stack) *Stack
    pickStack = func(a *Stack, b *Stack) *Stack {
        // if one stack is empty, pick another
        if len(a.Data) == 0 {
            return b
        }
        if len(b.Data) == 0 {
            return a
        }
        // if stack tops are different, pick smaller one
        if a.Data[0] < b.Data[0] {
            return a
        } else if a.Data[0] > b.Data[0] {
            return b
        }
        // now stack tops are same
        if b.Dup == len(b.Data) {
            a, b = b, a
        }
        if b.Dup == len(b.Data) {
            // all data in both stacks is the same, pick random one
            return a
        }
        if a.Dup == len(a.Data) {
            // a has only one kind, three cases:
            // a=BBB, b=BBA... => pick b (BBA...BBB)
            // a=BBB, b=BBC... => pick a (BBBBBC...)
            if a.Data[0] < b.Next {
                return a
            } else {
                return b
            }
        }
        // now both stacks have more than one kind
        if b.Dup < a.Dup {
            a, b = b, a
        }
        if a.Dup < b.Dup {
            // two cases:
            // a=BBA..., b=BBB... => pick a (BBA...BBB...)
            // a=BBC..., b=BBB... => pick b (BBB...BBC...)
            if a.Next < a.Data[0] {
                return a
            } else {
                return b
            }
        }
        // both stacks have the same leading kind, and different next, then pick one with smaller next
        // a=CCA..., b=CCB... => pick a (CCA...CCB...)
        // a=CCA..., b=CCD... => pick a (CCA...CCD...)
        // a=CCD..., b=CCE... => pick a (CCD...CCE...)
        if a.Next < b.Next {
            return a
        } else if a.Next > b.Next {
            return b
        }
        // now two stacks have long shared prefix
        // optional quick bailout: if next kinds are bigger, pick random one
        // a=BBC..., b=BBC... => pick any (BBBBC...C...)
        if a.Next > a.Data[0] {
            return a
        }
        if len(a.Data) > len(b.Data) {
            a, b = b, a
        }
        for i := a.Dup; i < len(a.Data); i++ {
            if a.Data[i] == b.Data[i] {
                // optional quick bailout: if we hit something bigger than leading kind, pick random one, because all kinds before [i] must be consumed first anyway
                if a.Data[i] > a.Data[0] {
                    return a
                }
                continue
            }
            // Seems true for these
            if a.Data[i] < b.Data[i] {
                return a
            } else {
                return b
            }
        }
        if len(a.Data) == len(b.Data) {
            // totally identical, pick anyone
            return a
        }
        ab := &Stack{append(a.Data, b.Data...), a.Dup, a.Next}
        ba := &Stack{append(b.Data, a.Data...), b.Dup, b.Next}
        if pickStack(ab, ba) == ab {
            return a
        } else {
            return b
        }
    }
    evalStack := func(stack *Stack, data []byte) {
        stack.Data = data
        if len(data) > 0 {
            dup := 0
            for _, b := range data {
                if b != data[0] {
                    stack.Next = b
                    break
                }
                dup += 1
            }
            stack.Dup = dup
        }
    }
    var stacks [2]Stack
    evalStack(&stacks[0], []byte(a))
    evalStack(&stacks[1], []byte(b))
    selected := make([]byte, 0, len(stacks[0].Data) + len(stacks[1].Data))
    for len(stacks[0].Data) + len(stacks[1].Data) > 0 {
        pick := pickStack(&stacks[0], &stacks[1])
        // pick all repeated
        selected = append(selected, pick.Data[:pick.Dup]...)
        evalStack(pick, pick.Data[pick.Dup:])
    }
    return string(selected)
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
        a := readLine(reader)

        b := readLine(reader)

        result := morganAndString(a, b)

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
