package types

import (
	"github.com/osmosis-labs/osmosis/osmomath"
)

// ConsensusMinFee is a governance set parameter created by prop 354 (https://www.mintscan.io/osmosis/proposals/354)
// and adjusted by prop 676 (https://www.mintscan.io/osmosis/proposals/676)
// It is intended to be .025 uosmo / gas
var ConsensusMinFee osmomath.Dec = osmomath.NewDecWithPrec(25, 3)
