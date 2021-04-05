package controllers

import (
	kdauditv1alpha1 "github.com/ulfox/kdaudit-operator/pkg/apis/kdaudit/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Deployment struct {
	data string
}

func (d *Deployment) newDeployment(kdAudit *kdauditv1alpha1.KDAudit) *appsv1.Deployment {
	labels := setLabels(kdAudit)
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      kdAudit.Spec.DeploymentName,
			Namespace: kdAudit.Namespace,
			// sets the appropriate OwnerReferences on the resource
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(
					kdAudit,
					kdauditv1alpha1.SchemeGroupVersion.WithKind("KDAudit"),
				),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: kdAudit.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: kdAudit.Spec.DeploymentName,
					Containers: []corev1.Container{
						{
							Name:            kdAudit.Spec.DeploymentName,
							Image:           kdAudit.Spec.Image,
							ImagePullPolicy: "Never",
							Env: []corev1.EnvVar{
								{
									Name:  "KDAUDIT_SERVICE",
									Value: kdAudit.Spec.Service,
								},
								{
									Name:  "SLACK_WEBHOOK",
									Value: kdAudit.Spec.SlackWebhook,
								},
								{
									Name:  "NAMESPACE_WATCHER_ENABLED",
									Value: kdAudit.Spec.NamespaceWatcher.Enabled,
								},
								{
									Name:  "NAMESPACE_WATCHER_CONFIGMAP",
									Value: d.data,
								},
							},
						},
					},
				},
			},
		},
	}
}

func newServiceAccount(kdAudit *kdauditv1alpha1.KDAudit) *corev1.ServiceAccount {
	return &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      kdAudit.Spec.DeploymentName,
			Namespace: kdAudit.Namespace,
			Labels:    setLabels(kdAudit),
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(kdAudit, kdauditv1alpha1.SchemeGroupVersion.WithKind("KDAudit")),
			},
		},
	}
}

func newClusterRole(kdAudit *kdauditv1alpha1.KDAudit) *v1.ClusterRole {
	return &v1.ClusterRole{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRole",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      kdAudit.Spec.DeploymentName,
			Namespace: kdAudit.Namespace,
			Labels:    setLabels(kdAudit),
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(kdAudit, kdauditv1alpha1.SchemeGroupVersion.WithKind("KDAudit")),
			},
		},
		Rules: []v1.PolicyRule{
			{
				APIGroups: []string{"", "apps"},
				Resources: []string{"*"},
				Verbs:     []string{"*"},
			},
			{
				APIGroups: []string{"kdaudit.k8s.io", "kdaudits.kdaudit.k8s.io"},
				Resources: []string{"*"},
				Verbs:     []string{"*"},
			},
		},
	}
}

func newClusterRoleBinding(kdAudit *kdauditv1alpha1.KDAudit) *v1.ClusterRoleBinding {
	return &v1.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRoleBinding",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      kdAudit.Spec.DeploymentName,
			Namespace: kdAudit.Namespace,
			Labels:    setLabels(kdAudit),
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(kdAudit, kdauditv1alpha1.SchemeGroupVersion.WithKind("KDAudit")),
			},
		},
		Subjects: []v1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      kdAudit.Spec.DeploymentName,
				Namespace: kdAudit.Namespace,
			},
		},
		RoleRef: v1.RoleRef{
			Kind:     "ClusterRole",
			APIGroup: "rbac.authorization.k8s.io",
			Name:     kdAudit.Spec.DeploymentName,
		},
	}
}
