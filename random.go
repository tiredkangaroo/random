package main

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

var (
	OutputModeHex    = "hex"
	OutputModeBase32 = "base32"
	OutputModeBase64 = "base64"
	OutputModeRaw    = "raw"
)

var OutputMode string
var OutputLength int
var OutputNewline bool
var RepeatCount int

func init() {
	flag.StringVar(&OutputMode, "output", OutputModeHex, "output encoding: hex, base32, base64, raw")
	flag.StringVar(&OutputMode, "o", OutputModeHex, "output encoding: hex, base32, base64, raw (shorthand)")

	flag.BoolVar(&OutputNewline, "newline", false, "add newline at the end of output")
	flag.BoolVar(&OutputNewline, "n", false, "add newline at the end of output (shorthand)")

	flag.IntVar(&OutputLength, "length", 16, "number of random bytes to generate")
	flag.IntVar(&OutputLength, "l", 16, "number of random bytes to generate (shorthand)")

	flag.IntVar(&RepeatCount, "repeat", 0, "number of times to repeat the output")
	flag.IntVar(&RepeatCount, "r", 0, "number of times to repeat the output (shorthand)")

	flag.Parse()
	switch OutputMode {
	case OutputModeHex, OutputModeBase32, OutputModeBase64, OutputModeRaw:
		// valid output mode
	default:
		fmt.Fprintf(os.Stderr, "invalid output mode (hex, base32, base64, raw): %s\n", OutputMode)
	}
	if OutputLength <= 0 {
		fmt.Fprintf(os.Stderr, "output length must be greater than zero: %d\n", OutputLength)
	}
	if RepeatCount < 0 {
		fmt.Fprintf(os.Stderr, "repeat count cannot be negative: %d\n", RepeatCount)
	}
}

func main() {
	for i := range RepeatCount + 1 {
		var buffer = make([]byte, OutputLength)
		rand.Read(buffer)
		var output string
		switch OutputMode {
		case OutputModeHex:
			output = hex.EncodeToString(buffer)
		case OutputModeBase32:
			output = base32.StdEncoding.EncodeToString(buffer)
		case OutputModeBase64:
			output = base64.StdEncoding.EncodeToString(buffer)
		case OutputModeRaw:
			output = string(buffer)
		}
		if i < RepeatCount {
			output += "\n"
		}
		fmt.Fprint(os.Stdout, output)
	}
	if OutputNewline {
		fmt.Fprint(os.Stdout, "\n")
	}
}
