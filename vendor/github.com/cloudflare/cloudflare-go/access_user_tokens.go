package cloudflare

import (
	"context"
	"fmt"
	"net/http"
)

type AccessUserEmail struct {
	Email string `json:"email"`
}

// RevokeAccessUserTokens revokes any outstanding tokens issued for a specific user
// Access User.
//
// API reference: https://api.cloudflare.com/#access-organizations-revoke-all-access-tokens-for-a-user
func (api *API) RevokeAccessUserTokens(ctx context.Context, accountID string, accessUserEmail AccessUserEmail) error {
	return api.revokeUserTokens(ctx, accountID, accessUserEmail, AccountRouteRoot)
}

// RevokeZoneLevelAccessUserTokens revokes any outstanding tokens issued for a specific user
// Access User.
//
// API reference: https://api.cloudflare.com/#zone-level-access-organizations-revoke-all-access-tokens-for-a-user
func (api *API) RevokeZoneLevelAccessUserTokens(ctx context.Context, zoneID string, accessUserEmail AccessUserEmail) error {
	return api.revokeUserTokens(ctx, zoneID, accessUserEmail, ZoneRouteRoot)
}

func (api *API) revokeUserTokens(ctx context.Context, id string, accessUserEmail AccessUserEmail, routeRoot RouteRoot) error {
	uri := fmt.Sprintf("/%s/%s/access/organizations/revoke_user", routeRoot, id)

	_, err := api.makeRequestContext(ctx, http.MethodPost, uri, accessUserEmail)
	if err != nil {
		return err
	}

	return nil
}
