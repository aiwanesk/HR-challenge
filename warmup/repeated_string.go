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
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
    /*   countA := int64(0)
    for _, char := range s {
        if char == 'a' {
            countA++
        }
    }

    length := int64(len(s))
    fullRepeats := n / length
    totalA := fullRepeats * countA

    remainder := n % length
    for i := int64(0); i < remainder; i++ {
        if s[i] == 'a' {
            totalA++
        }
    }

    return totalA*/
    
       var A, B int64
    size := int64(len(s))

    for i := int64(0); i < size; i++ {
        if s[i] == 'a' {
            A++
            if i < n%size {
                B++
            }
        }
    }

    return n/size*A + B
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    result := repeatedString(s, n)

    fmt.Fprintf(writer, "%d\n", result)

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
