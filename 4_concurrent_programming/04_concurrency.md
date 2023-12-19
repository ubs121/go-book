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

Параллел давталт ашигласнаар хугацаа маш их хожих болно. Гэхдээ давтаж байгаа зүйлс нь хоорондоо хамааралгүй байх хэрэгтэй.

Жишээ болгон тус бүр 200 миллисекундын ажил гүйцэтгэдэг хэд хэдэн go функц зэрэгцээ ажиллуулж нэг болон олон цөмтэй процессор дээр нийт ажиллах хугацааг нь харьцуулая. Цөмийн тоог тохируулахдаа `runtime.GOMAXPROCS()` функцийг ашиглаж болно. Энэ функцийг `main()` функцийн эхэнд дуудах хэрэгтэй.

200 миллисекундын зохиомол ажил (sleep функц ашиглан) гүйцэтгэх `doJob()` нэртэй функц үүсгэе. Энэ функцийг 60 удаа зэрэгцээ дуудаж ажиллуулая.

```go

type Request struct{ Name string }
type Response struct{ Data string }

func doJob(request *Request, ch chan *Response) {
	// 200 ms-н ажил
	time.Sleep(time.Millisecond * 200)

	// хариу илгээх
	ch <- &Response{Data: request.Name + " дууслаа"}
}

func main() {
	jobChan := make(chan *Response)

  // ажлын жагсаалт
  n := 60
	var requests []Request
	for i := 0; i < n; i++ {
		requests = append(requests, Request{"task" + strconv.Itoa(i)})
	}

	// параллел давталт
	for _, req := range requests {
		req1 := req
		go doJob(&req1, jobChan)
	}

	// бүх функцээс өгөгдөл хүлээх
	for i := 0; i < len(requests); i++ {
		resp := <-jobChan
		println(resp.Data)
	}
}
```

Сүүлийн давталтанд бүх сувгаас өгөгдөл ирж дуусахыг хүлээж байна. Бүх сувгаас мессеж ирсэн үед main() функцийн ажиллагаа дуусна.

Дээрх програмын ажилласан нийт хугацааг тодорхой гаргахын тулд Linux орчинд `time` програмыг ашиглаж болно.

```sh
$ time go run multicore.go
real	0m0.279s
user	0m0.071s
sys	0m0.012s
```

8 цөмтэй процессор дээр 0.279 секундын дотор 60 функц ажиллаж дууссан байна. 4 цөмтэй процессор дээр 0.325 секунд болж байсан. Ажиллах хугацааны хувьд компютерийн хүчин чадлаас хамаараад янз бүр байж болно, гол дүгнэлт бол процессорын цөмийн тоо ихсэх үед параллел давталтын нийт ажиллах хугацаа багасаж байгаа явдал юм.

## WaitGroup

Хэсэг go функц эхлүүлээд бүгдийг нь ажиллаж дуусахыг хүлээх хэрэгтэй бол `WaitGroup` ашиглаж болно.

```go
var wg sync.WaitGroup

for _, req := range requests {
  wg.Add(1) // группын тоолуурыг нэмэх
  go func() {
      // дууссан болохыг мэдэгдэх, defer хийх
      defer wg.Done() 

      // ажил хийх ...
      doJob(r, jobChan)
  }(r *Request)
}
// бүх go функц дуусахыг хүлээх
wg.Wait()
```

TODO: Ordered results from channel

```go
type ResultType struct {
	Order int
	Data  any
}

func TestOrderedResultsFromWorkers(t *testing.T) {
	n := 5 // number of workers

	// a work simulation
	DoWork := func(p int) ResultType {
		return ResultType{Order: p}
	}

	// create an array of channels
	resultChans := make([]chan ResultType, n)

	for i := 0; i < n; i++ {
		// create a channel for worker 'i'
		resultChans[i] = make(chan ResultType)

		// start a worker 'i'
		go func(order int) {
			result := DoWork(order)
			resultChans[order] <- result
		}(i)
	}

	// Collect and process results in order
	for i := 0; i < n; i++ {
		result := <-resultChans[i]
		// Process the result
		fmt.Println("result ", result)
	}
}

```

TODO: Worker pool pattern

TODO: Pipeline pattern