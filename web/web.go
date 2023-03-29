package web

import (
	"embed"
	"fitnessApp/key"
	pagehandler "fitnessApp/web/routes/pageHandler"
	posthandler "fitnessApp/web/routes/postHandler"
	serversideeventhandler "fitnessApp/web/routes/serverSideEventHandler"
	"fmt"
	"log"
	"net/http"
	"os"
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

	fmt.Println("Welcome to the fitnessApp \n if you downloaded your data you can use the folow link 'http://localhost:8000/' \n If you downloaded your data, you can use the follow link  'http://localhost:8080/init' \n If you have any issues, you can email me to putzmaximilian@gmail.com")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
