package keeper

import (
	"github.com/fivesixnine/splash/x/ibcchat/types"
)

var _ types.QueryServer = Keeper{}
