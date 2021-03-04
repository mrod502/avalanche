package platform

import (
	"encoding/json"
	"sync"

	"github.com/vmihailenco/msgpack/v5"
)

//Serializer - handles encoding and decoding of objects
type Serializer interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}

//MsgpackSerializer - uses msgpack
type MsgpackSerializer struct {
	mux *sync.Mutex
}

//Marshal - to bytes
func (m MsgpackSerializer) Marshal(v interface{}) ([]byte, error) {
	m.mux.Lock()
	defer m.mux.Unlock()
	return msgpack.Marshal(v)
}

//Unmarshal - from bytes
func (m MsgpackSerializer) Unmarshal(b []byte, v interface{}) error {
	m.mux.Lock()
	defer m.mux.Unlock()
	return msgpack.Unmarshal(b, v)
}

//NewMsgpackSerializer - returns initialized serializer
func NewMsgpackSerializer() *MsgpackSerializer {
	return &MsgpackSerializer{mux: new(sync.Mutex)}
}

//JSONSerializer - uses JSON
type JSONSerializer struct {
	mux *sync.Mutex
}

//Marshal - to bytes
func (j JSONSerializer) Marshal(v interface{}) ([]byte, error) {
	j.mux.Lock()
	defer j.mux.Unlock()
	return json.Marshal(v)
}

//Unmarshal - from bytes
func (j JSONSerializer) Unmarshal(b []byte, v interface{}) error {
	j.mux.Lock()
	defer j.mux.Unlock()
	return json.Unmarshal(b, v)
}

//NewJSONSerializer - returns initialized serializer
func NewJSONSerializer() *JSONSerializer {
	return &JSONSerializer{mux: new(sync.Mutex)}
}
