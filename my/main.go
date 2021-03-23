package main

import (
	"encoding/json"

	// "encoding/json"
	"fmt"
	"io"
	"os"
	"time"
	// "github.com/phper-go/frame/func/conv"
)

var chan_int = make(chan int, 10)

func out(stdout io.Writer) {

}

func t1() {

	defer func() {
		err := recover()
		fmt.Println("t1 err:", err)
	}()

	fmt.Println("t1")
	go t2()
}

func t2() {

	// defer func() {
	// 	err := recover()
	// 	fmt.Println("t2 err:", err)
	// }()
	fmt.Println("t2")
	// panic("t2 bong")
}

func maps() map[string]string {
	return map[string]string{"name": "home"}
}

type User struct {
	ID   int
	Name string
}

type Admin struct {
	*User
}

func (this *User) SetName(name string) {
	this.Name = name
	fmt.Println(this.Name)
}

func (this *User) Echo() {
	fmt.Println(this.Name)
}

func tt(arr *[]int) {

	*arr = append(*arr, 333)
}

func mm(arr map[string]int) {
	arr["age"] = 111
	arr["name"] = 8888
}
func main() {

	arr := map[string]int{
		"name": 3232,
	}

	mm(arr)

	fmt.Println(arr)

	return

	arr1 := []int{111, 222}

	tt(&arr1)

	fmt.Println(arr1)

	return
	admin := Admin{
		&User{
			ID: 123,
		},
	}
	admin.ID = 12345
	admin.Name = "hong"
	admin.Echo()

	admin.SetName("abc")
	admin.Echo()
	fmt.Println(admin.Name)

	return
	admin.SetName("abc")
	fmt.Println(admin.User.Name)

	err := json.Unmarshal([]byte(""), &admin)
	fmt.Println(err)
	return

	// fmt.Println(conv.String(false) == "false")
	// fmt.Println(conv.String("[123,443]"))
	// fmt.Println(conv.String(maps()))
	return
	fmt.Println(maps())

	fmt.Println(maps()["name"])

	return

	func(run func()) {
		run()
		fmt.Println(fmt.Sprintf("%v", run))
	}(t2)

	return
	x, y := 1, 2

	defer func() {
		fmt.Printf("x:%d,y:%d\n", x, y) // y 为闭包引用
	}() // 复制 x 的值

	x += 100
	y += 100
	fmt.Println(x, y)

	return
	for k := 0; k < 10; k++ {
		go func() {
			fmt.Println(k)
		}()
	}
	time.Sleep(2 * time.Second)
	return
	defer func() {
		err := recover()
		fmt.Println("main err:", err)
	}()

	go t1()

	time.Sleep(10 * time.Second)
	return
	out(os.Stderr)
	return
	go echo()

	for k := 0; k < 20; k++ {
		chan_int <- k
	}
}

func echo() {

	var num int
	for {
		num = <-chan_int
		fmt.Println(num)
		time.Sleep(1 * time.Second)
	}
}
