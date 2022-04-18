package protocol

import "github.com/ferru97/libocr/offchainreporting/types"

// Used only for testing
type XXXUnknownMessageType struct{}

// Conform to protocol.Message interface
func (XXXUnknownMessageType) process(*oracleState, types.OracleID) {}
