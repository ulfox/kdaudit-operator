package utils

import (
	"os"
	"strings"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

// clientAuth struct used configure kubernetes.Clientset
type clientAuth struct {
	client  *kubernetes.Clientset
	kubeCFG string
	config  *rest.Config
}

// getInClusterConfig for getting kube-apiserver connection info
// this option uses the pod's service account
func (c *clientAuth) getInClusterConfig() {
	inClusterAccess, err := rest.InClusterConfig()
	if err != nil {
		klog.Fatal(err)
	}

	c.config = inClusterAccess
}

// getLocalKubeConfig for loading local kubernetes client config
func (c *clientAuth) getLocalKubeConfig(localCFG string) {
	if strings.HasPrefix(localCFG, "~") {
		homePath := os.ExpandEnv("$HOME")
		c.kubeCFG = homePath + localCFG[1:]
	} else {
		c.kubeCFG = localCFG
	}

	if _, err := os.Stat(c.kubeCFG); err != nil {
		klog.Fatal(err)
	}

	cfg, err := clientcmd.BuildConfigFromFlags("", c.kubeCFG)
	if err != nil {
		klog.Fatal(err)
	}

	c.config = cfg
}

// setClient for setting up a client to work with kube-apiserver
// this will use either getLocalKubeConfig or getInClusterConfig
// configuration
func (c *clientAuth) setClient() {
	client, err := kubernetes.NewForConfig(c.config)
	if err != nil {
		klog.Fatal(err)
	}
	c.client = client
}

// NewClientAuth for creating a new kubernetes client using either
// inCluster for connecting withing kubernetes namespace by using a service account
// or with a local kubernetes client configuration
func NewClientAuth(c, l string) (*kubernetes.Clientset, *rest.Config) {
	clientAuth := clientAuth{}
	if c == "inCluster" {
		clientAuth.getInClusterConfig()
	} else {
		clientAuth.getLocalKubeConfig(l)
	}

	clientAuth.setClient()

	return clientAuth.client, clientAuth.config
}
