# Интерфэйс төрөл

Интерфэйсийг ерөнхий утгаар нь олон зүйлтэй нийтлэг зарчмаар харилцах "үүд" гэж ойлгож болох юм. Төрлүүдийн дунд байдаг нийтлэг, төстэй "чадваруудыг" багцлан шинэ интерфэйс тодорхойлж болно. Интерфэйс нь бүтэц төрөлтэй нэлээд төстэй боловч талбар байхгүй зөвхөн метод агуулна.

Жишээлбэл `Rectangle`, `Circle` зэрэг геометр дүрсийг төлөөлөх төрлүүд байг. Эдгээрийн дунд талбай, переметр гэсэн нийтлэг ойлголтууд байдаг, тиймээс "талбайг тооцоолох",  "приметрийг тооцоолох" гэсэн нийтлэг үйлдлүүд байж болно. Тэгэхээр бид эдгээрийг ерөнхийлөөд `Shape` нэртэй дараах интерфэйс үүсгэж болох юм.

```go
type Shape interface {
    area() float64
    perimeter() float64
}
```

`Rectangle`, `Circle` төрлүүд нийтлэг `Shape` интерфэйстэй болохын тулд `area()`, `perimeter()` методуудыг хэрэгжүүлэх хэрэгтэй. Ингэхдээ методуудыг өөр өөрсдийн хувилбараар хэрэгжүүлнэ, тухайлбал талбай, периметр тооцох аргууд дүрс бүрийн хувьд өөр өөр байна.

```go
/*   Тэгш өнцөгт */
func (r *Rectangle) area() float64 {
 // тэгш өнцөгтийн талбай тооцоолох
}

func (r *Rectangle) perimeter() float64 {
 // тэгш өнцөгтийн периметр тооцоолох
}

/*   Тойрог */
func (c *Circle) area() float64 {
 // тойргийн талбай тооцоолох
}

func (c *Circle) perimeter() float64 {
 // тойргийн перитемр тооцоолох
}
```

Нэг ширхэг обекттой интерфэйсээр дамжуулан харилцах нь ач холбогдол багатай байдаг. Харин олон төрлийн, олон обекттой нэг интерфэйсээр харилцах үед интерфэйсийн ач холбогдол харагдана. Жишээлбэл дараах `totalArea()` функц нь `Shape` интерфэйстэй обектуудыг хүлээн аваад бүх дүрсийн нийт талбайг олж байна:

```go
// энэ функц нь олон төрлийн дүрсийн олонлог хүлээн авна
func totalArea(shapes ...Shape) float64 {
    var area float64

    for _, s := range shapes {
        area += s.area() // тайлбайг тооцоолж хувьсагчид хураах
    }
    return area
}
```

Гурван дүрсийн талбайг олохын тулд функцийг дараах байдлаар дуудаж болно:

```go
fmt.Println( totalArea(&circle1, &circle2, &rect1) )
```

`totalArea()` функцийн `area += s.area()` мөрд хамгийн гол "нууц" нуугдаж байгаа юм. Ө.х `Shape` интерфэйсээр дамжуулан ямар ч дүрсийн талбайг тооцоолж чадаж байна. `s.area()` дуудалт бүрийн цаана дүрсүүд өөр өөрийн аргаар талбайг тооцоолно. 

Үүнийг _полиморфизм_ гэж нэрлэдэг. Полимофирзм нь нэг төрлийн зүйлсийн олон төрхтэй байхыг хэлдэг. Энэ нь обектууд нэг ижил нөхцөлд ялгаатай хариу үйлдэл хийх байдлаар илэрдэг.