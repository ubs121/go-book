# Өгөгдлийг шахах

Ихэнхи вэб, файл серверүүд өгөгдлийг ямар нэг хэлбэрээр шахаж илгээдэг. HTML, Javascript, CSS зэрэг текстэн өгөгдлийг шахаж илгээх нь оновчтой байдаг, харин png зэрэг зургуудын хувьд анхнаасаа шахалттай байдаг.

Өгөгдлийг шахаж илгээснээр сүлжээгээр дамжих өгөгдлийг их хэмжээгээр багасгана, үүний үр дүнд сүлжээгээр нэвтрэх хурд нэмэгдэнэ. Шахах үйлдэл нь CPU-д бага зэрэг ачаалал өгдөг, гэхдээ gzip, snappy, lz4 зэрэг сайн алгоритмын хувьд энэ нь бараг нөлөөгүй байдаг.

Gzip нь хамгийн түгээмэл шахалтын алгоритм юм. Apache, Nginx зэрэг ихэнхи вэб серверүүд gzip шахалтыг дэмждэг. Мөн сүүлийн үеийн бүх браузер gzip дэмжинэ.

`archive/zip` пакет нь ZIP форматаар өгөгдлийг шахах, задлах боломжийг олгоно. Жишээлбэл дараах програмд `readme.txt`, `todo.txt` файлуудыг динамикаар үүсгэж `readme.zip` нэртэй архив файл уруу шахаж байна.

```go
package main

import (
  "archive/zip"
  "bytes"
  "os"
  "log"
)

func main() {
    // шинэ zip буфер үүсгэх
    buf := new(bytes.Buffer)
    w := zip.NewWriter(buf)

    // архивт нэмэх файлууд
    var files = []struct {
        Name, Body string
    } {
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
```

ZIP файлыг задалж унших програмыг доор харуулав:

```go
package main 

import (
  "archive/zip"
  "log"
  "fmt"
  "io"
  "os"
)

func main() {
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

        // дэлгэц рүү хуулах
        _, err = io.Copy(os.Stdout, rc)
        if err != nil {
            log.Fatal(err)
        }
        println()
        rc.Close()
    }
}
```



