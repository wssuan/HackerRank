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
 * Complete the 'quadrants' function below.
 *
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY p
 *  2. STRING_ARRAY queries
 */

type QCount [4]int
func (this *QCount) add(a *QCount) *QCount {
    this[0] += a[0]
    this[1] += a[1]
    this[2] += a[2]
    this[3] += a[3]
    return this
}
func (this *QCount) sub(a *QCount) *QCount {
    this[0] -= a[0]
    this[1] -= a[1]
    this[2] -= a[2]
    this[3] -= a[3]
    return this
}
func (this *QCount) clone() *QCount {
    return &QCount{this[0], this[1], this[2], this[3]}
}
func (this *QCount) transform(seed byte) *QCount {
    this[seed], this[1 ^ seed], this[2 ^ seed], this[3 ^ seed] = this[0], this[1], this[2], this[3]
    return this
}

func quadrants(p [][]int32, queries []string) {
    accumulate := make([]QCount, len(p) + 1)
    curr := &accumulate[len(p)]
    for i, coord := range p {
        accumulate[i] = *curr
        var q byte
        if coord[0] > 0 {
            if coord[1] > 0 {
                q = 0   // first
            } else {
                q = 1   // fourth
            }
        } else {
            if coord[1] > 0 {
                q = 2   // second
            } else {
                q = 3   // third
            }
        }
        curr[q] += 1
    }
    
    type LookupNode struct {
        Begin, End  int
        Counts      QCount  // cache of sum of subsets, before transformation
        Transform   byte    // current transformation from super point of view
        Big, Small  *LookupNode
    }
    
    var transform func(node *LookupNode, begin, end int, seed byte) *QCount
    transform = func(node *LookupNode, begin, end int, seed byte) *QCount {
        if begin >= node.End || end <= node.Begin {
            // out of scope
            return &QCount{}
        }
        
        if begin <= node.Begin && end >= node.End {
            // full covered
            ori := node.Counts.clone().transform(node.Transform)
            node.Transform ^= seed
            return node.Counts.clone().transform(node.Transform).sub(ori)
        }
        
        // partial covered, go deeper
        if node.Big == nil {
            // split
            mid := (node.Begin + node.End) / 2
            node.Small = &LookupNode{node.Begin, mid, *accumulate[mid].clone().sub(&accumulate[node.Begin]), 0, nil, nil}
            node.Big = &LookupNode{mid, node.End, *accumulate[node.End].clone().sub(&accumulate[mid]), 0, nil, nil}
        }
        // go deeper
        d0, d1 := transform(node.Big, begin, end, seed), transform(node.Small, begin, end, seed)
        d := d0.clone().add(d1)
        // update myself
        node.Counts.add(d)
        return d.transform(node.Transform)
    }
    var calculate func(node *LookupNode, begin, end int) *QCount
    calculate = func(node *LookupNode, begin, end int) *QCount {
        if begin >= node.End || end <= node.Begin {
            // out of scope
            return &QCount{}
        }
        if begin <= node.Begin && end >= node.End {
            // full covered
            return node.Counts.clone().transform(node.Transform)
        }
        // partial covered, go deeper
        if node.Big == nil {
            // split
            mid := (node.Begin + node.End) / 2
            node.Small = &LookupNode{node.Begin, mid, *accumulate[mid].clone().sub(&accumulate[node.Begin]), 0, nil, nil}
            node.Big = &LookupNode{mid, node.End, *accumulate[node.End].clone().sub(&accumulate[mid]), 0, nil, nil}
        }
        // go deeper
        c0, c1 := calculate(node.Big, begin, end), calculate(node.Small, begin, end)
        c := c0.clone().add(c1).transform(node.Transform)
        return c
    }
    
    //xmirror := []byte{1, 0, 3, 2}
    //ymirror := []byte{2, 3, 0, 1}
    lookupTree := &LookupNode{0, len(p), accumulate[len(p)], 0, nil, nil}
    for _, query := range queries {
        var i, j int
        fmt.Sscanf(query[2:], "%d %d", &i, &j)
        i -= 1  // turn i to 0-based index
                // keep j as exclusive 0-based index
        
        switch query[0] {
        case 'X':
            transform(lookupTree, i, j, 1)
        case 'Y':
            transform(lookupTree, i, j, 2)
        case 'C':
            counts := calculate(lookupTree, i, j)
            fmt.Printf("%d %d %d %d\n", counts[0], counts[2], counts[3], counts[1])
        }
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    var p [][]int32
    for i := 0; i < int(n); i++ {
        pRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var pRow []int32
        for _, pRowItem := range pRowTemp {
            pItemTemp, err := strconv.ParseInt(pRowItem, 10, 64)
            checkError(err)
            pItem := int32(pItemTemp)
            pRow = append(pRow, pItem)
        }

        if len(pRow) != 2 {
            panic("Bad input")
        }

        p = append(p, pRow)
    }

    qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int32(qTemp)

    var queries []string

    for i := 0; i < int(q); i++ {
        queriesItem := readLine(reader)
        queries = append(queries, queriesItem)
    }

    quadrants(p, queries)
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
