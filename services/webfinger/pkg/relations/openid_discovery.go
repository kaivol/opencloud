package relations

import (
	"context"

	"github.com/opencloud-eu/opencloud/services/webfinger/pkg/service/v0"
	"github.com/opencloud-eu/opencloud/services/webfinger/pkg/webfinger"
)

const (
	OpenIDConnectRel = "http://openid.net/specs/connect/1.0/issuer"
)

// ClientIDProperty is the property URI for the OIDC client ID
const ClientIDProperty = "http://openid.net/specs/connect/1.0/client_id"

// ScopeProperty is the property URI for the OIDC scope
const ScopeProperty = "http://openid.net/specs/connect/1.0/scope"

type openIDDiscovery struct {
	Href     string
	ClientID string
	Scope    string
}

// OpenIDDiscovery adds the Openid Connect issuer relation
func OpenIDDiscovery(href string, clientId string, scope string) service.RelationProvider {
	return &openIDDiscovery{
		Href:     href,
		ClientID: clientId,
		Scope:    scope,
	}
}

func (l *openIDDiscovery) Add(_ context.Context, jrd *webfinger.JSONResourceDescriptor) {
	if jrd == nil {
		jrd = &webfinger.JSONResourceDescriptor{}
	}
	link := webfinger.Link{
		Rel:  OpenIDConnectRel,
		Href: l.Href,
	}
	if l.ClientID != "" {
		link.Properties = map[string]string{
			ClientIDProperty: l.ClientID,
		}
	}
	if l.Scope != "" {
		link.Properties = map[string]string{
			ScopeProperty: l.ClientID,
		}
	}
	jrd.Links = append(jrd.Links, link)
}
