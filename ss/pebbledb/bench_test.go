package pebbledb

import (
	"testing"

	"github.com/plume-protocol/plume-db/config"
	sstest "github.com/plume-protocol/plume-db/ss/test"
	"github.com/plume-protocol/plume-db/ss/types"
)

func BenchmarkDBBackend(b *testing.B) {
	s := &sstest.StorageBenchSuite{
		NewDB: func(dir string) (types.StateStore, error) {
			return New(dir, config.DefaultStateStoreConfig())
		},
		BenchBackendName: "PebbleDB",
	}

	s.BenchmarkGet(b)
	s.BenchmarkApplyChangeset(b)
	s.BenchmarkIterate(b)
}
