package main

import (
	"net"
	"net/url"
	"os"

	"github.com/shynome/aliyun-monitor-report/api"
	_ "github.com/shynome/aliyun-monitor-report/api/init"
)

func main() {
	e := api.Server
	endpoint := os.Getenv("LISTEN_ENDPOINT")
	if endpoint == "" {
		endpoint = "http://:3000"
	}
	uri, err := url.Parse(endpoint)
	if err != nil {
		e.Logger.Fatal(err)
		return
	}
	if uri.Scheme == "http" {
		e.Logger.Fatal(e.Start(uri.Host))
		return
	}
	if uri.Scheme == "unix" {
		os.Remove(uri.Path)
		l, err := net.Listen("unix", uri.Path)
		if err != nil {
			e.Logger.Fatal(err)
			return
		}
		e.Listener = l
		e.Logger.Fatal(e.Start(""))
		return
	}
	e.Logger.Fatal("SERVER_ENDPOINT must be http://:3000 or unix:///tmp/foo.sock")
}
