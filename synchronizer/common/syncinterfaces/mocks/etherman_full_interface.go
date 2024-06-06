// Code generated by mockery. DO NOT EDIT.

package mock_syncinterfaces

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	etherman "github.com/0xPolygonHermez/zkevm-node/etherman"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// EthermanFullInterface is an autogenerated mock type for the EthermanFullInterface type
type EthermanFullInterface struct {
	mock.Mock
}

type EthermanFullInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *EthermanFullInterface) EXPECT() *EthermanFullInterface_Expecter {
	return &EthermanFullInterface_Expecter{mock: &_m.Mock}
}

// EthBlockByNumber provides a mock function with given fields: ctx, blockNumber
func (_m *EthermanFullInterface) EthBlockByNumber(ctx context.Context, blockNumber uint64) (*types.Block, error) {
	ret := _m.Called(ctx, blockNumber)

	if len(ret) == 0 {
		panic("no return value specified for EthBlockByNumber")
	}

	var r0 *types.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*types.Block, error)); ok {
		return rf(ctx, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *types.Block); ok {
		r0 = rf(ctx, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthermanFullInterface_EthBlockByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EthBlockByNumber'
type EthermanFullInterface_EthBlockByNumber_Call struct {
	*mock.Call
}

// EthBlockByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - blockNumber uint64
func (_e *EthermanFullInterface_Expecter) EthBlockByNumber(ctx interface{}, blockNumber interface{}) *EthermanFullInterface_EthBlockByNumber_Call {
	return &EthermanFullInterface_EthBlockByNumber_Call{Call: _e.mock.On("EthBlockByNumber", ctx, blockNumber)}
}

func (_c *EthermanFullInterface_EthBlockByNumber_Call) Run(run func(ctx context.Context, blockNumber uint64)) *EthermanFullInterface_EthBlockByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *EthermanFullInterface_EthBlockByNumber_Call) Return(_a0 *types.Block, _a1 error) *EthermanFullInterface_EthBlockByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthermanFullInterface_EthBlockByNumber_Call) RunAndReturn(run func(context.Context, uint64) (*types.Block, error)) *EthermanFullInterface_EthBlockByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetFinalizedBlockNumber provides a mock function with given fields: ctx
func (_m *EthermanFullInterface) GetFinalizedBlockNumber(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetFinalizedBlockNumber")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthermanFullInterface_GetFinalizedBlockNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFinalizedBlockNumber'
type EthermanFullInterface_GetFinalizedBlockNumber_Call struct {
	*mock.Call
}

// GetFinalizedBlockNumber is a helper method to define mock.On call
//   - ctx context.Context
func (_e *EthermanFullInterface_Expecter) GetFinalizedBlockNumber(ctx interface{}) *EthermanFullInterface_GetFinalizedBlockNumber_Call {
	return &EthermanFullInterface_GetFinalizedBlockNumber_Call{Call: _e.mock.On("GetFinalizedBlockNumber", ctx)}
}

func (_c *EthermanFullInterface_GetFinalizedBlockNumber_Call) Run(run func(ctx context.Context)) *EthermanFullInterface_GetFinalizedBlockNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *EthermanFullInterface_GetFinalizedBlockNumber_Call) Return(_a0 uint64, _a1 error) *EthermanFullInterface_GetFinalizedBlockNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthermanFullInterface_GetFinalizedBlockNumber_Call) RunAndReturn(run func(context.Context) (uint64, error)) *EthermanFullInterface_GetFinalizedBlockNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestBatchNumber provides a mock function with given fields:
func (_m *EthermanFullInterface) GetLatestBatchNumber() (uint64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLatestBatchNumber")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthermanFullInterface_GetLatestBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestBatchNumber'
type EthermanFullInterface_GetLatestBatchNumber_Call struct {
	*mock.Call
}

// GetLatestBatchNumber is a helper method to define mock.On call
func (_e *EthermanFullInterface_Expecter) GetLatestBatchNumber() *EthermanFullInterface_GetLatestBatchNumber_Call {
	return &EthermanFullInterface_GetLatestBatchNumber_Call{Call: _e.mock.On("GetLatestBatchNumber")}
}

func (_c *EthermanFullInterface_GetLatestBatchNumber_Call) Run(run func()) *EthermanFullInterface_GetLatestBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EthermanFullInterface_GetLatestBatchNumber_Call) Return(_a0 uint64, _a1 error) *EthermanFullInterface_GetLatestBatchNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthermanFullInterface_GetLatestBatchNumber_Call) RunAndReturn(run func() (uint64, error)) *EthermanFullInterface_GetLatestBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestVerifiedBatchNum provides a mock function with given fields:
func (_m *EthermanFullInterface) GetLatestVerifiedBatchNum() (uint64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLatestVerifiedBatchNum")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthermanFullInterface_GetLatestVerifiedBatchNum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestVerifiedBatchNum'
type EthermanFullInterface_GetLatestVerifiedBatchNum_Call struct {
	*mock.Call
}

// GetLatestVerifiedBatchNum is a helper method to define mock.On call
func (_e *EthermanFullInterface_Expecter) GetLatestVerifiedBatchNum() *EthermanFullInterface_GetLatestVerifiedBatchNum_Call {
	return &EthermanFullInterface_GetLatestVerifiedBatchNum_Call{Call: _e.mock.On("GetLatestVerifiedBatchNum")}
}

func (_c *EthermanFullInterface_GetLatestVerifiedBatchNum_Call) Run(run func()) *EthermanFullInterface_GetLatestVerifiedBatchNum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EthermanFullInterface_GetLatestVerifiedBatchNum_Call) Return(_a0 uint64, _a1 error) *EthermanFullInterface_GetLatestVerifiedBatchNum_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthermanFullInterface_GetLatestVerifiedBatchNum_Call) RunAndReturn(run func() (uint64, error)) *EthermanFullInterface_GetLatestVerifiedBatchNum_Call {
	_c.Call.Return(run)
	return _c
}

// GetRollupInfoByBlockRange provides a mock function with given fields: ctx, fromBlock, toBlock
func (_m *EthermanFullInterface) GetRollupInfoByBlockRange(ctx context.Context, fromBlock uint64, toBlock *uint64) ([]etherman.Block, map[common.Hash][]etherman.Order, error) {
	ret := _m.Called(ctx, fromBlock, toBlock)

	if len(ret) == 0 {
		panic("no return value specified for GetRollupInfoByBlockRange")
	}

	var r0 []etherman.Block
	var r1 map[common.Hash][]etherman.Order
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *uint64) ([]etherman.Block, map[common.Hash][]etherman.Order, error)); ok {
		return rf(ctx, fromBlock, toBlock)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *uint64) []etherman.Block); ok {
		r0 = rf(ctx, fromBlock, toBlock)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]etherman.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, *uint64) map[common.Hash][]etherman.Order); ok {
		r1 = rf(ctx, fromBlock, toBlock)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[common.Hash][]etherman.Order)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint64, *uint64) error); ok {
		r2 = rf(ctx, fromBlock, toBlock)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EthermanFullInterface_GetRollupInfoByBlockRange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRollupInfoByBlockRange'
type EthermanFullInterface_GetRollupInfoByBlockRange_Call struct {
	*mock.Call
}

// GetRollupInfoByBlockRange is a helper method to define mock.On call
//   - ctx context.Context
//   - fromBlock uint64
//   - toBlock *uint64
func (_e *EthermanFullInterface_Expecter) GetRollupInfoByBlockRange(ctx interface{}, fromBlock interface{}, toBlock interface{}) *EthermanFullInterface_GetRollupInfoByBlockRange_Call {
	return &EthermanFullInterface_GetRollupInfoByBlockRange_Call{Call: _e.mock.On("GetRollupInfoByBlockRange", ctx, fromBlock, toBlock)}
}

func (_c *EthermanFullInterface_GetRollupInfoByBlockRange_Call) Run(run func(ctx context.Context, fromBlock uint64, toBlock *uint64)) *EthermanFullInterface_GetRollupInfoByBlockRange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(*uint64))
	})
	return _c
}

func (_c *EthermanFullInterface_GetRollupInfoByBlockRange_Call) Return(_a0 []etherman.Block, _a1 map[common.Hash][]etherman.Order, _a2 error) *EthermanFullInterface_GetRollupInfoByBlockRange_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EthermanFullInterface_GetRollupInfoByBlockRange_Call) RunAndReturn(run func(context.Context, uint64, *uint64) ([]etherman.Block, map[common.Hash][]etherman.Order, error)) *EthermanFullInterface_GetRollupInfoByBlockRange_Call {
	_c.Call.Return(run)
	return _c
}

// GetTrustedSequencerURL provides a mock function with given fields:
func (_m *EthermanFullInterface) GetTrustedSequencerURL() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTrustedSequencerURL")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthermanFullInterface_GetTrustedSequencerURL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTrustedSequencerURL'
type EthermanFullInterface_GetTrustedSequencerURL_Call struct {
	*mock.Call
}

// GetTrustedSequencerURL is a helper method to define mock.On call
func (_e *EthermanFullInterface_Expecter) GetTrustedSequencerURL() *EthermanFullInterface_GetTrustedSequencerURL_Call {
	return &EthermanFullInterface_GetTrustedSequencerURL_Call{Call: _e.mock.On("GetTrustedSequencerURL")}
}

func (_c *EthermanFullInterface_GetTrustedSequencerURL_Call) Run(run func()) *EthermanFullInterface_GetTrustedSequencerURL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EthermanFullInterface_GetTrustedSequencerURL_Call) Return(_a0 string, _a1 error) *EthermanFullInterface_GetTrustedSequencerURL_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthermanFullInterface_GetTrustedSequencerURL_Call) RunAndReturn(run func() (string, error)) *EthermanFullInterface_GetTrustedSequencerURL_Call {
	_c.Call.Return(run)
	return _c
}

// HeaderByNumber provides a mock function with given fields: ctx, number
func (_m *EthermanFullInterface) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ret := _m.Called(ctx, number)

	if len(ret) == 0 {
		panic("no return value specified for HeaderByNumber")
	}

	var r0 *types.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (*types.Header, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Header); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthermanFullInterface_HeaderByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HeaderByNumber'
type EthermanFullInterface_HeaderByNumber_Call struct {
	*mock.Call
}

// HeaderByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - number *big.Int
func (_e *EthermanFullInterface_Expecter) HeaderByNumber(ctx interface{}, number interface{}) *EthermanFullInterface_HeaderByNumber_Call {
	return &EthermanFullInterface_HeaderByNumber_Call{Call: _e.mock.On("HeaderByNumber", ctx, number)}
}

func (_c *EthermanFullInterface_HeaderByNumber_Call) Run(run func(ctx context.Context, number *big.Int)) *EthermanFullInterface_HeaderByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*big.Int))
	})
	return _c
}

func (_c *EthermanFullInterface_HeaderByNumber_Call) Return(_a0 *types.Header, _a1 error) *EthermanFullInterface_HeaderByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthermanFullInterface_HeaderByNumber_Call) RunAndReturn(run func(context.Context, *big.Int) (*types.Header, error)) *EthermanFullInterface_HeaderByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// L1ChainID provides a mock function with given fields: ctx
func (_m *EthermanFullInterface) L1ChainID(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for L1ChainID")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthermanFullInterface_L1ChainID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'L1ChainID'
type EthermanFullInterface_L1ChainID_Call struct {
	*mock.Call
}

// L1ChainID is a helper method to define mock.On call
//   - ctx context.Context
func (_e *EthermanFullInterface_Expecter) L1ChainID(ctx interface{}) *EthermanFullInterface_L1ChainID_Call {
	return &EthermanFullInterface_L1ChainID_Call{Call: _e.mock.On("L1ChainID", ctx)}
}

func (_c *EthermanFullInterface_L1ChainID_Call) Run(run func(ctx context.Context)) *EthermanFullInterface_L1ChainID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *EthermanFullInterface_L1ChainID_Call) Return(_a0 uint64, _a1 error) *EthermanFullInterface_L1ChainID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthermanFullInterface_L1ChainID_Call) RunAndReturn(run func(context.Context) (uint64, error)) *EthermanFullInterface_L1ChainID_Call {
	_c.Call.Return(run)
	return _c
}

// VerifyGenBlockNumber provides a mock function with given fields: ctx, genBlockNumber
func (_m *EthermanFullInterface) VerifyGenBlockNumber(ctx context.Context, genBlockNumber uint64) (bool, error) {
	ret := _m.Called(ctx, genBlockNumber)

	if len(ret) == 0 {
		panic("no return value specified for VerifyGenBlockNumber")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (bool, error)); ok {
		return rf(ctx, genBlockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) bool); ok {
		r0 = rf(ctx, genBlockNumber)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, genBlockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthermanFullInterface_VerifyGenBlockNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyGenBlockNumber'
type EthermanFullInterface_VerifyGenBlockNumber_Call struct {
	*mock.Call
}

// VerifyGenBlockNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - genBlockNumber uint64
func (_e *EthermanFullInterface_Expecter) VerifyGenBlockNumber(ctx interface{}, genBlockNumber interface{}) *EthermanFullInterface_VerifyGenBlockNumber_Call {
	return &EthermanFullInterface_VerifyGenBlockNumber_Call{Call: _e.mock.On("VerifyGenBlockNumber", ctx, genBlockNumber)}
}

func (_c *EthermanFullInterface_VerifyGenBlockNumber_Call) Run(run func(ctx context.Context, genBlockNumber uint64)) *EthermanFullInterface_VerifyGenBlockNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *EthermanFullInterface_VerifyGenBlockNumber_Call) Return(_a0 bool, _a1 error) *EthermanFullInterface_VerifyGenBlockNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthermanFullInterface_VerifyGenBlockNumber_Call) RunAndReturn(run func(context.Context, uint64) (bool, error)) *EthermanFullInterface_VerifyGenBlockNumber_Call {
	_c.Call.Return(run)
	return _c
}

// NewEthermanFullInterface creates a new instance of EthermanFullInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEthermanFullInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *EthermanFullInterface {
	mock := &EthermanFullInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
