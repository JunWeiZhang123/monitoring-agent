package v1

import (
	"testing"
	"github.com/vinkdong/monitoring-agent/common"
)

func TestRule_GetGroup(t *testing.T) {
	common.InitKube("//Users/vink/go/src/github.com/vinkdong/monitoring-agent/config.yaml")
	common.Namespace = "monitoring"
	common.PrometheusRuleConfigmap ="prometheus-rules"
	p:= Prometheus{ClientSet:common.ClientSet}
	p.GetGroups("host")
}