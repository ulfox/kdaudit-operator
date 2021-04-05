package controllers

import (
	"context"

	kdauditv1alpha1 "github.com/ulfox/kdaudit-operator/pkg/apis/kdaudit/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

type Controller struct {
	client kubernetes.Interface
}

func NewControllerDeployment(client kubernetes.Interface) *Controller {
	controllerDeployment := Controller{
		client: client,
	}
	return &controllerDeployment
}

func (c *Controller) Create(kdAudit *kdauditv1alpha1.KDAudit) (*appsv1.Deployment, error) {
	if err := c.preDeployment(kdAudit); err != nil {
		return nil, err
	}

	encodedData, err := getEncodedData(kdAudit)
	if err != nil {
		return nil, err
	}

	newDep := Deployment{
		data: encodedData,
	}
	newDeployment := newDep.newDeployment(kdAudit)

	deployment, err := c.client.AppsV1().Deployments(kdAudit.Namespace).Create(context.TODO(), newDeployment, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return deployment, nil
}

func (c *Controller) Update(kdAudit *kdauditv1alpha1.KDAudit) (*appsv1.Deployment, error) {
	if err := c.preDeployment(kdAudit); err != nil {
		return nil, err
	}

	encodedData, err := getEncodedData(kdAudit)
	if err != nil {
		return nil, err
	}

	newDep := Deployment{
		data: encodedData,
	}
	newDeployment := newDep.newDeployment(kdAudit)

	deployment, err := c.client.AppsV1().Deployments(kdAudit.Namespace).Update(context.TODO(), newDeployment, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	return deployment, nil
}

func (c *Controller) preDeployment(kdAudit *kdauditv1alpha1.KDAudit) error {
	_, err := c.client.CoreV1().ServiceAccounts(kdAudit.Namespace).Get(context.TODO(), kdAudit.Spec.DeploymentName, metav1.GetOptions{})
	if err != nil {
		_, err := c.client.CoreV1().ServiceAccounts(kdAudit.Namespace).Create(context.TODO(), newServiceAccount(kdAudit), metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}
	_, err = c.client.RbacV1().ClusterRoles().Get(context.TODO(), kdAudit.Spec.DeploymentName, metav1.GetOptions{})
	if err != nil {
		_, err := c.client.RbacV1().ClusterRoles().Create(context.TODO(), newClusterRole(kdAudit), metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	_, err = c.client.RbacV1().ClusterRoleBindings().Get(context.TODO(), kdAudit.Spec.DeploymentName, metav1.GetOptions{})
	if err != nil {
		_, err := c.client.RbacV1().ClusterRoleBindings().Create(context.TODO(), newClusterRoleBinding(kdAudit), metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Controller) GetObjDiff(depActive *appsv1.Deployment, kdAudit *kdauditv1alpha1.KDAudit) (bool, error) {
	encodedData, err := getEncodedData(kdAudit)
	if err != nil {
		return false, err
	}

	newDep := Deployment{
		data: encodedData,
	}
	depActual := newDep.newDeployment(kdAudit)

	if len(depActive.Spec.Template.Spec.Containers[0].Env) != len(depActual.Spec.Template.Spec.Containers[0].Env) {
		klog.Infof(
			"Found different: Deployment %s has different number env keys from actual. Issuing rollout",
			kdAudit.Spec.DeploymentName,
		)
		return true, nil
	}

	for v := range depActive.Spec.Template.Spec.Containers[0].Env {
		if depActive.Spec.Template.Spec.Containers[0].Env[v] != depActual.Spec.Template.Spec.Containers[0].Env[v] {
			klog.Infof(
				"Found different: Deployment %s has different active env key %s from actual. Issuing rollout",
				kdAudit.Spec.DeploymentName,
				depActive.Spec.Template.Spec.Containers[0].Env[v].Name,
			)
			return true, nil
		}
	}
	if depActive.Spec.Template.Spec.Containers[0].Image != depActual.Spec.Template.Spec.Containers[0].Image {
		klog.InfoS("Found different image for deployment %s. Issuing rollout", kdAudit.Spec.DeploymentName)
		return true, nil
	}
	return false, nil
}
