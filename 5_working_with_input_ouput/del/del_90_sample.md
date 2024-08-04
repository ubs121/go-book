#  Жишээ дасгал

1. `os.Stat()` функц нь файлын тухай мэдээлэл буцаана. Энэ функцийг ашиглан файлын хэмжээг олох програм бичээрэй.

  ```go
  package main

  import (
    "os"
    "fmt"
  )

  func main() {
    fi, err := os.Stat("file_size.go")
    if err != nil {
      // файлын мэдээллийг авч чадсангүй
      os.Exit(1)
    }

    fmt.Printf("%s файлын хэмжээ %d байт\n", fi.Name(), fi.Size())
  }
  ```

3. Файлаас эхний 10 мөрийг таслан дэлгэцэнд харуулах `Head` нэртэй програм бичээрэй.

  ```go
  package main

  import (
  	"bufio"
  	"fmt"
  	"os"
  )

  func main() {
  	if len(os.Args) != 2 {
  		fmt.Fprintf(os.Stderr, "Файлын нэр заана уу\n")
  		os.Exit(1)
  	}

  	f, err := os.Open(os.Args[1])
  	if err != nil {
  		fmt.Printf("Файлыг нээхэд алдаа гарлаа: %v\n", err)
  		os.Exit(1)
  	}
  	defer f.Close()

  	n := 10
  	scanner := bufio.NewScanner(f)
  	for n > 0 && scanner.Scan() {
  		fmt.Println(scanner.Text())
  		n--
  	}
  }
  ```

1. TCP/IP протокол ашиглан сүлжээгээр харилцах клиент болон сервер програмууд бичээрэй.

   **Шийдэл**: Сервер талд `8088` порт дээр ажиллах програм бичие.  Сервер талд `Listener` обект үүсгэсний дараа `Accept()` методыг дуудаж клиент холбогдохыг хүлээнэ. Клиент холбогдоход түүнд мэдээлэл (жнь: өөрийнх нь IP хаяг) дамжуулаад холболтыг хаана.

   ```go
   // tcp_server.go
   package main

   import (
    "fmt"
    "net"
    "os"
   )

   func main() {
    service := ":8088"
    listener, err := net.Listen("tcp", service)
    checkError(err)

    for {
      conn, err := listener.Accept()
      if err != nil {
       continue
      }
      // клиентэд хариу илгээх
      conn.Write([]byte("Амжилттай холбогдлоо. Таны хаяг: " +
           conn.RemoteAddr().String() + "\n" ))
      conn.Close()
    }
   }

   //  checkError(err error) функцийг энд оруулна
   ```

   Сервер програмын ажиллагааг хурдан туршиж үзэх бол серверээ асаагаад nc програмаар хүсэлт илгээж болно:

   ```sh
   $ echo -n "Hello" | nc localhost 8088
   Амжилттай холбогдлоо. Таны хаяг: 127.0.0.1:48653
   ```

   Одоо дээрх сервер програмтай холбогдох клиент програмыг бичие.

   ```go
   // tcp_client.go
   package main

   import (
    "fmt"
    "net"
    "io"
   )

   func main() {
    tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:8088")
    checkError(err)

    // холболт тогтоох
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)

    // серверээс мэдээлэл унших
    result, err := io.ReadAll(conn)
    checkError(err)

    // авсан хариуг дэлгэцэнд хэвлэж харуулах
    fmt.Println(string(result))
   }

   //  checkError() функцийг энд оруулна
   ```

2. Вэбээс сервер рүү файл ачаалах програм бичээрэй.

   **Шийдэл**: Вэбээс ирсэн хүсэлтийг боловсруулахын тулд вэб функц үүсгэх хэрэгтэй. `Upload` нэртэй дараах вэб функц үүсгэе.

   ```go
   func Upload(w http.ResponseWriter, r *http.Request) {
   // 1. илгээсэн файлын агуулгыг унших

   // 2. уншсан өгөгдлөө файл руу бичих
   }
   ```

   a) Илгээсэн файлын агуулгыг `http.Request` обектын `Body` талбараас уншиж болно.

   ```go
   body, err = io.ReadAll(r.Body)
   ```

   b) Уншсан өгөгдлөө файл руу бичихэд `os.WriteFile()` функцийг ашиглаж болно.

   ```go
   err = os.WriteFile("upload.dat", body, 0700)
   ```

   Ингээд `Upload` функцээ http үйлчилгээнд бүртгэх хэрэгтэй.

   ```go
   http.HandleFunc("/", Upload)
   ```


   1. Өгөгдсөн хавтас дотор \*.html файлуудыг самнаж доторхи холбоосуудыг ( &lt;a&gt; таагийн href атрибутад заасан URL) хэвлэж харуулах програм бичээрэй. Хавтас самнах үйлдлийг хурдасгахын тулд файл унших үйлдлийг зэрэгцээ гүйцэтгээрэй.

   **Шийдэл**. Програмыг хавтас доторх html файлуудыг самнаж бүртгэх, олсон html файл дотор боловсруулалт хийх гэсэн хоёр дэд хэсэгт хувааж болно.

   Дараах нэртэй функцүүд үүсгэе.  
   `crawl` – хавтас, дэд хавтасаар дамжин самнаж html файл хайна  
   `findLink` – html файл дотор холбоос хайна

   Дээрх хоёр функцийг зэрэг ажиллуулвал илүү үр дүнтэй байна. Зэрэг ажиллах үедээ хоорондоо мэдээлэл солилцоход нь зориулж `taskChannel` нэртэй суваг үүсгэж болно. Энэ сувгаар боловсруулах шаардлагатай файлын нэрсийг солилцоно.

   Програмыг бүр илүү үр дүнтэй ажиллуулахын тулд файл дотроос холбоос хайх (`findLink`) функцийг 2 буюу түүнээс дээш тоогоор зэрэг ажиллуулж болно.

   Дараах жишээнд 3 зэрэг ажиллуулсан байна.

   ```go
   func main() {
      // эхлэх хавтас
      rootFolder := "web_root"

      // ажлын даалгавар (файлын нэрс) солилцох суваг үүсгэх
      taskChannel := make(chan string)

      // самнаж эхлэх (go функц үүсгэж байна)
      go crawl(rootFolder, taskChannel)

      // 3 зэрэг findLink функц ажиллуулах
      for i := 0; i < 3; i++ {
          go findLink(taskChannel)
      }

      // самналт дуусахыг хүлээх
      <- done
   }
   ```

   Програмын үндсэн шийдэл нь ингээд боллоо. `Crawl` болон `findLink` функцүүдийн бүтэн кодыг доор харуулав.

   ```go
   // хавтасаар самнах
   func crawl(folder string, ch chan string) {
      files, _ := os.ReadDir(folder)

      for _, f := range files {

          if f.IsDir() {
              // хавтас бол цааш самнах
              crawl(folder+"/"+f.Name(), ch)
          } else {
              //  *.html файл эсэхийг шалгах
              if strings.HasSuffix(f.Name(), fileExt) {
                  // мөн бол ажлын дараалалд оруулах
                  ch <- folder + "/" + f.Name()
              }

          }

      }
      done <- true
   }

   // html файл дотроос холбоос хайх функц
   func findLink(ch chan string) {
      for {
          select {
          case fileName := <-ch:
              fmt.Println(fileName)
              // энд холбоос хайх хэсэг байна
          }
      }
   }
   ```