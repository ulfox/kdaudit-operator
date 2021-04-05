package utils

import (
	"context"
	"errors"
	"strconv"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

// KDAuditCFG for storing kdaudit configMap
// The configMap is used to check which kdaudit
// services will be loaded
type KDAuditCFG struct {
	client      *kubernetes.Clientset
	configMap   map[string]string
	Threadiness int
}

// NewKDAuditCFG for creating a new KDAuditCFG struct
func NewKDAuditCFG(client *kubernetes.Clientset) KDAuditCFG {
	return KDAuditCFG{
		client: client,
	}
}

// ReadConfig for reading kdaudit configMap
func (c *KDAuditCFG) ReadConfig() {
	for {
		if err := c.readUntil(); err != nil {
			klog.Error(err)
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
}

func (c *KDAuditCFG) readUntil() error {
	if configMap, err := c.client.CoreV1().ConfigMaps("kdaudit").Get(
		context.TODO(),
		"kdaudit-operator",
		metav1.GetOptions{},
	); err != nil {
		return errors.New("configmap kdaudit-operator could not be found. Retrying...")
	} else {
		c.configMap = configMap.Data
		return c.getOpts()
	}
}

func (c *KDAuditCFG) getOpts() error {
	if c.configMap["workerThreads"] == "" {
		c.Threadiness = 1
	} else {
		threadiness, err := strconv.Atoi(c.configMap["workerThreads"])
		if err != nil {
			return err
		}
		c.Threadiness = threadiness
	}
	return nil
}
