package services

// service entities are abstractions of upstream services. The main attribute of a service is its URL which can be set as a single string or by specifying the `protocol`, `host`, `port` and `path` individually.
type Service struct {
	// Array of `CA Certificate` object UUIDs that are used to build the trust store while verifying upstream server's TLS certificate. If set to `null` when Nginx default is respected. If default CA list in Nginx are not specified and TLS verification is enabled, then handshake with upstream server will always fail (because no CA are trusted).
	CaCertificates []string `json:"ca_certificates,omitempty"`
	// Certificate to be used as client certificate while TLS handshaking to the upstream server.
	ClientCertificate *ClientCertificate `json:"client_certificate,omitempty"`
	// The timeout in milliseconds for establishing a connection to the upstream server.
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the service is active. If set to `false`, the proxy behavior will be as if any routes attached to it do not exist (404). Default: `true`.
	Enabled *bool `json:"enabled,omitempty"`
	// The host of the upstream server. Note that the host value is case sensitive.
	Host *string `json:"host,omitempty"`
	Id   *string `json:"id,omitempty"`
	// The service name.
	Name *string `json:"name,omitempty"`
	// The path to be used in requests to the upstream server.
	Path *string `json:"path,omitempty"`
	// The upstream server port.
	Port *int64 `json:"port,omitempty"`
	// The protocol used to communicate with the upstream.
	Protocol *string `json:"protocol,omitempty"`
	// The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server.
	ReadTimeout *int64 `json:"read_timeout,omitempty"`
	// The number of retries to execute upon failure to proxy.
	Retries *int64 `json:"retries,omitempty"`
	// An optional set of strings associated with the service for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Whether to enable verification of upstream server TLS certificate. If set to `null`, then the Nginx default is respected.
	TlsVerify *bool `json:"tls_verify,omitempty"`
	// Maximum depth of chain while verifying Upstream server's TLS certificate. If set to `null`, then the Nginx default is respected.
	TlsVerifyDepth *int64 `json:"tls_verify_depth,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// Helper field to set `protocol`, `host`, `port` and `path` using a URL. This field is write-only and is not returned in responses.
	Url *string `json:"url,omitempty"`
	// The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server.
	WriteTimeout *int64 `json:"write_timeout,omitempty"`
}

func (s *Service) SetCaCertificates(caCertificates []string) {
	s.CaCertificates = caCertificates
}

func (s *Service) GetCaCertificates() []string {
	if s == nil {
		return nil
	}
	return s.CaCertificates
}

func (s *Service) SetClientCertificate(clientCertificate ClientCertificate) {
	s.ClientCertificate = &clientCertificate
}

func (s *Service) GetClientCertificate() *ClientCertificate {
	if s == nil {
		return nil
	}
	return s.ClientCertificate
}

func (s *Service) SetConnectTimeout(connectTimeout int64) {
	s.ConnectTimeout = &connectTimeout
}

func (s *Service) GetConnectTimeout() *int64 {
	if s == nil {
		return nil
	}
	return s.ConnectTimeout
}

func (s *Service) SetCreatedAt(createdAt int64) {
	s.CreatedAt = &createdAt
}

func (s *Service) GetCreatedAt() *int64 {
	if s == nil {
		return nil
	}
	return s.CreatedAt
}

func (s *Service) SetEnabled(enabled bool) {
	s.Enabled = &enabled
}

func (s *Service) GetEnabled() *bool {
	if s == nil {
		return nil
	}
	return s.Enabled
}

func (s *Service) SetHost(host string) {
	s.Host = &host
}

func (s *Service) GetHost() *string {
	if s == nil {
		return nil
	}
	return s.Host
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

func (s *Service) SetName(name string) {
	s.Name = &name
}

func (s *Service) GetName() *string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *Service) SetPath(path string) {
	s.Path = &path
}

func (s *Service) GetPath() *string {
	if s == nil {
		return nil
	}
	return s.Path
}

func (s *Service) SetPort(port int64) {
	s.Port = &port
}

func (s *Service) GetPort() *int64 {
	if s == nil {
		return nil
	}
	return s.Port
}

func (s *Service) SetProtocol(protocol string) {
	s.Protocol = &protocol
}

func (s *Service) GetProtocol() *string {
	if s == nil {
		return nil
	}
	return s.Protocol
}

func (s *Service) SetReadTimeout(readTimeout int64) {
	s.ReadTimeout = &readTimeout
}

func (s *Service) GetReadTimeout() *int64 {
	if s == nil {
		return nil
	}
	return s.ReadTimeout
}

func (s *Service) SetRetries(retries int64) {
	s.Retries = &retries
}

func (s *Service) GetRetries() *int64 {
	if s == nil {
		return nil
	}
	return s.Retries
}

func (s *Service) SetTags(tags []string) {
	s.Tags = tags
}

func (s *Service) GetTags() []string {
	if s == nil {
		return nil
	}
	return s.Tags
}

func (s *Service) SetTlsVerify(tlsVerify bool) {
	s.TlsVerify = &tlsVerify
}

func (s *Service) GetTlsVerify() *bool {
	if s == nil {
		return nil
	}
	return s.TlsVerify
}

func (s *Service) SetTlsVerifyDepth(tlsVerifyDepth int64) {
	s.TlsVerifyDepth = &tlsVerifyDepth
}

func (s *Service) GetTlsVerifyDepth() *int64 {
	if s == nil {
		return nil
	}
	return s.TlsVerifyDepth
}

func (s *Service) SetUpdatedAt(updatedAt int64) {
	s.UpdatedAt = &updatedAt
}

func (s *Service) GetUpdatedAt() *int64 {
	if s == nil {
		return nil
	}
	return s.UpdatedAt
}

func (s *Service) SetUrl(url string) {
	s.Url = &url
}

func (s *Service) GetUrl() *string {
	if s == nil {
		return nil
	}
	return s.Url
}

func (s *Service) SetWriteTimeout(writeTimeout int64) {
	s.WriteTimeout = &writeTimeout
}

func (s *Service) GetWriteTimeout() *int64 {
	if s == nil {
		return nil
	}
	return s.WriteTimeout
}

// Certificate to be used as client certificate while TLS handshaking to the upstream server.
type ClientCertificate struct {
	Id *string `json:"id,omitempty"`
}

func (c *ClientCertificate) SetId(id string) {
	c.Id = &id
}

func (c *ClientCertificate) GetId() *string {
	if c == nil {
		return nil
	}
	return c.Id
}
