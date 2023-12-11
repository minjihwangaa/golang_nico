package urlchecker

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// 3.0 Hiturl
var errRequestFailed = errors.New("request failed")
func HitURL(url string) error{
    resp, err := http.Get(url)
    fmt.Println("Checking",url)
    // 400이하는 리다이렉션,400부터 에러
    if err != nil || resp.StatusCode>=400 {
        return errRequestFailed
    }
    return nil
}

// 3.2 Go routines - 병행성(Concurrency)
func SexyCount(person string){
    for i := 0; i< 10; i++ {
        fmt.Println(person,"is sexy", i)
        time.Sleep(time.Second)
    }
}


