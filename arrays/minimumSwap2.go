package arrays

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

var v [100003][]int32
var visit [100003]bool

func dfs(i int32) int32{
    visit[i] = true
    z := int32(1)

    for _, x := range v[i] {
        if !visit[x] {
            z += dfs(x)
        }
    }

    return z
}

func minimumSwaps(A []int32) int32 {
    for i := range v {
        v[i] = nil
    }
    for i := range visit {
        visit[i] = false
    }

    for i := int32(0); i < int32(len(A)); i++ {
        v[i] = append(v[i], A[i]-1)
        v[A[i]-1] = append(v[A[i]-1], i)
    }

    c := int32(0)
    for i := int32(0); i < int32(len(A)); i++ {
        if !visit[i] {
            c += dfs(i) - 1
        }
    }

    return c
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    res := minimumSwaps(arr)

    fmt.Fprintf(writer, "%d\n", res)

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
