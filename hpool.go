// Package Hpool is an implementation of the Hpool API in Golang.
package hpool

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const (
	API_BASE = "https://www.hpool.com/api" // Hpool API endpoint
)

// New returns an instantiated hpool struct
func New(accessToken string) *Hpool {
	client := NewClient(accessToken)
	return &Hpool{client}
}

// NewWithCustomHttpClient returns an instantiated hpool struct with custom http client
func NewWithCustomHttpClient(accessToken string, httpClient *http.Client) *Hpool {
	client := NewClientWithCustomHttpConfig(accessToken, httpClient)
	return &Hpool{client}
}

// NewWithCustomTimeout returns an instantiated hpool struct with custom timeout
func NewWithCustomTimeout(accessToken string, timeout time.Duration) *Hpool {
	client := NewClientWithCustomTimeout(accessToken, timeout)
	return &Hpool{client}
}

// handleErr gets JSON response from hpool API en deal with error
func handleErr(r jsonResponse) error {
	if r.Code != 200 {
		return errors.New(r.Message)
	}
	return nil
}

// hpool represent a hpool client
type Hpool struct {
	client *client
}

// set enable/disable http request/response dump
func (i *Hpool) SetDebug(enable bool) {
	i.client.debug = enable
}

type Pool struct {
	ApiKey                          string  `json:"api_key"`
	BlockReward                     float64 `json:"block_reward,string"`
	BlockTime                       int     `json:"block_time"`
	Capacity                        int     `json:"capacity"`
	Coin                            string  `json:"coin"`
	DepositMortgageBalance          float64 `json:"deposit_mortgage_balance,string"`
	DepositMortgageEffectiveBalance float64 `json:"deposit_mortgage_effective_balance,string"`
	DepositMortgageFreeBalance      float64 `json:"deposit_mortgage_free_balance,string"`
	DepositRate                     float64 `json:"deposit_rate,string"`
	Fee                             float64 `json:"fee"`
	LoanMortgageBalance             float64 `json:"loan_mortgage_balance,string"`
	Mortgage                        float64 `json:"mortgage,string"`
	Name                            string  `json:"name"`
	Offline                         int     `json:"offline"`
	Online                          int     `json:"online"`
	PaymentTime                     string  `json:"payment_time"`
	PointDepositBalance             float64 `json:"point_deposit_balance,string"`
	PoolAddress                     string  `json:"pool_address"`
	PoolIncome                      float64 `json:"pool_income,string"`
	PoolType                        string  `json:"pool_type"`
	PreviousIncomePb                float64 `json:"previous_income_pb,string"`
	TheoryMortgageBalance           float64 `json:"theory_mortgage_balance,string"`
	Type                            string  `json:"type"`
	UndistributedIncome             float64 `json:"undistributed_income,string"`
}

type PoolType int

const (
	Opened PoolType = iota + 1
	All
)

func (w PoolType) String() string {
	return [...]string{"opened", "all"}[w-1]
}

func (i *Hpool) PoolList(poolType PoolType) (pools []Pool, err error) {
	payload := map[string]string{
		"type": poolType.String(),
	}
	r, err := i.client.do("GET", "pool/list", payload, true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var data listData
	if err = json.Unmarshal(response.Data, &data); err != nil {
		return
	}
	err = json.Unmarshal(data.List, &pools)
	return
}

func (i *Hpool) PoolDetail(pool string) (poolDetail Pool, err error) {
	payload := map[string]string{
		"type": pool,
		"page": "1",
		//"count"
	}
	r, err := i.client.do("GET", "pool/detail", payload, true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Data, &poolDetail)
	return
}

type Miner struct {
	Name       string `json:"miner_name"`
	Id         string `json:"id"`
	Capacity   int    `json:"capacity"`
	Online     bool   `json:"online"`
	UpdateTime int    `json:"update_time"`
}

func (i *Hpool) Miners(pool string) (miners []Miner, err error) {
	payload := map[string]string{
		"type": pool,
		"page": "1",
		//"count"
	}
	r, err := i.client.do("GET", "pool/miner", payload, true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var data listData
	if err = json.Unmarshal(response.Data, &data); err != nil {
		return
	}
	err = json.Unmarshal(data.List, &miners)
	return
}

type Plot struct {
	Capacity  int64  `json:"capacity"`
	PublicKey string `json:"public_key"`
	Size      int64  `json:"size"`
	Status    string `json:"status"`
	Uuid      string `json:"uuid"`
	UpdatedAt int    `json:"updated_at"`
}

func (i *Hpool) Plots(pool string) (plots []Plot, err error) {
	payload := map[string]string{
		"pool":  pool,
		"page":  "1",
		"count": "100",
	}
	r, err := i.client.do("GET", "pool/GetPlots", payload, true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var data listData
	if err = json.Unmarshal(response.Data, &data); err != nil {
		return
	}
	err = json.Unmarshal(data.List, &plots)
	return
}

type MiningIncome struct {
	Amount     float64 `json:"amount,string"`
	Coin       string  `json:"coin"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	RecordTime int     `json:"record_time"`
}

func (i *Hpool) MiningIncome(pool string) (miningIncomes []MiningIncome, err error) {
	payload := map[string]string{
		"type":  pool,
		"page":  "1",
		"count": "100",
	}
	r, err := i.client.do("GET", "pool/miningincomerecord", payload, true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var data listData
	if err = json.Unmarshal(response.Data, &data); err != nil {
		return
	}
	err = json.Unmarshal(data.List, &miningIncomes)
	return
}

type Assets struct {
	Balance           float64 `json:"balance,string"`
	CooperationAmount float64 `json:"cooperation_amount,string"`
	DepositAmount     float64 `json:"deposit_amount,string"`
	FreezeBalance     float64 `json:"freeze_balance,string"`
	Name              string  `json:"name"`
	Type              string  `json:"type"`
	TotalAssets       float64 `json:"total_assets,string"`
	WithdrawAmount    float64 `json:"withdraw_amount,string"`
}

func (i *Hpool) Totalassets() (assets []Assets, err error) {
	r, err := i.client.do("GET", "assets/totalassets", nil, true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var data listData
	if err = json.Unmarshal(response.Data, &data); err != nil {
		return
	}
	err = json.Unmarshal(data.List, &assets)
	return
}
