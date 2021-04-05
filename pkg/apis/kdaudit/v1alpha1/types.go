package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KDAudit is a specification for a KDAudit resource
type KDAudit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KDAuditSpec   `json:"spec"`
	Status KDAuditStatus `json:"status"`
}

type NamespaceWatcher struct {
	Enabled             string            `json:"enabled"`
	NamespacesConfigMap map[string]string `json:"namespacesConfigMap"`
}

// KDAuditSpec is the spec for a KDAudit resource
type KDAuditSpec struct {
	DeploymentName   string           `json:"deploymentName"`
	Replicas         *int32           `json:"replicas"`
	SlackWebhook     string           `json:"slackWebhook"`
	Service          string           `json:"service"`
	Image            string           `json:"image"`
	NamespaceWatcher NamespaceWatcher `json:"namespaceWatcher"`
}

// KDAuditStatus is the status for a KDAudit resource
type KDAuditStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KDAuditList is a list of KDAudit resources
type KDAuditList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []KDAudit `json:"items"`
}
