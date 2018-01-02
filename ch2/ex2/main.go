package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/artgillespie/gopl/ch2/ex1"
)

// parse an array of strings into an array of float64
func convertToFloat(s []string) []float64 {
	var ret []float64
	for _, v := range s {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Fatalf("Couldn't convert %s to int: %v", v, err)
		}
		ret = append(ret, n)
	}
	return ret
}

// given an array of strings, return the conversion to fahrenheit
func conversionString(s []string) string {
	f := convertToFloat(s)
	buf := bytes.NewBufferString("Conversions: ")
	for _, v := range f {
		buf.Write([]byte(fmt.Sprintf("%s = %s ", ex1.Celsius(v), ex1.CToF(ex1.Celsius(v)))))
	}
	return strings.TrimSpace(buf.String())
}

func main() {
	if len(os.Args) == 1 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Error reading stdin: %v", err)
		}
		t := strings.TrimSpace(string(b))
		log.Println(conversionString(strings.Split(string(t), " ")))
		return
	}
	log.Println(conversionString(os.Args[1:]))
}
