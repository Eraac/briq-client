package briq

var BaseURL = "https://www.givebriq.com/v0"

const (
	endpointOrganization = "/organizations/%s"

	endpointTransactions = endpointOrganization + "/transactions"
	endpointTransaction  = endpointTransactions + "/%s"

	endpointUsers = endpointOrganization + "/users"
	endpointUser  = endpointUsers + "/%s"
)
