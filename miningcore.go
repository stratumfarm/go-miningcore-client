package miningcore

import (
	"context"
	"fmt"
	"net/http"
)

// GetPools returns a list of all available pools.
func (c *Client) GetPools(ctx context.Context) ([]*PoolInfo, int, error) {
	e := "/api/pools"
	var res struct {
		Pools []*PoolInfo `json:"pools"`
	}
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil)
	if err != nil {
		return nil, s, err
	}
	return res.Pools, s, nil
}

// GetPool returns information about a specific pool.
func (c *Client) GetPool(ctx context.Context, id string) (*PoolInfo, int, error) {
	e := "/api/pools/" + id
	var res struct {
		Pool PoolInfo `json:"pool"`
	}
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil)
	if err != nil {
		return nil, s, err
	}
	return &res.Pool, s, nil
}

// GetPoolBlocks returns a list of blocks found by a pool.
func (c *Client) GetPoolBlocks(ctx context.Context, id string) ([]*Block, int, error) {
	e := fmt.Sprintf("/api/pools/%s/blocks", id)
	var res []*Block
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil)
	if err != nil {
		return nil, s, err
	}
	return res, s, nil
}

// GetPoolPayments returns a list of payments made by a pool.
func (c *Client) GetPoolPayments(ctx context.Context, id string) ([]*Payment, int, error) {
	e := fmt.Sprintf("/api/pools/%s/payments", id)
	var res []*Payment
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil)
	if err != nil {
		return nil, s, err
	}
	return res, s, nil
}

// GetMiner returns information about a specific miner in a pool.
func (c *Client) GetMiner(ctx context.Context, id, addr string) (*MinerStats, int, error) {
	e := fmt.Sprintf("/api/pools/%s/miners/%s", id, addr)
	var res MinerStats
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil)
	if err != nil {
		return nil, s, err
	}
	return &res, s, nil
}

// GetMinerPayments returns a list of payments of a miner.
func (c *Client) GetMinerPayments(ctx context.Context, id, addr string) ([]*Payment, int, error) {
	e := fmt.Sprintf("/api/pools/%s/miners/%s/payments", id, addr)
	var res []*Payment
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil)
	if err != nil {
		return nil, s, err
	}
	return res, s, nil
}

// GetMinerDailyEarnings returns a list of daily earnings of a miner.
func (c *Client) GetMinerDailyEarnings(ctx context.Context, id, addr string) ([]*DailyEarning, int, error) {
	e := fmt.Sprintf("/api/pools/%s/miners/%s/dailyearnings", id, addr)
	var res []*DailyEarning
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil)
	if err != nil {
		return nil, s, err
	}
	return res, s, nil
}

// GetMinerBalanceChanges returns a list of balance changes of a miner.
func (c *Client) GetMinerBalanceChanges(ctx context.Context, id, addr string) ([]*BalanceChange, int, error) {
	e := fmt.Sprintf("/api/pools/%s/miners/%s/balancechanges", id, addr)
	var res []*BalanceChange
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil)
	if err != nil {
		return nil, s, err
	}
	return res, s, nil
}
