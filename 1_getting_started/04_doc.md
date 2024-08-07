# Тайлбар

Тайлбар нь програмын элементүүдийг тайлбарлахад хэрэглэгдэнэ. Хөрвүүлэгч нь тайлбарыг алгасаж хөрвүүлдэг, ө.х эцсийн машины кодонд тайлбар нь орохгүй, нөлөө үзүүлэхгүй гэсэн үг.

Тайлбарыг хоёр хэлбэрээр хийж болно:

* Нэг мөр тайлбар нь `//`-аар эхлэх ба мөрийн төгсгөл дээр төгсөнө
* Блок тайлбар нь `/*  тайлбар  */`  хэлбэртэй бичигдэнэ

```go
// Энгийн нэг мөр тайлбар

/* Дараагийн мөрийг тайлбарласан энгийн тайлбар */

/*
 * Кодыг тайлбарласан олон мөр тайлбар
 * Зузаан шрифтээр тэмдэглэж болохгүй ч
 * бид **чухал хэсэг** гэж бичиж болно.
 */
```

Тайлбар хэрхэн хийх талаар Go хэлний стандарт сангуудаас олон жишээ харж болно. Жишээ нь `fmt.Fprint()` функцийн тайлбарыг харая:

```go
// Fprint formats using the default formats for its operands and
// writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error
// encountered.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
```

Дээрх тайлбарын эхний өгүүлбэр нь функцийн нэрээр эхэлсэн байгааг ажиглаарай. Тайлбарын эхний өгүүлбэр нь чухал үүрэгтэй байдаг. Баримтыг текст, HTML, Unix man хуудас гэх мэт олон төрлийн форматаар үүсгэх үед янз бүрээр тайлбар нь таслагдах шаардлага гарна. Энэ үед эхний өгүүлбэр үлдэх магадлал өндөр байдаг.

Сайн тайлбар, баримтжуулалт нь програмыг ойлгомжтой, засвар үйлчилгээ хийхэд хялбар болгодог. Програмын чухал хэсгүүдэд тайлбар заавал хийж сурах хэрэгтэй. Тайлбар нь энгийн, уншихад ойлгомжтой, товч тодорхой байх ёстой.

Тайлбар нь дараах агуулгатай байж болно.

* Програм, кодын хэсэг юу хийдэг болох тухай товч тодорхойлолт
* Яах гэж энэ кодыг бичсэн, юу хийдэг болох тухай
* Кодыг ашиглах жишээ, товч заавар

Зарим хүн програмыг зөвхөн компютерт зориулсан зааврууд гэж ойлгодог. Энэ нь буруу ойлголт юм. Зөвхөн машинд зориулж програм бичдэг гэвэл өрөөсгөл зүйл болно. Тийм програмууд нь уншиж ойлгох, засварлахад маш хүнд байдаг бөгөөд заримдаа зохиогч нь өөрөө ойлгохгүй байх нь ч бий. 

Ер нь "Би энэ програмыг зөвхөн өөрөө өөртөө бичиж байна, тиймээс надад тайлбар хэрэггүй" гэж бодох хэрэггүй. Эхлээд код бичиж байх үед бүх зүйл тодорхой харагдах боловч түүнийгээ хожим дахин харахад ойлгомжгүй зүйлүүд их гарах болно.

## Тайлбараас баримт үүсгэх

Баримтыг програмаас тусад нь хөтөлж болох боловч код өөрчлөгдөх үед дагуулж өөрчлөх нь хүндрэлтэй байдаг. Тиймээс програмын кодон дотор тайлбараа хамт оруулж бичдэг болсон. Энэ нь нэг талаар хөгжүүлэгчдэд амар, нөгөө талаар баримтжуулалт хоцрохгүй ач холбогдолтой байдаг.

Сүүлийн үеийн програмчлалын хэлүүд бүгд кодон дотроос тайлбарыг түүвэрлэж баримт үүсгэдэг багажтай болсон. Go хэл энэ хандлагыг дагаад `godoc` нэртэй багажтай. Энэ багаж нь Go эх файлыг задлан шинжилж (тайлбар, жишээ зэргийг) HTML, текстэн баримт гаргаж авдаг. Энэ баримт нь кодтойгоо нягт холбоотой байдаг, жишээлбэл баримт дотроос ямар нэг функцийн дотоод кодыг сонирхож байвал дээр нь дараад л код руу дамжиж  болно.

Тайлбар ямар форматтай бичих талаар хатуу дүрэм байхгүй. Гэхдээ `Godoc` багажид ойлгуулахын тулд зарим жижиг дүрмийг баримтлах хэрэгтэй:

* Төрөл, хувьсагч, тогтмол, функц, пакет зэрэгт тайлбар хийхдээ аль болох ойр байрлуулах хэрэгтэй, дунд нь сул мөр үлдээж болохгүй. 
* Цувуулж бичсэн өгүүлбэрүүдийг нэг фараграф гэж ойлгодог. Хэрэв тусдаа фараграф болгохыг хүсвэл дунд нь хоосон мөр оруулах хэрэгтэй.
* Форматтай текст (жишээ код г.м) оруулах бол тайлбар текстээс дотогшоо доголдож бичих хэрэгтэй.

  ```go
    /*  Фибоначийн тоон цуваа
    *     f(n) = f(n-1)+f(n-2)
    *   Энд n>1, f(1)=1, f(2)=1 байна.
    */
  ```