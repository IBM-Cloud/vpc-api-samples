package core

// IamToken - variable holds iam token
var IamToken string

// RiasEndpoint - rias endpoint prod
// endpoint for different regions would be
// https://{{region}}.iaas.cloud.ibm.com/v1" where region is the region you are referring to.
const RiasEndpoint = "https://us-south.iaas.cloud.ibm.com/v1"

// APIKey - api key
const APIKey = "Your API key here"

// IAMEndpoint - prod
const IAMEndpoint = "https://iam.cloud.ibm.com/identity/token"

// RiasVersion - rias api version
const RiasVersion = "2019-09-24"

// Generation - 1 is VPC on Classic and 2 is VPC
const Generation = "1"

// QueryParams -
const QueryParams = `?version=` + RiasVersion + `&generation=` + Generation
