package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/zcrawl/zcrawl/types"
)

const (
	defaultAddr = "https://zcrawl.org/api"
)

var (
	errTokenRequired = errors.New("Token required error")
)

// Config contains the client configuration.
type Config struct {
	Collection string
	Addr       string
	Token      string
}

// Client implements a client.
type Client struct {
	config *Config
}

// New initializes a new client.
func New(cfg *Config) (*Client, error) {
	if cfg.Token == "" {
		return nil, errTokenRequired
	}
	if cfg.Addr == "" {
		cfg.Addr = defaultAddr
	}
	c := &Client{
		config: cfg,
	}
	return c, nil
}

func (c *Client) newRequest(method string, path string) (*http.Request, error) {
	u, err := url.Parse(c.config.Addr)
	if err != nil {
		return nil, err
	}
	u.Path = u.Path + path
	return http.NewRequest(method, u.String(), nil)
}

func (c *Client) newRequestWithData(method string, path string, data []byte) (*http.Request, error) {
	u, err := url.Parse(c.config.Addr)
	if err != nil {
		return nil, err
	}
	u.Path = u.Path + path
	r := bytes.NewReader(data)
	return http.NewRequest(method, u.String(), r)
}

// Ping pings the server.
func (c *Client) Ping() (ok bool) {
	req, err := c.newRequest("GET", "/ping")
	if err != nil {
		return ok
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ok
	}
	if res.Body == nil {
		return ok
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ok
	}
	return bytes.Equal([]byte("pong\n"), data)
}

type storeRequest struct {
	Collection string                 `json:"collection"`
	Item       map[string]interface{} `json:"item"`
}

// CollectionStore stores a record in the specified collection.
func (c *Client) CollectionStore(item map[string]interface{}, collection string) {
	storeReq := storeRequest{
		Collection: collection,
		Item:       item,
	}
	reqJSON, _ := json.Marshal(&storeReq)
	req, err := c.newRequestWithData("POST", "/store", reqJSON)
	if err != nil {
		return
	}
	if req == nil {
		return
	}
	res, err := http.DefaultClient.Do(req)
	fmt.Println(res, err)
}

// Store stores a record in the global collection.
func (c *Client) Store(item map[string]interface{}) {
	c.CollectionStore(item, c.config.Collection)
}

type jobRequest struct {
	CrawlerID string `json:"crawler_id"`
}

// CreateJob creates a job.
func (c *Client) CreateJob(crawlerID string) (jobID string, err error) {
	jobReq := jobRequest{
		CrawlerID: crawlerID,
	}
	reqJSON, err := json.Marshal(&jobReq)
	if err != nil {
		return jobID, err
	}
	req, err := c.newRequestWithData("POST", "/jobs", reqJSON)
	if err != nil {
		return jobID, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return jobID, err
	}
	var job types.Job
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&job)
	if err != nil {
		return jobID, err
	}
	jobID = job.ID.Hex()
	return jobID, err
}
