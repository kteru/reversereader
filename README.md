reversereader
=============

Package reversereader provides basic interfaces to read.  
It traverse an io.Reader as a backward stream.

[![GoDoc](https://godoc.org/github.com/kteru/reversereader?status.svg)](https://godoc.org/github.com/kteru/reversereader)

Installation
------------

```
$ go get -u github.com/kteru/reversereader
```

Example
-------

```
package main

import (
	"bytes"
	"io"
	"os"

	"github.com/kteru/reversereader"
)

func main() {
	readSeeker := bytes.NewReader([]byte("hello world"))

	revrd := reversereader.NewReader(readSeeker)

	io.Copy(os.Stdout, revrd)
	// => "dlrow olleh"
}
```
