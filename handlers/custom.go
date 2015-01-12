package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func CustomHandler(w http.ResponseWriter, r *http.Request) {

    var err error
    
    log.Println("New cluster created", r.URL.RawQuery)
    
	custom, err := setupCustom(r.URL.RawQuery)
	
	if err != nil {
		log.Printf("setupToken returned: %v", err)
		http.Error(w, "Unable to generate token", 400)
		return
	}

	log.Println("New cluster created", custom)

	fmt.Fprintf(w, "http://discovery.isspaas.com/"+custom)
}
