package monitor
import (
	"io/ioutil"
	"strconv"
	"encoding/json"
	"gopkg.in/gomail.v1" 
	"time"
	"fmt"
	"net/http"
	"github.com/kylelemons/go-gypsy/yaml"
	// prom_v1 "github.com/vinkdong/monitoring-agent/prometheus/v1"
)

type es struct {
    Status string
}
func GoMail(body string) error {  
config, err := yaml.ReadFile("conf.yaml")
    if err != nil {
        fmt.Println(err)
    }
    from, err :=config.Get("from")
    to, err :=config.Get("to")     
    host, err :=config.Get("host")        
    password, err :=config.Get("password")
    portString, err :=config.Get("port")
    port,err:=strconv.Atoi(portString)  
    msg := gomail.NewMessage()  
    msg.SetHeader("From", from)       
    msg.SetHeader("To", to)  
    msg.SetHeader("Subject", "Alert")  
    msg.SetBody("text/html", body)  	
    mailer := gomail.NewMailer(host, from, password,port)
    if err := mailer.Send(msg); err != nil {  
        return err
    }  
    return nil  
}  


func addMonitorRoute()  {
       config, err := yaml.ReadFile("conf.yaml")
       ticker := time.NewTicker(time.Second * 20)
       client := &http.Client{}
	//生成要访问的url
	prometheus_url, err :=config.Get("prometheus_url")
	elasticSearch_url, err :=config.Get("elasticSearch_url")
	//提交请求
	elasticSearch_request, err := http.NewRequest("GET", elasticSearch_url, nil)	
	prometheus_request, err := http.NewRequest("GET", prometheus_url, nil)	
	if err != nil {
		panic(err)
	}	
        go func() {
	    for t := range ticker.C {
	        prometheus_response, _ := client.Do(prometheus_request)
		elasticSearch_response, _ := client.Do(elasticSearch_request)
		if(elasticSearch_response==nil || elasticSearch_response.Status!="200 ok"){
			GoMail("elasticSearch is failed")
			fmt.Println("elasticSearch is failed")
		  } else {
			var es es
			body, err := ioutil.ReadAll(elasticSearch_response.Body)
			json.Unmarshal(body,&es)
			if err != nil {
				panic(err)
			}	
                if(es.Status!="green"){
			GoMail("elasticSearch is failed")
			fmt.Println("elasticSearch is failed")
			}
		  }
		if(prometheus_response==nil || prometheus_response.Status!="200 ok"){
			GoMail("prometheus was closed")
			fmt.Println("prometheus is failed")
		  }
		fmt.Println(t)	  
	}
}()
}


