package service_test

import (
	"git/inspursoft/k8sclientdemo/apiserver/service"
	"git/inspursoft/k8sclientdemo/common/model"
	"testing"

	// "git/inspursoft/k8sclientdemo/common/utils"

	"github.com/astaxie/beego/logs"
	"github.com/stretchr/testify/assert"
)

const {
	k8sMasterIP  = "192.168.154.3"
}

var path = "./"
var configServiceStep = model.ConfigServiceStep{
	ProjectID:   1,
	Instance:    1,
	ServiceName: "testService",
	ContainerList: []model.Container{
		model.Container{
			Name: "testService",
			Image: model.ImageIndex{
				ImageName:   "nginx",
				ImageTag:    "1.15.7-alpine",
				ProjectName: "library",
			},
		},
	},
	ExternalServiceList: []model.ExternalService{
		model.ExternalService{
			ContainerName: "testService",
			NodeConfig: model.NodeType{
				TargetPort: 80,
				NodePort:   32080,
			},
		},
	},
}

var configStatefulSet = model.ConfigServiceStep{
	ProjectID:   1,
	ProjectName: "library",
	Instance:    1,
	ServiceName: "unitteststatefulset001",
	ServiceType: model.ServiceTypeStatefulSet,
	ClusterIP:   "None",
	ContainerList: []model.Container{
		model.Container{
			Name: "nginx",
			Image: model.ImageIndex{
				ImageName:   "nginx",
				ImageTag:    "1.15.7-alpine",
				ProjectName: "library",
			},
		},
	},
	ExternalServiceList: []model.ExternalService{
		model.ExternalService{
			ContainerName: "nginx",
			NodeConfig: model.NodeType{
				TargetPort: 80,
				Port:       80,
			},
		},
	},
}

var statufulsetName = "unitteststatefulset001"

// TODO: unit test case later
// TestDeployStatefulSet
func TestDeployStatefulSet(t *testing.T) {
	assert := assert.New(t)
	t.Log("Check KubeMaster")
	// masterIP := utils.GetStringValue("KUBE_MASTER_IP")
	masterIP := k8sMasterIP
	logs.Info("KUBE_MASTER_IP %s", masterIP)

	// registryURI := utils.GetStringValue("REGISTRY_BASE_URI")
	registryURI := k8sMasterIP
	logs.Info("REGISTRY_URI %s", registryURI)

	deployStatefulSetInfo, err := service.DeployStatefulSet(&configStatefulSet, registryURI)
	assert.Nil(err, "Failed, err when create test StatefulSet")
	assert.Equal(statufulsetName, deployStatefulSetInfo.Service.Name, "Failed to create StatefulSet")
	logs.Info("Created statefulset %v %v", deployStatefulSetInfo.Service, deployStatefulSetInfo.StatefulSet)

	//clean test
	t.Log("Clean TestDeployStatefulSet")
	cleanStatefulSet(configStatefulSet.ProjectName, configStatefulSet.ServiceName)
	t.Log("Tested TestDeployStatefulSet")
}

// clean test
func cleanStatefulSet(projectName string, serviceName string) {
	logs.Info("cleanStatefulSet %s %s", projectName, serviceName)
	err := service.StopStatefulSetK8s(&model.ServiceStatus{
		Name:        serviceName,
		ProjectName: projectName,
	})
	if err != nil {
		logs.Info("cleanStatefulSet failed %v", err)
		return
	}
	logs.Info("cleaned StatefulSet")
}
