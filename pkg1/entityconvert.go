package pkg1

import (
	"strings"
)

func ConvertAlertRow(alertRows []AlertRow) ([]Alert, error) {
	var alerts []Alert
	if len(alertRows) <= 0 {
		return alerts, nil
	}

	var lastRule = ""
	var currentRule = ""
	var alertItems []AlertItem
	rownum := len(alertRows)
	for i := 0; i < rownum; i++ {
		row := alertRows[i]
		currentRule = row.RuleName

		if len(lastRule) == 0 {
			// first row
			alertItems = newAlertItems(alertItems)
		} else if strings.Compare(currentRule, lastRule) != 0 {
			// new alert
			alerts = newAlert(lastRule, alertItems, alerts)
			alertItems = newAlertItems(alertItems)
		}
		// same alert
		alertItem := convert(row)
		alertItems = append(alertItems, alertItem)

		lastRule = row.RuleName

		if i == (rownum - 1) {
			// last row
			alerts = newAlert(lastRule, alertItems, alerts)
		}
	}

	return alerts, nil
}

func newAlertItems(alertItems []AlertItem) []AlertItem {
	alertItems = make([]AlertItem, 0, 16)
	return alertItems
}

func newAlert(lastRule string, alertRules []AlertItem, alerts []Alert) []Alert {
	alert := Alert{lastRule, alertRules}
	alerts = append(alerts, alert)
	return alerts
}

func convert(row AlertRow) AlertItem {
	var alertItem AlertItem

	alertItem.AlertName = row.AlertName
	alertItem.Message = row.Message
	alertItem.Duration = row.Duration
	alertItem.Expr = row.Expr
	alertItem.Severity = row.Severity
	if len(row.Labels) > 0 {
		alertItem.Labels = strings.Split(row.Labels, ";")
	}
	//fmt.Printf("labels: %v, len: %d, nil: %t", alertItem.Labels, len(alertItem.Labels), nil == alertItem.Labels)
	alertItem.Summary = row.Summary
	return alertItem
}
