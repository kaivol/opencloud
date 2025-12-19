package staticroutes

import (
	"encoding/json"
	"net/http"

	"github.com/opencloud-eu/opencloud/services/proxy/pkg/config"
)

// OIDCWellKnownRewrite is a handler that rewrites the /.well-known/openid-configuration endpoint for external IDPs.
func (s *StaticRouteHandler) oidcConfig(oidc config.OIDC) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		issuer := oidc.Issuer
		clientId := oidc.ClientID
		scopes := oidc.Scopes

		response := struct {
			Issuer   string `json:"issuer"`
			ClientId string `json:"client_id"`
			Scopes   string `json:"scopes"`
		}{
			issuer, clientId, scopes,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		encoder := json.NewEncoder(w)
		err := encoder.Encode(response)
		if err != nil {
			s.Logger.Error().
				Err(err).
				Str("handler", "oidc publish config").
				Msg("writing response body failed")
		}
	}
}
