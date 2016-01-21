package main

import (
	"log"
	"model"
	"net/http"
	"time"
)

// StartApi start api http listen.
func StartApi(addr string, s *model.Store) {
	go func() {
		var (
			err       error
			serverMux = http.NewServeMux()
		)
		serverMux.Handle("/version", httpVersionHandler{})
		serverMux.Handle("/get", httpGetHandler{s: s})
		serverMux.Handle("/upload", httpUploadHandler{})
		serverMux.Handle("/uploads", httpUploadsHandler{})
		serverMux.Handle("/del", httpDelHandler{})
		serverMux.Handle("/dels", httpDelsHandler{})
		if err = http.ListenAndServe(addr, serverMux); err != nil {
			log.Printf("http.ListenAndServer(\"%s\") error(%v)", addr, err)
		}
	}()
}

type httpVersionHandler struct {
}

func (h httpVersionHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var msg []byte = []byte(Version())
	rw.Write(msg)
}

type httpGetHandler struct {
	s *model.Store
}

func (h httpGetHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var (
		now = time.Now()
		ret = http.StatusOK
	)
	if r.Method != "GET" && r.Method != "HEAD" {
		http.Error(rw, "method not allowed", ret)
		return
	}
	log.Println(now)
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
