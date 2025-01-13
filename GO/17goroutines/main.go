package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	//agar hum defer wg.done() nhi likhte to program terminate ho jayega 
	// or kuch esa error dega 
	// 	fatal error: all goroutines are asleep - deadlock!

	// goroutine 1 [semacquire]:
	// sync.runtime_Semacquire(0xc000008120?)
	//         C:/Program Files/Go/src/runtime/sema.go:71 +0x25
	// sync.(*WaitGroup).Wait(0xc0000160d0?)
	//         C:/Program Files/Go/src/sync/waitgroup.go:118 +0x48
	// main.main()
	//         C:/Users/mraja/Downloads/non-recycle image/codes/GO/17goroutines/main.go:31 +0x9d
	// exit status 2
	fmt.Println("task id:", id)
}

// func main() {
// 	for i := 0; i < 10; i++ {
// 		go task(i) // go keyword introduce multhreading in go

// 	}
// 	// agar hum timer nhi lagate to ye program terminate ho jayega before goroutines complete ho
// 	time.Sleep(time.Second * 2)
// }

// but hum log upper wala method use nhi karte 

func main() {
	var wg sync.WaitGroup
	for i:=0;i<=10;i++{
		wg.Add(1)
		// ander hum ek pointer pass karte hai task ke sath 
		go task(i, &wg)
	}

	wg.Wait()
}