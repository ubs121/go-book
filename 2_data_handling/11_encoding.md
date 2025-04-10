# Өгөгдлийн хувиргалт

Бинари өгөгдлийг байт, битийн цуваа гэж үзэж болно. Бинари өгөгдлийг ихэвчлэн 8 битээр (заримдаа 7 битээр) кодлосон байдаг.

Юникод тесктээс байт хэлбэр рүү, байтаас текст хэлбэр рүү хувиргах үйлдлийг харгалзуулан энкодлох, декодлох гэж хэлдэг. Ихэвчлэн текстийг файлд хадгалах, сүлжээгээр дамжуулах зэрэгт энкодлолт хийх шаардлага гарна.

Дээхэн үед 8-бит өгөгдлийг илгээхэд ихээхэн асуудалтай байсан. Өгөгдөл замдаа эвдэрч ирэх магадлал өндөр байсан. Үүний оронд 7 бит өгөгдөл нь илүү найдвартай дамждаг байсан, учир нь сүүлийн 1 битийг шалгалтын цифр болгон ашиглаж алдаатай дамжсан эсэхийг мэдэж чаддаг байсан.

Гар утасны SMS, ASCII тэмдэгтийн хүснэгт зэрэг 7-битийн кодлолоор зохиогдсон зүйлс олон байдаг. Гэхдээ ихэнхи тохиолдолд 8 битийн бинари өгөгдөлтэй ажилладаг.

## UTF-8 энкодлолт

Текстийг дамжуулах, хадгалах зэрэгт олон төрлийн энкодлох (хувиргах) аргуудыг ашигладаг. Хамгийн түгээмэл арга нь UTF-8 юм. UTF-8 нь Юникод өгөгдлөөс (16 битээс ) илүүдэл байтуудыг танах замаар зай хэмнэж хувиргадаг. Тухайлбал латин тэмдэгтүүд (англи хэл дээрх текст) бүгд 8 битээр дүрслэгдэх боломжтой байдаг учраас 16 битээр хадгалах нь хадгалах төхөөрөмжийн хувьд үрэлгэн хандсан хэрэг юм.

Текст энкодлох, декодлох олон хэлбэр бий. Хамгийн өргөн хэрэглэгддэг нь:

1. UTF-8 (UCS Transformation Format, 8-бит хэлбэр) . Тэмдэгт бүр утгаасаа хамаараад 1 ээс 4 байтаар энкодлогддог. ASCII тэмдэгтүүд 1 байтад; `0x0080` ба `0x07ff` хоорондох тэмдэгтүүд 2 байтад; `0x0800`-ээс их тэмдэгтүүд 3 байтад хувирдаг. Орлуулагч нь 4 байтаар бичигддэг. Энкодлолтыг заагаагүй бол UTF-8 энкодлолт автоматаар сонгогддог.

2. UTF-16. Тэмдэгт бүр 2 байтаар дүрслэгддэг (орлуулагчаас бусад). Мөн Юникод энкодлолт гэж нэрлэгддэг.

3. ASCII. Тэмдэгт бүрийг 8-бит ASCII тэмдэгтээр кодлодог. Бүх тэмдэгтүүд ASCII  завсарт (`0x00` to `0x7F`) байж чадаж байгаа бол энэ кодлолтыг ашиглаж болно. Энэ засвараас гарсан тэмдэгт бичих үед тэмдэгтийн бага байтыг ашигладаг, их байт нь хаягдах эрсдэлтэй.

Энкодлолт болон декодлолт нь `unicode/utf8` сангаар хийгддэг. Энэ сан нь энкодлох, декодлох, зөв форматтай эсэхийг шалгах функцүүд агуулсан байдаг. Үүнд ASCII, UTF8, Unicode зэрэг багтана. Сүүлийнх нь UTF-16 энкодлолтод хэрэглэгддэг.

Энкодинг обект нь олон чухал методуудыг агуулдаг ба тэдгээрээс хамгийн чухал хоёр нь текстийг байт дараалал руу хувиргах `GetBytes`, байтуудыг текст рүү хувиргах `GetString` методууд юм.

```go
const nihongo = "日本語"
for i, w := 0, 0; i < len(nihongo); i += w {
  runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
  fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
  w = width
}
```

## Огноог текст рүү хувиргах

`time.Format` функц нь огноог текст рүү хувиргана. Энэ функц бусад програмчлалын хэлэнд байдаг форматлагч тэмдэгтүүдийг ашигладаггүй, өөр энгийн аргыг ашигладаг. Ингэхдээ бодит огноо утгыг форматад оруулан бичдэг.

Жишээлбэл:

```go
time.Now().Format("20060102150405")
```

Энэ командын үр дүнд одоогийн огноо `YYYYMMDDHHMMSS` форматаар хэвлэгдэх болно.  

Мөн урьдчилан бэлдсэн стандарт цагийн форматууд байдаг. Жишээлбэл `time.RFC3339`, `time.Kitchen` гэх мэт.

```go
time.Now().Format(time.Kitchen)
```

## JSON бүтцээр кодлох

JSON нь JavaScript Object Notation гэсэн үгний товчлол юм. JSON нь Javascript, вэб програмуудад өргөн ашиглагддаг, их энгийн бүтэцтэй, нүдээр уншигдах боломжтой, товч бичиглэлтэй өгөгдлийн бүтэц юм. Тийм учраас өгөгдөл хадгалах, дамжуулах, тохиргоо хадгалах, сүлжээгээр өгөгдөл тээвэрлэх, өгөгдлийн бүтэц тодорхойлох, мета хэл болгох гэх мэт олон төрлийн зориулалтаар өргөн ашигладаг.

Обект, массив, эгэл утга зэргийг JSON бүтэц уруу хялбар хувиргаж болно. Эгэл утганд тэмдэгт мөр, тоо, логик, `null` утгууд хамаарна. Массивын элементүүд нь таслалаар тусгаарлагдан JSON формат уруу хувирдаг. Массивын элементүүдийг "\[ ... \]" хаалтад хашиж бичнэ. Обектын хувьд "талбар: утга" байдлаар "{ ... }". хаалтад хашигдаж JSON уруу хувирна.

Жишээлбэл дараах програмд хүмүүсийн нэр, э-мэйл хаягийн мэдээллийг төлөөлөх обект, тэдгээрийн массивыг хэрхэн JSON бүтэц уруу хөрвүүлж файлд хадгалахыг харуулсан байна.

```go
package main

import (
  "encoding/json"
  "fmt"
  "os"
)

type Person struct {
  XMLName Name     `xml:"person" json:"person"`
  Name Name        `xml:"name"`
  Email []Email    `xml:"email"`
}

type Name struct {
  First string    `xml:"first" json:"firstName"`
  Last string     `xml:"last" json:"lastName"`
}

type Email struct {
  Type string     `xml:"type,attr"`
  Address string  `xml:",chardata"`
}

func getPerson() Person {
  // өгөгдөл
  person := Person{
  Name: Name{First: "Ууганбаяр", Last: "Сүхбаатар"},
  Email: []Email{
          {Kind: "хувийн", Address: "ub@gmail.com"},
          {Kind: "ажлын", Address: "ub@hotmail.com"},
        }
  }
  return person
}

func saveJSONSample() {
  outFile, err := os.Create("person.json")
  checkError(err)
  defer outFile.Close()

  encoder := json.NewEncoder(outFile)
  err = encoder.Encode(getPerson())
  checkError(err)
}

func checkError(err error) {
  if err != nil {
    fmt.Println("Алдаа ", err.Error())
    os.Exit(1)
  }
}
```

Дээрх програм нь үр дүнгээ `person.json` файлд хадгална. Энэ файлаас уншихдаа дараах функцийг ашиглаж болно.

```go
func loadJSON(fileName string, key interface{}) {
  inFile, err := os.Open("person.json")
  checkError(err)
  decoder := json.NewDecoder(inFile)
  err = decoder.Decode(key)
  checkError(err)
  inFile.Close()
}
```

## XML бүтцээр кодлох

XML нь мөн нийлмэл өгөгдлийг текстэн хэлбэрт хувиргахад маш өргөн ашиглагддаг бүтэц юм. Жишээлбэл Open Document, Microsoft Word файлууд, SVG зургийг XML бүтцээр кодлон файлд хадгалсан байдаг. Бидний сайн мэдэх HTML хэл нь XML дээр суурилсан, MathML болон CML (Chemistry Markup Language) зэрэг хэлүүд XML дээр суурилсан гэх мэтээр жишээ маш олныг нэрлэж болно. Мөн байгууллага хооронд мэдээлэл солилцоход өргөн ашиглагддаг веб үйлчилгээний SOAP протокол нь XML дээр суурилсан байдаг.

XML бүтэц нь таагуудаас тогтоно. Таагууд бие биенээ агуулж болно, тааг нь атрибуттай байж болно. Жишээлбэл бидний өмнө үзсэн `Person` обектыг XML бүтэц уруу хувиргавал дараах байдалтай харагдана.

Энд харуулснаар дараах дүрмээр хувиргалт хийнэ:

* талбарыг XML тааг болгон кодлох бол талбарын ард таагийн нэрийг \`\` хашилтад бичнэ
* талбарыг XML атрибут болгон кодлох бол `xml:tag,attr` бичиглэлийг ашиглана.
  бүтцийн нэрийг кодлох бол XMLName нэртэй талбар үүсгэж ард нь таагийн нэрийг тохируулна.

```go
package main

import (
  "encoding/xml"
  "fmt"
  "os"
)

// Person, Name, Email төрлүүд энд байрлана

func main() {
  str := `<?xml version="1.0" encoding="utf-8"?>
   <person>
    <name>
    <first>Ууганбаяр</first>
    <last>Сүхбаатар</last>
    </name>
    <email type="хувийн">ub@gmail.com</email>
    <email type="ажлын">ub@hotmail.com</email>
   </person>`

  var person Person
  err := xml.Unmarshal([]byte(str), &person)
  checkError(err)

  // обектыг талбаруудыг шалгах
  fmt.Println("Нэр: \"" + person.Name.First + "\"")
  fmt.Println("Э-мэйл 2: \"" + person.Email[1].Address + "\"")
}

func checkError(err error) {
  if err != nil {
    panic(err)
  }
}
```

## gob бүтцээр кодлох

Gob нь өгөгдлийг хувиргах, зөөвөрлөх, хадгалахад зориулагдсан бөгөөд бусад хэлэнд одоогоор ашиглах боломжгүй зөвхөн Go хэлэнд байдаг өгөгдөл хувиргалтын арга юм.

Gob нь суваг, функц, интерфэйсээс бусад Go хэлний бүх төрлийн өгөгдлийг хувиргаж чадна. Үүнд бүхэл тоон өгөгдөл, тэмдэг мөр, логик утга, struct, массив зэрэг багтана.

Gob нь өгөгдлийн төрлийн тухай мэдээллийг өгөгдөл дотор оруулж кодлодог. Энэ талаараа ASN бүтцээс их хэмжээтэй болдог, гэхдээ XML бүтэцтэй харьцуулахад бага хэмжээтэй байдаг. Өгөгдлийн төрлийн мэдээлэлд обектын төрөл, талбаруудын нэрс багтана.

Өгөгдлийн төрлийг багтааж кодлосноор өөрчлөлт орсон ч илүү тогвортой ажиллах баталгаатай болох юм. Тухайлбал обектын талбарын байрлал солигдсон ч нэрээр нь талбаруудыг зөв тайлж уншиж чадна.

Gob бүтцээр өгөгдөл кодлохын тулд `Encoder` обект үүсгэх хэрэгтэй. Энэ обектын `Encode()` функц нь өгөгдөл кодлоход ашиглагдана. Энэ функцийг хэдэн ч удаа дуудаж болно, энэ үед өгөгдлийн төрлийн талаархи мэдээлэл нь нэг л удаа бичигддэг.

Дараах програм нь хүмүүсийн мэдээллийг gob бүтэц уруу хувирган файлд бичнэ.

```go
func saveGobSample() {
  // файл үүсгэх
  outFile, err := os.Create("person.gob")
  checkError(err)
  defer outFile.Close()

  // энкодлогч үүсгэх
  encoder := gob.NewEncoder(outFile)
  // өгөгдлийг бичих
  err = encoder.Encode(getPerson())
  checkError(err)
}
```


## ASN бүтцээр кодлох

Abstract Syntax Notation One (ASN.1) нь анх 1984 онд харилцаа холбооны салбарт зориулан зохиосон бүтэц юм. Энэ бүтцийг одоо ч өргөн ашигладаг, тухайлбал үүрэн холбооны ярианы файлуудыг энэ бүтцээр хадгалдаг. ASN.1 бүтцийн давуу тал нь нийлмэл бүтэцтэй өгөгдлийг бага хэмжээтэй болгон бинари хэлбэрээр хадгалах, дамжуулахад тохиромжтой байдаг.

ASN.1 нь өөрөө нэлээд том стандарт бөгөөд түүний тодорхой хэсгийг `encoding/asn1` пакетад оруулсан байдаг. ASN.1 бүтцийг сан болгон оруулсан гол шалтгаан нь сүлжээний систем, нууцлалын сангуудад ASN бүтцийг ашигладагтай холбоотой. Тухайлбал X.509 сертификатыг тайлах, кодлох зэрэгт ASN.1 бүтцийг өргөн ашигладаг.

ASN.1 бүтэцтэй ажиллах хоёр функц бий.

```go
func Marshal(val interface{}) ([]byte, os.Error)
func Unmarshal(val interface{}, b []byte) (rest []byte, err os.Error)
```

Эхний функц нь өгөгдлийг ASN байт цуваа уруу хувиргана, дараагийн функц нь эсрэг үйлдэл гүйцэтгэнэ.

Хамгийн энгийн жишээ болгон бүхэл тоог ASN бүтцээр кодлох програм үзэе.

```go
package main

import (
  "encoding/asn1"
  "fmt"
  "os"
)

func main() {
  mdata, err := asn1.Marshal(123)
  checkError(err)
  var n int
  _, err1 := asn1.Unmarshal(mdata, &n)
  checkError(err1)
  fmt.Println("Задалж уншсан утга: ", n)
}

func checkError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Алдаа: %s", err.Error())
    os.Exit(1)
  }
}
```

ASN бүтцийн талаар нарийн судлахыг тулд түүний стандартыг сайтар судлах хэрэгтэй. Тухайлбал ASN стандартад өгөгдлийг төрөлжүүлсэн байдаг. Эдгээрийн талаар нарийн судлахыг хүсвэл Олон Улсын Харилцаа Холбооны Хорооны (ITU-T) стандартчилсан баримтуудыг үзээрэй.

## base64 энкодлолт

Бинари өгөгдлийг HTTP протоколоор дамжуулах үед текстэн хэлбэрт оруулах шаардлага их гардаг. Тухайлбал веб програмд зураг оруулахад түүнийг сервер уруу текст хэлбэрээр дамжуулах хэрэгтэй байж болно. Энэ үед зургийн байт бүрийг харгалзуулан текстээр кодлох хэрэгтэй болно.

Бинараас текст уруу хувиргах олон төрлийн боломж бий. Хамгийн түгээмэл аргуудын нэг нь Base64 формат ашиглах юм.

Дараах програм нь байт цувааг Base64 форматаар кодлож, буцаан уншиж харуулсан байна:

```go
package main

import (
  "bytes"
  "encoding/base64"
  "fmt"
)

func main() {
  data := []byte{1, 2, 3, 4, 5, 6, 7, 8}
  bb := &bytes.Buffer{}
  encoder := base64.NewEncoder(base64.StdEncoding, bb)
  encoder.Write(data)
  encoder.Close()

  fmt.Println(bb)

  dbuf := make([]byte, 12)
  decoder := base64.NewDecoder(base64.StdEncoding, bb)
  decoder.Read(dbuf)

  for _, ch := range dbuf {
    fmt.Print(ch)
  }
}
```



