package web

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"strava/key"
	pagehandler "strava/web/routes/pageHandler"
	posthandler "strava/web/routes/postHandler"
	serversideeventhandler "strava/web/routes/serverSideEventHandler"
)

//go:embed public/**
var staticFiles embed.FS

func RunWeb() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/init", pagehandler.ImportPage)
	http.HandleFunc("/exchange_token/", pagehandler.ExchangePage)
	http.HandleFunc("/syncPage", pagehandler.SyncPage)

	// dist, _ := fs.Sub(staticFiles, "public/runningApp")
	// fs := http.FileServer(http.FS(dist))
	fs := http.FileServer(http.Dir("./web/public/runningApp"))

	http.Handle("/", fs)

	http.HandleFunc("/activitiesStream", serversideeventhandler.StreamAvtivities)

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {

		response := posthandler.GetSessionKey(w, r)

		key.SetResponse((response))
		fmt.Printf("response: %v\n", response)
	})

	fmt.Println(">>>>>>> OClient started at:", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
