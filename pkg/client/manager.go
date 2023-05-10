package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type ClientManager struct {
	kubeconfigPath string               // kubeconfig 连接哪个集群
	insecureConfig *rest.Config         // 平台自用的rest.config
	insecureClient kubernetes.Interface //平台自用的client-interface
	k8sTokenKey    string               // 可配置的key的名字
}

func NewClientManager(kubeconfigPath, k8sTokenKey string) *ClientManager {
	cm := &ClientManager{
		kubeconfigPath: kubeconfigPath,
		insecureConfig: nil,
		insecureClient: nil,
		k8sTokenKey:    k8sTokenKey,
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	cm.insecureConfig = config
	cm.insecureClient = clientset

	return cm
}

func (selt *ClientManager) InsecureClient() kubernetes.Interface(
	return self.insecureClient
	)