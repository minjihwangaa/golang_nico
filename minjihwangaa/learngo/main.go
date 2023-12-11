package main

import (
	"fmt"
	// "strings"
)

//1.2 length, return
// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

//1.3 repeat
// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }

//1.4 naked return (return 없이 출력) defer(return 실행될때 실행되는 함수)
// func nakedReturn(name string) (length int, uppercase string) {
// 	defer fmt.Println(("nakedReturn finished"))
// 	length = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return
// }
//1.5,1.6 for, range
// func supperAdd(numbers ...int) int {
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}

// 	return total
// }
//1.7 if with Twist (go lang 은 초기 변수를 if에 작성 가능 - variable expression), else 생략 가능
// func canIDrink(age int) bool {
// 	// koreanAge := age +2
// 	// if koreanAge < 18 {
// 	// 	return false
// 	// }
// 	// return true

// 	if koreanAge := age +2; koreanAge < 18 {
// 		return false
// 	}
// 	return true
// }

// 1.8 Switch (if elseif 난무하는 경우 사용), variable expression 사용 가능
// func canYouDrink(age int) bool {
// 	switch koreanAge := age +2; koreanAge{
// 	case 10:
// 		return false
// 	case 18:
// 		return true
// 	}
// 	return false
// }

//1.8 pointers(memorie address) - low level programming
// 값이 복사되지 않기를 바랄때, 메모리 주소가 필요할때 사용
// &(변수) : 변수의 메모리 주소, *변수 : 변수 메모리의 값(value)
// func myPointers() {
// 	a := 2
// 	b := &a // b는 a의 메모리주소
// 	*b = 20 // a의 메모리 값(*b)을 20으로 변경
// 	fmt.Println(a, b, *b)
// }

//1.9 Arrays and Slices (Slices 는 기본적으로 Array인데 length가 없음)
// a := [2]string{"nico","minji"} // Arrays
// b := []string{"nico","minji"} // Slices
// 값 추가할시 append사용 가능하나 데이터에 변형 주지 않음
// func ArraysSlices () {
// 	names := []string{"nico","minji"}
// 	names = append(names, "junho")
// 	fmt.Println(names)
// }

//1.10 Structs
// type person struct{
// 	name string
// 	age  int
// 	favFood []string
// }
// func StructFunc() {
// 	favFood := []string {"kimchi","ramen"}
// 	nico := person{name:"nico", age: 33, favFood: favFood}
// 	fmt.Println(nico)
// }
func main() {
	// totalLen, _ := lenAndUpper("minji")
	// fmt.Println(totalLen)
	// repeatMe("minji","sieun","janghhon")
	// _, uppercase := nakedReturn("sie")
	// fmt.Println(uppercase)
	// fmt.Println(supperAdd(1,2,3,4,5) )

	// fmt.Println(canIDrink(18))
	// fmt.Println(canYouDrink(18))

	// myPointers()

	// ArraysSlices()

	// StructFunc()

	// account := banking.Account{Owner: "nico", Balance: 1000}
	// account.Owner ="pepe" // public struct variable can change
	// fmt.Println(account)

	// account := account.NewAccount("nico")
	// fmt.Println(account)
	// account.Deposit(10)
	// fmt.Println(account.Balance())
// account.WithDraw(20)
    // fmt.Println(account.Balance())

    // account.WithDraw2(20) // Go는 exception이 없기 때문에 return 을 바로 출력하지 않음
    // err := account.WithDraw2(20)
    // if err != nil {
    //  // log.Fatalln(err) // log.Fatalln => Println + kill 
    //  fmt.Println(err)
    // }
    // fmt.Println(account.Balance())

    // dictionary := mydict.Dictionary{}
    // dictionary["hello"] = "hello"

    // dictionary := mydict.Dictionary{"first": "First word"}
    // definition, err := dictionary.Search("second")
    // if err != nil {
    //  fmt.Println(err)
    // }else{
    //  fmt.Println(definition)
    // }
    
    // dictionary := mydict.Dictionary{}
    // word := "hello"
    // definition := "greeting"
    // err := dictionary.Add(word, definition)
    // if err != nil {
    //  fmt.Println(err)
    // }
    // hello, _ := dictionary.Search(word)
    // fmt.Println(hello)
    // err2 := dictionary.Add(word, definition)
    // if err2 != nil {
    //  fmt.Println(err2)
    // }

    // dictionary := mydict.Dictionary{}
    // baseWord := "hello"
    // dictionary.Add(baseWord, "First")
    // err := dictionary.Update(baseWord,"Second")
    // if err != nil {
    //  fmt.Println(err)
    // }
    // word, _ := dictionary.Search(baseWord)
    // fmt.Println(word)

    // dictionary := mydict.Dictionary{}
    // baseWord := "hello"
    // definition := "greeting"
    // dictionary.Add(baseWord, definition)
    //  err := dictionary.Delete("not hello")
    // if err == nil {
    //  fmt.Println("잘 지워짐")
    // } else  {
    //  fmt.Println( err)
    // }

    // var results = map[string]string{}
    // var results = make(map[string]string) //make : maeks map for you
    // urls := []string{
    //  "https://www.google.com",
    //  "https://www.youtube.com",
    //  "https://www.facebook.com",
    //  "https://www.baidu.com",
    //  "https://www.airbnb.com",
    // }
    // for _, url := range urls {
    //  result := "ok"
    //  err := urlchecker.HitURL(url)
    //  if err != nil {
    //      result = "failed"
    //  }
    //  results[url] = result
    // }
    // for url, result := range results {
    //  fmt.Println(url, result)
    // }

    // 3.2 go루틴이 작동하는 이유 - 하단에 go 루틴이 아닌 함수가 작동하기 때문
    // Main함수는 go루틴을 기다려 주지 않음
    // go urlchecker.SexyCount("minji")
    // urlchecker.SexyCount("nico")
    // time.Sleep(time.Second * 5)


	// 3-3. Channels
	// c := make(chan bool)
	// people :=  [2]string{"nico","flynn"}
	// for _, person := range people {
	// 	go channels.IsSexy(person, c)
	// }
	// result := <-c
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c) // println 3회 이상 하여 deadlock

	// 3-4. Channels recap
	// c <- : 채널로 부터 메시지를 얻는다
	// blocking operation: 프로그램, 메인함수가 뭔가 받기 전에 멈춘다
	// c := make(chan string)
	// people :=  [2]string{"nico","flynn"}
	// for _, person := range people {
	// 	go channels.IsSexy2(person, c)
	// }
	// resultOne := <-c
	// resultTwo := <-c
	// fmt.Println(resultOne) // blocking operation
	// fmt.Println(resultTwo)
		fmt.Println()

	// 3-5
	// for i:=0; i< len(people); i++ {
	// 	fmt.Println(<-c)
	// }
}