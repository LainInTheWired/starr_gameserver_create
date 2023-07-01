package main

import (
	"fmt"
	"log"

	// "time"

	agonesv1 "agones.dev/agones/pkg/apis/agones/v1"
	e2eframework "agones.dev/agones/test/e2e/framework"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	framework *e2eframework.Framework
)

func main() {
	fw, err := e2eframework.NewFromFlags()
	if err != nil {
		log.Fatalf("failed to init e2e e2eframework: %+v", err)
	}
	framework = fw

	namespace := "default"

	gsSpec := agonesv1.GameServerSpec{
		Ports: []agonesv1.GameServerPort{
			{
				ContainerPort: 7654,
				Name:          "tcp",
				PortPolicy:    agonesv1.Dynamic,
				Protocol:      corev1.ProtocolTCP,
			}},
		Template: corev1.PodTemplateSpec{
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{{
					Name:            "tcp-server",
					Image:           "gcr.io/agones-images/tcp-server:0.4",
					ImagePullPolicy: corev1.PullIfNotPresent,
					Resources: corev1.ResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("30m"),
							corev1.ResourceMemory: resource.MustParse("32Mi"),
						},
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("30m"),
							corev1.ResourceMemory: resource.MustParse("32Mi"),
						},
					},
				}},
			},
		},
	}
	fmt.Print(gsSpec)
	// Fleetを作る
	fleetDef := &agonesv1.Fleet{
		ObjectMeta: metav1.ObjectMeta{Name: "fleet-name", Namespace: namespace},
		Spec: agonesv1.FleetSpec{
			Replicas: 2,
			Template: agonesv1.GameServerTemplateSpec{
				Spec: gsSpec,
			},
		},
	}
	// flt, err := framework.AgonesClient.AgonesV1().Fleets(namespace).Create(fleetDef)
	framework.CreateGameServerAndWaitUntilReady(_, _, gsSpec)
	fmt.Print(fleetDef)
	// log.Printf("fleet created! (name: %s)", flt.ObjectMeta.Name)

}
