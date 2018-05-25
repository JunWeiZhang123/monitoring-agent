package common

import "k8s.io/client-go/kubernetes"

var (
	Namespace string
	PrometheusRuleConfigmap      string
	ClientSet *kubernetes.Clientset
)
