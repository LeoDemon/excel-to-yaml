package pkg1

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
)

var AlertItemFileName = "./templ/rule-files.xlsx"

func ReadRules() ([]AlertRow, error) {
	f, err := excelize.OpenFile(AlertItemFileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("alerts")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var alertRows []AlertRow
	for idx, row := range rows {
		if idx == 0 {
			continue
		}
		var alertRow AlertRow
		alertRow.RuleName = row[0]
		alertRow.AlertName = row[1]
		alertRow.Severity = row[2]
		alertRow.Expr = row[3]
		alertRow.Duration = row[4]
		alertRow.Summary = row[5]
		alertRow.Message = row[6]
		if len(row) > 7 {
			alertRow.Labels = strings.TrimSpace(row[7])
		}

		alertRows = append(alertRows, alertRow)
	}

	return alertRows, nil
}
