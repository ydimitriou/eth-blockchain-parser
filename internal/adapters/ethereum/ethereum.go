package ethereum

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ydimitriou/eth-blockchain-parser/internal/app/parser"
)

const (
	baseURL               = "https://cloudflare-eth.com"
	getCurrentBlockMethod = "eth_blockNumber"
	getBlockByNumber      = "eth_getBlockByNumber"
	jsonRpc               = "2.0"
)

type RpcRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

type GetCurrenBlockResponse struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}

type EthereumService struct {
	baseURL    string
	httpClient *http.Client
}

func NewEthereumService() *EthereumService {
	httpClient := &http.Client{}

	return &EthereumService{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

func (e *EthereumService) GetCurrentBlock(ctx context.Context) (*string, error) {
	req := RpcRequest{
		Jsonrpc: jsonRpc,
		Method:  getCurrentBlockMethod,
		Id:      83,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, e.baseURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	body, err := e.makeRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var res GetCurrenBlockResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &res.Result, nil
}

func (e *EthereumService) GetBlockByNumber(ctx context.Context, blockNumber string) (*parser.GetBlockByNumberResponse, error) {
	req := RpcRequest{
		Jsonrpc: jsonRpc,
		Method:  getBlockByNumber,
		Params:  []interface{}{blockNumber, true},
		Id:      1,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, e.baseURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	body, err := e.makeRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var res parser.GetBlockByNumberResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &res, nil
}

func (e *EthereumService) makeRequest(httpReq *http.Request) ([]byte, error) {
	res, err := e.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("http client.Do: %w", err)
	}

	return extractResponseBody(res)
}

func extractResponseBody(httpRes *http.Response) ([]byte, error) {
	defer func() {
		err := httpRes.Body.Close()
		// handle this erros better later
		if err != nil {
			log.Fatalln(err)
		}
	}()

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	return body, nil
}
