package main
import (
	"flag"
	"github.com/vinkdong/monitoring-agent/server"
	"github.com/vinkdong/monitoring-agent/common"
       )

var(
        kubeconfig = flag.String("kubeconfig","","kubeconfig path file")
	addr = flag.String("addr",":9800","which addr to bind")
   )

func main()  {
	flag.Parse()
	common.InitKube(*kubeconfig)
	initVars()
	server.StartServer(*addr)
}

func initVars()  {
	common.Namespace = "monitoring"
	common.PrometheusRuleConfigmap = "prometheus-rules"
}