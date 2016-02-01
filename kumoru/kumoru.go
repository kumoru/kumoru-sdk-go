package kumoru

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/kumoru/kumoru-sdk-go/RootCAs"
)

const (
	POST   = "POST"
	GET    = "GET"
	HEAD   = "HEAD"
	PUT    = "PUT"
	DELETE = "DELETE"
	PATCH  = "PATCH"
)

type (
	Request  *http.Request
	Response *http.Response

	KumoruClient struct {
		SliceData         []interface{}
		BasicAuth         struct{ UserName, Password string }
		BounceToRawString bool
		Client            *http.Client
		Data              map[string]interface{}
		Debug             bool
		EndPoint          *Endpoints
		Errors            []error
		FormData          url.Values
		Header            map[string]string
		Logger            *log.Logger
		Method            string
		QueryData         url.Values
		RawString         string
		Sign              bool
		TargetType        string
		Tokens            *Ktokens
		Transport         *http.Transport
		Url               string
	}
)

// Used to create a new KumoruClient Object.
func New() *KumoruClient {

	config := os.Getenv("KUMORU_CONFIG")

	if config == "" {
		usrHome := os.Getenv("HOME")

		config = usrHome + "/.kumoru/config"

	}

	e := LoadEndpoints(config, "endpoints")

	t, err := LoadTokens(config, "tokens")
	if err != nil {
		log.Warning("No tokens found...")
	}

	return &KumoruClient{
		BounceToRawString: false,
		Client:            &http.Client{},
		Data:              make(map[string]interface{}),
		Debug:             false,
		SliceData:         []interface{}{},
		EndPoint:          &e,
		Errors:            nil,
		FormData:          url.Values{},
		Header:            make(map[string]string),
		Logger:            log.New(),
		QueryData:         url.Values{},
		RawString:         "",
		Sign:              false,
		TargetType:        "form",
		Tokens:            &t,
		Transport:         &http.Transport{},
		Url:               "",
	}

}

func (k *KumoruClient) SignRequest(enable bool) {
	k.Sign = enable
}

func (k *KumoruClient) SetDebug(enable bool) {
	k.Debug = enable
}

func (k *KumoruClient) SetLogger(logger *log.Logger) {
	k.Logger = logger
}

// Clear KumoruClient data for a new request
func (k *KumoruClient) ClearKumoruClient() {
	k.BounceToRawString = false
	k.Data = make(map[string]interface{})
	k.Errors = nil
	k.FormData = url.Values{}
	k.Header = make(map[string]string)
	k.Method = ""
	k.QueryData = url.Values{}
	k.RawString = ""
	k.Sign = false
	k.TargetType = "form"
	k.Url = ""
	k.SliceData = []interface{}{}
}

func (k *KumoruClient) Get(targetUrl string) {
	k.ClearKumoruClient()
	k.Method = GET
	k.Url = targetUrl
	k.Errors = nil
}

func (k *KumoruClient) Patch(targetUrl string) {
	k.ClearKumoruClient()
	k.Method = PATCH
	k.Url = targetUrl
	k.Errors = nil
}

func (k *KumoruClient) Put(targetUrl string) {
	k.ClearKumoruClient()
	k.Method = PUT
	k.Url = targetUrl
	k.Errors = nil
}

func (k *KumoruClient) Delete(targetUrl string) {
	k.ClearKumoruClient()
	k.Method = DELETE
	k.Url = targetUrl
	k.Errors = nil
}

func (k *KumoruClient) Post(targetUrl string) {
	k.ClearKumoruClient()
	k.Method = POST
	k.Url = targetUrl
	k.Errors = nil
}

func (k *KumoruClient) Head(targetUrl string) {
	k.ClearKumoruClient()
	k.Method = POST
	k.Url = targetUrl
	k.Errors = nil
}

// SetHeader headers
// kumoru.New().
// POST("/application/B8658129-701E-432C-BD80-5D0F464EC932").
// SetHeader("Accept", "application/x-www-form-urlencoded")
func (k *KumoruClient) SetHeader(param string, value string) {
	k.Header[param] = value
}

func (k *KumoruClient) Param(key string, value string) {
	k.QueryData.Add(key, value)

}

// Probably can merge in the initialize section
func (k *KumoruClient) SetBasicAuth(username string, password string) {
	k.BasicAuth = struct{ UserName, Password string }{username, password}

}

// Query fucntion forms a query-string in the url of GET method or body of POST method.
// Usage Example:
//
// kumoru.New().
// Get("/applications/").
// Query(`{ query: 'myapp' }`).
// Query(`{ limit: '5' }`).
// End()
//
// kumoru.New().
// Get("/applications/").
// Query("query=myapp&limit=5").
// End()
//
// kumoru.New().
// Get("/applications/").
// Query("query=myapp&limit=5").
// Query(`{ sort: 'asc' }`).
// End()

func (k *KumoruClient) Query(content interface{}) {
	switch v := reflect.ValueOf(content); v.Kind() {

	case reflect.String:
		k.queryString(v.String())
	case reflect.Struct:
		k.queryStruct(v.Interface())
	default:
	}

}

func (k *KumoruClient) queryString(content string) {
	var val map[string]string
	if err := json.Unmarshal([]byte(content), &val); err == nil {
		for key, v := range val {
			k.QueryData.Add(key, v)
		}
	} else {
		if queryVal, err := url.ParseQuery(content); err == nil {
			for key, _ := range queryVal {
				k.QueryData.Add(key, queryVal.Get(key))
			}
		} else {
			k.Errors = append(k.Errors, err)
		}
	}

}

func (k *KumoruClient) queryStruct(content interface{}) {
	if marchalContent, err := json.Marshal(content); err != nil {
		k.Errors = append(k.Errors, err)
	} else {
		var val map[string]interface{}
		if err := json.Unmarshal(marchalContent, &val); err != nil {
			k.Errors = append(k.Errors, err)
		} else {
			for key, v := range val {
				key = strings.ToLower(key)
				k.QueryData.Add(key, v.(string))
			}
		}
	}

}

func (k *KumoruClient) TLSClientConfig(config *tls.Config) {
	k.Transport.TLSClientConfig = config

}

// End() or EndBytes() must be called to execute the call otherwise it won't do a thing.
func (k *KumoruClient) End(callback ...func(response Response, body string, errs []error)) (Response, string, []error) {
	var bytesCallback []func(response Response, body []byte, errs []error)
	if len(callback) > 0 {
		bytesCallback = []func(response Response, body []byte, errs []error){
			func(response Response, body []byte, errs []error) {
				callback[0](response, string(body), errs)
			},
		}
	}
	resp, body, errs := k.EndBytes(bytesCallback...)
	bodyString := string(body)
	return resp, bodyString, errs
}

// EndBytes should be used when you want the body as bytes.
func (k *KumoruClient) EndBytes(callback ...func(response Response, body []byte, errs []error)) (Response, []byte, []error) {
	// check whether there is an error. if yes, return all errors
	if len(k.Errors) != 0 {
		return nil, nil, k.Errors
	}

	req, err := k.NewRequest()

	if err != nil {
		k.Errors = append(k.Errors, err)
		return nil, nil, k.Errors
	}

	for key, v := range k.Header {
		req.Header.Set(key, v)
	}

	// Add all querystring from Query func
	q := req.URL.Query()
	for key, v := range k.QueryData {
		for _, vv := range v {
			q.Add(key, vv)
		}
	}
	req.URL.RawQuery = q.Encode()

	// Sign Request
	if !k.Sign {
		if k.BasicAuth != struct{ UserName, Password string }{} {
			req.SetBasicAuth(k.BasicAuth.UserName, k.BasicAuth.Password)
		}
	} else {
		date := time.Now().UTC()
		compliantDate := fmt.Sprintf(date.Format(time.RFC822Z))

		signingString := k.Method + "\n"

		d, err := ioutil.ReadAll(strings.NewReader(k.RawString))

		if err != nil {
			log.Fatal(err)
		}
		if len(d) != 0 {
			md5Sum := md5.Sum(d)
			req.Header.Set("Content-MD5", fmt.Sprintf("%x", string(md5Sum[:16])))

			signingString += fmt.Sprintf("content-md5:%x", string(md5Sum[:16])) + "\n"
			signingString += fmt.Sprintf("content-type:%v", req.Header.Get("Content-Type")+"\n")
		}

		u, _ := url.Parse(k.Url)
		signingString += "x-kumoru-date:" + compliantDate + "\n" + u.Path
		req.Header.Set("X-Kumoru-Date", compliantDate)

		h := hmac.New(sha256.New, []byte(k.Tokens.Private))
		h.Write([]byte(signingString))
		digest := fmt.Sprintf("%x", h.Sum(nil))

		req.Header.Set("Authorization", base64.StdEncoding.EncodeToString([]byte(k.Tokens.Public+":"+digest)))
	}

	// Set Transport
	certPool := *x509.NewCertPool()
	certPool.AppendCertsFromPEM(RootCAs.AlphaSSLCA)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: &certPool},
	}

	k.Client.Transport = tr

	// Log details of this request
	if k.Debug {
		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			k.Logger.Println("Error:", err)
		} else {
			k.Logger.Printf("HTTP Request: %s", string(dump))
		}
	}

	// Send request
	resp, err := k.Client.Do(req)
	if err != nil {
		k.Errors = append(k.Errors, err)
		return nil, nil, k.Errors
	}
	defer resp.Body.Close()

	// Log details of this response
	if k.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if nil != err {
			k.Logger.Println("Error:", err)
		} else {
			k.Logger.Printf("HTTP Response: %s", string(dump))
		}
	}

	body, _ := ioutil.ReadAll(resp.Body)
	// Reset resp.Body so it can be use again
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	// deep copy response to give it to both return and callback func
	respCallback := *resp
	if len(callback) != 0 {
		callback[0](&respCallback, body, k.Errors)
	}
	return resp, body, nil
}

func changeMapToURLValues(data map[string]interface{}) url.Values {
	var newUrlValues = url.Values{}
	for k, v := range data {
		switch val := v.(type) {
		case string:
			newUrlValues.Add(k, val)
		case []string:
			for _, element := range val {
				newUrlValues.Add(k, element)
			}
		// if a number, change to string
		// json.Number used to protect against a wrong (for GoRequest) default conversion
		// which always converts number to float64.
		// This type is caused by using Decoder.UseNumber()
		case json.Number:
			newUrlValues.Add(k, string(val))
		}
	}
	return newUrlValues
}
