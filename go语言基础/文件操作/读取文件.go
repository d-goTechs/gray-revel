package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"io"
	"bufio"
)
func check(e error){
	if e!=nil{
		panic(e)
	}
}
func main() {
	dat, err := ioutil.ReadFile("D:/git_projects/beegoApp/src/gray-revel/resources/goFile")
	check(err)
	fmt.Println(string(dat))


	f, err := os.Open("D:/git_projects/beegoApp/src/gray-revel/resources/goFile")
	check(err)
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
	//io 包提供了一些可以帮助我们进行文件读取的函数。例如，上面的读取可以使用 ReadAtLeast 得到一个更健壮的实现。
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	//没有内置的回转支持，但是使用 Seek(0, 0) 实现。
	_, err = f.Seek(0, 0)
	check(err)
	//bufio 包实现了带缓冲的读取，这不仅对有很多小的读取操作的能提升性能，也提供了很多附加的读取函数。
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	//任务结束后要关闭这个文件（通常这个操作应该在 Open操作后立即使用 defer 来完成）。
	f.Close()
}
