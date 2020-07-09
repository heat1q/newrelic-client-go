// Code generated by typegen; DO NOT EDIT.

package accounts

/* AccountOutline - The `AccountOutline` object provides basic data about an account. */
type AccountOutline struct {
	/* ID -  */
	ID int `json:"id"`
	/* Name -  */
	Name string `json:"name"`
	/* ReportingEventTypes - Returns event types that are currently reporting in the account. */
	ReportingEventTypes []string `json:"reportingEventTypes"`
}

/* RegionScope -  */
type RegionScope string

var RegionScopeTypes = struct {
	/* GLOBAL - Do not filter by region */
	GLOBAL RegionScope
	/* IN_REGION - Filter by region */
	IN_REGION RegionScope
}{
	GLOBAL:    "GLOBAL",
	IN_REGION: "IN_REGION",
}
