// Code generated for package gentemplates by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/consul_catalog-v1.tmpl
// templates/consul_catalog.tmpl
// templates/docker-v1.tmpl
// templates/docker.tmpl
// templates/ecs-v1.tmpl
// templates/ecs.tmpl
// templates/eureka.tmpl
// templates/kubernetes.tmpl
// templates/kv.tmpl
// templates/marathon-v1.tmpl
// templates/marathon.tmpl
// templates/mesos-v1.tmpl
// templates/mesos.tmpl
// templates/notFound.tmpl
// templates/rancher-v1.tmpl
// templates/rancher.tmpl
package gentemplates

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesConsul_catalogV1Tmpl = []byte(`[backends]
{{range $index, $node := .Nodes }}
  [backends."backend-{{ getBackend $node }}".servers."{{ getBackendName $node $index }}"]
    url = "{{ getAttribute "protocol" $node.Service.Tags "http" }}://{{ getBackendAddress $node }}:{{ $node.Service.Port }}"
    {{ $weight := getAttribute "backend.weight" $node.Service.Tags "0" }}
    {{with $weight }}
      weight = {{ $weight }}
    {{end}}
{{end}}

{{range .Services }}
  {{ $service := .ServiceName }}

  {{ $circuitBreaker := getAttribute "backend.circuitbreaker" .Attributes "" }}
  {{with $circuitBreaker }}
  [backends."backend-{{ $service }}".circuitbreaker]
    expression = "{{ $circuitBreaker }}"
  {{end}}

  [backends."backend-{{ $service }}".loadbalancer]
    method = "{{ getAttribute "backend.loadbalancer" .Attributes "wrr" }}"
    sticky = {{ getSticky .Attributes }}
    {{if hasStickinessLabel .Attributes }}
    [backends."backend-{{ $service }}".loadbalancer.stickiness]
      cookieName = "{{ getStickinessCookieName .Attributes }}"
    {{end}}

  {{if hasMaxconnAttributes .Attributes }}
  [backends."backend-{{ $service }}".maxconn]
    amount = {{ getAttribute "backend.maxconn.amount" .Attributes "" }}
    extractorfunc = "{{ getAttribute "backend.maxconn.extractorfunc" .Attributes "" }}"
  {{end}}

{{end}}

[frontends]
{{range .Services }}
  [frontends."frontend-{{ .ServiceName }}"]
  backend = "backend-{{ .ServiceName }}"
  passHostHeader = {{ getAttribute "frontend.passHostHeader" .Attributes "true" }}
  priority = {{ getAttribute "frontend.priority" .Attributes "0" }}

  {{ $entryPoints := getAttribute "frontend.entrypoints" .Attributes "" }}
  {{with $entryPoints }}
    entrypoints = [{{range getEntryPoints $entryPoints }}
      "{{ . }}",
    {{end}}]
  {{end}}

  basicAuth = [{{range getBasicAuth .Attributes }}
  "{{ . }}",
  {{end}}]

  [frontends."frontend-{{ .ServiceName }}".routes."route-host-{{ .ServiceName }}"]
    rule = "{{ getFrontendRule . }}"
{{end}}
`)

func templatesConsul_catalogV1TmplBytes() ([]byte, error) {
	return _templatesConsul_catalogV1Tmpl, nil
}

func templatesConsul_catalogV1Tmpl() (*asset, error) {
	bytes, err := templatesConsul_catalogV1TmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/consul_catalog-v1.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConsul_catalogTmpl = []byte(`[backends]
{{range $service := .Services}}
  {{ $backendName := getServiceBackendName $service }}

  {{ $circuitBreaker := getCircuitBreaker $service.TraefikLabels }}
  {{if $circuitBreaker }}
  [backends."backend-{{ $backendName }}".circuitBreaker]
    expression = "{{ $circuitBreaker.Expression }}"
  {{end}}

  {{ $responseForwarding := getResponseForwarding $service.TraefikLabels }}
  {{if $responseForwarding }}
  [backends."backend-{{ $backendName }}".responseForwarding]
    flushInterval = "{{ $responseForwarding.FlushInterval }}"
  {{end}}



  {{ $loadBalancer := getLoadBalancer $service.TraefikLabels }}
  {{if $loadBalancer }}
  [backends."backend-{{ $backendName }}".loadBalancer]
    method = "{{ $loadBalancer.Method }}"
    sticky = {{ $loadBalancer.Sticky }}
    {{if $loadBalancer.Stickiness }}
    [backends."backend-{{ $backendName }}".loadBalancer.stickiness]
      cookieName = "{{ $loadBalancer.Stickiness.CookieName }}"
      secure = {{ $loadBalancer.Stickiness.Secure }}
      httpOnly = {{ $loadBalancer.Stickiness.HTTPOnly }}
      sameSite = "{{ $loadBalancer.Stickiness.SameSite }}"
    {{end}}
  {{end}}

  {{ $maxConn := getMaxConn $service.TraefikLabels }}
  {{if $maxConn }}
  [backends."backend-{{ $backendName }}".maxConn]
    extractorFunc = "{{ $maxConn.ExtractorFunc }}"
    amount = {{ $maxConn.Amount }}
  {{end}}

  {{ $healthCheck := getHealthCheck $service.TraefikLabels }}
  {{if $healthCheck }}
  [backends."backend-{{ $backendName }}".healthCheck]
    scheme = "{{ $healthCheck.Scheme }}"
    path = "{{ $healthCheck.Path }}"
    port = {{ $healthCheck.Port }}
    interval = "{{ $healthCheck.Interval }}"
    hostname = "{{ $healthCheck.Hostname }}"
    {{if $healthCheck.Headers }}
    [backends."backend-{{ $backendName }}".healthCheck.headers]
      {{range $k, $v := $healthCheck.Headers }}
      {{$k}} = "{{$v}}"
      {{end}}
    {{end}}
  {{end}}

  {{ $buffering := getBuffering $service.TraefikLabels }}
  {{if $buffering }}
  [backends."backend-{{ $backendName }}".buffering]
    maxRequestBodyBytes = {{ $buffering.MaxRequestBodyBytes }}
    memRequestBodyBytes = {{ $buffering.MemRequestBodyBytes }}
    maxResponseBodyBytes = {{ $buffering.MaxResponseBodyBytes }}
    memResponseBodyBytes = {{ $buffering.MemResponseBodyBytes }}
    retryExpression = "{{ $buffering.RetryExpression }}"
  {{end}}

{{end}}
{{range $index, $node := .Nodes}}
  {{ $server := getServer $node }}
  [backends."backend-{{ getNodeBackendName $node }}".servers."{{ getServerName $node $index }}"]
    url = "{{ $server.URL }}"
    weight = {{ $server.Weight }}

{{end}}

[frontends]
{{range $service := .Services}}

  [frontends."frontend-{{ $service.ServiceName }}"]
    backend = "backend-{{ getServiceBackendName $service }}"
    priority = {{ getPriority $service.TraefikLabels }}
    passHostHeader = {{ getPassHostHeader $service.TraefikLabels }}
    passTLSCert = {{ getPassTLSCert $service.TraefikLabels }}

    entryPoints = [{{range getFrontEndEntryPoints $service.TraefikLabels }}
      "{{.}}",
      {{end}}]

    {{ $tlsClientCert := getPassTLSClientCert $service.TraefikLabels }}
    {{if $tlsClientCert }}
    [frontends."frontend-{{ $service.ServiceName }}".passTLSClientCert]
      pem = {{ $tlsClientCert.PEM }}
      {{ $infos := $tlsClientCert.Infos }}
      {{if $infos }}
      [frontends."frontend-{{ $service.ServiceName }}".passTLSClientCert.infos]
        notAfter = {{ $infos.NotAfter   }}
        notBefore = {{ $infos.NotBefore }}
        sans = {{ $infos.Sans }}
        {{ $subject := $infos.Subject }}
        {{if $subject }}
        [frontends."frontend-{{ $service.ServiceName }}".passTLSClientCert.infos.subject]
          country = {{ $subject.Country }}
          province = {{ $subject.Province }}
          locality = {{ $subject.Locality }}
          organization = {{ $subject.Organization }}
          commonName = {{ $subject.CommonName }}
          serialNumber = {{ $subject.SerialNumber }}
          domainComponent = {{ $subject.DomainComponent }}
        {{end}}
        {{ $issuer := $infos.Subject }}
        {{if $issuer }}
        [frontends."frontend-{{ $service.ServiceName }}".passTLSClientCert.infos.issuer]
          country = {{ $issuer.Country }}
          province = {{ $issuer.Province }}
          locality = {{ $issuer.Locality }}
          organization = {{ $issuer.Organization }}
          commonName = {{ $issuer.CommonName }}
          serialNumber = {{ $issuer.SerialNumber }}
          domainComponent = {{ $issuer.DomainComponent }}
        {{end}}
      {{end}}
    {{end}}

    {{ $auth := getAuth $service.TraefikLabels }}
    {{if $auth }}
    [frontends."frontend-{{ $service.ServiceName }}".auth]
      headerField = "{{ $auth.HeaderField }}"

      {{if $auth.Forward }}
      [frontends."frontend-{{ $service.ServiceName }}".auth.forward]
        address = "{{ $auth.Forward.Address }}"
        trustForwardHeader = {{ $auth.Forward.TrustForwardHeader }}
        {{if $auth.Forward.AuthResponseHeaders }}
        authResponseHeaders = [{{range $auth.Forward.AuthResponseHeaders }}
          "{{.}}",
          {{end}}]
        {{end}}

        {{if $auth.Forward.TLS }}
        [frontends."frontend-{{ $service.ServiceName }}".auth.forward.tls]
          ca = "{{ $auth.Forward.TLS.CA }}"
          caOptional = {{ $auth.Forward.TLS.CAOptional }}
          cert = """{{ $auth.Forward.TLS.Cert }}"""
          key = """{{ $auth.Forward.TLS.Key }}"""
          insecureSkipVerify = {{ $auth.Forward.TLS.InsecureSkipVerify }}
        {{end}}
      {{end}}

      {{if $auth.Basic }}
      [frontends."frontend-{{ $service.ServiceName }}".auth.basic]
        removeHeader = {{ $auth.Basic.RemoveHeader }}
        {{if $auth.Basic.Users }}
        users = [{{range $auth.Basic.Users }}
          "{{.}}",
          {{end}}]
        {{end}}
        usersFile = "{{ $auth.Basic.UsersFile }}"
      {{end}}

      {{if $auth.Digest }}
      [frontends."frontend-{{ $service.ServiceName }}".auth.digest]
        removeHeader = {{ $auth.Digest.RemoveHeader }}
        {{if $auth.Digest.Users }}
        users = [{{range $auth.Digest.Users }}
          "{{.}}",
          {{end}}]
        {{end}}
        usersFile = "{{ $auth.Digest.UsersFile }}"
      {{end}}
    {{end}}

    {{ $whitelist := getWhiteList $service.TraefikLabels }}
    {{if $whitelist }}
    [frontends."frontend-{{ $service.ServiceName }}".whiteList]
      sourceRange = [{{range $whitelist.SourceRange }}
        "{{.}}",
        {{end}}]
      useXForwardedFor = {{ $whitelist.UseXForwardedFor }}
    {{end}}

    {{ $redirect := getRedirect $service.TraefikLabels }}
    {{if $redirect }}
    [frontends."frontend-{{ $service.ServiceName }}".redirect]
      entryPoint = "{{ $redirect.EntryPoint }}"
      regex = "{{ $redirect.Regex }}"
      replacement = "{{ $redirect.Replacement }}"
      permanent = {{ $redirect.Permanent }}
    {{end}}

    {{ $errorPages := getErrorPages $service.TraefikLabels }}
    {{if $errorPages }}
    [frontends."frontend-{{ $service.ServiceName }}".errors]
      {{range $pageName, $page := $errorPages }}
      [frontends."frontend-{{ $service.ServiceName }}".errors."{{ $pageName }}"]
        status = [{{range $page.Status }}
          "{{.}}",
          {{end}}]
        backend = "backend-{{ $page.Backend }}"
        query = "{{ $page.Query }}"
      {{end}}
    {{end}}

    {{ $rateLimit := getRateLimit $service.TraefikLabels }}
    {{if $rateLimit }}
    [frontends."frontend-{{ $service.ServiceName }}".rateLimit]
      extractorFunc = "{{ $rateLimit.ExtractorFunc }}"
      [frontends."frontend-{{ $service.ServiceName }}".rateLimit.rateSet]
        {{ range $limitName, $limit := $rateLimit.RateSet }}
        [frontends."frontend-{{ $service.ServiceName }}".rateLimit.rateSet."{{ $limitName }}"]
          period = "{{ $limit.Period }}"
          average = {{ $limit.Average }}
          burst = {{ $limit.Burst }}
        {{end}}
    {{end}}

    {{ $headers := getHeaders $service.TraefikLabels }}
    {{if $headers }}
    [frontends."frontend-{{ $service.ServiceName }}".headers]
      SSLRedirect = {{ $headers.SSLRedirect }}
      SSLTemporaryRedirect = {{ $headers.SSLTemporaryRedirect }}
      SSLHost = "{{ $headers.SSLHost }}"
      SSLForceHost = {{ $headers.SSLForceHost }}
      STSSeconds = {{ $headers.STSSeconds }}
      STSIncludeSubdomains = {{ $headers.STSIncludeSubdomains }}
      STSPreload = {{ $headers.STSPreload }}
      ForceSTSHeader = {{ $headers.ForceSTSHeader }}
      FrameDeny = {{ $headers.FrameDeny }}
      CustomFrameOptionsValue = "{{ $headers.CustomFrameOptionsValue }}"
      ContentTypeNosniff = {{ $headers.ContentTypeNosniff }}
      BrowserXSSFilter = {{ $headers.BrowserXSSFilter }}
      CustomBrowserXSSValue = "{{ $headers.CustomBrowserXSSValue }}"
      ContentSecurityPolicy = "{{ $headers.ContentSecurityPolicy }}"
      PublicKey = "{{ $headers.PublicKey }}"
      ReferrerPolicy = "{{ $headers.ReferrerPolicy }}"
      IsDevelopment = {{ $headers.IsDevelopment }}

      {{if $headers.AllowedHosts }}
      AllowedHosts = [{{range $headers.AllowedHosts }}
        "{{.}}",
        {{end}}]
      {{end}}

      {{if $headers.HostsProxyHeaders }}
      HostsProxyHeaders = [{{range $headers.HostsProxyHeaders }}
        "{{.}}",
        {{end}}]
      {{end}}

      {{if $headers.CustomRequestHeaders }}
      [frontends."frontend-{{ $service.ServiceName }}".headers.customRequestHeaders]
        {{range $k, $v := $headers.CustomRequestHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}

      {{if $headers.CustomResponseHeaders }}
      [frontends."frontend-{{ $service.ServiceName }}".headers.customResponseHeaders]
        {{range $k, $v := $headers.CustomResponseHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}

      {{if $headers.SSLProxyHeaders }}
      [frontends."frontend-{{ $service.ServiceName }}".headers.SSLProxyHeaders]
        {{range $k, $v := $headers.SSLProxyHeaders}}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}
    {{end}}

    [frontends."frontend-{{ $service.ServiceName }}".routes."route-host-{{ $service.ServiceName }}"]
      rule = "{{ getFrontendRule $service }}"

{{end}}
`)

func templatesConsul_catalogTmplBytes() ([]byte, error) {
	return _templatesConsul_catalogTmpl, nil
}

func templatesConsul_catalogTmpl() (*asset, error) {
	bytes, err := templatesConsul_catalogTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/consul_catalog.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDockerV1Tmpl = []byte(`{{$backendServers := .Servers}}

[backends]
{{range $backendName, $backend := .Backends }}

  {{if hasCircuitBreakerLabel $backend }}
  [backends."backend-{{ $backendName }}".circuitbreaker]
    expression = "{{ getCircuitBreakerExpression $backend }}"
  {{end}}

  {{if hasLoadBalancerLabel $backend }}
  [backends."backend-{{ $backendName }}".loadbalancer]
    method = "{{ getLoadBalancerMethod $backend }}"
    sticky = {{ getSticky $backend }}
    {{if hasStickinessLabel $backend }}
    [backends."backend-{{ $backendName }}".loadbalancer.stickiness]
      cookieName = "{{ getStickinessCookieName $backend }}"
    {{end}}
  {{end}}

  {{if hasMaxConnLabels $backend }}
  [backends."backend-{{ $backendName }}".maxconn]
    amount = {{ getMaxConnAmount $backend }}
    extractorfunc = "{{ getMaxConnExtractorFunc $backend }}"
  {{end}}

  {{ $servers := index $backendServers $backendName }}
  {{range $serverName, $server := $servers }}
  {{if hasServices $server }}
    {{$services := getServiceNames $server }}
    {{range $serviceIndex, $serviceName := $services }}
    [backends."backend-{{ getServiceBackend $server $serviceName }}".servers."service-{{ $serverName }}"]
      url = "{{ getServiceProtocol $server $serviceName }}://{{ getIPAddress $server }}:{{ getServicePort $server $serviceName }}"
      weight = {{ getServiceWeight $server $serviceName }}
    {{end}}
  {{else}}
    [backends."backend-{{ $backendName }}".servers."server-{{$server.Name | replace "/" "" | replace "." "-"}}"]
    url = "{{ getProtocol $server }}://{{ getIPAddress $server }}:{{ getPort $server }}"
    weight = {{ getWeight $server }}
  {{end}}
  {{end}}

{{end}}

[frontends]
{{range $frontend, $containers := .Frontends}}
  {{$container := index $containers 0}}

  {{if hasServices $container }}
  {{ $services := getServiceNames $container }}
  {{range $serviceIndex, $serviceName := $services }}
  [frontends."frontend-{{ getServiceBackend $container $serviceName }}"]
    backend = "backend-{{ getServiceBackend $container $serviceName }}"
    passHostHeader = {{ getServicePassHostHeader $container $serviceName }}
    passTLSCert = {{ getServicePassTLSCert $container $serviceName }}

    {{if getWhitelistSourceRange $container }}
    whitelistSourceRange = [{{range getWhitelistSourceRange $container }}
      "{{.}}",
      {{end}}]
    {{end}}

    priority = {{ getServicePriority $container $serviceName }}

    entryPoints = [{{range getServiceEntryPoints $container $serviceName }}
      "{{.}}",
      {{end}}]

    basicAuth = [{{range getServiceBasicAuth $container $serviceName }}
      "{{.}}",
      {{end}}]

    {{if hasServiceRedirect $container $serviceName }}
      [frontends."frontend-{{ getServiceBackend $container $serviceName }}".redirect]
        entryPoint = "{{ getServiceRedirectEntryPoint $container $serviceName }}"
        regex = "{{ getServiceRedirectRegex $container $serviceName }}"
        replacement = "{{ getServiceRedirectReplacement $container $serviceName }}"
    {{end}}

    [frontends."frontend-{{ getServiceBackend $container $serviceName }}".routes."service-{{ $serviceName | replace "/" "" | replace "." "-" }}"]
      rule = "{{ getServiceFrontendRule $container $serviceName }}"
  {{end}}
  {{else}}
  [frontends."frontend-{{ $frontend }}"]
    backend = "backend-{{ getBackend $container }}"
    passHostHeader = {{ getPassHostHeader $container}}
    passTLSCert = {{ getPassTLSCert $container }}
    priority = {{ getPriority $container }}

    {{if getWhitelistSourceRange $container}}
      whitelistSourceRange = [{{range getWhitelistSourceRange $container}}
        "{{.}}",
      {{end}}]
    {{end}}

    entryPoints = [{{range getEntryPoints $container }}
      "{{.}}",
      {{end}}]

    basicAuth = [{{range getBasicAuth $container }}
      "{{.}}",
      {{end}}]

    {{if hasRedirect $container}}
      [frontends."frontend-{{$frontend}}".redirect]
        entryPoint = "{{getRedirectEntryPoint $container}}"
        regex = "{{getRedirectRegex $container}}"
        replacement = "{{getRedirectReplacement $container}}"
    {{end}}

    {{if hasHeaders $container }}
    [frontends."frontend-{{ $frontend }}".headers]
      {{if hasSSLRedirectHeaders $container}}
      SSLRedirect = {{getSSLRedirectHeaders $container}}
      {{end}}
      {{if hasSSLTemporaryRedirectHeaders $container}}
      SSLTemporaryRedirect = {{getSSLTemporaryRedirectHeaders $container}}
      {{end}}
      {{if hasSSLHostHeaders $container}}
      SSLHost = "{{getSSLHostHeaders $container}}"
      {{end}}
      {{if hasSTSSecondsHeaders $container}}
      STSSeconds = {{getSTSSecondsHeaders $container}}
      {{end}}
      {{if hasSTSIncludeSubdomainsHeaders $container}}
      STSIncludeSubdomains = {{getSTSIncludeSubdomainsHeaders $container}}
      {{end}}
      {{if hasSTSPreloadHeaders $container}}
      STSPreload = {{getSTSPreloadHeaders $container}}
      {{end}}
      {{if hasForceSTSHeaderHeaders $container}}
      ForceSTSHeader = {{getForceSTSHeaderHeaders $container}}
      {{end}}
      {{if hasFrameDenyHeaders $container}}
      FrameDeny = {{getFrameDenyHeaders $container}}
      {{end}}
      {{if hasCustomFrameOptionsValueHeaders $container}}
      CustomFrameOptionsValue = "{{getCustomFrameOptionsValueHeaders $container}}"
      {{end}}
      {{if hasContentTypeNosniffHeaders $container}}
      ContentTypeNosniff = {{getContentTypeNosniffHeaders $container}}
      {{end}}
      {{if hasBrowserXSSFilterHeaders $container}}
      BrowserXSSFilter = {{getBrowserXSSFilterHeaders $container}}
      {{end}}
      {{if hasContentSecurityPolicyHeaders $container}}
      ContentSecurityPolicy = "{{getContentSecurityPolicyHeaders $container}}"
      {{end}}
      {{if hasPublicKeyHeaders $container}}
      PublicKey = "{{getPublicKeyHeaders $container}}"
      {{end}}
      {{if hasReferrerPolicyHeaders $container}}
      ReferrerPolicy = "{{getReferrerPolicyHeaders $container}}"
      {{end}}
      {{if hasIsDevelopmentHeaders $container}}
      IsDevelopment = {{getIsDevelopmentHeaders $container}}
      {{end}}
      {{if hasAllowedHostsHeaders $container}}
      AllowedHosts = [{{range getAllowedHostsHeaders $container}}
        "{{.}}",
        {{end}}]
      {{end}}
      {{if hasHostsProxyHeaders $container}}
      HostsProxyHeaders = [{{range getHostsProxyHeaders $container}}
        "{{.}}",
        {{end}}]
      {{end}}
      {{if hasRequestHeaders $container}}
        [frontends."frontend-{{$frontend}}".headers.customrequestheaders]
        {{range $k, $v := getRequestHeaders $container}}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}
      {{if hasResponseHeaders $container}}
        [frontends."frontend-{{$frontend}}".headers.customresponseheaders]
        {{range $k, $v := getResponseHeaders $container}}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}
      {{if hasSSLProxyHeaders $container}}
        [frontends."frontend-{{$frontend}}".headers.SSLProxyHeaders]
        {{range $k, $v := getSSLProxyHeaders $container}}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}
    {{end}}

    [frontends."frontend-{{$frontend}}".routes."route-frontend-{{$frontend}}"]
      rule = "{{getFrontendRule $container}}"
  {{end}}

{{end}}
`)

func templatesDockerV1TmplBytes() ([]byte, error) {
	return _templatesDockerV1Tmpl, nil
}

func templatesDockerV1Tmpl() (*asset, error) {
	bytes, err := templatesDockerV1TmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/docker-v1.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDockerTmpl = []byte(`{{$backendServers := .Servers}}
[backends]
{{range $backendName, $servers := .Servers}}
{{ $backend := index $servers 0 }}

  {{ $circuitBreaker := getCircuitBreaker $backend.SegmentLabels }}
  {{if $circuitBreaker }}
  [backends."backend-{{ $backendName }}".circuitBreaker]
    expression = "{{ $circuitBreaker.Expression }}"
  {{end}}

  {{ $responseForwarding := getResponseForwarding $backend.SegmentLabels }}
  {{if $responseForwarding }}
  [backends."backend-{{ $backendName }}".responseForwarding]
    flushInterval = "{{ $responseForwarding.FlushInterval }}"
  {{end}}

  {{ $loadBalancer := getLoadBalancer $backend.SegmentLabels }}
  {{if $loadBalancer }}
    [backends."backend-{{ $backendName }}".loadBalancer]
      method = "{{ $loadBalancer.Method }}"
      sticky = {{ $loadBalancer.Sticky }}
      {{if $loadBalancer.Stickiness }}
      [backends."backend-{{ $backendName }}".loadBalancer.stickiness]
        cookieName = "{{ $loadBalancer.Stickiness.CookieName }}"
        secure = {{ $loadBalancer.Stickiness.Secure }}
        httpOnly = {{ $loadBalancer.Stickiness.HTTPOnly }}
        sameSite = "{{ $loadBalancer.Stickiness.SameSite }}"
      {{end}}
  {{end}}

  {{ $maxConn := getMaxConn $backend.SegmentLabels }}
  {{if $maxConn }}
  [backends."backend-{{ $backendName }}".maxConn]
    extractorFunc = "{{ $maxConn.ExtractorFunc }}"
    amount = {{ $maxConn.Amount }}
  {{end}}

  {{ $healthCheck := getHealthCheck $backend.SegmentLabels }}
  {{if $healthCheck }}
  [backends."backend-{{ $backendName }}".healthCheck]
    scheme = "{{ $healthCheck.Scheme }}"
    path = "{{ $healthCheck.Path }}"
    port = {{ $healthCheck.Port }}
    interval = "{{ $healthCheck.Interval }}"
    hostname = "{{ $healthCheck.Hostname }}"
    {{if $healthCheck.Headers }}
    [backends."backend-{{ $backendName }}".healthCheck.headers]
      {{range $k, $v := $healthCheck.Headers }}
      {{$k}} = "{{$v}}"
      {{end}}
    {{end}}
  {{end}}

  {{ $buffering := getBuffering $backend.SegmentLabels }}
  {{if $buffering }}
  [backends."backend-{{ $backendName }}".buffering]
    maxRequestBodyBytes = {{ $buffering.MaxRequestBodyBytes }}
    memRequestBodyBytes = {{ $buffering.MemRequestBodyBytes }}
    maxResponseBodyBytes = {{ $buffering.MaxResponseBodyBytes }}
    memResponseBodyBytes = {{ $buffering.MemResponseBodyBytes }}
    retryExpression = "{{ $buffering.RetryExpression }}"
  {{end}}

  {{range $serverName, $server := getServers $servers }}
  [backends."backend-{{ $backendName }}".servers."{{ $serverName }}"]
    url = "{{ $server.URL }}"
    weight = {{ $server.Weight }}
  {{end}}

{{end}}

[frontends]
{{range $frontendName, $containers := .Frontends }}
  {{ $container := index $containers 0 }}

  [frontends."frontend-{{ $frontendName }}"]
    backend = "backend-{{ getBackendName $container }}"
    priority = {{ getPriority $container.SegmentLabels }}
    passHostHeader = {{ getPassHostHeader $container.SegmentLabels }}
    passTLSCert = {{ getPassTLSCert $container.SegmentLabels }}

    entryPoints = [{{range getEntryPoints $container.SegmentLabels }}
      "{{.}}",
      {{end}}]

    {{ $tlsClientCert := getPassTLSClientCert $container.SegmentLabels }}
    {{if $tlsClientCert }}
    [frontends."frontend-{{ $frontendName }}".passTLSClientCert]
      pem = {{ $tlsClientCert.PEM }}
      {{ $infos := $tlsClientCert.Infos }}
      {{if $infos }}
      [frontends."frontend-{{ $frontendName }}".passTLSClientCert.infos]
        notAfter = {{ $infos.NotAfter   }}
        notBefore = {{ $infos.NotBefore }}
        sans = {{ $infos.Sans }}
        {{ $subject := $infos.Subject }}
        {{if $subject }}
        [frontends."frontend-{{ $frontendName }}".passTLSClientCert.infos.subject]
          country = {{ $subject.Country }}
          province = {{ $subject.Province }}
          locality = {{ $subject.Locality }}
          organization = {{ $subject.Organization }}
          commonName = {{ $subject.CommonName }}
          serialNumber = {{ $subject.SerialNumber }}
          domainComponent = {{ $subject.DomainComponent }}
        {{end}}
        {{ $issuer := $infos.Issuer }}
        {{if $issuer }}
        [frontends."frontend-{{ $frontendName }}".passTLSClientCert.infos.issuer]
          country = {{ $issuer.Country }}
          province = {{ $issuer.Province }}
          locality = {{ $issuer.Locality }}
          organization = {{ $issuer.Organization }}
          commonName = {{ $issuer.CommonName }}
          serialNumber = {{ $issuer.SerialNumber }}
          domainComponent = {{ $issuer.DomainComponent }}
        {{end}}
      {{end}}
    {{end}}

    {{ $auth := getAuth $container.SegmentLabels }}
    {{if $auth }}
    [frontends."frontend-{{ $frontendName }}".auth]
      headerField = "{{ $auth.HeaderField }}"

      {{if $auth.Forward }}
      [frontends."frontend-{{ $frontendName }}".auth.forward]
        address = "{{ $auth.Forward.Address }}"
        trustForwardHeader = {{ $auth.Forward.TrustForwardHeader }}
        {{if $auth.Forward.AuthResponseHeaders }}
        authResponseHeaders = [{{range $auth.Forward.AuthResponseHeaders }}
          "{{.}}",
          {{end}}]
        {{end}}

        {{if $auth.Forward.TLS }}
        [frontends."frontend-{{ $frontendName }}".auth.forward.tls]
          ca = "{{ $auth.Forward.TLS.CA }}"
          caOptional = {{ $auth.Forward.TLS.CAOptional }}
          cert = """{{ $auth.Forward.TLS.Cert }}"""
          key = """{{ $auth.Forward.TLS.Key }}"""
          insecureSkipVerify = {{ $auth.Forward.TLS.InsecureSkipVerify }}
        {{end}}
      {{end}}

      {{if $auth.Basic }}
      [frontends."frontend-{{ $frontendName }}".auth.basic]
        removeHeader = {{ $auth.Basic.RemoveHeader }}
        {{if $auth.Basic.Users }}
        users = [{{range $auth.Basic.Users }}
          "{{.}}",
          {{end}}]
        {{end}}
        usersFile = "{{ $auth.Basic.UsersFile }}"
      {{end}}

      {{if $auth.Digest }}
      [frontends."frontend-{{ $frontendName }}".auth.digest]
        removeHeader = {{ $auth.Digest.RemoveHeader }}
        {{if $auth.Digest.Users }}
        users = [{{range $auth.Digest.Users }}
          "{{.}}",
          {{end}}]
        {{end}}
        usersFile = "{{ $auth.Digest.UsersFile }}"
      {{end}}
    {{end}}

    {{ $whitelist := getWhiteList $container.SegmentLabels }}
    {{if $whitelist }}
    [frontends."frontend-{{ $frontendName }}".whiteList]
      sourceRange = [{{range $whitelist.SourceRange }}
        "{{.}}",
        {{end}}]
      useXForwardedFor = {{ $whitelist.UseXForwardedFor }}
    {{end}}

    {{ $redirect := getRedirect $container.SegmentLabels }}
    {{if $redirect }}
    [frontends."frontend-{{ $frontendName }}".redirect]
      entryPoint = "{{ $redirect.EntryPoint }}"
      regex = "{{ $redirect.Regex }}"
      replacement = "{{ $redirect.Replacement }}"
      permanent = {{ $redirect.Permanent }}
    {{end}}

    {{ $errorPages := getErrorPages $container.SegmentLabels }}
    {{if $errorPages }}
    [frontends."frontend-{{ $frontendName }}".errors]
      {{range $pageName, $page := $errorPages }}
      [frontends."frontend-{{ $frontendName }}".errors."{{ $pageName }}"]
        status = [{{range $page.Status }}
          "{{.}}",
          {{end}}]
        backend = "backend-{{ $page.Backend }}"
        query = "{{ $page.Query }}"
      {{end}}
    {{end}}

    {{ $rateLimit := getRateLimit $container.SegmentLabels }}
    {{if $rateLimit }}
    [frontends."frontend-{{ $frontendName }}".rateLimit]
      extractorFunc = "{{ $rateLimit.ExtractorFunc }}"
      [frontends."frontend-{{ $frontendName }}".rateLimit.rateSet]
        {{ range $limitName, $limit := $rateLimit.RateSet }}
        [frontends."frontend-{{ $frontendName }}".rateLimit.rateSet."{{ $limitName }}"]
          period = "{{ $limit.Period }}"
          average = {{ $limit.Average }}
          burst = {{ $limit.Burst }}
        {{end}}
    {{end}}

    {{ $headers := getHeaders $container.SegmentLabels }}
    {{if $headers }}
    [frontends."frontend-{{ $frontendName }}".headers]
      SSLRedirect = {{ $headers.SSLRedirect }}
      SSLTemporaryRedirect = {{ $headers.SSLTemporaryRedirect }}
      SSLHost = "{{ $headers.SSLHost }}"
      SSLForceHost = {{ $headers.SSLForceHost }}
      STSSeconds = {{ $headers.STSSeconds }}
      STSIncludeSubdomains = {{ $headers.STSIncludeSubdomains }}
      STSPreload = {{ $headers.STSPreload }}
      ForceSTSHeader = {{ $headers.ForceSTSHeader }}
      FrameDeny = {{ $headers.FrameDeny }}
      CustomFrameOptionsValue = "{{ $headers.CustomFrameOptionsValue }}"
      ContentTypeNosniff = {{ $headers.ContentTypeNosniff }}
      BrowserXSSFilter = {{ $headers.BrowserXSSFilter }}
      ContentSecurityPolicy = "{{ $headers.ContentSecurityPolicy }}"
      CustomBrowserXSSValue = "{{ $headers.CustomBrowserXSSValue }}"
      PublicKey = "{{ $headers.PublicKey }}"
      ReferrerPolicy = "{{ $headers.ReferrerPolicy }}"
      IsDevelopment = {{ $headers.IsDevelopment }}

      {{if $headers.AllowedHosts }}
      AllowedHosts = [{{range $headers.AllowedHosts }}
        "{{.}}",
        {{end}}]
      {{end}}

      {{if $headers.HostsProxyHeaders }}
      HostsProxyHeaders = [{{range $headers.HostsProxyHeaders }}
        "{{.}}",
        {{end}}]
      {{end}}

      {{if $headers.CustomRequestHeaders }}
      [frontends."frontend-{{ $frontendName }}".headers.customRequestHeaders]
        {{range $k, $v := $headers.CustomRequestHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}

      {{if $headers.CustomResponseHeaders }}
      [frontends."frontend-{{ $frontendName }}".headers.customResponseHeaders]
        {{range $k, $v := $headers.CustomResponseHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}

      {{if $headers.SSLProxyHeaders }}
      [frontends."frontend-{{ $frontendName }}".headers.SSLProxyHeaders]
        {{range $k, $v := $headers.SSLProxyHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}

    {{end}}

    [frontends."frontend-{{ $frontendName }}".routes."route-frontend-{{ $frontendName }}"]
      rule = "{{ getFrontendRule $container $container.SegmentLabels }}"

{{end}}
`)

func templatesDockerTmplBytes() ([]byte, error) {
	return _templatesDockerTmpl, nil
}

func templatesDockerTmpl() (*asset, error) {
	bytes, err := templatesDockerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/docker.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesEcsV1Tmpl = []byte(`[backends]
{{range $serviceName, $instances := .Services }}
  [backends."backend-{{ $serviceName }}".loadBalancer]
    method = "{{ getLoadBalancerMethod $instances }}"
    sticky = {{ getLoadBalancerSticky $instances }}

    {{if hasStickinessLabel $instances }}
    [backends."backend-{{ $serviceName }}".loadBalancer.stickiness]
      cookieName = "{{ getStickinessCookieName $instances }}"
    {{end}}

    {{ if hasHealthCheckLabels $instances }}
    [backends."backend-{{ $serviceName }}".healthCheck]
      path = "{{ getHealthCheckPath $instances }}"
      interval = "{{ getHealthCheckInterval $instances }}"
    {{end}}

    {{range $index, $i := $instances }}
    [backends."backend-{{ $serviceName }}".servers."server-{{ $i.Name }}{{ $i.ID }}"]
      url = "{{ getProtocol $i }}://{{ getHost $i }}:{{ getPort $i }}"
      weight = {{ getWeight $i }}
    {{end}}
{{end}}

[frontends]
{{range $serviceName, $instances := .Services}}
{{range filterFrontends $instances }}
  [frontends."frontend-{{ $serviceName }}"]
    backend = "backend-{{ $serviceName }}"
    passHostHeader = {{ getPassHostHeader . }}
    priority = {{ getPriority . }}

    entryPoints = [{{range  getEntryPoints . }}
      "{{.}}",
      {{end}}]

    basicAuth = [{{range getBasicAuth . }}
      "{{.}}",
    {{end}}]

    [frontends."frontend-{{ $serviceName }}".routes."route-frontend-{{ $serviceName }}"]
      rule = "{{getFrontendRule .}}"
{{end}}
{{end}}`)

func templatesEcsV1TmplBytes() ([]byte, error) {
	return _templatesEcsV1Tmpl, nil
}

func templatesEcsV1Tmpl() (*asset, error) {
	bytes, err := templatesEcsV1TmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/ecs-v1.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesEcsTmpl = []byte(`[backends]
{{range $serviceName, $instances := .Services }}
  {{ $firstInstance := index $instances 0 }}

  {{ $circuitBreaker := getCircuitBreaker $firstInstance.SegmentLabels }}
  {{if $circuitBreaker }}
  [backends."backend-{{ $serviceName }}".circuitBreaker]
    expression = "{{ $circuitBreaker.Expression }}"
  {{end}}

  {{ $responseForwarding := getResponseForwarding $firstInstance.SegmentLabels }}
  {{if $responseForwarding }}
  [backends."backend-{{ $serviceName }}".responseForwarding]
    flushInterval = "{{ $responseForwarding.FlushInterval }}"
  {{end}}

  {{ $loadBalancer := getLoadBalancer $firstInstance.SegmentLabels }}
  {{if $loadBalancer }}
  [backends."backend-{{ $serviceName }}".loadBalancer]
    method = "{{ $loadBalancer.Method }}"
    sticky = {{ $loadBalancer.Sticky }}
    {{if $loadBalancer.Stickiness }}
    [backends."backend-{{ $serviceName }}".loadBalancer.stickiness]
      cookieName = "{{ $loadBalancer.Stickiness.CookieName }}"
      secure = {{ $loadBalancer.Stickiness.Secure }}
      httpOnly = {{ $loadBalancer.Stickiness.HTTPOnly }}
      sameSite = "{{ $loadBalancer.Stickiness.SameSite }}"
    {{end}}
  {{end}}

  {{ $maxConn := getMaxConn $firstInstance.SegmentLabels }}
  {{if $maxConn }}
  [backends."backend-{{ $serviceName }}".maxConn]
    extractorFunc = "{{ $maxConn.ExtractorFunc }}"
    amount = {{ $maxConn.Amount }}
  {{end}}

  {{ $healthCheck := getHealthCheck $firstInstance.SegmentLabels }}
  {{if $healthCheck }}
  [backends."backend-{{ $serviceName }}".healthCheck]
    scheme = "{{ $healthCheck.Scheme }}"
    path = "{{ $healthCheck.Path }}"
    port = {{ $healthCheck.Port }}
    interval = "{{ $healthCheck.Interval }}"
    hostname = "{{ $healthCheck.Hostname }}"
    {{if $healthCheck.Headers }}
    [backends."backend-{{ $serviceName }}".healthCheck.headers]
      {{range $k, $v := $healthCheck.Headers }}
      {{$k}} = "{{$v}}"
      {{end}}
    {{end}}
  {{end}}

  {{ $buffering := getBuffering $firstInstance.SegmentLabels }}
  {{if $buffering }}
  [backends."backend-{{ $serviceName }}".buffering]
    maxRequestBodyBytes = {{ $buffering.MaxRequestBodyBytes }}
    memRequestBodyBytes = {{ $buffering.MemRequestBodyBytes }}
    maxResponseBodyBytes = {{ $buffering.MaxResponseBodyBytes }}
    memResponseBodyBytes = {{ $buffering.MemResponseBodyBytes }}
    retryExpression = "{{ $buffering.RetryExpression }}"
  {{end}}

  {{range $serverName, $server := getServers $instances }}
  [backends."backend-{{ $serviceName }}".servers."{{ $serverName }}"]
    url = "{{ $server.URL }}"
    weight = {{ $server.Weight }}
  {{end}}

{{end}}

[frontends]
{{range $serviceName, $instances := .Services }}
{{range $instance := filterFrontends $instances }}

  {{ $frontendName := getFrontendName $instance }}

  [frontends."frontend-{{ $frontendName }}"]
    backend = "backend-{{ $serviceName }}"
    priority = {{ getPriority $instance.SegmentLabels }}
    passHostHeader = {{ getPassHostHeader $instance.SegmentLabels }}
    passTLSCert = {{ getPassTLSCert $instance.SegmentLabels }}

    entryPoints = [{{range getEntryPoints $instance.SegmentLabels }}
      "{{.}}",
      {{end}}]

    {{ $tlsClientCert := getPassTLSClientCert $instance.SegmentLabels }}
    {{if $tlsClientCert }}
    [frontends."frontend-{{ $frontendName }}".passTLSClientCert]
      pem = {{ $tlsClientCert.PEM }}
      {{ $infos := $tlsClientCert.Infos }}
      {{if $infos }}
      [frontends."frontend-{{ $frontendName }}".passTLSClientCert.infos]
        notAfter = {{ $infos.NotAfter   }}
        notBefore = {{ $infos.NotBefore }}
        sans = {{ $infos.Sans }}
        {{ $subject := $infos.Subject }}
        {{if $subject }}
        [frontends."frontend-{{ $frontendName }}".passTLSClientCert.infos.subject]
          country = {{ $subject.Country }}
          province = {{ $subject.Province }}
          locality = {{ $subject.Locality }}
          organization = {{ $subject.Organization }}
          commonName = {{ $subject.CommonName }}
          serialNumber = {{ $subject.SerialNumber }}
          domainComponent = {{ $subject.DomainComponent }}
        {{end}}
        {{ $issuer := $infos.Issuer }}
        {{if $issuer }}
        [frontends."frontend-{{ $frontendName }}".passTLSClientCert.infos.issuer]
          country = {{ $issuer.Country }}
          province = {{ $issuer.Province }}
          locality = {{ $issuer.Locality }}
          organization = {{ $issuer.Organization }}
          commonName = {{ $issuer.CommonName }}
          serialNumber = {{ $issuer.SerialNumber }}
          domainComponent = {{ $issuer.DomainComponent }}
        {{end}}
      {{end}}
    {{end}}

    {{ $auth := getAuth $instance.SegmentLabels }}
    {{if $auth }}
    [frontends."frontend-{{ $frontendName }}".auth]
      headerField = "{{ $auth.HeaderField }}"

      {{if $auth.Forward }}
      [frontends."frontend-{{ $frontendName }}".auth.forward]
        address = "{{ $auth.Forward.Address }}"
        trustForwardHeader = {{ $auth.Forward.TrustForwardHeader }}
        {{if $auth.Forward.AuthResponseHeaders }}
        authResponseHeaders = [{{range $auth.Forward.AuthResponseHeaders }}
          "{{.}}",
          {{end}}]
        {{end}}

        {{if $auth.Forward.TLS }}
        [frontends."frontend-{{ $frontendName }}".auth.forward.tls]
          ca = "{{ $auth.Forward.TLS.CA }}"
          caOptional = {{ $auth.Forward.TLS.CAOptional }}
          cert = """{{ $auth.Forward.TLS.Cert }}"""
          key = """{{ $auth.Forward.TLS.Key }}"""
          insecureSkipVerify = {{ $auth.Forward.TLS.InsecureSkipVerify }}
        {{end}}
      {{end}}

      {{if $auth.Basic }}
      [frontends."frontend-{{ $frontendName }}".auth.basic]
        removeHeader = {{ $auth.Basic.RemoveHeader }}
        {{if $auth.Basic.Users }}
        users = [{{range $auth.Basic.Users }}
          "{{.}}",
          {{end}}]
        {{end}}
        usersFile = "{{ $auth.Basic.UsersFile }}"
      {{end}}

      {{if $auth.Digest }}
      [frontends."frontend-{{ $frontendName }}".auth.digest]
        removeHeader = {{ $auth.Digest.RemoveHeader }}
        {{if $auth.Digest.Users }}
        users = [{{range $auth.Digest.Users }}
         "{{.}}",
          {{end}}]
        {{end}}
        usersFile = "{{ $auth.Digest.UsersFile }}"
      {{end}}
    {{end}}

    {{ $whitelist := getWhiteList $instance.SegmentLabels }}
    {{if $whitelist }}
    [frontends."frontend-{{ $frontendName }}".whiteList]
      sourceRange = [{{range $whitelist.SourceRange }}
        "{{.}}",
        {{end}}]
      useXForwardedFor = {{ $whitelist.UseXForwardedFor }}
    {{end}}

    {{ $redirect := getRedirect $instance.SegmentLabels }}
    {{if $redirect }}
    [frontends."frontend-{{ $frontendName }}".redirect]
      entryPoint = "{{ $redirect.EntryPoint }}"
      regex = "{{ $redirect.Regex }}"
      replacement = "{{ $redirect.Replacement }}"
      permanent = {{ $redirect.Permanent }}
    {{end}}

    {{ $errorPages := getErrorPages $instance.SegmentLabels }}
    {{if $errorPages }}
    [frontends."frontend-{{ $frontendName }}".errors]
      {{range $pageName, $page := $errorPages }}
      [frontends."frontend-{{ $frontendName }}".errors."{{ $pageName }}"]
        status = [{{range $page.Status }}
          "{{.}}",
          {{end}}]
        backend = "backend-{{ $page.Backend }}"
        query = "{{ $page.Query }}"
      {{end}}
    {{end}}

    {{ $rateLimit := getRateLimit $instance.SegmentLabels }}
    {{if $rateLimit }}
    [frontends."frontend-{{ $frontendName }}".rateLimit]
      extractorFunc = "{{ $rateLimit.ExtractorFunc }}"
      [frontends."frontend-{{ $frontendName }}".rateLimit.rateSet]
        {{ range $limitName, $limit := $rateLimit.RateSet }}
        [frontends."frontend-{{ $frontendName }}".rateLimit.rateSet."{{ $limitName }}"]
          period = "{{ $limit.Period }}"
          average = {{ $limit.Average }}
          burst = {{ $limit.Burst }}
        {{end}}
    {{end}}

    {{ $headers := getHeaders $instance.SegmentLabels }}
    {{if $headers }}
    [frontends."frontend-{{ $frontendName }}".headers]
      SSLRedirect = {{ $headers.SSLRedirect }}
      SSLTemporaryRedirect = {{ $headers.SSLTemporaryRedirect }}
      SSLHost = "{{ $headers.SSLHost }}"
      SSLForceHost = {{ $headers.SSLForceHost }}
      STSSeconds = {{ $headers.STSSeconds }}
      STSIncludeSubdomains = {{ $headers.STSIncludeSubdomains }}
      STSPreload = {{ $headers.STSPreload }}
      ForceSTSHeader = {{ $headers.ForceSTSHeader }}
      FrameDeny = {{ $headers.FrameDeny }}
      CustomFrameOptionsValue = "{{ $headers.CustomFrameOptionsValue }}"
      ContentTypeNosniff = {{ $headers.ContentTypeNosniff }}
      BrowserXSSFilter = {{ $headers.BrowserXSSFilter }}
      CustomBrowserXSSValue = "{{ $headers.CustomBrowserXSSValue }}"
      ContentSecurityPolicy = "{{ $headers.ContentSecurityPolicy }}"
      PublicKey = "{{ $headers.PublicKey }}"
      ReferrerPolicy = "{{ $headers.ReferrerPolicy }}"
      IsDevelopment = {{ $headers.IsDevelopment }}

      {{if $headers.AllowedHosts }}
      AllowedHosts = [{{range $headers.AllowedHosts }}
        "{{.}}",
        {{end}}]
      {{end}}

      {{if $headers.HostsProxyHeaders }}
      HostsProxyHeaders = [{{range $headers.HostsProxyHeaders }}
        "{{.}}",
        {{end}}]
      {{end}}

      {{if $headers.CustomRequestHeaders }}
      [frontends."frontend-{{ $frontendName }}".headers.customRequestHeaders]
        {{range $k, $v := $headers.CustomRequestHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}

      {{if $headers.CustomResponseHeaders }}
      [frontends."frontend-{{ $frontendName }}".headers.customResponseHeaders]
        {{range $k, $v := $headers.CustomResponseHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}

      {{if $headers.SSLProxyHeaders }}
      [frontends."frontend-{{ $frontendName }}".headers.SSLProxyHeaders]
        {{range $k, $v := $headers.SSLProxyHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}
    {{end}}

    [frontends."frontend-{{ $frontendName }}".routes."route-frontend-{{ $frontendName }}"]
      rule = "{{ getFrontendRule $instance }}"

{{end}}
{{end}}`)

func templatesEcsTmplBytes() ([]byte, error) {
	return _templatesEcsTmpl, nil
}

func templatesEcsTmpl() (*asset, error) {
	bytes, err := templatesEcsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/ecs.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesEurekaTmpl = []byte(`[backends]
{{range $app := .Applications }}

  [backends.backend-{{ $app.Name }}]

    {{range $instance := .Instances }}
    [backends."backend-{{ $app.Name }}".servers."server-{{ getInstanceID $instance }}"]
      url = "{{ getProtocol $instance }}://{{ .IpAddr }}:{{ getPort $instance }}"
      weight = {{ getWeight $instance }}
    {{end}}

{{end}}

[frontends]
{{range $app := .Applications }}

  [frontends."frontend-{{ $app.Name }}"]
    backend = "backend-{{ $app.Name }}"
    entryPoints = ["http"]

    [frontends."frontend-{{ $app.Name }}".routes."route-host{{ $app.Name }}"]
      rule = "Host:{{ $app.Name | tolower }}"

{{end}}
`)

func templatesEurekaTmplBytes() ([]byte, error) {
	return _templatesEurekaTmpl, nil
}

func templatesEurekaTmpl() (*asset, error) {
	bytes, err := templatesEurekaTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/eureka.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesKubernetesTmpl = []byte(`[backends]
{{range $backendName, $backend := .Backends }}

  [backends."{{ $backendName }}"]

    {{if $backend.CircuitBreaker }}
    [backends."{{ $backendName }}".circuitBreaker]
      expression = "{{ $backend.CircuitBreaker.Expression }}"
    {{end}}

    {{if $backend.ResponseForwarding }}
    [backends."{{ $backendName }}".responseForwarding]
      flushInterval = "{{ $backend.ResponseForwarding.FlushInterval }}"
    {{end}}

    [backends."{{ $backendName }}".loadBalancer]
      method = "{{ $backend.LoadBalancer.Method }}"
      sticky = {{ $backend.LoadBalancer.Sticky }}
      {{if $backend.LoadBalancer.Stickiness }}
      [backends."{{ $backendName }}".loadBalancer.stickiness]
        cookieName = "{{ $backend.LoadBalancer.Stickiness.CookieName }}"
        secure = {{ $backend.LoadBalancer.Stickiness.Secure }}
        httpOnly = {{ $backend.LoadBalancer.Stickiness.HTTPOnly }}
        sameSite = "{{ $backend.LoadBalancer.Stickiness.SameSite }}"
      {{end}}

    {{if $backend.MaxConn }}
    [backends."{{ $backendName }}".maxConn]
      amount = {{ $backend.MaxConn.Amount }}
      extractorFunc = "{{ $backend.MaxConn.ExtractorFunc }}"
    {{end}}

    {{if $backend.Buffering }}
    [backends."{{ $backendName }}".buffering]
      maxRequestBodyBytes = {{ $backend.Buffering.MaxRequestBodyBytes }}
      memRequestBodyBytes = {{ $backend.Buffering.MemRequestBodyBytes }}
      maxResponseBodyBytes = {{ $backend.Buffering.MaxResponseBodyBytes }}
      memResponseBodyBytes = {{ $backend.Buffering.MemResponseBodyBytes }}
      retryExpression = "{{ $backend.Buffering.RetryExpression }}"
    {{end}}

    {{range $serverName, $server := $backend.Servers }}
    [backends."{{ $backendName }}".servers."{{ $serverName }}"]
      url = "{{ $server.URL }}"
      weight = {{ $server.Weight }}
    {{end}}

{{end}}

[frontends]
{{range $frontendName, $frontend := .Frontends }}

  [frontends."{{ $frontendName }}"]
    backend = "{{ $frontend.Backend }}"
    priority = {{ $frontend.Priority }}
    passHostHeader = {{ $frontend.PassHostHeader }}
    passTLSCert = {{ $frontend.PassTLSCert }}

    entryPoints = [{{range $frontend.EntryPoints }}
      "{{.}}",
      {{end}}]

    {{if $frontend.Auth }}
    [frontends."{{ $frontendName }}".auth]
      {{if $frontend.Auth.HeaderField }}
      headerField = "{{ $frontend.Auth.HeaderField }}"
      {{end}}

      {{if $frontend.Auth.Basic }}
      [frontends."{{ $frontendName }}".auth.basic]
        removeHeader = {{$frontend.Auth.Basic.RemoveHeader}}
        users = [{{range $frontend.Auth.Basic.Users }}
          "{{.}}",
          {{end}}]
      {{end}}

      {{if $frontend.Auth.Digest }}
      [frontends."{{ $frontendName }}".auth.digest]
        removeHeader = {{$frontend.Auth.Digest.RemoveHeader}}
        users = [{{range $frontend.Auth.Digest.Users }}
          "{{.}}",
          {{end}}]
      {{end}}

      {{if $frontend.Auth.Forward }}
        [frontends."{{ $frontendName }}".auth.forward]
          address = "{{ $frontend.Auth.Forward.Address }}"
          authResponseHeaders = [{{range $frontend.Auth.Forward.AuthResponseHeaders }}
            "{{.}}",
            {{end}}]
          trustForwardHeader = {{ $frontend.Auth.Forward.TrustForwardHeader }}
          {{if $frontend.Auth.Forward.TLS }}
          [frontends."{{ $frontendName }}".auth.forward.tls]
            cert = """{{ $frontend.Auth.Forward.TLS.Cert }}"""
            key = """{{ $frontend.Auth.Forward.TLS.Key }}"""
            insecureSkipVerify = {{ $frontend.Auth.Forward.TLS.InsecureSkipVerify }}
          {{end}}
      {{end}}

    {{end}}

    {{if $frontend.WhiteList }}
    [frontends."{{ $frontendName }}".whiteList]
      sourceRange = [{{range $frontend.WhiteList.SourceRange }}
        "{{.}}",
        {{end}}]
      useXForwardedFor = {{ $frontend.WhiteList.UseXForwardedFor }}
    {{end}}

    {{if $frontend.Redirect }}
    [frontends."{{ $frontendName }}".redirect]
      entryPoint = "{{ $frontend.Redirect.EntryPoint }}"
      regex = "{{ $frontend.Redirect.Regex }}"
      replacement = "{{ $frontend.Redirect.Replacement }}"
      permanent = {{ $frontend.Redirect.Permanent }}
    {{end}}

    {{if $frontend.Errors }}
    [frontends."{{ $frontendName }}".errors]
      {{range $pageName, $page := $frontend.Errors }}
      [frontends."{{ $frontendName }}".errors."{{ $pageName }}"]
        status = [{{range $page.Status }}
          "{{.}}",
          {{end}}]
        backend = "{{ $page.Backend }}"
        query = "{{ $page.Query }}"
      {{end}}
    {{end}}

    {{if $frontend.RateLimit }}
    [frontends."{{ $frontendName }}".rateLimit]
      extractorFunc = "{{ $frontend.RateLimit.ExtractorFunc }}"
      [frontends."{{ $frontendName }}".rateLimit.rateSet]
        {{range $limitName, $limit := $frontend.RateLimit.RateSet }}
        [frontends."{{ $frontendName }}".rateLimit.rateSet."{{ $limitName }}"]
          period = "{{ $limit.Period }}"
          average = {{ $limit.Average }}
          burst = {{ $limit.Burst }}
        {{end}}
    {{end}}

    {{if $frontend.PassTLSClientCert }}
    [frontends."{{ $frontendName }}".passTLSClientCert]
      pem = {{ $frontend.PassTLSClientCert.PEM }}
      {{ $infos := $frontend.PassTLSClientCert.Infos }}
      {{if $infos }}
      [frontends."{{ $frontendName }}".passTLSClientCert.infos]
        notAfter = {{ $infos.NotAfter   }}
        notBefore = {{ $infos.NotBefore }}
        sans = {{ $infos.Sans }}
        {{ $subject := $infos.Subject }}
        {{if $subject }}
        [frontends."{{ $frontendName }}".passTLSClientCert.infos.subject]
          country = {{ $subject.Country }}
          province = {{ $subject.Province }}
          locality = {{ $subject.Locality }}
          organization = {{ $subject.Organization }}
          commonName = {{ $subject.CommonName }}
          serialNumber = {{ $subject.SerialNumber }}
          domainComponent = {{ $subject.DomainComponent }}
        {{end}}
        {{ $issuer := $infos.Subject }}
        {{if $issuer }}
        [frontends."{{ $frontendName }}".passTLSClientCert.infos.issuer]
          country = {{ $issuer.Country }}
          province = {{ $issuer.Province }}
          locality = {{ $issuer.Locality }}
          organization = {{ $issuer.Organization }}
          commonName = {{ $issuer.CommonName }}
          serialNumber = {{ $issuer.SerialNumber }}
          domainComponent = {{ $issuer.DomainComponent }}
        {{end}}
      {{end}}
    {{end}}

  {{if $frontend.Headers }}
  [frontends."{{ $frontendName }}".headers]
    SSLRedirect = {{ $frontend.Headers.SSLRedirect }}
    SSLTemporaryRedirect = {{ $frontend.Headers.SSLTemporaryRedirect }}
    SSLHost = "{{ $frontend.Headers.SSLHost }}"
    SSLForceHost = {{ $frontend.Headers.SSLForceHost }}
    STSSeconds = {{ $frontend.Headers.STSSeconds }}
    STSIncludeSubdomains = {{ $frontend.Headers.STSIncludeSubdomains }}
    STSPreload = {{ $frontend.Headers.STSPreload }}
    ForceSTSHeader = {{ $frontend.Headers.ForceSTSHeader }}
    FrameDeny = {{ $frontend.Headers.FrameDeny }}
    CustomFrameOptionsValue = "{{ $frontend.Headers.CustomFrameOptionsValue }}"
    ContentTypeNosniff = {{ $frontend.Headers.ContentTypeNosniff }}
    BrowserXSSFilter = {{ $frontend.Headers.BrowserXSSFilter }}
    CustomBrowserXSSValue = "{{ $frontend.Headers.CustomBrowserXSSValue }}"
    ContentSecurityPolicy = "{{ $frontend.Headers.ContentSecurityPolicy }}"
    PublicKey = "{{ $frontend.Headers.PublicKey }}"
    ReferrerPolicy = "{{ $frontend.Headers.ReferrerPolicy }}"
    IsDevelopment = {{ $frontend.Headers.IsDevelopment }}
    {{if $frontend.Headers.AllowedHosts }}
    AllowedHosts = [{{range $frontend.Headers.AllowedHosts }}
      "{{.}}",
      {{end}}]
    {{end}}
    {{if $frontend.Headers.HostsProxyHeaders }}
    HostsProxyHeaders = [{{range $frontend.Headers.HostsProxyHeaders }}
      "{{.}}",
      {{end}}]
    {{end}}
    {{if $frontend.Headers.CustomRequestHeaders }}
    [frontends."{{ $frontendName }}".headers.customRequestHeaders]
      {{range $k, $v := $frontend.Headers.CustomRequestHeaders }}
      {{ $k }} = "{{ $v }}"
      {{end}}
    {{end}}
    {{if $frontend.Headers.CustomResponseHeaders }}
    [frontends."{{ $frontendName }}".headers.customResponseHeaders]
      {{range $k, $v := $frontend.Headers.CustomResponseHeaders }}
      {{ $k }} = "{{ $v }}"
      {{end}}
    {{end}}
    {{if $frontend.Headers.SSLProxyHeaders }}
    [frontends."{{ $frontendName }}".headers.SSLProxyHeaders]
      {{range $k, $v := $frontend.Headers.SSLProxyHeaders }}
      {{ $k }} = "{{ $v }}"
      {{end}}
    {{end}}
  {{end}}

    {{range $routeName, $route := $frontend.Routes }}
    [frontends."{{ $frontendName }}".routes."{{ $routeName }}"]
      rule = "{{ $route.Rule }}"
    {{end}}

{{end}}

{{range $tls := .TLS }}
[[tls]]
  entryPoints = [{{range $tls.EntryPoints }}
    "{{.}}",
    {{end}}]
  [tls.certificate]
    certFile = """{{ $tls.Certificate.CertFile }}"""
    keyFile = """{{ $tls.Certificate.KeyFile }}"""
{{end}}
`)

func templatesKubernetesTmplBytes() ([]byte, error) {
	return _templatesKubernetesTmpl, nil
}

func templatesKubernetesTmpl() (*asset, error) {
	bytes, err := templatesKubernetesTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/kubernetes.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesKvTmpl = []byte(`[backends]
{{range $backend := List .Prefix "/backends/" }}
  {{ $backendName := Last $backend }}

  {{ $circuitBreaker := getCircuitBreaker $backend }}
  {{if $circuitBreaker }}
  [backends."{{ $backendName }}".circuitBreaker]
    expression = "{{ $circuitBreaker.Expression }}"
  {{end}}
          
  {{ $responseForwarding := getResponseForwarding $backend }}
  {{if $responseForwarding }}
  [backends."{{ $backendName }}".responseForwarding]
    flushInterval = "{{ $responseForwarding.flushInterval }}"
  {{end}}

  {{ $loadBalancer := getLoadBalancer $backend }}
  {{if $loadBalancer }}
    [backends."{{ $backendName }}".loadBalancer]
      method = "{{ $loadBalancer.Method }}"
      sticky = {{ $loadBalancer.Sticky }}
      {{if $loadBalancer.Stickiness }}
      [backends."{{ $backendName }}".loadBalancer.stickiness]
        cookieName = "{{ $loadBalancer.Stickiness.CookieName }}"
        secure = {{ $loadBalancer.Stickiness.Secure }}
        httpOnly = {{ $loadBalancer.Stickiness.HTTPOnly }}
        sameSite = "{{ $loadBalancer.Stickiness.SameSite }}"
      {{end}}
  {{end}}

  {{ $maxConn := getMaxConn $backend }}
  {{if $maxConn }}
  [backends."{{ $backendName }}".maxConn]
    extractorFunc = "{{ $maxConn.ExtractorFunc }}"
    amount = {{ $maxConn.Amount }}
  {{end}}

  {{ $healthCheck := getHealthCheck $backend }}
  {{if $healthCheck }}
  [backends."{{ $backendName }}".healthCheck]
    scheme = "{{ $healthCheck.Scheme }}"
    path = "{{ $healthCheck.Path }}"
    port = {{ $healthCheck.Port }}
    interval = "{{ $healthCheck.Interval }}"
    hostname = "{{ $healthCheck.Hostname }}"
    {{if $healthCheck.Headers }}
    [backends."{{ $backendName }}".healthCheck.headers]
      {{range $k, $v := $healthCheck.Headers }}
      {{$k}} = "{{$v}}"
      {{end}}
    {{end}}
  {{end}}

  {{ $buffering := getBuffering $backend }}
  {{if $buffering }}
  [backends."{{ $backendName }}".buffering]
    maxRequestBodyBytes = {{ $buffering.MaxRequestBodyBytes }}
    memRequestBodyBytes = {{ $buffering.MemRequestBodyBytes }}
    maxResponseBodyBytes = {{ $buffering.MaxResponseBodyBytes }}
    memResponseBodyBytes = {{ $buffering.MemResponseBodyBytes }}
    retryExpression = "{{ $buffering.RetryExpression }}"
  {{end}}

  {{range $serverName, $server := getServers $backend}}
  [backends."{{ $backendName }}".servers."{{ $serverName }}"]
    url = "{{ $server.URL }}"
    weight = {{ $server.Weight }}
  {{end}}

{{end}}

[frontends]
{{range $frontend := List .Prefix "/frontends/" }}
  {{ $frontendName := Last $frontend }}

  [frontends."{{ $frontendName }}"]
    backend = "{{ getBackendName $frontend }}"
    priority = {{ getPriority $frontend }}
    passHostHeader = {{ getPassHostHeader $frontend }}
    passTLSCert = {{ getPassTLSCert $frontend }}

    entryPoints = [{{range getEntryPoints $frontend }}
      "{{.}}",
      {{end}}]

    {{ $tlsClientCert := getPassTLSClientCert $frontend }}
    {{if $tlsClientCert }}
    [frontends."{{ $frontendName }}".passTLSClientCert]
      pem = {{ $tlsClientCert.PEM }}
      {{ $infos := $tlsClientCert.Infos }}
      {{if $infos }}
      [frontends."{{ $frontendName }}".passTLSClientCert.infos]
        notAfter = {{ $infos.NotAfter   }}
        notBefore = {{ $infos.NotBefore }}
        sans = {{ $infos.Sans }}
        {{ $subject := $infos.Subject }}
        {{if $subject }}
        [frontends."{{ $frontendName }}".passTLSClientCert.infos.subject]
          country = {{ $subject.Country }}
          province = {{ $subject.Province }}
          locality = {{ $subject.Locality }}
          organization = {{ $subject.Organization }}
          commonName = {{ $subject.CommonName }}
          serialNumber = {{ $subject.SerialNumber }}
          domainComponent = {{ $subject.DomainComponent }}
        {{end}}
        {{ $issuer := $infos.Subject }}
        {{if $issuer }}
        [frontends."{{ $frontendName }}".passTLSClientCert.infos.issuer]
          country = {{ $issuer.Country }}
          province = {{ $issuer.Province }}
          locality = {{ $issuer.Locality }}
          organization = {{ $issuer.Organization }}
          commonName = {{ $issuer.CommonName }}
          serialNumber = {{ $issuer.SerialNumber }}
          domainComponent = {{ $issuer.DomainComponent }}
        {{end}}
      {{end}}
    {{end}}

    {{ $auth := getAuth $frontend }}
    {{if $auth }}
    [frontends."{{ $frontendName }}".auth]
      headerField = "{{ $auth.HeaderField }}"

      {{if $auth.Forward }}
      [frontends."{{ $frontendName }}".auth.forward]
        address = "{{ $auth.Forward.Address }}"
        trustForwardHeader = {{ $auth.Forward.TrustForwardHeader }}
        {{if $auth.Forward.AuthResponseHeaders }}
        authResponseHeaders = [{{range $auth.Forward.AuthResponseHeaders }}
          "{{.}}",
          {{end}}]
        {{end}}

        {{if $auth.Forward.TLS }}
        [frontends."{{ $frontendName }}".auth.forward.tls]
          ca = "{{ $auth.Forward.TLS.CA }}"
          caOptional = {{ $auth.Forward.TLS.CAOptional }}
          cert = """{{ $auth.Forward.TLS.Cert }}"""
          key = """{{ $auth.Forward.TLS.Key }}"""
          insecureSkipVerify = {{ $auth.Forward.TLS.InsecureSkipVerify }}
        {{end}}
      {{end}}

      {{if $auth.Basic }}
      [frontends."{{ $frontendName }}".auth.basic]
        removeHeader = {{ $auth.Basic.RemoveHeader }}
        {{if $auth.Basic.Users }}
        users = [{{range $auth.Basic.Users }}
          "{{.}}",
          {{end}}]
        {{end}}
        usersFile = "{{ $auth.Basic.UsersFile }}"
      {{end}}

      {{if $auth.Digest }}
      [frontends."{{ $frontendName }}".auth.digest]
        removeHeader = {{ $auth.Digest.RemoveHeader }}
        {{if $auth.Digest.Users }}
        users = [{{range $auth.Digest.Users }}
          "{{.}}",
          {{end}}]
        {{end}}
        usersFile = "{{ $auth.Digest.UsersFile }}"
      {{end}}
    {{end}}

    {{ $whitelist := getWhiteList $frontend }}
    {{if $whitelist }}
    [frontends."{{ $frontendName }}".whiteList]
      sourceRange = [{{range $whitelist.SourceRange }}
        "{{.}}",
        {{end}}]
      useXForwardedFor = {{ $whitelist.UseXForwardedFor }}
    {{end}}

    {{ $redirect := getRedirect $frontend }}
    {{if $redirect }}
    [frontends."{{ $frontendName }}".redirect]
      entryPoint = "{{ $redirect.EntryPoint }}"
      regex = "{{ $redirect.Regex }}"
      replacement = "{{ $redirect.Replacement }}"
      permanent = {{ $redirect.Permanent }}
    {{end}}

    {{ $errorPages := getErrorPages $frontend }}
    {{if $errorPages }}
    [frontends."{{ $frontendName }}".errors]
      {{range $pageName, $page := $errorPages }}
      [frontends."{{$frontendName}}".errors."{{ $pageName }}"]
        status = [{{range $page.Status }}
          "{{.}}",
          {{end}}]
        backend = "{{$page.Backend}}"
        query = "{{$page.Query}}"
      {{end}}
    {{end}}

    {{ $rateLimit := getRateLimit $frontend }}
    {{if $rateLimit }}
    [frontends."{{ $frontendName }}".rateLimit]
      extractorFunc = "{{ $rateLimit.ExtractorFunc }}"
      [frontends."{{ $frontendName }}".rateLimit.rateSet]
        {{range $limitName, $rateLimit := $rateLimit.RateSet }}
        [frontends."{{ $frontendName }}".rateLimit.rateSet."{{ $limitName }}"]
          period = "{{ $rateLimit.Period }}"
          average = {{ $rateLimit.Average }}
          burst = {{ $rateLimit.Burst }}
        {{end}}
    {{end}}

    {{ $headers := getHeaders $frontend }}
    {{if $headers }}
    [frontends."{{ $frontendName }}".headers]
      SSLRedirect = {{ $headers.SSLRedirect }}
      SSLTemporaryRedirect = {{ $headers.SSLTemporaryRedirect }}
      SSLHost = "{{ $headers.SSLHost }}"
      SSLForceHost = {{ $headers.SSLForceHost }}
      STSSeconds = {{ $headers.STSSeconds }}
      STSIncludeSubdomains = {{ $headers.STSIncludeSubdomains }}
      STSPreload = {{ $headers.STSPreload }}
      ForceSTSHeader = {{ $headers.ForceSTSHeader }}
      FrameDeny = {{ $headers.FrameDeny }}
      CustomFrameOptionsValue = "{{ $headers.CustomFrameOptionsValue }}"
      ContentTypeNosniff = {{ $headers.ContentTypeNosniff }}
      BrowserXSSFilter = {{ $headers.BrowserXSSFilter }}
      CustomBrowserXSSValue = "{{ $headers.CustomBrowserXSSValue }}"
      ContentSecurityPolicy = "{{ $headers.ContentSecurityPolicy }}"
      PublicKey = "{{ $headers.PublicKey }}"
      ReferrerPolicy = "{{ $headers.ReferrerPolicy }}"
      IsDevelopment = {{ $headers.IsDevelopment }}

      {{if $headers.AllowedHosts }}
      AllowedHosts = [{{range $headers.AllowedHosts }}
        "{{.}}",
        {{end}}]
      {{end}}

      {{if $headers.HostsProxyHeaders }}
      HostsProxyHeaders = [{{range $headers.HostsProxyHeaders }}
        "{{.}}",
        {{end}}]
      {{end}}

      {{if $headers.CustomRequestHeaders }}
      [frontends."{{ $frontendName }}".headers.customRequestHeaders]
        {{range $k, $v := $headers.CustomRequestHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}

      {{if $headers.CustomResponseHeaders }}
      [frontends."{{ $frontendName }}".headers.customResponseHeaders]
        {{range $k, $v := $headers.CustomResponseHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}

      {{if $headers.SSLProxyHeaders }}
      [frontends."{{ $frontendName }}".headers.SSLProxyHeaders]
        {{range $k, $v := $headers.SSLProxyHeaders}}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}
    {{end}}

    {{range $routeName, $route := getRoutes $frontend }}
    [frontends."{{ $frontendName }}".routes."{{ $routeName }}"]
      rule = "{{ $route.Rule }}"
    {{end}}

{{end}}

{{range $tls := getTLSSection .Prefix }}
[[tls]]

  entryPoints = [{{range $tls.EntryPoints }}
    "{{.}}",
    {{end}}]

  [tls.certificate]
    certFile = """{{ $tls.Certificate.CertFile }}"""
    keyFile = """{{ $tls.Certificate.KeyFile }}"""

{{end}}
`)

func templatesKvTmplBytes() ([]byte, error) {
	return _templatesKvTmpl, nil
}

func templatesKvTmpl() (*asset, error) {
	bytes, err := templatesKvTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/kv.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMarathonV1Tmpl = []byte(`{{$apps := .Applications}}

{{range $app := $apps }}
{{range $task := $app.Tasks }}
{{range $serviceIndex, $serviceName := getServiceNames $app }}
  [backends."{{ getBackend $app $serviceName }}".servers."server-{{ $task.ID | replace "." "-"}}{{getServiceNameSuffix $serviceName }}"]
    url = "{{ getProtocol $app $serviceName }}://{{ getBackendServer $task $app }}:{{ getPort $task $app $serviceName }}"
    weight = {{ getWeight $app $serviceName }}
{{end}}
{{end}}
{{end}}

{{range $app := $apps }}
{{range $serviceIndex, $serviceName := getServiceNames $app }}

[backends."{{ getBackend $app $serviceName }}"]
  {{if hasMaxConnLabels $app }}
  [backends."{{ getBackend $app $serviceName }}".maxConn]
    amount = {{ getMaxConnAmount $app }}
    extractorFunc = "{{ getMaxConnExtractorFunc $app }}"
  {{end}}

  {{if hasLoadBalancerLabels $app }}
  [backends."{{ getBackend $app $serviceName }}".loadBalancer]
    method = "{{ getLoadBalancerMethod $app }}"
    sticky = {{ getSticky $app }}
    {{if hasStickinessLabel $app }}
    [backends."{{ getBackend $app $serviceName }}".loadBalancer.stickiness]
      cookieName = "{{ getStickinessCookieName $app }}"
    {{end}}
  {{end}}

  {{if hasCircuitBreakerLabels $app }}
  [backends."{{ getBackend $app $serviceName }}".circuitBreaker]
    expression = "{{ getCircuitBreakerExpression $app }}"
  {{end}}

  {{if hasHealthCheckLabels $app }}
  [backends."{{ getBackend $app $serviceName }}".healthCheck]
    path = "{{ getHealthCheckPath $app }}"
    interval = "{{ getHealthCheckInterval $app }}"
  {{end}}

{{end}}
{{end}}

[frontends]
{{range $app := $apps }}
{{range $serviceIndex, $serviceName := getServiceNames . }}

  [frontends."{{ getFrontendName $app $serviceName | normalize }}"]
    backend = "{{ getBackend $app $serviceName }}"
    passHostHeader = {{ getPassHostHeader $app $serviceName }}
    priority = {{ getPriority $app $serviceName }}

    entryPoints = [{{range getEntryPoints $app $serviceName }}
      "{{.}}",
      {{end}}]

    basicAuth = [{{range getBasicAuth $app $serviceName }}
      "{{.}}",
      {{end}}]

    [frontends."{{ getFrontendName $app $serviceName | normalize }}".routes."route-host{{ $app.ID | replace "/" "-" }}{{ getServiceNameSuffix $serviceName }}"]
      rule = "{{ getFrontendRule $app $serviceName }}"

{{end}}
{{end}}
`)

func templatesMarathonV1TmplBytes() ([]byte, error) {
	return _templatesMarathonV1Tmpl, nil
}

func templatesMarathonV1Tmpl() (*asset, error) {
	bytes, err := templatesMarathonV1TmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/marathon-v1.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMarathonTmpl = []byte(`{{ $apps := .Applications }}

[backends]
{{range $backendName, $app := $apps }}

  [backends."{{ $backendName }}"]

    {{ $circuitBreaker := getCircuitBreaker $app.SegmentLabels }}
    {{if $circuitBreaker }}
    [backends."{{ $backendName }}".circuitBreaker]
      expression = "{{ $circuitBreaker.Expression }}"
    {{end}}
          
    {{ $responseForwarding := getResponseForwarding $app.SegmentLabels }}
    {{if $responseForwarding }}
    [backends."{{ $backendName }}".responseForwarding]
      flushInterval = "{{ $responseForwarding.FlushInterval }}"
    {{end}}

    {{ $loadBalancer := getLoadBalancer $app.SegmentLabels }}
    {{if $loadBalancer }}
    [backends."{{ $backendName }}".loadBalancer]
      method = "{{ $loadBalancer.Method }}"
      sticky = {{ $loadBalancer.Sticky }}
      {{if $loadBalancer.Stickiness }}
      [backends."{{ $backendName }}".loadBalancer.stickiness]
        cookieName = "{{ $loadBalancer.Stickiness.CookieName }}"
        secure = {{ $loadBalancer.Stickiness.Secure }}
        httpOnly = {{ $loadBalancer.Stickiness.HTTPOnly }}
        sameSite = "{{ $loadBalancer.Stickiness.SameSite }}"
      {{end}}
    {{end}}

    {{ $maxConn := getMaxConn $app.SegmentLabels }}
    {{if $maxConn }}
    [backends."{{ $backendName }}".maxConn]
      extractorFunc = "{{ $maxConn.ExtractorFunc }}"
      amount = {{ $maxConn.Amount }}
    {{end}}

    {{ $healthCheck := getHealthCheck $app.SegmentLabels }}
    {{if $healthCheck }}
    [backends."{{ $backendName }}".healthCheck]
      scheme = "{{ $healthCheck.Scheme }}"
      path = "{{ $healthCheck.Path }}"
      port = {{ $healthCheck.Port }}
      interval = "{{ $healthCheck.Interval }}"
      hostname = "{{ $healthCheck.Hostname }}"
      {{if $healthCheck.Headers }}
      [backends.{{ $backendName }}.healthCheck.headers]
        {{range $k, $v := $healthCheck.Headers }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}
    {{end}}

    {{ $buffering := getBuffering $app.SegmentLabels }}
    {{if $buffering }}
    [backends."{{ $backendName }}".buffering]
      maxRequestBodyBytes = {{ $buffering.MaxRequestBodyBytes }}
      memRequestBodyBytes = {{ $buffering.MemRequestBodyBytes }}
      maxResponseBodyBytes = {{ $buffering.MaxResponseBodyBytes }}
      memResponseBodyBytes = {{ $buffering.MemResponseBodyBytes }}
      retryExpression = "{{ $buffering.RetryExpression }}"
    {{end}}

    {{range $serverName, $server := getServers $app }}
    [backends."{{ $backendName }}".servers."{{ $serverName }}"]
      url = "{{ $server.URL }}"
      weight = {{ $server.Weight }}
    {{end}}

{{end}}

[frontends]
{{range $backendName, $app := $apps }}
  {{ $frontendName := getFrontendName $app }}

  [frontends."{{ $frontendName }}"]
    backend = "{{ $backendName }}"
    priority = {{ getPriority $app.SegmentLabels }}
    passHostHeader = {{ getPassHostHeader $app.SegmentLabels }}
    passTLSCert = {{ getPassTLSCert $app.SegmentLabels }}

    entryPoints = [{{range getEntryPoints $app.SegmentLabels }}
      "{{.}}",
      {{end}}]

    {{ $tlsClientCert := getPassTLSClientCert $app.SegmentLabels }}
    {{if $tlsClientCert }}
    [frontends."{{ $frontendName }}".passTLSClientCert]
      pem = {{ $tlsClientCert.PEM }}
      {{ $infos := $tlsClientCert.Infos }}
      {{if $infos }}
      [frontends."{{ $frontendName }}".passTLSClientCert.infos]
        notAfter = {{ $infos.NotAfter   }}
        notBefore = {{ $infos.NotBefore }}
        sans = {{ $infos.Sans }}
        {{ $subject := $infos.Subject }}
        {{if $subject }}
        [frontends."{{ $frontendName }}".passTLSClientCert.infos.subject]
          country = {{ $subject.Country }}
          province = {{ $subject.Province }}
          locality = {{ $subject.Locality }}
          organization = {{ $subject.Organization }}
          commonName = {{ $subject.CommonName }}
          serialNumber = {{ $subject.SerialNumber }}
          domainComponent = {{ $subject.DomainComponent }}
        {{end}}
        {{ $issuer := $infos.Subject }}
        {{if $issuer }}
        [frontends."{{ $frontendName }}".passTLSClientCert.infos.issuer]
          country = {{ $issuer.Country }}
          province = {{ $issuer.Province }}
          locality = {{ $issuer.Locality }}
          organization = {{ $issuer.Organization }}
          commonName = {{ $issuer.CommonName }}
          serialNumber = {{ $issuer.SerialNumber }}
          domainComponent = {{ $issuer.DomainComponent }}
        {{end}}
      {{end}}
    {{end}}

  {{ $auth := getAuth $app.SegmentLabels }}
    {{if $auth }}
    [frontends."{{ $frontendName }}".auth]
      headerField = "{{ $auth.HeaderField }}"

      {{if $auth.Forward }}
      [frontends."{{ $frontendName }}".auth.forward]
        address = "{{ $auth.Forward.Address }}"
        trustForwardHeader = {{ $auth.Forward.TrustForwardHeader }}
        {{if $auth.Forward.AuthResponseHeaders }}
        authResponseHeaders = [{{range $auth.Forward.AuthResponseHeaders }}
          "{{.}}",
          {{end}}]
        {{end}}

        {{if $auth.Forward.TLS }}
        [frontends."{{ $frontendName }}".auth.forward.tls]
          ca = "{{ $auth.Forward.TLS.CA }}"
          caOptional = {{ $auth.Forward.TLS.CAOptional }}
          cert = """{{ $auth.Forward.TLS.Cert }}"""
          key = """{{ $auth.Forward.TLS.Key }}"""
          insecureSkipVerify = {{ $auth.Forward.TLS.InsecureSkipVerify }}
        {{end}}
      {{end}}

      {{if $auth.Basic }}
      [frontends."{{ $frontendName }}".auth.basic]
        removeHeader = {{ $auth.Basic.RemoveHeader }}
        {{if $auth.Basic.Users }}
        users = [{{range $auth.Basic.Users }}
          "{{.}}",
          {{end}}]
        {{end}}
        usersFile = "{{ $auth.Basic.UsersFile }}"
      {{end}}

      {{if $auth.Digest }}
      [frontends."{{ $frontendName }}".auth.digest]
        removeHeader = {{ $auth.Digest.RemoveHeader }}
        {{if $auth.Digest.Users }}
        users = [{{range $auth.Digest.Users }}
          "{{.}}",
          {{end}}]
        {{end}}
        usersFile = "{{ $auth.Digest.UsersFile }}"
      {{end}}
    {{end}}

    {{ $whitelist := getWhiteList $app.SegmentLabels }}
    {{if $whitelist }}
    [frontends."{{ $frontendName }}".whiteList]
      sourceRange = [{{range $whitelist.SourceRange }}
        "{{.}}",
        {{end}}]
      useXForwardedFor = {{ $whitelist.UseXForwardedFor }}
    {{end}}

    {{ $redirect := getRedirect $app.SegmentLabels }}
    {{if $redirect }}
    [frontends."{{ $frontendName }}".redirect]
      entryPoint = "{{ $redirect.EntryPoint }}"
      regex = "{{ $redirect.Regex }}"
      replacement = "{{ $redirect.Replacement }}"
      permanent = {{ $redirect.Permanent }}
    {{end}}

    {{ $errorPages := getErrorPages $app.SegmentLabels }}
    {{if $errorPages }}
    [frontends."{{ $frontendName }}".errors]
      {{range $pageName, $page := $errorPages }}
      [frontends."{{ $frontendName }}".errors."{{ $pageName }}"]
        status = [{{range $page.Status }}
          "{{.}}",
          {{end}}]
        backend = "backend{{ $page.Backend }}"
        query = "{{ $page.Query }}"
      {{end}}
    {{end}}

    {{ $rateLimit := getRateLimit $app.SegmentLabels }}
    {{if $rateLimit }}
    [frontends."{{ $frontendName }}".rateLimit]
      extractorFunc = "{{ $rateLimit.ExtractorFunc }}"
      [frontends."{{ $frontendName }}".rateLimit.rateSet]
        {{ range $limitName, $limit := $rateLimit.RateSet }}
        [frontends."{{ $frontendName }}".rateLimit.rateSet."{{ $limitName }}"]
          period = "{{ $limit.Period }}"
          average = {{ $limit.Average }}
          burst = {{ $limit.Burst }}
        {{end}}
    {{end}}

    {{ $headers := getHeaders $app.SegmentLabels }}
    {{if $headers }}
    [frontends."{{ $frontendName }}".headers]
      SSLRedirect = {{ $headers.SSLRedirect }}
      SSLTemporaryRedirect = {{ $headers.SSLTemporaryRedirect }}
      SSLHost = "{{ $headers.SSLHost }}"
      SSLForceHost = {{ $headers.SSLForceHost }}
      STSSeconds = {{ $headers.STSSeconds }}
      STSIncludeSubdomains = {{ $headers.STSIncludeSubdomains }}
      STSPreload = {{ $headers.STSPreload }}
      ForceSTSHeader = {{ $headers.ForceSTSHeader }}
      FrameDeny = {{ $headers.FrameDeny }}
      CustomFrameOptionsValue = "{{ $headers.CustomFrameOptionsValue }}"
      ContentTypeNosniff = {{ $headers.ContentTypeNosniff }}
      BrowserXSSFilter = {{ $headers.BrowserXSSFilter }}
      CustomBrowserXSSValue = "{{ $headers.CustomBrowserXSSValue }}"
      ContentSecurityPolicy = "{{ $headers.ContentSecurityPolicy }}"
      PublicKey = "{{ $headers.PublicKey }}"
      ReferrerPolicy = "{{ $headers.ReferrerPolicy }}"
      IsDevelopment = {{ $headers.IsDevelopment }}

      {{if $headers.AllowedHosts }}
      AllowedHosts = [{{range $headers.AllowedHosts }}
        "{{.}}",
        {{end}}]
      {{end}}

      {{if $headers.HostsProxyHeaders }}
      HostsProxyHeaders = [{{range $headers.HostsProxyHeaders }}
        "{{.}}",
        {{end}}]
      {{end}}

      {{if $headers.CustomRequestHeaders }}
      [frontends."{{ $frontendName }}".headers.customRequestHeaders]
        {{range $k, $v := $headers.CustomRequestHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}

      {{if $headers.CustomResponseHeaders }}
      [frontends."{{ $frontendName }}".headers.customResponseHeaders]
        {{range $k, $v := $headers.CustomResponseHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}

      {{if $headers.SSLProxyHeaders }}
      [frontends."{{ $frontendName }}".headers.SSLProxyHeaders]
        {{range $k, $v := $headers.SSLProxyHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}
    {{end}}

  [frontends."{{ $frontendName }}".routes."route-host{{ $app.ID | replace "/" "-" }}{{ getSegmentNameSuffix $app.SegmentName }}"]
    rule = "{{ getFrontendRule $app }}"

{{end}}
`)

func templatesMarathonTmplBytes() ([]byte, error) {
	return _templatesMarathonTmpl, nil
}

func templatesMarathonTmpl() (*asset, error) {
	bytes, err := templatesMarathonTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/marathon.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMesosV1Tmpl = []byte(`{{$apps := .Applications}}

[backends]
{{range .Tasks}}

  [backends."backend-{{ getBackend . $apps }}".servers."server-{{ getID . }}"]
    url = "{{ getProtocol . $apps }}://{{ getHost . }}:{{ getPort . $apps }}"
    weight = {{ getWeight . $apps }}

{{end}}

[frontends]
{{range .Applications}}

  [frontends."frontend-{{getFrontEndName . }}"]
    backend = "backend-{{ getFrontendBackend . }}"
    passHostHeader = {{ getPassHostHeader . }}
    priority = {{ getPriority . }}

    entryPoints = [{{range getEntryPoints . }}
      "{{.}}",
      {{end}}]

    [frontends."frontend-{{ getFrontEndName . }}".routes."route-host-{{ getFrontEndName . }}"]
    rule = "{{ getFrontendRule . }}"

{{end}}
`)

func templatesMesosV1TmplBytes() ([]byte, error) {
	return _templatesMesosV1Tmpl, nil
}

func templatesMesosV1Tmpl() (*asset, error) {
	bytes, err := templatesMesosV1TmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/mesos-v1.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMesosTmpl = []byte(`[backends]
{{range $applicationName, $tasks := .ApplicationsTasks }}
  {{ $app := index $tasks 0 }}
  {{ $backendName := getBackendName $app }}

  [backends."backend-{{ $backendName }}"]

  {{ $circuitBreaker := getCircuitBreaker $app.TraefikLabels }}
  {{if $circuitBreaker }}
  [backends."backend-{{ $backendName }}".circuitBreaker]
    expression = "{{ $circuitBreaker.Expression }}"
  {{end}}

  {{ $responseForwarding := getResponseForwarding $app.TraefikLabels }}
  {{if $responseForwarding }}
  [backends."backend-{{ $backendName }}".responseForwarding]
    flushInterval = "{{ $responseForwarding.FlushInterval }}"
  {{end}}

  {{ $loadBalancer := getLoadBalancer $app.TraefikLabels }}
  {{if $loadBalancer }}
    [backends."backend-{{ $backendName }}".loadBalancer]
      method = "{{ $loadBalancer.Method }}"
      sticky = {{ $loadBalancer.Sticky }}
      {{if $loadBalancer.Stickiness }}
      [backends."backend-{{ $backendName }}".loadBalancer.stickiness]
        cookieName = "{{ $loadBalancer.Stickiness.CookieName }}"
        secure = {{ $loadBalancer.Stickiness.Secure }}
        httpOnly = {{ $loadBalancer.Stickiness.HTTPOnly }}
        sameSite = "{{ $loadBalancer.Stickiness.SameSite }}"
      {{end}}
  {{end}}

  {{ $maxConn := getMaxConn $app.TraefikLabels }}
  {{if $maxConn }}
  [backends."backend-{{ $backendName }}".maxConn]
    extractorFunc = "{{ $maxConn.ExtractorFunc }}"
    amount = {{ $maxConn.Amount }}
  {{end}}

  {{ $healthCheck := getHealthCheck $app.TraefikLabels }}
  {{if $healthCheck }}
  [backends."backend-{{ $backendName }}".healthCheck]
    scheme = "{{ $healthCheck.Scheme }}"
    path = "{{ $healthCheck.Path }}"
    port = {{ $healthCheck.Port }}
    interval = "{{ $healthCheck.Interval }}"
    hostname = "{{ $healthCheck.Hostname }}"
    {{if $healthCheck.Headers }}
    [backends."backend-{{ $backendName }}".healthCheck.headers]
      {{range $k, $v := $healthCheck.Headers }}
      {{$k}} = "{{$v}}"
      {{end}}
    {{end}}
  {{end}}

  {{ $buffering := getBuffering $app.TraefikLabels }}
  {{if $buffering }}
  [backends."backend-{{ $backendName }}".buffering]
    maxRequestBodyBytes = {{ $buffering.MaxRequestBodyBytes }}
    memRequestBodyBytes = {{ $buffering.MemRequestBodyBytes }}
    maxResponseBodyBytes = {{ $buffering.MaxResponseBodyBytes }}
    memResponseBodyBytes = {{ $buffering.MemResponseBodyBytes }}
    retryExpression = "{{ $buffering.RetryExpression }}"
  {{end}}

  {{range $serverName, $server := getServers $tasks }}
  [backends."backend-{{ $backendName }}".servers."{{ $serverName }}"]
    url = "{{ $server.URL }}"
    weight = {{ $server.Weight }}
  {{end}}
{{end}}

[frontends]
{{range $applicationName, $tasks := .ApplicationsTasks }}
  {{ $app := index $tasks 0 }}
  {{ $frontendName := getFrontEndName $app }}

  [frontends."frontend-{{ $frontendName }}"]
    backend = "backend-{{ getBackendName $app }}"
    priority = {{ getPriority $app.TraefikLabels }}
    passHostHeader = {{ getPassHostHeader $app.TraefikLabels }}
    passTLSCert = {{ getPassTLSCert $app.TraefikLabels }}

    entryPoints = [{{range getEntryPoints $app.TraefikLabels }}
      "{{.}}",
      {{end}}]

    {{ $tlsClientCert := getPassTLSClientCert $app.TraefikLabels }}
    {{if $tlsClientCert }}
    [frontends."frontend-{{ $frontendName }}".passTLSClientCert]
      pem = {{ $tlsClientCert.PEM }}
      {{ $infos := $tlsClientCert.Infos }}
      {{if $infos }}
      [frontends."frontend-{{ $frontendName }}".passTLSClientCert.infos]
        notAfter = {{ $infos.NotAfter   }}
        notBefore = {{ $infos.NotBefore }}
        sans = {{ $infos.Sans }}
        {{ $subject := $infos.Subject }}
        {{if $subject }}
        [frontends."frontend-{{ $frontendName }}".passTLSClientCert.infos.subject]
          country = {{ $subject.Country }}
          province = {{ $subject.Province }}
          locality = {{ $subject.Locality }}
          organization = {{ $subject.Organization }}
          commonName = {{ $subject.CommonName }}
          serialNumber = {{ $subject.SerialNumber }}
          domainComponent = {{ $subject.DomainComponent }}
        {{end}}
        {{ $issuer := $infos.Subject }}
        {{if $issuer }}
        [frontends."frontend-{{ $frontendName }}".passTLSClientCert.infos.issuer]
          country = {{ $issuer.Country }}
          province = {{ $issuer.Province }}
          locality = {{ $issuer.Locality }}
          organization = {{ $issuer.Organization }}
          commonName = {{ $issuer.CommonName }}
          serialNumber = {{ $issuer.SerialNumber }}
          domainComponent = {{ $issuer.DomainComponent }}
        {{end}}
      {{end}}
    {{end}}

    {{ $auth := getAuth $app.TraefikLabels }}
    {{if $auth }}
    [frontends."frontend-{{ $frontendName }}".auth]
      headerField = "{{ $auth.HeaderField }}"

      {{if $auth.Forward }}
      [frontends."frontend-{{ $frontendName }}".auth.forward]
        address = "{{ $auth.Forward.Address }}"
        trustForwardHeader = {{ $auth.Forward.TrustForwardHeader }}
        {{if $auth.Forward.AuthResponseHeaders }}
        authResponseHeaders = [{{range $auth.Forward.AuthResponseHeaders }}
          "{{.}}",
          {{end}}]
        {{end}}

        {{if $auth.Forward.TLS }}
        [frontends."frontend-{{ $frontendName }}".auth.forward.tls]
          ca = "{{ $auth.Forward.TLS.CA }}"
          caOptional = {{ $auth.Forward.TLS.CAOptional }}
          cert = """{{ $auth.Forward.TLS.Cert }}"""
          key = """{{ $auth.Forward.TLS.Key }}"""
          insecureSkipVerify = {{ $auth.Forward.TLS.InsecureSkipVerify }}
        {{end}}
      {{end}}

      {{if $auth.Basic }}
      [frontends."frontend-{{ $frontendName }}".auth.basic]
        removeHeader = {{ $auth.Basic.RemoveHeader}}
        {{if $auth.Basic.Users }}
        users = [{{range $auth.Basic.Users }}
          "{{.}}",
          {{end}}]
        {{end}}
        usersFile = "{{ $auth.Basic.UsersFile }}"
      {{end}}

      {{if $auth.Digest }}
      [frontends."frontend-{{ $frontendName }}".auth.digest]
        removeHeader = {{ $auth.Digest.RemoveHeader}}
        {{if $auth.Digest.Users }}
        users = [{{range $auth.Digest.Users }}
          "{{.}}",
          {{end}}]
        {{end}}
        usersFile = "{{ $auth.Digest.UsersFile }}"
      {{end}}
    {{end}}
          
    {{ $whitelist := getWhiteList $app.TraefikLabels }}
    {{if $whitelist }}
    [frontends."frontend-{{ $frontendName }}".whiteList]
      sourceRange = [{{range $whitelist.SourceRange }}
        "{{.}}",
        {{end}}]
      useXForwardedFor = {{ $whitelist.UseXForwardedFor }}
    {{end}}

    {{ $redirect := getRedirect $app.TraefikLabels }}
    {{if $redirect }}
    [frontends."frontend-{{ $frontendName }}".redirect]
      entryPoint = "{{ $redirect.EntryPoint }}"
      regex = "{{ $redirect.Regex }}"
      replacement = "{{ $redirect.Replacement }}"
      permanent = {{ $redirect.Permanent }}
    {{end}}

    {{ $errorPages := getErrorPages $app.TraefikLabels }}
    {{if $errorPages }}
    [frontends."frontend-{{ $frontendName }}".errors]
      {{range $pageName, $page := $errorPages }}
      [frontends."frontend-{{ $frontendName }}".errors."{{ $pageName }}"]
        status = [{{range $page.Status }}
        "{{.}}",
        {{end}}]
        backend = "backend-{{ $page.Backend }}"
        query = "{{ $page.Query }}"
      {{end}}
    {{end}}

    {{ $rateLimit := getRateLimit $app.TraefikLabels }}
    {{if $rateLimit }}
    [frontends."frontend-{{ $frontendName }}".rateLimit]
      extractorFunc = "{{ $rateLimit.ExtractorFunc }}"
      [frontends."frontend-{{ $frontendName }}".rateLimit.rateSet]
        {{ range $limitName, $limit := $rateLimit.RateSet }}
        [frontends."frontend-{{ $frontendName }}".rateLimit.rateSet."{{ $limitName }}"]
          period = "{{ $limit.Period }}"
          average = {{ $limit.Average }}
          burst = {{ $limit.Burst }}
        {{end}}
    {{end}}

    {{ $headers := getHeaders $app.TraefikLabels }}
    {{if $headers }}
    [frontends."frontend-{{ $frontendName }}".headers]
      SSLRedirect = {{ $headers.SSLRedirect }}
      SSLTemporaryRedirect = {{ $headers.SSLTemporaryRedirect }}
      SSLHost = "{{ $headers.SSLHost }}"
      SSLForceHost = {{ $headers.SSLForceHost }}
      STSSeconds = {{ $headers.STSSeconds }}
      STSIncludeSubdomains = {{ $headers.STSIncludeSubdomains }}
      STSPreload = {{ $headers.STSPreload }}
      ForceSTSHeader = {{ $headers.ForceSTSHeader }}
      FrameDeny = {{ $headers.FrameDeny }}
      CustomFrameOptionsValue = "{{ $headers.CustomFrameOptionsValue }}"
      ContentTypeNosniff = {{ $headers.ContentTypeNosniff }}
      BrowserXSSFilter = {{ $headers.BrowserXSSFilter }}
      CustomBrowserXSSValue = "{{ $headers.CustomBrowserXSSValue }}"
      ContentSecurityPolicy = "{{ $headers.ContentSecurityPolicy }}"
      PublicKey = "{{ $headers.PublicKey }}"
      ReferrerPolicy = "{{ $headers.ReferrerPolicy }}"
      IsDevelopment = {{ $headers.IsDevelopment }}

      {{if $headers.AllowedHosts }}
      AllowedHosts = [{{range $headers.AllowedHosts }}
        "{{.}}",
        {{end}}]
      {{end}}

      {{if $headers.HostsProxyHeaders }}
      HostsProxyHeaders = [{{range $headers.HostsProxyHeaders }}
        "{{.}}",
        {{end}}]
      {{end}}

      {{if $headers.CustomRequestHeaders }}
      [frontends."frontend-{{ $frontendName }}".headers.customRequestHeaders]
        {{range $k, $v := $headers.CustomRequestHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}

      {{if $headers.CustomResponseHeaders }}
      [frontends."frontend-{{ $frontendName }}".headers.customResponseHeaders]
        {{range $k, $v := $headers.CustomResponseHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}

      {{if $headers.SSLProxyHeaders }}
      [frontends."frontend-{{ $frontendName }}".headers.SSLProxyHeaders]
        {{range $k, $v := $headers.SSLProxyHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}
    {{end}}

    [frontends."frontend-{{ $frontendName }}".routes."route-host-{{ $frontendName }}"]
    rule = "{{ getFrontendRule $app }}"

{{end}}`)

func templatesMesosTmplBytes() ([]byte, error) {
	return _templatesMesosTmpl, nil
}

func templatesMesosTmpl() (*asset, error) {
	bytes, err := templatesMesosTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/mesos.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNotfoundTmpl = []byte(`<!DOCTYPE html>
<html>
<head>
    <title>Traefik</title>
</head>
<body>
    Ohhhh man, this is bad...
</body>
</html>`)

func templatesNotfoundTmplBytes() ([]byte, error) {
	return _templatesNotfoundTmpl, nil
}

func templatesNotfoundTmpl() (*asset, error) {
	bytes, err := templatesNotfoundTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/notFound.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRancherV1Tmpl = []byte(`{{$backendServers := .Backends}}

[backends]
{{range $backendName, $backend := .Backends }}
  {{if hasCircuitBreakerLabel $backend }}
  [backends."backend-{{ $backendName }}".circuitBreaker]
    expression = "{{ getCircuitBreakerExpression $backend }}"
  {{end}}

  {{if hasLoadBalancerLabel $backend }}
  [backends."backend-{{ $backendName }}".loadBalancer]
    method = "{{ getLoadBalancerMethod $backend }}"
    sticky = {{ getSticky $backend }}
    {{if hasStickinessLabel $backend }}
    [backends."backend-{{ $backendName }}".loadBalancer.stickiness]
      cookieName = "{{ getStickinessCookieName $backend }}"
    {{end}}
  {{end}}

  {{if hasMaxConnLabels $backend }}
  [backends."backend-{{ $backendName }}".maxConn]
    amount = {{ getMaxConnAmount $backend }}
    extractorFunc = "{{ getMaxConnExtractorFunc $backend }}"
  {{end}}

  {{range $index, $ip := $backend.Containers }}
  [backends."backend-{{ $backendName }}".servers."server-{{ $index }}"]
    url = "{{ getProtocol $backend }}://{{ $ip }}:{{ getPort $backend }}"
    weight = {{ getWeight $backend }}
  {{end}}

{{end}}

[frontends]
{{range $frontendName, $service := .Frontends }}
  [frontends."frontend-{{ $frontendName }}"]
    backend = "backend-{{ getBackend $service }}"
    passHostHeader = {{ getPassHostHeader $service }}
    priority = {{ getPriority $service }}

  entryPoints = [{{range getEntryPoints $service }}
      "{{.}}",
    {{end}}]

  basicAuth = [{{range getBasicAuth $service }}
      "{{.}}",
    {{end}}]

  {{if hasRedirect $service }}
  [frontends."frontend-{{ $frontendName }}".redirect]
    entryPoint = "{{ getRedirectEntryPoint $service }}"
    regex = "{{ getRedirectRegex $service }}"
    replacement = "{{ getRedirectReplacement $service }}"
  {{end}}

  [frontends."frontend-{{ $frontendName }}".routes."route-frontend-{{ $frontendName }}"]
    rule = "{{ getFrontendRule $service }}"
{{end}}
`)

func templatesRancherV1TmplBytes() ([]byte, error) {
	return _templatesRancherV1Tmpl, nil
}

func templatesRancherV1Tmpl() (*asset, error) {
	bytes, err := templatesRancherV1TmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/rancher-v1.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRancherTmpl = []byte(`{{ $backendServers := .Backends }}
[backends]
{{range $backendName, $backend := .Backends }}

  [backends."backend-{{ $backendName }}"]

  {{ $circuitBreaker := getCircuitBreaker $backend.SegmentLabels }}
  {{if $circuitBreaker }}
  [backends."backend-{{ $backendName }}".circuitBreaker]
    expression = "{{ $circuitBreaker.Expression }}"
  {{end}}

  {{ $responseForwarding := getResponseForwarding $backend.SegmentLabels }}
  {{if $responseForwarding }}
  [backends."backend-{{ $backendName }}".responseForwarding]
    flushInterval = "{{ $responseForwarding.FlushInterval }}"
  {{end}}

  {{ $loadBalancer := getLoadBalancer $backend.SegmentLabels }}
  {{if $loadBalancer }}
    [backends."backend-{{ $backendName }}".loadBalancer]
      method = "{{ $loadBalancer.Method }}"
      sticky = {{ $loadBalancer.Sticky }}
      {{if $loadBalancer.Stickiness }}
      [backends."backend-{{ $backendName }}".loadBalancer.stickiness]
        cookieName = "{{ $loadBalancer.Stickiness.CookieName }}"
        secure = {{ $loadBalancer.Stickiness.Secure }}
        httpOnly = {{ $loadBalancer.Stickiness.HTTPOnly }}
        sameSite = "{{ $loadBalancer.Stickiness.SameSite }}"
      {{end}}
  {{end}}

  {{ $maxConn := getMaxConn $backend.SegmentLabels }}
  {{if $maxConn }}
  [backends."backend-{{ $backendName }}".maxConn]
    extractorFunc = "{{ $maxConn.ExtractorFunc }}"
    amount = {{ $maxConn.Amount }}
  {{end}}

  {{ $healthCheck := getHealthCheck $backend.SegmentLabels }}
  {{if $healthCheck }}
  [backends."backend-{{ $backendName }}".healthCheck]
    scheme = "{{ $healthCheck.Scheme }}"
    path = "{{ $healthCheck.Path }}"
    port = {{ $healthCheck.Port }}
    interval = "{{ $healthCheck.Interval }}"
    hostname = "{{ $healthCheck.Hostname }}"
    {{if $healthCheck.Headers }}
    [backends."backend-{{ $backendName }}".healthCheck.headers]
      {{range $k, $v := $healthCheck.Headers }}
      {{$k}} = "{{$v}}"
      {{end}}
    {{end}}
  {{end}}

  {{ $buffering := getBuffering $backend.SegmentLabels }}
  {{if $buffering }}
  [backends."backend-{{ $backendName }}".buffering]
    maxRequestBodyBytes = {{ $buffering.MaxRequestBodyBytes }}
    memRequestBodyBytes = {{ $buffering.MemRequestBodyBytes }}
    maxResponseBodyBytes = {{ $buffering.MaxResponseBodyBytes }}
    memResponseBodyBytes = {{ $buffering.MemResponseBodyBytes }}
    retryExpression = "{{ $buffering.RetryExpression }}"
  {{end}}

  {{range $serverName, $server := getServers $backend}}
  [backends."backend-{{ $backendName }}".servers."{{ $serverName }}"]
    url = "{{ $server.URL }}"
    weight = {{ $server.Weight }}
  {{end}}

{{end}}

[frontends]
{{range $frontendName, $service := .Frontends }}

  [frontends."frontend-{{ $frontendName }}"]
    backend = "backend-{{ getBackendName $service }}"
    priority = {{ getPriority $service.SegmentLabels }}
    passHostHeader = {{ getPassHostHeader $service.SegmentLabels }}
    passTLSCert = {{ getPassTLSCert $service.SegmentLabels }}

    entryPoints = [{{range getEntryPoints $service.SegmentLabels }}
      "{{.}}",
      {{end}}]

    {{ $tlsClientCert := getPassTLSClientCert $service.SegmentLabels }}
    {{if $tlsClientCert }}
    [frontends."frontend-{{ $frontendName }}".passTLSClientCert]
      pem = {{ $tlsClientCert.PEM }}
      {{ $infos := $tlsClientCert.Infos }}
      {{if $infos }}
      [frontends."frontend-{{ $frontendName }}".passTLSClientCert.infos]
        notAfter = {{ $infos.NotAfter   }}
        notBefore = {{ $infos.NotBefore }}
        sans = {{ $infos.Sans }}
        {{ $subject := $infos.Subject }}
        {{if $subject }}
        [frontends."frontend-{{ $frontendName }}".passTLSClientCert.infos.subject]
          country = {{ $subject.Country }}
          province = {{ $subject.Province }}
          locality = {{ $subject.Locality }}
          organization = {{ $subject.Organization }}
          commonName = {{ $subject.CommonName }}
          serialNumber = {{ $subject.SerialNumber }}
          domainComponent = {{ $subject.DomainComponent }}
        {{end}}
        {{ $issuer := $infos.Subject }}
        {{if $issuer }}
        [frontends."frontend-{{ $frontendName }}".passTLSClientCert.infos.issuer]
          country = {{ $issuer.Country }}
          province = {{ $issuer.Province }}
          locality = {{ $issuer.Locality }}
          organization = {{ $issuer.Organization }}
          commonName = {{ $issuer.CommonName }}
          serialNumber = {{ $issuer.SerialNumber }}
          domainComponent = {{ $issuer.DomainComponent }}
        {{end}}
      {{end}}
    {{end}}

    {{ $auth := getAuth $service.SegmentLabels }}
    {{if $auth }}
    [frontends."frontend-{{ $frontendName }}".auth]
      headerField = "{{ $auth.HeaderField }}"

      {{if $auth.Forward }}
      [frontends."frontend-{{ $frontendName }}".auth.forward]
        address = "{{ $auth.Forward.Address }}"
        trustForwardHeader = {{ $auth.Forward.TrustForwardHeader }}
        {{if $auth.Forward.AuthResponseHeaders }}
        authResponseHeaders = [{{range $auth.Forward.AuthResponseHeaders }}
          "{{.}}",
          {{end}}]
        {{end}}

        {{if $auth.Forward.TLS }}
        [frontends."frontend-{{ $frontendName }}".auth.forward.tls]
          ca = "{{ $auth.Forward.TLS.CA }}"
          caOptional = {{ $auth.Forward.TLS.CAOptional }}
          cert = """{{ $auth.Forward.TLS.Cert }}"""
          key = """{{ $auth.Forward.TLS.Key }}"""
          insecureSkipVerify = {{ $auth.Forward.TLS.InsecureSkipVerify }}
        {{end}}
      {{end}}

      {{if $auth.Basic }}
      [frontends."frontend-{{ $frontendName }}".auth.basic]
        removeHeader = {{ $auth.Basic.RemoveHeader }}
        {{if $auth.Basic.Users }}
        users = [{{range $auth.Basic.Users }}
          "{{.}}",
          {{end}}]
        {{end}}
        usersFile = "{{ $auth.Basic.UsersFile }}"
      {{end}}

      {{if $auth.Digest }}
      [frontends."frontend-{{ $frontendName }}".auth.digest]
        removeHeader = {{ $auth.Digest.RemoveHeader }}
        {{if $auth.Digest.Users }}
        users = [{{range $auth.Digest.Users }}
          "{{.}}",
          {{end}}]
        {{end}}
        usersFile = "{{ $auth.Digest.UsersFile }}"
      {{end}}
    {{end}}

    {{ $whitelist := getWhiteList $service.SegmentLabels }}
    {{if $whitelist }}
    [frontends."frontend-{{ $frontendName }}".whiteList]
      sourceRange = [{{range $whitelist.SourceRange }}
        "{{.}}",
        {{end}}]
      useXForwardedFor = {{ $whitelist.UseXForwardedFor }}
    {{end}}

    {{ $redirect := getRedirect $service.SegmentLabels }}
    {{if $redirect }}
    [frontends."frontend-{{ $frontendName }}".redirect]
      entryPoint = "{{ $redirect.EntryPoint }}"
      regex = "{{ $redirect.Regex }}"
      replacement = "{{ $redirect.Replacement }}"
      permanent = {{ $redirect.Permanent }}
    {{end}}

    {{ $errorPages := getErrorPages $service.SegmentLabels }}
    {{if $errorPages }}
    [frontends."frontend-{{ $frontendName }}".errors]
      {{range $pageName, $page := $errorPages }}
      [frontends."frontend-{{ $frontendName }}".errors."{{ $pageName }}"]
        status = [{{range $page.Status }}
        "{{.}}",
        {{end}}]
        backend = "backend-{{ $page.Backend }}"
        query = "{{ $page.Query }}"
      {{end}}
    {{end}}

    {{ $rateLimit := getRateLimit $service.SegmentLabels }}
    {{if $rateLimit }}
    [frontends."frontend-{{ $frontendName }}".rateLimit]
      extractorFunc = "{{ $rateLimit.ExtractorFunc }}"
      [frontends."frontend-{{ $frontendName }}".rateLimit.rateSet]
        {{ range $limitName, $limit := $rateLimit.RateSet }}
        [frontends."frontend-{{ $frontendName }}".rateLimit.rateSet."{{ $limitName }}"]
          period = "{{ $limit.Period }}"
          average = {{ $limit.Average }}
          burst = {{ $limit.Burst }}
        {{end}}
    {{end}}

    {{ $headers := getHeaders $service.SegmentLabels }}
    {{if $headers }}
    [frontends."frontend-{{ $frontendName }}".headers]
      SSLRedirect = {{ $headers.SSLRedirect }}
      SSLTemporaryRedirect = {{ $headers.SSLTemporaryRedirect }}
      SSLHost = "{{ $headers.SSLHost }}"
      SSLForceHost = {{ $headers.SSLForceHost }}
      STSSeconds = {{ $headers.STSSeconds }}
      STSIncludeSubdomains = {{ $headers.STSIncludeSubdomains }}
      STSPreload = {{ $headers.STSPreload }}
      ForceSTSHeader = {{ $headers.ForceSTSHeader }}
      FrameDeny = {{ $headers.FrameDeny }}
      CustomFrameOptionsValue = "{{ $headers.CustomFrameOptionsValue }}"
      ContentTypeNosniff = {{ $headers.ContentTypeNosniff }}
      BrowserXSSFilter = {{ $headers.BrowserXSSFilter }}
      CustomBrowserXSSValue = "{{ $headers.CustomBrowserXSSValue }}"
      ContentSecurityPolicy = "{{ $headers.ContentSecurityPolicy }}"
      PublicKey = "{{ $headers.PublicKey }}"
      ReferrerPolicy = "{{ $headers.ReferrerPolicy }}"
      IsDevelopment = {{ $headers.IsDevelopment }}

      {{if $headers.AllowedHosts }}
      AllowedHosts = [{{range $headers.AllowedHosts }}
        "{{.}}",
        {{end}}]
      {{end}}

      {{if $headers.HostsProxyHeaders }}
      HostsProxyHeaders = [{{range $headers.HostsProxyHeaders }}
        "{{.}}",
        {{end}}]
      {{end}}

      {{if $headers.CustomRequestHeaders }}
      [frontends."frontend-{{ $frontendName }}".headers.customRequestHeaders]
        {{range $k, $v := $headers.CustomRequestHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}

      {{if $headers.CustomResponseHeaders }}
      [frontends."frontend-{{ $frontendName }}".headers.customResponseHeaders]
        {{range $k, $v := $headers.CustomResponseHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}

      {{if $headers.SSLProxyHeaders }}
      [frontends."frontend-{{ $frontendName }}".headers.SSLProxyHeaders]
        {{range $k, $v := $headers.SSLProxyHeaders }}
        {{$k}} = "{{$v}}"
        {{end}}
      {{end}}
    {{end}}

    [frontends."frontend-{{ $frontendName }}".routes."route-frontend-{{ $frontendName }}"]
      rule = "{{ getFrontendRule $service.Name $service.SegmentLabels }}"

{{end}}
`)

func templatesRancherTmplBytes() ([]byte, error) {
	return _templatesRancherTmpl, nil
}

func templatesRancherTmpl() (*asset, error) {
	bytes, err := templatesRancherTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/rancher.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/consul_catalog-v1.tmpl": templatesConsul_catalogV1Tmpl,
	"templates/consul_catalog.tmpl":    templatesConsul_catalogTmpl,
	"templates/docker-v1.tmpl":         templatesDockerV1Tmpl,
	"templates/docker.tmpl":            templatesDockerTmpl,
	"templates/ecs-v1.tmpl":            templatesEcsV1Tmpl,
	"templates/ecs.tmpl":               templatesEcsTmpl,
	"templates/eureka.tmpl":            templatesEurekaTmpl,
	"templates/kubernetes.tmpl":        templatesKubernetesTmpl,
	"templates/kv.tmpl":                templatesKvTmpl,
	"templates/marathon-v1.tmpl":       templatesMarathonV1Tmpl,
	"templates/marathon.tmpl":          templatesMarathonTmpl,
	"templates/mesos-v1.tmpl":          templatesMesosV1Tmpl,
	"templates/mesos.tmpl":             templatesMesosTmpl,
	"templates/notFound.tmpl":          templatesNotfoundTmpl,
	"templates/rancher-v1.tmpl":        templatesRancherV1Tmpl,
	"templates/rancher.tmpl":           templatesRancherTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"templates": {nil, map[string]*bintree{
		"consul_catalog-v1.tmpl": {templatesConsul_catalogV1Tmpl, map[string]*bintree{}},
		"consul_catalog.tmpl":    {templatesConsul_catalogTmpl, map[string]*bintree{}},
		"docker-v1.tmpl":         {templatesDockerV1Tmpl, map[string]*bintree{}},
		"docker.tmpl":            {templatesDockerTmpl, map[string]*bintree{}},
		"ecs-v1.tmpl":            {templatesEcsV1Tmpl, map[string]*bintree{}},
		"ecs.tmpl":               {templatesEcsTmpl, map[string]*bintree{}},
		"eureka.tmpl":            {templatesEurekaTmpl, map[string]*bintree{}},
		"kubernetes.tmpl":        {templatesKubernetesTmpl, map[string]*bintree{}},
		"kv.tmpl":                {templatesKvTmpl, map[string]*bintree{}},
		"marathon-v1.tmpl":       {templatesMarathonV1Tmpl, map[string]*bintree{}},
		"marathon.tmpl":          {templatesMarathonTmpl, map[string]*bintree{}},
		"mesos-v1.tmpl":          {templatesMesosV1Tmpl, map[string]*bintree{}},
		"mesos.tmpl":             {templatesMesosTmpl, map[string]*bintree{}},
		"notFound.tmpl":          {templatesNotfoundTmpl, map[string]*bintree{}},
		"rancher-v1.tmpl":        {templatesRancherV1Tmpl, map[string]*bintree{}},
		"rancher.tmpl":           {templatesRancherTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
