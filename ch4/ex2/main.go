package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var use384 = flag.Bool("sha384", false, "Use SHA384")
	var use512 = flag.Bool("sha512", false, "Use SHA512")
	flag.Parse()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Couldn't read stdin: %s", err)
	}
	var d []byte
	if *use384 {
		t := sha512.Sum384(b)
		d = t[:]
	} else if *use512 {
		t := sha512.Sum512(b)
		d = t[:]
	} else {
		t := sha256.Sum256(b)
		d = t[:]
	}
	log.Printf("%x\n", d)
}
