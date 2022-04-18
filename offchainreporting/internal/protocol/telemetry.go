package protocol

import "github.com/ferru97/libocr/offchainreporting/types"

type TelemetrySender interface {
	RoundStarted(
		configDigest types.ConfigDigest,
		epoch uint32,
		round uint8,
		leader types.OracleID,
	)
}
