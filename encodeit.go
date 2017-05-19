package main

import (
	"bufio"
	"encoding/base32"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

// Integer codes for the encoding constants
const (
	Base32 = iota
	Base64
)

// Error definition
var (
	ErrInvalidEncoding = errors.New("encoding not specified")
	ErrFile            = errors.New("invalid file")
)

func main() {

	var err error

	// flags
	encPtr := flag.String("encoding", "base64", "encoding format (base64 (default), base32)")
	srcPtr := flag.String("source", "", "source for the input value (defaults to StdIn)")

	flag.Parse()

	// encoding
	var enc int
	switch *encPtr {
	case "base64":
		enc = Base64
	case "base32":
		enc = Base32
	default:
		enc = Base64
	}

	// source
	var input string
	var source io.Reader

	switch *srcPtr {
	case "":
		source = os.Stdin
	default:
		source, err = os.Open(*srcPtr)
		if err != nil {
			fmt.Printf("Aborting with error %s", err)
			return
		}
	}

	// tailing args, has precedence over -source flag
	l := len(flag.Args())
	switch {
	case l > 1:
		fmt.Println("Ambiguous, please enclose sentence inside \"'s")
		return
	case l < 1:
		input, err = fetchInput(source)
		if err != nil {
			fmt.Printf("Aborting with error %s", err)
			return
		}
	default:
		input = flag.Arg(0)
	}

	res, err := encodeValue(enc, []byte(input))
	if err != nil {
		fmt.Printf("Aborting with error %s", err)
		return
	}
	fmt.Printf("Result: %s", res)

	fmt.Println()
}

// encodeValue encodes the inputed value using the encoding paramter.
// Defaults to an error.
func encodeValue(base int, value []byte) (string, error) {
	var err error
	var res string

	switch base {
	case Base32:
		res = base32.StdEncoding.EncodeToString(value)
	case Base64:
		res = base64.StdEncoding.EncodeToString(value)
	default:
		err = ErrInvalidEncoding
	}

	return res, err
}

// fetchInput scans text from the io.Reader (can be an os.File (os.Stdin for
// standard input), or insert a -source=filename so that 1st line is read).
// Scan strips the \n and avoids io.EOF (https://golang.org/pkg/io/#pkg-variables)
// Note: only the fist line of a multi line file is used here. Its easy to extend
// for other needs.
func fetchInput(f io.Reader) (string, error) {
	if f == os.Stdin {
		fmt.Println("Insert a value to encode:")
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	return scanner.Text(), scanner.Err()
}
