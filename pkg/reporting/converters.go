package reporting

import (
	"github.com/ddeep2007/KopCapital/pkg/stock"
	"github.com/gocarina/gocsv"
	"os"
)

func ConvertFundamentalsToCsv(fileName string, input []*stock.AVFundamentals) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()
	err = gocsv.MarshalFile(&input, file)
	if err != nil {
		return err
	}

	return nil
}

func ConvertBalanceSheetToCsv(fileName string, input []ReportBalanceSheet) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()
	err = gocsv.MarshalFile(&input, file)
	if err != nil {
		return err
	}

	return nil
}

func ConvertIncomeStatementToCsv(fileName string, input []ReportIncomeStatement) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()
	err = gocsv.MarshalFile(&input, file)
	if err != nil {
		return err
	}

	return nil
}

func ConvertCashFlowToCsv(fileName string, input []ReportCashFlow) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()
	err = gocsv.MarshalFile(&input, file)
	if err != nil {
		return err
	}

	return nil
}

func ConvertEarningsToCsv(fileName string, input []ReportEarnings) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()
	err = gocsv.MarshalFile(&input, file)
	if err != nil {
		return err
	}

	return nil
}
