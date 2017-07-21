package client

import "github.com/pkg/errors"

type Block struct {
	Hash              string
	Confirmations     int
	Size              int
	Height            int
	Version           int
	Payload           string
	PayloadHash       string
	MerkleRoot        string
	Time              int64
	MedianTime        int64
	Creator           string
	CreatorSignature  string
	Signatures        int
	AdminSignatures   int
	ChainSignature    string
	AdminMultiSig     string
	MissingCreatorIds []string
	AdminSignerIds    []string
	Tx                []string
	PreviousBlockHash string
	NextBlockHash     string
}

// GetBlockCount is the wrapper around getblockcount rpc
func (c *Client) GetBlockCount() (int, error) {
	response, err := c.c.Call("getblockcount")
	if err != nil {
		return 0, errors.Wrap(err, "getblockcount: call")
	}

	var blockCount int
	err = response.GetObject(&blockCount)
	if err != nil {
		return 0, errors.Wrap(err, "getblockcount: unmarshal")
	}

	return blockCount, nil
}

// GetBlockHash is the wrapper around getblockhash rpc
func (c *Client) GetBlockHash(index int) (string, error) {
	response, err := c.c.Call("getblockhash", index)
	if err != nil {
		return "", errors.Wrap(err, "getblockhash: call")
	}

	var blockHash string
	err = response.GetObject(&blockHash)
	if err != nil {
		return "", errors.Wrap(err, "getblockhash: unmarshal")
	}

	return blockHash, nil

}

// GetBlock is the wrapper around getblock rpc
func (c *Client) GetBlock(hash string, verbose bool, mode int) (*Block, error) {
	response, err := c.c.Call("getblock", hash)
	if err != nil {
		return nil, errors.Wrap(err, "getblock: call")
	}

	block := Block{}
	err = response.GetObject(&block)
	if err != nil {
		return nil, errors.Wrap(err, "getblock: unmarshal")
	}

	return &block, nil
}
