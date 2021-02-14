package cvr

// Official client for https://cvr.dev.
// For more information, see https://docs.cvr.dev/.

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/micvbang/go-helpy/inty"
)

const (
	APIBaseAddress          = "https://api.cvr.dev"
	HTTPAuthorizationHeader = "Authorization"
)

// Client provides functionality to request data from cvr.dev.
type Client struct {
	client         *http.Client
	apiBaseAddress string
}

// NewClient returns a client sending requests to the cvr.dev servers.
func NewClient(apiKey string) *Client {
	c := &http.Client{Transport: newAuthTransport(apiKey, http.DefaultTransport)}
	return &Client{
		client:         c,
		apiBaseAddress: APIBaseAddress,
	}
}

// NewClientBaseAddress returns a Client sending requests to the given
// apiBaseAddress.
func NewClientBaseAddress(apiKey string, apiBaseAddress string) *Client {
	c := &http.Client{Transport: newAuthTransport(apiKey, http.DefaultTransport)}
	return &Client{
		client:         c,
		apiBaseAddress: apiBaseAddress,
	}
}

func (c *Client) httpGET(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func (c *Client) statusCodeToError(r *http.Response) error {
	switch {
	case r.StatusCode == http.StatusUnauthorized:
		return ErrUnauthorized

	case r.StatusCode == http.StatusNotFound:
		return ErrNotFound

	case r.StatusCode > http.StatusOK:
		bs, _ := ioutil.ReadAll(r.Body)
		return ErrServerError{
			Message: string(bs),
		}

	default:
		return nil
	}
}

func (c *Client) buildURL(endpoint string, params map[string]string) (string, error) {
	u, err := url.Parse(fmt.Sprintf("%s/api/%s", c.apiBaseAddress, endpoint))
	if err != nil {
		return "", err
	}

	urlValues := url.Values{}
	for key, value := range params {
		urlValues.Add(key, value)
	}

	return fmt.Sprintf("%s?%s", u.String(), urlValues.Encode()), nil
}

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNotFound     = errors.New("not found")
)

type ErrServerError struct {
	Message string
}

func (e ErrServerError) Error() string {
	return fmt.Sprintf("server error: %s", e.Message)
}

// TestAPIKey returns nil if Client successfully authenticated with the
// server.
func (c *Client) TestAPIKey() error {
	endpoint, err := c.buildURL("test/apikey", nil)
	if err != nil {
		return err
	}

	r, err := c.httpGET(endpoint)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = c.statusCodeToError(r)
	if err != nil {
		return err
	}

	return nil
}

// CVRVirksomhederByCVRNumre returns a list of Virksomheder with the given
// CVR numre. At most 10 virksomheder can be requested at once.
// NOTE: this data originates directly from CVR and is not validated in any way.
func (c *Client) CVRVirksomhederByCVRNumre(cvrNumre ...int) ([]Virksomhed, error) {
	cvrNumre = cvrNumre[0:inty.Min(9, len(cvrNumre))]

	cvrNumreStr := make([]string, len(cvrNumre))
	for i, cvrNummer := range cvrNumre {
		cvrNumreStr[i] = strconv.Itoa(cvrNummer)
	}

	endpoint, err := c.buildURL("cvr/virksomhed", map[string]string{
		"cvr_nummer": strings.Join(cvrNumreStr, ","),
	})
	if err != nil {
		return nil, nil
	}

	r, err := c.httpGET(endpoint)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	err = c.statusCodeToError(r)
	if err != nil {
		return nil, err
	}

	v := []Virksomhed{}
	err = json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// CVRVirksomhederByNavn returns a list of Virksomheder with names similar to
// the given navn. At most 25 virksomheder are returned.
// NOTE: this data originates directly from CVR and is not validated in any way.
func (c *Client) CVRVirksomhederByNavn(navn string) ([]Virksomhed, error) {
	endpoint, err := c.buildURL("cvr/virksomhed", map[string]string{
		"navn": navn,
	})
	if err != nil {
		return nil, nil
	}

	r, err := c.httpGET(endpoint)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	err = c.statusCodeToError(r)
	if err != nil {
		return nil, err
	}

	vs := []Virksomhed{}
	err = json.NewDecoder(r.Body).Decode(&vs)
	if err != nil {
		return nil, err
	}

	return vs, nil
}

// CVRProduktionsenhederByPNumre returns a list of Produktionsenheder with the
// given CVR numre. At most 10 produktionsenheder can be requested at once.
// NOTE: this data originates directly from CVR and is not validated in any way.
func (c *Client) CVRProduktionsenhederByPNumre(pNumre ...int) ([]Produktionsenhed, error) {
	pNumre = pNumre[0:inty.Min(9, len(pNumre))]

	pNumreStr := make([]string, len(pNumre))
	for i, cvrNummer := range pNumre {
		pNumreStr[i] = strconv.Itoa(cvrNummer)
	}

	endpoint, err := c.buildURL("cvr/produktionsenhed", map[string]string{
		"p_nummer": strings.Join(pNumreStr, ","),
	})
	if err != nil {
		return nil, nil
	}

	r, err := c.httpGET(endpoint)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	err = c.statusCodeToError(r)
	if err != nil {
		return nil, err
	}

	v := []Produktionsenhed{}
	err = json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (c *Client) CVRProduktionsenhederByAdresse(adresse string) ([]Produktionsenhed, error) {
	endpoint, err := c.buildURL("cvr/produktionsenhed", map[string]string{
		"adresse": adresse,
	})
	if err != nil {
		return nil, nil
	}

	r, err := c.httpGET(endpoint)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	err = c.statusCodeToError(r)
	if err != nil {
		return nil, err
	}

	v := []Produktionsenhed{}
	err = json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// authTransport is used to add Authentication headers to HTTP requests
type authTransport struct {
	apiKey string
	orig   http.RoundTripper
}

func newAuthTransport(apiKey string, orig http.RoundTripper) http.RoundTripper {
	return &authTransport{
		apiKey: apiKey,
		orig:   orig,
	}
}

func (at *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", at.apiKey)
	return at.orig.RoundTrip(req)
}
