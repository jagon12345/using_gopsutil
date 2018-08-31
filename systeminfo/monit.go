package systeminfo

import (
	"net/http"
)

func Handle() {
	http.HandleFunc("/cpuinfo", getCpuInfo)
	http.HandleFunc("/diskinfo", getDiskInfo)
	http.HandleFunc("/netinfo", getNetInfo)
	http.HandleFunc("/meminfo", getMemInfo)
	http.HandleFunc("/hostinfo", getHostInfo)
	//http.HandleFunc("/dockerinfo", GetDockerInfo)
	http.HandleFunc("/loadinfo", getLoadInfo)
	http.ListenAndServe(":7373", nil)
}
