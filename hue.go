// package hue is a command line client for interacting with Philips Hue lighting
package main

import (
	"bytes"
	"os"
)

func main() {
	os.Stdout.Write(bytes.NewBufferString("woohoo from hue").Bytes())
}
