package handlers

import (
	"net/http"
)

func CustomHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r,
		"http://www.baidu.com",
		http.StatusMovedPermanently,
	)
}
