package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	naclient "nodeagent/clients/agent/clientset/versioned"
	"testing"
)

func TestNodeActionClient (t *testing.T){
	cfg, err := clientcmd.BuildConfigFromFlags("", "~/.kube/config")
	if err != nil {
		panic(err)
	}
	client := naclient.NewForConfigOrDie(cfg)
	nodeActions, err := client.AgentV1().NodeActions("").List(context.Background(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(nodeActions)
}
