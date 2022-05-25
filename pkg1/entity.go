package pkg1

type Alert struct {
	RuleName string
	Rules    []AlertItem
}

type AlertItem struct {
	AlertName string
	Severity  string
	Expr      string
	Duration  string
	Summary   string
	Message   string
	Labels    []string
}

// AlertRow for Excel Row Data
type AlertRow struct {
	RuleName  string
	AlertName string
	Severity  string
	Expr      string
	Duration  string
	Summary   string
	Message   string
	Labels    string
}
