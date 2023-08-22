package practise

import (
	"log"
	"net/http"
)

func SetRoute() {
	http.HandleFunc("/", webPage)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 練習頁面
func webPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`my first website`))
}
