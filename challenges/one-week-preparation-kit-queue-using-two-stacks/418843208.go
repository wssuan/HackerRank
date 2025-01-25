package main
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

type Stack struct {
    Data    []int
}

func NewStack() *Stack {
    return &Stack{make([]int, 0)}
}

func (this *Stack) Push(x int) {
    this.Data = append(this.Data, x)
}

func (this *Stack) Pop() int {
    x := this.Data[len(this.Data) - 1]
    this.Data = this.Data[:len(this.Data) - 1]
    return x
}

func (this *Stack) Peek() int {
    return this.Data[len(this.Data) - 1]
}

func (this *Stack) Size() int {
    return len(this.Data)
}

func processQueries(queries []string) []string {
    stacks := []*Stack{NewStack(), NewStack()}
    var output []string
    
    for _, query := range queries {
        switch query[0] {
        case '1':   // enqueue
            var x int
            fmt.Sscanf(query[2:], "%d", &x)
            for stacks[0].Size() > 0 {
                stacks[1].Push(stacks[0].Pop())
            }
            stacks[1].Push(x)
        case '2':   // dequeue
            for stacks[1].Size() > 0 {
                stacks[0].Push(stacks[1].Pop())
            }
            stacks[0].Pop()
        case '3':   // peek
            for stacks[1].Size() > 0 {
                stacks[0].Push(stacks[1].Pop())
            }
            output = append(output, fmt.Sprintf("%d", stacks[0].Peek()))
        }
    }
    return output
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int32(tTemp)
    queries := make([]string, 0, q)

    for range q {
        queries = append(queries, strings.TrimSpace(readLine(reader)))
    }
    for _, output := range processQueries(queries) {
        fmt.Fprintln(writer, output)
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
