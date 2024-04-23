package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ForkReader(r io.Reader, w1, w2 io.Writer) error {
	i := 0
	b := make([]byte, 1)
	for {
		_, err := r.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if i%2 == 0 {
			_, err = w1.Write(b)
			if err != nil {
				return err
			}
		} else {
			_, err = w2.Write(b)
			if err != nil {
				return err
			}
		}
		i += 1
	}
	return nil
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	str := readLine(reader)

	r := strings.NewReader(str)
	rw1, rw2 := bytes.NewBuffer([]byte{}), bytes.NewBuffer([]byte{})
	err = ForkReader(r, rw1, rw2)
	if err != nil {
		panic(err)
	}
	res1, err := ioutil.ReadAll(rw1)
	if err != nil {
		panic(err)
	}
	res2, err := ioutil.ReadAll(rw2)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(writer, "%s\n", string(res1))
	fmt.Fprintf(writer, "%s\n", string(res2))

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
