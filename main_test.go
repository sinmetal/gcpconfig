package gcpconfig_test

import (
	"context"
	"os"
	"testing"

	"github.com/sinmetal/gcpconfig"
)

func TestGetFromDatastore(t *testing.T) {
	setDatastoreEmulatorHost(t)
	config := newGCPConfigService(t)

	ctx := context.Background()

	const keyName = "from-datastore"
	const value = "hogeeeee"
	if err := config.Set(ctx, keyName, value); err != nil {
		t.Fatal(err)
	}

	v := config.Get(ctx, keyName)
	if e, g := value, v; e != g {
		t.Errorf("want %v but got %v", e, g)
	}
}

func TestGetFromOSEnv(t *testing.T) {
	setDatastoreEmulatorHost(t)
	config := newGCPConfigService(t)

	ctx := context.Background()

	const keyName = "from-osenv"
	const value = "hogeeeee"
	if err := os.Setenv(keyName, value); err != nil {
		t.Fatal(err)
	}

	v := config.Get(ctx, keyName)
	if e, g := value, v; e != g {
		t.Errorf("want %v but got %v", e, g)
	}
}

func newGCPConfigService(t *testing.T) *gcpconfig.GCPConfigService {
	ctx := context.Background()
	config, err := gcpconfig.NewGCPConfigService(ctx, "local")
	if err != nil {
		t.Fatal(err)
	}
	return config
}

func setDatastoreEmulatorHost(t *testing.T) {
	if err := os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:8081"); err != nil {
		t.Fatal(err)
	}
}
