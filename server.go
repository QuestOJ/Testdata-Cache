package main

import (
	"net/http"

	"github.com/QuestOJ/testdata-cache/testdata"
	"github.com/QuestOJ/testdata-cache/verify"
)

func startServer() {
	log(2, "Listening on 0.0.0.0:"+port)

	http.HandleFunc("/", sayHello)
	http.HandleFunc("/download", download)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)

	if err != nil {
		log(1, err.Error())
	}
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	log(3, "Recevie request from "+request.RemoteAddr+" path "+request.URL.Path)
	writer.Write([]byte(VERSION))
}

func download(writer http.ResponseWriter, request *http.Request) {
	log(3, "Recevie request from "+request.RemoteAddr+" path "+request.URL.Path)

	judgerName := request.PostFormValue("judger_name")

	if judgerName == "" {
		writer.WriteHeader(403)
		log(3, "403 Authentication name required")
		return
	}

	password := request.PostFormValue("password")

	if password == "" {
		writer.WriteHeader(403)
		log(3, "403 Authentication password required")
		return
	}

	res, err := verify.Verify(judgerName, password, config)

	if err != nil {
		writer.WriteHeader(403)
		log(3, "403 Authentication failed")
		log(2, err.Error())
		return
	}

	if res == false {
		writer.WriteHeader(403)
		log(3, "403 Authentication failed")
		return
	}

	id := request.PostFormValue("id")

	if id == "" {
		writer.WriteHeader(400)
		log(3, "400 Empty Testdata ID")
		return
	}

	filetype := request.PostFormValue("filetype")

	if filetype == "" {
		writer.WriteHeader(400)
		log(3, "400 Empty Testdata Filetype")
		return
	}

	err = testdata.Get(id, filetype, dataDir, config.OSS, writer)
	if err != nil {
		log(2, err.Error())
	}
}
