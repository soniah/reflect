package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// Ding Ding - Shau Kei Wan, HKU, Kennedy Town, Admiralty.
func dingding() {
	fmt.Printf("\n%s\n\n", "-- dingding() -- interfaces --")

	const scratch = "/var/tmp/scratch"
	var err error
	var n int
	buf := make([]byte, 20)

	cmd := exec.Command("cp", "-f", "/etc/hosts", scratch)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// --------------------

	var foo io.ReadWriter
	foo, err = os.OpenFile(scratch, os.O_RDWR, 0)
	if err != nil {
		log.Fatal(err)
	}

	// foo (io.ReaderWriter) can be used for reading
	n, err = foo.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("foo: read %d bytes: %s\n\n", n, string(buf))

	// foo (io.ReaderWriter) can be used for writing
	n, err = foo.Write(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("foo: wrote %d bytes\n\n", n)

	// but foo can't be used for anything else (even though the concrete value
	// inside is a *File) - io.ReaderWriter only reads and writes.
	// "foo.WriteAt undefined (type io.ReadWriter has no field or method WriteAt)"

	// --------------------

	// assert foo to an io.Reader and it can be used for reading
	r := foo.(io.Reader)
	n, err = r.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("r: read %d bytes: %s\n\n", n, string(buf))

	// but it can't be used for writing
	// "r.Write undefined (type io.Reader has no field or method Write)"
	// n, err = r.Write(buf)

	// --------------------

	// assert foo to an io.Writer and it can be used for writing
	w := foo.(io.Writer)
	n, err = w.Write(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("w: wrote %d bytes\n\n", n)

	// but it can't be used for reading
	// "w.Read undefined (type io.Writer has no field or method Read)"
	// n, err = w.Read(buf)

	// or anything else, an io.Writer only writes
	// "w.Sync undefined (type io.Writer has no field or method Sync)"
	// err = w.Sync()

}
