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