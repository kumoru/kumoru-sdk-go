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

// Constant Methods
const (
	POST   = "POST"
	GET    = "GET"
	HEAD   = "HEAD"
	PUT    = "PUT"
	DELETE = "DELETE"
	PATCH  = "PATCH"
)

type (
	// Request declaration
	Request *http.Request
	// Response declaration
	Response *http.Response

	// Client declartation
	Client struct {
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
		ProxyRequestData  *http.Request
		QueryData         url.Values
		RawString         string
		Sign              bool
		SliceData         []interface{}
		TargetType        string
		Tokens            *Ktokens
		Transport         *http.Transport
		URL               string
	}
)

// New creates a Client Object.
func New() *Client {

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

	envDebug := false
	if strings.ToLower(os.Getenv("KUMORU_SDK_DEBUG")) == "true" {
		envDebug = true
	}

	logger := log.New()

	return &Client{
		BounceToRawString: false,
		Client:            &http.Client{},
		Data:              make(map[string]interface{}),
		Debug:             envDebug,
		EndPoint:          &e,
		Errors:            nil,
		FormData:          url.Values{},
		Header:            make(map[string]string),
		Logger:            logger,
		ProxyRequestData:  nil,
		QueryData:         url.Values{},
		RawString:         "",
		Sign:              false,
		SliceData:         []interface{}{},
		TargetType:        "form",
		Tokens:            &t,
		Transport:         &http.Transport{},
		URL:               "",
	}

}

// SignRequest enables kumoru's authentication
func (k *Client) SignRequest(enable bool) {
	k.Sign = enable
}

// SetDebug enables debugging
func (k *Client) SetDebug(enable bool) {
	k.Debug = enable
}

// SetLogger enable logger
func (k *Client) SetLogger(logger *log.Logger) {
	k.Logger = logger
}

// ClearClient clears data for a new request
func (k *Client) ClearClient() {
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
	k.URL = ""
	k.SliceData = []interface{}{}
}

// Get method
func (k *Client) Get(targetURL string) {
	k.ClearClient()
	k.Method = GET
	k.URL = targetURL
	k.Errors = nil
}

// Patch method
func (k *Client) Patch(targetURL string) {
	k.ClearClient()
	k.Method = PATCH
	k.URL = targetURL
	k.Errors = nil
}

// Put method
func (k *Client) Put(targetURL string) {
	k.ClearClient()
	k.Method = PUT
	k.URL = targetURL
	k.Errors = nil
}

// Delete method
func (k *Client) Delete(targetURL string) {
	k.ClearClient()
	k.Method = DELETE
	k.URL = targetURL
	k.Errors = nil
}

// Post method
func (k *Client) Post(targetURL string) {
	k.ClearClient()
	k.Method = POST
	k.URL = targetURL
	k.Errors = nil
}

// Head method
func (k *Client) Head(targetURL string) {
	k.ClearClient()
	k.Method = POST
	k.URL = targetURL
	k.Errors = nil
}

// SetHeader headers
// kumoru.New().
// POST("/application/B8658129-701E-432C-BD80-5D0F464EC932").
// SetHeader("Accept", "application/x-www-form-urlencoded")
func (k *Client) SetHeader(param string, value string) {
	k.Header[param] = value
}

// Param function adds a key value pair to the list of parameters
func (k *Client) Param(key string, value string) {
	k.QueryData.Add(key, value)

}

// SetBasicAuth user name and password
func (k *Client) SetBasicAuth(username string, password string) {
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
func (k *Client) Query(content interface{}) {
	switch v := reflect.ValueOf(content); v.Kind() {

	case reflect.String:
		k.queryString(v.String())
	case reflect.Struct:
		k.queryStruct(v.Interface())
	default:
	}

}

func (k *Client) queryString(content string) {
	var val map[string]string
	if err := json.Unmarshal([]byte(content), &val); err == nil {
		for key, v := range val {
			k.QueryData.Add(key, v)
		}
	} else {
		if queryVal, err := url.ParseQuery(content); err == nil {
			for key := range queryVal {
				k.QueryData.Add(key, queryVal.Get(key))
			}
		} else {
			k.Errors = append(k.Errors, err)
		}
	}

}

func (k *Client) queryStruct(content interface{}) {
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

// TLSClientConfig set TLS configuration
func (k *Client) TLSClientConfig(config *tls.Config) {
	k.Transport.TLSClientConfig = config

}

// ProxyRequest set ProxyRequest Headers
func (k *Client) ProxyRequest(r *http.Request) {
	k.ProxyRequestData = r
}

func genProxyRequestHeader(r *http.Request) string {
	components := r.Method + "\n"

	for _, value := range []string{"Content-MD5", "Content-Type", "Proxy-Authorization"} {
		if r.Header.Get(value) != "" {
			components += fmt.Sprintf("%s:%s\n", strings.ToLower(value), r.Header.Get(value))
		}
	}

	if r.Header.Get("X-Kumoru-Date") == "" {
		components += fmt.Sprintf("date:%s\n", r.Header.Get("Date")) + r.URL.Path
	} else {
		components += fmt.Sprintf("x-kumoru-date:%s\n", r.Header.Get("X-Kumoru-Date")) + r.URL.Path
	}

	tmpAuthHeader, err := base64.StdEncoding.DecodeString(r.Header.Get("Authorization"))
	if err != nil {
		return ""
	}
	fmt.Println("genProxyRequestHeaders.components: ", components)

	return base64.StdEncoding.EncodeToString([]byte(string(tmpAuthHeader) + ":" + base64.StdEncoding.EncodeToString([]byte(components))))
}

// End or EndBytes() must be called to execute the call otherwise it won't do a thing.
func (k *Client) End(callback ...func(response Response, body string, errs []error)) (Response, string, []error) {
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
func (k *Client) EndBytes(callback ...func(response Response, body []byte, errs []error)) (Response, []byte, []error) {
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

		d, readErr := ioutil.ReadAll(strings.NewReader(k.RawString))

		if readErr != nil {
			k.Logger.Fatal(err)
		}

		if len(d) != 0 {
			md5Sum := md5.Sum(d)
			req.Header.Set("Content-MD5", fmt.Sprintf("%x", string(md5Sum[:16])))

			signingString += fmt.Sprintf("content-md5:%v\n", req.Header.Get("Content-MD5"))
			signingString += fmt.Sprintf("content-type:%v\n", req.Header.Get("Content-Type"))
		}

		if k.ProxyRequestData != nil {
			req.Header.Set("Proxy-Authorization", genProxyRequestHeader(k.ProxyRequestData))
			signingString += fmt.Sprintf("proxy-authorization:%v\n", req.Header.Get("Proxy-Authorization"))

			if k.ProxyRequestData.Header.Get("X-Kumoru-Context") != "" {
				req.Header.Set("X-Kumoru-Context", k.ProxyRequestData.Header.Get("X-Kumoru-Context"))
				signingString += fmt.Sprintf("x-kumoru-context:%v\n", req.Header.Get("x-kumoru-context"))
			}

			k.Logger.Info("signingString", signingString)
		}

		u, _ := url.Parse(k.URL)
		k.Logger.Info("k.Url", k.URL)
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
		dump, logErr := httputil.DumpRequest(req, true)
		k.check(logErr, dump)
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
		k.check(err, dump)
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

func (k *Client) check(e error, dump []byte) {
	if e != nil {
		k.Logger.Println("Error:", e)
	} else {
		k.Logger.Printf("HTTP Request: %s", string(dump))
	}

}
