package systeminfo

import (
	"log"
	"net/http"
)

func getCpuInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	info, err := HandleCpuInfo()
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write([]byte(info))
}

func getDiskInfo(w http.ResponseWriter, r *http.Request) {
	info, err := HandleDiskInfo()
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write([]byte(info))
}

func getHostInfo(w http.ResponseWriter, r *http.Request) {
	info, err := HandleHostInfo()
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write([]byte(info))
}

func getLoadInfo(w http.ResponseWriter, r *http.Request) {
	info, err := HandleLoadInfo()
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write([]byte(info))
}

func getMemInfo(w http.ResponseWriter, r *http.Request) {
	info, err := HandleMemInfo()
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write([]byte(info))
}

func getNetInfo(w http.ResponseWriter, r *http.Request) {
	info, err := HandleNetInfo()
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write([]byte(info))
}
