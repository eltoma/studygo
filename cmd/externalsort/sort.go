package main

import (
	"os"
	"elto/hello/pipeline"
	"bufio"
	"fmt"
)

func main() {
	p := createPipeline("large.in", 800000000, 4)
	writeToFile(p, "large.out")
	printFile("large.out", 100)
}

func createPipeline(
	filename string,
	fileSize, chunkCount int) <- chan int {

	chunkSize := fileSize / chunkCount

	sortResult := [] <- chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if(err != nil) {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReadSource(bufio.NewReader(file), chunkSize)

		sortResult = append(sortResult, pipeline.InMemSort(source))
	}

	return pipeline.MergeN(sortResult...)
}

func writeToFile(p <- chan int, filename string) {
	file, err := os.Create(filename)
	if(err != nil) {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	pipeline.WriterSink(writer, p)
}

func printFile(filname string, num int) {
	file, err := os.Open(filname)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	p := pipeline.ReadSource(file, -1);

	count := 0;
	for v := range p {
		fmt.Println(v)
		count ++
		if count > num {
			break
		}
	}
}