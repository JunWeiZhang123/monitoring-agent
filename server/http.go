package server

import (
	"net/http"
	"k8s.io/client-go/kubernetes"
	"github.com/emicklei/go-restful"
    "github.com/vinkdong/monitoring-agent/monitor"
	prom_v1 "github.com/vinkdong/monitoring-agent/prometheus/v1"
)

var Clientset *kubernetes.Clientset

func initServer(addr string) {  
	monitor.InitMonitor()
	addRoute()	
	server := &http.Server{Addr:addr,Handler:restful.DefaultContainer}
	server.ListenAndServe()
}

func addRoute()  {
	addPrometheusRoute()
}

/*
GET /rules/categories              *** get all rules file
GET /rules/category/groups/        *** get groups of one category
GET /rules/category/group/rules    *** get rules of one group

PUT /rules/category                *** create a category
PUT /rules/category/group          *** create a group of one category
PUT /rules/category/group/rule     *** create a rule of one group
*/
func addPrometheusRoute()  {
	ws := new(restful.WebService)
	ws.Path("/prometheus/rules").Consumes(restful.MIME_JSON,restful.MIME_XML).Produces(restful.MIME_JSON,restful.MIME_XML)
	ws.Route(ws.GET("categories").
		To(prom_v1.RouteGetCategories).
			Doc("get all categories"))
	ws.Route(ws.GET("{category-id}/groups").
		To(prom_v1.RouteGetGroups).
			Doc("get groups of one category").
				Param(ws.PathParameter("category-id","identifier of the category").
					DataType("string")))
	restful.Add(ws)
}

func StartServer(addr string)  {
	initServer(addr)
}