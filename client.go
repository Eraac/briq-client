package briq

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type (
	Briq struct {
		Key         string
		AccessToken string
	}

	Response struct {
		Error error
		Page  *Link
	}
)

var (
	Client = http.Client{
		Timeout: 10 * time.Second,
	}

	statusCode = map[string]int{
		http.MethodGet:    http.StatusOK,
		http.MethodPost:   http.StatusCreated,
		http.MethodDelete: http.StatusNoContent,
	}
)

func NewBriq(organizationKey, apiKey string) Briq {
	return Briq{
		Key:         organizationKey,
		AccessToken: base64.StdEncoding.EncodeToString([]byte(apiKey + ":")),
	}
}

func (b Briq) key() string {
	return url.PathEscape(b.Key)
}

func (b Briq) uri(endpoint string, params ...interface{}) string {
	return fmt.Sprintf(endpoint, params...)
}

func (b Briq) request(method, uri string, p *Pagination, body io.Reader, out interface{}) Response {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s%s", BaseURL, uri, p.query()), body)
	if err != nil {
		return Response{Error: err}
	}

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", b.AccessToken))
	req.Header.Set("Content-Type", "application/json")

	res, err := Client.Do(req)
	if err != nil {
		return Response{Error: err}
	}

	if res.StatusCode != statusCode[method] {
		return Response{Error: fmt.Errorf("server response %d", res.StatusCode)}
	}

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Response{Error: err}
	}

	return Response{Error: json.Unmarshal(bs, out), Page: linkFromResponse(res)}
}
