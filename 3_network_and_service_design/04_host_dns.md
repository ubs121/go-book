# Хоост нэр, DNS

IP буюу тоон хаягыг хүн тогтооход хэцүү учраас хоост нэр ашигладаг. Хоост нэр нь нэг буюу хоёр үетэй энгийн үг байдаг. Компютерийн сүлжээнд хоост нэрийг тоон хаягтай холбосон бүртгэл байдаг. Энэ бүртгэлийг Domain Name System (DNS) гэж нэрлэдэг.

DNS-г гар утасны бүртгэлтэй төстэй гэж ойлгож болно. Бид хүмүүсийн нэрний цаана гар утасны дугаарыг хадгалдаг бөгөөд ихэвчлэн нэрээр нь хэрэглэдэг, харин дугаарыг нь төдийлөн санадаггүй.

DNS бүртгэл нь олон сервер дээр тархан байрласан байдаг бөгөөд хамгийн ойр байх серверээс хайлтыг эхлэдэг. Эцэст нь хайлтын үрд дүнд олдсон тоон хаягаар компютерүүд холбогддог.

<img src="res/dns_lookup.drawio.svg"/>

`net.LookupHost()` функц нь өгөгдсөн хоост нэрийн цаана холбосон IP хаягийг олоход тусладаг. Зарим тохиолдолд хоост нэрийн цаана олон ялгаатай IP хаяг байж болно. Өөрөөр хэлбэл тухайн компютер олон сүлжээ карттай тохиолдолд ийм байж болно.

Дараах жижиг програм нь хоост нэрээр IP хаягийг хэрхэн олж болохыг харуулсан байна.

```go
// lookup_host.go
package main

import (
  "net"
  "os"
  "fmt"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "host нэр заана уу\n")
        os.Exit(1)
    }
    name := os.Args[1]
    addrs, err := net.LookupHost(name)

    if err != nil {
        fmt.Println("Error: ", err.Error())
        os.Exit(2)
    }

    for _, s := range addrs {
        fmt.Println(s)
    }
    os.Exit(0)
}
```

Програмын үр дүн:

```sh
$ go run lookup_host.go google.com
122.201.16.249
122.201.16.251
122.201.16.208
122.201.16.212
122.201.16.216
122.201.16.218
122.201.16.219
```



