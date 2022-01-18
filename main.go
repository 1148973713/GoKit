package main

import (
	. "GoKit/Services"
	"GoKit/util"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	routeMux "github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	user := UserService{}
	endp := GenUserEndpoint(user)

	serverHanlder := httptransport.NewServer(endp, DecodeUserRequest, EncodeUserResponse)
	router := routeMux.NewRouter()
	{
		//r.Handle(`/user/(uid:\d+)`,serverHanlder)
		router.Methods("GET").Path(`/user/(uid:\d+)`).Handler(serverHanlder)
		router.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-type", "application/json")
			writer.Write([]byte(`{"status":"ok"}`))
		})
	}

	errChan := make(chan error)

	//注册服务
	go (func() {
		util.RegService()
		err := http.ListenAndServe(":8090", router)
		if err != nil {
			log.Println(err)
			errChan <- err
		}
	})()

	go (func() {
		sig_c := make(chan os.Signal)
		signal.Notify(sig_c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sig_c)
	})()

	//如果errChan没有问题一直会被阻塞不会进入下线服务操作
	getErr := <-errChan
	util.UnRegService()
	log.Println(getErr)
}
