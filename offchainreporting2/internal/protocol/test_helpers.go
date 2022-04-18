package protocol

import (
	"github.com/ferru97/libocr/commontypes"
	"github.com/ferru97/libocr/offchainreporting2/types"
)

// Used only for testing
type XXXUnknownMessageType struct{}

// Conform to protocol.Message interface
func (XXXUnknownMessageType) CheckSize(types.ReportingPluginLimits) bool { return true }
func (XXXUnknownMessageType) process(*oracleState, commontypes.OracleID) {}
