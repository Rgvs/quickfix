//Package allocationack msg type = P.
package allocationack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a AllocationACK wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//ClientID is a non-required field for AllocationACK.
func (m Message) ClientID() (*field.ClientIDField, quickfix.MessageRejectError) {
	f := &field.ClientIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from AllocationACK.
func (m Message) GetClientID(f *field.ClientIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for AllocationACK.
func (m Message) ExecBroker() (*field.ExecBrokerField, quickfix.MessageRejectError) {
	f := &field.ExecBrokerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from AllocationACK.
func (m Message) GetExecBroker(f *field.ExecBrokerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a required field for AllocationACK.
func (m Message) AllocID() (*field.AllocIDField, quickfix.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from AllocationACK.
func (m Message) GetAllocID(f *field.AllocIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for AllocationACK.
func (m Message) TradeDate() (*field.TradeDateField, quickfix.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from AllocationACK.
func (m Message) GetTradeDate(f *field.TradeDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for AllocationACK.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from AllocationACK.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocStatus is a required field for AllocationACK.
func (m Message) AllocStatus() (*field.AllocStatusField, quickfix.MessageRejectError) {
	f := &field.AllocStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocStatus reads a AllocStatus from AllocationACK.
func (m Message) GetAllocStatus(f *field.AllocStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocRejCode is a non-required field for AllocationACK.
func (m Message) AllocRejCode() (*field.AllocRejCodeField, quickfix.MessageRejectError) {
	f := &field.AllocRejCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocRejCode reads a AllocRejCode from AllocationACK.
func (m Message) GetAllocRejCode(f *field.AllocRejCodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for AllocationACK.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from AllocationACK.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds AllocationACK messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for AllocationACK.
func Builder(
	allocid *field.AllocIDField,
	tradedate *field.TradeDateField,
	allocstatus *field.AllocStatusField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header().Set(field.NewMsgType("P"))
	builder.Body().Set(allocid)
	builder.Body().Set(tradedate)
	builder.Body().Set(allocstatus)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX40, "P", r
}
