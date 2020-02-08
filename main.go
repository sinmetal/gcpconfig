package gcpconfig

import (
	"context"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/sinmetal/gcpmetadata"
)

type GCPConfigService struct {
	store *gcpConfigStore
}

func NewGCPConfigService(ctx context.Context, projectID string) (*GCPConfigService, error) {
	ds, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	store := newGCPConfigStore(ctx, ds)
	return &GCPConfigService{store: store}, nil
}

func (s *GCPConfigService) Get(ctx context.Context, key string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}

	v, err := gcpmetadata.GetInstanceAttribute(key)
	if err != nil {
		return ""
	}
	if v != "" {
		return v
	}

	v, err = s.store.Get(ctx, key)
	if err != nil {
		return ""
	}
	return v
}

func (s *GCPConfigService) Set(ctx context.Context, key string, value string) error {
	return s.store.Set(ctx, key, value)
}
