package docs

import "nx-go-api/app/account/adapters"

// swagger:route POST /v1/accounts/ v1-account-tag v1-account
// aacoun does some amazing stuff.
// responses:
//   200: v1-post-account-response

// This text will appear as description of your response body.
// swagger:response v1-post-account-response
type accountResponse struct {
	// in:body
	Body adapters.AccountRequest
}

// swagger:parameters v1-account
type accountRequest struct {
	// This text will appear as description of your request body.
	// in:body
	Body adapters.AccountRequest
}
