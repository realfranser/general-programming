package services

import (
	"bytes"
	"encoding/gob"
	"os"
	"time"

	"microservice_example/models"

	"github.com/bradfitz/gomemcache/memcache"
)

type Client struct {
	client *memcache.Client
}

var memcacheClient *Client

func NewMemcached() (*Client, error) {
	// XXX Assuming environment variable contains only one server
	client := memcache.New(os.Getenv("MEMCACHED"))

	if err := client.Ping(); err != nil {
		return nil, err
	}

	client.Timeout = 100 * time.Millisecond
	client.MaxIdleConns = 100

	memcacheClient = &Client{
		client: client,
	}

	return memcacheClient, nil
}

func (c *Client) CloseMemcacheClient() {
	c.client.Close()
}

func SetName(n models.Name) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(n); err != nil {
		return err
	}

	return memcacheClient.client.Set(&memcache.Item{
		Key:        n.NConst,
		Value:      b.Bytes(),
		Expiration: int32(time.Now().Add(25 * time.Second).Unix()),
	})
}

func GetName(nconst string) (models.Name, error) {
	item, err := memcacheClient.client.Get(nconst)
	if err != nil {
		return models.Name{}, err
	}

	b := bytes.NewReader(item.Value)

	var res models.Name

	if err := gob.NewDecoder(b).Decode(&res); err != nil {
		return models.Name{}, err
	}

	return res, nil
}

func Setame(n models.Name) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(n); err != nil {
		return err
	}

	return memcacheClient.client.Set(&memcache.Item{
		Key:        n.NConst,
		Value:      b.Bytes(),
		Expiration: int32(time.Now().Add(25 * time.Second).Unix()),
	})
}
