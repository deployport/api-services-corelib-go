package configurator

import (
	"errors"
	"os"

	"go.deployport.com/api-services-corelib/configurator/signingv1"
	sdk "go.deployport.com/specular-runtime/client"
)

// DefaultClientOptions returns the default options for the client
func DefaultClientOptions(options ...sdk.Option) (sdk.Options, error) {
	var c Config
	if err := sdk.ApplyOptions(&c, options...); err != nil {
		return nil, err
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return sdk.Options{
		sdk.WithOnSubmission(NewOnSubmissionForCredentials(
			*c.Credentials,
			c.Region,
		)),
	}, nil
}

// Config is the configuration for the client
type Config struct {
	Credentials *signingv1.Credentials
	Region      string
}

// Validate validates the configuration
func (c *Config) Validate() error {
	// TODO: Load from file
	if c.Credentials == nil {
		c.Credentials = &signingv1.Credentials{}
	}
	if c.Credentials.KeyID == "" {
		if keyID := os.Getenv("DPP_ACCESS_KEY_ID"); keyID != "" {
			c.Credentials.KeyID = keyID
		}
	}
	if c.Credentials.Secret == "" {
		if secret := os.Getenv("DPP_SECRET_ACCESS_KEY"); secret != "" {
			c.Credentials.Secret = secret
		}
	}
	if region := os.Getenv("DPP_REGION"); region != "" && c.Region == "" {
		c.Region = region
	}
	if c.Credentials.KeyID == "" {
		return errors.New("key id is required")
	}
	if c.Credentials.Secret == "" {
		return errors.New("secret is required")
	}
	if c.Region == "" {
		return errors.New("region is required")
	}
	return nil
}

// WithCredentials sets the credentials
func WithCredentials(creds signingv1.Credentials) sdk.OptionFunc {
	return func(o any) error {
		c, ok := o.(*Config)
		if !ok {
			return nil
		}
		c.Credentials = &creds
		return nil
	}
}

// WithRegion sets the region
func WithRegion(region string) sdk.OptionFunc {
	return func(o any) error {
		c, ok := o.(*Config)
		if !ok {
			return nil
		}
		c.Region = region
		return nil
	}
}
