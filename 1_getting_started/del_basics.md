## Арифметик үйлдэл

| Үйлдэл | Утга | Жишээ |
| --- | --- | --- |
| \* | Үржих | x \* y |
| / | Хуваах | x / y |
| % | Хуваасан үлдэгдэл олох | x % y |
| + | Нэмэх | x + y |
| - | Хасах | x - y |
| + \(унар\) | Эерэг тэмдэг | +x |
| - \(унар\) | Сөрөг тэмдэг | -x |

Эдгээр үйлдлээс зөвхөн `％` үйлдэл нь бүхэл тоон операнд шаардана. Бусад үйлдлийг нь дурын тоон төрөл дээр хэрэглэж болно.

## Утга олгох үйлдэл

Олгох үйлдэл нь түүний баруун талд байгаа илэрхийллийг тооцоолоод үр дүнг зүүн талын обектод онооно. Зүүн талын операнд нь өөрчлөгдөх боломжтой обект \(хувьсагч гэж нэрлэдэг\) байна.

| Үйлдэл | Утга | Жишээ | Үр дүн |
| --- | --- | --- | --- |
| = | Утга олгох | x = y | х-д y утгыг онооно |
| := | Зарлаад олгох | x:=10 | x хувьсагчийг шинээр зарлаад 10 утга онооно |
| +=, -=, \*=, /=, %=, &=, ^=, \|=, &lt;&lt;=, &gt;&gt;= | Нийлмэл олгох | x \*= y | `x oper= y` нь `x = x oper (y)` –тэй ижил |

## Нэмэгдүүлэх, хорогдуулах

| Үйлдэл | Утга | Нөлөө | Илэрхийллийн үр дүн |
| --- | --- | --- | --- |
| x++ | Нэмэх | x = x + 1 | x –н утгыг 1-ээр нэмнэ |
| x-- | Хасах | x = x - 1 | x –н утгыг 1-ээр хорогдуулна |

## Харьцуулах

Харьцуулах үйлдэл нь хоёр операндыг харьцуулаад `true` \(үнэн\) эсвэл  `false` \(худал\) утгыг гаргана.

| Үйлдэл | Утга | Жишээ | Үр дүн |
| --- | --- | --- | --- |
| &lt; | Бага | x &lt; y | хэрэв x нь y-ээс бага бол `true`, бусад үед  `false` |
| &lt;= | Бага буюу тэнцүү | x &lt;= y | хэрэв x нь y-тэй тэнцүү буюу бага бол `true`, эсрэг тохиолдолд `false` |
| &gt; | Их | x &gt; y | x нь y-ээс их бол `true`, эсрэг тохиолдолд `false` |
| &gt;= | Их буюу тэнцүү | x &gt;= y | хэрэв x нь y-тэй тэнцүү буюу их бол `true`, эсрэг тохиолдолд `false` |
| == | Тэнцүү | x == y | x нь y-тэй тэнцүү бол `true`, эсрэг тохиолдолд `false` |
| != | Тэнцүү биш | x != y | хэрэв x нь y-тэй тэнцүү биш бол `true`, эсрэг тохиолдолд  `false` |

## Логик үйлдэл

Логик үйлдлүүдийг холбон нийлмэл логик илэрхийлэл зохиож болно. Логик илэрхийлэл нь ихэвчлэн үсэргэх болон заавруудын дарааллыг өөрчлөхөд ашиглагддаг. Go хэлэнд дараах логик үйлдлүүдийг ашигладаг.

| Үйлдэл | Утга | Жишээ | Үр дүн |
| --- | --- | --- | --- |
| && | Логик AND | x && y | x ба y нь хоёулаа `true` үед `true`, бусад үед `false` |
| \|\| | Логик OR | x \|\| y | x ба y нь хоёулаа `false` үед `false`, бусад үед `true` |
| ! | Логик NOT | !x | x нь `false` бол `true`, эсрэг тохиолдолд `false` |

Жишээ нь `deviation` утга `[-0.2, 0.2]` завсараас гадна орших эсэхийг хоёр хэлбэрээр шалгаж болно: `(deviation <  -0.2) || (deviation >  0.2)` эсвэл `!(deviation >= -0.2  &&  deviation <= 0.2)`

## Бит үйлдлүүд

Бит үйлдлүүд нь байт доторхи битүүдтэй “нэг нэгээр” нь ажиллах боломж олгоно. Тухайлбал битүүдийг цэвэрлэх, сэргээх, бүлгээр нь инверс хийх гэх мэт. Мөн битүүдийн байрлалыг шилжүүлэн хөдөлгөж болно. Бит үйлдэл ашигладаг хамгийн энгийн жишээ бол Linux үйлдлийн системийн файл руу хандах зөвшөөрлийн бүтэц юм.

Битүүд нь баруунаасаа зүүн тийш 0-ээс эхлэн дугаарлагдана. Жишээлбэл, `*` тэмдэгтийн бит бүтцийг авч үзэе. Энэ тэмдэгтийн ASCII код нь 42 бөгөөд хоёртын системд 101010 болох юм. 101010 утгыг 8 битээр дүрслэхийн тулд урд нь 00  залгах хэрэгтэй.

```
Бит бүтэц      0 0 1 0 1 0 1 0
Бит дугаар     7 6 5 4 3 2 1 0
```

Go хэлний бит үйлдлүүдийг дараах хүснэгтэд харуулав.

| Үйлдэл | Утга | Жишээ | Үр дүн \(битээр\) |
| --- | --- | --- | --- |
| & | Бит AND | x & y | 1&1=1 , 1&0=0, 0&0=0 |
| \| | Бит OR | x \| y | 1\|1=1, 1\|0=1, 0\|0=0 |
| ^ | Бит XOR | x ^ y | 1^1=1, 1^0=0, 0^1=0, 0^0=1 |
| ~ | Бит NOT \(нэгийн гүйцээлт\) | ~x | ~1=0, ~0=1 |

Бит үйлдлийн операнд нь бүхэл тоон төрөл байх ба үр дүн нь мөн бүхэл тоо байна.

Жишээ болгон `a:=6`, `b:=11` утгуудын хооронд бит үйлдэл хийе.

```tex
 a & b
```

Бүхэл тоон утгын тодорхой битүүдийг бит AND үйлдэл ашиглан 0-ээр цэвэрлэж болно. Цэвэрлэхэд ашиглаж байгаа битүүдийг бит маск гэж хэлдэг. Жишээ нь, `0xFF` бит маскаар AND үйлдэл хийвэл тооны бага найман битээс бусад бүх битийг цэвэрлэнэ.

```go
a &= 0xFF;       // a = a & 0xFF; бичиглэлтэй адил
```

Энэ жишээнд харж байгаагаар нийлмэл олгох `&=` үйлдлийг мөн ашиглаж болж байна. Мөн бусад бит үйлдлүүдийг олгох үйлдэлтэй нийлүүлэн хэрэглэж болно.

Бит үйлдлүүд нь мөн дараагийн бит үйлдэлд зориулан бит маск бэлдэхэд ашиглагдана. Жишээ нь, `0x20` бит бүтцэд зөвхөн 5-р бит нь 1 байна. Тиймээс `~0x20` илэрхийллийг 5-р битээс бусад битийг 1 болгох маск болгон ашиглаж болж байна:

```go
a &= ~0x20;    // 5-р битийг цэвэрлэх
```

`~0x20` бит маскийг `0xFFFFFFDF` гэж ашиглах нь илүү дээр юм, учир нь энэ нь машины бүтцээс хамааралгүй хүссэн үр дүнг гарган авахад илүү тохиромжтой, мөн хүн уншихад илүү ойлгомжтой байна.

Мөн `|` \(OR\) ба `^` \(exclusive OR\) үйлдлүүдийг тодорхой битүүдийг сэргээх, цэвэрлэхэд ашиглаж болно.

Жишээлбэл:

```go
int mask = 0xC;  // битээр 00001100
a |= mask;       // 2 ба 3 дугаартай битүүдийг сэргээх
a ^= mask;       // 2 ба 3 дугаартай битүүдийг инверслэх
```

Нэгэн ижил бит маск ашигласан хоёр дахь инверс нь эхний утгатай адил үр дүнг гаргана. Ө.х, `b^mask^mask` нь `b` хувьсагчийн анхны утгыг гаргана. Энэ чанарыг ашиглан хоёр бүхэл тооны утгуудыг гуравдагч хувьсагч ашиглалгүй хооронд нь сольж болно:

```go
a ^= b;          // a = a ^ b;
b ^= a;          // b –д a-н анхны утгыг оноох
a ^= b;          // a –д b-н анхны утгыг оноох
```

Энэ жишээний эхний хоёр илэрхийлэл нь `b = b^(a^b)` эсвэл  `b = (a^b)^b` илэрхийлэлтэй ижил юм. Үр дүн нь `b = a` болно. Энэ үед гуравдахь илэрхийлэл \(`a` ба `b` –ийн анхны утгуудыг ашиглан\) нь `a = (a^b)^a` буюу `a = b` болгоно.

## Хаяглах үйлдлүүд

Обектийн санах ой дахь хаягийг тодорхойлох, массивын элемент индекслэх, бүтцийн гишүүнд хандах зэрэг нь хаяглах үйлдэл юм.

| Үйлдэл | Утга | Жишээ | Үр дүн |
| --- | --- | --- | --- |
| & | Хаяг авах | &x | х-г заах хаяг |
| \[\] | Индекслэх | x\[y\] | x  массивын y байрлал дахь элемент |
| . | Бүтцийн гишүүнд хандах | x.y | x бүтцийн y нэртэй гишүүн |

## Бусад үйлдлүүд

Дээрх ангиллуудын  алинд ч орохооргүй дараах хэдэн үйлдэл байна.

| Үйлдэл | Утга | Жишээ | Үр дүн |
| --- | --- | --- | --- |
| \( \) | Функц дуудах | log\(x\) | Заасан функцэд  удирдлагыг шилжүүлнэ |
| len\(\) | Обектийн хэмжээ олох | len\(x\) | х-н санах ойд эзлэх хэмжээ, урт |
| make\(\) | Обект эсвэл обектуудын олонлог үүсгэх | make\(\[\]int, 10, 10\) | 10 урттай бүхэл тоон массив үүсгэнэ |
| new\(\) | Шинэ обектод санах ой хувиарлах | new\(T\) | T төрлийн хэмжээтэй санах ой хувиарлаад түүний хаягийг буцаана |
| append\(\) | Хоёр массивыг залгах | append\(s0, 2\) | s0 массив дээр 2 утгыг нэмж залгана |
| copy\(\) | Массивыг хуулах | copy\(b, "Hello"\) | b массив руу Hello текстийг хуулна |
| delete\(\) | Майпаас элемент хасах | delete\(m, k\) | m майпаас k түлхүүртэй өгөгдлийг устгана |
| .\(type name\) | Төрөл хувиргах | x.\(string\) | х-г string болгон хөрвүүлнэ |