package common

type SearchModel struct {
	Page        uint   `query:"page"`
	Limit       uint   `query:"limit"`
	Count       bool   `query:"count"`
	OrderBy     string `query:"order_by"`
	OrderType   string `query:"order_type"`
	Between     string `query:"between"`
	BetweenDate string `query:"between_date"`
}
