package main
import (
	"fmt"
	"sync"
)

var globalInt int
var lock sync.Mutex
var sum int
var wg sync.WaitGroup

type Address struct {
	street string
	number int
}

func (adr *Address) PPAddress2() {
	fmt.Printf("Street: %s and number: %d\n", adr.street, adr.number)
}

func PPAddress(adr Address) {
	fmt.Printf("Street: %s and number: %d\n", adr.street, adr.number)
}

func sumElements(arr []int) {
	wg.Add(1)
	localSum := 0
	for _, e := range arr {
		localSum = localSum + e
	}
	lock.Lock()
	sum = sum + localSum
	fmt.Println(sum)
	lock.Unlock()
	wg.Done()
}

func main() {
	adr := Address{
		street : "Ruberlei Boareto da Silva",
		number : 361,
	}
	PPAddress(adr)
	adr.PPAddress2()

	myArray := []int{}
	for i := 0; i < 1000000 ; i++ {
		myArray = append(myArray, i)
	}
	nParallel := 8
	sliceSize := len(myArray) / nParallel
	sum = 0
	for chunk := 0; chunk < nParallel; chunk++ {
		offset := chunk*sliceSize
		offsetEnd := (chunk + 1) * sliceSize
		payload := myArray[offset:offsetEnd]
		go sumElements(payload)
	}
	
	fmt.Println(sum)
	wg.Wait()
	fmt.Println(sum)
}