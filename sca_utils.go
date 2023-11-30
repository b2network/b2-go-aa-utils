package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	scaRegistryABI    = `[{"constant":true,"inputs":[{"name":"id","type":"bytes32"}],"name":"getSCAAddress","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"}]`
	kernelFactoryABI  = `[{"constant":true,"inputs":[{"name":"data","type":"bytes"},{"name":"index","type":"uint256"}],"name":"getAccountAddress","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"}]`
	zeroAddressString = "0x0000000000000000000000000000000000000000"
)

func getDeterministicAddress(client *ethclient.Client, factoryAddress common.Address, ownerHash common.Hash) (common.Address, error) {
	parsedABI, err := abi.JSON(strings.NewReader(kernelFactoryABI))
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to parse ABI: %v", err)
	}

	// convert ownerHash to account index
	accountIndex := new(big.Int)
	accountIndex.SetBytes(ownerHash[:])

	callData, err := parsedABI.Pack("getAccountAddress", []byte{}, accountIndex)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to pack getAccountAddress call: %v", err)
	}

	callMsg := ethereum.CallMsg{
		To:   &factoryAddress,
		Data: callData,
	}

	outBz, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to call contract: %v", err)
	}
	parsedOutputs, err := parsedABI.Unpack("getAccountAddress", outBz)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to parse outputs: %v", err)
	}

	return parsedOutputs[0].(common.Address), nil
}

func GetSCAAddress(client *ethclient.Client, scaRegistryAddress common.Address, factoryAddress common.Address, owner string) (common.Address, error) {
	parsedABI, err := abi.JSON(strings.NewReader(scaRegistryABI))
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to parse ABI: %v", err)
	}

	ownerHash := crypto.Keccak256Hash([]byte(strings.ToLower(owner)))

	callData, err := parsedABI.Pack("getSCAAddress", ownerHash)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to pack getSCAAddress call: %v", err)
	}

	callMsg := ethereum.CallMsg{
		To:   &scaRegistryAddress,
		Data: callData,
	}

	outBz, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to call contract: %v", err)
	}

	parsedOutputs, err := parsedABI.Unpack("getSCAAddress", outBz)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to unpack function output: %v", err)
	}

	registeredAddress := parsedOutputs[0].(common.Address)
	// if ownerHash is not registered, return derived address
	if registeredAddress == common.HexToAddress(zeroAddressString) {
		return getDeterministicAddress(client, factoryAddress, ownerHash)
	}

	return registeredAddress, nil
}
