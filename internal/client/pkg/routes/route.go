package routes

import (
	"github.com/stevan-sdk/internal/utils"
)

// Route entities define rules to match client requests. Every request matching a given route will be proxied to its associated service.
type Route struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// One or more lists of values indexed by header name that will cause this route to match if present in the request. The `Host` header cannot be used with this attribute: hosts should be specified using the `hosts` attribute. When `headers` contains only one value and that value starts with the special prefix `~*`, the value is interpreted as a regular expression.
	Headers map[string]string `json:"headers,omitempty"`
	// A list of domain names that match this route. Note that the hosts value is case sensitive.
	Hosts []string `json:"hosts,omitempty"`
	// The status code Kong responds with when all properties of a route match except the protocol i.e. if the protocol of the request is `HTTP` instead of `HTTPS`. `Location` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the route is configured to only accept the `https` protocol.
	HttpsRedirectStatusCode *int64  `json:"https_redirect_status_code,omitempty"`
	Id                      *string `json:"id,omitempty"`
	// A list of HTTP methods that match this route.
	Methods []string `json:"methods,omitempty"`
	// The name of the route. Route names must be unique, and they are case sensitive. For example, there can be two different routes named test and Test.
	Name *string `json:"name,omitempty"`
	// Controls how the service path, route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior.
	PathHandling *string `json:"path_handling,omitempty"`
	// A list of paths that match this route.
	Paths []string `json:"paths,omitempty"`
	// When matching a route via one of the `hosts` domain names, use the request `Host` header in the upstream request headers. If set to `false`, the upstream `Host` header will be that of the service's `host`.
	PreserveHost *bool `json:"preserve_host,omitempty"`
	// An array of the protocols this route should allow. See the [route Object](#route-object) section for a list of accepted protocols. When set to only `https`, HTTP requests are answered with an upgrade error. When set to only `http`, HTTPS requests are answered with an error.
	Protocols []string `json:"protocols,omitempty"`
	// A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. When two routes match the path and have the same `regex_priority`, the older one (lowest `created_at`) is used. Note that the priority for non-regex routes is different (longer non-regex routes are matched before shorter ones).
	RegexPriority *int64 `json:"regex_priority,omitempty"`
	// Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding.
	RequestBuffering *bool `json:"request_buffering,omitempty"`
	// Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding.
	ResponseBuffering *bool `json:"response_buffering,omitempty"`
	// The service this route is associated to. This is where the route proxies traffic to.
	Service *Service `json:"service,omitempty"`
	// A list of SNIs that match this route when using stream routing.
	Snis []string `json:"snis,omitempty"`
	// When matching a route via one of the `paths`, strip the matching prefix from the upstream request URL.
	StripPath *bool `json:"strip_path,omitempty"`
	// An optional set of strings associated with the route for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (r *Route) SetCreatedAt(createdAt int64) {
	r.CreatedAt = &createdAt
}

func (r *Route) GetCreatedAt() *int64 {
	if r == nil {
		return nil
	}
	return r.CreatedAt
}

func (r *Route) SetHeaders(headers map[string]string) {
	r.Headers = utils.CloneMap(headers)
}

func (r *Route) GetHeaders() map[string]string {
	if r == nil {
		return nil
	}
	return r.Headers
}

func (r *Route) SetHosts(hosts []string) {
	r.Hosts = hosts
}

func (r *Route) GetHosts() []string {
	if r == nil {
		return nil
	}
	return r.Hosts
}

func (r *Route) SetHttpsRedirectStatusCode(httpsRedirectStatusCode int64) {
	r.HttpsRedirectStatusCode = &httpsRedirectStatusCode
}

func (r *Route) GetHttpsRedirectStatusCode() *int64 {
	if r == nil {
		return nil
	}
	return r.HttpsRedirectStatusCode
}

func (r *Route) SetId(id string) {
	r.Id = &id
}

func (r *Route) GetId() *string {
	if r == nil {
		return nil
	}
	return r.Id
}

func (r *Route) SetMethods(methods []string) {
	r.Methods = methods
}

func (r *Route) GetMethods() []string {
	if r == nil {
		return nil
	}
	return r.Methods
}

func (r *Route) SetName(name string) {
	r.Name = &name
}

func (r *Route) GetName() *string {
	if r == nil {
		return nil
	}
	return r.Name
}

func (r *Route) SetPathHandling(pathHandling string) {
	r.PathHandling = &pathHandling
}

func (r *Route) GetPathHandling() *string {
	if r == nil {
		return nil
	}
	return r.PathHandling
}

func (r *Route) SetPaths(paths []string) {
	r.Paths = paths
}

func (r *Route) GetPaths() []string {
	if r == nil {
		return nil
	}
	return r.Paths
}

func (r *Route) SetPreserveHost(preserveHost bool) {
	r.PreserveHost = &preserveHost
}

func (r *Route) GetPreserveHost() *bool {
	if r == nil {
		return nil
	}
	return r.PreserveHost
}

func (r *Route) SetProtocols(protocols []string) {
	r.Protocols = protocols
}

func (r *Route) GetProtocols() []string {
	if r == nil {
		return nil
	}
	return r.Protocols
}

func (r *Route) SetRegexPriority(regexPriority int64) {
	r.RegexPriority = &regexPriority
}

func (r *Route) GetRegexPriority() *int64 {
	if r == nil {
		return nil
	}
	return r.RegexPriority
}

func (r *Route) SetRequestBuffering(requestBuffering bool) {
	r.RequestBuffering = &requestBuffering
}

func (r *Route) GetRequestBuffering() *bool {
	if r == nil {
		return nil
	}
	return r.RequestBuffering
}

func (r *Route) SetResponseBuffering(responseBuffering bool) {
	r.ResponseBuffering = &responseBuffering
}

func (r *Route) GetResponseBuffering() *bool {
	if r == nil {
		return nil
	}
	return r.ResponseBuffering
}

func (r *Route) SetService(service Service) {
	r.Service = &service
}

func (r *Route) GetService() *Service {
	if r == nil {
		return nil
	}
	return r.Service
}

func (r *Route) SetSnis(snis []string) {
	r.Snis = snis
}

func (r *Route) GetSnis() []string {
	if r == nil {
		return nil
	}
	return r.Snis
}

func (r *Route) SetStripPath(stripPath bool) {
	r.StripPath = &stripPath
}

func (r *Route) GetStripPath() *bool {
	if r == nil {
		return nil
	}
	return r.StripPath
}

func (r *Route) SetTags(tags []string) {
	r.Tags = tags
}

func (r *Route) GetTags() []string {
	if r == nil {
		return nil
	}
	return r.Tags
}

func (r *Route) SetUpdatedAt(updatedAt int64) {
	r.UpdatedAt = &updatedAt
}

func (r *Route) GetUpdatedAt() *int64 {
	if r == nil {
		return nil
	}
	return r.UpdatedAt
}

// The service this route is associated to. This is where the route proxies traffic to.
type Service struct {
	Id *string `json:"id,omitempty"`
}

func (s *Service) SetId(id string) {
	s.Id = &id
}

func (s *Service) GetId() *string {
	if s == nil {
		return nil
	}
	return s.Id
}
