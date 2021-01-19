package provider

import (
	"context"

	"github.com/netlify/gotrue/conf"
	"golang.org/x/oauth2"
)

const (
	defaultTwitterAuthBase  = "twitter.com"
	defaultTwitterTokenBase = "api.twitter.com"
	defaultTwitterAPIBase   = "api.twitter.com"
)

type twitterProvider struct {
	*oauth2.Config
	ProfileURL string
}

type twitterUser struct {
	Email string `json:"email"`
}

// NewTwitterProvider creates a Facebook account provider.
func NewTwitterProvider(ext conf.OAuthProviderConfiguration) (OAuthProvider, error) {
	authHost := chooseHost(ext.URL, defaultTwitterAuthBase)
	tokenHost := chooseHost(ext.URL, defaultTwitterTokenBase)
	profileURL := chooseHost(ext.URL, defaultTwitterAPIBase) + "/me?fields=email,first_name,last_name,name,picture" // TODO: fix this url

	return &twitterProvider{
		Config: &oauth2.Config{
			ClientID:     ext.ClientID,
			ClientSecret: ext.Secret,
			RedirectURL:  ext.RedirectURI,
			Endpoint: oauth2.Endpoint{
				AuthURL:  authHost + "/oauth/request_token",
				TokenURL: tokenHost + "/oauth/access_token",
			},
			Scopes: []string{"email"},
		},
		ProfileURL: profileURL,
	}, nil
}

func (p twitterProvider) GetOAuthToken(code string) (*oauth2.Token, error) {
	return p.Exchange(oauth2.NoContext, code)
}

func (p twitterProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	// TODO: make external call for getting user info
	// TODO: fill struct with correct user info
	return &UserProvidedData{}, nil
}
