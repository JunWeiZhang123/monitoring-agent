package common

import (
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	"net/http"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
)

func InitKube(kubeconfig string)  {
	config, err := clientcmd.BuildConfigFromFlags("",kubeconfig)
	if err != nil{
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	ClientSet = clientset

	selector := "annotations=choerodon.io/logs-parser"

	deployments :=  clientset.AppsV1beta1().Deployments("")
	d , e := deployments.List(metav1.ListOptions{FieldSelector:selector})
	fmt.Println(e)
	for _,v := range d.Items{
		fmt.Println(v.Name)
	}


}

func apiHandler(resp http.ResponseWriter, req *http.Request) {

	fmt.Println(req.Host)
}
