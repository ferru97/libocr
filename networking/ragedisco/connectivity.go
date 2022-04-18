package ragedisco

import ragetypes "github.com/ferru97/libocr/ragep2p/types"

type connectivityMsgType int

const (
	_ connectivityMsgType = iota
	connectivityAdd
	connectivityRemove
)

type connectivityMsg struct {
	msgType connectivityMsgType
	peerID  ragetypes.PeerID
}
