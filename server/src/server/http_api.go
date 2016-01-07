package main

import (
	"log"
	"net/http"
)

// StartApi start api http listen.
func StartApi(addr string) {
	go func() {
		var (
			err       error
			serverMux = http.NewServeMux()
		)
		serverMux.Handle("/get", httpGetHandler{})
		serverMux.Handle("/upload", httpUploadHandler{})
		serverMux.Handle("/uploads", httpUploadsHandler{})
		serverMux.Handle("/del", httpDelHandler{})
		serverMux.Handle("/dels", httpDelsHandler{})
		if err = http.ListenAndServe(addr, serverMux); err != nil {
			log.Printf("http.ListenAndServer(\"%s\") error(%v)", addr, err)
		}
	}()
}

type httpGetHandler struct {
}

func (h httpGetHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var msg []byte = []byte("get")
	rw.Write(msg)
}

type httpUploadHandler struct {
}

func (h httpUploadHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var msg []byte = []byte("upload")
	rw.Write(msg)
}

type httpUploadsHandler struct {
}

func (h httpUploadsHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var msg []byte = []byte("uploads")
	rw.Write(msg)
}

type httpDelHandler struct {
}

func (h httpDelHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var msg []byte = []byte("del")
	rw.Write(msg)
}

type httpDelsHandler struct {
}

func (h httpDelsHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var msg []byte = []byte("dels")
	rw.Write(msg)
}
