package server

import "fmt"

func NewBaiduServer(address, listen string, verbose bool) *Server {
	upstream := NewUpstream(
		WithAddress(fmt.Sprintf("%s:443", address)),
		WithHeaders(Headers{"X-T5-Auth": "ZjQxNDIh", "User-Agent": "okhttp/3.11.0 Dalvik/2.1.0 (Linux; U; Android 11; Redmi K30 5G Build/RKQ1.200826.002) baiduboxapp/11.0.5.12 (Baidu; P1 11)"}),
	)
	return &Server{
		Listen:        listen,
		HttpUpstream:  upstream,
		HttpsUpstream: upstream,
		Verbose:       verbose,
	}
}

func NewUcServer(listen string, verbose bool) *Server {
	upstream := NewUpstream(
		WithAddress("101.71.140.5:8128"),
		WithHeaders(Headers{"Proxy-Authorization": "Basic dWMxMC4xMDMuMjcuMTgyOjFmNDdkM2VmNTNiMDM1NDQzNDUxYzdlZTc4NzNmZjM4=="}),
	)
	return &Server{
		Listen:        listen,
		HttpUpstream:  upstream,
		HttpsUpstream: upstream,
		Verbose:       verbose,
	}
}
