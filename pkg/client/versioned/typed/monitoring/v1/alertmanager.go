// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1"
	scheme "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/scheme"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// AlertmanagersGetter has a method to return a AlertmanagerInterface.
// A group's client should implement this interface.
type AlertmanagersGetter interface {
	Alertmanagers(namespace string) AlertmanagerInterface
}

// AlertmanagerInterface has methods to work with Alertmanager resources.
type AlertmanagerInterface interface {
	Create(ctx context.Context, alertmanager *v1.Alertmanager, opts metav1.CreateOptions) (*v1.Alertmanager, error)
	Update(ctx context.Context, alertmanager *v1.Alertmanager, opts metav1.UpdateOptions) (*v1.Alertmanager, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, alertmanager *v1.Alertmanager, opts metav1.UpdateOptions) (*v1.Alertmanager, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Alertmanager, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AlertmanagerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Alertmanager, err error)
	Apply(ctx context.Context, alertmanager *monitoringv1.AlertmanagerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Alertmanager, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, alertmanager *monitoringv1.AlertmanagerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Alertmanager, err error)
	GetScale(ctx context.Context, alertmanagerName string, options metav1.GetOptions) (*autoscalingv1.Scale, error)
	UpdateScale(ctx context.Context, alertmanagerName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (*autoscalingv1.Scale, error)

	AlertmanagerExpansion
}

// alertmanagers implements AlertmanagerInterface
type alertmanagers struct {
	*gentype.ClientWithListAndApply[*v1.Alertmanager, *v1.AlertmanagerList, *monitoringv1.AlertmanagerApplyConfiguration]
}

// newAlertmanagers returns a Alertmanagers
func newAlertmanagers(c *MonitoringV1Client, namespace string) *alertmanagers {
	return &alertmanagers{
		gentype.NewClientWithListAndApply[*v1.Alertmanager, *v1.AlertmanagerList, *monitoringv1.AlertmanagerApplyConfiguration](
			"alertmanagers",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.Alertmanager { return &v1.Alertmanager{} },
			func() *v1.AlertmanagerList { return &v1.AlertmanagerList{} }),
	}
}

// GetScale takes name of the alertmanager, and returns the corresponding autoscalingv1.Scale object, and an error if there is any.
func (c *alertmanagers) GetScale(ctx context.Context, alertmanagerName string, options metav1.GetOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.GetClient().Get().
		Namespace(c.GetNamespace()).
		Resource("alertmanagers").
		Name(alertmanagerName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *alertmanagers) UpdateScale(ctx context.Context, alertmanagerName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.GetClient().Put().
		Namespace(c.GetNamespace()).
		Resource("alertmanagers").
		Name(alertmanagerName).
		SubResource("scale").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scale).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied alertmanager.
func (c *alertmanagers) Apply(ctx context.Context, alertmanager *monitoringv1.AlertmanagerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Alertmanager, err error) {
	if alertmanager == nil {
		return nil, fmt.Errorf("alertmanager provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(alertmanager)
	if err != nil {
		return nil, err
	}
	name := alertmanager.Name
	if name == nil {
		return nil, fmt.Errorf("alertmanager.Name must be provided to Apply")
	}
	result = &v1.Alertmanager{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("alertmanagers").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *alertmanagers) ApplyStatus(ctx context.Context, alertmanager *monitoringv1.AlertmanagerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Alertmanager, err error) {
	if alertmanager == nil {
		return nil, fmt.Errorf("alertmanager provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(alertmanager)
	if err != nil {
		return nil, err
	}

	name := alertmanager.Name
	if name == nil {
		return nil, fmt.Errorf("alertmanager.Name must be provided to Apply")
	}

	result = &v1.Alertmanager{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("alertmanagers").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
