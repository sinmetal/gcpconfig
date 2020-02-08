package gcpconfig

import (
	"context"

	"cloud.google.com/go/datastore"
)

type Config struct {
	Value string
}

type gcpConfigStore struct {
	ds *datastore.Client
}

func newGCPConfigStore(ctx context.Context, ds *datastore.Client) *gcpConfigStore {
	return &gcpConfigStore{ds: ds}
}

func (s *gcpConfigStore) Key(key string) *datastore.Key {
	return datastore.NameKey("SinmetalGCPConfig", key, nil)
}

func (s *gcpConfigStore) Get(ctx context.Context, key string) (string, error) {
	var v Config
	if err := s.ds.Get(ctx, s.Key(key), &v); err != nil {
		return "", err
	}
	return v.Value, nil
}

func (s *gcpConfigStore) Set(ctx context.Context, key string, value string) error {
	var v Config
	v.Value = value

	_, err := s.ds.Put(ctx, s.Key(key), &v)
	if err != nil {
		return err
	}
	return nil
}
