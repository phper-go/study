package main

import (
	"encoding/json"
	"strings"

	// "database/sql"
	// "sync"

	// "sync/atomic"

	_ "github.com/go-sql-driver/mysql"

	// "bytes"
	"fmt"

	// "strings"

	// "net/http"
	// _ "net/http/pprof"
	"os"
	"path/filepath"
	"syscall"

	// "reflect"
	"sort"
	"time"
	// "github.com/phper-go/frame/func/object"
)

type User struct {
	ID   int
	List []string
	Body map[string]interface{}
}
type T fmt.Stringer

func main() {

	user := &User{
		ID:   333,
		Body: map[string]interface{}{},
	}
	user.Body["a"] = "aa"
	user.Body["bb"] = "bb"
	user.Body["ff"] = "ff"
	fmt.Println(user)
	delete(user.Body, "a")
	fmt.Println(user)

	return
	msg, _ := json.Marshal(logdata())
	fmt.Println(string(msg))
	return
	var ch = make(chan int, 1)

	go func() {
		var k int
		time.Sleep(2 * (time.Second))

		for {
			select {
			case k = <-ch:

				fmt.Println("chan:", k)
				time.Sleep(2 * (time.Second))
			}
		}
	}()

	for i := 1; i < 6; i++ {
		fmt.Println(i)
		ch <- i

		fmt.Println(i)
	}

	return
	var buf = []byte("bvtewf")
	tt(&buf)

	// bf := bytes.NewBuffer(buf)
	// bf.WriteString("44444")
	// fmt.Println(bf.String())
	fmt.Println(string(buf))

	return
	var action = "indexAction"

	fmt.Println(action[len(action)-6:])

	return
	fmt.Println(os.Getenv(("goapp_conf_example")))

	return

	if ex, err := os.Executable(); err != nil {
		fmt.Println(err)
	} else {
		path := filepath.Dir(ex) + "/t"
		fmt.Println(os.MkdirAll(path, os.ModePerm))
		fmt.Println(syscall.Access(path, syscall.O_WRONLY))

	}

	return
	// var values = map[string]interface{}{}

	// var user = &User{}

	// fmt.Println(object.Values(values))

	// fmt.Println(object.Values(user))

	// http.ListenAndServe(":8090", nil)

	return
	fmt.Println(time.Now())

	// var nums = []int{0, -100, 100}
	var nums = []int{1, 2147483647}

	fmt.Println(moveArray(nums))

	fmt.Println(time.Now())

	// for i := 0; i < 5; i++ {
	// 	defer fmt.Println(i)
	// }
	// fmt.Println(f1())

	// for i := 0; i < 5; i++ {
	// 	defer func() {
	// 		// fmt.Println(i)
	// 	}()
	// }

	// arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// arr1 := arr[:0]

	// fmt.Println(arr1)

	// return

	// fd, _ := os.OpenFile("t.log", os.O_RDWR, 0777)
	// info, _ := fd.Stat()

	// mmap, err := syscall.Mmap(int(fd.Fd()), 0, int(info.Size()), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)

	// mmap[1] = 'k'

	// fmt.Println(string(mmap))

	// syscall.Munmap(mmap)

	// os.Remove("my.db")

	// db, err := bolt.Open("my.db", 0600, nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer db.Close()

	// err = db.Update(func(tx *bolt.Tx) error {

	// 	b, err := tx.CreateBucketIfNotExists([]byte("1111"))
	// 	if err != nil {
	// 		return fmt.Errorf("create bucket: %v", err)
	// 	}

	// 	for i := 0; i < 500; i++ {
	// 		b.Put([]byte(fmt.Sprintf("%d", i)), []byte("test"))
	// 	}

	// 	// fmt.Println(string(b.Get([]byte("100"))))

	// 	return nil
	// })
	// fmt.Println("--------------------------")

	// log.Fatal(err)
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"study/lib"
// 	"time"
// 	// "strings"
// 	// "sync/atomic"
// 	// "time"
// 	// "github.com/webapp-go/frame/util/object"
// )

// // "strings"
// // "sync/atomic"
// // "github.com/webapp-go/frame/util/format"
// // "reflect"
// // "strings"
// // "github.com/webapp-go/frame/plugin/cache"
// // "github.com/webapp-go/frame/util/format"
// // "sync"
// // "time"

// type TIntf interface {
// 	Output()
// }

// func main() {

// 	// str := "{\"id\":3,\"oper\":{\"name\":\"hong\"}}"

// 	// begin := time.Now()

// 	// time.Sleep(2 * time.Second)

// 	// fmt.Println(time.Since(begin).String())
// 	// fmt.Println(time.Since(begin).Nanoseconds())
// 	// fmt.Printf("%0.4f", time.Since(begin).Seconds())

// 	// return

// 	var tv TIntf
// 	tv = &lib.T{}

// 	fmt.Println(tv)
// 	tv.Output()

// 	return
// 	str := "{\"id\":3,\"oper\":\"4\"}"
// 	var admin interface{}
// 	admin = &Admin{}

// 	json.Unmarshal([]byte(str), admin)

// 	out := admin.(*Admin)
// 	out.Name = "hong"
// 	out.Output = out.output
// 	out.Output()

// 	return
// 	var num Number

// 	defer num.Print()
// 	defer num.PPrint()
// 	// defer func() {
// 	// 	num.Print()
// 	// }()

// 	// defer func() {
// 	// 	num.PPrint()
// 	// }()
// 	num = 3
// 	return
// 	for i := 0; i < 100; i++ {
// 		go func() {
// 			loop(i, "loop:")
// 		}()
// 	}

// 	time.Sleep(1 * time.Second)
// 	return
// 	s := make([]int, 5, 10)
// 	fmt.Println(s)
// 	s1 := s[3:]
// 	fmt.Println(s1)
// 	fmt.Println(len(s1), cap(s1))

// 	return
// 	// admin := map[string]Admin{"my": Admin{ID: 123}}
// 	// a := admin["my"]
// 	// a.ID = 1000
// 	// fmt.Println(a)

// 	return
// 	// admin := &Admin{ID: 1000}
// 	// admin.Name = "hong"
// 	// admin.Pass = nil
// 	// fmt.Println(object.Values(admin))
// 	// return

// 	// var req uint32 = 10
// 	// admin.Name = "hong"
// 	// reqNo := strings.Replace(time.Now().Format("0405.00"), ".", "", 1)
// 	// reqNo += fmt.Sprintf("%010d", atomic.AddUint32(&req, 1))
// 	// fmt.Println(reqNo)

// 	// content := "0123456789abcdefghijklmn"
// 	// fmt.Println(content[:16], content[16:])

// 	// ref := reflect.ValueOf(admin)
// 	// fo := reflect.Indirect(ref).FieldByName("Info")
// 	// u, ok := fo.Interface().(*User)

// 	// var maps = &sync.Map{}

// 	// maps.Store("a", 1)
// 	// maps.Store("b", 2)
// 	// maps.Store("c", 3)

// 	// maps.Range(func(key, value interface{}) bool {
// 	// 	fmt.Println(key, value)
// 	// 	return true
// 	// })
// 	// fmt.Println(strings.Split("123123,2323", ","))

// 	// Cache := &cache.Momery{}
// 	// Cache.Construct()
// 	// Cache.Set("name", "hong")
// 	// fmt.Println(Cache.Get("name"))

// 	// sc := make(chan string)

// 	// maps.Store("11000", sc)

// 	// go func() {

// 	// 	ctmp, ok := maps.Load("11000")
// 	// 	fmt.Println(ctmp, ok)
// 	// 	c, ok := ctmp.(chan string)
// 	// 	fmt.Println(c, ok)
// 	// 	time.Sleep(3 * time.Second)
// 	// 	c <- "{1:2}"

// 	// }()

// 	// str := <-sc

// 	// fmt.Println(str)

// 	// time.Sleep(10 * time.Second)

// }

// type User struct {
// 	ID   int
// 	Name string
// 	Pass interface{}
// }

// func (this *User) Output() {
// 	fmt.Println("table:", this.Table())
// }

// func (this *User) Table() int {

// 	return this.ID + 2
// }

// type Admin struct {
// 	User
// 	Output func()
// }

// func (this *Admin) output() {
// 	fmt.Println(this.Name)
// }

// func (this *Admin) Table() int {
// 	return this.ID - 2
// }

// type Number int

// func (this Number) Print() {
// 	fmt.Println("num", this)
// }

// func (this *Number) PPrint() {
// 	fmt.Println("*num", *this)
// }

// func loop(i int, b string) {
// 	fmt.Println(b, i)
// }

func f1() (r int) {

	// r = 3
	// defer func() {
	// 	r++
	// }()
	// return r

	// r = 1
	// defer func() {
	// 	r++
	// 	fmt.Println(r)
	// }()
	// // r = 2
	// fmt.Println(r)
	return r
}

func moveArray(nums []int) int {
	sort.Ints(nums)
	len := len(nums)
	if len < 2 {
		return 0
	}

	var index = 0
	for {
		if nums[0] == nums[len-1] {
			return index
		}
		for i := 0; i < len-2; i++ {
			nums[i]++
		}
		if nums[len-1] < nums[len-2] {
			t := nums[len-2]
			nums[len-2] = nums[len-1]
			nums[len-1] = t
		}

		index++
	}
	return index
}

func join(tokens []T, delim string) string {
	res := ""
	for _, token := range tokens {
		if res != "" {
			res += delim
		}

		res += token.String()
	}

	return res
}

func tt(buf *[]byte) {

	*buf = append(*buf, 56)

}

var name = "hong"

func logdata() map[string]interface{} {

	var body interface{}
	var user = &User{
		ID:   1234,
		List: []string{},
	}
	body = user
	var resp interface{}
	var data = map[string]interface{}{}
	var logdata = map[string]interface{}{}
	var delkeys = []string{"list", "friends", "following"}

	if data_byte, err := json.Marshal(body); err == nil {
		if nil == json.Unmarshal(data_byte, &data) {
			for key, val := range data {
				var isin = false
				for _, v := range delkeys {
					if v == strings.ToLower(key) {
						isin = true
						break
					}
				}
				if isin == true {
					tval, _ := json.Marshal(val)
					fmt.Println(len(tval))
					break
				}
				logdata[key] = val
			}
			resp = logdata
		} else {
			resp = body
		}
	} else {
		resp = body
	}

	return map[string]interface{}{
		"status": 200,
		"msg":    "ok",
		"data":   resp,
	}
}
