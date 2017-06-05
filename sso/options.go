package sso

type Options struct {
	// ClientID provided by the EVE Online Developer Portal.
	ClientID string
	// ClientSecret provided by the EVE Online Developer Portal.
	ClientSecret string
	// CallbackAddress is the address the user is redirected back to after logging
	// in with EVE Online. The address must match exactly with the value set in
	// the EVE Online Developer Portal, otherwise EVE's SSO service will reject.
	CallbackAddress string
	// OAuthRoot is the root address of the EVE Online login service the client
	// should use. By default, the client uses the Tranquility cluster. If you
	// wish to use Singularity instead, set this value to sso.SingularityOAuth.
	OAuthRoot string
}

// Validate returns an error of any option is invalide, nil otherwise.
func (opts *Options) Validate() (err error) {
	if opts.ClientID == "" {
		return ErrClientID
	}
	if opts.ClientSecret == "" {
		return ErrClientSecret
	}
	if opts.CallbackAddress == "" {
		return ErrCallbackAddress
	}
	if opts.OAuthRoot == "" {
		opts.OAuthRoot = TranquilityOAuth
	}
	if opts.OAuthRoot != TranquilityOAuth && opts.OAuthRoot != SingularityOAuth {
		return ErrBadOAuthAddress
	}
	return
}
