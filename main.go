package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"os"
	"strconv"

	. "github.com/mickael-kerjean/webpty/common"
	"github.com/mickael-kerjean/webpty/ctrl"
)

var port int = 3456
var addr string = "localhost"

func init() {
	if pStr := os.Getenv("PORT"); pStr != "" {
		if pInt, err := strconv.Atoi(pStr); err == nil {
			port = pInt
		}
	}
	if pStr := os.Getenv("ADDR"); pStr != "" {
		addr = pStr
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ctrl.Main)
	mux.HandleFunc("/setup", ctrl.SetupTunnel)
	mux.HandleFunc("/tunnel.js", ctrl.RedirectTunnel)
	mux.HandleFunc("/healthz", ctrl.HealthCheck)
	mux.HandleFunc("/favicon.ico", ctrl.ServeFavicon)

	msg := `
    ██╗    ██╗███████╗██████╗ ██████╗ ████████╗██╗   ██╗
    ██║    ██║██╔════╝██╔══██╗██╔══██╗╚══██╔══╝╚██╗ ██╔╝
    ██║ █╗ ██║█████╗  ██████╔╝██████╔╝   ██║    ╚████╔╝
    ██║███╗██║██╔══╝  ██╔══██╗██╔═══╝    ██║     ╚██╔╝
    ╚███╔███╔╝███████╗██████╔╝██║        ██║      ██║
     ╚══╝╚══╝ ╚══════╝╚═════╝ ╚═╝        ╚═╝      ╚═╝

    Web Interface:
`
	msg += fmt.Sprintf("    - http://%s:%d\n", addr, port)
	Log.Stdout(msg + "\nLOGS:")
	Log.Info("WebPty is ready to go")

	if err := (&http.Server{
		Addr:     fmt.Sprintf("%s:%d", addr, port),
		Handler:  mux,
		ErrorLog: NewNilLogger(),
	}).ListenAndServe(); err != nil {
		Log.Error("[https]: listen_serve %v", err)
	}
}
