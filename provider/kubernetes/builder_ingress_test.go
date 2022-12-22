package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func buildIngress(opts ...func(*networkingv1.Ingress)) *networkingv1.Ingress {
	i := &networkingv1.Ingress{}
	for _, opt := range opts {
		opt(i)
	}
	return i
}

func iNamespace(value string) func(*networkingv1.Ingress) {
	return func(i *networkingv1.Ingress) {
		i.Namespace = value
	}
}

func iAnnotation(name string, value string) func(*networkingv1.Ingress) {
	return func(i *networkingv1.Ingress) {
		if i.Annotations == nil {
			i.Annotations = make(map[string]string)
		}
		i.Annotations[name] = value
	}
}

func iRules(opts ...func(*networkingv1.IngressSpec)) func(*networkingv1.Ingress) {
	return func(i *networkingv1.Ingress) {
		s := &networkingv1.IngressSpec{}
		for _, opt := range opts {
			opt(s)
		}
		i.Spec = *s
	}
}

func iSpecBackends(opts ...func(*networkingv1.IngressSpec)) func(*networkingv1.Ingress) {
	return func(i *networkingv1.Ingress) {
		s := &networkingv1.IngressSpec{}
		for _, opt := range opts {
			opt(s)
		}
		i.Spec = *s
	}
}

func iSpecBackend(opts ...func(*networkingv1.IngressBackend)) func(*networkingv1.IngressSpec) {
	return func(s *networkingv1.IngressSpec) {
		p := &networkingv1.IngressBackend{}
		for _, opt := range opts {
			opt(p)
		}
		s.DefaultBackend = p
	}
}

func iIngressBackend(name string, port networkingv1.ServiceBackendPort) func(*networkingv1.IngressBackend) {
	return func(p *networkingv1.IngressBackend) {
		p.Service = new(networkingv1.IngressServiceBackend)
		p.Service.Name = name
		p.Service.Port = port
	}
}

func iRule(opts ...func(*networkingv1.IngressRule)) func(*networkingv1.IngressSpec) {
	return func(spec *networkingv1.IngressSpec) {
		r := &networkingv1.IngressRule{}
		for _, opt := range opts {
			opt(r)
		}
		spec.Rules = append(spec.Rules, *r)
	}
}

func iHost(name string) func(*networkingv1.IngressRule) {
	return func(rule *networkingv1.IngressRule) {
		rule.Host = name
	}
}

func iPaths(opts ...func(*networkingv1.HTTPIngressRuleValue)) func(*networkingv1.IngressRule) {
	return func(rule *networkingv1.IngressRule) {
		rule.HTTP = &networkingv1.HTTPIngressRuleValue{}
		for _, opt := range opts {
			opt(rule.HTTP)
		}
	}
}

func onePath(opts ...func(*networkingv1.HTTPIngressPath)) func(*networkingv1.HTTPIngressRuleValue) {
	return func(irv *networkingv1.HTTPIngressRuleValue) {
		p := &networkingv1.HTTPIngressPath{}
		for _, opt := range opts {
			opt(p)
		}
		irv.Paths = append(irv.Paths, *p)
	}
}

func iPath(name string) func(*networkingv1.HTTPIngressPath) {
	return func(p *networkingv1.HTTPIngressPath) {
		p.Path = name
	}
}

func iBackend(name string, port networkingv1.ServiceBackendPort) func(*networkingv1.HTTPIngressPath) {
	return func(p *networkingv1.HTTPIngressPath) {
		p.Backend = networkingv1.IngressBackend{
			Service: &networkingv1.IngressServiceBackend{
				Name: name,
				Port: port,
			},
		}
	}
}

func iTLSes(opts ...func(*networkingv1.IngressTLS)) func(*networkingv1.Ingress) {
	return func(i *networkingv1.Ingress) {
		for _, opt := range opts {
			iTLS := networkingv1.IngressTLS{}
			opt(&iTLS)
			i.Spec.TLS = append(i.Spec.TLS, iTLS)
		}
	}
}

func iTLS(secret string, hosts ...string) func(*networkingv1.IngressTLS) {
	return func(i *networkingv1.IngressTLS) {
		i.SecretName = secret
		i.Hosts = hosts
	}
}

// Test

func TestBuildIngress(t *testing.T) {
	i := buildIngress(
		iNamespace("testing"),
		iRules(
			iRule(iHost("foo"), iPaths(
				onePath(iPath("/bar"), iBackend("service1", networkingv1.ServiceBackendPort{Number: 80})),
				onePath(iPath("/namedthing"), iBackend("service4", networkingv1.ServiceBackendPort{Name: "https"}))),
			),
			iRule(iHost("bar"), iPaths(
				onePath(iBackend("service3", networkingv1.ServiceBackendPort{Name: "https"})),
				onePath(iBackend("service2", networkingv1.ServiceBackendPort{Number: 802})),
			),
			),
		),
		iTLSes(
			iTLS("tls-secret", "foo"),
		),
	)

	assert.EqualValues(t, sampleIngress(), i)
}

func sampleIngress() *networkingv1.Ingress {
	return &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "testing",
		},
		Spec: networkingv1.IngressSpec{
			Rules: []networkingv1.IngressRule{
				{
					Host: "foo",
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path: "/bar",
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: "service1",
											Port: networkingv1.ServiceBackendPort{
												Number: 80,
											},
										},
									},
								},
								{
									Path: "/namedthing",
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: "service4",
											Port: networkingv1.ServiceBackendPort{
												Name: "https",
											},
										},
									},
								},
							},
						},
					},
				},
				{
					Host: "bar",
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: "service3",
											Port: networkingv1.ServiceBackendPort{
												Name: "https",
											},
										},
									},
								},
								{
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: "service2",
											Port: networkingv1.ServiceBackendPort{
												Number: 802,
											},
										},
									},
								},
							},
						},
					},
				},
			},
			TLS: []networkingv1.IngressTLS{
				{
					Hosts:      []string{"foo"},
					SecretName: "tls-secret",
				},
			},
		},
	}
}
