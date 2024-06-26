/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha2

const (
	TopologyPublic  = "public"
	TopologyPrivate = "private"
)

type TopologySpec struct {
	// Masters is not used.
	// +k8s:conversion-gen=false
	Masters string `json:"masters,omitempty"`

	// Nodes is not used.
	// +k8s:conversion-gen=false
	Nodes string `json:"nodes,omitempty"`

	// Bastion provide an external facing point of entry into a network
	// containing private network instances. This host can provide a single
	// point of fortification or audit and can be started and stopped to enable
	// or disable inbound SSH communication from the Internet, some call bastion
	// as the "jump server".
	Bastion *BastionSpec `json:"bastion,omitempty"`

	DNS DNSType `json:"-"`

	// DNS configures options relating to DNS, in particular whether we use a public or a private hosted zone
	// +k8s:conversion-gen=false
	LegacyDNS *DNSSpec `json:"dns,omitempty"`
}

type DNSSpec struct {
	Type DNSType `json:"type,omitempty"`
}

type DNSType string

const (
	DNSTypePublic  DNSType = "Public"
	DNSTypePrivate DNSType = "Private"
)
