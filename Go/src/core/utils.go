package core

// GetStringForBool - return string for bool - to be used in building json
func GetStringForBool(boolValue bool) string {
	if boolValue {
		return "true"
	}
	return "false"
}

// ResourceByID - to create json with with id
type ResourceByID struct {
	ID string `json:"id"`
}

// ResourceByName - to create json with with id
type ResourceByName struct {
	Name string `json:"name"`
}

// Reference - ID, Name and Href
type Reference struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Href string `json:"href,omitempty"`
	CRN  string `json:"crn,omitempty"`
}
