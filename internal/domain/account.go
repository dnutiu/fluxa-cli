package domain

type Account struct {
	ID                      ID      `json:"id"`
	EntityID                ID      `json:"entity_id"`
	Name                    string  `json:"name"`
	Kind                    string  `json:"kind"`
	Currency                string  `json:"currency"`
	OpeningBalance          string  `json:"opening_balance"`
	Balance                 string  `json:"balance"`
	BalanceRON              string  `json:"balance_ron"`
	PortfolioValue          *string `json:"portfolio_value"`
	PortfolioValueUpdatedAt *string `json:"portfolio_value_updated_at"`
	AnnualInterest          string  `json:"annual_interest"`
	InterestTax             string  `json:"interest_tax"`
	AccountGroupID          *ID     `json:"account_group_id"`
	Archived                bool    `json:"archived"`
	ArchivedAt              *string `json:"archived_at"`
	CreatedAt               *string `json:"created_at"`
	UpdatedAt               *string `json:"updated_at"`
}

type AccountFilter struct{ IncludeArchived bool }
