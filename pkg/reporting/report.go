package reporting

import (
	"errors"
	"github.com/ddeep2007/KopCapital/pkg/stock"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	fileNameTimeFormat = "2006-01-02"

	ReportTypeFundamentals           = "Fundamentals"
	ReportTypeBalanceSheetQuarter    = "BalanceSheetQuarter"
	ReportTypeBalanceSheetAnnual     = "BalanceSheetAnnual"
	ReportTypeIncomeStatementQuarter = "IncomeStatementQuarter"
	ReportTypeIncomeStatementAnnual  = "IncomeStatementAnnual"
	ReportTypeEarningsQuarter        = "EarningsQuarter"
	ReportTypeEarningsAnnual         = "EarningsAnnual"
	ReportTypeCashFlowQuarter        = "CashFlowQuarter"
	ReportTypeCashFlowAnnual         = "CashFlowAnnual"
	ReportTypeSaas                   = "Saas"
)

var SaaSTickers = []string{
	"SNOW",
	"DDOG",
	"ESTC",
	"ZM",
	"TWLO",
	"OKTA",
	"ZS",
	"NET",
}

type Generator struct {
	Logger *log.Logger
	Av     *stock.AlphaVantage
}

func (r *Generator) GenerateTickerReport(ticker, report string) error {
	switch report {
	case ReportTypeFundamentals:
		resp, err := r.Av.GetFundamentals(ticker)
		if err != nil {
			r.Logger.Errorf("failed to get fundamentals %s", err.Error())
			return err
		}

		fundamentalTickers := []*stock.AVFundamentals{resp}
		err = ConvertFundamentalsToCsv(constructFileName(ticker, report), fundamentalTickers)
		if err != nil {
			r.Logger.Errorf("failed to convert to csv %s", err.Error())
			return err
		}

	case ReportTypeBalanceSheetQuarter:
		resp, err := r.Av.GetBalanceSheetQuarterly(ticker)
		if err != nil {
			r.Logger.Errorf("failed to get balance sheet quarter %s", err.Error())
			return err
		}

		reportBS := make([]ReportBalanceSheet, 0)
		bs := ReportBalanceSheet{Symbol: ticker}
		bs.AVBalanceSheetInfo = resp[0]
		reportBS = append(reportBS, bs)
		err = ConvertBalanceSheetToCsv(constructFileName(ticker, report), reportBS)
		if err != nil {
			r.Logger.Errorf("failed to convert to csv %s", err.Error())
			return err
		}

	case ReportTypeBalanceSheetAnnual:
		resp, err := r.Av.GetBalanceSheetAnnual(ticker)
		if err != nil {
			r.Logger.Errorf("failed to get balance sheet quarter %s", err.Error())
			return err
		}

		reports := make([]ReportBalanceSheet, 0)
		rep := ReportBalanceSheet{Symbol: ticker}
		rep.AVBalanceSheetInfo = resp[0]
		reports = append(reports, rep)
		err = ConvertBalanceSheetToCsv(constructFileName(ticker, report), reports)
		if err != nil {
			r.Logger.Errorf("failed to convert to csv %s", err.Error())
			return err
		}

	case ReportTypeIncomeStatementQuarter:
		resp, err := r.Av.GetIncomeStatementQuarterly(ticker)
		if err != nil {
			r.Logger.Errorf("failed to get balance sheet quarter %s", err.Error())
			return err
		}

		reports := make([]ReportIncomeStatement, 0)
		rep := ReportIncomeStatement{Symbol: ticker}
		rep.AVIncomeInfo = resp[0]
		reports = append(reports, rep)

		err = ConvertIncomeStatementToCsv(constructFileName(ticker, report), reports)
		if err != nil {
			r.Logger.Errorf("failed to convert to csv %s", err.Error())
			return err
		}

	case ReportTypeIncomeStatementAnnual:
		resp, err := r.Av.GetIncomeStatementAnnual(ticker)
		if err != nil {
			r.Logger.Errorf("failed to get balance sheet quarter %s", err.Error())
			return err
		}

		reports := make([]ReportIncomeStatement, 0)
		rep := ReportIncomeStatement{Symbol: ticker}
		rep.AVIncomeInfo = resp[0]
		reports = append(reports, rep)

		err = ConvertIncomeStatementToCsv(constructFileName(ticker, report), reports)
		if err != nil {
			r.Logger.Errorf("failed to convert to csv %s", err.Error())
			return err
		}

	case ReportTypeEarningsQuarter:
		resp, err := r.Av.GetEarningsQuarterly(ticker)
		if err != nil {
			r.Logger.Errorf("failed to get balance sheet quarter %s", err.Error())
			return err
		}

		reports := make([]ReportEarnings, 0)
		rep := ReportEarnings{Symbol: ticker}
		rep.AVEarningsInfo = resp[0]
		reports = append(reports, rep)

		err = ConvertEarningsToCsv(constructFileName(ticker, report), reports)
		if err != nil {
			r.Logger.Errorf("failed to convert to csv %s", err.Error())
			return err
		}

	case ReportTypeEarningsAnnual:
		resp, err := r.Av.GetEarningsAnnual(ticker)
		if err != nil {
			r.Logger.Errorf("failed to get balance sheet quarter %s", err.Error())
			return err
		}

		reports := make([]ReportEarnings, 0)
		rep := ReportEarnings{Symbol: ticker}
		rep.AVEarningsInfo = resp[0]
		reports = append(reports, rep)

		err = ConvertEarningsToCsv(constructFileName(ticker, report), reports)
		if err != nil {
			r.Logger.Errorf("failed to convert to csv %s", err.Error())
			return err
		}

	case ReportTypeCashFlowQuarter:
		resp, err := r.Av.GetCashFlowQuarterly(ticker)
		if err != nil {
			r.Logger.Errorf("failed to get balance sheet quarter %s", err.Error())
			return err
		}

		reports := make([]ReportCashFlow, 0)
		rep := ReportCashFlow{Symbol: ticker}
		rep.AVCashFlowInfo = resp[0]
		reports = append(reports, rep)

		err = ConvertCashFlowToCsv(constructFileName(ticker, report), reports)
		if err != nil {
			r.Logger.Errorf("failed to convert to csv %s", err.Error())
			return err
		}

	case ReportTypeCashFlowAnnual:
		resp, err := r.Av.GetCashFlowQuarterly(ticker)
		if err != nil {
			r.Logger.Errorf("failed to get balance sheet quarter %s", err.Error())
			return err
		}

		reports := make([]ReportCashFlow, 0)
		rep := ReportCashFlow{Symbol: ticker}
		rep.AVCashFlowInfo = resp[0]
		reports = append(reports, rep)

		err = ConvertCashFlowToCsv(constructFileName(ticker, report), reports)
		if err != nil {
			r.Logger.Errorf("failed to convert to csv %s", err.Error())
			return err
		}

	case ReportTypeSaas:
		fundamentals := make([]*stock.AVFundamentals, 0)
		bsQuarterly := make([]ReportBalanceSheet, 0)
		bsAnnual := make([]ReportBalanceSheet, 0)
		isQuarterly := make([]ReportIncomeStatement, 0)
		isAnnual := make([]ReportIncomeStatement, 0)
		cfQuarterly := make([]ReportCashFlow, 0)
		cfAnnual := make([]ReportCashFlow, 0)
		earningsAnnual := make([]ReportEarnings, 0)
		earningsQuarterly := make([]ReportEarnings, 0)

		for _, ticker := range SaaSTickers {
			respF, err := r.Av.GetFundamentals(ticker)
			if err != nil {
				r.Logger.Errorf("failed to get fundamentals for %s ticker due to error %s", ticker, err.Error())
				continue
			}

			fundamentals = append(fundamentals, respF)

			respBSA, err := r.Av.GetBalanceSheetAnnual(ticker)
			if err != nil {
				r.Logger.Errorf("failed to get balance sheet annual for %s ticker due to error %s", ticker, err.Error())
				continue
			}

			for _, bsa := range respBSA {
				r := ReportBalanceSheet{Symbol: ticker}
				r.AVBalanceSheetInfo = bsa
				bsAnnual = append(bsAnnual, r)
			}

			respBSQ, err := r.Av.GetBalanceSheetQuarterly(ticker)
			if err != nil {
				r.Logger.Errorf("failed to get balance sheet quarterly for %s ticker due to error %s", ticker, err.Error())
				continue
			}

			for _, bsq := range respBSQ {
				r := ReportBalanceSheet{Symbol: ticker}
				r.AVBalanceSheetInfo = bsq
				bsQuarterly = append(bsQuarterly, r)
			}

			respISA, err := r.Av.GetIncomeStatementAnnual(ticker)
			if err != nil {
				r.Logger.Errorf("failed to get income statement annual for %s ticker due to error %s", ticker, err.Error())
				continue
			}

			for _, isa := range respISA {
				r := ReportIncomeStatement{Symbol: ticker}
				r.AVIncomeInfo = isa
				isAnnual = append(isAnnual, r)
			}

			respISQ, err := r.Av.GetIncomeStatementQuarterly(ticker)
			if err != nil {
				r.Logger.Errorf("failed to get income statement quarterly for %s ticker due to error %s", ticker, err.Error())
				continue
			}

			for _, isq := range respISQ {
				r := ReportIncomeStatement{Symbol: ticker}
				r.AVIncomeInfo = isq
				isQuarterly = append(isQuarterly, r)
			}

			respCFA, err := r.Av.GetCashFlowAnnual(ticker)
			if err != nil {
				r.Logger.Errorf("failed to get cash flow annual for %s ticker due to error %s", ticker, err.Error())
				continue
			}

			for _, cfa := range respCFA {
				r := ReportCashFlow{Symbol: ticker}
				r.AVCashFlowInfo = cfa
				cfAnnual = append(cfAnnual, r)
			}

			respCFQ, err := r.Av.GetCashFlowQuarterly(ticker)
			if err != nil {
				r.Logger.Errorf("failed to get cash flow quarterly for %s ticker due to error %s", ticker, err.Error())
				continue
			}

			for _, cfq := range respCFQ {
				r := ReportCashFlow{Symbol: ticker}
				r.AVCashFlowInfo = cfq
				cfQuarterly = append(cfQuarterly, r)
			}

			respEA, err := r.Av.GetEarningsAnnual(ticker)
			if err != nil {
				r.Logger.Errorf("failed to get earnings annual for %s ticker due to error %s", ticker, err.Error())
				continue
			}

			for _, ea := range respEA {
				r := ReportEarnings{Symbol: ticker}
				r.AVEarningsInfo = ea
				earningsAnnual = append(earningsAnnual, r)
			}

			respEQ, err := r.Av.GetEarningsQuarterly(ticker)
			if err != nil {
				r.Logger.Errorf("failed to get earnings quarterly for %s ticker due to error %s", ticker, err.Error())
				continue
			}

			for _, eq := range respEQ {
				r := ReportEarnings{Symbol: ticker}
				r.AVEarningsInfo = eq
				earningsQuarterly = append(earningsQuarterly, r)
			}
		}

		err := ConvertFundamentalsToCsv(constructFileName("Saas", ReportTypeFundamentals), fundamentals)
		if err != nil {
			r.Logger.Errorf("failed to generated fundamentals csv error %s", err.Error())
		}

		err = ConvertCashFlowToCsv(constructFileName("Saas", ReportTypeCashFlowAnnual), cfAnnual)
		if err != nil {
			r.Logger.Errorf("failed to generated cf annual csv error %s", err.Error())
		}

		err = ConvertCashFlowToCsv(constructFileName("Saas", ReportTypeCashFlowQuarter), cfQuarterly)
		if err != nil {
			r.Logger.Errorf("failed to generated cf quarter csv error %s", err.Error())
		}

		err = ConvertIncomeStatementToCsv(constructFileName("Saas", ReportTypeIncomeStatementAnnual), isAnnual)
		if err != nil {
			r.Logger.Errorf("failed to generated is annual csv error %s", err.Error())
		}

		err = ConvertIncomeStatementToCsv(constructFileName("Saas", ReportTypeIncomeStatementQuarter), isQuarterly)
		if err != nil {
			r.Logger.Errorf("failed to generated is quarterly csv error %s", err.Error())
		}

		err = ConvertBalanceSheetToCsv(constructFileName("Saas", ReportTypeBalanceSheetAnnual), bsAnnual)
		if err != nil {
			r.Logger.Errorf("failed to generated bs annual csv error %s", err.Error())
		}

		err = ConvertBalanceSheetToCsv(constructFileName("Saas", ReportTypeBalanceSheetQuarter), bsQuarterly)
		if err != nil {
			r.Logger.Errorf("failed to generated bs quarterly csv error %s", err.Error())
		}

		err = ConvertEarningsToCsv(constructFileName("Saas", ReportTypeEarningsAnnual), earningsAnnual)
		if err != nil {
			r.Logger.Errorf("failed to generated earnings annual csv error %s", err.Error())
		}

		err = ConvertEarningsToCsv(constructFileName("Saas", ReportTypeEarningsQuarter), earningsQuarterly)
		if err != nil {
			r.Logger.Errorf("failed to generated earnings quarterly csv error %s", err.Error())
		}

	default:
		r.Logger.Errorf("unknown report type specified %s for ticker %S", report, ticker)
		return errors.New("unknown function specified")
	}
	return nil
}

func constructFileName(ticker, functionName string) string {
	return ticker + functionName + time.Now().Format(fileNameTimeFormat) + ".csv"
}
