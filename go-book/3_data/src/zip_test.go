package data

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestWriteZip(t *testing.T) {
	buf := new(bytes.Buffer)

	// шинэ zip буфер үүсгэх
	w := zip.NewWriter(buf)

	// архивт нэмэх файлууд
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "Энэ файл нь заавар мэдээлэл агуулна"},
		{"todo.txt", "Энэ файл нь хийх зүйлсийн жагсаалтыг агуулна"},
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// буферыг хаах
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	// zip өгөгдлийг файлд бичих
	os.WriteFile("readme.zip", buf.Bytes(), 0777)
}

func TestReadZip(t *testing.T) {
	r, err := zip.OpenReader("readme.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Архив дахь файлуудаар давтаж агуулгыг хэвлэх
	for _, f := range r.File {
		fmt.Printf("'%s' файлын агуулга:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(os.Stdout, rc)
		if err != nil {
			log.Fatal(err)
		}
		println()
		rc.Close()
	}
}
