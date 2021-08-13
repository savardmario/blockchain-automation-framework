package chaincode_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
	"github.com/stretchr/testify/require"
	"drone-reservation/chaincode"
	"drone-reservation/chaincode/mocks"
)

//go:generate counterfeiter -o mocks/transaction.go -fake-name TransactionContext . transactionContext
type transactionContext interface {
	contractapi.TransactionContextInterface
}

//go:generate counterfeiter -o mocks/chaincodestub.go -fake-name ChaincodeStub . chaincodeStub
type chaincodeStub interface {
	shim.ChaincodeStubInterface
}

//go:generate counterfeiter -o mocks/statequeryiterator.go -fake-name StateQueryIterator . stateQueryIterator
type stateQueryIterator interface {
	shim.StateQueryIteratorInterface
}

func TestGivenNoErrorOnBlockchain_whenInitLedger_shouldNotReturnError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	assetTransfer := chaincode.SmartContract{}
	err := assetTransfer.InitLedger(transactionContext)

	require.NoError(t, err)
}

func TestGivenErrorOnBlockchain_whenInitLedger_shouldReturnError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	assetTransfer := chaincode.SmartContract{}
	chaincodeStub.PutStateReturns(fmt.Errorf("failed inserting key"))
	err := assetTransfer.InitLedger(transactionContext)

	require.EqualError(t, err, "failed to put to world state. failed inserting key")
}

func TestGivenValidAsset_whenCreate_shouldNotReturnError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	assetTransfer := chaincode.SmartContract{}
	err := assetTransfer.CreateAsset(transactionContext, "", "", "", "")

	require.NoError(t, err)
}

func TestGivenAssetWithSameIdAlreadyInBlockchain_whenCreate_shouldReturnError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	assetTransfer := chaincode.SmartContract{}
	chaincodeStub.GetStateReturns([]byte{}, nil)
	err := assetTransfer.CreateAsset(transactionContext, "asset1", "", "", "")

	require.EqualError(t, err, "the asset asset1 already exists")

}

func TestGivenBlockchainCodeError_whenCreate_shouldReturnError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	assetTransfer := chaincode.SmartContract{}
	chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve asset"))
	err := assetTransfer.CreateAsset(transactionContext, "asset1", "", "", "")

	require.EqualError(t, err, "failed to read from world state: unable to retrieve asset")
}

func TestGivenAssetPresent_whenRead_shouldReturnRightAsset(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	expectedAsset := &chaincode.Asset{ID: "asset1"}
	bytes, err := json.Marshal(expectedAsset)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(bytes, nil)
	assetTransfer := chaincode.SmartContract{}
	asset, err := assetTransfer.ReadAsset(transactionContext, "")

	require.NoError(t, err)
	require.Equal(t, expectedAsset, asset)
}

func TestGivenBlockchainError_whenRead_shouldReturnError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve asset"))
	assetTransfer := chaincode.SmartContract{}
	_, err := assetTransfer.ReadAsset(transactionContext, "")

	require.EqualError(t, err, "failed to read from world state: unable to retrieve asset")
}

func TestGivenAssetMissing_whenRead_shouldReturnError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	chaincodeStub.GetStateReturns(nil, nil)
	assetTransfer := chaincode.SmartContract{}
	asset, err := assetTransfer.ReadAsset(transactionContext, "asset1")

	require.EqualError(t, err, "the asset asset1 does not exist")
	require.Nil(t, asset)
}

func TestGivenAssetPresent_whenUpdate_shouldReturnNoError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	expectedAsset := &chaincode.Asset{ID: "asset1"}
	bytes, err := json.Marshal(expectedAsset)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(bytes, nil)
	assetTransfer := chaincode.SmartContract{}
	err = assetTransfer.UpdateAsset(transactionContext, "", "", "")

	require.NoError(t, err)
}

func TestGivenAssetMissing_whenUpdate_shouldReturnError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	chaincodeStub.GetStateReturns(nil, nil)
	assetTransfer := chaincode.SmartContract{}
	err := assetTransfer.UpdateAsset(transactionContext, "asset1", "", "")

	require.EqualError(t, err, "the asset asset1 does not exist")
}

func TestGivenChaincodeError_whenUpdate_shouldReturnError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve asset"))
	assetTransfer := chaincode.SmartContract{}
	err := assetTransfer.UpdateAsset(transactionContext, "asset1", "", "")

	require.EqualError(t, err, "failed to read from world state: unable to retrieve asset")
}

func TestGivenAssetPresent_whenDelete_shouldReturnNoError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	asset := &chaincode.Asset{ID: "asset1"}
	bytes, err := json.Marshal(asset)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(bytes, nil)
	chaincodeStub.DelStateReturns(nil)
	assetTransfer := chaincode.SmartContract{}
	err = assetTransfer.DeleteAsset(transactionContext, "")

	require.NoError(t, err)
}

func TestGivenAssetMissing_whenDelete_shouldReturnError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	chaincodeStub.GetStateReturns(nil, nil)
	assetTransfer := chaincode.SmartContract{}
	err := assetTransfer.DeleteAsset(transactionContext, "asset1")

	require.EqualError(t, err, "the asset asset1 does not exist")
}

func TestGivenChaincodeError_whenDelete_shouldReturnError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve asset"))
	assetTransfer := chaincode.SmartContract{}
	err := assetTransfer.DeleteAsset(transactionContext, "")

	require.EqualError(t, err, "failed to read from world state: unable to retrieve asset")
}

func TestGivenAssetPresent_whenTransfer_shouldReturnNoError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	asset := &chaincode.Asset{ID: "asset1"}
	bytes, err := json.Marshal(asset)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(bytes, nil)
	assetTransfer := chaincode.SmartContract{}
	err = assetTransfer.TransferAsset(transactionContext, "", "")

	require.NoError(t, err)
}

func TestGivenErrorReading_whenTransfer_shouldReturnError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	assetTransfer := chaincode.SmartContract{}
	chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve asset"))
	err := assetTransfer.TransferAsset(transactionContext, "", "")

	require.EqualError(t, err, "failed to read from world state: unable to retrieve asset")
}

func TestGivenNoErrorOnBlockchain_whenGettingAllAssets_shouldReturnNoError(t *testing.T) {
	asset := &chaincode.Asset{ID: "asset1", Date: "1994-11-05T13:15:30Z", Request: "125, 134, 898", Status: "requested", Owner: "ClientOrg", Location: ""}
	bytes, err := json.Marshal(asset)
	require.NoError(t, err)

	iterator := &mocks.StateQueryIterator{}
	iterator.HasNextReturnsOnCall(0, true)
	iterator.HasNextReturnsOnCall(1, false)
	iterator.NextReturns(&queryresult.KV{Value: bytes}, nil)

	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	chaincodeStub.GetStateByRangeReturns(iterator, nil)
	assetTransfer := &chaincode.SmartContract{}
	assets, err := assetTransfer.GetAllAssets(transactionContext)

	require.NoError(t, err)
	require.Equal(t, []*chaincode.Asset{asset}, assets)
}

func TestGivenErrorWhileIterating_whenGettingAllAssets_shouldReturnError(t *testing.T) {
	iterator := &mocks.StateQueryIterator{}
	iterator.HasNextReturns(true)
	iterator.NextReturns(nil, fmt.Errorf("failed retrieving next item"))

	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	chaincodeStub.GetStateByRangeReturns(iterator, nil)
	assetTransfer := &chaincode.SmartContract{}
	assets, err := assetTransfer.GetAllAssets(transactionContext)

	require.EqualError(t, err, "failed retrieving next item")
	require.Nil(t, assets)
}

func TestGivenErrorWhenGettingRange_whenGettingAllAssets_shouldReturnError(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	chaincodeStub.GetStateByRangeReturns(nil, fmt.Errorf("failed retrieving all assets"))
	assetTransfer := &chaincode.SmartContract{}
	assets, err := assetTransfer.GetAllAssets(transactionContext)

	require.EqualError(t, err, "failed retrieving all assets")
	require.Nil(t, assets)
}