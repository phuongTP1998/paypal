package paypal

import (
	"fmt"
	"time"
)

type Balances struct {
	Currency 	     string 	`json:"currency"`
	Primary 	     bool     	`json:"primary"`
	TotalBalance 	 *Money 	`json:"total_balance"`
	AvailableBalance *Money 	`json:"available_balance"`
	WithheldBalance  *Money 	`json:"withheld_balance"`
}

type BalancesRequest struct {
	AsOfTime         *time.Time
}

type BalancesResponse struct {
	Balances  		 	[]Balances 			`json:"balances"`
	AccountID       	string              `json:"account_id"`
	AsOfTime            JSONTime            `json:"as_of_time"`
	LastRefreshTime 	JSONTime            `json:"last_refresh_time"`
}

// GetBalance - Use this to get PayPal balances in all currency of your account.
// Endpoint: GET /v1/reporting/balances
func (c *Client) GetBalances(req *BalancesRequest) (*BalancesResponse, error) {
	response := &BalancesResponse{}

	r, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.APIBase, "/v1/reporting/balances"), nil)
	if err != nil {
		return nil, err
	}

	q := r.URL.Query()

	if req.AsOfTime != nil {
		q.Add("as_of_time", req.AsOfTime.Format(time.RFC3339))
	}

	r.URL.RawQuery = q.Encode()

	if err = c.SendWithAuth(r, response); err != nil {
		return nil, err
	}

	return response, nil
}
