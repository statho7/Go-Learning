package main

import (
	"bytes"
	"fmt"
	"io"
)

// func main() {
// 	var w Writer = ConsoleWriter{}
// 	w.Write([]byte("Hello Go!"))
// }

// type Writer interface {
// 	Write([]byte) (int, error)
// }

// type ConsoleWriter struct {}

// func(cw ConsoleWriter) Write(data []byte) (int, error){
// 	n, err := fmt.Println(string(data))
// 	return n, err
// }

// func main(){
// 	myInt := IntCounter(0)
// 	var inc Incrementer = &myInt
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(inc.Increment())
// 	}
// }

// type Incrementer interface{
// 	Increment() int
// }

// type IntCounter int

// func (ic *IntCounter) Increment() int{
// 	*ic++
// 	return int(*ic)
// }

func main() {
	// // var wc WriterCloser = NewBufferedWriterCloser()
	// // wc.Write([]byte("Hello Youtube listeners, this is a test"))
	// // wc.Close()
	// var wc WriterCloser = NewBufferedWriterCloser()
	// wc.Write([]byte("Hello Youtube listeners, this is a test"))
	// wc.Close()

	// // bwc := wc.(*BufferedWriterCloser)
	// // bwc := wc.(io.Reader)
	// // fmt.Println(bwc)

	// // r, ok := wc.(io.Reader)
	// // r, ok := wc.(*BufferedWriterCloser)
	// r, ok := wc.(BufferedWriterCloser)
	// if ok {
	// 	fmt.Println(r)
	// } else {
	// 	fmt.Println("Conversion failed")
	// }

	var myObj interface{} = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("Hello Youtube listeners, this is a test"))
		wc.Close()
	}
	r, ok := myObj.(io.Reader)
	// r, ok := wc.(*BufferedWriterCloser)
	// r, ok := wc.(BufferedWriterCloser)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func(bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}