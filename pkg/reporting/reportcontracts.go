package reporting

import "github.com/ddeep2007/KopCapital/pkg/stock"

type ReportBalanceSheet struct {
	Symbol string `csv:"symbol" json:"symbol,omitempty"`
	*stock.AVBalanceSheetInfo
}

type ReportEarnings struct {
	Symbol string `csv:"symbol" json:"symbol,omitempty"`
	*stock.AVEarningsInfo
}

type ReportCashFlow struct {
	Symbol string `csv:"symbol" json:"symbol,omitempty"`
	*stock.AVCashFlowInfo
}

type ReportIncomeStatement struct {
	Symbol string `csv:"symbol" json:"symbol,omitempty"`
	*stock.AVIncomeInfo
}