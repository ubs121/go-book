## Параллел давталт

Массив зэрэг өгөгдлийн олонлог дээр нэгэн зэрэг боловсруулалт хийж хугацаа хожих шаардлага олон гардаг. C#, Fortress зэрэг бусад хэлэнд параллел давталт гүйцэтгэхэд зориулсан тусгай заавар байдаг, харин Go хэлэнд үүнийг тусгай зааваргүйгээр хялбархан хийж болно. Дараах хэлбэрийн `for` давталтыг go функцтэй хослуулан хэрэглэж болно:

```go
for i, v := range data {
  go func (i int, v float64) {
		 doSomething(i, v)
		 ...
  } (i, v)
}
```

Параллел давталт ашигласнаар нийт ажлын хугацааг багасгаж болно. Гэхдээ давтаж байгаа зүйлс нь хоорондоо хамааралгүй байх хэрэгтэй.

Жишээ болгон тус бүр 200 миллисекундын ажил гүйцэтгэдэг хэд хэдэн go функц зэрэгцээ ажиллуулж нэг болон олон цөмтэй процессор дээр нийт ажиллах хугацааг нь харьцуулая. 

Эхлээд 200 миллисекундын зохиомол ажил (Sleep функц ашиглан) гүйцэтгэх `doJob()` нэртэй функц үүсгэе.

```go
type Request struct{ Name string }
type Response struct{ Data string }

// ажил гүйцэтгэх
func doJob(request *Request) *Response {
	time.Sleep(200 * time.Millisecond) // 200 ms-н ажил

	// хариу илгээх
	return &Response{Data: request.Name + " дууслаа"}
}
```

Энэ функцийг 60 удаа зэрэгцээ дуудаж ажиллуулая.

```go
func parallelFor() {
	n := 60
	jobChan := make(chan *Response)

	// ажлын жагсаалт үүсгэх
	var requests []Request
	for i := 0; i < n; i++ {
		requests = append(requests, Request{"task" + strconv.Itoa(i)})
	}

	// параллел давталт
	for _, req := range requests {
		go func(r *Request) {
			reply := doJob(r)
			jobChan <- reply
		}(&req)
	}

	// бүх функцээс өгөгдөл хүлээх
	for i := 0; i < len(requests); i++ {
		resp := <-jobChan
		println(resp.Data)
	}
}
```

Сүүлийн давталтанд бүх сувгаас өгөгдөл ирж дуусахыг хүлээж байна. Бүх сувгаас мессеж ирсэн үед parallelFor() функцийн ажиллагаа дуусна.

Дээрх програмын ажилласан нийт хугацааг тодорхой гаргахын тулд Linux орчинд `time` програмыг ашиглаж болно.

```sh
$ time go run multicore.go
real	0m0.279s
user	0m0.071s
sys	0m0.012s
```

8 цөмтэй процессор дээр 0.279 секундын дотор 60 функц ажиллаж дууссан байна. 4 цөмтэй процессор дээр 0.325 секунд болж байсан. Ажиллах хугацааны хувьд компютерийн хүчин чадлаас хамаараад янз бүр байж болно, гол дүгнэлт бол процессорын цөмийн тоо ихсэх үед параллел давталтын нийт ажиллах хугацаа багасаж байгаа явдал юм. 

Цөмийн тоог тохируулахдаа `runtime.GOMAXPROCS()` функцийг ашиглаж болно. Энэ функцийг ихэвчлэн `main()` функцийн эхэнд дууддаг.

## Ordered results from channel

```go
func orderedResults(n int) {
	// сувгийн слайс үүсгэх
	results := make([]chan *Response, n)

	for i := 0; i < n; i++ {
		//  'i' ажилд зориулж суваг үүсгэх
		results[i] = make(chan *Response)

		// шинэ ажил эхлүүлэх
		go func(order int) {
			reply := doJob(&Request{Name: "task" + strconv.Itoa(order)})
			results[order] <- reply
		}(i)
	}

	// үр дүнг дарааллаар нь цуглуулах
	for i := 0; i < n; i++ {
		result := <-results[i]
		fmt.Println("result ", result)
	}
}
```

## WaitGroup

Зарим тохиолдолд хэдэн go функц хэрэгтэй болохыг мэдэх боломжгүй байдаг. Энэ тохиолдолд `WaitGroup` обект ашиглаж болно. Хэсэг go функц эхлүүлээд бүгдийг нь ажиллаж дуусахыг хүлээх хэрэгтэй бол `WaitGroup` ашиглаж болно.

```go
var wg sync.WaitGroup

for _, req := range requests {
  wg.Add(1) // группын тоолуурыг нэмэх

  go func(r *Request) {
      defer wg.Done()  // дууссан болохыг мэдэгдэх, defer хийх

	  reply:=doJob(r) // ажил хийх
	  jobChan<-reply  // хариу илгээх
  }(&req)
}

// бүх go функц дуусахыг хүлээх
wg.Wait()
```

TODO: Worker pool pattern

TODO: Pipeline pattern