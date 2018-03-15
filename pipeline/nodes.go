
package pipeline

import (
	"sort"
	"io"
	"encoding/binary"
	"math/rand"
	"time"
	"fmt"
)

/*
 * 输出为只读的channel
 */
func ArraySource(a ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _,v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

var starTime time.Time
func Init() {
	starTime = time.Now()
}

func InMemSort(in <- chan int) <- chan int {
	out := make(chan int)

	go func() {
		// read into memory
		a := []int{}
		for v := range in {
			a = append(a, v)
		}

		fmt.Println("Read into Memory done: ", time.Now().Sub(starTime))

		sort.Ints(a);

		fmt.Println("sort in memoey done: ", time.Now().Sub(starTime))

		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}


func Merge(in1 , in2 <- chan int) <- chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <- in1
		v2, ok2 := <- in2
		for ok1 || ok2  {
			if !ok2 || (ok1 && v1 <= v2){
				out <- v1
				v1, ok1 = <- in1
			} else {
				out <- v2
				v2, ok2 = <- in2
			}
		}

		close(out)
		fmt.Println("merge in memory done: ", time.Now().Sub(starTime))
	}()
	return out
}

func ReadSource(reader io.Reader, chunckSize int) <-chan int {
	out := make(chan int)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				bytesRead += n
				out <- v
			}
			if err != nil || chunckSize != -1 && bytesRead >= chunckSize {
				break
			}
		}
		close(out)
	}()

	return out
}

func WriterSink(writer io.Writer, in <- chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))

		writer.Write(buffer)
	}
}

func RandomSource(count int) <- chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}

		close(out)
	}()
	return out
}

func MergeN(inputs ...<- chan  int) <- chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}

	m := len(inputs) / 2
	// merger inputs [0 .. m) and inputs [m..end)
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))
}