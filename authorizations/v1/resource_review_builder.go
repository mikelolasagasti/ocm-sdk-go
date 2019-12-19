/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/authorizations/v1

// ResourceReviewBuilder contains the data and logic needed to build 'resource_review' objects.
//
// Contains the result of performing a resource access review.
type ResourceReviewBuilder struct {
	accountUsername *string
	action          *string
	clusterIDs      []string
	clusterUUIDs    []string
	organizationIDs []string
	resourceType    *string
	subscriptionIDs []string
}

// NewResourceReview creates a new builder of 'resource_review' objects.
func NewResourceReview() *ResourceReviewBuilder {
	return new(ResourceReviewBuilder)
}

// AccountUsername sets the value of the 'account_username' attribute
// to the given value.
//
//
func (b *ResourceReviewBuilder) AccountUsername(value string) *ResourceReviewBuilder {
	b.accountUsername = &value
	return b
}

// Action sets the value of the 'action' attribute
// to the given value.
//
//
func (b *ResourceReviewBuilder) Action(value string) *ResourceReviewBuilder {
	b.action = &value
	return b
}

// ClusterIDs sets the value of the 'cluster_IDs' attribute
// to the given values.
//
//
func (b *ResourceReviewBuilder) ClusterIDs(values ...string) *ResourceReviewBuilder {
	b.clusterIDs = make([]string, len(values))
	copy(b.clusterIDs, values)
	return b
}

// ClusterUUIDs sets the value of the 'cluster_UUIDs' attribute
// to the given values.
//
//
func (b *ResourceReviewBuilder) ClusterUUIDs(values ...string) *ResourceReviewBuilder {
	b.clusterUUIDs = make([]string, len(values))
	copy(b.clusterUUIDs, values)
	return b
}

// OrganizationIDs sets the value of the 'organization_IDs' attribute
// to the given values.
//
//
func (b *ResourceReviewBuilder) OrganizationIDs(values ...string) *ResourceReviewBuilder {
	b.organizationIDs = make([]string, len(values))
	copy(b.organizationIDs, values)
	return b
}

// ResourceType sets the value of the 'resource_type' attribute
// to the given value.
//
//
func (b *ResourceReviewBuilder) ResourceType(value string) *ResourceReviewBuilder {
	b.resourceType = &value
	return b
}

// SubscriptionIDs sets the value of the 'subscription_IDs' attribute
// to the given values.
//
//
func (b *ResourceReviewBuilder) SubscriptionIDs(values ...string) *ResourceReviewBuilder {
	b.subscriptionIDs = make([]string, len(values))
	copy(b.subscriptionIDs, values)
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ResourceReviewBuilder) Copy(object *ResourceReview) *ResourceReviewBuilder {
	if object == nil {
		return b
	}
	b.accountUsername = object.accountUsername
	b.action = object.action
	if len(object.clusterIDs) > 0 {
		b.clusterIDs = make([]string, len(object.clusterIDs))
		copy(b.clusterIDs, object.clusterIDs)
	} else {
		b.clusterIDs = nil
	}
	if len(object.clusterUUIDs) > 0 {
		b.clusterUUIDs = make([]string, len(object.clusterUUIDs))
		copy(b.clusterUUIDs, object.clusterUUIDs)
	} else {
		b.clusterUUIDs = nil
	}
	if len(object.organizationIDs) > 0 {
		b.organizationIDs = make([]string, len(object.organizationIDs))
		copy(b.organizationIDs, object.organizationIDs)
	} else {
		b.organizationIDs = nil
	}
	b.resourceType = object.resourceType
	if len(object.subscriptionIDs) > 0 {
		b.subscriptionIDs = make([]string, len(object.subscriptionIDs))
		copy(b.subscriptionIDs, object.subscriptionIDs)
	} else {
		b.subscriptionIDs = nil
	}
	return b
}

// Build creates a 'resource_review' object using the configuration stored in the builder.
func (b *ResourceReviewBuilder) Build() (object *ResourceReview, err error) {
	object = new(ResourceReview)
	if b.accountUsername != nil {
		object.accountUsername = b.accountUsername
	}
	if b.action != nil {
		object.action = b.action
	}
	if b.clusterIDs != nil {
		object.clusterIDs = make([]string, len(b.clusterIDs))
		copy(object.clusterIDs, b.clusterIDs)
	}
	if b.clusterUUIDs != nil {
		object.clusterUUIDs = make([]string, len(b.clusterUUIDs))
		copy(object.clusterUUIDs, b.clusterUUIDs)
	}
	if b.organizationIDs != nil {
		object.organizationIDs = make([]string, len(b.organizationIDs))
		copy(object.organizationIDs, b.organizationIDs)
	}
	if b.resourceType != nil {
		object.resourceType = b.resourceType
	}
	if b.subscriptionIDs != nil {
		object.subscriptionIDs = make([]string, len(b.subscriptionIDs))
		copy(object.subscriptionIDs, b.subscriptionIDs)
	}
	return
}