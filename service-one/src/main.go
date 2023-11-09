package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	http.HandleFunc("/pods/amount/update/", podsHandler)
	http.HandleFunc("/pods/health/", healthHandler)

	http.ListenAndServe(":8081", nil)
}

func podsHandler(w http.ResponseWriter, r *http.Request) {
	amountStr := extractAmount(r.URL.Path, "/pods/amount/update/")
	amount := normalizeValue(amountStr)
	err := scaleDeploymentPods("default", "service-two", amount)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, "Pods up: %s", amountStr)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey, I'm alive!")
}

func normalizeValue(value string) int32 {
	intVal, err := strconv.Atoi(value)

	if err != nil {
		panic(err.Error())
	}

	return int32(intVal)
}

func extractAmount(path, prefix string) string {
	return path[len(prefix):]
}

func outClusterConfig() kubernetes.Clientset {

	kubeconfigPath := "kubeconfig.yaml"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)

	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	return *clientset

}

func inClusterConfig() kubernetes.Clientset {
	config, err := rest.InClusterConfig()

	if err != nil {
		if err != nil {
			panic(err.Error())
		}
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	return *clientset
}

func scaleDeploymentPods(namespace, deploymentName string, replicas int32) error {

	clientset := outClusterConfig()
	// clientset := inClusterConfig()

	deploymentsClient := clientset.AppsV1().Deployments(namespace)

	deployment, err := deploymentsClient.Get(context.TODO(), deploymentName, metav1.GetOptions{})

	if err != nil {
		panic((err.Error()))
	}

	deployment.Spec.Replicas = &replicas

	_, err = deploymentsClient.Update(context.TODO(), deployment, metav1.UpdateOptions{})

	return err
}
