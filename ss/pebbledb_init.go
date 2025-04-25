package ss

import (
	"github.com/plume-protocol/plume-db/common/utils"
	"github.com/plume-protocol/plume-db/config"
	"github.com/plume-protocol/plume-db/ss/pebbledb"
	"github.com/plume-protocol/plume-db/ss/types"
)

func init() {
	initializer := func(dir string, configs config.StateStoreConfig) (types.StateStore, error) {
		dbHome := utils.GetStateStorePath(dir, configs.Backend)
		if configs.DBDirectory != "" {
			dbHome = configs.DBDirectory
		}
		return pebbledb.New(dbHome, configs)
	}
	RegisterBackend(PebbleDBBackend, initializer)
}
