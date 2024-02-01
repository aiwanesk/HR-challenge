package warmup

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	//"sort"
)

/*
* Complete the 'sockMerchant' function below.
*
* The function is expected to return an INTEGER.
* The function accepts following parameters:
*  1. INTEGER n
*  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
	/* Naive O(nlogn) + O(n)
	   	    ret := int32(0)

	          sort.Slice(ar, func(i, j int) bool { return ar[i] < ar[j] })

	          for i := 0; i < len(ar) - 1; i++{
	              if (ar[i] == ar[i + 1]){
	                  ret++
	                  i++
	              }
	          }

	          return ret
	*/

	/* O(n^2)
	    socketsMap := make(map[int32]int32)

	    var ret int32

	    for _, a := range ar{
	        if _, found:= socketsMap[a]; !found {
	            socketsMap[a] = 1
	        }else{
	            socketsMap[a]++
	            if socketsMap[a] % 2 == 0{
	                ret++
	            }
	        }
	    }

	    var ret int32

	    for _, v := range socketsMap{
	        ret += v / 2
	    }

	return ret
	*/

	socketMap := make(map[int32]int32)

	var ret int32
	//o(n)
	for _, a := range ar {
		if _, found := socketMap[a]; !found {
			socketMap[a] = 1
		} else {
			socketMap[a]++
			if socketMap[a]%2 == 0 {
				ret++
			}
		}
	}

	return ret
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := sockMerchant(n, ar)

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
