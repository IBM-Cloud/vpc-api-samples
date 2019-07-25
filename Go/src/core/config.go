package core

// IamToken - variable holds iam token
var IamToken string

// RiasVersion - rias api version
// generation - 1 is GC and 2 is NG
const RiasVersion = "?version=2019-06-01&generation=1"

// RiasEndpoint - rias endpoint prod
// endpoint for different regions would be
// https://{{region}}.iaas.cloud.ibm.com/v1" where region is the region you are referring to.
const RiasEndpoint = "https://us-south.iaas.cloud.ibm.com/v1"

// APIKey - api key
const APIKey = "Your API key here"

// IAMEndpoint - prod
const IAMEndpoint = "https://iam.cloud.ibm.com/identity/token"
