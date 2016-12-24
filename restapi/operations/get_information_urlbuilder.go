package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
)

// GetInformationURL generates an URL for the get information operation
type GetInformationURL struct {
}

// Build a url path and query string
func (o *GetInformationURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/blockchain/info"

	result.Path = _path

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetInformationURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetInformationURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetInformationURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetInformationURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetInformationURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetInformationURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
