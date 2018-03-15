package main

import (
	"elto/hello/pipeline"
	"fmt"
	"os"
	"bufio"
)

func main() {
	const filename = "large.in"
	const n = 64

	file , err := os.Create(filename)
	if err != nil {
		// 不知道怎么办了
		panic(err)
	}
	// 保证文件最后会被关闭
	defer file.Close()

	p := pipeline.RandomSource(n)

	// 使用buffer io，利用缓存，默认是 direct io
	bw := bufio.NewWriter(file)
	pipeline.WriterSink(bw, p)
	bw.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p = pipeline.ReadSource(bufio.NewReader(file), -1)
	for v := range p  {
		fmt.Println(v)
	}
}

func mergeDemo() {
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(4,2,3,6,8)),			pipeline.InMemSort(
			pipeline.ArraySource(7,4,0,2,8,13)))
	/*for {
		if num, ok := <- p; ok {
			fmt.Println(num)
		} else {
			break
		}
	}*/

	// 更简便的写法，底层自动通过锁实现同步
	count := 0;
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}
