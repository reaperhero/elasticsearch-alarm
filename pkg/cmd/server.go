package cmd

import "github.com/reaperhero/elasticsearch-alarm/handler/http"

func Run() {
	http.RunHttpserver()
}
