package shared

type RuleNotifiers struct {
	Emails  []string `json:"emails,omitempty"`
	Webhook *string  `json:"webhook,omitempty"`
}

// Rule
// Configuration to alert to a url or list of email addresses based on the given type / condition.
type Rule struct {
	CompanyID *string       `json:"companyId,omitempty"`
	ID        string        `json:"id"`
	Notifiers RuleNotifiers `json:"notifiers"`
	Type      string        `json:"type"`
}
