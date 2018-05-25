package v1

import (
	"k8s.io/client-go/kubernetes"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/vinkdong/gox/log"
	"github.com/bitly/go-simplejson"
	"strings"
	"github.com/vinkdong/monitoring-agent/common"
	"k8s.io/api/core/v1"
	"gopkg.in/yaml.v2"
)

type Prometheus struct {
	ClientSet *kubernetes.Clientset
}

type Category struct {
	Groups []Group `yaml:groups`
}
type Group struct {
	Name string `yaml:name`
	Rules []Rule `yaml:rules`
}

type Rule struct {
	Alert string `yaml:alert`
	Record string `yaml:record`
	Expr  string `yaml:expr`
	For   string `yaml:for`
	Annotations map[string]string `yaml:annotations`
}

func (p *Prometheus) getConfigmap() (*v1.ConfigMap,error) {
	namespace := common.Namespace
	ruleName := common.PrometheusRuleConfigmap
	return p.ClientSet.CoreV1().ConfigMaps(namespace).Get(ruleName, meta_v1.GetOptions{})
}

func (p *Prometheus) GetCategories() *simplejson.Json {
	configmap, err := p.getConfigmap()
	if err != nil {
		log.Error(err)
	}
	keys := make([]string, len(configmap.Data))
	i := 0
	for k := range configmap.Data {
		keys[i] = strings.Replace(k, ".rules", "", 1)
		i ++
	}
	js := common.BuildCommonJson()
	js.Set("categories",keys)
	return js
}

func (p *Prometheus) GetGroups(category string) *simplejson.Json {
	configmap, err := p.getConfigmap()
	if err != nil{
		log.Error(err)
	}
	js := common.BuildCommonJson()
	if value, ok := configmap.Data[category+".rules"];ok{
		c := &Category{}
		yaml.Unmarshal([]byte(value),c)
		groups := make([]string,len(c.Groups))
		for i:=0;i<len(c.Groups);i++{
			groups[i] = c.Groups[i].Name
		}
		js.Set("groups",groups)
		js.Set("category",category)
	}else {
		js.Set("success",false)
	}
	return js
}

func (p *Prometheus) GetRules(category,group string) *simplejson.Json {
	_, err := p.ClientSet.CoreV1().ConfigMaps(common.Namespace).Get(common.PrometheusRuleConfigmap, meta_v1.GetOptions{})
	if err != nil {
		log.Error(err)
	}
	js := simplejson.New()
	//searchName := name+".rules"
	//
	//if value, ok := configmap.Data[searchName];ok{
	//	js.Set("success",true)
	//	js.Set("value",value)
	//}
	//js.Set("version","v1")
	return js
}

func (p *Prometheus) CreateRule( ) {
}

func (p *Prometheus) UpdateRule() {

}
