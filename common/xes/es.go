package xes

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/jinzhu/copier"
)

type Config struct {
	Addresses  []string
	Username   string
	Password   string
	MaxRetries int
}

func New(config Config) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{}
	if err := copier.Copy(&cfg, &config); err != nil {
		return nil, err
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func MustNew(config Config) *elasticsearch.Client {
	es, err := New(config)
	if err != nil {
		panic(err)
	}

	return es
}

func NewType(config Config) (*elasticsearch.TypedClient, error) {
	cfg := elasticsearch.Config{}
	if err := copier.Copy(&cfg, &config); err != nil {
		return nil, err
	}

	client, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func MustNewType(config Config) *elasticsearch.TypedClient {
	es, err := NewType(config)
	if err != nil {
		panic(err)
	}

	return es
}
