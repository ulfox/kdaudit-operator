package main

import (
	"time"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/klog/v2"

	"github.com/ulfox/kdaudit-operator/controllers"
	clientset "github.com/ulfox/kdaudit-operator/pkg/generated/clientset/versioned"
	informers "github.com/ulfox/kdaudit-operator/pkg/generated/informers/externalversions"
	"github.com/ulfox/kdaudit-operator/utils"
)

func main() {
	klog.InitFlags(nil)

	stopCh := utils.SetupSignalHandler()

	kdAuditFlags := utils.ParseFlags()
	kubeClient, cfg := utils.NewClientAuth(
		kdAuditFlags.ClientAuthType,
		kdAuditFlags.LocalKubeCFG,
	)
	kdAuditClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatal(err)
	}

	kdAuditCFG := utils.NewKDAuditCFG(kubeClient)
	kdAuditCFG.ReadConfig()

	kubeInformer := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	kdauditInformer := informers.NewSharedInformerFactory(kdAuditClient, time.Second*30)

	controller := controllers.NewController(
		kubeClient,
		kdAuditClient,
		kubeInformer.Apps().V1().Deployments(),
		kdauditInformer.Kdaudit().V1alpha1().KDAudits(),
	)

	kubeInformer.Start(stopCh)
	kdauditInformer.Start(stopCh)

	if err = controller.Run(kdAuditCFG.Threadiness, stopCh); err != nil {
		klog.Fatal(err)
	}
}
