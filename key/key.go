package key

import (
	"fitnessApp/structs"
	"sync"
)

var (
	responseMutex sync.Mutex
	response      structs.Response
)

func ResponseKey() structs.Response {

	return getResponse()
}

func SetResponse(r structs.Response) {
	responseMutex.Lock()
	defer responseMutex.Unlock()
	response = r
}

func getResponse() structs.Response {
	responseMutex.Lock()
	defer responseMutex.Unlock()
	return response
}
