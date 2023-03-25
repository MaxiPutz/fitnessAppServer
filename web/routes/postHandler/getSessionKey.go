package posthandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strava/key"
	"strava/structs"
)

func GetSessionKey(w http.ResponseWriter, r *http.Request) structs.Response {
	var data struct {
		Code         string `json:"code"`
		ClientID     string `json:"clientId"`
		ClientSecret string `json:"clientSecret"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return structs.Response{}
	}

	fmt.Printf("Code: %s, Client ID: %s, Client Secret: %s\n", data.Code, data.ClientID, data.ClientSecret)

	form := url.Values{}
	form.Add("client_id", data.ClientID)
	form.Add("client_secret", data.ClientSecret)
	form.Add("code", data.Code)
	form.Add("grant_type", "authorization_code")
	form.Add("Content-Type", "application/x-www-form-urlencoded")

	stravaUrl := "https://www.strava.com/oauth/token"
	req, _ := http.NewRequest("POST", stravaUrl, bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)

	}

	defer res.Body.Close()

	response := structs.Response{}
	json.NewDecoder(res.Body).Decode(&response)

	fmt.Println(response.AccessToken)
	fmt.Println(response.RefreshToken)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonRes, _ := json.Marshal(response)

	prettyJson := bytes.Buffer{}

	json.Indent(&prettyJson, jsonRes, "", "\t")
	// os.WriteFile("./key/responseKey.json", prettyJson.Bytes(), 0222)

	token := key.ResponseKey().AccessToken
	fmt.Printf("token: %v\n", token)
	w.Write(jsonRes)
	return response
}
