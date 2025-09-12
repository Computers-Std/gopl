package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var SHA384 = flag.Bool("sha384", false, "Use SHA-384")
var SHA512 = flag.Bool("sha512", false, "Use SHA-512")

func main() {
	flag.Parse()
	var args []string

	if *SHA512 || *SHA384 {
		args = os.Args[2:]
	} else {
		args = os.Args[1:]
	}

	for _, arg := range args {
		// Declare sha with specific types
		var sha256Hash [32]byte
		var sha384Hash [48]byte
		var sha512Hash [64]byte

		switch {
		case *SHA384:
			sha384Hash = sha512.Sum384([]byte(arg))
			fmt.Printf("%x %T\n", sha384Hash, sha384Hash)
		case *SHA512:
			sha512Hash = sha512.Sum512([]byte(arg))
			fmt.Printf("%x %T\n", sha512Hash, sha512Hash)
		default:
			sha256Hash = sha256.Sum256([]byte(arg))
			fmt.Printf("%x %T\n", sha256Hash, sha256Hash)
		}
	}
}

// NOTE: need a revision
