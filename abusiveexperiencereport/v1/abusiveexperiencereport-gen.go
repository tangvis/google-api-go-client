// Copyright 2025 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package abusiveexperiencereport provides access to the Abusive Experience Report API.
//
// For product documentation, see: https://developers.google.com/abusive-experience-report/
//
// # Library status
//
// These client libraries are officially supported by Google. However, this
// library is considered complete and is in maintenance mode. This means
// that we will address critical bugs and security issues but will not add
// any new features.
//
// When possible, we recommend using our newer
// [Cloud Client Libraries for Go](https://pkg.go.dev/cloud.google.com/go)
// that are still actively being worked and iterated on.
//
// # Creating a client
//
// Usage example:
//
//	import "google.golang.org/api/abusiveexperiencereport/v1"
//	...
//	ctx := context.Background()
//	abusiveexperiencereportService, err := abusiveexperiencereport.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for
// authentication. For information on how to create and obtain Application
// Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// # Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API
// keys), use [google.golang.org/api/option.WithAPIKey]:
//
//	abusiveexperiencereportService, err := abusiveexperiencereport.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth
// flow, use [google.golang.org/api/option.WithTokenSource]:
//
//	config := &oauth2.Config{...}
//	// ...
//	token, err := config.Exchange(ctx, ...)
//	abusiveexperiencereportService, err := abusiveexperiencereport.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See [google.golang.org/api/option.ClientOption] for details on options.
package abusiveexperiencereport // import "google.golang.org/api/abusiveexperiencereport/v1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	internal "google.golang.org/api/internal"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint
var _ = internal.Version

const apiId = "abusiveexperiencereport:v1"
const apiName = "abusiveexperiencereport"
const apiVersion = "v1"
const basePath = "https://abusiveexperiencereport.googleapis.com/"
const basePathTemplate = "https://abusiveexperiencereport.UNIVERSE_DOMAIN/"
const mtlsBasePath = "https://abusiveexperiencereport.mtls.googleapis.com/"

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultEndpointTemplate(basePathTemplate))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	opts = append(opts, internaloption.EnableNewAuthLibrary())
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Sites = NewSitesService(s)
	s.ViolatingSites = NewViolatingSitesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Sites *SitesService

	ViolatingSites *ViolatingSitesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewSitesService(s *Service) *SitesService {
	rs := &SitesService{s: s}
	return rs
}

type SitesService struct {
	s *Service
}

func NewViolatingSitesService(s *Service) *ViolatingSitesService {
	rs := &ViolatingSitesService{s: s}
	return rs
}

type ViolatingSitesService struct {
	s *Service
}

// SiteSummaryResponse: Response message for GetSiteSummary.
type SiteSummaryResponse struct {
	// AbusiveStatus: The site's Abusive Experience Report status.
	//
	// Possible values:
	//   "UNKNOWN" - Not reviewed.
	//   "PASSING" - Passing.
	//   "FAILING" - Failing.
	AbusiveStatus string `json:"abusiveStatus,omitempty"`
	// EnforcementTime: The time at which enforcement
	// (https://support.google.com/webtools/answer/7538608) against the site began
	// or will begin. Not set when the filter_status is OFF.
	EnforcementTime string `json:"enforcementTime,omitempty"`
	// FilterStatus: The site's enforcement status
	// (https://support.google.com/webtools/answer/7538608).
	//
	// Possible values:
	//   "UNKNOWN" - N/A.
	//   "ON" - Enforcement is on.
	//   "OFF" - Enforcement is off.
	//   "PAUSED" - Enforcement is paused.
	//   "PENDING" - Enforcement is pending.
	FilterStatus string `json:"filterStatus,omitempty"`
	// LastChangeTime: The time at which the site's status last changed.
	LastChangeTime string `json:"lastChangeTime,omitempty"`
	// ReportUrl: A link to the full Abusive Experience Report for the site. Not
	// set in ViolatingSitesResponse. Note that you must complete the Search
	// Console verification process
	// (https://support.google.com/webmasters/answer/9008080) for the site before
	// you can access the full report.
	ReportUrl string `json:"reportUrl,omitempty"`
	// ReviewedSite: The name of the reviewed site, e.g. `google.com`.
	ReviewedSite string `json:"reviewedSite,omitempty"`
	// UnderReview: Whether the site is currently under review.
	UnderReview bool `json:"underReview,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "AbusiveStatus") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "AbusiveStatus") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s SiteSummaryResponse) MarshalJSON() ([]byte, error) {
	type NoMethod SiteSummaryResponse
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// ViolatingSitesResponse: Response message for ListViolatingSites.
type ViolatingSitesResponse struct {
	// ViolatingSites: The list of violating sites.
	ViolatingSites []*SiteSummaryResponse `json:"violatingSites,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "ViolatingSites") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "ViolatingSites") to include in
	// API requests with the JSON null value. By default, fields with empty values
	// are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s ViolatingSitesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ViolatingSitesResponse
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

type SitesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a site's Abusive Experience Report summary.
//
//   - name: The name of the site whose summary to get, e.g.
//     `sites/http%3A%2F%2Fwww.google.com%2F`. Format: `sites/{site}`.
func (r *SitesService) Get(name string) *SitesGetCall {
	c := &SitesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *SitesGetCall) Fields(s ...googleapi.Field) *SitesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets an optional parameter which makes the operation fail if the
// object's ETag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *SitesGetCall) IfNoneMatch(entityTag string) *SitesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *SitesGetCall) Context(ctx context.Context) *SitesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *SitesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SitesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "", c.header_)
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "abusiveexperiencereport.sites.get" call.
// Any non-2xx status code is an error. Response headers are in either
// *SiteSummaryResponse.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified was
// returned.
func (c *SitesGetCall) Do(opts ...googleapi.CallOption) (*SiteSummaryResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &SiteSummaryResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

type ViolatingSitesListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists sites that are failing in the Abusive Experience Report.
func (r *ViolatingSitesService) List() *ViolatingSitesListCall {
	c := &ViolatingSitesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *ViolatingSitesListCall) Fields(s ...googleapi.Field) *ViolatingSitesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets an optional parameter which makes the operation fail if the
// object's ETag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ViolatingSitesListCall) IfNoneMatch(entityTag string) *ViolatingSitesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *ViolatingSitesListCall) Context(ctx context.Context) *ViolatingSitesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *ViolatingSitesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ViolatingSitesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "", c.header_)
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/violatingSites")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "abusiveexperiencereport.violatingSites.list" call.
// Any non-2xx status code is an error. Response headers are in either
// *ViolatingSitesResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified was
// returned.
func (c *ViolatingSitesListCall) Do(opts ...googleapi.CallOption) (*ViolatingSitesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &ViolatingSitesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}
