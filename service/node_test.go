package service_test

import (
	"fmt"

	"testing"
)

const (
	nodeName  = "10.110.18.71"
	groupName = "testClientGO"
)

func TestGetNode(t *testing.T) {
	node, err := GetNode(nodeName)
	fmt.Println(node, err)

	a, b := SuspendNode(nodeName)
	fmt.Println(a, b)

	a, b = ResumeNode(nodeName)
	fmt.Println(a, b)

	nodeList := GetNodeList()
	fmt.Println(nodeList)

	err = AddNodeToGroup(nodeName, groupName)
	if err != nil {
		fmt.Printf("Add node '%s' to group '%s' successfully\n", nodeName, groupName)
	} else {
		fmt.Printf("Fail to add node '%s' to group '%s'\n", nodeName, groupName)
	}

	nodeGroupNames, err := GetGroupOfNode(nodeName)
	fmt.Println(nodeGroupNames, err)

	success := RemoveNodeFromGroup(nodeName, groupName)
	fmt.Println(success)

	nodeResources, err := GetNodesAvailableResources()
	fmt.Println(nodeResources, err)

	a, b = NodeExists(nodeName)
	fmt.Println(a, b)
}
