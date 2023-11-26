package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Example01() {
    // create a context with timeout to call our request if takes to long
    timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond * 500)
    defer cancel()

    // create HTTP request
    req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://placehold.it/200x200", nil)
    if err != nil {
        panic(err)
    }

    // perform HTTP request
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    // get data from HTTP response
    imageData, err := io.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }
    fmt.Printf("downloaded image of size %d\n", len(imageData))
}
