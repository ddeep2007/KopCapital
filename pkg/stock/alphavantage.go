package stock

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	urlPath             = "https://www.alphavantage.co/query?"
	queryParamSeparator = "&"

	functionParam = "function="
	symbolParam   = "symbol="
	apiKeyParam   = "apikey="

	functionNameFundamentals    = "OVERVIEW"
	functionNameIncomeStatement = "INCOME_STATEMENT"
	functionNameEarnings        = "EARNINGS"
	functionNameBalanceSheet    = "BALANCE_SHEET"
	functionNameCashFlow        = "CASH_FLOW"
)

type AlphaVantage struct {
	logger *log.Logger
	apiKey string
	client *http.Client
}

func InitAlphaVantage(logger *log.Logger, apiKey string) *AlphaVantage {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &AlphaVantage{
		logger: logger,
		apiKey: apiKey,
		client: client,
	}
}

func (a *AlphaVantage) GetFundamentals(ticker string) (*AVFundamentals, error) {
	resp, err := a.performHttpRequest(ticker, functionNameFundamentals)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	avFundamentals := &AVFundamentals{}
	err = json.NewDecoder(resp.Body).Decode(avFundamentals)
	if err != nil {
		a.logger.WithFields(log.Fields{
			"url":   a.constructURL(ticker, functionNameFundamentals),
			"error": err.Error(),
		}).Error("failed to decode http response")
	}

	return avFundamentals, err
}

func (a *AlphaVantage) GetIncomeStatementQuarterly(ticker string) ([]*AVIncomeInfo, error) {
	resp, err := a.performHttpRequest(ticker, functionNameIncomeStatement)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	avIncomeInfo := &AVIncomeStatement{}
	err = json.NewDecoder(resp.Body).Decode(avIncomeInfo)
	if err != nil {
		a.logger.WithFields(log.Fields{
			"url":   a.constructURL(ticker, functionNameIncomeStatement),
			"error": err.Error(),
		}).Error("failed to decode http response")
	}

	return avIncomeInfo.QuarterlyReports, err
}

func (a *AlphaVantage) GetIncomeStatementAnnual(ticker string) ([]*AVIncomeInfo, error) {
	resp, err := a.performHttpRequest(ticker, functionNameIncomeStatement)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	avIncomeInfo := &AVIncomeStatement{}
	err = json.NewDecoder(resp.Body).Decode(avIncomeInfo)
	if err != nil {
		a.logger.WithFields(log.Fields{
			"url":   a.constructURL(ticker, functionNameIncomeStatement),
			"error": err.Error(),
		}).Error("failed to decode http response")
	}

	return avIncomeInfo.AnnualReports, err
}

func (a *AlphaVantage) GetEarningsQuarterly(ticker string) ([]*AVEarningsInfo, error) {
	resp, err := a.performHttpRequest(ticker, functionNameEarnings)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	avEarnings := &AVEarnings{}
	err = json.NewDecoder(resp.Body).Decode(avEarnings)
	if err != nil {
		a.logger.WithFields(log.Fields{
			"url":   a.constructURL(ticker, functionNameEarnings),
			"error": err.Error(),
		}).Error("failed to decode http response")
	}

	return avEarnings.QuarterlyEarnings, err
}

func (a *AlphaVantage) GetEarningsAnnual(ticker string) ([]*AVEarningsInfo, error) {
	resp, err := a.performHttpRequest(ticker, functionNameEarnings)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	avEarnings := &AVEarnings{}
	err = json.NewDecoder(resp.Body).Decode(avEarnings)
	if err != nil {
		a.logger.WithFields(log.Fields{
			"url":   a.constructURL(ticker, functionNameEarnings),
			"error": err.Error(),
		}).Error("failed to decode http response")
	}

	return avEarnings.AnnualEarnings, err
}

func (a *AlphaVantage) GetBalanceSheetQuarterly(ticker string) ([]*AVBalanceSheetInfo, error) {
	resp, err := a.performHttpRequest(ticker, functionNameBalanceSheet)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	avBalanceSheet := &AVBalanceSheet{}
	err = json.NewDecoder(resp.Body).Decode(avBalanceSheet)
	if err != nil {
		a.logger.WithFields(log.Fields{
			"url":   a.constructURL(ticker, functionNameBalanceSheet),
			"error": err.Error(),
		}).Error("failed to decode http response")
	}

	return avBalanceSheet.QuarterlyBalanceSheet, err
}

func (a *AlphaVantage) GetBalanceSheetAnnual(ticker string) ([]*AVBalanceSheetInfo, error) {
	resp, err := a.performHttpRequest(ticker, functionNameBalanceSheet)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	avBalanceSheet := &AVBalanceSheet{}
	err = json.NewDecoder(resp.Body).Decode(avBalanceSheet)
	if err != nil {
		a.logger.WithFields(log.Fields{
			"url":   a.constructURL(ticker, functionNameBalanceSheet),
			"error": err.Error(),
		}).Error("failed to decode http response")
	}

	return avBalanceSheet.AnnualBalanceSheet, err
}

func (a *AlphaVantage) GetCashFlowQuarterly(ticker string) ([]*AVCashFlowInfo, error) {
	resp, err := a.performHttpRequest(ticker, functionNameCashFlow)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	avCashFlow := &AVCashFlow{}
	err = json.NewDecoder(resp.Body).Decode(avCashFlow)
	if err != nil {
		a.logger.WithFields(log.Fields{
			"url":   a.constructURL(ticker, functionNameCashFlow),
			"error": err.Error(),
		}).Error("failed to decode http response")
	}

	return avCashFlow.QuarterlyCashFlow, err
}

func (a *AlphaVantage) GetCashFlowAnnual(ticker string) ([]*AVCashFlowInfo, error) {
	resp, err := a.performHttpRequest(ticker, functionNameCashFlow)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	avCashFlow := &AVCashFlow{}
	err = json.NewDecoder(resp.Body).Decode(avCashFlow)
	if err != nil {
		a.logger.WithFields(log.Fields{
			"url":   a.constructURL(ticker, functionNameCashFlow),
			"error": err.Error(),
		}).Error("failed to decode http response")
	}

	return avCashFlow.AnnualCashFlow, err
}

func (a *AlphaVantage) performHttpRequest(ticker, function string) (*http.Response, error) {
	// Alphavantage enforces a limit of 5 API calls per minute
	time.Sleep(time.Second * 15)
	url := a.constructURL(ticker, function)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		a.logger.WithFields(log.Fields{
			"url":   a.constructURL(ticker, function),
			"error": err.Error(),
		}).Error("failed to create http request")

		return nil, err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		a.logger.WithFields(log.Fields{
			"url":   a.constructURL(ticker, function),
			"error": err.Error(),
		}).Error("http request failed due to a client error")

		return nil, err
	}

	if resp.StatusCode != 200 {
		a.logger.WithFields(log.Fields{
			"url":        a.constructURL(ticker, function),
			"statusCode": resp.StatusCode,
			"status":     resp.Status,
		}).Error("http request failed at server")

		return nil, errors.New("http request failed at server")
	}

	return resp, nil
}

func (a *AlphaVantage) constructURL(ticker, function string) string {
	return urlPath + functionParam + function + queryParamSeparator + symbolParam + ticker + queryParamSeparator +
		apiKeyParam + a.apiKey
}
