package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/artgillespie/gopl/ch2/ex1"
	"github.com/artgillespie/gopl/ch2/ex2/lengthconv"
	"github.com/artgillespie/gopl/ch2/ex2/weightconv"
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
	buf := bytes.NewBufferString("")
	tab := tabwriter.NewWriter(buf, 0, 0, 4, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(tab, "Value\tTemperature (F)\tLength (M)\tWeight (Kg)\t")
	for _, v := range f {
		fmt.Fprintf(tab, "%.2f\t%.2f\t%.2f\t%.2f\t\n", v, ex1.CToF(ex1.Celsius(v)), lengthconv.FToM(v), weightconv.LToK(v))
	}
	tab.Flush()
	return buf.String()
}

func main() {
	if len(os.Args) == 1 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Error reading stdin: %v", err)
		}
		t := strings.TrimSpace(string(b))
		fmt.Println(conversionString(strings.Split(string(t), " ")))
		return
	}
	fmt.Println(conversionString(os.Args[1:]))
}
