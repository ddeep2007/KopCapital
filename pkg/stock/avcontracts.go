package stock

type AVFundamentals struct {
	Symbol                     string `csv:"Symbol" json:"Symbol,omitempty"`
	AssetType                  string `csv:"AssetType" json:"AssetType,omitempty"`
	Name                       string `csv:"Name" json:"Name,omitempty"`
	Exchange                   string `csv:"Exchange" json:"Exchange,omitempty"`
	Currency                   string `csv:"Currency" json:"Currency,omitempty"`
	Country                    string `csv:"Country" json:"Country,omitempty"`
	Sector                     string `csv:"Sector" json:"Sector,omitempty"`
	Industry                   string `csv:"Industry" json:"Industry,omitempty"`
	FullTimeEmployees          string `csv:"FullTimeEmployees" json:"FullTimeEmployees,omitempty"`
	FiscalYearEnd              string `csv:"FiscalYearEnd" json:"FiscalYearEnd,omitempty"`
	LatestQuarter              string `csv:"LatestQuarter" json:"LatestQuarter,omitempty"`
	MarketCapitalization       string `csv:"MarketCapitalization" json:"MarketCapitalization,omitempty"`
	EBITDA                     string `csv:"EBITDA" json:"EBITDA,omitempty"`
	PERatio                    string `csv:"PERatio" json:"PERatio,omitempty"`
	PEGRatio                   string `csv:"PEGRatio" json:"PEGRatio,omitempty"`
	BookValue                  string `csv:"BookValue" json:"BookValue,omitempty"`
	DividendPerShare           string `csv:"DividendPerShare" json:"DividendPerShare,omitempty"`
	DividendYield              string `csv:"DividendYield" json:"DividendYield,omitempty"`
	EPS                        string `csv:"EPS" json:"EPS,omitempty"`
	RevenuePerShareTTM         string `csv:"RevenuePerShareTTM" json:"RevenuePerShareTTM,omitempty"`
	ProfitMargin               string `csv:"ProfitMargin" json:"ProfitMargin,omitempty"`
	OperatingMarginTTM         string `csv:"OperatingMarginTTM" json:"OperatingMarginTTM,omitempty"`
	ReturnOnAssetsTTM          string `csv:"ReturnOnAssetsTTM" json:"ReturnOnAssetsTTM,omitempty"`
	ReturnOnEquityTTM          string `csv:"ReturnOnEquityTTM" json:"ReturnOnEquityTTM,omitempty"`
	RevenueTTM                 string `csv:"RevenueTTM" json:"RevenueTTM,omitempty"`
	GrossProfitTTM             string `csv:"GrossProfitTTM" json:"GrossProfitTTM,omitempty"`
	DilutedEPSTTM              string `csv:"DilutedEPSTTM" json:"DilutedEPSTTM,omitempty"`
	QuarterlyEarningsGrowthYOY string `csv:"QuarterlyEarningsGrowthYOY" json:"QuarterlyEarningsGrowthYOY,omitempty"`
	QuarterlyRevenueGrowthYOY  string `csv:"QuarterlyRevenueGrowthYOY" json:"QuarterlyRevenueGrowthYOY,omitempty"`
	AnalystTargetPrice         string `csv:"AnalystTargetPrice" json:"AnalystTargetPrice,omitempty"`
	TrailingPE                 string `csv:"TrailingPE" json:"TrailingPE,omitempty"`
	ForwardPE                  string `csv:"ForwardPE" json:"ForwardPE,omitempty"`
	PriceToSalesRatioTTM       string `csv:"PriceToSalesRatioTTM" json:"PriceToSalesRatioTTM,omitempty"`
	PriceToBookRatio           string `csv:"PriceToBookRatio" json:"PriceToBookRatio,omitempty"`
	EVToRevenue                string `csv:"EVToRevenue" json:"EVToRevenue,omitempty"`
	EVToEBITDA                 string `csv:"EVToEBITDA" json:"EVToEBITDA,omitempty"`
	Beta                       string `csv:"Beta" json:"Beta,omitempty"`
	FiftyTwoWeekHigh           string `csv:"52WeekHigh" json:"52WeekHigh,omitempty"`
	FiftyTwoWeekLow            string `csv:"52WeekLow" json:"52WeekLow,omitempty"`
	FiftyDayMovingAverage      string `csv:"50DayMovingAverage" json:"50DayMovingAverage,omitempty"`
	TwoHundredDayMovingAverage string `csv:"200DayMovingAverage" json:"200DayMovingAverage,omitempty"`
	SharesOutstanding          string `csv:"SharesOutstanding" json:"SharesOutstanding,omitempty"`
	SharesFloat                string `csv:"SharesFloat" json:"SharesFloat,omitempty"`
	SharesShort                string `csv:"SharesShort" json:"SharesShort,omitempty"`
	SharesShortPriorMonth      string `csv:"SharesShortPriorMonth" json:"SharesShortPriorMonth,omitempty"`
	ShortRatio                 string `csv:" ShortRatio" json:"ShortRatio,omitempty"`
	ShortPercentOutstanding    string `csv:"ShortPercentOutstanding" json:"ShortPercentOutstanding,omitempty"`
	ShortPercentFloat          string `csv:"ShortPercentFloat" json:"ShortPercentFloat,omitempty"`
	PercentInsiders            string `csv:"PercentInsiders" json:"PercentInsiders,omitempty"`
	PercentInstitutions        string `csv:"PercentInstitutions" json:"PercentInstitutions,omitempty"`
	ForwardAnnualDividendRate  string `csv:"ForwardAnnualDividendRate" json:"ForwardAnnualDividendRate,omitempty"`
	ForwardAnnualDividendYield string `csv:"ForwardAnnualDividendYield" json:"ForwardAnnualDividendYield,omitempty"`
	PayoutRatio                string `csv:"PayoutRatio" json:"PayoutRatio,omitempty"`
	DividendDate               string `csv:"DividendDate" json:"DividendDate,omitempty"`
	ExDividendDate             string `csv:"ExDividendDate" json:"ExDividendDate,omitempty"`
	LastSplitFactor            string `csv:"LastSplitFactor" json:"LastSplitFactor,omitempty"`
	LastSplitDate              string `csv:"LastSplitDate" json:"LastSplitDate,omitempty"`
}

type AVIncomeInfo struct {
	FiscalDateEnding                  string `csv:"fiscalDateEnding" json:"fiscalDateEnding,omitempty"`
	ReportedCurrency                  string `csv:"reportedCurrency" json:"reportedCurrency,omitempty"`
	GrossProfit                       string `csv:"grossProfit" json:"grossProfit,omitempty"`
	TotalRevenue                      string `csv:"totalRevenue" json:"totalRevenue,omitempty"`
	CostOfRevenue                     string `csv:"costOfRevenue" json:"costOfRevenue,omitempty"`
	CostOfGoodsAndServicesSold        string `csv:"costOfGoodsAndServicesSold" json:"costOfGoodsAndServicesSold,omitempty"`
	OperatingIncome                   string `csv:"operatingIncome" json:"operatingIncome,omitempty"`
	SellingGeneralAndAdministrative   string `csv:"sellingGeneralAndAdministrative" json:"sellingGeneralAndAdministrative,omitempty"`
	ResearchAndDevelopment            string `csv:"researchAndDevelopment" json:"researchAndDevelopment,omitempty"`
	OperatingExpenses                 string `csv:"operatingExpenses" json:"operatingExpenses,omitempty"`
	InvestmentIncomeNet               string `csv:"investmentIncomeNet" json:"investmentIncomeNet,omitempty"`
	NetInterestIncome                 string `csv:"netInterestIncome" json:"netInterestIncome,omitempty"`
	InterestIncome                    string `csv:"interestIncome" json:"interestIncome,omitempty"`
	InterestExpense                   string `csv:"interestExpense" json:"interestExpense,omitempty"`
	NonInterestIncome                 string `csv:"nonInterestIncome" json:"nonInterestIncome,omitempty"`
	OtherNonOperatingIncome           string `csv:"otherNonOperatingIncome" json:"otherNonOperatingIncome,omitempty"`
	Depreciation                      string `csv:"depreciation" json:"depreciation,omitempty"`
	DepreciationAndAmortization       string `csv:"depreciationAndAmortization" json:"depreciationAndAmortization,omitempty"`
	IncomeBeforeTax                   string `csv:"incomeBeforeTax" json:"incomeBeforeTax,omitempty"`
	IncomeTaxExpense                  string `csv:"incomeTaxExpense" json:"incomeTaxExpense,omitempty"`
	InterestAndDebtExpense            string `csv:"interestAndDebtExpense" json:"interestAndDebtExpense,omitempty"`
	NetIncomeFromContinuingOperations string `csv:"netIncomeFromContinuingOperations" json:"netIncomeFromContinuingOperations,omitempty"`
	ComprehensiveIncomeNetOfTax       string `csv:"comprehensiveIncomeNetOfTax" json:"comprehensiveIncomeNetOfTax,omitempty"`
	EBIT                              string `csv:"ebit" json:"ebit,omitempty"`
	EBITDA                            string `csv:"ebitda" json:"ebitda,omitempty"`
	NetIncome                         string `csv:"netIncome" json:"netIncome,omitempty"`
}

type AVIncomeStatement struct {
	Symbol           string          `csv:"symbol" json:"symbol,omitempty"`
	AnnualReports    []*AVIncomeInfo `csv:"annualReports" json:"annualReports,omitempty"`
	QuarterlyReports []*AVIncomeInfo `csv:"quarterlyReports" json:"quarterlyReports,omitempty"`
}

type AVEarningsInfo struct {
	FiscalDateEnding   string `csv:"fiscalDateEnding" json:"fiscalDateEnding,omitempty"`
	ReportedDate       string `csv:"reportedDate" json:"reportedDate,omitempty"`
	ReportedEPS        string `csv:"reportedEPS" json:"reportedEPS,omitempty"`
	EstimatedEPS       string `csv:"estimatedEPS" json:"estimatedEPS,omitempty"`
	Surprise           string `csv:"surprise" json:"surprise,omitempty"`
	SurprisePercentage string `csv:"surprisePercentage" json:"surprisePercentage,omitempty"`
}

type AVEarnings struct {
	Symbol            string            `csv:"symbol" json:"symbol,omitempty"`
	AnnualEarnings    []*AVEarningsInfo `csv:"annualEarnings" json:"annualEarnings,omitempty"`
	QuarterlyEarnings []*AVEarningsInfo `csv:"quarterlyEarnings" json:"quarterlyEarnings,omitempty"`
}

type AVBalanceSheetInfo struct {
	FiscalDateEnding                       string `csv:"fiscalDateEnding" json:"fiscalDateEnding,omitempty"`
	ReportedCurrency                       string `csv:"reportedCurrency" json:"reportedCurrency,omitempty"`
	TotalAssets                            string `csv:"totalAssets" json:"totalAssets,omitempty"`
	TotalCurrentAssets                     string `csv:"totalCurrentAssets" json:"totalCurrentAssets,omitempty"`
	CashAndCashEquivalentsAtCarryingValue  string `csv:"cashAndCashEquivalentsAtCarryingValue" json:"cashAndCashEquivalentsAtCarryingValue,omitempty"`
	CashAndShortTermInvestments            string `csv:"cashAndShortTermInvestments" json:"cashAndShortTermInvestments,omitempty"`
	Inventory                              string `csv:"inventory" json:"inventory,omitempty"`
	CurrentNetReceivables                  string `csv:"currentNetReceivables" json:"currentNetReceivables,omitempty"`
	TotalNonCurrentAssets                  string `csv:"totalNonCurrentAssets" json:"totalNonCurrentAssets,omitempty"`
	PropertyPlantEquipment                 string `csv:"propertyPlantEquipment" json:"propertyPlantEquipment,omitempty"`
	AccumulatedDepreciationAmortizationPPE string `csv:"accumulatedDepreciationAmortizationPPE" json:"accumulatedDepreciationAmortizationPPE,omitempty"`
	IntangibleAssets                       string `csv:"intangibleAssets" json:"intangibleAssets,omitempty"`
	IntangibleAssetsExcludingGoodwill      string `csv:"intangibleAssetsExcludingGoodwill" json:"intangibleAssetsExcludingGoodwill,omitempty"`
	Goodwill                               string `csv:"goodwill" json:"goodwill,omitempty"`
	Investments                            string `csv:"investments" json:"investments,omitempty"`
	LongTermInvestments                    string `csv:"longTermInvestments" json:"longTermInvestments,omitempty"`
	ShortTermInvestments                   string `csv:"shortTermInvestments" json:"shortTermInvestments,omitempty"`
	OtherCurrentAssets                     string `csv:"otherCurrentAssets" json:"otherCurrentAssets,omitempty"`
	OtherNonCurrentAssets                  string `csv:"otherNonCurrentAssets" json:"otherNonCurrentAssets,omitempty"`
	TotalLiabilities                       string `csv:"totalLiabilities" json:"totalLiabilities,omitempty"`
	TotalCurrentLiabilities                string `csv:"totalCurrentLiabilities" json:"totalCurrentLiabilities,omitempty"`
	CurrentAccountsPayable                 string `csv:"currentAccountsPayable" json:"currentAccountsPayable,omitempty"`
	DeferredRevenue                        string `csv:"deferredRevenue" json:"deferredRevenue,omitempty"`
	CurrentDebt                            string `csv:"currentDebt" json:"currentDebt,omitempty"`
	ShortTermDebt                          string `csv:"shortTermDebt" json:"shortTermDebt,omitempty"`
	TotalNonCurrentLiabilities             string `csv:"totalNonCurrentLiabilities" json:"totalNonCurrentLiabilities,omitempty"`
	CapitalLeaseObligations                string `csv:"capitalLeaseObligations" json:"capitalLeaseObligations,omitempty"`
	LongTermDebt                           string `csv:"longTermDebt" json:"longTermDebt,omitempty"`
	CurrentLongTermDebt                    string `csv:"currentLongTermDebt" json:"currentLongTermDebt,omitempty"`
	LongTermDebtNoncurrent                 string `csv:"longTermDebtNoncurrent" json:"longTermDebtNoncurrent,omitempty"`
	ShortLongTermDebtTotal                 string `csv:"shortLongTermDebtTotal" json:"shortLongTermDebtTotal,omitempty"`
	OtherCurrentLiabilities                string `csv:"otherCurrentLiabilities" json:"otherCurrentLiabilities,omitempty"`
	OtherNonCurrentLiabilities             string `csv:"otherNonCurrentLiabilities json:"otherNonCurrentLiabilities,omitempty"`
	TotalShareholderEquity                 string `csv:"totalShareholderEquity" json:"totalShareholderEquity,omitempty"`
	TreasuryStocks                         string `csv:"treasuryStocks" json:"treasuryStocks,omitempty"`
	RetainedEarnings                       string `csv:"retainedEarnings" json:"retainedEarnings,omitempty"`
	CommonStock                            string `csv:"commonStock" json:"commonStock,omitempty"`
	CommonStockSharesOutstanding           string `csv:"commonStockSharesOutstanding" json:"commonStockSharesOutstanding,omitempty"`
}

type AVBalanceSheet struct {
	Symbol                string                `csv:"symbol" json:"symbol,omitempty"`
	AnnualBalanceSheet    []*AVBalanceSheetInfo `csv:"annualReports" json:"annualReports,omitempty"`
	QuarterlyBalanceSheet []*AVBalanceSheetInfo `csv:"quarterlyReports" json:"quarterlyReports,omitempty"`
}

type AVCashFlowInfo struct {
	FiscalDateEnding                                          string `csv:"fiscalDateEnding" json:"fiscalDateEnding,omitempty"`
	ReportedCurrency                                          string `csv:"reportedCurrency" json:"reportedCurrency,omitempty"`
	OperatingCashflow                                         string `csv:"operatingCashflow" json:"operatingCashflow,omitempty"`
	PaymentsForOperatingActivities                            string `csv:"paymentsForOperatingActivities" json:"paymentsForOperatingActivities,omitempty"`
	ProceedsFromOperatingActivities                           string `csv:"proceedsFromOperatingActivities" json:"proceedsFromOperatingActivities,omitempty"`
	ChangeInOperatingLiabilities                              string `csv:"changeInOperatingLiabilities" json:"changeInOperatingLiabilities,omitempty"`
	DepreciationDepletionAndAmortization                      string `csv:"depreciationDepletionAndAmortization" json:"depreciationDepletionAndAmortization,omitempty"`
	CapitalExpenditures                                       string `csv:"capitalExpenditures" json:"capitalExpenditures,omitempty"`
	ChangeInReceivables                                       string `csv:"changeInReceivables" json:"changeInReceivables,omitempty"`
	ChangeInInventory                                         string `csv:"changeInInventory" json:"changeInInventory,omitempty"`
	ProfitLoss                                                string `csv:"profitLoss" json:"profitLoss,omitempty"`
	CashflowFromInvestment                                    string `csv:"cashflowFromInvestment" json:"cashflowFromInvestment,omitempty"`
	CashflowFromFinancing                                     string `csv:"cashflowFromFinancing" json:"cashflowFromFinancing,omitempty"`
	ProceedsFromRepaymentsOfShortTermDebt                     string `csv:"proceedsFromRepaymentsOfShortTermDebt" json:"proceedsFromRepaymentsOfShortTermDebt,omitempty"`
	PaymentsForRepurchaseOfCommonStock                        string `csv:"paymentsForRepurchaseOfCommonStock" json:"paymentsForRepurchaseOfCommonStock,omitempty"`
	PaymentsForRepurchaseOfEquity                             string `csv:"paymentsForRepurchaseOfEquity" json:"paymentsForRepurchaseOfEquity,omitempty"`
	PaymentsForRepurchaseOfPreferredStock                     string `csv:"paymentsForRepurchaseOfPreferredStock" json:"paymentsForRepurchaseOfPreferredStock,omitempty"`
	DividendPayout                                            string `csv:"dividendPayout" json:"dividendPayout,omitempty"`
	DividendPayoutCommonStock                                 string `csv:"dividendPayoutCommonStock" json:"dividendPayoutCommonStock,omitempty"`
	DividendPayoutPreferredStock                              string `csv:"dividendPayoutPreferredStock" json:"dividendPayoutPreferredStock,omitempty"`
	ProceedsFromIssuanceOfCommonStock                         string `csv:"proceedsFromIssuanceOfCommonStock" json:"proceedsFromIssuanceOfCommonStock,omitempty"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet string `csv:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet" json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet,omitempty"`
	ProceedsFromIssuanceOfPreferredStock                      string `csv:"proceedsFromIssuanceOfPreferredStock" json:"proceedsFromIssuanceOfPreferredStock,omitempty"`
	ProceedsFromRepurchaseOfEquity                            string `csv:"proceedsFromRepurchaseOfEquity" json:"proceedsFromRepurchaseOfEquity,omitempty"`
	ProceedsFromSaleOfTreasuryStock                           string `csv:"proceedsFromSaleOfTreasuryStock" json:"proceedsFromSaleOfTreasuryStock,omitempty"`
	ChangeInCashAndCashEquivalents                            string `csv:"changeInCashAndCashEquivalents" json:"changeInCashAndCashEquivalents,omitempty"`
	ChangeInExchangeRate                                      string `csv:"changeInExchangeRate" json:"changeInExchangeRate,omitempty"`
	NetIncome                                                 string `csv:"netIncome" json:"netIncome,omitempty"`
}

type AVCashFlow struct {
	Symbol            string            `csv:"symbol" json:"symbol,omitempty"`
	AnnualCashFlow    []*AVCashFlowInfo `csv:"annualReports" json:"annualReports,omitempty"`
	QuarterlyCashFlow []*AVCashFlowInfo `csv:"quarterlyReports" json:"quarterlyReports,omitempty"`
}
