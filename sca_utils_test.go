package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

const (
	B2DevRpcUrl   = "http://43.135.203.73:8123"
	SCARegistry   = "0x231aec684Ad0e63c2F4d176EddCE97A1B666247c"
	KernelFactory = "0x7516283Ff7090B8286E23a16f8b5b35B3ba541A2"
)

func TestGetFreshSCAAddress(t *testing.T) {
	client, err := ethclient.Dial(B2DevRpcUrl)
	assert.NoError(t, err)

	scaRegistryAddress := common.HexToAddress(SCARegistry)
	factoryAddress := common.HexToAddress(KernelFactory)

	result, err := GetSCAAddress(client, scaRegistryAddress, factoryAddress, "0x223A645679A72E7cE7ef250f10c77cFbA5d75cc7")
	assert.NoError(t, err)
	assert.Equal(t, common.HexToAddress("0x4319f8b4B5373a783D69c2F4E35557F4E5775a43"), result)

	result, err = GetSCAAddress(client, scaRegistryAddress, factoryAddress, "bc1qe40k9zyjyndl2t7f0fxws7h94pjrfz0zan5yak")
	assert.NoError(t, err)
	assert.Equal(t, common.HexToAddress("0xE86eB17dD5eCaCD2FebA905DDA25127723F4EC7c"), result)
}
