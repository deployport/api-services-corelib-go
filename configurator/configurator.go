package configurator

import (
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
	c.LoadFromEnv()
	opts := make(sdk.Options, 0, 2)
	if c.Region != "" {
		opts = append(
			opts,
			sdk.WithOnSubmission(NewOnSubmissionForRegionOnly(
				c.Region,
			)),
		)
	}
	if c.IsComplete() {
		opts = append(
			opts,
			sdk.WithOnSubmission(NewOnSubmissionForCredentials(
				*c.Credentials,
				c.Region,
			)),
		)
	}
	return opts, nil
}

// Config is the configuration for the client
type Config struct {
	Credentials *signingv1.Credentials
	Region      string
}

// IsComplete checks if the configuration is complete
func (c *Config) IsComplete() bool {
	if c.Credentials.IsEmpty() {
		return false
	}
	if c.Region == "" {
		return false
	}
	return true
}

// LoadFromEnv loads the configuration from the environment variables
// DPP_ACCESS_KEY_ID, DPP_SECRET_ACCESS_KEY and DPP_REGION
func (c *Config) LoadFromEnv() {
	if c.Credentials == nil {
		tempCreds := signingv1.Credentials{}
		if keyID := os.Getenv("DPP_ACCESS_KEY_ID"); keyID != "" {
			tempCreds.KeyID = keyID
		}
		if secret := os.Getenv("DPP_SECRET_ACCESS_KEY"); secret != "" {
			tempCreds.Secret = secret
		}
		if !tempCreds.IsEmpty() {
			c.Credentials = &tempCreds
		}
	}
	if c.Region == "" {
		if region := os.Getenv("DPP_REGION"); region != "" {
			c.Region = region
		}
	}
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
