package client

import (
	"github.com/emicklei/go-restful/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
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

func (self *ClientManager) InsecureClient() kubernetes.Interface {
	return self.insecureClient
}

func (self *ClientManager) GetK8sClientByRequest(req *restful.Request) (kubernetes.Interface, error) {
	if req == nil {
		return nil, errors.NewBadRequest("request can not be nil")
	}
	return self.secureClient(req)
}

func (cm *ClientManager) secureClient(req *restful.Request) (kubernetes.Interface, error) {
	// 首先要根据req 解析k8sToken
	k8sToken := req.HeaderParameter(cm.k8sTokenKey)
	if k8sToken == "" {
		msgStr := "req.HeaderParameter.find.k8sToken.empty"
		klog.Errorf(msgStr)
		return nil, errors.NewBadRequest(msgStr)
	}
}

//根据http的request来拿到k8s_token,初始化一个kubernetes.interface
