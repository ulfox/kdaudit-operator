package controllers

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	kdauditv1alpha1 "github.com/ulfox/kdaudit-operator/pkg/apis/kdaudit/v1alpha1"
)

func setLabels(kdAudit *kdauditv1alpha1.KDAudit) map[string]string {
	labels := map[string]string{
		"app":     "kdaudit-namespace-watcher",
		"watcher": kdAudit.Name,
	}
	return labels
}

func getType(t string) (ControllerType, error) {
	var rt ControllerType
	switch t {
	case "namespaceWatcher":
		rt = NamespaceWatcher
	default:
		return 0, errors.New("Given type is not a known type")
	}
	return rt, nil
}

func getEncodedData(kdAudit *kdauditv1alpha1.KDAudit) (string, error) {
	deploymentType, err := getType(kdAudit.Spec.Service)
	if err != nil {
		return "", err
	}

	var data map[string]string
	switch deploymentType {
	case NamespaceWatcher:
		data = kdAudit.Spec.NamespaceWatcher.NamespacesConfigMap
	default:
		return "", nil
	}

	jsonByte, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	encodedData := base64.StdEncoding.EncodeToString([]byte(jsonByte))
	return encodedData, nil
}
