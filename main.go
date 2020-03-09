package main

import (
	"flag"
	"fmt"
	"git/inspursoft/k8sclientdemo/common/model"
	"git/inspursoft/k8sclientdemo/common/k8sassist"
)

const (
	kubeConfig = "./conf/k8s.conf"
	version    = "v1.8"
)

var namespace = model.Namespace{
	ObjectMeta: model.ObjectMeta{
		Name:   "k8sassist",
		Labels: map[string]string{"app": "namespace"},
	},
}

func main() {
	var config k8sassist.K8sAssistConfig
	config.KubeConfigPath = flag.String("kubeconfig", "./conf/config", "path to the kubeconfig file")
	flag.Parse()
	k8sclient := k8sassist.NewK8sAssistClient(&config)
	
	ns, err := k8sclient.AppV1().Namespace().Create(&namespace)
	if err != nil {
		fmt.Println("update namespace error: ", err.Error())
		return
	}
	fmt.Printf("namespace:%+v", ns)
}
