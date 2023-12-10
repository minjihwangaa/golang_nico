package channels

import (
	"fmt"
	"time"
)

func IsSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}

func IsSexy2(person string, c chan string) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- person + "is sexy"
}