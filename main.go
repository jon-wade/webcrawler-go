package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func traverse(n *html.Node) {
	//fmt.Printf("%+v\n", n)
	//fmt.Println("type", n.Type)
	//fmt.Println("attr", n.Attr)
	//fmt.Println("data", n.Data)
	//fmt.Println("atom", n.DataAtom)
	//fmt.Println("")
	if n.Type == html.ElementNode && n.Data == "link" {
		// Do something with n...
		fmt.Println(n.Attr)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c)
	}
}

func parsePage(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			os.Exit(0)
		}
	}()

	doc, err := html.Parse(res.Body)
	traverse(doc)
}

func main() {
	parsePage("https://bbc.co.uk/news")
}

//func lrt(r chan int32, done chan bool) {
//	i := 0
//	for {
//		i++
//		time.Sleep(time.Second * 1)
//		r <- rand.Int31n(100)
//		if i == 6 {
//			close(r)
//			done <- true
//			break
//		}
//	}
//}

//func main() {
//	//parsePage("https://bbc.co.uk/news")
//	r := <-longRunningTask()
//	fmt.Println(r)
//}

//func worker(resultSlice *[]int32, wg *sync.WaitGroup, mutex *sync.Mutex) {
//	defer wg.Done()
//	mutex.Lock()
//	*resultSlice = append(*resultSlice, rand.Int31n(100))
//	mutex.Unlock()
//}
//
//func main() {
//	var resultSlice []int32
//	var wg sync.WaitGroup
//	var mutex = sync.Mutex{}
//
//	for i := 1; i <= 5; i++ {
//		wg.Add(1)
//		go worker(&resultSlice, &wg, &mutex)
//	}
//
//	wg.Wait()
//	fmt.Println("resultSlice", resultSlice);
//}

//func main() {
//	resultChan := make(chan int32)
//	done := make(chan bool)
//	var resultSlice []int32
//
//	go lrt(resultChan, done)
//
//	for {
//		result, more := <-resultChan
//		if more {
//			exists := false
//			for _, v := range resultSlice {
//				fmt.Println("v", v)
//				if v == result {
//					exists = true
//					break
//				}
//			}
//			if !exists {
//				resultSlice = append(resultSlice, result)
//			}
//			fmt.Println(resultSlice)
//			continue
//		}
//		fmt.Println("done")
//		break
//	}
//
//	<-done
//}