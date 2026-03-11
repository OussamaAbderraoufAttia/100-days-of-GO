package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(string)
}

// ReadWriter embeds both Reader and Writer
type ReadWriter interface {
	Reader
	Writer
}

type Device struct {
	Name string
}

func (d Device) Read() string {
	return "Data from " + d.Name
}

func (d Device) Write(s string) {
	fmt.Printf("Writing '%s' to %s\n", s, d.Name)
}

func Process(rw ReadWriter) {
	fmt.Println(rw.Read())
	rw.Write("processed data")
}

func main() {
	d := Device{Name: "HardDrive"}
	Process(d)
}
