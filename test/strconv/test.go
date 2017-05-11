package main

import (
	"time"
	"strconv"
	"math/rand"
	"fmt"
)

func GetRandomString(length int) string {
	str := "0123456789abcdefABCDEF"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}
func main() {
	//str := GetRandomString(16)
	str := 0xdaAEbA39B7EAF8Ff
	fmt.Println(str)
	//fmt.Println(int8(str))

	num, err := strconv.ParseInt(str,16,64)
	//s := strconv.FormatUint(num,16)
	//fmt.Println(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
	fmt.Printf("%v",num)
}