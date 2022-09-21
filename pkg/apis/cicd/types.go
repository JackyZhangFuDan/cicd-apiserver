package cicd

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type JenkinsService struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   JenkinsServiceSpec
	Status JenkinsServiceStatus
}

type JenkinsServiceSpec struct {
	InstanceAmount int
	InstanceCpu    int
}

type JenkinsServiceStatus struct {
	ApprovalStatus string
	Instances      []JenkinsServerInstance
}

type JenkinsServerInstance struct {
	Cpu     int
	Running bool
}
