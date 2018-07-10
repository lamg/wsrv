package main

import (
	"flag"
	"fmt"
	h "net/http"
	"os"
	"path"
)

func main() {
	var addr, dir, wfile string
	flag.StringVar(&addr, "a", ":3000", "Address to serve")
	flag.StringVar(&dir, "d", ".", "Directory with WASM files")
	flag.StringVar(&wfile, "f", "", "WASM file")
	flag.Parse()
	h.Handle("/", h.FileServer(h.Dir(dir)))
	s := &wasmSrv{file: wfile}
	h.HandleFunc(path.Join("/", wfile), s.srvWASM)
	e := h.ListenAndServe(addr, h.DefaultServeMux)
	if e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
	}
}

type wasmSrv struct {
	file string
}

func (s *wasmSrv) srvWASM(w h.ResponseWriter, r *h.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	h.ServeFile(w, r, s.file)
}
