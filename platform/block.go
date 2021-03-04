package platform

import (
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/vms/components/core"
	"github.com/vmihailenco/msgpack/v5"
)

//Block - a unit block
type Block struct {
	*core.Block `serialize:"true"`
	Timestamp   int64 `serialize:"true"`
	Data        [64]byte
}

//ID - return block's unique ID
func (b Block) ID() ids.ID {

	return b.Metadata.ID()
}

//Status - return status of the block
func (b Block) Status() choices.Status {
	return b.Status()
}

//Parent - return the block's parent
func (b Block) Parent() snowman.Block {
	block, err := b.VM.GetBlock(b.PrntID)
	if err != nil {
		return nil
	}
	return block
}

//Verify - verify the block
func (b Block) Verify() (err error) {
	if accepted, err := b.Block.Verify(); err != nil || accepted {
		return err
	}

	parent, ok := b.Parent().(*Block)
	if !ok {
		return errors.New("Error converting interface to Block")
	}

	if b.Timestamp < time.Unix(parent.Timestamp, 0).Unix() {
		return errors.New("block's timestamp is before parent's timestamp")
	}

	if b.Timestamp >= time.Now().Add(time.Hour).Unix() {
		return errors.New("block's timestamp is more than 1 hour ahead of local time")
	}

	b.VM.SaveBlock(b.VM.DB, b)

	return b.VM.DB.Commit()
}

//Bytes - return bytes representation of block
func (b Block) Bytes() []byte {
	bs, _ := msgpack.Marshal(b)
	return bs
}
