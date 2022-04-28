package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pkg/pb"
	"sync"
)

func main() {
	TestPublishOrder()
}

func TestPublishOrder() {
	req := &pb.PublishOrderRequest{
		PassengerId:   214,
		StartTime:     1242566774356,
		Origin:        "srfsdgs",
		Destination:   "hdfdsfd",
		PassengerName: "treywfw",
	}
	data, _ := json.Marshal(req)
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				httpReq, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:10020/passenger/publish-order", bytes.NewReader(data))
				if err != nil {
					fmt.Println("NewRequest error", err)
					return
				}
				rsp, err := http.DefaultClient.Do(httpReq)
				if err != nil {
					fmt.Println("client request error", err)
				}
				body, err := ioutil.ReadAll(rsp.Body)
				if err != nil {
					fmt.Println("ReadAll error", err)
				}
				fmt.Println(string(body))
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
