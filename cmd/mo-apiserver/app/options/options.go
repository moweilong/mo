package options

import (
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	cliflag "k8s.io/component-base/cli/flag"

	"github.com/moweilong/mo/internal/apiserver"
	"github.com/moweilong/mo/pkg/app"
	genericoptions "github.com/moweilong/mo/pkg/options"
)

type ServerOptions struct {
	HTTPOptions *genericoptions.HTTPOptions `json:"http" mapstructure:"http"`
}

var _ app.NamedFlagSetOptions = (*ServerOptions)(nil)

// Flags implements app.NamedFlagSetOptions.
func (o *ServerOptions) Flags() (fss cliflag.NamedFlagSets) {
	o.HTTPOptions.AddFlags(fss.FlagSet("http"))
	return fss
}

// Complete implements app.NamedFlagSetOptions.
func (s *ServerOptions) Complete() error {
	return s.HTTPOptions.Complete()
}

// Validate implements app.NamedFlagSetOptions.
func (s *ServerOptions) Validate() error {
	errs := []error{}
	errs = append(errs, s.HTTPOptions.Validate()...)
	return utilerrors.NewAggregate(errs)
}

// NewServerOptions creates and returns a new ServerOptions object with default parameters.
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		HTTPOptions: genericoptions.NewHTTPOptions(),
	}
}

// Config builds an apiserver.Config based on ServerOptions.
func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		HTTPOptions: o.HTTPOptions,
	}, nil
}
