package miningcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// GetPools returns a list of all available pools.
func (c *Client) GetPools(ctx context.Context) ([]*PoolInfo, int, error) {
	var res struct {
		Pools []*PoolInfo `json:"pools"`
	}
	s, err := c.UnmarshalPools(ctx, &res)
	if err != nil {
		return nil, s, err
	}
	return res.Pools, s, nil
}

func (c *Client) UnmarshalPools(ctx context.Context, res any) (int, error) {
	e := "/api/pools"
	return c.doRequest(ctx, e, http.MethodGet, res, nil)
}

// GetPool returns information about a specific pool.
func (c *Client) GetPool(ctx context.Context, id string) (*PoolInfo, int, error) {
	var res struct {
		Pool PoolInfo `json:"pool"`
	}
	s, err := c.UnmarshalPool(ctx, id, &res)
	if err != nil {
		return nil, s, err
	}
	return &res.Pool, s, nil
}

func (c *Client) UnmarshalPool(ctx context.Context, id string, res any) (int, error) {
	e := "/api/pools/" + id
	return c.doRequest(ctx, e, http.MethodGet, res, nil)
}

// GetPoolBlocks returns a list of blocks found by a pool.
// This endpoint implements pagination using the `page` and `perPage` parameters.
func (c *Client) GetPoolBlocks(ctx context.Context, id string, params ...map[string]string) (*BlocksRes, int, error) {
	var res BlocksRes
	s, err := c.UnmarshalPoolBlocks(ctx, id, &res, params...)
	if err != nil {
		return nil, s, err
	}
	return &res, s, nil
}

func (c *Client) UnmarshalPoolBlocks(ctx context.Context, id string, res any, params ...map[string]string) (int, error) {
	e := fmt.Sprintf("/api/v2/pools/%s/blocks", id)
	return c.doRequest(ctx, e, http.MethodGet, res, nil, params...)
}

// GetPoolPayments returns a list of payments made by a pool.
// This endpoint implements pagination using the `page` and `perPage` parameters.
func (c *Client) GetPoolPayments(ctx context.Context, id string, params ...map[string]string) (*PaymentRes, int, error) {
	var res PaymentRes
	s, err := c.UnmarshalPoolPayments(ctx, id, &res, params...)
	if err != nil {
		return nil, s, err
	}
	return &res, s, nil
}

func (c *Client) UnmarshalPoolPayments(ctx context.Context, id string, res any, params ...map[string]string) (int, error) {
	e := fmt.Sprintf("/api/v2/pools/%s/payments", id)
	return c.doRequest(ctx, e, http.MethodGet, res, nil, params...)
}

// GetMiners returns a list of all miners in a pool.
// This endpoint implements pagination using the `page` and `perPage` parameters.
func (c *Client) GetMiners(ctx context.Context, id string, params ...map[string]string) ([]*MinerPerformanceStats, int, error) {
	var res []*MinerPerformanceStats
	s, err := c.UnmarshalMiners(ctx, id, &res, params...)
	if err != nil {
		return nil, s, err
	}
	return res, s, nil
}

func (c *Client) UnmarshalMiners(ctx context.Context, id string, res any, params ...map[string]string) (int, error) {
	e := fmt.Sprintf("/api/pools/%s/miners", id)
	return c.doRequest(ctx, e, http.MethodGet, res, nil, params...)
}

// GetMiner returns information about a specific miner in a pool.
// This endpoints allows to specify the performance mode using the`perfMode` parameter.
// Possible values are:
// 		"Hour"
// 		"Day"
// 		"Month"
func (c *Client) GetMiner(ctx context.Context, id, addr string, params ...map[string]string) (*MinerStats, int, error) {
	var res MinerStats
	s, err := c.UnmarshalMiner(ctx, id, addr, &res, params...)
	if err != nil {
		return nil, s, err
	}
	return &res, s, nil
}

func (c *Client) UnmarshalMiner(ctx context.Context, id, addr string, res any, params ...map[string]string) (int, error) {
	e := fmt.Sprintf("/api/pools/%s/miners/%s", id, addr)
	return c.doRequest(ctx, e, http.MethodGet, res, nil, params...)
}

// GetMinerPayments returns a list of payments of a miner.
// This endpoint implements pagination using the `page` and `perPage` parameters.
func (c *Client) GetMinerPayments(ctx context.Context, id, addr string, params ...map[string]string) (*PaymentRes, int, error) {
	var res PaymentRes
	s, err := c.UnmarshalMinerPayments(ctx, id, addr, &res, params...)
	if err != nil {
		return nil, s, err
	}
	return &res, s, nil
}

func (c *Client) UnmarshalMinerPayments(ctx context.Context, id, addr string, res any, params ...map[string]string) (int, error) {
	e := fmt.Sprintf("/api/v2/pools/%s/miners/%s/payments", id, addr)
	return c.doRequest(ctx, e, http.MethodGet, res, nil, params...)
}

// GetMinerDailyEarnings returns a list of daily earnings of a miner.
// This endpoint implements pagination using the `page` and `perPage` parameters.
func (c *Client) GetMinerDailyEarnings(ctx context.Context, id, addr string, params ...map[string]string) (*DailyEarningRes, int, error) {
	var res DailyEarningRes
	s, err := c.UnmarshalMinerDailyEarnings(ctx, id, addr, &res, params...)
	if err != nil {
		return nil, s, err
	}
	return &res, s, nil
}

func (c *Client) UnmarshalMinerDailyEarnings(ctx context.Context, id, addr string, res any, params ...map[string]string) (int, error) {
	e := fmt.Sprintf("/api/v2/pools/%s/miners/%s/earnings/daily", id, addr)
	return c.doRequest(ctx, e, http.MethodGet, res, nil, params...)
}

// GetMinerBalanceChanges returns a list of balance changes of a miner.
// This endpoint implements pagination using the `page` and `perPage` parameters.
func (c *Client) GetMinerBalanceChanges(ctx context.Context, id, addr string, params ...map[string]string) (*BalanceChangeRes, int, error) {
	var res BalanceChangeRes
	s, err := c.UnmarshalMinerBalanceChanges(ctx, id, addr, &res, params...)
	if err != nil {
		return nil, s, err
	}
	return &res, s, nil
}

func (c *Client) UnmarshalMinerBalanceChanges(ctx context.Context, id, addr string, res any, params ...map[string]string) (int, error) {
	e := fmt.Sprintf("/api/v2/pools/%s/miners/%s/balancechanges", id, addr)
	return c.doRequest(ctx, e, http.MethodGet, res, nil, params...)
}

// GetMinerPerformance returns a list of performance samples of a miner.
// This endpoints allows to specify the sample range using the`sampleRange` parameter.
// Possible values are:
// 		"Hour"
// 		"Day"
// 		"Month"
func (c *Client) GetMinerPerformance(ctx context.Context, id, addr string, params ...map[string]string) ([]*WorkerStats, int, error) {
	var res []*WorkerStats
	s, err := c.UnmarshalMinerPerformance(ctx, id, addr, &res, params...)
	if err != nil {
		return nil, s, err
	}
	return nil, 0, nil
}

func (c *Client) UnmarshalMinerPerformance(ctx context.Context, id, addr string, res any, params ...map[string]string) (int, error) {
	e := fmt.Sprintf("/api/pools/%s/miners/%s/performance", id, addr)
	return c.doRequest(ctx, e, http.MethodGet, res, nil)
}

// GetMinerSettings returns the current miner settings of a pool.
func (c *Client) GetMinerSettings(ctx context.Context, id, addr string) (*MinerSettings, int, error) {
	var res MinerSettings
	s, err := c.UnmarshalMinerSettings(ctx, id, addr, &res)
	if err != nil {
		return nil, s, err
	}
	return &res, s, nil
}

func (c *Client) UnmarshalMinerSettings(ctx context.Context, id, addr string, res any) (int, error) {
	e := fmt.Sprintf("/api/pools/%s/miners/%s/settings", id, addr)
	return c.doRequest(ctx, e, http.MethodGet, res, nil)
}

// PostMinerSettings updates the miner settings of a pool.
func (c *Client) PostMinerSettings(ctx context.Context, id, addr string, settings *MinerSettingsUpdateReq) (*MinerSettingsUpdateRes, int, error) {
	// TODO: implement
	return nil, 0, errors.New("not implemented yet")
}

// GetPerformance returns a list of performance stats of a pool.
// This endpoint allows to specify the sample range using the `r` parameter and the sample interval using the `i` parameter.
// Possible values for `r` are:
// 		"Hour"
// 		"Day"
// 		"Month"
// Possible values for `i` are:
// 		"Hour"
// 		"Day"
func (c *Client) GetPerformance(ctx context.Context, id string, params ...map[string]string) ([]*PoolPerformance, int, error) {
	var res struct {
		Stats []*PoolPerformance `json:"stats"`
	}
	s, err := c.UnmarshalPoolPerformance(ctx, id, &res, params...)
	if err != nil {
		return nil, s, err
	}
	return res.Stats, s, nil
}

func (c *Client) UnmarshalPoolPerformance(ctx context.Context, id string, res any, params ...map[string]string) (int, error) {
	e := fmt.Sprintf("/api/pools/%s/performance", id)
	return c.doRequest(ctx, e, http.MethodGet, res, nil, params...)
}
