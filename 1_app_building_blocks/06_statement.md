# Заавар

Зааврууд нь компютерт юу хийхийг зааварлах элементүүд юм. Бүх зааврууд нийлээд програм юу хийдгийг тодорхойлно.

Жишээ нь:

* хувьсагчид утга оноо
* функцэд параметр дамжуул
* функцэд удирдлагыг шилжүүл
* дараах заавар руу үсэр гэх мэт...

## Заавруудыг бүлэглэх

Хэд хэдэн заавруудыг угалзан хаалтаар `{ }` хашиж нэг бүлэг заавар болгодог. Үүнийг _блок заавар_ гэж хэлнэ. Блок заавар нь нэг заавар мэтээр програмын аль ч хэсэгт хэрэглэгдэж болно.

Жишээ нь:

```go
{
   var x = 0   // зарлалт
   x++         // заавар
   x = y + 5   // заавар
}
```

Блок дотор хувьсагч зарлах бол ихэвчлэн эхэнд нь бичдэг. Блок дотор тодорхойлогдсон хувьсагч нь тухайн блок дотор л хүчинтэй байдаг тухай бид өмнө үзсэн.

## Салаалах заавар

Салаалах заавар нь програмын биелэлтийг өгөгдсөн нөхцөлөөс хамааруулан өөр өөр замуудаар гүйцэтгэхэд хэрэглэгдэнэ. Go хэлэнд `if` ба `switch` гэсэн хоёр салаалах заавар байдаг. Функцийн дуудалт нь мөн нэг төрлийн салаалах заавар юм, гэхдээ нөхцөлгүйгээр програмын удирдлагыг өөр тийш салаалуулна.

### if заавар

`if` заавар нь дараах хэлбэртэй бичигдэнэ.

```go
if n % 2 == 0 {
  println("n бол тэгш тоо" )
} else {
  println("n бол сондгой тоо" )
}
```

Дээрх кодыг компютер гүйцэтгэхдээ эхлээд `n % 2 == 0` илэрхийллийг үнэлнэ. Хэрэв энэ илэрхийллийн утга үнэн бол эхний заавар биелэнэ, эсрэг тохиолдолд `else` блокт бичсэн заавар биелэнэ. `else` хэсэг нь шаардлагагүй бол орхиж болно.

### switch заавар

`switch` заавар нь `else if` цуваатай төстэй, гэхдээ илүү энгийн, уншихад хялбар байдаг.

Жишээ нь дараах өгөгдсөн сар хэд хоногтойг тооцоолох програмыг үзэе.

```go
package main
import "fmt"

var year int      /* он */
var month int     /* сар */

func main() {
 fmt.Print("Он, сар ? ")
 fmt.Scanf("%d %d", &year, &month)

 fmt.Print(year, month, " сар ")
 switch month {
  case 1,3,5,7,8,10,12:
    fmt.Println("31 хоногтой")
  case 4,6,9,11:
    fmt.Println( "30 хоногтой" )
  case 2: // өндөр жил эсэх
    if year == 400 || (year % 4 == 0 && year % 100 == 0) {
         fmt.Println( "29 хоногтой" )
    } else {
       fmt.Println( "28 хоногтой" )
      }
  default:
        fmt.Println( "Сарын дугаар буруу!" )
 }
}
```

`switch` доторх `month` утга `case`-н ард бичсэн утгатай тохирч байвал уг `case`-н ард бичигдсэн зааврыг гүйцэтгэнэ. Дээрх жишээнд гурван `case` байна, мөн эхний `case`-н ард `1,3,5,7,8,10,12` гэж олон утга бичсэн байна.

Хэрэв `month` утга аль ч `case` утгатай таарахгүй бол `default` –н ард бичигдсэн заавар биелэх болно. Ерөнхийдөө `default` нь орхиж болох боловч `switch` дотор бичиж занших нь зүйтэй байдаг.

## Давтах заавар

`for` давтах заавар нь хэсэг кодыг тодорхой тооны удаа эсвэл тодорхой нөхцөлөөр давтан гүйцэтгэх боломжийг олгоно. Тухайлбал ямар нэг зүйлийг тоолох, нийлбэрчлэх зэрэгт давтах шаардлага гарна.

Жишээ нь өгөгдсөн 5 тооны нийлбэр олоход давталтыг ашиглая:

```go
package main
import "fmt"

var (
  total int      /* нийт дүн */
  current int    /* хэрэглэгчээс  өгөх утга */
  counter int    /* давталтын тоолуур */
)

func main() {
    total = 0

    // 5 удаа давтаж тоо асуух
    for counter = 0; counter < 5; counter++ {
        print("Тоо? ")
        fmt.Scanf("%d", &current) // тоог оруулах

        total += current // total хувьсагчид нийлбэрийг хураах
    }

    // үр дүнг хэвлэх
    fmt.Printf("Нийт дүн %d\n", total)
}
```

Энд `counter` нь 0, 1, 2, 3, 4 гэж тоолж 5 удаа давтана. Давтах блок дотор хэрэглэгчээс тоо асуух, оруулсан тоог хувьсагчид нэмж хураах гэсэн хоёр зааврыг давтан гүйцэтгэж байна.

### Нөхцөлгүй давталт

Зарим тохиолдолд нөхцөлгүй давталтыг ашиглах шаардлага гардаг. Жишээлбэл, хэчнээн тоо хэрэглэгчээс оруулахыг мэдэхгүй байж болно. Энэ үед нөхцөлгүй давталт зохион байгуулж болно. Давталтыг дуусгахдаа `break` зааврыг ашиглаж болно.

```go
for {
   // энд break хэрэгтэй
}
```

`break` зааврыг мартах эсвэл буруу ашиглавал зарим тохиолдолд *төгсгөлгүй давталт* үүсдэг. Энэ тохиолдолд програм тасралтгүй ажиллаж CPU-д их хэмжээний ачаалал өгдөг. Ийм програмыг хүчээр зогсоохоос өөр аргагүй болдог.