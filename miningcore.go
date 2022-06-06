package miningcore

import (
	"context"
	"errors"
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
// This endpoint implements pagination using the `page` and `perPage` parameters.
func (c *Client) GetPoolBlocks(ctx context.Context, id string, params ...map[string]string) (*BlocksRes, int, error) {
	e := fmt.Sprintf("/api/v2/pools/%s/blocks", id)
	var res BlocksRes
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil, params...)
	if err != nil {
		return nil, s, err
	}
	return &res, s, nil
}

// GetPoolPayments returns a list of payments made by a pool.
// This endpoint implements pagination using the `page` and `perPage` parameters.
func (c *Client) GetPoolPayments(ctx context.Context, id string, params ...map[string]string) ([]*Payment, int, error) {
	e := fmt.Sprintf("/api/pools/%s/payments", id)
	var res []*Payment
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil, params...)
	if err != nil {
		return nil, s, err
	}
	return res, s, nil
}

// GetMiners returns a list of all miners in a pool.
// This endpoint implements pagination using the `page` and `perPage` parameters.
func (c *Client) GetMiners(ctx context.Context, id string, params ...map[string]string) ([]*MinerPerformanceStats, int, error) {
	e := fmt.Sprintf("/api/pools/%s/miners", id)
	var res []*MinerPerformanceStats
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil, params...)
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
// This endpoint implements pagination using the `page` and `perPage` parameters.
func (c *Client) GetMinerPayments(ctx context.Context, id, addr string, params ...map[string]string) ([]*Payment, int, error) {
	e := fmt.Sprintf("/api/pools/%s/miners/%s/payments", id, addr)
	var res []*Payment
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil, params...)
	if err != nil {
		return nil, s, err
	}
	return res, s, nil
}

// GetMinerDailyEarnings returns a list of daily earnings of a miner.
// This endpoint implements pagination using the `page` and `perPage` parameters.
func (c *Client) GetMinerDailyEarnings(ctx context.Context, id, addr string, params ...map[string]string) ([]*DailyEarning, int, error) {
	e := fmt.Sprintf("/api/pools/%s/miners/%s/earnings/daily", id, addr)
	var res []*DailyEarning
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil, params...)
	if err != nil {
		return nil, s, err
	}
	return res, s, nil
}

// GetMinerBalanceChanges returns a list of balance changes of a miner.
// This endpoint implements pagination using the `page` and `perPage` parameters.
func (c *Client) GetMinerBalanceChanges(ctx context.Context, id, addr string, params ...map[string]string) ([]*BalanceChange, int, error) {
	e := fmt.Sprintf("/api/pools/%s/miners/%s/balancechanges", id, addr)
	var res []*BalanceChange
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil, params...)
	if err != nil {
		return nil, s, err
	}
	return res, s, nil
}

// GetMinerSettings returns the current miner settings of a pool.
func (c *Client) GetMinerSettings(ctx context.Context, id, addr string) (*MinerSettings, int, error) {
	e := fmt.Sprintf("/api/pools/%s/miners/%s/settings", id, addr)
	var res MinerSettings
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil)
	if err != nil {
		return nil, s, err
	}
	return &res, s, nil
}

// PostMinerSettings updates the miner settings of a pool.
func (c *Client) PostMinerSettings(ctx context.Context, id, addr string, settings *MinerSettingsUpdateReq) (*MinerSettingsUpdateRes, int, error) {
	// TODO: implement
	return nil, 0, errors.New("not implemented yet")
}

// GetPerformance returns a list of performance stats of a pool.
func (c *Client) GetPerformance(ctx context.Context, id string) ([]*PoolPerformance, int, error) {
	e := fmt.Sprintf("/api/pools/%s/performance", id)
	var res []*PoolPerformance
	s, err := c.doRequest(ctx, e, http.MethodGet, &res, nil)
	if err != nil {
		return nil, s, err
	}
	return res, s, nil
}
