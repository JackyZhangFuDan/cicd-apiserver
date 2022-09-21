package v1alpha

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type JenkinsService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,name=metadata"`

	Spec   JenkinsServiceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status JenkinsServiceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type JenkinsServiceSpec struct {
	InstanceAmount int `json:"instanceamount,omitempty" protobuf:"int32,1,opt,name=instanceamount"`
	InstanceCpu    int `json:"metadata,omitempty" protobuf:"int32,2,opt,name=instancecpu"`
}

type JenkinsServiceStatus struct {
	ApprovalStatus string                  `json:"approvalstatus" protobuf:"bytes,1,name=approvalstatus"`
	Instances      []JenkinsServerInstance `json:"instances" protobuf:"bytes,2,rep,name=instances"`
}

type JenkinsServerInstance struct {
	Cpu     int  `json:"cpu" protobuf:"int32,1,name=cpu"`
	Running bool `json:"running" protobuf:"bool,2,name=running"`
}
