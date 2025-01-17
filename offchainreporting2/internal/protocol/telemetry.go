package protocol

import (
	"github.com/ferru97/libocr/commontypes"
	"github.com/ferru97/libocr/offchainreporting2/types"
)

type TelemetrySender interface {
	RoundStarted(
		configDigest types.ConfigDigest,
		epoch uint32,
		round uint8,
		leader commontypes.OracleID,
	)
}
