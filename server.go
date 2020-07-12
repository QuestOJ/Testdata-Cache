package main

import (
	"net/http"

	"github.com/QuestOJ/testdata-cache/testdata"
	"github.com/QuestOJ/testdata-cache/verify"
)

func startServer() {
	http.HandleFunc("/download", download)
	http.ListenAndServe("0.0.0.0:"+port, nil)
}

func download(writer http.ResponseWriter, request *http.Request) {
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

	testdata.Get(id, filetype, dataDir, config.OSS, writer)
}
