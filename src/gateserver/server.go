package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"

	Handlers "gserver_gateway_demo/src/gateserver/handlers"
	Msg "gserver_gateway_demo/src/gateserver/protomsg"

	"github.com/gfandada/gserver"
)

func StartAdminHttp(webaddr string) error {
	adminServeMux := http.NewServeMux()
	adminServeMux.HandleFunc("/debug/pprof/", pprof.Index)
	adminServeMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	adminServeMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	adminServeMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	err := http.ListenAndServe(webaddr, adminServeMux)
	if err != nil {
		x := fmt.Sprintf("http.ListenAdServe(\"%s\") failed (%s)", webaddr, err.Error())
		fmt.Println(x)
		return err
	}
	return nil
}

func pproftest() {
	go func() {
		if err := StartAdminHttp("127.0.0.1:7081"); err != nil {
			os.Exit(-1)
		}
	}()
}

func main() {
	// for test
	pproftest()
	Handlers.NewHandlers()
	logger := "C:/Users/Administrator/go/src/gserver_gateway_demo/cfg/gatewaylogger.xml"
	gateway := "C:/Users/Administrator/go/src/gserver_gateway_demo/cfg/tcpgateway.json"
	dipath := "C:/Users/Administrator/go/src/gserver_gateway_demo/cfg/discovery.json"
	coder := Msg.NewMsgCoder()
	gserver.RunTCPGateway(logger, gateway, dipath, coder)
}
