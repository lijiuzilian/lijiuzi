package rpc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	_ "gitlab.33.cn/chain33/chain33/plugin"
	tokenty "gitlab.33.cn/chain33/chain33/plugin/dapp/token/types"
	_ "gitlab.33.cn/chain33/chain33/system"
	"gitlab.33.cn/chain33/chain33/types"
	"gitlab.33.cn/chain33/chain33/util"
	"gitlab.33.cn/chain33/chain33/util/testnode"
)

func TestRPCTokenPreCreate(t *testing.T) {
	// 启动RPCmocker
	mock33 := testnode.New("", nil)
	defer func() {
		mock33.Close()
	}()
	mock33.Listen()
	//precreate
	mock33.SendHot()
	block := mock33.GetLastBlock()
	acc := mock33.GetAccount(block.StateHash, mock33.GetGenesisAddress())
	assert.Equal(t, acc.Balance, 99990000*types.Coin)
	acc = mock33.GetAccount(block.StateHash, mock33.GetHotAddress())
	assert.Equal(t, acc.Balance, 10000*types.Coin)

	//开始发行percreate
	param := tokenty.TokenPreCreate{
		Name:   "Test",
		Symbol: "TEST",
		Total:  10000 * types.Coin,
		Owner:  mock33.GetHotAddress(),
	}
	var txhex string
	err := mock33.GetJsonC().Call("token.CreateRawTokenPreCreateTx", param, &txhex)
	assert.Nil(t, err)
	hash, err := mock33.SendAndSign(mock33.GetHotKey(), txhex)
	assert.Nil(t, err)
	if hash == nil {
		return
	}
	assert.NotNil(t, hash)
	detail, err := mock33.WaitTx(hash)
	assert.Nil(t, err)
	assert.Equal(t, detail.Receipt.Ty, types.ExecOk)
	util.JsonPrint(t, detail)
}
