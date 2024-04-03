/*
 * Order Uploader
 *
 * Takes cleaned and validated order data records as input and enters the record into the database tables
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type Payment struct {

	CCType string `json:"CC_type,omitempty"`

	CcNumber string `json:"Cc_number,omitempty"`

	CcExp string `json:"Cc_exp,omitempty"`

	CardSource string `json:"Card_source,omitempty"`

	CheckAccount string `json:"Check_account,omitempty"`

	CheckRouting string `json:"Check_routing,omitempty"`

	CheckSsn string `json:"Check_ssn,omitempty"`

	CardBin string `json:"Card_bin,omitempty"`

	CardBinShort string `json:"Card_bin_short,omitempty"`

	PrepaidMatch bool `json:"Prepaid_match,omitempty"`
}

// AssertPaymentRequired checks if the required fields are not zero-ed
func AssertPaymentRequired(obj Payment) error {
	return nil
}

// AssertPaymentConstraints checks if the values respects the defined constraints
func AssertPaymentConstraints(obj Payment) error {
	return nil
}
