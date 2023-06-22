package main

import (
	"io"
	"log"
	"os"

	"github.com/masahiro331/go-ext4-filesystem/ext4"
)

func main() {
	f, err := os.Open("filesystem.ext4")
	if err != nil {
		log.Fatal(err)
	}
	info, _ := f.Stat()
	filesystem, err := ext4.NewFS(io.NewSectionReader(f, 0, info.Size()), nil)
	if err != nil {
		log.Fatal(err)
	}
	print(filesystem)
}