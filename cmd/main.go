package main

import (
	"flag"
	"k8s.io/klog/v2"
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

	// 打印集群基本信息，判断一下是否连通
	versionInfo, err := clientManager.InsecureClient().Discovery().ServerVersion()
	if err != nil {
		klog.Errorf("clientManager.InsecureClient().Discovery().ServerVersion().err:%v", err)
	}
	klog.Infof("[clientManager.init.success][version:%+v]", versionInfo.GitVersion, versionInfo.GoVersion, versionInfo.Compiler)
}
