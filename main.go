package main

import (
	"fmt"

	"git/inspursoft/k8sclientdemo/service"
	"git/inspursoft/k8sclientdemo/common/model"

)

const (
	nodeName  = "192.168.154.5"
	groupName = "testClientGO"

	k8sMasterIP  = "192.168.154.3"
)

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

func main() {
	fmt.Println("##############################################################")
	fmt.Println("testing node:\n")
	testNode()

	fmt.Println("##############################################################\n\n\n\n\n")

	fmt.Println("testing statefulset:\n")
	testDeployStatefulSet()
}

func testNode() {
	fmt.Println("\n*************get node:*************")
	node, err := service.GetNode(nodeName)
	if err != nil {
		fmt.Println("Fail to get node")
	} else {
		fmt.Println("get node:\n", node)
	}


	fmt.Println("\n*************Suspend node:*************")
	unschedulable, err := service.SuspendNode(nodeName)
	if err != nil {
		fmt.Println("Fail to Suspend node")
	} else {
		fmt.Println("unschedulable:", unschedulable)
	}


	fmt.Println("\n*************Resume node:*************")
	unschedulable, err = service.ResumeNode(nodeName)
	if err != nil {
		fmt.Println("Fail to Resume node")
	} else {
		fmt.Println("unschedulable:", unschedulable)
	}


	fmt.Println("\n*************Get node list:*************")
	nodeList := service.GetNodeList()
	fmt.Println(nodeList)


	fmt.Println("\n*************Add Node To Group:*************")
	err = service.AddNodeToGroup(nodeName, groupName)
	if err != nil {
		fmt.Printf("Fail to add node '%s' to group '%s'\n", nodeName, groupName)
	} else {
		fmt.Printf("Add node '%s' to group '%s' successfully\n", nodeName, groupName)
	}


	fmt.Println("\n*************Get Group Of Node:*************")
	nodeGroupNames, err := service.GetGroupOfNode(nodeName)
	if err != nil {
		fmt.Println("Fail to get node groups\n")
	} else {
		fmt.Printf("node groups:\n", nodeGroupNames)
	}


	fmt.Println("\n*************Remove Node From Group:*************")
	err = service.RemoveNodeFromGroup(nodeName, groupName)
	if err != nil {
		fmt.Printf("Fail to Remove node '%s' from group '%s'\n", nodeName, groupName)
	} else {
		fmt.Printf("Remove node '%s' from group '%s' successfully\n", nodeName, groupName)
	}

	fmt.Println("\n*************Get Group Of Node:*************")
	nodeGroupNames, err = service.GetGroupOfNode(nodeName)
	if err != nil {
		fmt.Println("Fail to get node groups\n")
	} else {
		fmt.Printf("node groups:\n", nodeGroupNames)
	}


	fmt.Println("\n*************Get Nodes Available Resources:*************")
	nodeResources, err := service.GetNodesAvailableResources()
	fmt.Println(nodeResources, err)
}


// TODO: unit test case later
// TestDeployStatefulSet
func testDeployStatefulSet() {
	masterIP := k8sMasterIP
	fmt.Printf("KUBE_MASTER_IP %s\n", masterIP)

	registryURI := k8sMasterIP
	fmt.Printf("REGISTRY_URI %s\n", registryURI)

	deployStatefulSetInfo, err := service.DeployStatefulSet(&configStatefulSet, registryURI)
	
	if err != nil {
		fmt.Println("Failed, err when create test StatefulSet")
		return
	} 
	
	fmt.Printf("\nCreated statefulset %v %v\n", deployStatefulSetInfo.Service, deployStatefulSetInfo.StatefulSet)

	//clean test
	cleanStatefulSet(configStatefulSet.ProjectName, configStatefulSet.ServiceName)
}

// clean test
func cleanStatefulSet(projectName string, serviceName string) {
	fmt.Printf("cleanStatefulSet %s %s", projectName, serviceName)
	err := service.StopStatefulSetK8s(&model.ServiceStatus{
		Name:        serviceName,
		ProjectName: projectName,
	})
	if err != nil {
		fmt.Printf("cleanStatefulSet failed %v", err)
		return
	}
	fmt.Printf("cleaned StatefulSet")
}