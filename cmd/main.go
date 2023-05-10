package main

import (
	"flag"
	"k8s_management/pkg/client"
)

var (
	kubeconfig  *string
	httpAddr    string
	k8sTokenKey *string
)

func main() {
	// 初始化clientManager

	kubeconfig = flag.String("kubeconfig", "kubeconfig", "absolute path to the kubeconfig file")
	k8sTokenKey = flag.String("k8sTokenKey", "k8sToken", "absolute path to the kubeconfig file")
	flag.StringVar(&httpAddr, "http.addr", "0.0.0.0:8087", "The http addr")
	flag.Parse()
	// 初始化这个clientManager
	clientManager := client.NewClientManager(*kubeconfig, *k8sTokenKey)
	clientManager.InsecureClient().Discovery()
}
