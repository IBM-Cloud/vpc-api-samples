package core

// Iam_token - variable holds iam token
var Iam_token string

// VPC_api_endpoint - api endpoint prod
// endpoint for different regions would be
// https://{{region}}.iaas.cloud.ibm.com/v1" where region is the region you are referring to.
const VPC_api_endpoint = "https://us-south.iaas.cloud.ibm.com/v1"

// API_key - api key
const API_key = "Your API key here"

// IAM_endpoint - prod
const IAM_endpoint = "https://iam.cloud.ibm.com/identity/token"

// API_version - api version
const API_version = "2019-09-24"

// Generation - 1 is VPC on Classic and 2 is VPC
const Generation = "1"

// QueryParams -
const QueryParams = `?version=` + API_version + `&generation=` + Generation
