package platform

import (
	"time"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/vms/components/core"
)

//VM - the virtual machine
type VM struct {
	core.SnowmanVM
	db         *database.Database
	serializer Serializer
	mempool    [][64]byte
}

//Initialize the VM
func (v *VM) Initialize(
	ctx *snow.Context,
	db database.Database,
	genesisBytes []byte,
	toEngine chan<- common.Message,
	fxs []*common.Fx,
) (err error) {
	return
}

//NewVM - return an initialized vm
func NewVM(
	ctx *snow.Context,
	db database.Database,
	genesisBytes []byte,
	toEngine chan<- common.Message,
	fxs []*common.Fx,
) (vm *VM, err error) {

	vm = &VM{}

	err = vm.Initialize(ctx, db, genesisBytes, toEngine, fxs)

	return
}

//Shutdown - turn off the vm
func (v *VM) Shutdown() {
	return
}

func (v *VM) CreateHandlers() (m map[string]*common.HTTPHandler) {

	m = make(map[string]*common.HTTPHandler)

	return
}

func (v *VM) BuildBlock() (b snowman.Block, err error) { return }

func (v *VM) ParseBlock(b []byte) (bb snowman.Block, err error) {
	block := new(Block)

	err = v.serializer.Unmarshal(b, block)
	bb = *block

	return
}

func (v *VM) GetBlock(id ids.ID) (b snowman.Block, err error) { return }

func (v *VM) SetPreference(id ids.ID) {}

func (v *VM) LastAccepted() (id ids.ID) { return }

func (v *VM) NewBlock(parentID ids.ID, data [64]byte, timestamp time.Time) (b *Block, err error) {

	b = &Block{
		Block:     core.NewBlock(parentID, b.Parent().Height()+1),
		Data:      data,
		Timestamp: timestamp.Unix(),
	}
	var blockBytes []byte
	blockBytes, err = v.serializer.Marshal(b)
	if err != nil {
		return nil, err
	}
	b.Initialize(blockBytes, &v.SnowmanVM)

	return
}
