// auth/auth.go

package auth

import (
	"context"
	"log"
	"github.com/ranjanistic/alumnus/config"
	"golang.org/x/oauth2"
	oidc "github.com/coreos/go-oidc"
)

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

func NewAuthenticator() (*Authenticator, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "https://alumnus.us.auth0.com/")
	if err != nil {
		log.Printf("failed to get provider: %v", err)
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     "jRiCVXAKYgxCBWBGMNFnS8RE1TfW4SyR",
		ClientSecret: config.Env.AUTH0SEC,
		RedirectURL:  config.Env.SITE+"/callback",
		Endpoint: 	  provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil
}