package main

import (
	"fmt"
	"github.com/manhhnv/sk-gcs/pkg/gcs"
	"github.com/manhhnv/sk-gcs/pkg/setting"
	"github.com/manhhnv/sk-gcs/routers"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	gcs.Setup()
}

func main() {
	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", setting.AppSetting.HttpPort)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Start server error %v", err)
	}
}
//package main
//
//import (
//	"errors"
//	"sync"
//)
//
//func test(i int) (int, error) {
//	if i > 2 {
//		return 0, errors.New("test error")
//	}
//	return i + 5, nil
//}
//
//func test2(i int) (int, error) {
//	if i > 3 {
//		return 0, errors.New("test2 error")
//	}
//	return i + 7, nil
//}
//
//func main() {
//	results := make(chan int, 2)
//	errors := make(chan error, 2)
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		result, err := test(2)
//		if err != nil {
//			errors <- err
//			return
//		}
//		results <- result
//	}()
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		result, err := test2(3)
//		if err != nil {
//			errors <- err
//			return
//		}
//		results <- result
//	}()
//
//	// here we wait in other goroutine to all jobs done and close the channels
//	go func() {
//		wg.Wait()
//		close(results)
//		close(errors)
//	}()
//	for err := range errors {
//		// here error happend u could exit your caller function
//		println(err.Error())
//		//return
//
//	}
//	//fmt.Printf("results %v", results)
//	for res := range results {
//		println("--------- ", res, " ------------")
//	}
//}
