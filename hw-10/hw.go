package hw10

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

var from, to string
var limit, offset int

func init() {
	flag.StringVar(&from, "from", "", "file from")
	flag.StringVar(&to, "to", "", "file to")
	flag.IntVar(&limit, "limit", 0, "limit")
	flag.IntVar(&offset, "offset", 0, "offset")
}

func Copy(from, to string, limit, offset int) error {
	file_from, err := os.OpenFile(from, os.O_RDWR, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("no such file")
		}
	}
	defer file_from.Close()
	file_to, _ := os.Create(to)

	defer file_to.Close()

	_, err = file_from.Seek(int64(offset), 0)
	if err == io.EOF {
		fmt.Printf("%s", err)
	}

	buf := make([]byte, 1024*1024)
	offset = 0
	if limit == 0 {
		limit = math.MaxUint32
	}
	fmt.Printf("Alredy read %d bytes\n", offset)
	for offset < limit {
		read, err := file_from.Read(buf[offset:])
		offset += read
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panicf("failed to read: %v", err)
		}
		fmt.Printf("Alredy read %d bytes\n", offset)
	}

	_, err = file_to.Write(buf[:offset])
	if err != nil {
		log.Panicf("failed to write: %v", err)
	}
	return nil
}

func main() {
	flag.Parse()
	err := Copy(from, to, limit, offset)
	fmt.Printf("%s\n", err)
}
