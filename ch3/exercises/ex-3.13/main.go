package main

type kiloByte float64

const (
	_   kiloByte = iota             // discard the first value of iota
	KiB          = 1 << (10 * iota) // 1024^1
	MiB          = 1 << (10 * iota) // 1024^2
	GiB          = 1 << (10 * iota) // 1024^3
)

/*
NOTE: Is it Valid,
Need to learn about "bits and their arthematic"
*/

func main() {
	// fmt.Printf("1 KiB = %v kB\n", KiB/1000)
	// Convert from bytes to kilobytes (1000 bytes = 1 kB in metric system)
}
