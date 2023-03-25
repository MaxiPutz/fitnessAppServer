package pagehandler

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func RunningApp(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("/home/max/workspace/fitness/fitnessAppFrontend/my-fitness-app2/build", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ImportPage(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("web", "public", "importPage.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SyncPage(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("web", "public", "sync.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ExchangePage(w http.ResponseWriter, r *http.Request) {
	myCode := r.URL.Query().Get("code")

	fmt.Println(myCode)
	fp := path.Join("web", "public", "exchange.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
