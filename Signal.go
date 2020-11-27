package godecoder

import (
	"encoding/json"
)

// MessageSignal is for Signal Model
type MessageSignal struct {
	MsgId      int32   `json:"msg_id"`
	Timestamp  int64   `json:"timestamp"`
	Epoch      int32   `json:"epoch"`
	Vin        string  `json:"vin"`
	Vlan       string  `json:"vlan"`
	MsgName    string  `json:"msg_name"`
	SignalName string  `json:"signal_name"`
	Value      float64 `json:"value"`
}

// NewMessagSignal return new MessageSignal
func NewMessageSignal(msgId int32, timestamp int64, epoch int32, vin string, vlan string, msgName string, signalName string, value float64) *MessageSignal {
	return &MessageSignal{MsgId: msgId, Timestamp: timestamp, Epoch: epoch, Vin: vin, Vlan: vlan, MsgName: msgName, SignalName: signalName, Value: value}
}

// String return String Representation
func (s MessageSignal) String() string {
	str, _ := json.Marshal(s)
	return string(str)
}
