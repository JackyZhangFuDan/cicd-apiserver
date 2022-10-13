package jenkinsservice

import (
	"cicd-apiserver/pkg/apis/cicd"
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
)

type jenkinsServiceStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func NewStrategy(typer runtime.ObjectTyper) jenkinsServiceStrategy {
	return jenkinsServiceStrategy{typer, names.SimpleNameGenerator}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	object, ok := obj.(*cicd.JenkinsService)
	if !ok {
		return nil, nil, fmt.Errorf("the object isn't a JenkinsService")
	}
	fs := generic.ObjectMetaFieldsSet(&object.ObjectMeta, true)
	return labels.Set(object.ObjectMeta.Labels), fs, nil
}

func MatchJenkinsService(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// CreateStrategy接口定义的方法
func (jenkinsServiceStrategy) AllowCreateOnUpdate() bool {
	return false
}
func (jenkinsServiceStrategy) Canonicalize(obj runtime.Object) {

}
func (jenkinsServiceStrategy) NamespaceScoped() bool {
	return true
}
func (jenkinsServiceStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {

}
func (jenkinsServiceStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	errs := field.ErrorList{} //承载发现的错误

	js := obj.(*cicd.JenkinsService)
	if js.Spec.InstanceAmount > 10 {
		errs = append(errs, field.TooMany(field.NewPath("spec").Key("instanceamount"), js.Spec.InstanceAmount, 10))
	}
	if len(errs) > 0 {
		return errs
	} else {
		return nil
	}
}
func (jenkinsServiceStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return []string{}
}

// UpdateStrategy接口定义的方法
func (jenkinsServiceStrategy) AllowUnconditionalUpdate() bool {
	return false
}
func (jenkinsServiceStrategy) PrepareForUpdate(ctx context.Context, obj runtime.Object, old runtime.Object) {

}
func (jenkinsServiceStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
func (jenkinsServiceStrategy) WarningsOnUpdate(ctx context.Context, obj runtime.Object, old runtime.Object) []string {
	return []string{}
}
