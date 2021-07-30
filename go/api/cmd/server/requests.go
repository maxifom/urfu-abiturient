package main

type Range struct {
	LTE *int64 `json:"lte,omitempty"`
	GTE *int64 `json:"gte,omitempty"`
}

type QueryRequest struct {
	Basis          []string `json:"basis,omitempty"`
	Form           []string `json:"form,omitempty"`
	Sum            Range    `json:"sum,omitempty"`
	Program        string   `json:"program,omitempty"`
	StatementGiven bool     `json:"statement_given"`
	OriginalGiven  bool     `json:"original_given"`
	Type           []string `json:"type,omitempty"`
	Status         []string `json:"status,omitempty"`
	Name           string   `json:"name,omitempty"`
	Number         int64    `json:"number,omitempty"`
}
