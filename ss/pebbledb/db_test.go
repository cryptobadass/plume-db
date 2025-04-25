package pebbledb

import (
	"testing"

	"github.com/plume-protocol/plume-db/config"
	sstest "github.com/plume-protocol/plume-db/ss/test"
	"github.com/plume-protocol/plume-db/ss/types"
	"github.com/stretchr/testify/suite"
)

func TestStorageTestSuite(t *testing.T) {
	s := &sstest.StorageTestSuite{
		NewDB: func(dir string, config config.StateStoreConfig) (types.StateStore, error) {
			return New(dir, config)
		},
		Config:         config.DefaultStateStoreConfig(),
		EmptyBatchSize: 12,
	}

	suite.Run(t, s)
}
