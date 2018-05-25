package v1

import (
	"github.com/emicklei/go-restful"
	"github.com/vinkdong/monitoring-agent/common"
	"github.com/vinkdong/gox/log"
)

func RouteGetCategories(request *restful.Request, response *restful.Response) {
	p := Prometheus{ClientSet: common.ClientSet }
	js := p.GetCategories()
	b, err := js.MarshalJSON()
	if err != nil{
		log.Error(err)
	}
	response.Header().Set("content-type","application/json; charset=UTF-8")
	response.Write(b)
}

func RouteGetGroups(request *restful.Request, response *restful.Response) {
	category := request.PathParameter("category-id")
	p := Prometheus{ClientSet: common.ClientSet }
	js := p.GetGroups(category)
	b, err := js.MarshalJSON()
	if err != nil{
		log.Error(err)
	}
	response.Header().Set("content-type","application/json; charset=UTF-8")
	response.Write(b)
}

