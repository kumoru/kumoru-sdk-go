package resources

import (
	"fmt"
	"net/http"

	"github.com/kumoru/kumoru-sdk-go/kumoru"
)

// Find resources that are accesible to the requester
func Find(rType, action, UUID string, wrappedRequest *http.Request) (*http.Response, string, []error) {
	params := "select_by="
	params += fmt.Sprintf("type=%s,", rType)
	params += fmt.Sprintf("action=%s", action)
	if UUID != "" {
		params += fmt.Sprintf(",uuid=%s,", UUID)
	}

	fmt.Println("Params: ", params)

	k := kumoru.New()
	k.Get(fmt.Sprintf("%s/v1/resources/", k.EndPoint.Authorization))
	k.Query(params)
	if wrappedRequest != nil {
		k.ProxyRequest(wrappedRequest)
		if wrappedRequest.Header.Get("X-Kumoru-Context") != "" {
			k.SetHeader("X-Kumoru-Context", wrappedRequest.Header.Get("X-Kumoru-Context"))
		}
	}
	k.SignRequest(true)

	return k.End()
}
