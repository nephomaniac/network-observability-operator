//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/autoscaling/v2beta2"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateReference) DeepCopyInto(out *CertificateReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateReference.
func (in *CertificateReference) DeepCopy() *CertificateReference {
	if in == nil {
		return nil
	}
	out := new(CertificateReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientTLS) DeepCopyInto(out *ClientTLS) {
	*out = *in
	out.CACert = in.CACert
	out.UserCert = in.UserCert
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientTLS.
func (in *ClientTLS) DeepCopy() *ClientTLS {
	if in == nil {
		return nil
	}
	out := new(ClientTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworkOperatorConfig) DeepCopyInto(out *ClusterNetworkOperatorConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworkOperatorConfig.
func (in *ClusterNetworkOperatorConfig) DeepCopy() *ClusterNetworkOperatorConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkOperatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsolePluginPortConfig) DeepCopyInto(out *ConsolePluginPortConfig) {
	*out = *in
	if in.PortNames != nil {
		in, out := &in.PortNames, &out.PortNames
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsolePluginPortConfig.
func (in *ConsolePluginPortConfig) DeepCopy() *ConsolePluginPortConfig {
	if in == nil {
		return nil
	}
	out := new(ConsolePluginPortConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FLPMetrics) DeepCopyInto(out *FLPMetrics) {
	*out = *in
	in.Server.DeepCopyInto(&out.Server)
	if in.IgnoreTags != nil {
		in, out := &in.IgnoreTags, &out.IgnoreTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FLPMetrics.
func (in *FLPMetrics) DeepCopy() *FLPMetrics {
	if in == nil {
		return nil
	}
	out := new(FLPMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowCollector) DeepCopyInto(out *FlowCollector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowCollector.
func (in *FlowCollector) DeepCopy() *FlowCollector {
	if in == nil {
		return nil
	}
	out := new(FlowCollector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlowCollector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowCollectorAgent) DeepCopyInto(out *FlowCollectorAgent) {
	*out = *in
	out.IPFIX = in.IPFIX
	in.EBPF.DeepCopyInto(&out.EBPF)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowCollectorAgent.
func (in *FlowCollectorAgent) DeepCopy() *FlowCollectorAgent {
	if in == nil {
		return nil
	}
	out := new(FlowCollectorAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowCollectorConsolePlugin) DeepCopyInto(out *FlowCollectorConsolePlugin) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	in.Autoscaler.DeepCopyInto(&out.Autoscaler)
	in.PortNaming.DeepCopyInto(&out.PortNaming)
	if in.QuickFilters != nil {
		in, out := &in.QuickFilters, &out.QuickFilters
		*out = make([]QuickFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowCollectorConsolePlugin.
func (in *FlowCollectorConsolePlugin) DeepCopy() *FlowCollectorConsolePlugin {
	if in == nil {
		return nil
	}
	out := new(FlowCollectorConsolePlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowCollectorEBPF) DeepCopyInto(out *FlowCollectorEBPF) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Sampling != nil {
		in, out := &in.Sampling, &out.Sampling
		*out = new(int32)
		**out = **in
	}
	if in.Interfaces != nil {
		in, out := &in.Interfaces, &out.Interfaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeInterfaces != nil {
		in, out := &in.ExcludeInterfaces, &out.ExcludeInterfaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowCollectorEBPF.
func (in *FlowCollectorEBPF) DeepCopy() *FlowCollectorEBPF {
	if in == nil {
		return nil
	}
	out := new(FlowCollectorEBPF)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowCollectorExporter) DeepCopyInto(out *FlowCollectorExporter) {
	*out = *in
	out.Kafka = in.Kafka
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowCollectorExporter.
func (in *FlowCollectorExporter) DeepCopy() *FlowCollectorExporter {
	if in == nil {
		return nil
	}
	out := new(FlowCollectorExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowCollectorFLP) DeepCopyInto(out *FlowCollectorFLP) {
	*out = *in
	in.Metrics.DeepCopyInto(&out.Metrics)
	in.Resources.DeepCopyInto(&out.Resources)
	in.KafkaConsumerAutoscaler.DeepCopyInto(&out.KafkaConsumerAutoscaler)
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowCollectorFLP.
func (in *FlowCollectorFLP) DeepCopy() *FlowCollectorFLP {
	if in == nil {
		return nil
	}
	out := new(FlowCollectorFLP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowCollectorHPA) DeepCopyInto(out *FlowCollectorHPA) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]v2beta2.MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowCollectorHPA.
func (in *FlowCollectorHPA) DeepCopy() *FlowCollectorHPA {
	if in == nil {
		return nil
	}
	out := new(FlowCollectorHPA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowCollectorIPFIX) DeepCopyInto(out *FlowCollectorIPFIX) {
	*out = *in
	out.ClusterNetworkOperator = in.ClusterNetworkOperator
	out.OVNKubernetes = in.OVNKubernetes
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowCollectorIPFIX.
func (in *FlowCollectorIPFIX) DeepCopy() *FlowCollectorIPFIX {
	if in == nil {
		return nil
	}
	out := new(FlowCollectorIPFIX)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowCollectorKafka) DeepCopyInto(out *FlowCollectorKafka) {
	*out = *in
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowCollectorKafka.
func (in *FlowCollectorKafka) DeepCopy() *FlowCollectorKafka {
	if in == nil {
		return nil
	}
	out := new(FlowCollectorKafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowCollectorList) DeepCopyInto(out *FlowCollectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlowCollector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowCollectorList.
func (in *FlowCollectorList) DeepCopy() *FlowCollectorList {
	if in == nil {
		return nil
	}
	out := new(FlowCollectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlowCollectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowCollectorLoki) DeepCopyInto(out *FlowCollectorLoki) {
	*out = *in
	out.BatchWait = in.BatchWait
	out.Timeout = in.Timeout
	out.MinBackoff = in.MinBackoff
	out.MaxBackoff = in.MaxBackoff
	if in.StaticLabels != nil {
		in, out := &in.StaticLabels, &out.StaticLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowCollectorLoki.
func (in *FlowCollectorLoki) DeepCopy() *FlowCollectorLoki {
	if in == nil {
		return nil
	}
	out := new(FlowCollectorLoki)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowCollectorSpec) DeepCopyInto(out *FlowCollectorSpec) {
	*out = *in
	in.Agent.DeepCopyInto(&out.Agent)
	in.Processor.DeepCopyInto(&out.Processor)
	in.Loki.DeepCopyInto(&out.Loki)
	in.ConsolePlugin.DeepCopyInto(&out.ConsolePlugin)
	out.Kafka = in.Kafka
	if in.Exporters != nil {
		in, out := &in.Exporters, &out.Exporters
		*out = make([]*FlowCollectorExporter, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FlowCollectorExporter)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowCollectorSpec.
func (in *FlowCollectorSpec) DeepCopy() *FlowCollectorSpec {
	if in == nil {
		return nil
	}
	out := new(FlowCollectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowCollectorStatus) DeepCopyInto(out *FlowCollectorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowCollectorStatus.
func (in *FlowCollectorStatus) DeepCopy() *FlowCollectorStatus {
	if in == nil {
		return nil
	}
	out := new(FlowCollectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsServerConfig) DeepCopyInto(out *MetricsServerConfig) {
	*out = *in
	in.TLS.DeepCopyInto(&out.TLS)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsServerConfig.
func (in *MetricsServerConfig) DeepCopy() *MetricsServerConfig {
	if in == nil {
		return nil
	}
	out := new(MetricsServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OVNKubernetesConfig) DeepCopyInto(out *OVNKubernetesConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OVNKubernetesConfig.
func (in *OVNKubernetesConfig) DeepCopy() *OVNKubernetesConfig {
	if in == nil {
		return nil
	}
	out := new(OVNKubernetesConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuickFilter) DeepCopyInto(out *QuickFilter) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuickFilter.
func (in *QuickFilter) DeepCopy() *QuickFilter {
	if in == nil {
		return nil
	}
	out := new(QuickFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerTLS) DeepCopyInto(out *ServerTLS) {
	*out = *in
	if in.Provided != nil {
		in, out := &in.Provided, &out.Provided
		*out = new(CertificateReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerTLS.
func (in *ServerTLS) DeepCopy() *ServerTLS {
	if in == nil {
		return nil
	}
	out := new(ServerTLS)
	in.DeepCopyInto(out)
	return out
}
