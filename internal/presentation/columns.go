package presentation

type Column struct {
	Name string
	Key  string
}

var Entities = []Column{{"ID", "id"}, {"NAME", "name"}, {"REPORT NAME", "report_name"}}
var Accounts = []Column{{"ID", "id"}, {"NAME", "name"}, {"KIND", "kind"}, {"CURRENCY", "currency"}, {"BALANCE", "balance"}, {"ARCHIVED", "archived"}}
var Categories = []Column{{"ID", "id"}, {"NAME", "name"}, {"KIND", "kind"}, {"PACK", "category_pack_id"}}
var Transactions = []Column{{"ID", "id"}, {"DATE", "date"}, {"KIND", "kind"}, {"AMOUNT", "amount"}, {"CURRENCY", "currency"}, {"CATEGORY", "category_id"}, {"ACCOUNT", "account_id"}, {"DESCRIPTION", "description"}}
var Subscriptions = []Column{{"ID", "id"}, {"NAME", "name"}, {"AMOUNT", "amount"}, {"CURRENCY", "currency"}, {"NEXT", "next_occurrence_date"}, {"ACTIVE", "active"}}
