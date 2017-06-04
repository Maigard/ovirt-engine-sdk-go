//
// Copyright (c) 2017 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package ovirtsdk4

import (
    "fmt"
)

//
// This service manages a single affinity group.
//
type AffinityGroupService struct {
    Service

    VmsService  *VmsService
}

func NewAffinityGroupService(connection *Connection, path string) *AffinityGroupService {
    var result AffinityGroupService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieve the affinity group details.
// [source,xml]
// ----
// <affinity_group id="00000000-0000-0000-0000-000000000000">
//   <name>AF_GROUP_001</name>
//   <cluster id="00000000-0000-0000-0000-000000000000"/>
//   <positive>true</positive>
//   <enforcing>true</enforcing>
// </affinity_group>
// ----
//
func (op *AffinityGroupService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Remove the affinity group.
// [source]
// ----
// DELETE /ovirt-engine/api/clusters/000-000/affinitygroups/123-456
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the removal should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AffinityGroupService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Update the affinity group.
// This method supports the following parameters:
// `Group`:: The affinity group.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AffinityGroupService) Update (
    group AffinityGroup,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(group, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AffinityGroupService) Service(path string) *AffinityGroupService {
    if path == nil {
        return op
    }
    if path == "vms" {
        return &VmsService{}
    }
    if path.startswith('vms/') {
        return VmsService().service(path[4:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *AffinityGroupService) String() string {
    return fmt.Sprintf("AffinityGroupService:%s", op.Path)
}


//
// This service manages a single virtual machine to affinity group assignment.
//
type AffinityGroupVmService struct {
    Service

}

func NewAffinityGroupVmService(connection *Connection, path string) *AffinityGroupVmService {
    var result AffinityGroupVmService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Remove this virtual machine from the affinity group.
// This method supports the following parameters:
// `Async`:: Indicates if the removal should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AffinityGroupVmService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AffinityGroupVmService) Service(path string) *AffinityGroupVmService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *AffinityGroupVmService) String() string {
    return fmt.Sprintf("AffinityGroupVmService:%s", op.Path)
}


//
// This service manages a collection of all the virtual machines assigned to an affinity group.
//
type AffinityGroupVmsService struct {
    Service

    VmService  *VmService
}

func NewAffinityGroupVmsService(connection *Connection, path string) *AffinityGroupVmsService {
    var result AffinityGroupVmsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add a virtual machine to the affinity group.
// For example, to add the virtual machine 000-000 to affinity group 123-456 send a request to:
// [source]
// ----
// POST /ovirt-engine/api/clusters/000-000/affinitygroups/123-456/vms
// ----
// With the following body:
// [source,xml]
// ----
// <vm id="000-000"/>
// ----
//
func (op *AffinityGroupVmsService) Add (
    vm Vm,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(vm, headers, query, wait)
}

//
// List all virtual machines assigned to this affinity group.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of virtual machines to return. If not specified, all the virtual machines are
// returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AffinityGroupVmsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AffinityGroupVmsService) Service(path string) *AffinityGroupVmsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Vm_service(path)
    return self.Vm_service(path[:index]).Service(path[index + 1:])
}

func (op *AffinityGroupVmsService) String() string {
    return fmt.Sprintf("AffinityGroupVmsService:%s", op.Path)
}


//
// The affinity groups service manages virtual machine relationships and dependencies.
//
type AffinityGroupsService struct {
    Service

    GroupService  *GroupService
}

func NewAffinityGroupsService(connection *Connection, path string) *AffinityGroupsService {
    var result AffinityGroupsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Create a new affinity group.
// Post a request like in the example below to create a new affinity group:
// [source]
// ----
// POST /ovirt-engine/api/clusters/000-000/affinitygroups
// ----
// And use the following example in its body:
// [source,xml]
// ----
// <affinity_group>
//   <name>AF_GROUP_001</name>
//   <positive>true</positive>
//   <enforcing>true</enforcing>
// </affinity_group>
// ----
// This method supports the following parameters:
// `Group`:: The affinity group object to create.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AffinityGroupsService) Add (
    group AffinityGroup,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(group, headers, query, wait)
}

//
// List existing affinity groups.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of affinity groups to return. If not specified all the affinity groups are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AffinityGroupsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AffinityGroupsService) Service(path string) *AffinityGroupsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Group_service(path)
    return self.Group_service(path[:index]).Service(path[index + 1:])
}

func (op *AffinityGroupsService) String() string {
    return fmt.Sprintf("AffinityGroupsService:%s", op.Path)
}


//
// The details of a single affinity label.
//
type AffinityLabelService struct {
    Service

    HostsService  *HostsService
    VmsService  *VmsService
}

func NewAffinityLabelService(connection *Connection, path string) *AffinityLabelService {
    var result AffinityLabelService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieves the details of a label.
//
func (op *AffinityLabelService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes a label from the system and clears all assignments
// of the removed label.
//
func (op *AffinityLabelService) Remove (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates a label. This call will update all metadata, such as the name
// or description.
//
func (op *AffinityLabelService) Update (
    label AffinityLabel,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalUpdate(label, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AffinityLabelService) Service(path string) *AffinityLabelService {
    if path == nil {
        return op
    }
    if path == "hosts" {
        return &HostsService{}
    }
    if path.startswith('hosts/') {
        return HostsService().service(path[6:])
    }
    if path == "vms" {
        return &VmsService{}
    }
    if path.startswith('vms/') {
        return VmsService().service(path[4:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *AffinityLabelService) String() string {
    return fmt.Sprintf("AffinityLabelService:%s", op.Path)
}


//
// This service represents a host that has a specific
// label when accessed through the affinitylabels/hosts
// subcollection.
//
type AffinityLabelHostService struct {
    Service

}

func NewAffinityLabelHostService(connection *Connection, path string) *AffinityLabelHostService {
    var result AffinityLabelHostService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieves details about a host that has this label assigned.
//
func (op *AffinityLabelHostService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Remove a label from a host.
//
func (op *AffinityLabelHostService) Remove (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AffinityLabelHostService) Service(path string) *AffinityLabelHostService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *AffinityLabelHostService) String() string {
    return fmt.Sprintf("AffinityLabelHostService:%s", op.Path)
}


//
// This service represents list of hosts that have a specific
// label when accessed through the affinitylabels/hosts
// subcollection.
//
type AffinityLabelHostsService struct {
    Service

    HostService  *HostService
}

func NewAffinityLabelHostsService(connection *Connection, path string) *AffinityLabelHostsService {
    var result AffinityLabelHostsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add a label to a host.
//
func (op *AffinityLabelHostsService) Add (
    host Host,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(host, headers, query, wait)
}

//
// List all hosts with the label.
//
func (op *AffinityLabelHostsService) List (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AffinityLabelHostsService) Service(path string) *AffinityLabelHostsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Host_service(path)
    return self.Host_service(path[:index]).Service(path[index + 1:])
}

func (op *AffinityLabelHostsService) String() string {
    return fmt.Sprintf("AffinityLabelHostsService:%s", op.Path)
}


//
// This service represents a vm that has a specific
// label when accessed through the affinitylabels/vms
// subcollection.
//
type AffinityLabelVmService struct {
    Service

}

func NewAffinityLabelVmService(connection *Connection, path string) *AffinityLabelVmService {
    var result AffinityLabelVmService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieves details about a vm that has this label assigned.
//
func (op *AffinityLabelVmService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Remove a label from a vm.
//
func (op *AffinityLabelVmService) Remove (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AffinityLabelVmService) Service(path string) *AffinityLabelVmService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *AffinityLabelVmService) String() string {
    return fmt.Sprintf("AffinityLabelVmService:%s", op.Path)
}


//
// This service represents list of vms that have a specific
// label when accessed through the affinitylabels/vms
// subcollection.
//
type AffinityLabelVmsService struct {
    Service

    VmService  *VmService
}

func NewAffinityLabelVmsService(connection *Connection, path string) *AffinityLabelVmsService {
    var result AffinityLabelVmsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add a label to a vm.
//
func (op *AffinityLabelVmsService) Add (
    vm Vm,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(vm, headers, query, wait)
}

//
// List all vms with the label.
//
func (op *AffinityLabelVmsService) List (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AffinityLabelVmsService) Service(path string) *AffinityLabelVmsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Vm_service(path)
    return self.Vm_service(path[:index]).Service(path[index + 1:])
}

func (op *AffinityLabelVmsService) String() string {
    return fmt.Sprintf("AffinityLabelVmsService:%s", op.Path)
}


//
// Manages the affinity labels available in the system.
//
type AffinityLabelsService struct {
    Service

    LabelService  *LabelService
}

func NewAffinityLabelsService(connection *Connection, path string) *AffinityLabelsService {
    var result AffinityLabelsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a new label. The label is automatically attached
// to all entities mentioned in the vms or hosts lists.
//
func (op *AffinityLabelsService) Add (
    label AffinityLabel,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(label, headers, query, wait)
}

//
// Lists all labels present in the system.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of labels to return. If not specified all the labels are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AffinityLabelsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AffinityLabelsService) Service(path string) *AffinityLabelsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Label_service(path)
    return self.Label_service(path[:index]).Service(path[index + 1:])
}

func (op *AffinityLabelsService) String() string {
    return fmt.Sprintf("AffinityLabelsService:%s", op.Path)
}


//
// This annotation is intended to specify what oVirt area is the annotated concept related to. Currently the following
// areas are in use, and they are closely related to the oVirt teams, but not necessarily the same:
// - Infrastructure
// - Network
// - SLA
// - Storage
// - Virtualization
// A concept may be associated to more than one area, or to no area.
// The value of this annotation is intended for reporting only, and it doesn't affect at all the generated code or the
// validity of the model
//
type AreaService struct {
    Service

}

func NewAreaService(connection *Connection, path string) *AreaService {
    var result AreaService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AreaService) Service(path string) *AreaService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *AreaService) String() string {
    return fmt.Sprintf("AreaService:%s", op.Path)
}


//
// This service represents one label to entity assignment
// when accessed using the entities/affinitylabels subcollection.
//
type AssignedAffinityLabelService struct {
    Service

}

func NewAssignedAffinityLabelService(connection *Connection, path string) *AssignedAffinityLabelService {
    var result AssignedAffinityLabelService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieves details about the attached label.
//
func (op *AssignedAffinityLabelService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes the label from an entity. Does not touch the label itself.
//
func (op *AssignedAffinityLabelService) Remove (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AssignedAffinityLabelService) Service(path string) *AssignedAffinityLabelService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *AssignedAffinityLabelService) String() string {
    return fmt.Sprintf("AssignedAffinityLabelService:%s", op.Path)
}


//
// This service is used to list and manipulate affinity labels that are
// assigned to supported entities when accessed using entities/affinitylabels.
//
type AssignedAffinityLabelsService struct {
    Service

    LabelService  *LabelService
}

func NewAssignedAffinityLabelsService(connection *Connection, path string) *AssignedAffinityLabelsService {
    var result AssignedAffinityLabelsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Attaches a label to an entity.
//
func (op *AssignedAffinityLabelsService) Add (
    label AffinityLabel,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(label, headers, query, wait)
}

//
// Lists all labels that are attached to an entity.
//
func (op *AssignedAffinityLabelsService) List (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AssignedAffinityLabelsService) Service(path string) *AssignedAffinityLabelsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Label_service(path)
    return self.Label_service(path[:index]).Service(path[index + 1:])
}

func (op *AssignedAffinityLabelsService) String() string {
    return fmt.Sprintf("AssignedAffinityLabelsService:%s", op.Path)
}


//
//
type AssignedCpuProfileService struct {
    Service

}

func NewAssignedCpuProfileService(connection *Connection, path string) *AssignedCpuProfileService {
    var result AssignedCpuProfileService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *AssignedCpuProfileService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AssignedCpuProfileService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AssignedCpuProfileService) Service(path string) *AssignedCpuProfileService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *AssignedCpuProfileService) String() string {
    return fmt.Sprintf("AssignedCpuProfileService:%s", op.Path)
}


//
//
type AssignedCpuProfilesService struct {
    Service

    ProfileService  *ProfileService
}

func NewAssignedCpuProfilesService(connection *Connection, path string) *AssignedCpuProfilesService {
    var result AssignedCpuProfilesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *AssignedCpuProfilesService) Add (
    profile CpuProfile,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(profile, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of profiles to return. If not specified all the profiles are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AssignedCpuProfilesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AssignedCpuProfilesService) Service(path string) *AssignedCpuProfilesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Profile_service(path)
    return self.Profile_service(path[:index]).Service(path[index + 1:])
}

func (op *AssignedCpuProfilesService) String() string {
    return fmt.Sprintf("AssignedCpuProfilesService:%s", op.Path)
}


//
//
type AssignedDiskProfileService struct {
    Service

}

func NewAssignedDiskProfileService(connection *Connection, path string) *AssignedDiskProfileService {
    var result AssignedDiskProfileService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *AssignedDiskProfileService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AssignedDiskProfileService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AssignedDiskProfileService) Service(path string) *AssignedDiskProfileService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *AssignedDiskProfileService) String() string {
    return fmt.Sprintf("AssignedDiskProfileService:%s", op.Path)
}


//
//
type AssignedDiskProfilesService struct {
    Service

    ProfileService  *ProfileService
}

func NewAssignedDiskProfilesService(connection *Connection, path string) *AssignedDiskProfilesService {
    var result AssignedDiskProfilesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *AssignedDiskProfilesService) Add (
    profile DiskProfile,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(profile, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of profiles to return. If not specified all the profiles are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AssignedDiskProfilesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AssignedDiskProfilesService) Service(path string) *AssignedDiskProfilesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Profile_service(path)
    return self.Profile_service(path[:index]).Service(path[index + 1:])
}

func (op *AssignedDiskProfilesService) String() string {
    return fmt.Sprintf("AssignedDiskProfilesService:%s", op.Path)
}


//
//
type AssignedNetworkService struct {
    Service

}

func NewAssignedNetworkService(connection *Connection, path string) *AssignedNetworkService {
    var result AssignedNetworkService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *AssignedNetworkService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AssignedNetworkService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
//
func (op *AssignedNetworkService) Update (
    network Network,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(network, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AssignedNetworkService) Service(path string) *AssignedNetworkService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *AssignedNetworkService) String() string {
    return fmt.Sprintf("AssignedNetworkService:%s", op.Path)
}


//
//
type AssignedNetworksService struct {
    Service

    NetworkService  *NetworkService
}

func NewAssignedNetworksService(connection *Connection, path string) *AssignedNetworksService {
    var result AssignedNetworksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *AssignedNetworksService) Add (
    network Network,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(network, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of networks to return. If not specified all the networks are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AssignedNetworksService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AssignedNetworksService) Service(path string) *AssignedNetworksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Network_service(path)
    return self.Network_service(path[:index]).Service(path[index + 1:])
}

func (op *AssignedNetworksService) String() string {
    return fmt.Sprintf("AssignedNetworksService:%s", op.Path)
}


//
// Represents a permission sub-collection, scoped by user, group or some entity type.
//
type AssignedPermissionsService struct {
    Service

    PermissionService  *PermissionService
}

func NewAssignedPermissionsService(connection *Connection, path string) *AssignedPermissionsService {
    var result AssignedPermissionsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Assign a new permission to a user or group for specific entity.
// For example, to assign the `UserVmManager` role to the virtual machine with id `123` to the user with id `456`
// send a request like this:
// ....
// POST /ovirt-engine/api/vms/123/permissions
// ....
// With a request body like this:
// [source,xml]
// ----
// <permission>
//   <role>
//     <name>UserVmManager</name>
//   </role>
//   <user id="456"/>
// </permission>
// ----
// To assign the `SuperUser` role to the system to the user with id `456` send a request like this:
// ....
// POST /ovirt-engine/api/permissions
// ....
// With a request body like this:
// [source,xml]
// ----
// <permission>
//   <role>
//     <name>SuperUser</name>
//   </role>
//   <user id="456"/>
// </permission>
// ----
// If you want to assign permission to the group instead of the user please replace the `user` element with the
// `group` element with proper `id` of the group. For example to assign the `UserRole` role to the cluster with
// id `123` to the group with id `789` send a request like this:
// ....
// POST /ovirt-engine/api/clusters/123/permissions
// ....
// With a request body like this:
// [source,xml]
// ----
// <permission>
//   <role>
//     <name>UserRole</name>
//   </role>
//   <group id="789"/>
// </permission>
// ----
// This method supports the following parameters:
// `Permission`:: The permission.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AssignedPermissionsService) Add (
    permission Permission,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(permission, headers, query, wait)
}

//
// List all the permissions of the specific entity.
// For example to list all the permissions of the cluster with id `123` send a request like this:
// ....
// GET /ovirt-engine/api/clusters/123/permissions
// ....
// [source,xml]
// ----
// <permissions>
//   <permission id="456">
//     <cluster id="123"/>
//     <role id="789"/>
//     <user id="451"/>
//   </permission>
//   <permission id="654">
//     <cluster id="123"/>
//     <role id="789"/>
//     <group id="127"/>
//   </permission>
// </permissions>
// ----
//
func (op *AssignedPermissionsService) List (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AssignedPermissionsService) Service(path string) *AssignedPermissionsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Permission_service(path)
    return self.Permission_service(path[:index]).Service(path[index + 1:])
}

func (op *AssignedPermissionsService) String() string {
    return fmt.Sprintf("AssignedPermissionsService:%s", op.Path)
}


//
// Represents a roles sub-collection, for example scoped by user.
//
type AssignedRolesService struct {
    Service

    RoleService  *RoleService
}

func NewAssignedRolesService(connection *Connection, path string) *AssignedRolesService {
    var result AssignedRolesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of roles to return. If not specified all the roles are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AssignedRolesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AssignedRolesService) Service(path string) *AssignedRolesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Role_service(path)
    return self.Role_service(path[:index]).Service(path[index + 1:])
}

func (op *AssignedRolesService) String() string {
    return fmt.Sprintf("AssignedRolesService:%s", op.Path)
}


//
// A service to manage assignment of specific tag to specific entities in system.
//
type AssignedTagService struct {
    Service

}

func NewAssignedTagService(connection *Connection, path string) *AssignedTagService {
    var result AssignedTagService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Gets the information about the assigned tag.
// For example to retrieve the information about the tag with the id `456` which is assigned to virtual machine
// with id `123` send a request like this:
// ....
// GET /ovirt-engine/api/vms/123/tags/456
// ....
// [source,xml]
// ----
// <tag href="/ovirt-engine/api/tags/456" id="456">
//   <name>root</name>
//   <description>root</description>
//   <vm href="/ovirt-engine/api/vms/123" id="123"/>
// </tag>
// ----
//
func (op *AssignedTagService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Unassign tag from specific entity in the system.
// For example to unassign the tag with id `456` from virtual machine with id `123` send a request like this:
// ....
// DELETE /ovirt-engine/api/vms/123/tags/456
// ....
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AssignedTagService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AssignedTagService) Service(path string) *AssignedTagService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *AssignedTagService) String() string {
    return fmt.Sprintf("AssignedTagService:%s", op.Path)
}


//
// A service to manage collection of assignment of tags to specific entities in system.
//
type AssignedTagsService struct {
    Service

    TagService  *TagService
}

func NewAssignedTagsService(connection *Connection, path string) *AssignedTagsService {
    var result AssignedTagsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Assign tag to specific entity in the system.
// For example to assign tag `mytag` to virtual machine with the id `123` send a request like this:
// ....
// POST /ovirt-engine/api/vms/123/tags
// ....
// With a request body like this:
// [source,xml]
// ----
// <tag>
//   <name>mytag</name>
// </tag>
// ----
// This method supports the following parameters:
// `Tag`:: The assigned tag.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AssignedTagsService) Add (
    tag Tag,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(tag, headers, query, wait)
}

//
// List all tags assigned to the specific entity.
// For example to list all the tags of the virtual machine with id `123` send a request like this:
// ....
// GET /ovirt-engine/api/vms/123/tags
// ....
// [source,xml]
// ----
// <tags>
//   <tag href="/ovirt-engine/api/tags/222" id="222">
//     <name>mytag</name>
//     <description>mytag</description>
//     <vm href="/ovirt-engine/api/vms/123" id="123"/>
//   </tag>
// </tags>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of tags to return. If not specified all the tags are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AssignedTagsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AssignedTagsService) Service(path string) *AssignedTagsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Tag_service(path)
    return self.Tag_service(path[:index]).Service(path[index + 1:])
}

func (op *AssignedTagsService) String() string {
    return fmt.Sprintf("AssignedTagsService:%s", op.Path)
}


//
//
type AssignedVnicProfileService struct {
    Service

    PermissionsService  *PermissionsService
}

func NewAssignedVnicProfileService(connection *Connection, path string) *AssignedVnicProfileService {
    var result AssignedVnicProfileService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *AssignedVnicProfileService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AssignedVnicProfileService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AssignedVnicProfileService) Service(path string) *AssignedVnicProfileService {
    if path == nil {
        return op
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *AssignedVnicProfileService) String() string {
    return fmt.Sprintf("AssignedVnicProfileService:%s", op.Path)
}


//
//
type AssignedVnicProfilesService struct {
    Service

    ProfileService  *ProfileService
}

func NewAssignedVnicProfilesService(connection *Connection, path string) *AssignedVnicProfilesService {
    var result AssignedVnicProfilesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *AssignedVnicProfilesService) Add (
    profile VnicProfile,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(profile, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of profiles to return. If not specified all the profiles are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AssignedVnicProfilesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AssignedVnicProfilesService) Service(path string) *AssignedVnicProfilesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Profile_service(path)
    return self.Profile_service(path[:index]).Service(path[index + 1:])
}

func (op *AssignedVnicProfilesService) String() string {
    return fmt.Sprintf("AssignedVnicProfilesService:%s", op.Path)
}


//
//
type AttachedStorageDomainService struct {
    Service

    DisksService  *DisksService
}

func NewAttachedStorageDomainService(connection *Connection, path string) *AttachedStorageDomainService {
    var result AttachedStorageDomainService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This operation activates an attached storage domain.
// Once the storage domain is activated it is ready for use with the data center.
// [source]
// ----
// POST /ovirt-engine/api/datacenters/123/storagedomains/456/activate
// ----
// The activate action does not take any action specific parameters,
// so the request body should contain an empty `action`:
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the activation should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AttachedStorageDomainService) Activate (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'activate', nil, headers, query, wait)
}

//
// This operation deactivates an attached storage domain.
// Once the storage domain is deactivated it will not be used with the data center.
// [source]
// ----
// POST /ovirt-engine/api/datacenters/123/storagedomains/456/deactivate
// ----
// The deactivate action does not take any action specific parameters,
// so the request body should contain an empty `action`:
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the deactivation should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AttachedStorageDomainService) Deactivate (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'deactivate', nil, headers, query, wait)
}

//
//
func (op *AttachedStorageDomainService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AttachedStorageDomainService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AttachedStorageDomainService) Service(path string) *AttachedStorageDomainService {
    if path == nil {
        return op
    }
    if path == "disks" {
        return &DisksService{}
    }
    if path.startswith('disks/') {
        return DisksService().service(path[6:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *AttachedStorageDomainService) String() string {
    return fmt.Sprintf("AttachedStorageDomainService:%s", op.Path)
}


//
// Manages the collection of disks available inside an storage domain that is attached to a data center.
//
type AttachedStorageDomainDisksService struct {
    Service

    DiskService  *DiskService
}

func NewAttachedStorageDomainDisksService(connection *Connection, path string) *AttachedStorageDomainDisksService {
    var result AttachedStorageDomainDisksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Adds or registers a disk.
// IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
// compatibility. It will be removed in the future. To add a new disk use the <<services/disks/methods/add, add>>
// operation of the service that manages the disks of the system. To register an unregistered disk use the
// <<services/attached_storage_domain_disk/methods/register, register>> operation of the service that manages
// that disk.
// This method supports the following parameters:
// `Disk`:: The disk to add or register.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AttachedStorageDomainDisksService) Add (
    disk Disk,
    unregistered bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if unregistered != nil {
        query["unregistered"] = unregistered
    }

    // Send the request
    return op.internalAdd(disk, headers, query, wait)
}

//
// Retrieve the list of disks that are available in the storage domain.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of disks to return. If not specified all the disks are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AttachedStorageDomainDisksService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AttachedStorageDomainDisksService) Service(path string) *AttachedStorageDomainDisksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Disk_service(path)
    return self.Disk_service(path[:index]).Service(path[index + 1:])
}

func (op *AttachedStorageDomainDisksService) String() string {
    return fmt.Sprintf("AttachedStorageDomainDisksService:%s", op.Path)
}


//
//
type AttachedStorageDomainsService struct {
    Service

    StorageDomainService  *StorageDomainService
}

func NewAttachedStorageDomainsService(connection *Connection, path string) *AttachedStorageDomainsService {
    var result AttachedStorageDomainsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *AttachedStorageDomainsService) Add (
    storageDomain StorageDomain,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(storageDomain, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of storage domains to return. If not specified all the storage domains are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AttachedStorageDomainsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AttachedStorageDomainsService) Service(path string) *AttachedStorageDomainsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.StorageDomain_service(path)
    return self.StorageDomain_service(path[:index]).Service(path[index + 1:])
}

func (op *AttachedStorageDomainsService) String() string {
    return fmt.Sprintf("AttachedStorageDomainsService:%s", op.Path)
}


//
//
type BalanceService struct {
    Service

}

func NewBalanceService(connection *Connection, path string) *BalanceService {
    var result BalanceService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *BalanceService) Get (
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *BalanceService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *BalanceService) Service(path string) *BalanceService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *BalanceService) String() string {
    return fmt.Sprintf("BalanceService:%s", op.Path)
}


//
//
type BalancesService struct {
    Service

    BalanceService  *BalanceService
}

func NewBalancesService(connection *Connection, path string) *BalancesService {
    var result BalancesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *BalancesService) Add (
    balance Balance,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(balance, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of balances to return. If not specified all the balances are returned.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *BalancesService) List (
    filter bool,
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *BalancesService) Service(path string) *BalancesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Balance_service(path)
    return self.Balance_service(path[:index]).Service(path[index + 1:])
}

func (op *BalancesService) String() string {
    return fmt.Sprintf("BalancesService:%s", op.Path)
}


//
// A service to manage a bookmark.
//
type BookmarkService struct {
    Service

}

func NewBookmarkService(connection *Connection, path string) *BookmarkService {
    var result BookmarkService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Get a bookmark.
// An example for getting a bookmark:
// [source]
// ----
// GET /ovirt-engine/api/bookmarks/123
// ----
// [source,xml]
// ----
// <bookmark href="/ovirt-engine/api/bookmarks/123" id="123">
//   <name>example_vm</name>
//   <value>vm: name=example*</value>
// </bookmark>
// ----
//
func (op *BookmarkService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Remove a bookmark.
// An example for removing a bookmark:
// [source]
// ----
// DELETE /ovirt-engine/api/bookmarks/123
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *BookmarkService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Update a bookmark.
// An example for updating a bookmark:
// [source]
// ----
// PUT /ovirt-engine/api/bookmarks/123
// ----
// With the request body:
// [source,xml]
// ----
// <bookmark>
//   <name>new_example_vm</name>
//   <value>vm: name=new_example*</value>
// </bookmark>
// ----
// This method supports the following parameters:
// `Bookmark`:: The updated bookmark.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *BookmarkService) Update (
    bookmark Bookmark,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(bookmark, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *BookmarkService) Service(path string) *BookmarkService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *BookmarkService) String() string {
    return fmt.Sprintf("BookmarkService:%s", op.Path)
}


//
// A service to manage bookmarks.
//
type BookmarksService struct {
    Service

    BookmarkService  *BookmarkService
}

func NewBookmarksService(connection *Connection, path string) *BookmarksService {
    var result BookmarksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Adding a new bookmark.
// Example of adding a bookmark:
// [source]
// ----
// POST /ovirt-engine/api/bookmarks
// ----
// [source,xml]
// ----
// <bookmark>
//   <name>new_example_vm</name>
//   <value>vm: name=new_example*</value>
// </bookmark>
// ----
// This method supports the following parameters:
// `Bookmark`:: The added bookmark.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *BookmarksService) Add (
    bookmark Bookmark,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(bookmark, headers, query, wait)
}

//
// Listing all the available bookmarks.
// Example of listing bookmarks:
// [source]
// ----
// GET /ovirt-engine/api/bookmarks
// ----
// [source,xml]
// ----
// <bookmarks>
//   <bookmark href="/ovirt-engine/api/bookmarks/123" id="123">
//     <name>database</name>
//     <value>vm: name=database*</value>
//   </bookmark>
//   <bookmark href="/ovirt-engine/api/bookmarks/456" id="456">
//     <name>example</name>
//     <value>vm: name=example*</value>
//   </bookmark>
// </bookmarks>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of bookmarks to return. If not specified all the bookmarks are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *BookmarksService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *BookmarksService) Service(path string) *BookmarksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Bookmark_service(path)
    return self.Bookmark_service(path[:index]).Service(path[index + 1:])
}

func (op *BookmarksService) String() string {
    return fmt.Sprintf("BookmarksService:%s", op.Path)
}


//
// A service to manage specific cluster.
//
type ClusterService struct {
    Service

    AffinityGroupsService  *AffinityGroupsService
    CpuProfilesService  *CpuProfilesService
    GlusterHooksService  *GlusterHooksService
    GlusterVolumesService  *GlusterVolumesService
    NetworkFiltersService  *NetworkFiltersService
    NetworksService  *NetworksService
    PermissionsService  *PermissionsService
}

func NewClusterService(connection *Connection, path string) *ClusterService {
    var result ClusterService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Get information about the cluster.
// An example of getting a cluster:
// [source]
// ----
// GET /ovirt-engine/api/clusters/123
// ----
// [source,xml]
// ----
// <cluster href="/ovirt-engine/api/clusters/123" id="123">
//   <actions>
//     <link href="/ovirt-engine/api/clusters/123/resetemulatedmachine" rel="resetemulatedmachine"/>
//   </actions>
//   <name>Default</name>
//   <description>The default server cluster</description>
//   <link href="/ovirt-engine/api/clusters/123/networks" rel="networks"/>
//   <link href="/ovirt-engine/api/clusters/123/permissions" rel="permissions"/>
//   <link href="/ovirt-engine/api/clusters/123/glustervolumes" rel="glustervolumes"/>
//   <link href="/ovirt-engine/api/clusters/123/glusterhooks" rel="glusterhooks"/>
//   <link href="/ovirt-engine/api/clusters/123/affinitygroups" rel="affinitygroups"/>
//   <link href="/ovirt-engine/api/clusters/123/cpuprofiles" rel="cpuprofiles"/>
//   <ballooning_enabled>false</ballooning_enabled>
//   <cpu>
//     <architecture>x86_64</architecture>
//     <type>Intel Penryn Family</type>
//   </cpu>
//   <error_handling>
//     <on_error>migrate</on_error>
//   </error_handling>
//   <fencing_policy>
//     <enabled>true</enabled>
//     <skip_if_connectivity_broken>
//       <enabled>false</enabled>
//       <threshold>50</threshold>
//     </skip_if_connectivity_broken>
//     <skip_if_sd_active>
//       <enabled>false</enabled>
//     </skip_if_sd_active>
//   </fencing_policy>
//   <gluster_service>false</gluster_service>
//   <ha_reservation>false</ha_reservation>
//   <ksm>
//     <enabled>true</enabled>
//     <merge_across_nodes>true</merge_across_nodes>
//   </ksm>
//   <maintenance_reason_required>false</maintenance_reason_required>
//   <memory_policy>
//     <over_commit>
//       <percent>100</percent>
//     </over_commit>
//     <transparent_hugepages>
//       <enabled>true</enabled>
//     </transparent_hugepages>
//   </memory_policy>
//   <migration>
//     <auto_converge>inherit</auto_converge>
//     <bandwidth>
//       <assignment_method>auto</assignment_method>
//     </bandwidth>
//     <compressed>inherit</compressed>
//   </migration>
//   <optional_reason>false</optional_reason>
//   <required_rng_sources>
//     <required_rng_source>random</required_rng_source>
//   </required_rng_sources>
//   <scheduling_policy href="/ovirt-engine/api/schedulingpolicies/456" id="456"/>
//   <threads_as_cores>false</threads_as_cores>
//   <trusted_service>false</trusted_service>
//   <tunnel_migration>false</tunnel_migration>
//   <version>
//     <major>4</major>
//     <minor>0</minor>
//   </version>
//   <virt_service>true</virt_service>
//   <data_center href="/ovirt-engine/api/datacenters/111" id="111"/>
// </cluster>
// ----
// This method supports the following parameters:
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *ClusterService) Get (
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes cluster from the system.
// [source]
// ----
// DELETE /ovirt-engine/api/clusters/00000000-0000-0000-0000-000000000000
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *ClusterService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the reset should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *ClusterService) ResetEmulatedMachine (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'resetemulatedmachine', nil, headers, query, wait)
}

//
// Updates information about the cluster.
// Only specified fields are updated, others remain unchanged.
// E.g. update cluster's CPU:
// [source]
// ----
// PUT /ovirt-engine/api/clusters/123
// ----
// With request body like:
// [source,xml]
// ----
// <cluster>
//   <cpu>
//     <type>Intel Haswell-noTSX Family</type>
//   </cpu>
// </cluster>
// ----
//
func (op *ClusterService) Update (
    cluster Cluster,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(cluster, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ClusterService) Service(path string) *ClusterService {
    if path == nil {
        return op
    }
    if path == "affinitygroups" {
        return &AffinityGroupsService{}
    }
    if path.startswith('affinitygroups/') {
        return AffinityGroupsService().service(path[15:])
    }
    if path == "cpuprofiles" {
        return &CpuProfilesService{}
    }
    if path.startswith('cpuprofiles/') {
        return CpuProfilesService().service(path[12:])
    }
    if path == "glusterhooks" {
        return &GlusterHooksService{}
    }
    if path.startswith('glusterhooks/') {
        return GlusterHooksService().service(path[13:])
    }
    if path == "glustervolumes" {
        return &GlusterVolumesService{}
    }
    if path.startswith('glustervolumes/') {
        return GlusterVolumesService().service(path[15:])
    }
    if path == "networkfilters" {
        return &NetworkFiltersService{}
    }
    if path.startswith('networkfilters/') {
        return NetworkFiltersService().service(path[15:])
    }
    if path == "networks" {
        return &NetworksService{}
    }
    if path.startswith('networks/') {
        return NetworksService().service(path[9:])
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *ClusterService) String() string {
    return fmt.Sprintf("ClusterService:%s", op.Path)
}


//
// Provides information about a specific cluster level. See the <<services/cluster_levels,ClusterLevels>> service for
// more information.
//
type ClusterLevelService struct {
    Service

}

func NewClusterLevelService(connection *Connection, path string) *ClusterLevelService {
    var result ClusterLevelService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Provides the information about the capabilities of the specific cluster level managed by this service.
// For example, to find what CPU types are supported by level 3.6 you can send a request like this:
// [source]
// ----
// GET /ovirt-engine/api/clusterlevels/3.6
// ----
// That will return a <<types/cluster_level, ClusterLevel>> object containing the supported CPU types, and other
// information which describes the cluster level:
// [source,xml]
// ----
// <cluster_level id="3.6">
//   <cpu_types>
//     <cpu_type>
//       <name>Intel Conroe Family</name>
//       <level>3</level>
//       <architecture>x86_64</architecture>
//     </cpu_type>
//     ...
//   </cpu_types>
//   <permits>
//     <permit id="1">
//       <name>create_vm</name>
//       <administrative>false</administrative>
//     </permit>
//     ...
//   </permits>
// </cluster_level>
// ----
//
func (op *ClusterLevelService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ClusterLevelService) Service(path string) *ClusterLevelService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *ClusterLevelService) String() string {
    return fmt.Sprintf("ClusterLevelService:%s", op.Path)
}


//
// Provides information about the capabilities of different cluster levels supported by the engine. Version 4.0 of the
// engine supports levels 4.0 and 3.6. Each of these levels support different sets of CPU types, for example. This
// service provides that information.
//
type ClusterLevelsService struct {
    Service

    LevelService  *LevelService
}

func NewClusterLevelsService(connection *Connection, path string) *ClusterLevelsService {
    var result ClusterLevelsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Lists the cluster levels supported by the system.
// [source]
// ----
// GET /ovirt-engine/api/clusterlevels
// ----
// This will return a list of available cluster levels.
// [source,xml]
// ----
// <cluster_levels>
//   <cluster_level id="4.0">
//      ...
//   </cluster_level>
//   ...
// </cluster_levels>
// ----
//
func (op *ClusterLevelsService) List (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ClusterLevelsService) Service(path string) *ClusterLevelsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Level_service(path)
    return self.Level_service(path[:index]).Service(path[index + 1:])
}

func (op *ClusterLevelsService) String() string {
    return fmt.Sprintf("ClusterLevelsService:%s", op.Path)
}


//
// A service to manage clusters.
//
type ClustersService struct {
    Service

    ClusterService  *ClusterService
}

func NewClustersService(connection *Connection, path string) *ClustersService {
    var result ClustersService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a new cluster.
// This requires the `name`, `cpu.type` and `data_center` attributes. Identify the data center with either the `id`
// or `name` attributes.
// [source]
// ----
// POST /ovirt-engine/api/clusters
// ----
// With a request body like this:
// [source,xml]
// ----
// <cluster>
//   <name>mycluster</name>
//   <cpu>
//     <type>Intel Penryn Family</type>
//   </cpu>
//   <data_center id="123"/>
// </cluster>
// ----
//
func (op *ClustersService) Add (
    cluster Cluster,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(cluster, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of clusters to return. If not specified all the clusters are returned.
// `Search`:: A query string used to restrict the returned clusters.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *ClustersService) List (
    caseSensitive bool,
    filter bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if filter != nil {
        query["filter"] = filter
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ClustersService) Service(path string) *ClustersService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Cluster_service(path)
    return self.Cluster_service(path[:index]).Service(path[index + 1:])
}

func (op *ClustersService) String() string {
    return fmt.Sprintf("ClustersService:%s", op.Path)
}


//
//
type CopyableService struct {
    Service

}

func NewCopyableService(connection *Connection, path string) *CopyableService {
    var result CopyableService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the copy should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *CopyableService) Copy (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'copy', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *CopyableService) Service(path string) *CopyableService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *CopyableService) String() string {
    return fmt.Sprintf("CopyableService:%s", op.Path)
}


//
//
type CpuProfileService struct {
    Service

    PermissionsService  *PermissionsService
}

func NewCpuProfileService(connection *Connection, path string) *CpuProfileService {
    var result CpuProfileService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *CpuProfileService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *CpuProfileService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
//
func (op *CpuProfileService) Update (
    profile CpuProfile,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(profile, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *CpuProfileService) Service(path string) *CpuProfileService {
    if path == nil {
        return op
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *CpuProfileService) String() string {
    return fmt.Sprintf("CpuProfileService:%s", op.Path)
}


//
//
type CpuProfilesService struct {
    Service

    ProfileService  *ProfileService
}

func NewCpuProfilesService(connection *Connection, path string) *CpuProfilesService {
    var result CpuProfilesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *CpuProfilesService) Add (
    profile CpuProfile,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(profile, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of profiles to return. If not specified all the profiles are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *CpuProfilesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *CpuProfilesService) Service(path string) *CpuProfilesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Profile_service(path)
    return self.Profile_service(path[:index]).Service(path[index + 1:])
}

func (op *CpuProfilesService) String() string {
    return fmt.Sprintf("CpuProfilesService:%s", op.Path)
}


//
// A service to manage a data center.
//
type DataCenterService struct {
    Service

    ClustersService  *ClustersService
    IscsiBondsService  *IscsiBondsService
    NetworksService  *NetworksService
    PermissionsService  *PermissionsService
    QossService  *QossService
    QuotasService  *QuotasService
    StorageDomainsService  *StorageDomainsService
}

func NewDataCenterService(connection *Connection, path string) *DataCenterService {
    var result DataCenterService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Get a data center.
// An example of getting a data center:
// [source]
// ----
// GET /ovirt-engine/api/datacenters/123
// ----
// [source,xml]
// ----
// <data_center href="/ovirt-engine/api/datacenters/123" id="123">
//   <name>Default</name>
//   <description>The default Data Center</description>
//   <link href="/ovirt-engine/api/datacenters/123/clusters" rel="clusters"/>
//   <link href="/ovirt-engine/api/datacenters/123/storagedomains" rel="storagedomains"/>
//   <link href="/ovirt-engine/api/datacenters/123/permissions" rel="permissions"/>
//   <link href="/ovirt-engine/api/datacenters/123/networks" rel="networks"/>
//   <link href="/ovirt-engine/api/datacenters/123/quotas" rel="quotas"/>
//   <link href="/ovirt-engine/api/datacenters/123/qoss" rel="qoss"/>
//   <link href="/ovirt-engine/api/datacenters/123/iscsibonds" rel="iscsibonds"/>
//   <local>false</local>
//   <quota_mode>disabled</quota_mode>
//   <status>up</status>
//   <storage_format>v3</storage_format>
//   <supported_versions>
//     <version>
//       <major>4</major>
//       <minor>0</minor>
//    </version>
//   </supported_versions>
//   <version>
//     <major>4</major>
//     <minor>0</minor>
//   </version>
//   <mac_pool href="/ovirt-engine/api/macpools/456" id="456"/>
// </data_center>
// ----
// This method supports the following parameters:
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DataCenterService) Get (
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes the data center.
// [source]
// ----
// DELETE /ovirt-engine/api/datacenters/123
// ----
// Without any special parameters, the storage domains attached to the data center are detached and then removed
// from the storage. If something fails when performing this operation, for example if there is no host available to
// remove the storage domains from the storage, the complete operation will fail.
// If the `force` parameter is `true` then the operation will always succeed, even if something fails while removing
// one storage domain, for example. The failure is just ignored and the data center is removed from the database
// anyway.
// This method supports the following parameters:
// `Force`:: Indicates if the operation should succeed, and the storage domain removed from the database, even if
// something fails during the operation.
// This parameter is optional, and the default value is `false`.
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DataCenterService) Remove (
    force bool,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if force != nil {
        query["force"] = force
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates the data center.
// The `name`, `description`, `storage_type`, `version`, `storage_format` and `mac_pool` elements are updatable
// post-creation. For example, to change the name and description of data center `123` send a request like this:
// [source]
// ----
// PUT /ovirt-engine/api/datacenters/123
// ----
// With a request body like this:
// [source,xml]
// ----
// <data_center>
//   <name>myupdatedname</name>
//   <description>An updated description for the data center</description>
// </data_center>
// ----
// This method supports the following parameters:
// `DataCenter`:: The data center that is being updated.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DataCenterService) Update (
    dataCenter DataCenter,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(dataCenter, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DataCenterService) Service(path string) *DataCenterService {
    if path == nil {
        return op
    }
    if path == "clusters" {
        return &ClustersService{}
    }
    if path.startswith('clusters/') {
        return ClustersService().service(path[9:])
    }
    if path == "iscsibonds" {
        return &IscsiBondsService{}
    }
    if path.startswith('iscsibonds/') {
        return IscsiBondsService().service(path[11:])
    }
    if path == "networks" {
        return &NetworksService{}
    }
    if path.startswith('networks/') {
        return NetworksService().service(path[9:])
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    if path == "qoss" {
        return &QossService{}
    }
    if path.startswith('qoss/') {
        return QossService().service(path[5:])
    }
    if path == "quotas" {
        return &QuotasService{}
    }
    if path.startswith('quotas/') {
        return QuotasService().service(path[7:])
    }
    if path == "storagedomains" {
        return &StorageDomainsService{}
    }
    if path.startswith('storagedomains/') {
        return StorageDomainsService().service(path[15:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *DataCenterService) String() string {
    return fmt.Sprintf("DataCenterService:%s", op.Path)
}


//
// A service to manage data centers.
//
type DataCentersService struct {
    Service

    DataCenterService  *DataCenterService
}

func NewDataCentersService(connection *Connection, path string) *DataCentersService {
    var result DataCentersService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a new data center.
// Creation of a new data center requires the `name` and `local` elements. For example, to create a data center
// named `mydc` that uses shared storage (NFS, iSCSI or fibre channel) send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/datacenters
// ----
// With a request body like this:
// [source,xml]
// ----
// <data_center>
//   <name>mydc</name>
//   <local>false</local>
// </data_center>
// ----
// This method supports the following parameters:
// `DataCenter`:: The data center that is being added.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DataCentersService) Add (
    dataCenter DataCenter,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(dataCenter, headers, query, wait)
}

//
// Lists the data centers.
// The following request retrieves a representation of the data centers:
// [source]
// ----
// GET /ovirt-engine/api/datacenters
// ----
// The above request performed with `curl`:
// [source,bash]
// ----
// curl \
// --request GET \
// --cacert /etc/pki/ovirt-engine/ca.pem \
// --header "Version: 4" \
// --header "Accept: application/xml" \
// --user "admin@internal:mypassword" \
// https://myengine.example.com/ovirt-engine/api/datacenters
// ----
// This is what an example response could look like:
// [source,xml]
// ----
// <data_center href="/ovirt-engine/api/datacenters/123" id="123">
//   <name>Default</name>
//   <description>The default Data Center</description>
//   <link href="/ovirt-engine/api/datacenters/123/networks" rel="networks"/>
//   <link href="/ovirt-engine/api/datacenters/123/storagedomains" rel="storagedomains"/>
//   <link href="/ovirt-engine/api/datacenters/123/permissions" rel="permissions"/>
//   <link href="/ovirt-engine/api/datacenters/123/clusters" rel="clusters"/>
//   <link href="/ovirt-engine/api/datacenters/123/qoss" rel="qoss"/>
//   <link href="/ovirt-engine/api/datacenters/123/iscsibonds" rel="iscsibonds"/>
//   <link href="/ovirt-engine/api/datacenters/123/quotas" rel="quotas"/>
//   <local>false</local>
//   <quota_mode>disabled</quota_mode>
//   <status>up</status>
//   <supported_versions>
//     <version>
//       <major>4</major>
//       <minor>0</minor>
//     </version>
//   </supported_versions>
//   <version>
//     <major>4</major>
//     <minor>0</minor>
//   </version>
// </data_center>
// ----
// Note the `id` code of your `Default` data center. This code identifies this data center in relation to other
// resources of your virtual environment.
// The data center also contains a link to the storage domains collection. The data center uses this collection to
// attach storage domains from the storage domains main collection.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of data centers to return. If not specified all the data centers are returned.
// `Search`:: A query string used to restrict the returned data centers.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DataCentersService) List (
    caseSensitive bool,
    filter bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if filter != nil {
        query["filter"] = filter
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DataCentersService) Service(path string) *DataCentersService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.DataCenter_service(path)
    return self.DataCenter_service(path[:index]).Service(path[index + 1:])
}

func (op *DataCentersService) String() string {
    return fmt.Sprintf("DataCentersService:%s", op.Path)
}


//
// This service manages the attachment of a disk to a virtual machine.
//
type DiskAttachmentService struct {
    Service

}

func NewDiskAttachmentService(connection *Connection, path string) *DiskAttachmentService {
    var result DiskAttachmentService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Returns the details of the attachment, including the bootable flag and link to the disk.
// An example of getting a disk attachment:
// [source]
// ----
// GET /ovirt-engine/api/vms/123/diskattachments/456
// ----
// [source,xml]
// ----
// <disk_attachment href="/ovirt-engine/api/vms/123/diskattachments/456" id="456">
//   <active>true</active>
//   <bootable>true</bootable>
//   <interface>virtio</interface>
//   <disk href="/ovirt-engine/api/disks/456" id="456"/>
//   <vm href="/ovirt-engine/api/vms/123" id="123"/>
// </disk_attachment>
// ----
//
func (op *DiskAttachmentService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes the disk attachment.
// This will only detach the disk from the virtual machine, but won't remove it from
// the system, unless the `detach_only` parameter is `false`.
// An example of removing a disk attachment:
// [source]
// ----
// DELETE /ovirt-engine/api/vms/123/diskattachments/456?detach_only=true
// ----
// This method supports the following parameters:
// `DetachOnly`:: Indicates if the disk should only be detached from the virtual machine, but not removed from the system.
// The default value is `true`, which won't remove the disk from the system.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DiskAttachmentService) Remove (
    detachOnly bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if detachOnly != nil {
        query["detach_only"] = detachOnly
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Update the disk attachment and the disk properties within it.
// [source]
// ----
// PUT /vms/{vm:id}/disksattachments/{attachment:id}
// <disk_attachment>
//   <bootable>true</bootable>
//   <interface>ide</interface>
//   <active>true</active>
//   <disk>
//     <name>mydisk</name>
//     <provisioned_size>1024</provisioned_size>
//     ...
//   </disk>
// </disk_attachment>
// ----
//
func (op *DiskAttachmentService) Update (
    diskAttachment DiskAttachment,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalUpdate(diskAttachment, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DiskAttachmentService) Service(path string) *DiskAttachmentService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *DiskAttachmentService) String() string {
    return fmt.Sprintf("DiskAttachmentService:%s", op.Path)
}


//
// This service manages the set of disks attached to a virtual machine. Each attached disk is represented by a
// <<types/disk_attachment,DiskAttachment>>, containing the bootable flag, the disk interface and the reference to
// the disk.
//
type DiskAttachmentsService struct {
    Service

    AttachmentService  *AttachmentService
}

func NewDiskAttachmentsService(connection *Connection, path string) *DiskAttachmentsService {
    var result DiskAttachmentsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Adds a new disk attachment to the virtual machine. The `attachment` parameter can contain just a reference, if
// the disk already exists:
// [source,xml]
// ----
// <disk_attachment>
//   <bootable>true</bootable>
//   <pass_discard>true</pass_discard>
//   <interface>ide</interface>
//   <active>true</active>
//   <disk id="123"/>
// </disk_attachment>
// ----
// Or it can contain the complete representation of the disk, if the disk doesn't exist yet:
// [source,xml]
// ----
// <disk_attachment>
//   <bootable>true</bootable>
//   <pass_discard>true</pass_discard>
//   <interface>ide</interface>
//   <active>true</active>
//   <disk>
//     <name>mydisk</name>
//     <provisioned_size>1024</provisioned_size>
//     ...
//   </disk>
// </disk_attachment>
// ----
// In this case the disk will be created and then attached to the virtual machine.
// In both cases, use the following URL for a virtual machine with an id `345`:
// [source]
// ----
// POST /ovirt-engine/api/vms/345/diskattachments
// ----
// IMPORTANT: The server accepts requests that don't contain the `active` attribute, but the effect is
// undefined. In some cases the disk will be automatically activated and in other cases it won't. To
// avoid issues it is strongly recommended to always include the `active` attribute with the desired
// value.
//
func (op *DiskAttachmentsService) Add (
    attachment DiskAttachment,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(attachment, headers, query, wait)
}

//
// List the disk that are attached to the virtual machine.
//
func (op *DiskAttachmentsService) List (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DiskAttachmentsService) Service(path string) *DiskAttachmentsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Attachment_service(path)
    return self.Attachment_service(path[:index]).Service(path[index + 1:])
}

func (op *DiskAttachmentsService) String() string {
    return fmt.Sprintf("DiskAttachmentsService:%s", op.Path)
}


//
//
type DiskProfileService struct {
    Service

    PermissionsService  *PermissionsService
}

func NewDiskProfileService(connection *Connection, path string) *DiskProfileService {
    var result DiskProfileService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *DiskProfileService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DiskProfileService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
//
func (op *DiskProfileService) Update (
    profile DiskProfile,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(profile, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DiskProfileService) Service(path string) *DiskProfileService {
    if path == nil {
        return op
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *DiskProfileService) String() string {
    return fmt.Sprintf("DiskProfileService:%s", op.Path)
}


//
//
type DiskProfilesService struct {
    Service

    DiskProfileService  *DiskProfileService
}

func NewDiskProfilesService(connection *Connection, path string) *DiskProfilesService {
    var result DiskProfilesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *DiskProfilesService) Add (
    profile DiskProfile,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(profile, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of profiles to return. If not specified all the profiles are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DiskProfilesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DiskProfilesService) Service(path string) *DiskProfilesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.DiskProfile_service(path)
    return self.DiskProfile_service(path[:index]).Service(path[index + 1:])
}

func (op *DiskProfilesService) String() string {
    return fmt.Sprintf("DiskProfilesService:%s", op.Path)
}


//
//
type DiskSnapshotService struct {
    Service

}

func NewDiskSnapshotService(connection *Connection, path string) *DiskSnapshotService {
    var result DiskSnapshotService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *DiskSnapshotService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DiskSnapshotService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DiskSnapshotService) Service(path string) *DiskSnapshotService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *DiskSnapshotService) String() string {
    return fmt.Sprintf("DiskSnapshotService:%s", op.Path)
}


//
//
type DiskSnapshotsService struct {
    Service

    SnapshotService  *SnapshotService
}

func NewDiskSnapshotsService(connection *Connection, path string) *DiskSnapshotsService {
    var result DiskSnapshotsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of snapshots to return. If not specified all the snapshots are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DiskSnapshotsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DiskSnapshotsService) Service(path string) *DiskSnapshotsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Snapshot_service(path)
    return self.Snapshot_service(path[:index]).Service(path[index + 1:])
}

func (op *DiskSnapshotsService) String() string {
    return fmt.Sprintf("DiskSnapshotsService:%s", op.Path)
}


//
// Manages the collection of disks available in the system.
//
type DisksService struct {
    Service

    DiskService  *DiskService
}

func NewDisksService(connection *Connection, path string) *DisksService {
    var result DisksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Adds a new floating disk.
// There are three types of disks that can be added - disk image, direct LUN and
//  https://wiki.openstack.org/wiki/Cinder[Cinder] disk.
// *Adding a new image disk:*
// When creating a new floating image <<types/disk,Disk>>, the API requires the `storage_domain`, `provisioned_size`
// and `format` attributes.
// To create a new floating image disk with specified `provisioned_size`, `format` and `name` on a storage domain
// with an id `123`, send a request as follows:
// [source]
// ----
// POST /ovirt-engine/api/disks
// ----
// With a request body as follows:
// [source,xml]
// ----
// <disk>
//   <storage_domains>
//     <storage_domain id="123"/>
//   </storage_domains>
//   <name>mydisk</name>
//   <provisioned_size>1048576</provisioned_size>
//   <format>cow</format>
// </disk>
// ----
// *Adding a new direct LUN disk:*
// When adding a new floating direct LUN via the API, there are two flavors that can be used:
// . With a `host` element - in this case, the host is used for sanity checks (e.g., that the LUN is visible) and
// to retrieve basic information about the LUN (e.g., size and serial).
// . Without a `host` element - in this case, the operation is a database-only operation, and the storage is never
// accessed.
// To create a new floating direct LUN disk with a `host` element with an id `123`, specified `alias`, `type` and
// `logical_unit` with an id `456` (that has the attributes `address`, `port` and `target`),
// send a request as follows:
// [source]
// ----
// POST /ovirt-engine/api/disks
// ----
// With a request body as follows:
// [source,xml]
// ----
// <disk>
//   <alias>mylun</alias>
//   <lun_storage>
//     <host id="123"/>
//     <type>iscsi</type>
//     <logical_units>
//       <logical_unit id="456">
//         <address>10.35.10.20</address>
//         <port>3260</port>
//         <target>iqn.2017-01.com.myhost:444</target>
//       </logical_unit>
//     </logical_units>
//   </lun_storage>
// </disk>
// ----
// To create a new floating direct LUN disk without using a host, remove the `host` element.
// *Adding a new Cinder disk:*
// To create a new floating Cinder disk, send a request as follows:
// [source]
// ----
// POST /ovirt-engine/api/disks
// ----
// With a request body as follows:
// [source,xml]
// ----
// <disk>
//   <openstack_volume_type>
//     <name>myceph</name>
//   </openstack_volume_type>
//   <storage_domains>
//     <storage_domain>
//       <name>cinderDomain</name>
//     </storage_domain>
//   </storage_domains>
//   <provisioned_size>1073741824</provisioned_size>
//   <interface>virtio</interface>
//   <format>raw</format>
// </disk>
// ----
// This method supports the following parameters:
// `Disk`:: The disk.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DisksService) Add (
    disk Disk,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(disk, headers, query, wait)
}

//
// Get list of disks.
// [source]
// ----
// GET /ovirt-engine/api/disks
// ----
// You will get a XML response which will look like this one:
// [source,xml]
// ----
// <disks>
//   <disk id="123">
//     <actions>...</actions>
//     <name>MyDisk</name>
//     <description>MyDisk description</description>
//     <link href="/ovirt-engine/api/disks/123/permissions" rel="permissions"/>
//     <link href="/ovirt-engine/api/disks/123/statistics" rel="statistics"/>
//     <actual_size>5345845248</actual_size>
//     <alias>MyDisk alias</alias>
//     ...
//     <status>ok</status>
//     <storage_type>image</storage_type>
//     <wipe_after_delete>false</wipe_after_delete>
//     <disk_profile id="123"/>
//     <quota id="123"/>
//     <storage_domains>...</storage_domains>
//   </disk>
//   ...
// </disks>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of disks to return. If not specified all the disks are returned.
// `Search`:: A query string used to restrict the returned disks.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DisksService) List (
    caseSensitive bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DisksService) Service(path string) *DisksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Disk_service(path)
    return self.Disk_service(path[:index]).Service(path[index + 1:])
}

func (op *DisksService) String() string {
    return fmt.Sprintf("DisksService:%s", op.Path)
}


//
// A service to view details of an authentication domain in the system.
//
type DomainService struct {
    Service

    GroupsService  *GroupsService
    UsersService  *UsersService
}

func NewDomainService(connection *Connection, path string) *DomainService {
    var result DomainService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Gets the authentication domain information.
// Usage:
// ....
// GET /ovirt-engine/api/domains/5678
// ....
// Will return the domain information:
// [source,xml]
// ----
// <domain href="/ovirt-engine/api/domains/5678" id="5678">
//   <name>internal-authz</name>
//   <link href="/ovirt-engine/api/domains/5678/users" rel="users"/>
//   <link href="/ovirt-engine/api/domains/5678/groups" rel="groups"/>
//   <link href="/ovirt-engine/api/domains/5678/users?search={query}" rel="users/search"/>
//   <link href="/ovirt-engine/api/domains/5678/groups?search={query}" rel="groups/search"/>
// </domain>
// ----
//
func (op *DomainService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DomainService) Service(path string) *DomainService {
    if path == nil {
        return op
    }
    if path == "groups" {
        return &GroupsService{}
    }
    if path.startswith('groups/') {
        return GroupsService().service(path[7:])
    }
    if path == "users" {
        return &UsersService{}
    }
    if path.startswith('users/') {
        return UsersService().service(path[6:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *DomainService) String() string {
    return fmt.Sprintf("DomainService:%s", op.Path)
}


//
//
type DomainGroupService struct {
    Service

}

func NewDomainGroupService(connection *Connection, path string) *DomainGroupService {
    var result DomainGroupService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *DomainGroupService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DomainGroupService) Service(path string) *DomainGroupService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *DomainGroupService) String() string {
    return fmt.Sprintf("DomainGroupService:%s", op.Path)
}


//
//
type DomainGroupsService struct {
    Service

    GroupService  *GroupService
}

func NewDomainGroupsService(connection *Connection, path string) *DomainGroupsService {
    var result DomainGroupsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of groups to return. If not specified all the groups are returned.
// `Search`:: A query string used to restrict the returned groups.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DomainGroupsService) List (
    caseSensitive bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DomainGroupsService) Service(path string) *DomainGroupsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Group_service(path)
    return self.Group_service(path[:index]).Service(path[index + 1:])
}

func (op *DomainGroupsService) String() string {
    return fmt.Sprintf("DomainGroupsService:%s", op.Path)
}


//
// A service to view a domain user in the system.
//
type DomainUserService struct {
    Service

}

func NewDomainUserService(connection *Connection, path string) *DomainUserService {
    var result DomainUserService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Gets the domain user information.
// Usage:
// ....
// GET /ovirt-engine/api/domains/5678/users/1234
// ....
// Will return the domain user information:
// [source,xml]
// ----
// <user href="/ovirt-engine/api/users/1234" id="1234">
//   <name>admin</name>
//   <namespace>*</namespace>
//   <principal>admin</principal>
//   <user_name>admin@internal-authz</user_name>
//   <domain href="/ovirt-engine/api/domains/5678" id="5678">
//     <name>internal-authz</name>
//   </domain>
//   <groups/>
// </user>
// ----
//
func (op *DomainUserService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DomainUserService) Service(path string) *DomainUserService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *DomainUserService) String() string {
    return fmt.Sprintf("DomainUserService:%s", op.Path)
}


//
// A service to list all domain users in the system.
//
type DomainUsersService struct {
    Service

    UserService  *UserService
}

func NewDomainUsersService(connection *Connection, path string) *DomainUsersService {
    var result DomainUsersService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// List all the users in the domain.
// Usage:
// ....
// GET /ovirt-engine/api/domains/5678/users
// ....
// Will return the list of users in the domain:
// [source,xml]
// ----
// <users>
//   <user href="/ovirt-engine/api/domains/5678/users/1234" id="1234">
//     <name>admin</name>
//     <namespace>*</namespace>
//     <principal>admin</principal>
//     <user_name>admin@internal-authz</user_name>
//     <domain href="/ovirt-engine/api/domains/5678" id="5678">
//       <name>internal-authz</name>
//     </domain>
//     <groups/>
//   </user>
// </users>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of users to return. If not specified all the users are returned.
// `Search`:: A query string used to restrict the returned users.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DomainUsersService) List (
    caseSensitive bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DomainUsersService) Service(path string) *DomainUsersService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.User_service(path)
    return self.User_service(path[:index]).Service(path[index + 1:])
}

func (op *DomainUsersService) String() string {
    return fmt.Sprintf("DomainUsersService:%s", op.Path)
}


//
// A service to list all authentication domains in the system.
//
type DomainsService struct {
    Service

    DomainService  *DomainService
}

func NewDomainsService(connection *Connection, path string) *DomainsService {
    var result DomainsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// List all the authentication domains in the system.
// Usage:
// ....
// GET /ovirt-engine/api/domains
// ....
// Will return the list of domains:
// [source,xml]
// ----
// <domains>
//   <domain href="/ovirt-engine/api/domains/5678" id="5678">
//     <name>internal-authz</name>
//     <link href="/ovirt-engine/api/domains/5678/users" rel="users"/>
//     <link href="/ovirt-engine/api/domains/5678/groups" rel="groups"/>
//     <link href="/ovirt-engine/api/domains/5678/users?search={query}" rel="users/search"/>
//     <link href="/ovirt-engine/api/domains/5678/groups?search={query}" rel="groups/search"/>
//   </domain>
// </domains>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of domains to return. If not specified all the domains are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DomainsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DomainsService) Service(path string) *DomainsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Domain_service(path)
    return self.Domain_service(path[:index]).Service(path[index + 1:])
}

func (op *DomainsService) String() string {
    return fmt.Sprintf("DomainsService:%s", op.Path)
}


//
// A service to manage an event in the system.
//
type EventService struct {
    Service

}

func NewEventService(connection *Connection, path string) *EventService {
    var result EventService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Get an event.
// An example of getting an event:
// [source]
// ----
// GET /ovirt-engine/api/events/123
// ----
// [source,xml]
// ----
// <event href="/ovirt-engine/api/events/123" id="123">
//   <description>Host example.com was added by admin@internal-authz.</description>
//   <code>42</code>
//   <correlation_id>135</correlation_id>
//   <custom_id>-1</custom_id>
//   <flood_rate>30</flood_rate>
//   <origin>oVirt</origin>
//   <severity>normal</severity>
//   <time>2016-12-11T11:13:44.654+02:00</time>
//   <cluster href="/ovirt-engine/api/clusters/456" id="456"/>
//   <host href="/ovirt-engine/api/hosts/789" id="789"/>
//   <user href="/ovirt-engine/api/users/987" id="987"/>
// </event>
// ----
// Note that the number of fields changes according to the information that resides on the event.
// For example, for storage domain related events you will get the storage domain reference,
// as well as the reference for the data center this storage domain resides in.
//
func (op *EventService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes an event from internal audit log.
// An event can be removed by sending following request
// [source]
// ----
// DELETE /ovirt-engine/api/events/123
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *EventService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *EventService) Service(path string) *EventService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *EventService) String() string {
    return fmt.Sprintf("EventService:%s", op.Path)
}


//
// A service to manage events in the system.
//
type EventsService struct {
    Service

    EventService  *EventService
}

func NewEventsService(connection *Connection, path string) *EventsService {
    var result EventsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Adds an external event to the internal audit log.
// This is intended for integration with external systems that detect or produce events relevant for the
// administrator of the system. For example, an external monitoring tool may be able to detect that a file system
// is full inside the guest operating system of a virtual machine. This event can be added to the internal audit
// log sending a request like this:
// [source]
// ----
// POST /ovirt-engine/api/events
// <event>
//   <description>File system /home is full</description>
//   <severity>alert</severity>
//   <origin>mymonitor</origin>
//   <custom_id>1467879754</custom_id>
// </event>
// ----
// Events can also be linked to specific objects. For example, the above event could be linked to the specific
// virtual machine where it happened, using the `vm` link:
// [source]
// ----
// POST /ovirt-engine/api/events
// <event>
//   <description>File system /home is full</description>
//   <severity>alert</severity>
//   <origin>mymonitor</origin>
//   <custom_id>1467879754</custom_id>
//   <vm id="aae98225-5b73-490d-a252-899209af17e9"/>
// </event>
// ----
// NOTE: When using links, like the `vm` in the previous example, only the `id` attribute is accepted. The `name`
// attribute, if provided, is simply ignored.
//
func (op *EventsService) Add (
    event Event,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(event, headers, query, wait)
}

//
// Get list of events.
// [source]
// ----
// GET /ovirt-engine/api/events
// ----
// To the above request we get following response:
// [source,xml]
// ----
// <events>
//   <event href="/ovirt-engine/api/events/2" id="2">
//     <description>User admin@internal-authz logged out.</description>
//     <code>31</code>
//     <correlation_id>1e892ea9</correlation_id>
//     <custom_id>-1</custom_id>
//     <flood_rate>30</flood_rate>
//     <origin>oVirt</origin>
//     <severity>normal</severity>
//     <time>2016-09-14T12:14:34.541+02:00</time>
//     <user href="/ovirt-engine/api/users/57d91d48-00da-0137-0138-000000000244" id="57d91d48-00da-0137-0138-000000000244"/>
//   </event>
//   <event href="/ovirt-engine/api/events/1" id="1">
//     <description>User admin logged in.</description>
//     <code>30</code>
//     <correlation_id>1fbd81f4</correlation_id>
//     <custom_id>-1</custom_id>
//     <flood_rate>30</flood_rate>
//     <origin>oVirt</origin>
//     <severity>normal</severity>
//     <time>2016-09-14T11:54:35.229+02:00</time>
//     <user href="/ovirt-engine/api/users/57d91d48-00da-0137-0138-000000000244" id="57d91d48-00da-0137-0138-000000000244"/>
//   </event>
// </events>
// ----
// The following events occur:
// * id="1" - The API logs in the admin user account.
// * id="2" - The API logs out of the admin user account.
// This method supports the following parameters:
// `From`:: Indicates the identifier of the the first event that should be returned. The identifiers of events are
// strictly increasing, so when this parameter is used only the events with that identifiers equal or greater
// than the given value will be returned. For example, the following request will return only the events
// with identifiers greater or equal than `123`:
// [source]
// ----
// GET /ovirt-engine/api/events?from=123
// ----
// This parameter is optional, and if not specified then the first event returned will be most recently
// generated.
// `Max`:: Sets the maximum number of events to return. If not specified all the events are returned.
// `Search`:: The events service provides search queries similar to other resource services.
// We can search by providing specific severity.
// [source]
// ----
// GET /ovirt-engine/api/events?search=severity%3Dnormal
// ----
// To the above request we get a list of events which severity is equal to `normal`:
// [source,xml]
// ----
// <events>
//   <event href="/ovirt-engine/api/events/2" id="2">
//     <description>User admin@internal-authz logged out.</description>
//     <code>31</code>
//     <correlation_id>1fbd81f4</correlation_id>
//     <custom_id>-1</custom_id>
//     <flood_rate>30</flood_rate>
//     <origin>oVirt</origin>
//     <severity>normal</severity>
//     <time>2016-09-14T11:54:35.229+02:00</time>
//     <user href="/ovirt-engine/api/users/57d91d48-00da-0137-0138-000000000244" id="57d91d48-00da-0137-0138-000000000244"/>
//   </event>
//   <event href="/ovirt-engine/api/events/1" id="1">
//     <description>Affinity Rules Enforcement Manager started.</description>
//     <code>10780</code>
//     <custom_id>-1</custom_id>
//     <flood_rate>30</flood_rate>
//     <origin>oVirt</origin>
//     <severity>normal</severity>
//     <time>2016-09-14T11:52:18.861+02:00</time>
//   </event>
// </events>
// ----
// A virtualization environment generates a large amount of events after
// a period of time. However, the API only displays a default number of
// events for one search query. To display more than the default, the API
// separates results into pages with the page command in a search query.
// The following search query tells the API to paginate results using a
// page value in combination with the sortby clause:
// [source]
// ----
// sortby time asc page 1
// ----
// Below example paginates event resources. The URL-encoded request is:
// [source]
// ----
// GET /ovirt-engine/api/events?search=sortby%20time%20asc%20page%201
// ----
// Increase the page value to view the next page of results.
// [source]
// ----
// GET /ovirt-engine/api/events?search=sortby%20time%20asc%20page%202
// ----
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *EventsService) List (
    caseSensitive bool,
    from int64,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if from != nil {
        query["from"] = from
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the un-delete should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *EventsService) Undelete (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'undelete', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *EventsService) Service(path string) *EventsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Event_service(path)
    return self.Event_service(path[:index]).Service(path[index + 1:])
}

func (op *EventsService) String() string {
    return fmt.Sprintf("EventsService:%s", op.Path)
}


//
//
type ExternalComputeResourceService struct {
    Service

}

func NewExternalComputeResourceService(connection *Connection, path string) *ExternalComputeResourceService {
    var result ExternalComputeResourceService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *ExternalComputeResourceService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ExternalComputeResourceService) Service(path string) *ExternalComputeResourceService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *ExternalComputeResourceService) String() string {
    return fmt.Sprintf("ExternalComputeResourceService:%s", op.Path)
}


//
//
type ExternalComputeResourcesService struct {
    Service

    ResourceService  *ResourceService
}

func NewExternalComputeResourcesService(connection *Connection, path string) *ExternalComputeResourcesService {
    var result ExternalComputeResourcesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of resources to return. If not specified all the resources are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *ExternalComputeResourcesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ExternalComputeResourcesService) Service(path string) *ExternalComputeResourcesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Resource_service(path)
    return self.Resource_service(path[:index]).Service(path[index + 1:])
}

func (op *ExternalComputeResourcesService) String() string {
    return fmt.Sprintf("ExternalComputeResourcesService:%s", op.Path)
}


//
//
type ExternalDiscoveredHostService struct {
    Service

}

func NewExternalDiscoveredHostService(connection *Connection, path string) *ExternalDiscoveredHostService {
    var result ExternalDiscoveredHostService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *ExternalDiscoveredHostService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ExternalDiscoveredHostService) Service(path string) *ExternalDiscoveredHostService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *ExternalDiscoveredHostService) String() string {
    return fmt.Sprintf("ExternalDiscoveredHostService:%s", op.Path)
}


//
//
type ExternalDiscoveredHostsService struct {
    Service

    HostService  *HostService
}

func NewExternalDiscoveredHostsService(connection *Connection, path string) *ExternalDiscoveredHostsService {
    var result ExternalDiscoveredHostsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of hosts to return. If not specified all the hosts are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *ExternalDiscoveredHostsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ExternalDiscoveredHostsService) Service(path string) *ExternalDiscoveredHostsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Host_service(path)
    return self.Host_service(path[:index]).Service(path[index + 1:])
}

func (op *ExternalDiscoveredHostsService) String() string {
    return fmt.Sprintf("ExternalDiscoveredHostsService:%s", op.Path)
}


//
//
type ExternalHostService struct {
    Service

}

func NewExternalHostService(connection *Connection, path string) *ExternalHostService {
    var result ExternalHostService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *ExternalHostService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ExternalHostService) Service(path string) *ExternalHostService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *ExternalHostService) String() string {
    return fmt.Sprintf("ExternalHostService:%s", op.Path)
}


//
//
type ExternalHostGroupService struct {
    Service

}

func NewExternalHostGroupService(connection *Connection, path string) *ExternalHostGroupService {
    var result ExternalHostGroupService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *ExternalHostGroupService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ExternalHostGroupService) Service(path string) *ExternalHostGroupService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *ExternalHostGroupService) String() string {
    return fmt.Sprintf("ExternalHostGroupService:%s", op.Path)
}


//
//
type ExternalHostGroupsService struct {
    Service

    GroupService  *GroupService
}

func NewExternalHostGroupsService(connection *Connection, path string) *ExternalHostGroupsService {
    var result ExternalHostGroupsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of groups to return. If not specified all the groups are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *ExternalHostGroupsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ExternalHostGroupsService) Service(path string) *ExternalHostGroupsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Group_service(path)
    return self.Group_service(path[:index]).Service(path[index + 1:])
}

func (op *ExternalHostGroupsService) String() string {
    return fmt.Sprintf("ExternalHostGroupsService:%s", op.Path)
}


//
//
type ExternalHostProvidersService struct {
    Service

    ProviderService  *ProviderService
}

func NewExternalHostProvidersService(connection *Connection, path string) *ExternalHostProvidersService {
    var result ExternalHostProvidersService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *ExternalHostProvidersService) Add (
    provider ExternalHostProvider,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(provider, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of providers to return. If not specified all the providers are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *ExternalHostProvidersService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ExternalHostProvidersService) Service(path string) *ExternalHostProvidersService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Provider_service(path)
    return self.Provider_service(path[:index]).Service(path[index + 1:])
}

func (op *ExternalHostProvidersService) String() string {
    return fmt.Sprintf("ExternalHostProvidersService:%s", op.Path)
}


//
//
type ExternalHostsService struct {
    Service

    HostService  *HostService
}

func NewExternalHostsService(connection *Connection, path string) *ExternalHostsService {
    var result ExternalHostsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of hosts to return. If not specified all the hosts are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *ExternalHostsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ExternalHostsService) Service(path string) *ExternalHostsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Host_service(path)
    return self.Host_service(path[:index]).Service(path[index + 1:])
}

func (op *ExternalHostsService) String() string {
    return fmt.Sprintf("ExternalHostsService:%s", op.Path)
}


//
//
type ExternalProviderService struct {
    Service

    CertificatesService  *CertificatesService
}

func NewExternalProviderService(connection *Connection, path string) *ExternalProviderService {
    var result ExternalProviderService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *ExternalProviderService) ImportCertificates (
    certificates []*Certificate,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Certificates: certificates,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'importcertificates', nil, headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the test should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *ExternalProviderService) TestConnectivity (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'testconnectivity', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ExternalProviderService) Service(path string) *ExternalProviderService {
    if path == nil {
        return op
    }
    if path == "certificates" {
        return &CertificatesService{}
    }
    if path.startswith('certificates/') {
        return CertificatesService().service(path[13:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *ExternalProviderService) String() string {
    return fmt.Sprintf("ExternalProviderService:%s", op.Path)
}


//
//
type ExternalProviderCertificateService struct {
    Service

}

func NewExternalProviderCertificateService(connection *Connection, path string) *ExternalProviderCertificateService {
    var result ExternalProviderCertificateService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *ExternalProviderCertificateService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ExternalProviderCertificateService) Service(path string) *ExternalProviderCertificateService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *ExternalProviderCertificateService) String() string {
    return fmt.Sprintf("ExternalProviderCertificateService:%s", op.Path)
}


//
//
type ExternalProviderCertificatesService struct {
    Service

    CertificateService  *CertificateService
}

func NewExternalProviderCertificatesService(connection *Connection, path string) *ExternalProviderCertificatesService {
    var result ExternalProviderCertificatesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of certificates to return. If not specified all the certificates are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *ExternalProviderCertificatesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ExternalProviderCertificatesService) Service(path string) *ExternalProviderCertificatesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Certificate_service(path)
    return self.Certificate_service(path[:index]).Service(path[index + 1:])
}

func (op *ExternalProviderCertificatesService) String() string {
    return fmt.Sprintf("ExternalProviderCertificatesService:%s", op.Path)
}


//
// Provides capability to import external virtual machines.
//
type ExternalVmImportsService struct {
    Service

}

func NewExternalVmImportsService(connection *Connection, path string) *ExternalVmImportsService {
    var result ExternalVmImportsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This operation is used to import a virtual machine from external hypervisor, such as KVM, XEN or VMware.
// For example import of a virtual machine from VMware can be facilitated using the following request:
// [source]
// ----
// POST /externalvmimports
// ----
// With request body of type <<types/external_vm_import,ExternalVmImport>>, for example:
// [source,xml]
// ----
// <external_vm_import>
//   <vm>
//     <name>my_vm</name>
//   </vm>
//   <cluster id="360014051136c20574f743bdbd28177fd" />
//   <storage_domain id="8bb5ade5-e988-4000-8b93-dbfc6717fe50" />
//   <name>vm_name_as_is_in_vmware</name>
//   <sparse>true</sparse>
//   <username>vmware_user</username>
//   <password>123456</password>
//   <provider>VMWARE</provider>
//   <url>vpx://wmware_user@vcenter-host/DataCenter/Cluster/esxi-host?no_verify=1</url>
//   <drivers_iso id="virtio-win-1.6.7.iso" />
// </external_vm_import>
// ----
//
func (op *ExternalVmImportsService) Add (
    import ExternalVmImport,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(import, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ExternalVmImportsService) Service(path string) *ExternalVmImportsService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *ExternalVmImportsService) String() string {
    return fmt.Sprintf("ExternalVmImportsService:%s", op.Path)
}


//
//
type FenceAgentService struct {
    Service

}

func NewFenceAgentService(connection *Connection, path string) *FenceAgentService {
    var result FenceAgentService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *FenceAgentService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *FenceAgentService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
//
func (op *FenceAgentService) Update (
    agent Agent,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(agent, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *FenceAgentService) Service(path string) *FenceAgentService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *FenceAgentService) String() string {
    return fmt.Sprintf("FenceAgentService:%s", op.Path)
}


//
//
type FenceAgentsService struct {
    Service

    AgentService  *AgentService
}

func NewFenceAgentsService(connection *Connection, path string) *FenceAgentsService {
    var result FenceAgentsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *FenceAgentsService) Add (
    agent Agent,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(agent, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of agents to return. If not specified all the agents are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *FenceAgentsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *FenceAgentsService) Service(path string) *FenceAgentsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Agent_service(path)
    return self.Agent_service(path[:index]).Service(path[index + 1:])
}

func (op *FenceAgentsService) String() string {
    return fmt.Sprintf("FenceAgentsService:%s", op.Path)
}


//
//
type FileService struct {
    Service

}

func NewFileService(connection *Connection, path string) *FileService {
    var result FileService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *FileService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *FileService) Service(path string) *FileService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *FileService) String() string {
    return fmt.Sprintf("FileService:%s", op.Path)
}


//
// Provides a way for clients to list available files.
// This services is specifically targeted to ISO storage domains, which contain ISO images and virtual floppy disks
// (VFDs) that an administrator uploads.
// The addition of a CDROM device to a virtual machine requires an ISO image from the files of an ISO storage domain.
//
type FilesService struct {
    Service

    FileService  *FileService
}

func NewFilesService(connection *Connection, path string) *FilesService {
    var result FilesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of files to return. If not specified all the files are returned.
// `Search`:: A query string used to restrict the returned files.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *FilesService) List (
    caseSensitive bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *FilesService) Service(path string) *FilesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.File_service(path)
    return self.File_service(path[:index]).Service(path[index + 1:])
}

func (op *FilesService) String() string {
    return fmt.Sprintf("FilesService:%s", op.Path)
}


//
//
type FilterService struct {
    Service

}

func NewFilterService(connection *Connection, path string) *FilterService {
    var result FilterService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *FilterService) Get (
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *FilterService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *FilterService) Service(path string) *FilterService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *FilterService) String() string {
    return fmt.Sprintf("FilterService:%s", op.Path)
}


//
//
type FiltersService struct {
    Service

    FilterService  *FilterService
}

func NewFiltersService(connection *Connection, path string) *FiltersService {
    var result FiltersService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *FiltersService) Add (
    filter Filter,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(filter, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of filters to return. If not specified all the filters are returned.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *FiltersService) List (
    filter bool,
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *FiltersService) Service(path string) *FiltersService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Filter_service(path)
    return self.Filter_service(path[:index]).Service(path[index + 1:])
}

func (op *FiltersService) String() string {
    return fmt.Sprintf("FiltersService:%s", op.Path)
}


//
// This service manages the gluster bricks in a gluster volume
//
type GlusterBricksService struct {
    Service

    BrickService  *BrickService
}

func NewGlusterBricksService(connection *Connection, path string) *GlusterBricksService {
    var result GlusterBricksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Activate the bricks post data migration of remove brick operation.
// Used to activate brick(s) once the data migration from bricks is complete but user no longer wishes to remove
// bricks. The bricks that were previously marked for removal will now be used as normal bricks.
// For example, to retain the bricks that on glustervolume `123` from which data was migrated, send a request like
// this:
// [source]
// ----
// POST /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks/activate
// ----
// With a request body like this:
// [source,xml]
// ----
// <action>
//   <bricks>
//     <brick>
//       <name>host1:/rhgs/brick1</name>
//     </brick>
//   </bricks>
// </action>
// ----
// This method supports the following parameters:
// `Bricks`:: The list of bricks that need to be re-activated.
// `Async`:: Indicates if the activation should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterBricksService) Activate (
    async bool,
    bricks []*GlusterBrick,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Bricks: bricks,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'activate', nil, headers, query, wait)
}

//
// Adds a list of bricks to gluster volume.
// Used to expand a gluster volume by adding bricks. For replicated volume types, the parameter `replica_count`
// needs to be passed. In case the replica count is being increased, then the number of bricks needs to be
// equivalent to the number of replica sets.
// For example, to add bricks to gluster volume `123`, send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks
// ----
// With a request body like this:
// [source,xml]
// ----
// <bricks>
//   <brick>
//     <server_id>111</server_id>
//     <brick_dir>/export/data/brick3</brick_dir>
//   </brick>
// </bricks>
// ----
// This method supports the following parameters:
// `Bricks`:: The list of bricks to be added to the volume
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterBricksService) Add (
    bricks []*GlusterBrick,
    replicaCount int64,
    stripeCount int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if replicaCount != nil {
        query["replica_count"] = replicaCount
    }
    if stripeCount != nil {
        query["stripe_count"] = stripeCount
    }

    // Send the request
    return op.internalAdd(bricks, headers, query, wait)
}

//
// Lists the bricks of a gluster volume.
// For example, to list bricks of gluster volume `123`, send a request like this:
// [source]
// ----
// GET /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks
// ----
// Provides an output as below:
// [source,xml]
// ----
// <bricks>
//   <brick id="234">
//     <name>host1:/rhgs/data/brick1</name>
//     <brick_dir>/rhgs/data/brick1</brick_dir>
//     <server_id>111</server_id>
//     <status>up</status>
//   </brick>
//   <brick id="233">
//     <name>host2:/rhgs/data/brick1</name>
//     <brick_dir>/rhgs/data/brick1</brick_dir>
//     <server_id>222</server_id>
//     <status>up</status>
//   </brick>
// </bricks>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of bricks to return. If not specified all the bricks are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterBricksService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Start migration of data prior to removing bricks.
// Removing bricks is a two-step process, where the data on bricks to be removed, is first migrated to remaining
// bricks. Once migration is completed the removal of bricks is confirmed via the API
// <<services/gluster_bricks/methods/remove, remove>>. If at any point, the action needs to be cancelled
// <<services/gluster_bricks/methods/stop_migrate, stopmigrate>> has to be called.
// For instance, to delete a brick from a gluster volume with id `123`, send a request:
// [source]
// ----
// POST /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks/migrate
// ----
// With a request body like this:
// [source,xml]
// ----
// <action>
//   <bricks>
//     <brick>
//       <name>host1:/rhgs/brick1</name>
//     </brick>
//   </bricks>
// </action>
// ----
// The migration process can be tracked from the job id returned from the API using
// <<services/job/methods/get, job>> and steps in job using <<services/step/methods/get, step>>
// This method supports the following parameters:
// `Bricks`:: List of bricks for which data migration needs to be started.
// `Async`:: Indicates if the migration should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterBricksService) Migrate (
    async bool,
    bricks []*GlusterBrick,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Bricks: bricks,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'migrate', nil, headers, query, wait)
}

//
// Removes bricks from gluster volume.
// The recommended way to remove bricks without data loss is to first migrate the data using
// <<services/gluster_bricks/methods/stop_migrate, stopmigrate>> and then removing them. If migrate was not called on
// bricks prior to remove, the bricks are removed without data migration which may lead to data loss.
// For example, to delete the bricks from gluster volume `123`, send a request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks
// ----
// With a request body like this:
// [source,xml]
// ----
// <bricks>
//   <brick>
//     <name>host:brick_directory</name>
//   </brick>
// </bricks>
// ----
// This method supports the following parameters:
// `Bricks`:: The list of bricks to be removed
// `ReplicaCount`:: Replica count of volume post add operation.
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterBricksService) Remove (
    bricks []*GlusterBrick,
    replicaCount int64,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if bricks != nil {
        query["bricks"] = bricks
    }
    if replicaCount != nil {
        query["replica_count"] = replicaCount
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Stops migration of data from bricks for a remove brick operation.
// To cancel data migration that was started as part of the 2-step remove brick process in case the user wishes to
// continue using the bricks. The bricks that were marked for removal will function as normal bricks post this
// operation.
// For example, to stop migration of data from the bricks of gluster volume `123`, send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks/stopmigrate
// ----
// With a request body like this:
// [source,xml]
// ----
// <bricks>
//   <brick>
//     <name>host:brick_directory</name>
//   </brick>
// </bricks>
// ----
// This method supports the following parameters:
// `Bricks`:: List of bricks for which data migration needs to be stopped. This list should match the arguments passed to
// <<services/gluster_bricks/methods/migrate, migrate>>.
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterBricksService) StopMigrate (
    async bool,
    bricks []*GlusterBrick,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Bricks: bricks,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'stopmigrate', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *GlusterBricksService) Service(path string) *GlusterBricksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Brick_service(path)
    return self.Brick_service(path[:index]).Service(path[index + 1:])
}

func (op *GlusterBricksService) String() string {
    return fmt.Sprintf("GlusterBricksService:%s", op.Path)
}


//
//
type GlusterHookService struct {
    Service

}

func NewGlusterHookService(connection *Connection, path string) *GlusterHookService {
    var result GlusterHookService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Resolves status conflict of hook among servers in cluster by disabling Gluster hook in all servers of the
// cluster. This updates the hook status to `DISABLED` in database.
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterHookService) Disable (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'disable', nil, headers, query, wait)
}

//
// Resolves status conflict of hook among servers in cluster by disabling Gluster hook in all servers of the
// cluster. This updates the hook status to `DISABLED` in database.
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterHookService) Enable (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'enable', nil, headers, query, wait)
}

//
//
func (op *GlusterHookService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes the this Gluster hook from all servers in cluster and deletes it from the database.
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterHookService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Resolves missing hook conflict depending on the resolution type.
// For `ADD` resolves by copying hook stored in engine database to all servers where the hook is missing. The
// engine maintains a list of all servers where hook is missing.
// For `COPY` resolves conflict in hook content by copying hook stored in engine database to all servers where
// the hook is missing. The engine maintains a list of all servers where the content is conflicting. If a host
// id is passed as parameter, the hook content from the server is used as the master to copy to other servers
// in cluster.
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterHookService) Resolve (
    async bool,
    host Host,
    resolutionType string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Host: host,
        ResolutionType: resolutionType,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'resolve', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *GlusterHookService) Service(path string) *GlusterHookService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *GlusterHookService) String() string {
    return fmt.Sprintf("GlusterHookService:%s", op.Path)
}


//
//
type GlusterHooksService struct {
    Service

    HookService  *HookService
}

func NewGlusterHooksService(connection *Connection, path string) *GlusterHooksService {
    var result GlusterHooksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of hooks to return. If not specified all the hooks are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterHooksService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *GlusterHooksService) Service(path string) *GlusterHooksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Hook_service(path)
    return self.Hook_service(path[:index]).Service(path[index + 1:])
}

func (op *GlusterHooksService) String() string {
    return fmt.Sprintf("GlusterHooksService:%s", op.Path)
}


//
// This service manages a collection of gluster volumes available in a cluster.
//
type GlusterVolumesService struct {
    Service

    VolumeService  *VolumeService
}

func NewGlusterVolumesService(connection *Connection, path string) *GlusterVolumesService {
    var result GlusterVolumesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a new gluster volume.
// The volume is created based on properties of the `volume` parameter. The properties `name`, `volume_type` and
// `bricks` are required.
// For example, to add a volume with name `myvolume` to the cluster `123`, send the following request:
// [source]
// ----
// POST /ovirt-engine/api/clusters/123/glustervolumes
// ----
// With the following request body:
// [source,xml]
// ----
// <gluster_volume>
//   <name>myvolume</name>
//   <volume_type>replicate</volume_type>
//   <replica_count>3</replica_count>
//   <bricks>
//     <brick>
//       <server_id>server1</server_id>
//       <brick_dir>/exp1</brick_dir>
//     </brick>
//     <brick>
//       <server_id>server2</server_id>
//       <brick_dir>/exp1</brick_dir>
//     </brick>
//     <brick>
//       <server_id>server3</server_id>
//       <brick_dir>/exp1</brick_dir>
//     </brick>
//   <bricks>
// </gluster_volume>
// ----
// This method supports the following parameters:
// `Volume`:: The gluster volume definition from which to create the volume is passed as input and the newly created
// volume is returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterVolumesService) Add (
    volume GlusterVolume,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(volume, headers, query, wait)
}

//
// Lists all gluster volumes in the cluster.
// For example, to list all Gluster Volumes in cluster `456`, send a request like
// this:
// [source]
// ----
// GET /ovirt-engine/api/clusters/456/glustervolumes
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of volumes to return. If not specified all the volumes are returned.
// `Search`:: A query string used to restrict the returned volumes.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterVolumesService) List (
    caseSensitive bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *GlusterVolumesService) Service(path string) *GlusterVolumesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Volume_service(path)
    return self.Volume_service(path[:index]).Service(path[index + 1:])
}

func (op *GlusterVolumesService) String() string {
    return fmt.Sprintf("GlusterVolumesService:%s", op.Path)
}


//
//
type GroupService struct {
    Service

    PermissionsService  *PermissionsService
    RolesService  *RolesService
    TagsService  *TagsService
}

func NewGroupService(connection *Connection, path string) *GroupService {
    var result GroupService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *GroupService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GroupService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *GroupService) Service(path string) *GroupService {
    if path == nil {
        return op
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    if path == "roles" {
        return &RolesService{}
    }
    if path.startswith('roles/') {
        return RolesService().service(path[6:])
    }
    if path == "tags" {
        return &TagsService{}
    }
    if path.startswith('tags/') {
        return TagsService().service(path[5:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *GroupService) String() string {
    return fmt.Sprintf("GroupService:%s", op.Path)
}


//
//
type GroupsService struct {
    Service

    GroupService  *GroupService
}

func NewGroupsService(connection *Connection, path string) *GroupsService {
    var result GroupsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add group from a directory service. Please note that domain name is name of the authorization provider.
// For example, to add the `Developers` group from the `internal-authz` authorization provider send a request
// like this:
// [source]
// ----
// POST /ovirt-engine/api/groups
// ----
// With a request body like this:
// [source,xml]
// ----
// <group>
//   <name>Developers</name>
//   <domain>
//     <name>internal-authz</name>
//   </domain>
// </group>
// ----
//
func (op *GroupsService) Add (
    group Group,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(group, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of groups to return. If not specified all the groups are returned.
// `Search`:: A query string used to restrict the returned groups.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GroupsService) List (
    caseSensitive bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *GroupsService) Service(path string) *GroupsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Group_service(path)
    return self.Group_service(path[:index]).Service(path[index + 1:])
}

func (op *GroupsService) String() string {
    return fmt.Sprintf("GroupsService:%s", op.Path)
}


//
// A service to access a particular device of a host.
//
type HostDeviceService struct {
    Service

}

func NewHostDeviceService(connection *Connection, path string) *HostDeviceService {
    var result HostDeviceService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieve information about a particular host's device.
// An example of getting a host device:
// [source]
// ----
// GET /ovirt-engine/api/hosts/123/devices/456
// ----
// [source,xml]
// ----
// <host_device href="/ovirt-engine/api/hosts/123/devices/456" id="456">
//   <name>usb_1_9_1_1_0</name>
//   <capability>usb</capability>
//   <host href="/ovirt-engine/api/hosts/123" id="123"/>
//   <parent_device href="/ovirt-engine/api/hosts/123/devices/789" id="789">
//     <name>usb_1_9_1</name>
//   </parent_device>
// </host_device>
// ----
//
func (op *HostDeviceService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *HostDeviceService) Service(path string) *HostDeviceService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *HostDeviceService) String() string {
    return fmt.Sprintf("HostDeviceService:%s", op.Path)
}


//
// A service to access host devices.
//
type HostDevicesService struct {
    Service

    DeviceService  *DeviceService
}

func NewHostDevicesService(connection *Connection, path string) *HostDevicesService {
    var result HostDevicesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// List the devices of a host.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of devices to return. If not specified all the devices are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostDevicesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *HostDevicesService) Service(path string) *HostDevicesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Device_service(path)
    return self.Device_service(path[:index]).Service(path[index + 1:])
}

func (op *HostDevicesService) String() string {
    return fmt.Sprintf("HostDevicesService:%s", op.Path)
}


//
//
type HostHookService struct {
    Service

}

func NewHostHookService(connection *Connection, path string) *HostHookService {
    var result HostHookService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *HostHookService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *HostHookService) Service(path string) *HostHookService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *HostHookService) String() string {
    return fmt.Sprintf("HostHookService:%s", op.Path)
}


//
//
type HostHooksService struct {
    Service

    HookService  *HookService
}

func NewHostHooksService(connection *Connection, path string) *HostHooksService {
    var result HostHooksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of hooks to return. If not specified all the hooks are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostHooksService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *HostHooksService) Service(path string) *HostHooksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Hook_service(path)
    return self.Hook_service(path[:index]).Service(path[index + 1:])
}

func (op *HostHooksService) String() string {
    return fmt.Sprintf("HostHooksService:%s", op.Path)
}


//
// A service to manage the network interfaces of a host.
//
type HostNicsService struct {
    Service

    NicService  *NicService
}

func NewHostNicsService(connection *Connection, path string) *HostNicsService {
    var result HostNicsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of NICs to return. If not specified all the NICs are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostNicsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *HostNicsService) Service(path string) *HostNicsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Nic_service(path)
    return self.Nic_service(path[:index]).Service(path[index + 1:])
}

func (op *HostNicsService) String() string {
    return fmt.Sprintf("HostNicsService:%s", op.Path)
}


//
//
type HostNumaNodesService struct {
    Service

    NodeService  *NodeService
}

func NewHostNumaNodesService(connection *Connection, path string) *HostNumaNodesService {
    var result HostNumaNodesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of nodes to return. If not specified all the nodes are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostNumaNodesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *HostNumaNodesService) Service(path string) *HostNumaNodesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Node_service(path)
    return self.Node_service(path[:index]).Service(path[index + 1:])
}

func (op *HostNumaNodesService) String() string {
    return fmt.Sprintf("HostNumaNodesService:%s", op.Path)
}


//
// A service to manage host storages.
//
type HostStorageService struct {
    Service

    StorageService  *StorageService
}

func NewHostStorageService(connection *Connection, path string) *HostStorageService {
    var result HostStorageService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Get list of storages.
// [source]
// ----
// GET /ovirt-engine/api/hosts/123/storage
// ----
// The XML response you get will be like this one:
// [source,xml]
// ----
// <host_storages>
//   <host_storage id="123">
//     ...
//   </host_storage>
//   ...
// </host_storages>
// ----
// This method supports the following parameters:
// `ReportStatus`:: Indicates if the status of the LUNs in the storage should be checked.
// Checking the status of the LUN is an heavy weight operation and
// this data is not always needed by the user.
// This parameter will give the option to not perform the status check of the LUNs.
// The default is `true` for backward compatibility.
// Here an example with the LUN status :
// [source,xml]
// ----
// <host_storage id="123">
//   <logical_units>
//     <logical_unit id="123">
//       <lun_mapping>0</lun_mapping>
//       <paths>1</paths>
//       <product_id>lun0</product_id>
//       <serial>123</serial>
//       <size>10737418240</size>
//       <status>used</status>
//       <vendor_id>LIO-ORG</vendor_id>
//       <volume_group_id>123</volume_group_id>
//     </logical_unit>
//   </logical_units>
//   <type>iscsi</type>
//   <host id="123"/>
// </host_storage>
// ----
// Here an example without the LUN status :
// [source,xml]
// ----
// <host_storage id="123">
//   <logical_units>
//     <logical_unit id="123">
//       <lun_mapping>0</lun_mapping>
//       <paths>1</paths>
//       <product_id>lun0</product_id>
//       <serial>123</serial>
//       <size>10737418240</size>
//       <vendor_id>LIO-ORG</vendor_id>
//       <volume_group_id>123</volume_group_id>
//     </logical_unit>
//   </logical_units>
//   <type>iscsi</type>
//   <host id="123"/>
// </host_storage>
// ----
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostStorageService) List (
    reportStatus bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if reportStatus != nil {
        query["report_status"] = reportStatus
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *HostStorageService) Service(path string) *HostStorageService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Storage_service(path)
    return self.Storage_service(path[:index]).Service(path[index + 1:])
}

func (op *HostStorageService) String() string {
    return fmt.Sprintf("HostStorageService:%s", op.Path)
}


//
// A service that manages hosts.
//
type HostsService struct {
    Service

    HostService  *HostService
}

func NewHostsService(connection *Connection, path string) *HostsService {
    var result HostsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a new host.
// The host is created based on the attributes of the `host` parameter. The `name`, `address` and `root_password`
// properties are required.
// For example, to add a host send the following request:
// [source]
// ----
// POST /ovirt-engine/api/hosts
// ----
// With the following request body:
// [source,xml]
// ----
// <host>
//   <name>myhost</name>
//   <address>myhost.example.com</address>
//   <root_password>myrootpassword</root_password>
// </host>
// ----
// NOTE: The `root_password` element is only included in the client-provided initial representation and is not
// exposed in the representations returned from subsequent requests.
// To add a hosted engine host, use the optional `deploy_hosted_engine` parameter:
// [source]
// ----
// POST /ovirt-engine/api/hosts?deploy_hosted_engine=true
// ----
// This method supports the following parameters:
// `Host`:: The host definition from which to create the new host is passed as parameter, and the newly created host
// is returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostsService) Add (
    host Host,
    deployHostedEngine bool,
    undeployHostedEngine bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if deployHostedEngine != nil {
        query["deploy_hosted_engine"] = deployHostedEngine
    }
    if undeployHostedEngine != nil {
        query["undeploy_hosted_engine"] = undeployHostedEngine
    }

    // Send the request
    return op.internalAdd(host, headers, query, wait)
}

//
// Get a list of all available hosts.
// For example, to list the hosts send the following request:
// ....
// GET /ovirt-engine/api/hosts
// ....
// The response body will be something like this:
// [source,xml]
// ----
// <hosts>
//   <host href="/ovirt-engine/api/hosts/123" id="123">
//     ...
//   </host>
//   <host href="/ovirt-engine/api/hosts/456" id="456">
//     ...
//   </host>
//   ...
// </host>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of hosts to return. If not specified all the hosts are returned.
// `Search`:: A query string used to restrict the returned hosts.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostsService) List (
    caseSensitive bool,
    filter bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if filter != nil {
        query["filter"] = filter
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *HostsService) Service(path string) *HostsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Host_service(path)
    return self.Host_service(path[:index]).Service(path[index + 1:])
}

func (op *HostsService) String() string {
    return fmt.Sprintf("HostsService:%s", op.Path)
}


//
// A service to manage an icon (read-only).
//
type IconService struct {
    Service

}

func NewIconService(connection *Connection, path string) *IconService {
    var result IconService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Get an icon.
// [source]
// ----
// GET /ovirt-engine/api/icons/123
// ----
// You will get a XML response like this one:
// [source,xml]
// ----
// <icon id="123">
//   <data>Some binary data here</data>
//   <media_type>image/png</media_type>
// </icon>
// ----
//
func (op *IconService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *IconService) Service(path string) *IconService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *IconService) String() string {
    return fmt.Sprintf("IconService:%s", op.Path)
}


//
// A service to manage icons.
//
type IconsService struct {
    Service

    IconService  *IconService
}

func NewIconsService(connection *Connection, path string) *IconsService {
    var result IconsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Get a list of icons.
// [source]
// ----
// GET /ovirt-engine/api/icons
// ----
// You will get a XML response which is similar to this one:
// [source,xml]
// ----
// <icons>
//   <icon id="123">
//     <data>...</data>
//     <media_type>image/png</media_type>
//   </icon>
//   ...
// </icons>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of icons to return. If not specified all the icons are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *IconsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *IconsService) Service(path string) *IconsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Icon_service(path)
    return self.Icon_service(path[:index]).Service(path[index + 1:])
}

func (op *IconsService) String() string {
    return fmt.Sprintf("IconsService:%s", op.Path)
}


//
//
type ImageService struct {
    Service

}

func NewImageService(connection *Connection, path string) *ImageService {
    var result ImageService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *ImageService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Cluster`:: Cluster where the image should be imported. Has effect only in case `import_as_template` parameter
// is set to `true`.
// `Disk`:: The disk which should be imported.
// `ImportAsTemplate`:: Specify if template should be created from the imported disk.
// `Template`:: Name of the template, which should be created. Has effect only in case `import_as_template` parameter
// is set to `true`.
// `StorageDomain`:: Storage domain where disk should be imported.
// `Async`:: Indicates if the import should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *ImageService) Import (
    async bool,
    cluster Cluster,
    disk Disk,
    importAsTemplate bool,
    storageDomain StorageDomain,
    template Template,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Cluster: cluster,
        Disk: disk,
        ImportAsTemplate: importAsTemplate,
        StorageDomain: storageDomain,
        Template: template,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'import', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ImageService) Service(path string) *ImageService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *ImageService) String() string {
    return fmt.Sprintf("ImageService:%s", op.Path)
}


//
// This service provides a mechanism to control an image transfer. The client will have
// to create a transfer by using <<services/image_transfers/methods/add, add>>
// of the <<services/image_transfers>> service, stating the image to transfer
// data to/from.
// After doing that, the transfer is managed by this service.
// E.g., for uploading to the disk image with id `52cb593f-837c-4633-a444-35a0a0383706`,
// the client can use oVirt's Python's SDK as follows:
// [source,python]
// ----
// transfers_service = system_service.image_transfers_service()
// transfer = transfers_service.add(
//    types.ImageTransfer(
//       image=types.Image(
//          id='52cb593f-837c-4633-a444-35a0a0383706'
//       )
//    )
// )
// ----
// If the user wishes to download a disk rather than upload, he/she should specify
// `download` as the <<types/image_transfer_direction, direction>> attribute of the transfer.
// This will grant a read permission from the image, instead of a write permission.
// E.g:
// [source,python]
// ----
// transfers_service = system_service.image_transfers_service()
// transfer = transfers_service.add(
//    types.ImageTransfer(
//       image=types.Image(
//          id='52cb593f-837c-4633-a444-35a0a0383706'
//       ),
//       direction=types.ImageTransferDirection.DOWNLOAD
//    )
// )
// ----
// Transfers have phases, which govern the flow of the upload/download.
// A client implementing such a flow should poll/check the transfer's phase and
// act accordingly. All the possible phases can be found in
// <<types/image_transfer_phase, ImageTransferPhase>>.
// After adding a new transfer, its phase will be <<types/image_transfer_phase, initializing>>.
// The client will have to poll on the transfer's phase until it changes.
// When the phase becomes <<types/image_transfer_phase, transferring>>,
// the session is ready to start the transfer.
// For example:
// [source,python]
// ----
// transfer_service = transfers_service.image_transfer_service(transfer.id)
// while transfer.phase == types.ImageTransferPhase.INITIALIZING:
//    time.sleep(3)
//    transfer = transfer_service.get()
// ----
// At that stage, if the transfer's phase is <<types/image_transfer_phase, paused_system>>, then the session was
// not successfully established. One possible reason for that is that the ovirt-imageio-daemon is not running
// in the host that was selected for transfer.
// The transfer can be resumed by calling <<services/image_transfer/methods/resume, resume>>
// of the service that manages it.
// If the session was successfully established - the returned transfer entity will
// contain the <<types/image_transfer, proxy_url>> and <<types/image_transfer, signed_ticket>> attributes,
// which the client needs to use in order to transfer the required data. The client can choose whatever
// technique and tool for sending the HTTPS request with the image's data.
// - `proxy_url` is the address of a proxy server to the image, to do I/O to.
// - `signed_ticket` is the content that needs to be added to the `Authentication`
//    header in the HTTPS request, in order to perform a trusted communication.
// For example, Python's HTTPSConnection can be used in order to perform a transfer,
// so an `transfer_headers` dict is set for the upcoming transfer:
// [source,python]
// ----
// transfer_headers = {
//    'Authorization' :  transfer.signed_ticket,
// }
// ----
// Using Python's `HTTPSConnection`, a new connection is established:
// [source,python]
// ----
// # Extract the URI, port, and path from the transfer's proxy_url.
// url = urlparse.urlparse(transfer.proxy_url)
// # Create a new instance of the connection.
// proxy_connection = HTTPSConnection(
//    url.hostname,
//    url.port,
//    context=ssl.SSLContext(ssl.PROTOCOL_SSLv23)
// )
// ----
// For upload, the specific content range being sent must be noted in the `Content-Range` HTTPS
// header. This can be used in order to split the transfer into several requests for
// a more flexible process.
// For doing that, the client will have to repeatedly extend the transfer session
// to keep the channel open. Otherwise, the session will terminate and the transfer will
// get into `paused_system` phase, and HTTPS requests to the server will be rejected.
// E.g., the client can iterate on chunks of the file, and send them to the
// proxy server while asking the service to extend the session:
// [source,python]
// ----
// path = "/path/to/image"
// MB_per_request = 32
// with open(path, "rb") as disk:
//    size = os.path.getsize(path)
//    chunk_size = 1024*1024*MB_per_request
//    pos = 0
//    while (pos < size):
//       transfer_service.extend()
//       transfer_headers['Content-Range'] = "bytes %d-%d/%d" % (pos, min(pos + chunk_size, size)-1, size)
//       proxy_connection.request(
//          'PUT',
//          url.path,
//          disk.read(chunk_size),
//          headers=transfer_headers
//       )
//       r = proxy_connection.getresponse()
//       print r.status, r.reason, "Completed", "{:.0%}".format(pos/ float(size))
//       pos += chunk_size
// ----
// Similarly, for a download transfer, a `Range` header must be sent, making the download process
// more easily managed by downloading the disk in chunks.
// E.g., the client will again iterate on chunks of the disk image, but this time he/she will download
// it to a local file, rather than uploading its own file to the image:
// [source,python]
// ----
// output_file = "/home/user/downloaded_image"
// MiB_per_request = 32
// chunk_size = 1024*1024*MiB_per_request
// total = disk_size
// with open(output_file, "wb") as disk:
//    pos = 0
//    while pos < total:
//       transfer_service.extend()
//       transfer_headers['Range'] = "bytes=%d-%d" %  (pos, min(total, pos + chunk_size) - 1)
//       proxy_connection.request('GET', proxy_url.path, headers=transfer_headers)
//       r = proxy_connection.getresponse()
//       disk.write(r.read())
//       print "Completed", "{:.0%}".format(pos/ float(total))
//       pos += chunk_size
// ----
// When finishing the transfer, the user should call
// <<services/image_transfer/methods/finalize, finalize>>. This will make the
// final adjustments and verifications for finishing the transfer process.
// For example:
// [source,python]
// ----
// transfer_service.finalize()
// ----
// In case of an error, the transfer's phase will be changed to
// <<types/image_transfer_phase, finished_failure>>, and
// the disk's status will be changed to `Illegal`. Otherwise it will be changed to
// <<types/image_transfer_phase, finished_success>>, and the disk will be ready
// to be used. In both cases, the transfer entity will be removed shortly after.
//
type ImageTransferService struct {
    Service

}

func NewImageTransferService(connection *Connection, path string) *ImageTransferService {
    var result ImageTransferService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Extend the image transfer session.
//
func (op *ImageTransferService) Extend (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
    }

    // Send the request and wait for the response:
    return internalAction(action, 'extend', nil, headers, query, wait)
}

//
// After finishing to transfer the data, finalize the transfer.
// This will make sure that the data being transferred is valid and fits the
// image entity that was targeted in the transfer. Specifically, will verify that
// if the image entity is a QCOW disk, the data uploaded is indeed a QCOW file,
// and that the image doesn't have a backing file.
//
func (op *ImageTransferService) Finalize (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
    }

    // Send the request and wait for the response:
    return internalAction(action, 'finalize', nil, headers, query, wait)
}

//
// Get the image transfer entity.
//
func (op *ImageTransferService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Pause the image transfer session.
//
func (op *ImageTransferService) Pause (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
    }

    // Send the request and wait for the response:
    return internalAction(action, 'pause', nil, headers, query, wait)
}

//
// Resume the image transfer session. The client will need to poll the transfer's phase until
// it is different than `resuming`. For example:
// [source,python]
// ----
// transfer_service = transfers_service.image_transfer_service(transfer.id)
// transfer_service.resume()
// transfer = transfer_service.get()
// while transfer.phase == types.ImageTransferPhase.RESUMING:
//    time.sleep(1)
//    transfer = transfer_service.get()
// ----
//
func (op *ImageTransferService) Resume (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
    }

    // Send the request and wait for the response:
    return internalAction(action, 'resume', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ImageTransferService) Service(path string) *ImageTransferService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *ImageTransferService) String() string {
    return fmt.Sprintf("ImageTransferService:%s", op.Path)
}


//
// This service manages image transfers, for performing Image I/O API in oVirt.
// Please refer to <<services/image_transfer, image transfer>> for further
// documentation.
//
type ImageTransfersService struct {
    Service

    ImageTransferService  *ImageTransferService
}

func NewImageTransfersService(connection *Connection, path string) *ImageTransfersService {
    var result ImageTransfersService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add a new image transfer. An image needs to be specified in order to make
// a new transfer.
//
func (op *ImageTransfersService) Add (
    imageTransfer ImageTransfer,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(imageTransfer, headers, query, wait)
}

//
// Retrieves the list of image transfers that are currently
// being performed.
//
func (op *ImageTransfersService) List (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ImageTransfersService) Service(path string) *ImageTransfersService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.ImageTransfer_service(path)
    return self.ImageTransfer_service(path[:index]).Service(path[index + 1:])
}

func (op *ImageTransfersService) String() string {
    return fmt.Sprintf("ImageTransfersService:%s", op.Path)
}


//
//
type ImagesService struct {
    Service

    ImageService  *ImageService
}

func NewImagesService(connection *Connection, path string) *ImagesService {
    var result ImagesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of images to return. If not specified all the images are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *ImagesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ImagesService) Service(path string) *ImagesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Image_service(path)
    return self.Image_service(path[:index]).Service(path[index + 1:])
}

func (op *ImagesService) String() string {
    return fmt.Sprintf("ImagesService:%s", op.Path)
}


//
//
type InstanceTypeService struct {
    Service

    GraphicsConsolesService  *GraphicsConsolesService
    NicsService  *NicsService
    WatchdogsService  *WatchdogsService
}

func NewInstanceTypeService(connection *Connection, path string) *InstanceTypeService {
    var result InstanceTypeService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Get a specific instance type and it's attributes.
// [source]
// ----
// GET /ovirt-engine/api/instancetypes/123
// ----
//
func (op *InstanceTypeService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes a specific instance type from the system.
// If a virtual machine was created using an instance type X after removal of the instance type
// the virtual machine's instance type will be set to `custom`.
// [source]
// ----
// DELETE /ovirt-engine/api/instancetypes/123
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *InstanceTypeService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Update a specific instance type and it's attributes.
// All the attributes are editable after creation.
// If a virtual machine was created using an instance type X and some configuration in instance
// type X was updated, the virtual machine's configuration will be updated automatically by the
// engine.
// [source]
// ----
// PUT /ovirt-engine/api/instancetypes/123
// ----
// For example, to update the memory of instance type `123` to 1 GiB and set the cpu topology
// to 2 sockets and 1 core, send a request like this:
// [source, xml]
// ----
// <instance_type>
//   <memory>1073741824</memory>
//   <cpu>
//     <topology>
//       <cores>1</cores>
//       <sockets>2</sockets>
//       <threads>1</threads>
//     </topology>
//   </cpu>
// </instance_type>
// ----
//
func (op *InstanceTypeService) Update (
    instanceType InstanceType,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(instanceType, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *InstanceTypeService) Service(path string) *InstanceTypeService {
    if path == nil {
        return op
    }
    if path == "graphicsconsoles" {
        return &GraphicsConsolesService{}
    }
    if path.startswith('graphicsconsoles/') {
        return GraphicsConsolesService().service(path[17:])
    }
    if path == "nics" {
        return &NicsService{}
    }
    if path.startswith('nics/') {
        return NicsService().service(path[5:])
    }
    if path == "watchdogs" {
        return &WatchdogsService{}
    }
    if path.startswith('watchdogs/') {
        return WatchdogsService().service(path[10:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *InstanceTypeService) String() string {
    return fmt.Sprintf("InstanceTypeService:%s", op.Path)
}


//
//
type InstanceTypeGraphicsConsoleService struct {
    Service

}

func NewInstanceTypeGraphicsConsoleService(connection *Connection, path string) *InstanceTypeGraphicsConsoleService {
    var result InstanceTypeGraphicsConsoleService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Gets graphics console configuration of the instance type.
//
func (op *InstanceTypeGraphicsConsoleService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Remove the graphics console from the instance type.
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *InstanceTypeGraphicsConsoleService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *InstanceTypeGraphicsConsoleService) Service(path string) *InstanceTypeGraphicsConsoleService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *InstanceTypeGraphicsConsoleService) String() string {
    return fmt.Sprintf("InstanceTypeGraphicsConsoleService:%s", op.Path)
}


//
//
type InstanceTypeGraphicsConsolesService struct {
    Service

    ConsoleService  *ConsoleService
}

func NewInstanceTypeGraphicsConsolesService(connection *Connection, path string) *InstanceTypeGraphicsConsolesService {
    var result InstanceTypeGraphicsConsolesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add new graphics console to the instance type.
//
func (op *InstanceTypeGraphicsConsolesService) Add (
    console GraphicsConsole,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(console, headers, query, wait)
}

//
// Lists all the configured graphics consoles of the instance type.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of consoles to return. If not specified all the consoles are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *InstanceTypeGraphicsConsolesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *InstanceTypeGraphicsConsolesService) Service(path string) *InstanceTypeGraphicsConsolesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Console_service(path)
    return self.Console_service(path[:index]).Service(path[index + 1:])
}

func (op *InstanceTypeGraphicsConsolesService) String() string {
    return fmt.Sprintf("InstanceTypeGraphicsConsolesService:%s", op.Path)
}


//
//
type InstanceTypeNicService struct {
    Service

}

func NewInstanceTypeNicService(connection *Connection, path string) *InstanceTypeNicService {
    var result InstanceTypeNicService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Gets network interface configuration of the instance type.
//
func (op *InstanceTypeNicService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Remove the network interface from the instance type.
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *InstanceTypeNicService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates the network interface configuration of the instance type.
//
func (op *InstanceTypeNicService) Update (
    nic Nic,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(nic, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *InstanceTypeNicService) Service(path string) *InstanceTypeNicService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *InstanceTypeNicService) String() string {
    return fmt.Sprintf("InstanceTypeNicService:%s", op.Path)
}


//
//
type InstanceTypeNicsService struct {
    Service

    NicService  *NicService
}

func NewInstanceTypeNicsService(connection *Connection, path string) *InstanceTypeNicsService {
    var result InstanceTypeNicsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add new network interface to the instance type.
//
func (op *InstanceTypeNicsService) Add (
    nic Nic,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(nic, headers, query, wait)
}

//
// Lists all the configured network interface of the instance type.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of NICs to return. If not specified all the NICs are returned.
// `Search`:: A query string used to restrict the returned templates.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *InstanceTypeNicsService) List (
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *InstanceTypeNicsService) Service(path string) *InstanceTypeNicsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Nic_service(path)
    return self.Nic_service(path[:index]).Service(path[index + 1:])
}

func (op *InstanceTypeNicsService) String() string {
    return fmt.Sprintf("InstanceTypeNicsService:%s", op.Path)
}


//
//
type InstanceTypeWatchdogService struct {
    Service

}

func NewInstanceTypeWatchdogService(connection *Connection, path string) *InstanceTypeWatchdogService {
    var result InstanceTypeWatchdogService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Gets watchdog configuration of the instance type.
//
func (op *InstanceTypeWatchdogService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Remove a watchdog from the instance type.
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *InstanceTypeWatchdogService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates the watchdog configuration of the instance type.
//
func (op *InstanceTypeWatchdogService) Update (
    watchdog Watchdog,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(watchdog, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *InstanceTypeWatchdogService) Service(path string) *InstanceTypeWatchdogService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *InstanceTypeWatchdogService) String() string {
    return fmt.Sprintf("InstanceTypeWatchdogService:%s", op.Path)
}


//
//
type InstanceTypeWatchdogsService struct {
    Service

    WatchdogService  *WatchdogService
}

func NewInstanceTypeWatchdogsService(connection *Connection, path string) *InstanceTypeWatchdogsService {
    var result InstanceTypeWatchdogsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add new watchdog to the instance type.
//
func (op *InstanceTypeWatchdogsService) Add (
    watchdog Watchdog,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(watchdog, headers, query, wait)
}

//
// Lists all the configured watchdogs of the instance type.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of watchdogs to return. If not specified all the watchdogs are
// returned.
// `Search`:: A query string used to restrict the returned templates.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *InstanceTypeWatchdogsService) List (
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *InstanceTypeWatchdogsService) Service(path string) *InstanceTypeWatchdogsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Watchdog_service(path)
    return self.Watchdog_service(path[:index]).Service(path[index + 1:])
}

func (op *InstanceTypeWatchdogsService) String() string {
    return fmt.Sprintf("InstanceTypeWatchdogsService:%s", op.Path)
}


//
//
type InstanceTypesService struct {
    Service

    InstanceTypeService  *InstanceTypeService
}

func NewInstanceTypesService(connection *Connection, path string) *InstanceTypesService {
    var result InstanceTypesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a new instance type.
// This requires only a name attribute and can include all hardware configurations of the
// virtual machine.
// [source]
// ----
// POST /ovirt-engine/api/instancetypes
// ----
// With a request body like this:
// [source,xml]
// ----
// <instance_type>
//   <name>myinstancetype</name>
// </template>
// ----
// Creating an instance type with all hardware configurations with a request body like this:
// [source,xml]
// ----
// <instance_type>
//   <name>myinstancetype</name>
//   <console>
//     <enabled>true</enabled>
//   </console>
//   <cpu>
//     <topology>
//       <cores>2</cores>
//       <sockets>2</sockets>
//       <threads>1</threads>
//     </topology>
//   </cpu>
//   <custom_cpu_model>AMD Opteron_G2</custom_cpu_model>
//   <custom_emulated_machine>q35</custom_emulated_machine>
//   <display>
//     <monitors>1</monitors>
//     <single_qxl_pci>true</single_qxl_pci>
//     <smartcard_enabled>true</smartcard_enabled>
//     <type>spice</type>
//   </display>
//   <high_availability>
//     <enabled>true</enabled>
//     <priority>1</priority>
//   </high_availability>
//   <io>
//     <threads>2</threads>
//   </io>
//   <memory>4294967296</memory>
//   <memory_policy>
//     <ballooning>true</ballooning>
//     <guaranteed>268435456</guaranteed>
//   </memory_policy>
//   <migration>
//     <auto_converge>inherit</auto_converge>
//     <compressed>inherit</compressed>
//     <policy id="00000000-0000-0000-0000-000000000000"/>
//   </migration>
//   <migration_downtime>2</migration_downtime>
//   <os>
//     <boot>
//       <devices>
//         <device>hd</device>
//       </devices>
//     </boot>
//   </os>
//   <rng_device>
//     <rate>
//       <bytes>200</bytes>
//       <period>2</period>
//     </rate>
//     <source>urandom</source>
//   </rng_device>
//   <soundcard_enabled>true</soundcard_enabled>
//   <usb>
//     <enabled>true</enabled>
//     <type>native</type>
//   </usb>
//   <virtio_scsi>
//     <enabled>true</enabled>
//   </virtio_scsi>
// </instance_type>
// ----
//
func (op *InstanceTypesService) Add (
    instanceType InstanceType,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(instanceType, headers, query, wait)
}

//
// Lists all existing instance types in the system.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of instance types to return. If not specified all the instance
// types are returned.
// `Search`:: A query string used to restrict the returned templates.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed
// taking case into account. The default value is `true`, which means that case is taken
// into account. If you want to search ignoring case set it to `false`.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *InstanceTypesService) List (
    caseSensitive bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *InstanceTypesService) Service(path string) *InstanceTypesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.InstanceType_service(path)
    return self.InstanceType_service(path[:index]).Service(path[index + 1:])
}

func (op *InstanceTypesService) String() string {
    return fmt.Sprintf("InstanceTypesService:%s", op.Path)
}


//
//
type IscsiBondService struct {
    Service

    NetworksService  *NetworksService
    StorageServerConnectionsService  *StorageServerConnectionsService
}

func NewIscsiBondService(connection *Connection, path string) *IscsiBondService {
    var result IscsiBondService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *IscsiBondService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes of an existing iSCSI bond.
// For example, to remove the iSCSI bond `456` send a request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/datacenters/123/iscsibonds/456
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *IscsiBondService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates an iSCSI bond.
// Updating of an iSCSI bond can be done on the `name` and the `description` attributes only. For example, to
// update the iSCSI bond `456` of data center `123`, send a request like this:
// [source]
// ----
// PUT /ovirt-engine/api/datacenters/123/iscsibonds/1234
// ----
// The request body should look like this:
// [source,xml]
// ----
// <iscsi_bond>
//    <name>mybond</name>
//    <description>My iSCSI bond</description>
// </iscsi_bond>
// ----
//
func (op *IscsiBondService) Update (
    bond IscsiBond,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(bond, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *IscsiBondService) Service(path string) *IscsiBondService {
    if path == nil {
        return op
    }
    if path == "networks" {
        return &NetworksService{}
    }
    if path.startswith('networks/') {
        return NetworksService().service(path[9:])
    }
    if path == "storageserverconnections" {
        return &StorageServerConnectionsService{}
    }
    if path.startswith('storageserverconnections/') {
        return StorageServerConnectionsService().service(path[25:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *IscsiBondService) String() string {
    return fmt.Sprintf("IscsiBondService:%s", op.Path)
}


//
//
type IscsiBondsService struct {
    Service

    IscsiBondService  *IscsiBondService
}

func NewIscsiBondsService(connection *Connection, path string) *IscsiBondsService {
    var result IscsiBondsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Create a new iSCSI bond on a data center.
// For example, to create a new iSCSI bond on data center `123` using storage connections `456` and `789`, send a
// request like this:
// [source]
// ----
// POST /ovirt-engine/api/datacenters/123/iscsibonds
// ----
// The request body should look like this:
// [source,xml]
// ----
// <iscsi_bond>
//   <name>mybond</name>
//   <storage_connections>
//     <storage_connection id="456"/>
//     <storage_connection id="789"/>
//   </storage_connections>
//   <networks>
//     <network id="abc"/>
//   </networks>
// </iscsi_bond>
// ----
//
func (op *IscsiBondsService) Add (
    bond IscsiBond,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(bond, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of bonds to return. If not specified all the bonds are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *IscsiBondsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *IscsiBondsService) Service(path string) *IscsiBondsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.IscsiBond_service(path)
    return self.IscsiBond_service(path[:index]).Service(path[index + 1:])
}

func (op *IscsiBondsService) String() string {
    return fmt.Sprintf("IscsiBondsService:%s", op.Path)
}


//
// A service to manage a job.
//
type JobService struct {
    Service

    StepsService  *StepsService
}

func NewJobService(connection *Connection, path string) *JobService {
    var result JobService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Set an external job execution to be cleared by the system.
// For example, to set a job with identifier `123` send the following request:
// [source]
// ----
// POST /ovirt-engine/api/jobs/clear
// ----
// With the following request body:
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *JobService) Clear (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'clear', nil, headers, query, wait)
}

//
// Marks an external job execution as ended.
// For example, to terminate a job with identifier `123` send the following request:
// [source]
// ----
// POST /ovirt-engine/api/jobs/end
// ----
// With the following request body:
// [source,xml]
// ----
// <action>
//   <force>true</force>
//   <status>finished</status>
// </action>
// ----
// This method supports the following parameters:
// `Force`:: Indicates if the job should be forcibly terminated.
// `Succeeded`:: Indicates if the job should be marked as successfully finished or as failed.
// This parameter is optional, and the default value is `true`.
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *JobService) End (
    async bool,
    force bool,
    succeeded bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Force: force,
        Succeeded: succeeded,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'end', nil, headers, query, wait)
}

//
// Retrieves a job.
// [source]
// ----
// GET /ovirt-engine/api/jobs/123
// ----
// You will receive response in XML like this one:
// [source,xml]
// ----
// <job href="/ovirt-engine/api/jobs/123" id="123">
//   <actions>
//     <link href="/ovirt-engine/api/jobs/123/clear" rel="clear"/>
//     <link href="/ovirt-engine/api/jobs/123/end" rel="end"/>
//   </actions>
//   <description>Adding Disk</description>
//   <link href="/ovirt-engine/api/jobs/123/steps" rel="steps"/>
//   <auto_cleared>true</auto_cleared>
//   <end_time>2016-12-12T23:07:29.758+02:00</end_time>
//   <external>false</external>
//   <last_updated>2016-12-12T23:07:29.758+02:00</last_updated>
//   <start_time>2016-12-12T23:07:26.593+02:00</start_time>
//   <status>failed</status>
//   <owner href="/ovirt-engine/api/users/456" id="456"/>
// </job>
// ----
//
func (op *JobService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *JobService) Service(path string) *JobService {
    if path == nil {
        return op
    }
    if path == "steps" {
        return &StepsService{}
    }
    if path.startswith('steps/') {
        return StepsService().service(path[6:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *JobService) String() string {
    return fmt.Sprintf("JobService:%s", op.Path)
}


//
// A service to manage jobs.
//
type JobsService struct {
    Service

    JobService  *JobService
}

func NewJobsService(connection *Connection, path string) *JobsService {
    var result JobsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add an external job.
// For example, to add a job with the following request:
// [source]
// ----
// POST /ovirt-engine/api/jobs
// ----
// With the following request body:
// [source,xml]
// ----
// <job>
//   <description>Doing some work</description>
//   <auto_cleared>true</auto_cleared>
// </job>
// ----
// The response should look like:
// [source,xml]
// ----
// <job href="/ovirt-engine/api/jobs/123" id="123">
//   <actions>
//     <link href="/ovirt-engine/api/jobs/123/clear" rel="clear"/>
//     <link href="/ovirt-engine/api/jobs/123/end" rel="end"/>
//   </actions>
//   <description>Doing some work</description>
//   <link href="/ovirt-engine/api/jobs/123/steps" rel="steps"/>
//   <auto_cleared>true</auto_cleared>
//   <external>true</external>
//   <last_updated>2016-12-13T02:15:42.130+02:00</last_updated>
//   <start_time>2016-12-13T02:15:42.130+02:00</start_time>
//   <status>started</status>
//   <owner href="/ovirt-engine/api/users/456" id="456"/>
// </job>
// ----
// This method supports the following parameters:
// `Job`:: Job that will be added.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *JobsService) Add (
    job Job,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(job, headers, query, wait)
}

//
// Retrieves the representation of the jobs.
// [source]
// ----
// GET /ovirt-engine/api/jobs
// ----
// You will receive response in XML like this one:
// [source,xml]
// ----
// <jobs>
//   <job href="/ovirt-engine/api/jobs/123" id="123">
//     <actions>
//       <link href="/ovirt-engine/api/jobs/123/clear" rel="clear"/>
//       <link href="/ovirt-engine/api/jobs/123/end" rel="end"/>
//     </actions>
//     <description>Adding Disk</description>
//     <link href="/ovirt-engine/api/jobs/123/steps" rel="steps"/>
//     <auto_cleared>true</auto_cleared>
//     <end_time>2016-12-12T23:07:29.758+02:00</end_time>
//     <external>false</external>
//     <last_updated>2016-12-12T23:07:29.758+02:00</last_updated>
//     <start_time>2016-12-12T23:07:26.593+02:00</start_time>
//     <status>failed</status>
//     <owner href="/ovirt-engine/api/users/456" id="456"/>
//   </job>
//   ...
// </jobs>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of jobs to return. If not specified all the jobs are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *JobsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *JobsService) Service(path string) *JobsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Job_service(path)
    return self.Job_service(path[:index]).Service(path[index + 1:])
}

func (op *JobsService) String() string {
    return fmt.Sprintf("JobsService:%s", op.Path)
}


//
// A service to manage Katello errata.
// The information is retrieved from Katello.
//
type KatelloErrataService struct {
    Service

    KatelloErratumService  *KatelloErratumService
}

func NewKatelloErrataService(connection *Connection, path string) *KatelloErrataService {
    var result KatelloErrataService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieves the representation of the Katello errata.
// [source]
// ----
// GET /ovirt-engine/api/katelloerrata
// ----
// You will receive response in XML like this one:
// [source,xml]
// ----
// <katello_errata>
//   <katello_erratum href="/ovirt-engine/api/katelloerrata/123" id="123">
//     <name>RHBA-2013:XYZ</name>
//     <description>The description of the erratum</description>
//     <title>some bug fix update</title>
//     <type>bugfix</type>
//     <issued>2013-11-20T02:00:00.000+02:00</issued>
//     <solution>Few guidelines regarding the solution</solution>
//     <summary>Updated packages that fix one bug are now available for XYZ</summary>
//     <packages>
//       <package>
//         <name>libipa_hbac-1.9.2-82.11.el6_4.i686</name>
//       </package>
//       ...
//     </packages>
//   </katello_erratum>
//   ...
// </katello_errata>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of errata to return. If not specified all the errata are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *KatelloErrataService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *KatelloErrataService) Service(path string) *KatelloErrataService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.KatelloErratum_service(path)
    return self.KatelloErratum_service(path[:index]).Service(path[index + 1:])
}

func (op *KatelloErrataService) String() string {
    return fmt.Sprintf("KatelloErrataService:%s", op.Path)
}


//
// A service to manage a Katello erratum.
//
type KatelloErratumService struct {
    Service

}

func NewKatelloErratumService(connection *Connection, path string) *KatelloErratumService {
    var result KatelloErratumService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieves a Katello erratum.
// [source]
// ----
// GET /ovirt-engine/api/katelloerrata/123
// ----
// You will receive response in XML like this one:
// [source,xml]
// ----
// <katello_erratum href="/ovirt-engine/api/katelloerrata/123" id="123">
//   <name>RHBA-2013:XYZ</name>
//   <description>The description of the erratum</description>
//   <title>some bug fix update</title>
//   <type>bugfix</type>
//   <issued>2013-11-20T02:00:00.000+02:00</issued>
//   <solution>Few guidelines regarding the solution</solution>
//   <summary>Updated packages that fix one bug are now available for XYZ</summary>
//   <packages>
//     <package>
//       <name>libipa_hbac-1.9.2-82.11.el6_4.i686</name>
//     </package>
//     ...
//   </packages>
// </katello_erratum>
// ----
//
func (op *KatelloErratumService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *KatelloErratumService) Service(path string) *KatelloErratumService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *KatelloErratumService) String() string {
    return fmt.Sprintf("KatelloErratumService:%s", op.Path)
}


//
//
type MacPoolService struct {
    Service

}

func NewMacPoolService(connection *Connection, path string) *MacPoolService {
    var result MacPoolService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *MacPoolService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes a MAC address pool.
// For example, to remove the MAC address pool having id `123` send a request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/macpools/123
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *MacPoolService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates a MAC address pool.
// The `name`, `description`, `allow_duplicates`, and `ranges` attributes can be updated.
// For example, to update the MAC address pool of id `123` send a request like this:
// [source]
// ----
// PUT /ovirt-engine/api/macpools/123
// ----
// With a request body like this:
// [source,xml]
// ----
// <mac_pool>
//   <name>UpdatedMACPool</name>
//   <description>An updated MAC address pool</description>
//   <allow_duplicates>false</allow_duplicates>
//   <ranges>
//     <range>
//       <from>00:1A:4A:16:01:51</from>
//       <to>00:1A:4A:16:01:e6</to>
//     </range>
//     <range>
//       <from>02:1A:4A:01:00:00</from>
//       <to>02:1A:4A:FF:FF:FF</to>
//     </range>
//   </ranges>
// </mac_pool>
// ----
//
func (op *MacPoolService) Update (
    pool MacPool,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(pool, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *MacPoolService) Service(path string) *MacPoolService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *MacPoolService) String() string {
    return fmt.Sprintf("MacPoolService:%s", op.Path)
}


//
//
type MacPoolsService struct {
    Service

    MacPoolService  *MacPoolService
}

func NewMacPoolsService(connection *Connection, path string) *MacPoolsService {
    var result MacPoolsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a new MAC address pool.
// Creation of a MAC address pool requires values for the `name` and `ranges` attributes.
// For example, to create MAC address pool send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/macpools
// ----
// With a request body like this:
// [source,xml]
// ----
// <mac_pool>
//   <name>MACPool</name>
//   <description>A MAC address pool</description>
//   <allow_duplicates>true</allow_duplicates>
//   <default_pool>false</default_pool>
//   <ranges>
//     <range>
//       <from>00:1A:4A:16:01:51</from>
//       <to>00:1A:4A:16:01:e6</to>
//     </range>
//   </ranges>
// </mac_pool>
// ----
//
func (op *MacPoolsService) Add (
    pool MacPool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(pool, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of pools to return. If not specified all the pools are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *MacPoolsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *MacPoolsService) Service(path string) *MacPoolsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.MacPool_service(path)
    return self.MacPool_service(path[:index]).Service(path[index + 1:])
}

func (op *MacPoolsService) String() string {
    return fmt.Sprintf("MacPoolsService:%s", op.Path)
}


//
//
type MeasurableService struct {
    Service

    StatisticsService  *StatisticsService
}

func NewMeasurableService(connection *Connection, path string) *MeasurableService {
    var result MeasurableService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *MeasurableService) Service(path string) *MeasurableService {
    if path == nil {
        return op
    }
    if path == "statistics" {
        return &StatisticsService{}
    }
    if path.startswith('statistics/') {
        return StatisticsService().service(path[11:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *MeasurableService) String() string {
    return fmt.Sprintf("MeasurableService:%s", op.Path)
}


//
//
type MoveableService struct {
    Service

}

func NewMoveableService(connection *Connection, path string) *MoveableService {
    var result MoveableService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the move should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *MoveableService) Move (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'move', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *MoveableService) Service(path string) *MoveableService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *MoveableService) String() string {
    return fmt.Sprintf("MoveableService:%s", op.Path)
}


//
// A service managing a network
//
type NetworkService struct {
    Service

    NetworkLabelsService  *NetworkLabelsService
    PermissionsService  *PermissionsService
    VnicProfilesService  *VnicProfilesService
}

func NewNetworkService(connection *Connection, path string) *NetworkService {
    var result NetworkService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Gets a logical network.
// For example:
// [source]
// ----
// GET /ovirt-engine/api/networks/123
// ----
// Will respond:
// [source,xml]
// ----
// <network href="/ovirt-engine/api/networks/123" id="123">
//   <name>ovirtmgmt</name>
//   <description>Default Management Network</description>
//   <link href="/ovirt-engine/api/networks/123/permissions" rel="permissions"/>
//   <link href="/ovirt-engine/api/networks/123/vnicprofiles" rel="vnicprofiles"/>
//   <link href="/ovirt-engine/api/networks/123/networklabels" rel="networklabels"/>
//   <mtu>0</mtu>
//   <stp>false</stp>
//   <usages>
//     <usage>vm</usage>
//   </usages>
//   <data_center href="/ovirt-engine/api/datacenters/456" id="456"/>
// </network>
// ----
//
func (op *NetworkService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes a logical network, or the association of a logical network to a data center.
// For example, to remove the logical network `123` send a request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/networks/123
// ----
// Each network is bound exactly to one data center. So if we disassociate network with data center it has the same
// result as if we would just remove that network. However it might be more specific to say we're removing network
// `456` of data center `123`.
// For example, to remove the association of network `456` to data center `123` send a request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/datacenters/123/networks/456
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *NetworkService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates a logical network.
// The `name`, `description`, `ip`, `vlan`, `stp` and `display` attributes can be updated.
// For example, to update the description of the logical network `123` send a request like this:
// [source]
// ----
// PUT /ovirt-engine/api/networks/123
// ----
// With a request body like this:
// [source,xml]
// ----
// <network>
//   <description>My updated description</description>
// </network>
// ----
// The maximum transmission unit of a network is set using a PUT request to
// specify the integer value of the `mtu` attribute.
// For example, to set the maximum transmission unit send a request like this:
// [source]
// ----
// PUT /ovirt-engine/api/datacenters/123/networks/456
// ----
// With a request body like this:
// [source,xml]
// ----
// <network>
//   <mtu>1500</mtu>
// </network>
// ----
//
func (op *NetworkService) Update (
    network Network,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(network, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *NetworkService) Service(path string) *NetworkService {
    if path == nil {
        return op
    }
    if path == "networklabels" {
        return &NetworkLabelsService{}
    }
    if path.startswith('networklabels/') {
        return NetworkLabelsService().service(path[14:])
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    if path == "vnicprofiles" {
        return &VnicProfilesService{}
    }
    if path.startswith('vnicprofiles/') {
        return VnicProfilesService().service(path[13:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *NetworkService) String() string {
    return fmt.Sprintf("NetworkService:%s", op.Path)
}


//
//
type NetworkAttachmentService struct {
    Service

}

func NewNetworkAttachmentService(connection *Connection, path string) *NetworkAttachmentService {
    var result NetworkAttachmentService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *NetworkAttachmentService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *NetworkAttachmentService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
//
func (op *NetworkAttachmentService) Update (
    attachment NetworkAttachment,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(attachment, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *NetworkAttachmentService) Service(path string) *NetworkAttachmentService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *NetworkAttachmentService) String() string {
    return fmt.Sprintf("NetworkAttachmentService:%s", op.Path)
}


//
//
type NetworkAttachmentsService struct {
    Service

    AttachmentService  *AttachmentService
}

func NewNetworkAttachmentsService(connection *Connection, path string) *NetworkAttachmentsService {
    var result NetworkAttachmentsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *NetworkAttachmentsService) Add (
    attachment NetworkAttachment,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(attachment, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of attachments to return. If not specified all the attachments are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *NetworkAttachmentsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *NetworkAttachmentsService) Service(path string) *NetworkAttachmentsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Attachment_service(path)
    return self.Attachment_service(path[:index]).Service(path[index + 1:])
}

func (op *NetworkAttachmentsService) String() string {
    return fmt.Sprintf("NetworkAttachmentsService:%s", op.Path)
}


//
// Manages a network filter.
// [source,xml]
// ----
// <network_filter id="00000019-0019-0019-0019-00000000026b">
//   <name>example-network-filter-b</name>
//   <version>
//     <major>4</major>
//     <minor>0</minor>
//     <build>-1</build>
//     <revision>-1</revision>
//   </version>
// </network_filter>
// ----
// Please note that version is referring to the minimal support version for the specific filter.
//
type NetworkFilterService struct {
    Service

}

func NewNetworkFilterService(connection *Connection, path string) *NetworkFilterService {
    var result NetworkFilterService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieves a representation of the network filter.
//
func (op *NetworkFilterService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *NetworkFilterService) Service(path string) *NetworkFilterService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *NetworkFilterService) String() string {
    return fmt.Sprintf("NetworkFilterService:%s", op.Path)
}


//
// This service manages a parameter for a network filter.
//
type NetworkFilterParameterService struct {
    Service

}

func NewNetworkFilterParameterService(connection *Connection, path string) *NetworkFilterParameterService {
    var result NetworkFilterParameterService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieves a representation of the network filter parameter.
//
func (op *NetworkFilterParameterService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes the filter parameter.
// For example, to remove the filter parameter with id `123` on NIC `456` of virtual machine `789`
// send a request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/vms/789/nics/456/networkfilterparameters/123
// ----
//
func (op *NetworkFilterParameterService) Remove (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates the network filter parameter.
// For example, to update the network filter parameter having with with id `123` on NIC `456` of
// virtual machine `789` send a request like this:
// [source]
// ----
// PUT /ovirt-engine/api/vms/789/nics/456/networkfilterparameters/123
// ----
// With a request body like this:
// [source,xml]
// ----
// <network_filter_parameter>
//   <name>updatedName</name>
//   <value>updatedValue</value>
// </network_filter_parameter>
// ----
// This method supports the following parameters:
// `Parameter`:: The network filter parameter that is being updated.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *NetworkFilterParameterService) Update (
    parameter NetworkFilterParameter,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalUpdate(parameter, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *NetworkFilterParameterService) Service(path string) *NetworkFilterParameterService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *NetworkFilterParameterService) String() string {
    return fmt.Sprintf("NetworkFilterParameterService:%s", op.Path)
}


//
// This service manages a collection of parameters for network filters.
//
type NetworkFilterParametersService struct {
    Service

    ParameterService  *ParameterService
}

func NewNetworkFilterParametersService(connection *Connection, path string) *NetworkFilterParametersService {
    var result NetworkFilterParametersService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add a network filter parameter.
// For example, to add the parameter for the network filter on NIC `456` of
// virtual machine `789` send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/vms/789/nics/456/networkfilterparameters
// ----
// With a request body like this:
// [source,xml]
// ----
// <network_filter_parameter>
//   <name>IP</name>
//   <value>10.0.1.2</value>
// </network_filter_parameter>
// ----
// This method supports the following parameters:
// `Parameter`:: The network filter parameter that is being added.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *NetworkFilterParametersService) Add (
    parameter NetworkFilterParameter,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(parameter, headers, query, wait)
}

//
// Retrieves the representations of the network filter parameters.
//
func (op *NetworkFilterParametersService) List (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *NetworkFilterParametersService) Service(path string) *NetworkFilterParametersService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Parameter_service(path)
    return self.Parameter_service(path[:index]).Service(path[index + 1:])
}

func (op *NetworkFilterParametersService) String() string {
    return fmt.Sprintf("NetworkFilterParametersService:%s", op.Path)
}


//
// Represents a readonly network filters sub-collection.
// The network filter enables to filter packets send to/from the VM's nic according to defined rules.
// For more information please refer to <<services/network_filter,NetworkFilter>> service documentation
// Network filters are supported in different versions, starting from version 3.0.
// A network filter is defined for each vnic profile.
// A vnic profile is defined for a specific network.
// A network can be assigned to several different clusters. In the future, each network will be defined in
// cluster level.
// Currently, each network is being defined at data center level. Potential network filters for each network
// are determined by the network's data center compatibility version V.
// V must be >= the network filter version in order to configure this network filter for a specific network.
// Please note, that if a network is assigned to cluster with a version supporting a network filter, the filter
// may not be available due to the data center version being smaller then the network filter's version.
// Example of listing all of the supported network filters for a specific cluster:
// [source]
// ----
// GET http://localhost:8080/ovirt-engine/api/clusters/{cluster:id}/networkfilters
// ----
// Output:
// [source,xml]
// ----
// <network_filters>
//   <network_filter id="00000019-0019-0019-0019-00000000026c">
//     <name>example-network-filter-a</name>
//     <version>
//       <major>4</major>
//       <minor>0</minor>
//       <build>-1</build>
//       <revision>-1</revision>
//     </version>
//   </network_filter>
//   <network_filter id="00000019-0019-0019-0019-00000000026b">
//     <name>example-network-filter-b</name>
//     <version>
//       <major>4</major>
//       <minor>0</minor>
//       <build>-1</build>
//       <revision>-1</revision>
//     </version>
//   </network_filter>
//   <network_filter id="00000019-0019-0019-0019-00000000026a">
//     <name>example-network-filter-a</name>
//     <version>
//       <major>3</major>
//       <minor>0</minor>
//       <build>-1</build>
//       <revision>-1</revision>
//     </version>
//   </network_filter>
// </network_filters>
// ----
//
type NetworkFiltersService struct {
    Service

    NetworkFilterService  *NetworkFilterService
}

func NewNetworkFiltersService(connection *Connection, path string) *NetworkFiltersService {
    var result NetworkFiltersService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieves the representations of the network filters.
//
func (op *NetworkFiltersService) List (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *NetworkFiltersService) Service(path string) *NetworkFiltersService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.NetworkFilter_service(path)
    return self.NetworkFilter_service(path[:index]).Service(path[index + 1:])
}

func (op *NetworkFiltersService) String() string {
    return fmt.Sprintf("NetworkFiltersService:%s", op.Path)
}


//
//
type NetworkLabelService struct {
    Service

}

func NewNetworkLabelService(connection *Connection, path string) *NetworkLabelService {
    var result NetworkLabelService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *NetworkLabelService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes a label from a logical network.
// For example, to remove the label `exemplary` from a logical network having id `123` send the following request:
// [source]
// ----
// DELETE /ovirt-engine/api/networks/123/labels/exemplary
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *NetworkLabelService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *NetworkLabelService) Service(path string) *NetworkLabelService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *NetworkLabelService) String() string {
    return fmt.Sprintf("NetworkLabelService:%s", op.Path)
}


//
//
type NetworkLabelsService struct {
    Service

    LabelService  *LabelService
}

func NewNetworkLabelsService(connection *Connection, path string) *NetworkLabelsService {
    var result NetworkLabelsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Attaches label to logical network.
// You can attach labels to a logical network to automate the association of that logical network with physical host
// network interfaces to which the same label has been attached.
// For example, to attach the label `mylabel` to a logical network having id `123` send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/networks/123/labels
// ----
// With a request body like this:
// [source,xml]
// ----
// <label id="mylabel"/>
// ----
//
func (op *NetworkLabelsService) Add (
    label NetworkLabel,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(label, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of labels to return. If not specified all the labels are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *NetworkLabelsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *NetworkLabelsService) Service(path string) *NetworkLabelsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Label_service(path)
    return self.Label_service(path[:index]).Service(path[index + 1:])
}

func (op *NetworkLabelsService) String() string {
    return fmt.Sprintf("NetworkLabelsService:%s", op.Path)
}


//
// Manages logical networks.
// The engine creates a default `ovirtmgmt` network on installation. This network acts as the management network for
// access to hypervisor hosts. This network is associated with the `Default` cluster and is a member of the `Default`
// data center.
//
type NetworksService struct {
    Service

    NetworkService  *NetworkService
}

func NewNetworksService(connection *Connection, path string) *NetworksService {
    var result NetworksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a new logical network, or associates an existing network with a data center.
// Creation of a new network requires the `name` and `data_center` elements.
// For example, to create a network named `mynetwork` for data center `123` send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/networks
// ----
// With a request body like this:
// [source,xml]
// ----
// <network>
//   <name>mynetwork</name>
//   <data_center id="123"/>
// </network>
// ----
// To associate the existing network `456` with the data center `123` send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/datacenters/123/networks
// ----
// With a request body like this:
// [source,xml]
// ----
// <network>
//   <name>ovirtmgmt</name>
// </network>
// ----
//
func (op *NetworksService) Add (
    network Network,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(network, headers, query, wait)
}

//
// List logical networks.
// For example:
// [source]
// ----
// GET /ovirt-engine/api/networks
// ----
// Will respond:
// [source,xml]
// ----
// <networks>
//   <network href="/ovirt-engine/api/networks/123" id="123">
//     <name>ovirtmgmt</name>
//     <description>Default Management Network</description>
//     <link href="/ovirt-engine/api/networks/123/permissions" rel="permissions"/>
//     <link href="/ovirt-engine/api/networks/123/vnicprofiles" rel="vnicprofiles"/>
//     <link href="/ovirt-engine/api/networks/123/networklabels" rel="networklabels"/>
//     <mtu>0</mtu>
//     <stp>false</stp>
//     <usages>
//       <usage>vm</usage>
//     </usages>
//     <data_center href="/ovirt-engine/api/datacenters/456" id="456"/>
//   </network>
//   ...
// </networks>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of networks to return. If not specified all the networks are returned.
// `Search`:: A query string used to restrict the returned networks.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *NetworksService) List (
    caseSensitive bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *NetworksService) Service(path string) *NetworksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Network_service(path)
    return self.Network_service(path[:index]).Service(path[index + 1:])
}

func (op *NetworksService) String() string {
    return fmt.Sprintf("NetworksService:%s", op.Path)
}


//
//
type OpenstackImageService struct {
    Service

}

func NewOpenstackImageService(connection *Connection, path string) *OpenstackImageService {
    var result OpenstackImageService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *OpenstackImageService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Imports a virtual machine from a Glance image storage domain.
// For example, to import the image with identifier `456` from the
// storage domain with identifier `123` send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/openstackimageproviders/123/images/456/import
// ----
// With a request body like this:
// [source,xml]
// ----
// <action>
//   <storage_domain>
//     <name>images0</name>
//   </storage_domain>
//   <cluster>
//     <name>images0</name>
//   </cluster>
// </action>
// ----
// This method supports the following parameters:
// `ImportAsTemplate`:: Indicates whether the image should be imported as a template.
// `Cluster`:: This parameter is mandatory in case of using `import_as_template` and indicates which cluster should be used
// for import glance image as template.
// `Async`:: Indicates if the import should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackImageService) Import (
    async bool,
    cluster Cluster,
    disk Disk,
    importAsTemplate bool,
    storageDomain StorageDomain,
    template Template,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Cluster: cluster,
        Disk: disk,
        ImportAsTemplate: importAsTemplate,
        StorageDomain: storageDomain,
        Template: template,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'import', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackImageService) Service(path string) *OpenstackImageService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *OpenstackImageService) String() string {
    return fmt.Sprintf("OpenstackImageService:%s", op.Path)
}


//
//
type OpenstackImageProviderService struct {
    ExternalProviderService

    CertificatesService  *CertificatesService
    ImagesService  *ImagesService
}

func NewOpenstackImageProviderService(connection *Connection, path string) *OpenstackImageProviderService {
    var result OpenstackImageProviderService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *OpenstackImageProviderService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
//
func (op *OpenstackImageProviderService) ImportCertificates (
    certificates []*Certificate,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Certificates: certificates,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'importcertificates', nil, headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackImageProviderService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the test should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackImageProviderService) TestConnectivity (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'testconnectivity', nil, headers, query, wait)
}

//
//
func (op *OpenstackImageProviderService) Update (
    provider OpenStackImageProvider,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(provider, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackImageProviderService) Service(path string) *OpenstackImageProviderService {
    if path == nil {
        return op
    }
    if path == "certificates" {
        return &CertificatesService{}
    }
    if path.startswith('certificates/') {
        return CertificatesService().service(path[13:])
    }
    if path == "images" {
        return &ImagesService{}
    }
    if path.startswith('images/') {
        return ImagesService().service(path[7:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *OpenstackImageProviderService) String() string {
    return fmt.Sprintf("OpenstackImageProviderService:%s", op.Path)
}


//
//
type OpenstackImageProvidersService struct {
    Service

    ProviderService  *ProviderService
}

func NewOpenstackImageProvidersService(connection *Connection, path string) *OpenstackImageProvidersService {
    var result OpenstackImageProvidersService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *OpenstackImageProvidersService) Add (
    provider OpenStackImageProvider,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(provider, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of providers to return. If not specified all the providers are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackImageProvidersService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackImageProvidersService) Service(path string) *OpenstackImageProvidersService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Provider_service(path)
    return self.Provider_service(path[:index]).Service(path[index + 1:])
}

func (op *OpenstackImageProvidersService) String() string {
    return fmt.Sprintf("OpenstackImageProvidersService:%s", op.Path)
}


//
//
type OpenstackImagesService struct {
    Service

    ImageService  *ImageService
}

func NewOpenstackImagesService(connection *Connection, path string) *OpenstackImagesService {
    var result OpenstackImagesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Lists the images of a Glance image storage domain.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of images to return. If not specified all the images are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackImagesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackImagesService) Service(path string) *OpenstackImagesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Image_service(path)
    return self.Image_service(path[:index]).Service(path[index + 1:])
}

func (op *OpenstackImagesService) String() string {
    return fmt.Sprintf("OpenstackImagesService:%s", op.Path)
}


//
//
type OpenstackNetworkService struct {
    Service

    SubnetsService  *SubnetsService
}

func NewOpenstackNetworkService(connection *Connection, path string) *OpenstackNetworkService {
    var result OpenstackNetworkService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *OpenstackNetworkService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This operation imports an external network into oVirt.
// The network will be added to the data center specified.
// This method supports the following parameters:
// `DataCenter`:: The data center into which the network is to be imported.
// Data center is mandatory, and can be specified
// using the `id` or `name` attributes, the rest of
// the attributes will be ignored.
// `Async`:: Indicates if the import should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackNetworkService) Import (
    async bool,
    dataCenter DataCenter,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        DataCenter: dataCenter,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'import', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackNetworkService) Service(path string) *OpenstackNetworkService {
    if path == nil {
        return op
    }
    if path == "subnets" {
        return &SubnetsService{}
    }
    if path.startswith('subnets/') {
        return SubnetsService().service(path[8:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *OpenstackNetworkService) String() string {
    return fmt.Sprintf("OpenstackNetworkService:%s", op.Path)
}


//
// This service manages OpenStack network provider.
//
type OpenstackNetworkProviderService struct {
    ExternalProviderService

    CertificatesService  *CertificatesService
    NetworksService  *NetworksService
}

func NewOpenstackNetworkProviderService(connection *Connection, path string) *OpenstackNetworkProviderService {
    var result OpenstackNetworkProviderService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Returns the representation of the object managed by this service.
// For example, to get the OpenStack network provider with identifier `1234`, send a request like this:
// [source]
// ----
// GET /ovirt-engine/api/openstacknetworkproviders/1234
// ----
//
func (op *OpenstackNetworkProviderService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
//
func (op *OpenstackNetworkProviderService) ImportCertificates (
    certificates []*Certificate,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Certificates: certificates,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'importcertificates', nil, headers, query, wait)
}

//
// Removes the provider.
// For example, to remove the OpenStack network provider with identifier `1234`, send a request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/openstacknetworkproviders/1234
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackNetworkProviderService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the test should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackNetworkProviderService) TestConnectivity (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'testconnectivity', nil, headers, query, wait)
}

//
// Updates the provider.
// For example, to update `provider_name`, `requires_authentication`, `url`, `tenant_name` and `type` properties,
// for the OpenStack network provider with identifier `1234`, send a request like this:
// [source]
// ----
// PUT /ovirt-engine/api/openstacknetworkproviders/1234
// ----
// With a request body like this:
// [source,xml]
// ----
// <openstack_network_provider>
//   <name>ovn-network-provider</name>
//   <requires_authentication>false</requires_authentication>
//   <url>http://some_server_url.domain.com:9696</url>
//   <tenant_name>oVirt</tenant_name>
//   <type>external</type>
// </openstack_network_provider>
// ----
// This method supports the following parameters:
// `Provider`:: The provider to update.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackNetworkProviderService) Update (
    provider OpenStackNetworkProvider,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(provider, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackNetworkProviderService) Service(path string) *OpenstackNetworkProviderService {
    if path == nil {
        return op
    }
    if path == "certificates" {
        return &CertificatesService{}
    }
    if path.startswith('certificates/') {
        return CertificatesService().service(path[13:])
    }
    if path == "networks" {
        return &NetworksService{}
    }
    if path.startswith('networks/') {
        return NetworksService().service(path[9:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *OpenstackNetworkProviderService) String() string {
    return fmt.Sprintf("OpenstackNetworkProviderService:%s", op.Path)
}


//
// This service manages OpenStack network providers.
//
type OpenstackNetworkProvidersService struct {
    Service

    ProviderService  *ProviderService
}

func NewOpenstackNetworkProvidersService(connection *Connection, path string) *OpenstackNetworkProvidersService {
    var result OpenstackNetworkProvidersService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// The operation adds a new network provider to the system.
// If the `type` property is not present, a default value of `NEUTRON` will be used.
//
func (op *OpenstackNetworkProvidersService) Add (
    provider OpenStackNetworkProvider,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(provider, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of providers to return. If not specified all the providers are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackNetworkProvidersService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackNetworkProvidersService) Service(path string) *OpenstackNetworkProvidersService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Provider_service(path)
    return self.Provider_service(path[:index]).Service(path[index + 1:])
}

func (op *OpenstackNetworkProvidersService) String() string {
    return fmt.Sprintf("OpenstackNetworkProvidersService:%s", op.Path)
}


//
//
type OpenstackNetworksService struct {
    Service

    NetworkService  *NetworkService
}

func NewOpenstackNetworksService(connection *Connection, path string) *OpenstackNetworksService {
    var result OpenstackNetworksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of networks to return. If not specified all the networks are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackNetworksService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackNetworksService) Service(path string) *OpenstackNetworksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Network_service(path)
    return self.Network_service(path[:index]).Service(path[index + 1:])
}

func (op *OpenstackNetworksService) String() string {
    return fmt.Sprintf("OpenstackNetworksService:%s", op.Path)
}


//
//
type OpenstackSubnetService struct {
    Service

}

func NewOpenstackSubnetService(connection *Connection, path string) *OpenstackSubnetService {
    var result OpenstackSubnetService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *OpenstackSubnetService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackSubnetService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackSubnetService) Service(path string) *OpenstackSubnetService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *OpenstackSubnetService) String() string {
    return fmt.Sprintf("OpenstackSubnetService:%s", op.Path)
}


//
//
type OpenstackSubnetsService struct {
    Service

    SubnetService  *SubnetService
}

func NewOpenstackSubnetsService(connection *Connection, path string) *OpenstackSubnetsService {
    var result OpenstackSubnetsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *OpenstackSubnetsService) Add (
    subnet OpenStackSubnet,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(subnet, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of sub-networks to return. If not specified all the sub-networks are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackSubnetsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackSubnetsService) Service(path string) *OpenstackSubnetsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Subnet_service(path)
    return self.Subnet_service(path[:index]).Service(path[index + 1:])
}

func (op *OpenstackSubnetsService) String() string {
    return fmt.Sprintf("OpenstackSubnetsService:%s", op.Path)
}


//
//
type OpenstackVolumeAuthenticationKeyService struct {
    Service

}

func NewOpenstackVolumeAuthenticationKeyService(connection *Connection, path string) *OpenstackVolumeAuthenticationKeyService {
    var result OpenstackVolumeAuthenticationKeyService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *OpenstackVolumeAuthenticationKeyService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackVolumeAuthenticationKeyService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
//
func (op *OpenstackVolumeAuthenticationKeyService) Update (
    key OpenstackVolumeAuthenticationKey,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalUpdate(key, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackVolumeAuthenticationKeyService) Service(path string) *OpenstackVolumeAuthenticationKeyService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *OpenstackVolumeAuthenticationKeyService) String() string {
    return fmt.Sprintf("OpenstackVolumeAuthenticationKeyService:%s", op.Path)
}


//
//
type OpenstackVolumeAuthenticationKeysService struct {
    Service

    KeyService  *KeyService
}

func NewOpenstackVolumeAuthenticationKeysService(connection *Connection, path string) *OpenstackVolumeAuthenticationKeysService {
    var result OpenstackVolumeAuthenticationKeysService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *OpenstackVolumeAuthenticationKeysService) Add (
    key OpenstackVolumeAuthenticationKey,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(key, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of keys to return. If not specified all the keys are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackVolumeAuthenticationKeysService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackVolumeAuthenticationKeysService) Service(path string) *OpenstackVolumeAuthenticationKeysService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Key_service(path)
    return self.Key_service(path[:index]).Service(path[index + 1:])
}

func (op *OpenstackVolumeAuthenticationKeysService) String() string {
    return fmt.Sprintf("OpenstackVolumeAuthenticationKeysService:%s", op.Path)
}


//
//
type OpenstackVolumeProviderService struct {
    ExternalProviderService

    AuthenticationKeysService  *AuthenticationKeysService
    CertificatesService  *CertificatesService
    VolumeTypesService  *VolumeTypesService
}

func NewOpenstackVolumeProviderService(connection *Connection, path string) *OpenstackVolumeProviderService {
    var result OpenstackVolumeProviderService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *OpenstackVolumeProviderService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
//
func (op *OpenstackVolumeProviderService) ImportCertificates (
    certificates []*Certificate,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Certificates: certificates,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'importcertificates', nil, headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackVolumeProviderService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the test should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackVolumeProviderService) TestConnectivity (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'testconnectivity', nil, headers, query, wait)
}

//
//
func (op *OpenstackVolumeProviderService) Update (
    provider OpenStackVolumeProvider,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(provider, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackVolumeProviderService) Service(path string) *OpenstackVolumeProviderService {
    if path == nil {
        return op
    }
    if path == "authenticationkeys" {
        return &AuthenticationKeysService{}
    }
    if path.startswith('authenticationkeys/') {
        return AuthenticationKeysService().service(path[19:])
    }
    if path == "certificates" {
        return &CertificatesService{}
    }
    if path.startswith('certificates/') {
        return CertificatesService().service(path[13:])
    }
    if path == "volumetypes" {
        return &VolumeTypesService{}
    }
    if path.startswith('volumetypes/') {
        return VolumeTypesService().service(path[12:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *OpenstackVolumeProviderService) String() string {
    return fmt.Sprintf("OpenstackVolumeProviderService:%s", op.Path)
}


//
//
type OpenstackVolumeProvidersService struct {
    Service

    ProviderService  *ProviderService
}

func NewOpenstackVolumeProvidersService(connection *Connection, path string) *OpenstackVolumeProvidersService {
    var result OpenstackVolumeProvidersService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Adds a new volume provider.
// For example:
// [source]
// ----
// POST /ovirt-engine/api/openstackvolumeproviders
// ----
// With a request body like this:
// [source,xml]
// ----
// <openstack_volume_provider>
//   <name>mycinder</name>
//   <url>https://mycinder.example.com:8776</url>
//   <data_center>
//     <name>mydc</name>
//   </data_center>
//   <requires_authentication>true</requires_authentication>
//   <username>admin</username>
//   <password>mypassword</password>
//   <tenant_name>mytenant</tenant_name>
// </openstack_volume_provider>
// ----
//
func (op *OpenstackVolumeProvidersService) Add (
    provider OpenStackVolumeProvider,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(provider, headers, query, wait)
}

//
// Retrieves the list of volume providers.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of providers to return. If not specified all the providers are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackVolumeProvidersService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackVolumeProvidersService) Service(path string) *OpenstackVolumeProvidersService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Provider_service(path)
    return self.Provider_service(path[:index]).Service(path[index + 1:])
}

func (op *OpenstackVolumeProvidersService) String() string {
    return fmt.Sprintf("OpenstackVolumeProvidersService:%s", op.Path)
}


//
//
type OpenstackVolumeTypeService struct {
    Service

}

func NewOpenstackVolumeTypeService(connection *Connection, path string) *OpenstackVolumeTypeService {
    var result OpenstackVolumeTypeService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *OpenstackVolumeTypeService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackVolumeTypeService) Service(path string) *OpenstackVolumeTypeService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *OpenstackVolumeTypeService) String() string {
    return fmt.Sprintf("OpenstackVolumeTypeService:%s", op.Path)
}


//
//
type OpenstackVolumeTypesService struct {
    Service

    TypeService  *TypeService
}

func NewOpenstackVolumeTypesService(connection *Connection, path string) *OpenstackVolumeTypesService {
    var result OpenstackVolumeTypesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of volume types to return. If not specified all the volume types are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OpenstackVolumeTypesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OpenstackVolumeTypesService) Service(path string) *OpenstackVolumeTypesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Type_service(path)
    return self.Type_service(path[:index]).Service(path[index + 1:])
}

func (op *OpenstackVolumeTypesService) String() string {
    return fmt.Sprintf("OpenstackVolumeTypesService:%s", op.Path)
}


//
//
type OperatingSystemService struct {
    Service

}

func NewOperatingSystemService(connection *Connection, path string) *OperatingSystemService {
    var result OperatingSystemService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *OperatingSystemService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OperatingSystemService) Service(path string) *OperatingSystemService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *OperatingSystemService) String() string {
    return fmt.Sprintf("OperatingSystemService:%s", op.Path)
}


//
//
type OperatingSystemsService struct {
    Service

    OperatingSystemService  *OperatingSystemService
}

func NewOperatingSystemsService(connection *Connection, path string) *OperatingSystemsService {
    var result OperatingSystemsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of networks to return. If not specified all the networks are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *OperatingSystemsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *OperatingSystemsService) Service(path string) *OperatingSystemsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.OperatingSystem_service(path)
    return self.OperatingSystem_service(path[:index]).Service(path[index + 1:])
}

func (op *OperatingSystemsService) String() string {
    return fmt.Sprintf("OperatingSystemsService:%s", op.Path)
}


//
//
type PermissionService struct {
    Service

}

func NewPermissionService(connection *Connection, path string) *PermissionService {
    var result PermissionService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *PermissionService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *PermissionService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *PermissionService) Service(path string) *PermissionService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *PermissionService) String() string {
    return fmt.Sprintf("PermissionService:%s", op.Path)
}


//
// A service to manage a specific permit of the role.
//
type PermitService struct {
    Service

}

func NewPermitService(connection *Connection, path string) *PermitService {
    var result PermitService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Gets the information about the permit of the role.
// For example to retrieve the information about the permit with the id `456` of the role with the id `123`
// send a request like this:
// ....
// GET /ovirt-engine/api/roles/123/permits/456
// ....
// [source,xml]
// ----
// <permit href="/ovirt-engine/api/roles/123/permits/456" id="456">
//   <name>change_vm_cd</name>
//   <administrative>false</administrative>
//   <role href="/ovirt-engine/api/roles/123" id="123"/>
// </permit>
// ----
//
func (op *PermitService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes the permit from the role.
// For example to remove the permit with id `456` from the role with id `123` send a request like this:
// ....
// DELETE /ovirt-engine/api/roles/123/permits/456
// ....
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *PermitService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *PermitService) Service(path string) *PermitService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *PermitService) String() string {
    return fmt.Sprintf("PermitService:%s", op.Path)
}


//
// Represents a permits sub-collection of the specific role.
//
type PermitsService struct {
    Service

    PermitService  *PermitService
}

func NewPermitsService(connection *Connection, path string) *PermitsService {
    var result PermitsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Adds a permit to the role. The permit name can be retrieved from the <<services/cluster_levels>> service.
// For example to assign a permit `create_vm` to the role with id `123` send a request like this:
// ....
// POST /ovirt-engine/api/roles/123/permits
// ....
// With a request body like this:
// [source,xml]
// ----
// <permit>
//   <name>create_vm</name>
// </permit>
// ----
// This method supports the following parameters:
// `Permit`:: The permit to add.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *PermitsService) Add (
    permit Permit,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(permit, headers, query, wait)
}

//
// List the permits of the role.
// For example to list the permits of the role with the id `123` send a request like this:
// ....
// GET /ovirt-engine/api/roles/123/permits
// ....
// [source,xml]
// ----
// <permits>
//   <permit href="/ovirt-engine/api/roles/123/permits/5" id="5">
//     <name>change_vm_cd</name>
//     <administrative>false</administrative>
//     <role href="/ovirt-engine/api/roles/123" id="123"/>
//   </permit>
//   <permit href="/ovirt-engine/api/roles/123/permits/7" id="7">
//     <name>connect_to_vm</name>
//     <administrative>false</administrative>
//     <role href="/ovirt-engine/api/roles/123" id="123"/>
//   </permit>
// </permits>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of permits to return. If not specified all the permits are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *PermitsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *PermitsService) Service(path string) *PermitsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Permit_service(path)
    return self.Permit_service(path[:index]).Service(path[index + 1:])
}

func (op *PermitsService) String() string {
    return fmt.Sprintf("PermitsService:%s", op.Path)
}


//
//
type QosService struct {
    Service

}

func NewQosService(connection *Connection, path string) *QosService {
    var result QosService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *QosService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *QosService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
//
func (op *QosService) Update (
    qos Qos,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(qos, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *QosService) Service(path string) *QosService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *QosService) String() string {
    return fmt.Sprintf("QosService:%s", op.Path)
}


//
//
type QossService struct {
    Service

    QosService  *QosService
}

func NewQossService(connection *Connection, path string) *QossService {
    var result QossService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *QossService) Add (
    qos Qos,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(qos, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of QoS descriptors to return. If not specified all the descriptors are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *QossService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *QossService) Service(path string) *QossService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Qos_service(path)
    return self.Qos_service(path[:index]).Service(path[index + 1:])
}

func (op *QossService) String() string {
    return fmt.Sprintf("QossService:%s", op.Path)
}


//
//
type QuotaService struct {
    Service

    PermissionsService  *PermissionsService
    QuotaClusterLimitsService  *QuotaClusterLimitsService
    QuotaStorageLimitsService  *QuotaStorageLimitsService
}

func NewQuotaService(connection *Connection, path string) *QuotaService {
    var result QuotaService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieves a quota.
// An example of retrieving a quota:
// [source]
// ----
// GET /ovirt-engine/api/datacenters/123/quotas/456
// ----
// [source,xml]
// ----
// <quota id="456">
//   <name>myquota</name>
//   <description>My new quota for virtual machines</description>
//   <cluster_hard_limit_pct>20</cluster_hard_limit_pct>
//   <cluster_soft_limit_pct>80</cluster_soft_limit_pct>
//   <storage_hard_limit_pct>20</storage_hard_limit_pct>
//   <storage_soft_limit_pct>80</storage_soft_limit_pct>
// </quota>
// ----
//
func (op *QuotaService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Delete a quota.
// An example of deleting a quota:
// [source]
// ----
// DELETE /ovirt-engine/api/datacenters/123-456/quotas/654-321
// -0472718ab224 HTTP/1.1
// Accept: application/xml
// Content-type: application/xml
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *QuotaService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates a quota.
// An example of updating a quota:
// [source]
// ----
// PUT /ovirt-engine/api/datacenters/123/quotas/456
// ----
// [source,xml]
// ----
// <quota>
//   <cluster_hard_limit_pct>30</cluster_hard_limit_pct>
//   <cluster_soft_limit_pct>70</cluster_soft_limit_pct>
//   <storage_hard_limit_pct>20</storage_hard_limit_pct>
//   <storage_soft_limit_pct>80</storage_soft_limit_pct>
// </quota>
// ----
//
func (op *QuotaService) Update (
    quota Quota,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(quota, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *QuotaService) Service(path string) *QuotaService {
    if path == nil {
        return op
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    if path == "quotaclusterlimits" {
        return &QuotaClusterLimitsService{}
    }
    if path.startswith('quotaclusterlimits/') {
        return QuotaClusterLimitsService().service(path[19:])
    }
    if path == "quotastoragelimits" {
        return &QuotaStorageLimitsService{}
    }
    if path.startswith('quotastoragelimits/') {
        return QuotaStorageLimitsService().service(path[19:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *QuotaService) String() string {
    return fmt.Sprintf("QuotaService:%s", op.Path)
}


//
//
type QuotaClusterLimitService struct {
    Service

}

func NewQuotaClusterLimitService(connection *Connection, path string) *QuotaClusterLimitService {
    var result QuotaClusterLimitService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *QuotaClusterLimitService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *QuotaClusterLimitService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *QuotaClusterLimitService) Service(path string) *QuotaClusterLimitService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *QuotaClusterLimitService) String() string {
    return fmt.Sprintf("QuotaClusterLimitService:%s", op.Path)
}


//
//
type QuotaClusterLimitsService struct {
    Service

    LimitService  *LimitService
}

func NewQuotaClusterLimitsService(connection *Connection, path string) *QuotaClusterLimitsService {
    var result QuotaClusterLimitsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *QuotaClusterLimitsService) Add (
    limit QuotaClusterLimit,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(limit, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of limits to return. If not specified all the limits are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *QuotaClusterLimitsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *QuotaClusterLimitsService) Service(path string) *QuotaClusterLimitsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Limit_service(path)
    return self.Limit_service(path[:index]).Service(path[index + 1:])
}

func (op *QuotaClusterLimitsService) String() string {
    return fmt.Sprintf("QuotaClusterLimitsService:%s", op.Path)
}


//
//
type QuotaStorageLimitService struct {
    Service

}

func NewQuotaStorageLimitService(connection *Connection, path string) *QuotaStorageLimitService {
    var result QuotaStorageLimitService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *QuotaStorageLimitService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the update should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *QuotaStorageLimitService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *QuotaStorageLimitService) Service(path string) *QuotaStorageLimitService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *QuotaStorageLimitService) String() string {
    return fmt.Sprintf("QuotaStorageLimitService:%s", op.Path)
}


//
//
type QuotaStorageLimitsService struct {
    Service

    LimitService  *LimitService
}

func NewQuotaStorageLimitsService(connection *Connection, path string) *QuotaStorageLimitsService {
    var result QuotaStorageLimitsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *QuotaStorageLimitsService) Add (
    limit QuotaStorageLimit,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(limit, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of limits to return. If not specified all the limits are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *QuotaStorageLimitsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *QuotaStorageLimitsService) Service(path string) *QuotaStorageLimitsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Limit_service(path)
    return self.Limit_service(path[:index]).Service(path[index + 1:])
}

func (op *QuotaStorageLimitsService) String() string {
    return fmt.Sprintf("QuotaStorageLimitsService:%s", op.Path)
}


//
//
type QuotasService struct {
    Service

    QuotaService  *QuotaService
}

func NewQuotasService(connection *Connection, path string) *QuotasService {
    var result QuotasService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a new quota.
// An example of creating a new quota:
// [source]
// ----
// POST /ovirt-engine/api/datacenters/123/quotas
// ----
// [source,xml]
// ----
// <quota>
//   <name>myquota</name>
//   <description>My new quota for virtual machines</description>
// </quota>
// ----
//
func (op *QuotasService) Add (
    quota Quota,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(quota, headers, query, wait)
}

//
// Lists quotas of a data center
// This method supports the following parameters:
// `Max`:: Sets the maximum number of quota descriptors to return. If not specified all the descriptors are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *QuotasService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *QuotasService) Service(path string) *QuotasService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Quota_service(path)
    return self.Quota_service(path[:index]).Service(path[index + 1:])
}

func (op *QuotasService) String() string {
    return fmt.Sprintf("QuotasService:%s", op.Path)
}


//
//
type RoleService struct {
    Service

    PermitsService  *PermitsService
}

func NewRoleService(connection *Connection, path string) *RoleService {
    var result RoleService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Get the role.
// [source]
// ----
// GET /ovirt-engine/api/roles/123
// ----
// You will receive XML response like this one:
// [source,xml]
// ----
// <role id="123">
//   <name>MyRole</name>
//   <description>MyRole description</description>
//   <link href="/ovirt-engine/api/roles/123/permits" rel="permits"/>
//   <administrative>true</administrative>
//   <mutable>false</mutable>
// </role>
// ----
//
func (op *RoleService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes the role.
// To remove the role you need to know its id, then send request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/roles/{role_id}
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *RoleService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates a role. You are allowed to update `name`, `description` and `administrative` attributes after role is
// created. Within this endpoint you can't add or remove roles permits you need to use
// <<services/permits, service>> that manages permits of role.
// For example to update role's `name`, `description` and `administrative` attributes send a request like this:
// [source]
// ----
// PUT /ovirt-engine/api/roles/123
// ----
// With a request body like this:
// [source,xml]
// ----
// <role>
//   <name>MyNewRoleName</name>
//   <description>My new description of the role</description>
//   <administrative>true</administrative>
// </group>
// ----
// This method supports the following parameters:
// `Role`:: Updated role.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *RoleService) Update (
    role Role,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(role, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *RoleService) Service(path string) *RoleService {
    if path == nil {
        return op
    }
    if path == "permits" {
        return &PermitsService{}
    }
    if path.startswith('permits/') {
        return PermitsService().service(path[8:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *RoleService) String() string {
    return fmt.Sprintf("RoleService:%s", op.Path)
}


//
// Provides read-only access to the global set of roles
//
type RolesService struct {
    Service

    RoleService  *RoleService
}

func NewRolesService(connection *Connection, path string) *RolesService {
    var result RolesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Create a new role. The role can be administrative or non-administrative and can have different permits.
// For example, to add the `MyRole` non-administrative role with permits to login and create virtual machines
// send a request like this (note that you have to pass permit id):
// [source]
// ----
// POST /ovirt-engine/api/roles
// ----
// With a request body like this:
// [source,xml]
// ----
// <role>
//   <name>MyRole</name>
//   <description>My custom role to create virtual machines</description>
//   <administrative>false</administrative>
//   <permits>
//     <permit id="1"/>
//     <permit id="1300"/>
//   </permits>
// </group>
// ----
// This method supports the following parameters:
// `Role`:: Role that will be added.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *RolesService) Add (
    role Role,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(role, headers, query, wait)
}

//
// List roles.
// [source]
// ----
// GET /ovirt-engine/api/roles
// ----
// You will receive response in XML like this one:
// [source,xml]
// ----
// <roles>
//   <role id="123">
//      <name>SuperUser</name>
//      <description>Roles management administrator</description>
//      <link href="/ovirt-engine/api/roles/123/permits" rel="permits"/>
//      <administrative>true</administrative>
//      <mutable>false</mutable>
//   </role>
//   ...
// </roles>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of roles to return. If not specified all the roles are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *RolesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *RolesService) Service(path string) *RolesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Role_service(path)
    return self.Role_service(path[:index]).Service(path[index + 1:])
}

func (op *RolesService) String() string {
    return fmt.Sprintf("RolesService:%s", op.Path)
}


//
//
type SchedulingPoliciesService struct {
    Service

    PolicyService  *PolicyService
}

func NewSchedulingPoliciesService(connection *Connection, path string) *SchedulingPoliciesService {
    var result SchedulingPoliciesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *SchedulingPoliciesService) Add (
    policy SchedulingPolicy,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(policy, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of policies to return. If not specified all the policies are returned.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SchedulingPoliciesService) List (
    filter bool,
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SchedulingPoliciesService) Service(path string) *SchedulingPoliciesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Policy_service(path)
    return self.Policy_service(path[:index]).Service(path[index + 1:])
}

func (op *SchedulingPoliciesService) String() string {
    return fmt.Sprintf("SchedulingPoliciesService:%s", op.Path)
}


//
//
type SchedulingPolicyService struct {
    Service

    BalancesService  *BalancesService
    FiltersService  *FiltersService
    WeightsService  *WeightsService
}

func NewSchedulingPolicyService(connection *Connection, path string) *SchedulingPolicyService {
    var result SchedulingPolicyService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SchedulingPolicyService) Get (
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SchedulingPolicyService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
//
func (op *SchedulingPolicyService) Update (
    policy SchedulingPolicy,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(policy, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SchedulingPolicyService) Service(path string) *SchedulingPolicyService {
    if path == nil {
        return op
    }
    if path == "balances" {
        return &BalancesService{}
    }
    if path.startswith('balances/') {
        return BalancesService().service(path[9:])
    }
    if path == "filters" {
        return &FiltersService{}
    }
    if path.startswith('filters/') {
        return FiltersService().service(path[8:])
    }
    if path == "weights" {
        return &WeightsService{}
    }
    if path.startswith('weights/') {
        return WeightsService().service(path[8:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *SchedulingPolicyService) String() string {
    return fmt.Sprintf("SchedulingPolicyService:%s", op.Path)
}


//
//
type SchedulingPolicyUnitService struct {
    Service

}

func NewSchedulingPolicyUnitService(connection *Connection, path string) *SchedulingPolicyUnitService {
    var result SchedulingPolicyUnitService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SchedulingPolicyUnitService) Get (
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SchedulingPolicyUnitService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SchedulingPolicyUnitService) Service(path string) *SchedulingPolicyUnitService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *SchedulingPolicyUnitService) String() string {
    return fmt.Sprintf("SchedulingPolicyUnitService:%s", op.Path)
}


//
//
type SchedulingPolicyUnitsService struct {
    Service

    UnitService  *UnitService
}

func NewSchedulingPolicyUnitsService(connection *Connection, path string) *SchedulingPolicyUnitsService {
    var result SchedulingPolicyUnitsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of policy units to return. If not specified all the policy units are returned.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SchedulingPolicyUnitsService) List (
    filter bool,
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SchedulingPolicyUnitsService) Service(path string) *SchedulingPolicyUnitsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Unit_service(path)
    return self.Unit_service(path[:index]).Service(path[index + 1:])
}

func (op *SchedulingPolicyUnitsService) String() string {
    return fmt.Sprintf("SchedulingPolicyUnitsService:%s", op.Path)
}


//
//
type SnapshotService struct {
    Service

    CdromsService  *CdromsService
    DisksService  *DisksService
    NicsService  *NicsService
}

func NewSnapshotService(connection *Connection, path string) *SnapshotService {
    var result SnapshotService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *SnapshotService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `AllContent`:: Indicates if all the attributes of the virtual machine snapshot should be included in the response.
// By default the attribute `initialization.configuration.data` is excluded.
// For example, to retrieve the complete representation of the snapshot with id `456` of the virtual machine
// with id `123` send a request like this:
// ....
// GET /ovirt-engine/api/vms/123/snapshots/456?all_content=true
// ....
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SnapshotService) Remove (
    async bool,
    allContent bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }
    if allContent != nil {
        query["all_content"] = allContent
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Restores a virtual machine snapshot.
// For example, to restore the snapshot with identifier `456` of virtual machine with identifier `123` send a
// request like this:
// [source]
// ----
// POST /ovirt-engine/api/vms/123/snapshots/456/restore
// ----
// With an empty `action` in the body:
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the restore should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SnapshotService) Restore (
    async bool,
    disks []*Disk,
    restoreMemory bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Disks: disks,
        RestoreMemory: restoreMemory,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'restore', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SnapshotService) Service(path string) *SnapshotService {
    if path == nil {
        return op
    }
    if path == "cdroms" {
        return &CdromsService{}
    }
    if path.startswith('cdroms/') {
        return CdromsService().service(path[7:])
    }
    if path == "disks" {
        return &DisksService{}
    }
    if path.startswith('disks/') {
        return DisksService().service(path[6:])
    }
    if path == "nics" {
        return &NicsService{}
    }
    if path.startswith('nics/') {
        return NicsService().service(path[5:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *SnapshotService) String() string {
    return fmt.Sprintf("SnapshotService:%s", op.Path)
}


//
//
type SnapshotCdromService struct {
    Service

}

func NewSnapshotCdromService(connection *Connection, path string) *SnapshotCdromService {
    var result SnapshotCdromService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *SnapshotCdromService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SnapshotCdromService) Service(path string) *SnapshotCdromService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *SnapshotCdromService) String() string {
    return fmt.Sprintf("SnapshotCdromService:%s", op.Path)
}


//
//
type SnapshotCdromsService struct {
    Service

    CdromService  *CdromService
}

func NewSnapshotCdromsService(connection *Connection, path string) *SnapshotCdromsService {
    var result SnapshotCdromsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of CDROMS to return. If not specified all the CDROMS are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SnapshotCdromsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SnapshotCdromsService) Service(path string) *SnapshotCdromsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Cdrom_service(path)
    return self.Cdrom_service(path[:index]).Service(path[index + 1:])
}

func (op *SnapshotCdromsService) String() string {
    return fmt.Sprintf("SnapshotCdromsService:%s", op.Path)
}


//
//
type SnapshotDiskService struct {
    Service

}

func NewSnapshotDiskService(connection *Connection, path string) *SnapshotDiskService {
    var result SnapshotDiskService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *SnapshotDiskService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SnapshotDiskService) Service(path string) *SnapshotDiskService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *SnapshotDiskService) String() string {
    return fmt.Sprintf("SnapshotDiskService:%s", op.Path)
}


//
//
type SnapshotDisksService struct {
    Service

    DiskService  *DiskService
}

func NewSnapshotDisksService(connection *Connection, path string) *SnapshotDisksService {
    var result SnapshotDisksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of disks to return. If not specified all the disks are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SnapshotDisksService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SnapshotDisksService) Service(path string) *SnapshotDisksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Disk_service(path)
    return self.Disk_service(path[:index]).Service(path[index + 1:])
}

func (op *SnapshotDisksService) String() string {
    return fmt.Sprintf("SnapshotDisksService:%s", op.Path)
}


//
//
type SnapshotNicService struct {
    Service

}

func NewSnapshotNicService(connection *Connection, path string) *SnapshotNicService {
    var result SnapshotNicService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *SnapshotNicService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SnapshotNicService) Service(path string) *SnapshotNicService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *SnapshotNicService) String() string {
    return fmt.Sprintf("SnapshotNicService:%s", op.Path)
}


//
//
type SnapshotNicsService struct {
    Service

    NicService  *NicService
}

func NewSnapshotNicsService(connection *Connection, path string) *SnapshotNicsService {
    var result SnapshotNicsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of NICs to return. If not specified all the NICs are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SnapshotNicsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SnapshotNicsService) Service(path string) *SnapshotNicsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Nic_service(path)
    return self.Nic_service(path[:index]).Service(path[index + 1:])
}

func (op *SnapshotNicsService) String() string {
    return fmt.Sprintf("SnapshotNicsService:%s", op.Path)
}


//
//
type SnapshotsService struct {
    Service

    SnapshotService  *SnapshotService
}

func NewSnapshotsService(connection *Connection, path string) *SnapshotsService {
    var result SnapshotsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a virtual machine snapshot.
// For example, to create a new snapshot for virtual machine `123` send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/vms/123/snapshots
// ----
// With a request body like this:
// [source,xml]
// ----
// <snapshot>
//   <description>My snapshot</description>
// </snapshot>
// ----
//
func (op *SnapshotsService) Add (
    snapshot Snapshot,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(snapshot, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of snapshots to return. If not specified all the snapshots are returned.
// `AllContent`:: Indicates if all the attributes of the virtual machine snapshot should be included in the response.
// By default the attribute `initialization.configuration.data` is excluded.
// For example, to retrieve the complete representation of the virtual machine with id `123` snapshots send a
// request like this:
// ....
// GET /ovirt-engine/api/vms/123/snapshots?all_content=true
// ....
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SnapshotsService) List (
    allContent bool,
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if allContent != nil {
        query["all_content"] = allContent
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SnapshotsService) Service(path string) *SnapshotsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Snapshot_service(path)
    return self.Snapshot_service(path[:index]).Service(path[index + 1:])
}

func (op *SnapshotsService) String() string {
    return fmt.Sprintf("SnapshotsService:%s", op.Path)
}


//
//
type SshPublicKeyService struct {
    Service

}

func NewSshPublicKeyService(connection *Connection, path string) *SshPublicKeyService {
    var result SshPublicKeyService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *SshPublicKeyService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SshPublicKeyService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
//
func (op *SshPublicKeyService) Update (
    key SshPublicKey,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(key, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SshPublicKeyService) Service(path string) *SshPublicKeyService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *SshPublicKeyService) String() string {
    return fmt.Sprintf("SshPublicKeyService:%s", op.Path)
}


//
//
type SshPublicKeysService struct {
    Service

    KeyService  *KeyService
}

func NewSshPublicKeysService(connection *Connection, path string) *SshPublicKeysService {
    var result SshPublicKeysService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *SshPublicKeysService) Add (
    key SshPublicKey,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(key, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of keys to return. If not specified all the keys are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SshPublicKeysService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SshPublicKeysService) Service(path string) *SshPublicKeysService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Key_service(path)
    return self.Key_service(path[:index]).Service(path[index + 1:])
}

func (op *SshPublicKeysService) String() string {
    return fmt.Sprintf("SshPublicKeysService:%s", op.Path)
}


//
//
type StatisticService struct {
    Service

}

func NewStatisticService(connection *Connection, path string) *StatisticService {
    var result StatisticService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *StatisticService) Get (
    statistic Statistic,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if statistic != nil {
        query["statistic"] = statistic
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StatisticService) Service(path string) *StatisticService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *StatisticService) String() string {
    return fmt.Sprintf("StatisticService:%s", op.Path)
}


//
//
type StatisticsService struct {
    Service

    StatisticService  *StatisticService
}

func NewStatisticsService(connection *Connection, path string) *StatisticsService {
    var result StatisticsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieves a list of statistics.
// For example, to retrieve the statistics for virtual machine `123` send a
// request like this:
// [source]
// ----
// GET /ovirt-engine/api/vms/123/statistics
// ----
// The result will be like this:
// [source,xml]
// ----
// <statistics>
//   <statistic href="/ovirt-engine/api/vms/123/statistics/456" id="456">
//     <name>memory.installed</name>
//     <description>Total memory configured</description>
//     <kind>gauge</kind>
//     <type>integer</type>
//     <unit>bytes</unit>
//     <values>
//       <value>
//         <datum>1073741824</datum>
//       </value>
//     </values>
//     <vm href="/ovirt-engine/api/vms/123" id="123"/>
//   </statistic>
//   ...
// </statistics>
// ----
// Just a single part of the statistics can be retrieved by specifying its id at the end of the URI. That means:
// [source]
// ----
// GET /ovirt-engine/api/vms/123/statistics/456
// ----
// Outputs:
// [source,xml]
// ----
// <statistic href="/ovirt-engine/api/vms/123/statistics/456" id="456">
//   <name>memory.installed</name>
//   <description>Total memory configured</description>
//   <kind>gauge</kind>
//   <type>integer</type>
//   <unit>bytes</unit>
//   <values>
//     <value>
//       <datum>1073741824</datum>
//     </value>
//   </values>
//   <vm href="/ovirt-engine/api/vms/123" id="123"/>
// </statistic>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of statistics to return. If not specified all the statistics are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StatisticsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StatisticsService) Service(path string) *StatisticsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Statistic_service(path)
    return self.Statistic_service(path[:index]).Service(path[index + 1:])
}

func (op *StatisticsService) String() string {
    return fmt.Sprintf("StatisticsService:%s", op.Path)
}


//
// A service to manage a step.
//
type StepService struct {
    MeasurableService

    StatisticsService  *StatisticsService
}

func NewStepService(connection *Connection, path string) *StepService {
    var result StepService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Marks an external step execution as ended.
// For example, to terminate a step with identifier `456` which belongs to a `job` with identifier `123` send the
// following request:
// [source]
// ----
// POST /ovirt-engine/api/jobs/123/steps/456/end
// ----
// With the following request body:
// [source,xml]
// ----
// <action>
//   <force>true</force>
//   <succeeded>true</succeeded>
// </action>
// ----
// This method supports the following parameters:
// `Force`:: Indicates if the step should be forcibly terminated.
// `Succeeded`:: Indicates if the step should be marked as successfully finished or as failed.
// This parameter is optional, and the default value is `true`.
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StepService) End (
    async bool,
    force bool,
    succeeded bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Force: force,
        Succeeded: succeeded,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'end', nil, headers, query, wait)
}

//
// Retrieves a step.
// [source]
// ----
// GET /ovirt-engine/api/jobs/123/steps/456
// ----
// You will receive response in XML like this one:
// [source,xml]
// ----
// <step href="/ovirt-engine/api/jobs/123/steps/456" id="456">
//   <actions>
//     <link href="/ovirt-engine/api/jobs/123/steps/456/end" rel="end"/>
//   </actions>
//   <description>Validating</description>
//   <end_time>2016-12-12T23:07:26.627+02:00</end_time>
//   <external>false</external>
//   <number>0</number>
//   <start_time>2016-12-12T23:07:26.605+02:00</start_time>
//   <status>finished</status>
//   <type>validating</type>
//   <job href="/ovirt-engine/api/jobs/123" id="123"/>
// </step>
// ----
//
func (op *StepService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StepService) Service(path string) *StepService {
    if path == nil {
        return op
    }
    if path == "statistics" {
        return &StatisticsService{}
    }
    if path.startswith('statistics/') {
        return StatisticsService().service(path[11:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *StepService) String() string {
    return fmt.Sprintf("StepService:%s", op.Path)
}


//
// A service to manage steps.
//
type StepsService struct {
    Service

    StepService  *StepService
}

func NewStepsService(connection *Connection, path string) *StepsService {
    var result StepsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add an external step to an existing job or to an existing step.
// For example, to add a step to `job` with identifier `123` send the
// following request:
// [source]
// ----
// POST /ovirt-engine/api/jobs/123/steps
// ----
// With the following request body:
// [source,xml]
// ----
// <step>
//   <description>Validating</description>
//   <start_time>2016-12-12T23:07:26.605+02:00</start_time>
//   <status>started</status>
//   <type>validating</type>
// </step>
// ----
// The response should look like:
// [source,xml]
// ----
// <step href="/ovirt-engine/api/jobs/123/steps/456" id="456">
//   <actions>
//     <link href="/ovirt-engine/api/jobs/123/steps/456/end" rel="end"/>
//   </actions>
//   <description>Validating</description>
//   <link href="/ovirt-engine/api/jobs/123/steps/456/statistics" rel="statistics"/>
//   <external>true</external>
//   <number>2</number>
//   <start_time>2016-12-13T01:06:15.380+02:00</start_time>
//   <status>started</status>
//   <type>validating</type>
//   <job href="/ovirt-engine/api/jobs/123" id="123"/>
// </step>
// ----
// This method supports the following parameters:
// `Step`:: Step that will be added.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StepsService) Add (
    step Step,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(step, headers, query, wait)
}

//
// Retrieves the representation of the steps.
// [source]
// ----
// GET /ovirt-engine/api/job/123/steps
// ----
// You will receive response in XML like this one:
// [source,xml]
// ----
// <steps>
//   <step href="/ovirt-engine/api/jobs/123/steps/456" id="456">
//     <actions>
//       <link href="/ovirt-engine/api/jobs/123/steps/456/end" rel="end"/>
//     </actions>
//     <description>Validating</description>
//     <link href="/ovirt-engine/api/jobs/123/steps/456/statistics" rel="statistics"/>
//     <external>true</external>
//     <number>2</number>
//     <start_time>2016-12-13T01:06:15.380+02:00</start_time>
//     <status>started</status>
//     <type>validating</type>
//     <job href="/ovirt-engine/api/jobs/123" id="123"/>
//   </step>
//   ...
// </steps>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of steps to return. If not specified all the steps are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StepsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StepsService) Service(path string) *StepsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Step_service(path)
    return self.Step_service(path[:index]).Service(path[index + 1:])
}

func (op *StepsService) String() string {
    return fmt.Sprintf("StepsService:%s", op.Path)
}


//
//
type StorageService struct {
    Service

}

func NewStorageService(connection *Connection, path string) *StorageService {
    var result StorageService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `ReportStatus`:: Indicates if the status of the LUNs in the storage should be checked.
// Checking the status of the LUN is an heavy weight operation and
// this data is not always needed by the user.
// This parameter will give the option to not perform the status check of the LUNs.
// The default is `true` for backward compatibility.
// Here an example with the LUN status :
// [source,xml]
// ----
// <host_storage id="360014051136c20574f743bdbd28177fd">
//   <logical_units>
//     <logical_unit id="360014051136c20574f743bdbd28177fd">
//       <lun_mapping>0</lun_mapping>
//       <paths>1</paths>
//       <product_id>lun0</product_id>
//       <serial>SLIO-ORG_lun0_1136c205-74f7-43bd-bd28-177fd5ce6993</serial>
//       <size>10737418240</size>
//       <status>used</status>
//       <vendor_id>LIO-ORG</vendor_id>
//       <volume_group_id>O9Du7I-RahN-ECe1-dZ1w-nh0b-64io-MNzIBZ</volume_group_id>
//     </logical_unit>
//   </logical_units>
//   <type>iscsi</type>
//   <host id="8bb5ade5-e988-4000-8b93-dbfc6717fe50"/>
// </host_storage>
// ----
// Here an example without the LUN status :
// [source,xml]
// ----
// <host_storage id="360014051136c20574f743bdbd28177fd">
//   <logical_units>
//     <logical_unit id="360014051136c20574f743bdbd28177fd">
//       <lun_mapping>0</lun_mapping>
//       <paths>1</paths>
//       <product_id>lun0</product_id>
//       <serial>SLIO-ORG_lun0_1136c205-74f7-43bd-bd28-177fd5ce6993</serial>
//       <size>10737418240</size>
//       <vendor_id>LIO-ORG</vendor_id>
//       <volume_group_id>O9Du7I-RahN-ECe1-dZ1w-nh0b-64io-MNzIBZ</volume_group_id>
//     </logical_unit>
//   </logical_units>
//   <type>iscsi</type>
//   <host id="8bb5ade5-e988-4000-8b93-dbfc6717fe50"/>
// </host_storage>
// ----
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageService) Get (
    reportStatus bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if reportStatus != nil {
        query["report_status"] = reportStatus
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageService) Service(path string) *StorageService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *StorageService) String() string {
    return fmt.Sprintf("StorageService:%s", op.Path)
}


//
//
type StorageDomainService struct {
    Service

    DiskProfilesService  *DiskProfilesService
    DiskSnapshotsService  *DiskSnapshotsService
    DisksService  *DisksService
    FilesService  *FilesService
    ImagesService  *ImagesService
    PermissionsService  *PermissionsService
    StorageConnectionsService  *StorageConnectionsService
    TemplatesService  *TemplatesService
    VmsService  *VmsService
}

func NewStorageDomainService(connection *Connection, path string) *StorageDomainService {
    var result StorageDomainService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainService) Get (
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainService) IsAttached (
    async bool,
    host Host,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Host: host,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'isattached', 'isAttached', headers, query, wait)
}

//
// This operation reduces logical units from the storage domain.
// In order to do so the data stored on the provided logical units will be moved to other logical units of the
// storage domain and only then they will be reduced from the storage domain.
// For example, in order to reduce two logical units from a storage domain send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/storagedomains/123/reduceluns
// ----
// With a request body like this:
// [source,xml]
// ----
//  <action>
//    <logical_units>
//      <logical_unit id="1IET_00010001"/>
//      <logical_unit id="1IET_00010002"/>
//    </logical_units>
//  </action>
// ----
// This method supports the following parameters:
// `LogicalUnits`:: The logical units that needs to be reduced from the storage domain.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainService) ReduceLuns (
    logicalUnits []*LogicalUnit,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        LogicalUnits: logicalUnits,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'reduceluns', nil, headers, query, wait)
}

//
// This operation refreshes the LUN size.
// After increasing the size of the underlying LUN on the storage server,
// the user can refresh the LUN size.
// This action forces a rescan of the provided LUNs and
// updates the database with the new size if required.
// For example, in order to refresh the size of two LUNs send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/storagedomains/262b056b-aede-40f1-9666-b883eff59d40/refreshluns
// ----
// With a request body like this:
// [source,xml]
// ----
//  <action>
//    <logical_units>
//      <logical_unit id="1IET_00010001"/>
//      <logical_unit id="1IET_00010002"/>
//    </logical_units>
//  </action>
// ----
// This method supports the following parameters:
// `LogicalUnits`:: The LUNs that need to be refreshed.
// `Async`:: Indicates if the refresh should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainService) RefreshLuns (
    async bool,
    logicalUnits []*LogicalUnit,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        LogicalUnits: logicalUnits,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'refreshluns', nil, headers, query, wait)
}

//
// Removes the storage domain.
// Without any special parameters, the storage domain is detached from the system and removed from the database. The
// storage domain can then be imported to the same or different setup, with all the data on it. If the storage isn't
// accessible the operation will fail.
// If the `destroy` parameter is `true` then the operation will always succeed, even if the storage isn't
// accessible, the failure is just ignored and the storage domain is removed from the database anyway.
// If the `format` parameter is `true` then the actual storage is formatted, and the metadata is removed from the
// LUN or directory, so it can no longer be imported to the same or a different setup.
// This method supports the following parameters:
// `Host`:: Indicates what host should be used to remove the storage domain.
// This parameter is mandatory, and it can contain the name or the identifier of the host. For example, to use
// the host named `myhost` to remove the storage domain with identifier `123` send a request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/storagedomains/123?host=myhost
// ----
// `Format`:: Indicates if the actual storage should be formatted, removing all the metadata from the underlying LUN or
// directory:
// [source]
// ----
// DELETE /ovirt-engine/api/storagedomains/123?format=true
// ----
// This parameter is optional, and the default value is `false`.
// `Destroy`:: Indicates if the operation should succeed, and the storage domain removed from the database, even if the
// storage isn't accessible.
// [source]
// ----
// DELETE /ovirt-engine/api/storagedomains/123?destroy=true
// ----
// This parameter is optional, and the default value is `false`.
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainService) Remove (
    host string,
    format bool,
    destroy bool,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if host != nil {
        query["host"] = host
    }
    if format != nil {
        query["format"] = format
    }
    if destroy != nil {
        query["destroy"] = destroy
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates a storage domain.
// Not all of the <<types/storage_domain,StorageDomain>>'s attributes are updatable post-creation. Those that can be
// updated are: `name`, `description`, `comment`, `warning_low_space_indicator`, `critical_space_action_blocker` and
// `wipe_after_delete` (note that changing the `wipe_after_delete` attribute will not change the wipe after delete
// property of disks that already exist).
// To update the `name` and `wipe_after_delete` attributes of a storage domain with an identifier `123`, send a
// request as follows:
// [source]
// ----
// PUT /ovirt-engine/api/storagedomains/123
// ----
// With a request body as follows:
// [source,xml]
// ----
// <storage_domain>
//   <name>data2</name>
//   <wipe_after_delete>true</wipe_after_delete>
// </storage_domain>
// ----
//
func (op *StorageDomainService) Update (
    storageDomain StorageDomain,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(storageDomain, headers, query, wait)
}

//
// This operation forces the update of the `OVF_STORE`
// of this storage domain.
// The `OVF_STORE` is a disk image that contains the meta-data
// of virtual machines and disks that reside in the
// storage domain. This meta-data is used in case the
// domain is imported or exported to or from a different
// data center or a different installation.
// By default the `OVF_STORE` is updated periodically
// (set by default to 60 minutes) but users might want to force an
// update after an important change, or when the they believe the
// `OVF_STORE` is corrupt.
// When initiated by the user, `OVF_STORE` update will be performed whether
// an update is needed or not.
// This method supports the following parameters:
// `Async`:: Indicates if the `OVF_STORE` update should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainService) UpdateOvfStore (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'updateovfstore', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageDomainService) Service(path string) *StorageDomainService {
    if path == nil {
        return op
    }
    if path == "diskprofiles" {
        return &DiskProfilesService{}
    }
    if path.startswith('diskprofiles/') {
        return DiskProfilesService().service(path[13:])
    }
    if path == "disksnapshots" {
        return &DiskSnapshotsService{}
    }
    if path.startswith('disksnapshots/') {
        return DiskSnapshotsService().service(path[14:])
    }
    if path == "disks" {
        return &DisksService{}
    }
    if path.startswith('disks/') {
        return DisksService().service(path[6:])
    }
    if path == "files" {
        return &FilesService{}
    }
    if path.startswith('files/') {
        return FilesService().service(path[6:])
    }
    if path == "images" {
        return &ImagesService{}
    }
    if path.startswith('images/') {
        return ImagesService().service(path[7:])
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    if path == "storageconnections" {
        return &StorageConnectionsService{}
    }
    if path.startswith('storageconnections/') {
        return StorageConnectionsService().service(path[19:])
    }
    if path == "templates" {
        return &TemplatesService{}
    }
    if path.startswith('templates/') {
        return TemplatesService().service(path[10:])
    }
    if path == "vms" {
        return &VmsService{}
    }
    if path.startswith('vms/') {
        return VmsService().service(path[4:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *StorageDomainService) String() string {
    return fmt.Sprintf("StorageDomainService:%s", op.Path)
}


//
//
type StorageDomainContentDiskService struct {
    Service

}

func NewStorageDomainContentDiskService(connection *Connection, path string) *StorageDomainContentDiskService {
    var result StorageDomainContentDiskService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainContentDiskService) Get (
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageDomainContentDiskService) Service(path string) *StorageDomainContentDiskService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *StorageDomainContentDiskService) String() string {
    return fmt.Sprintf("StorageDomainContentDiskService:%s", op.Path)
}


//
//
type StorageDomainContentDisksService struct {
    Service

    DiskService  *DiskService
}

func NewStorageDomainContentDisksService(connection *Connection, path string) *StorageDomainContentDisksService {
    var result StorageDomainContentDisksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of disks to return. If not specified all the disks are returned.
// `Search`:: A query string used to restrict the returned disks.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainContentDisksService) List (
    caseSensitive bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageDomainContentDisksService) Service(path string) *StorageDomainContentDisksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Disk_service(path)
    return self.Disk_service(path[:index]).Service(path[index + 1:])
}

func (op *StorageDomainContentDisksService) String() string {
    return fmt.Sprintf("StorageDomainContentDisksService:%s", op.Path)
}


//
// Manages a single disk available in a storage domain.
// IMPORTANT: Since version 4.2 of the engine this service is intended only to list disks available in the storage
// domain, and to register unregistered disks. All the other operations, like copying a disk, moving a disk, etc, have
// been deprecated and will be removed in the future. To perform those operations use the <<services/disks, service
// that manages all the disks of the system>>, or the <<services/disk, service that manages an specific disk>>.
//
type StorageDomainDiskService struct {
    MeasurableService

    PermissionsService  *PermissionsService
    StatisticsService  *StatisticsService
}

func NewStorageDomainDiskService(connection *Connection, path string) *StorageDomainDiskService {
    var result StorageDomainDiskService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Copies a disk to the specified storage domain.
// IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
// compatibility. It will be removed in the future. To copy a disk use the <<services/disk/methods/copy, copy>>
// operation of the service that manages that disk.
// This method supports the following parameters:
// `Disk`:: Description of the resulting disk.
// `StorageDomain`:: The storage domain where the new disk will be created.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainDiskService) Copy (
    disk Disk,
    storageDomain StorageDomain,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Disk: disk,
        StorageDomain: storageDomain,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'copy', nil, headers, query, wait)
}

//
// Exports a disk to an export storage domain.
// IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
// compatibility. It will be removed in the future. To export a disk use the <<services/disk/methods/export, export>>
// operation of the service that manages that disk.
// This method supports the following parameters:
// `StorageDomain`:: The export storage domain where the disk should be exported to.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainDiskService) Export (
    storageDomain StorageDomain,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        StorageDomain: storageDomain,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'export', nil, headers, query, wait)
}

//
// Retrieves the description of the disk.
//
func (op *StorageDomainDiskService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Moves a disk to another storage domain.
// IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
// compatibility. It will be removed in the future. To move a disk use the <<services/disk/methods/move, move>>
// operation of the service that manages that disk.
// This method supports the following parameters:
// `StorageDomain`:: The storage domain where the disk will be moved to.
// `Async`:: Indicates if the move should be performed asynchronously.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainDiskService) Move (
    async bool,
    filter bool,
    storageDomain StorageDomain,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Filter: filter,
        StorageDomain: storageDomain,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'move', nil, headers, query, wait)
}

//
// Removes a disk.
// IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
// compatibility. It will be removed in the future. To remove a disk use the <<services/disk/methods/remove, remove>>
// operation of the service that manages that disk.
//
func (op *StorageDomainDiskService) Remove (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Sparsify the disk.
// IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
// compatibility. It will be removed in the future. To remove a disk use the <<services/disk/methods/remove, remove>>
// operation of the service that manages that disk.
//
func (op *StorageDomainDiskService) Sparsify (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
    }

    // Send the request and wait for the response:
    return internalAction(action, 'sparsify', nil, headers, query, wait)
}

//
// Updates the disk.
// IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
// compatibility. It will be removed in the future. To update a disk use the
// <<services/disk/methods/update, update>> operation of the service that manages that disk.
// This method supports the following parameters:
// `Disk`:: The update to apply to the disk.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainDiskService) Update (
    disk Disk,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalUpdate(disk, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageDomainDiskService) Service(path string) *StorageDomainDiskService {
    if path == nil {
        return op
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    if path == "statistics" {
        return &StatisticsService{}
    }
    if path.startswith('statistics/') {
        return StatisticsService().service(path[11:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *StorageDomainDiskService) String() string {
    return fmt.Sprintf("StorageDomainDiskService:%s", op.Path)
}


//
// Manages the collection of disks available inside an specific storage domain.
//
type StorageDomainDisksService struct {
    Service

    DiskService  *DiskService
}

func NewStorageDomainDisksService(connection *Connection, path string) *StorageDomainDisksService {
    var result StorageDomainDisksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Adds or registers a disk.
// IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
// compatibility. It will be removed in the future. To add a new disk use the <<services/disks/methods/add, add>>
// operation of the service that manages the disks of the system. To register an unregistered disk use the
// <<services/attached_storage_domain_disk/methods/register, register>> operation of the service that manages
// that disk.
// This method supports the following parameters:
// `Disk`:: The disk to add or register.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainDisksService) Add (
    disk Disk,
    unregistered bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if unregistered != nil {
        query["unregistered"] = unregistered
    }

    // Send the request
    return op.internalAdd(disk, headers, query, wait)
}

//
// Retrieve the list of disks that are available in the storage domain.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of disks to return. If not specified all the disks are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainDisksService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageDomainDisksService) Service(path string) *StorageDomainDisksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Disk_service(path)
    return self.Disk_service(path[:index]).Service(path[index + 1:])
}

func (op *StorageDomainDisksService) String() string {
    return fmt.Sprintf("StorageDomainDisksService:%s", op.Path)
}


//
//
type StorageDomainServerConnectionService struct {
    Service

}

func NewStorageDomainServerConnectionService(connection *Connection, path string) *StorageDomainServerConnectionService {
    var result StorageDomainServerConnectionService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *StorageDomainServerConnectionService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Detaches a storage connection from storage.
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainServerConnectionService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageDomainServerConnectionService) Service(path string) *StorageDomainServerConnectionService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *StorageDomainServerConnectionService) String() string {
    return fmt.Sprintf("StorageDomainServerConnectionService:%s", op.Path)
}


//
//
type StorageDomainServerConnectionsService struct {
    Service

    ConnectionService  *ConnectionService
}

func NewStorageDomainServerConnectionsService(connection *Connection, path string) *StorageDomainServerConnectionsService {
    var result StorageDomainServerConnectionsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *StorageDomainServerConnectionsService) Add (
    connection StorageConnection,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(connection, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of connections to return. If not specified all the connections are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainServerConnectionsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageDomainServerConnectionsService) Service(path string) *StorageDomainServerConnectionsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Connection_service(path)
    return self.Connection_service(path[:index]).Service(path[index + 1:])
}

func (op *StorageDomainServerConnectionsService) String() string {
    return fmt.Sprintf("StorageDomainServerConnectionsService:%s", op.Path)
}


//
//
type StorageDomainTemplateService struct {
    Service

    DisksService  *DisksService
}

func NewStorageDomainTemplateService(connection *Connection, path string) *StorageDomainTemplateService {
    var result StorageDomainTemplateService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *StorageDomainTemplateService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Action to import a template from an export storage domain.
// For example, to import the template `456` from the storage domain `123` send the following request:
// [source]
// ----
// POST /ovirt-engine/api/storagedomains/123/templates/456/import
// ----
// With the following request body:
// [source, xml]
// ----
// <action>
//   <storage_domain>
//     <name>myexport</name>
//   </storage_domain>
//   <cluster>
//     <name>mycluster</name>
//   </cluster>
// </action>
// ----
// This method supports the following parameters:
// `Clone`:: Use the optional `clone` parameter to generate new UUIDs for the imported template and its entities.
// The user might want to import a template with the `clone` parameter set to `false` when importing a template
// from an export domain, with templates that was exported by a different {product-name} environment.
// `Async`:: Indicates if the import should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainTemplateService) Import (
    async bool,
    clone bool,
    cluster Cluster,
    exclusive bool,
    storageDomain StorageDomain,
    template Template,
    vm Vm,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Clone: clone,
        Cluster: cluster,
        Exclusive: exclusive,
        StorageDomain: storageDomain,
        Template: template,
        Vm: vm,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'import', nil, headers, query, wait)
}

//
// This method supports the following parameters:
// `AllowPartialImport`:: Indicates whether a template is allowed to be registered with only some of its disks.
// If this flag is `true`, the engine will not fail in the validation process if an image is not found, but
// instead it will allow the template to be registered without the missing disks. This is mainly used during
// registration of a template when some of the storage domains are not available. The default value is `false`.
// `Async`:: Indicates if the registration should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainTemplateService) Register (
    allowPartialImport bool,
    async bool,
    clone bool,
    cluster Cluster,
    exclusive bool,
    template Template,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        AllowPartialImport: allowPartialImport,
        Async: async,
        Clone: clone,
        Cluster: cluster,
        Exclusive: exclusive,
        Template: template,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'register', nil, headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainTemplateService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageDomainTemplateService) Service(path string) *StorageDomainTemplateService {
    if path == nil {
        return op
    }
    if path == "disks" {
        return &DisksService{}
    }
    if path.startswith('disks/') {
        return DisksService().service(path[6:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *StorageDomainTemplateService) String() string {
    return fmt.Sprintf("StorageDomainTemplateService:%s", op.Path)
}


//
//
type StorageDomainTemplatesService struct {
    Service

    TemplateService  *TemplateService
}

func NewStorageDomainTemplatesService(connection *Connection, path string) *StorageDomainTemplatesService {
    var result StorageDomainTemplatesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of templates to return. If not specified all the templates are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainTemplatesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageDomainTemplatesService) Service(path string) *StorageDomainTemplatesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Template_service(path)
    return self.Template_service(path[:index]).Service(path[index + 1:])
}

func (op *StorageDomainTemplatesService) String() string {
    return fmt.Sprintf("StorageDomainTemplatesService:%s", op.Path)
}


//
//
type StorageDomainVmService struct {
    Service

    DiskAttachmentsService  *DiskAttachmentsService
    DisksService  *DisksService
}

func NewStorageDomainVmService(connection *Connection, path string) *StorageDomainVmService {
    var result StorageDomainVmService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *StorageDomainVmService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Imports a virtual machine from an export storage domain.
// For example, send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/storagedomains/123/vms/456/import
// ----
// With a request body like this:
// [source,xml]
// ----
// <action>
//   <storage_domain>
//     <name>mydata</name>
//   </storage_domain>
//   <cluster>
//     <name>mycluster</name>
//   </cluster>
// </action>
// ----
// To import a virtual machine as a new entity add the `clone` parameter:
// [source,xml]
// ----
// <action>
//   <storage_domain>
//     <name>mydata</name>
//   </storage_domain>
//   <cluster>
//     <name>mycluster</name>
//   </cluster>
//   <clone>true</clone>
//   <vm>
//     <name>myvm</name>
//   </vm>
// </action>
// ----
// Include an optional `disks` parameter to choose which disks to import. For example, to import the disks
// of the template that have the identifiers `123` and `456` send the following request body:
// [source,xml]
// ----
// <action>
//   <cluster>
//     <name>mycluster</name>
//   </cluster>
//   <vm>
//     <name>myvm</name>
//   </vm>
//   <disks>
//     <disk id="123"/>
//     <disk id="456"/>
//   </disks>
// </action>
// ----
// This method supports the following parameters:
// `Clone`:: Indicates if the identifiers of the imported virtual machine
// should be regenerated.
// By default when a virtual machine is imported the identifiers
// are preserved. This means that the same virtual machine can't
// be imported multiple times, as that identifiers needs to be
// unique. To allow importing the same machine multiple times set
// this parameter to `true`, as the default is `false`.
// `CollapseSnapshots`:: Indicates of the snapshots of the virtual machine that is imported
// should be collapsed, so that the result will be a virtual machine
// without snapshots.
// This parameter is optional, and if it isn't explicitly specified the
// default value is `false`.
// `Async`:: Indicates if the import should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainVmService) Import (
    async bool,
    clone bool,
    cluster Cluster,
    collapseSnapshots bool,
    storageDomain StorageDomain,
    vm Vm,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Clone: clone,
        Cluster: cluster,
        CollapseSnapshots: collapseSnapshots,
        StorageDomain: storageDomain,
        Vm: vm,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'import', nil, headers, query, wait)
}

//
// This method supports the following parameters:
// `AllowPartialImport`:: Indicates whether a virtual machine is allowed to be registered with only some of its disks.
// If this flag is `true`, the engine will not fail in the validation process if an image is not found, but
// instead it will allow the virtual machine to be registered without the missing disks. This is mainly used
// during registration of a virtual machine when some of the storage domains are not available. The default
// value is `false`.
// `VnicProfileMappings`:: Mapping rules for virtual NIC profiles that will be applied during the import process.
// `ReassignBadMacs`:: Indicates if the problematic MAC addresses should be re-assigned during the import process by the engine.
// A MAC address would be considered as a problematic one if one of the following is true:
// - It conflicts with a MAC address that is already allocated to a virtual machine in the target environment.
// - It's out of the range of the target MAC address pool.
// `Async`:: Indicates if the registration should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainVmService) Register (
    allowPartialImport bool,
    async bool,
    clone bool,
    cluster Cluster,
    reassignBadMacs bool,
    vm Vm,
    vnicProfileMappings []*VnicProfileMapping,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        AllowPartialImport: allowPartialImport,
        Async: async,
        Clone: clone,
        Cluster: cluster,
        ReassignBadMacs: reassignBadMacs,
        Vm: vm,
        VnicProfileMappings: vnicProfileMappings,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'register', nil, headers, query, wait)
}

//
// Deletes a virtual machine from an export storage domain.
// For example, to delete the virtual machine `456` from the storage domain `123`, send a request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/storagedomains/123/vms/456
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainVmService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageDomainVmService) Service(path string) *StorageDomainVmService {
    if path == nil {
        return op
    }
    if path == "diskattachments" {
        return &DiskAttachmentsService{}
    }
    if path.startswith('diskattachments/') {
        return DiskAttachmentsService().service(path[16:])
    }
    if path == "disks" {
        return &DisksService{}
    }
    if path.startswith('disks/') {
        return DisksService().service(path[6:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *StorageDomainVmService) String() string {
    return fmt.Sprintf("StorageDomainVmService:%s", op.Path)
}


//
// Returns the details of the disks attached to a virtual machine in the export domain.
//
type StorageDomainVmDiskAttachmentService struct {
    Service

}

func NewStorageDomainVmDiskAttachmentService(connection *Connection, path string) *StorageDomainVmDiskAttachmentService {
    var result StorageDomainVmDiskAttachmentService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Returns the details of the attachment with all its properties and a link to the disk.
//
func (op *StorageDomainVmDiskAttachmentService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageDomainVmDiskAttachmentService) Service(path string) *StorageDomainVmDiskAttachmentService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *StorageDomainVmDiskAttachmentService) String() string {
    return fmt.Sprintf("StorageDomainVmDiskAttachmentService:%s", op.Path)
}


//
// Returns the details of a disk attached to a virtual machine in the export domain.
//
type StorageDomainVmDiskAttachmentsService struct {
    Service

    AttachmentService  *AttachmentService
}

func NewStorageDomainVmDiskAttachmentsService(connection *Connection, path string) *StorageDomainVmDiskAttachmentsService {
    var result StorageDomainVmDiskAttachmentsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// List the disks that are attached to the virtual machine.
//
func (op *StorageDomainVmDiskAttachmentsService) List (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageDomainVmDiskAttachmentsService) Service(path string) *StorageDomainVmDiskAttachmentsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Attachment_service(path)
    return self.Attachment_service(path[:index]).Service(path[index + 1:])
}

func (op *StorageDomainVmDiskAttachmentsService) String() string {
    return fmt.Sprintf("StorageDomainVmDiskAttachmentsService:%s", op.Path)
}


//
// Lists the virtual machines of an export storage domain.
// For example, to retrieve the virtual machines that are available in the storage domain with identifier `123` send the
// following request:
// [source]
// ----
// GET /ovirt-engine/api/storagedomains/123/vms
// ----
// This will return the following response body:
// [source,xml]
// ----
// <vms>
//   <vm id="456" href="/api/storagedomains/123/vms/456">
//     <name>vm1</name>
//     ...
//     <storage_domain id="123" href="/api/storagedomains/123"/>
//     <actions>
//       <link rel="import" href="/api/storagedomains/123/vms/456/import"/>
//     </actions>
//   </vm>
// </vms>
// ----
// Virtual machines and templates in these collections have a similar representation to their counterparts in the
// top-level <<types/vm, Vm>> and <<types/template, Template>> collections, except they also contain a
// <<types/storage_domain, StorageDomain>> reference and an <<services/storage_domain_vm/methods/import, import>>
// action.
//
type StorageDomainVmsService struct {
    Service

    VmService  *VmService
}

func NewStorageDomainVmsService(connection *Connection, path string) *StorageDomainVmsService {
    var result StorageDomainVmsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of virtual machines to return. If not specified all the virtual machines are
// returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainVmsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageDomainVmsService) Service(path string) *StorageDomainVmsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Vm_service(path)
    return self.Vm_service(path[:index]).Service(path[index + 1:])
}

func (op *StorageDomainVmsService) String() string {
    return fmt.Sprintf("StorageDomainVmsService:%s", op.Path)
}


//
//
type StorageDomainsService struct {
    Service

    StorageDomainService  *StorageDomainService
}

func NewStorageDomainsService(connection *Connection, path string) *StorageDomainsService {
    var result StorageDomainsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Adds a new storage domain.
// Creation of a new <<types/storage_domain,StorageDomain>> requires the `name`, `type`, `host` and `storage`
// attributes. Identify the `host` attribute with the `id` or `name` attributes. In oVirt 3.6 and later you can
// enable the wipe after delete option by default on the storage domain. To configure this, specify
// `wipe_after_delete` in the POST request. This option can be edited after the domain is created, but doing so will
// not change the wipe after delete property of disks that already exist.
// To add a new storage domain with specified `name`, `type`, `storage.type`, `storage.address` and `storage.path`
// and by using a host with an id `123`, send a request as follows:
// [source]
// ----
// POST /ovirt-engine/api/storagedomains
// ----
// With a request body as follows:
// [source,xml]
// ----
// <storage_domain>
//   <name>mydata</name>
//   <type>data</type>
//   <storage>
//     <type>nfs</type>
//     <address>mynfs.example.com</address>
//     <path>/exports/mydata</path>
//   </storage>
//   <host>
//     <name>myhost</name>
//   </host>
// </storage_domain>
// ----
// To create a new NFS ISO storage domain send a request like this:
// [source,xml]
// ----
// <storage_domain>
//   <name>myisos</name>
//   <type>iso</type>
//   <storage>
//     <type>nfs</type>
//     <address>mynfs.example.com</address>
//     <path>/export/myisos</path>
//   </storage>
//   <host>
//     <name>myhost</name>
//   </host>
// </storage_domain>
// ----
// To create a new iSCSI storage domain send a request like this:
// [source,xml]
// ----
// <storage_domain>
//   <name>myiscsi</name>
//   <type>data</type>
//   <storage>
//     <type>iscsi</type>
//     <logical_units>
//       <logical_unit id="3600144f09dbd050000004eedbd340001"/>
//       <logical_unit id="3600144f09dbd050000004eedbd340002"/>
//     </logical_units>
//   </storage>
//   <host>
//     <name>myhost</name>
//   </host>
// </storage_domain>
// ----
//
func (op *StorageDomainsService) Add (
    storageDomain StorageDomain,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(storageDomain, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of storage domains to return. If not specified all the storage domains are returned.
// `Search`:: A query string used to restrict the returned storage domains.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageDomainsService) List (
    caseSensitive bool,
    filter bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if filter != nil {
        query["filter"] = filter
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageDomainsService) Service(path string) *StorageDomainsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.StorageDomain_service(path)
    return self.StorageDomain_service(path[:index]).Service(path[index + 1:])
}

func (op *StorageDomainsService) String() string {
    return fmt.Sprintf("StorageDomainsService:%s", op.Path)
}


//
//
type StorageServerConnectionService struct {
    Service

}

func NewStorageServerConnectionService(connection *Connection, path string) *StorageServerConnectionService {
    var result StorageServerConnectionService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *StorageServerConnectionService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes a storage connection.
// A storage connection can only be deleted if neither storage domain nor LUN disks reference it. The host name or
// id is optional; providing it disconnects (unmounts) the connection from that host.
// This method supports the following parameters:
// `Host`:: The name or identifier of the host from which the connection would be unmounted (disconnected). If not
// provided, no host will be disconnected.
// For example, to use the host with identifier `456` to delete the storage connection with identifier `123`
// send a request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/storageconnections/123?host=456
// ----
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageServerConnectionService) Remove (
    host string,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if host != nil {
        query["host"] = host
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates the storage connection.
// For example, to change the address of the storage server send a request like this:
// [source,xml]
// ----
// PUT /ovirt-engine/api/storageconnections/123
// ----
// With a request body like this:
// [source,xml]
// ----
// <storage_connection>
//   <address>mynewnfs.example.com</address>
//   <host>
//     <name>myhost</name>
//   </host>
// </storage_connection>
// ----
//
func (op *StorageServerConnectionService) Update (
    connection StorageConnection,
    async bool,
    force bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }
    if force != nil {
        query["force"] = force
    }

    // Send the request
    return op.internalUpdate(connection, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageServerConnectionService) Service(path string) *StorageServerConnectionService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *StorageServerConnectionService) String() string {
    return fmt.Sprintf("StorageServerConnectionService:%s", op.Path)
}


//
//
type StorageServerConnectionExtensionService struct {
    Service

}

func NewStorageServerConnectionExtensionService(connection *Connection, path string) *StorageServerConnectionExtensionService {
    var result StorageServerConnectionExtensionService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *StorageServerConnectionExtensionService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageServerConnectionExtensionService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Update a storage server connection extension for the given host.
// To update the storage connection `456` of host `123` send a request like this:
// [source]
// ----
// PUT /ovirt-engine/api/hosts/123/storageconnectionextensions/456
// ----
// With a request body like this:
// [source,xml]
// ----
// <storage_connection_extension>
//   <target>iqn.2016-01.com.example:mytarget</target>
//   <username>myuser</username>
//   <password>mypassword</password>
// </storage_connection_extension>
// ----
//
func (op *StorageServerConnectionExtensionService) Update (
    extension StorageConnectionExtension,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(extension, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageServerConnectionExtensionService) Service(path string) *StorageServerConnectionExtensionService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *StorageServerConnectionExtensionService) String() string {
    return fmt.Sprintf("StorageServerConnectionExtensionService:%s", op.Path)
}


//
//
type StorageServerConnectionExtensionsService struct {
    Service

    StorageConnectionExtensionService  *StorageConnectionExtensionService
}

func NewStorageServerConnectionExtensionsService(connection *Connection, path string) *StorageServerConnectionExtensionsService {
    var result StorageServerConnectionExtensionsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a new storage server connection extension for the given host.
// The extension lets the user define credentials for an iSCSI target for a specific host. For example to use
// `myuser` and `mypassword` as the credentials when connecting to the iSCSI target from host `123` send a request
// like this:
// [source]
// ----
// POST /ovirt-engine/api/hosts/123/storageconnectionextensions
// ----
// With a request body like this:
// [source,xml]
// ----
// <storage_connection_extension>
//   <target>iqn.2016-01.com.example:mytarget</target>
//   <username>myuser</username>
//   <password>mypassword</password>
// </storage_connection_extension>
// ----
//
func (op *StorageServerConnectionExtensionsService) Add (
    extension StorageConnectionExtension,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(extension, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of extensions to return. If not specified all the extensions are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageServerConnectionExtensionsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageServerConnectionExtensionsService) Service(path string) *StorageServerConnectionExtensionsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.StorageConnectionExtension_service(path)
    return self.StorageConnectionExtension_service(path[:index]).Service(path[index + 1:])
}

func (op *StorageServerConnectionExtensionsService) String() string {
    return fmt.Sprintf("StorageServerConnectionExtensionsService:%s", op.Path)
}


//
//
type StorageServerConnectionsService struct {
    Service

    StorageConnectionService  *StorageConnectionService
}

func NewStorageServerConnectionsService(connection *Connection, path string) *StorageServerConnectionsService {
    var result StorageServerConnectionsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a new storage connection.
// For example, to create a new storage connection for the NFS server `mynfs.example.com` and NFS share
// `/export/mydata` send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/storageconnections
// ----
// With a request body like this:
// [source,xml]
// ----
// <storage_connection>
//   <type>nfs</type>
//   <address>mynfs.example.com</address>
//   <path>/export/mydata</path>
//   <host>
//     <name>myhost</name>
//   </host>
// </storage_connection>
// ----
//
func (op *StorageServerConnectionsService) Add (
    connection StorageConnection,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(connection, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of connections to return. If not specified all the connections are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *StorageServerConnectionsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *StorageServerConnectionsService) Service(path string) *StorageServerConnectionsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.StorageConnection_service(path)
    return self.StorageConnection_service(path[:index]).Service(path[index + 1:])
}

func (op *StorageServerConnectionsService) String() string {
    return fmt.Sprintf("StorageServerConnectionsService:%s", op.Path)
}


//
//
type SystemService struct {
    Service

    AffinityLabelsService  *AffinityLabelsService
    BookmarksService  *BookmarksService
    ClusterLevelsService  *ClusterLevelsService
    ClustersService  *ClustersService
    CpuProfilesService  *CpuProfilesService
    DataCentersService  *DataCentersService
    DiskProfilesService  *DiskProfilesService
    DisksService  *DisksService
    DomainsService  *DomainsService
    EventsService  *EventsService
    ExternalHostProvidersService  *ExternalHostProvidersService
    ExternalVmImportsService  *ExternalVmImportsService
    GroupsService  *GroupsService
    HostsService  *HostsService
    IconsService  *IconsService
    ImageTransfersService  *ImageTransfersService
    InstanceTypesService  *InstanceTypesService
    JobsService  *JobsService
    KatelloErrataService  *KatelloErrataService
    MacPoolsService  *MacPoolsService
    NetworkFiltersService  *NetworkFiltersService
    NetworksService  *NetworksService
    OpenstackImageProvidersService  *OpenstackImageProvidersService
    OpenstackNetworkProvidersService  *OpenstackNetworkProvidersService
    OpenstackVolumeProvidersService  *OpenstackVolumeProvidersService
    OperatingSystemsService  *OperatingSystemsService
    PermissionsService  *PermissionsService
    RolesService  *RolesService
    SchedulingPoliciesService  *SchedulingPoliciesService
    SchedulingPolicyUnitsService  *SchedulingPolicyUnitsService
    StorageConnectionsService  *StorageConnectionsService
    StorageDomainsService  *StorageDomainsService
    TagsService  *TagsService
    TemplatesService  *TemplatesService
    UsersService  *UsersService
    VmPoolsService  *VmPoolsService
    VmsService  *VmsService
    VnicProfilesService  *VnicProfilesService
}

func NewSystemService(connection *Connection, path string) *SystemService {
    var result SystemService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Returns basic information describing the API, like the product name, the version number and a summary of the
// number of relevant objects.
// [source]
// ----
// GET /ovirt-engine/api
// ----
// We get following response:
// [source,xml]
// ----
// <api>
//   <link rel="capabilities" href="/api/capabilities"/>
//   <link rel="clusters" href="/api/clusters"/>
//   <link rel="clusters/search" href="/api/clusters?search={query}"/>
//   <link rel="datacenters" href="/api/datacenters"/>
//   <link rel="datacenters/search" href="/api/datacenters?search={query}"/>
//   <link rel="events" href="/api/events"/>
//   <link rel="events/search" href="/api/events?search={query}"/>
//   <link rel="hosts" href="/api/hosts"/>
//   <link rel="hosts/search" href="/api/hosts?search={query}"/>
//   <link rel="networks" href="/api/networks"/>
//   <link rel="roles" href="/api/roles"/>
//   <link rel="storagedomains" href="/api/storagedomains"/>
//   <link rel="storagedomains/search" href="/api/storagedomains?search={query}"/>
//   <link rel="tags" href="/api/tags"/>
//   <link rel="templates" href="/api/templates"/>
//   <link rel="templates/search" href="/api/templates?search={query}"/>
//   <link rel="users" href="/api/users"/>
//   <link rel="groups" href="/api/groups"/>
//   <link rel="domains" href="/api/domains"/>
//   <link rel="vmpools" href="/api/vmpools"/>
//   <link rel="vmpools/search" href="/api/vmpools?search={query}"/>
//   <link rel="vms" href="/api/vms"/>
//   <link rel="vms/search" href="/api/vms?search={query}"/>
//   <product_info>
//     <name>oVirt Engine</name>
//     <vendor>ovirt.org</vendor>
//     <version>
//       <build>4</build>
//       <full_version>4.0.4</full_version>
//       <major>4</major>
//       <minor>0</minor>
//       <revision>0</revision>
//     </version>
//   </product_info>
//   <special_objects>
//     <blank_template href="/ovirt-engine/api/templates/00000000-0000-0000-0000-000000000000" id="00000000-0000-0000-0000-000000000000"/>
//     <root_tag href="/ovirt-engine/api/tags/00000000-0000-0000-0000-000000000000" id="00000000-0000-0000-0000-000000000000"/>
//   </special_objects>
//   <summary>
//     <hosts>
//       <active>0</active>
//       <total>0</total>
//     </hosts>
//     <storage_domains>
//       <active>0</active>
//       <total>1</total>
//     </storage_domains>
//     <users>
//       <active>1</active>
//       <total>1</total>
//     </users>
//     <vms>
//       <active>0</active>
//       <total>0</total>
//     </vms>
//   </summary>
//   <time>2016-09-14T12:00:48.132+02:00</time>
// </api>
// ----
// The entry point provides a user with links to the collections in a
// virtualization environment. The `rel` attribute of each collection link
// provides a reference point for each link.
// The entry point also contains other data such as `product_info`,
// `special_objects` and `summary`.
//
func (op *SystemService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the reload should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SystemService) ReloadConfigurations (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'reloadconfigurations', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SystemService) Service(path string) *SystemService {
    if path == nil {
        return op
    }
    if path == "affinitylabels" {
        return &AffinityLabelsService{}
    }
    if path.startswith('affinitylabels/') {
        return AffinityLabelsService().service(path[15:])
    }
    if path == "bookmarks" {
        return &BookmarksService{}
    }
    if path.startswith('bookmarks/') {
        return BookmarksService().service(path[10:])
    }
    if path == "clusterlevels" {
        return &ClusterLevelsService{}
    }
    if path.startswith('clusterlevels/') {
        return ClusterLevelsService().service(path[14:])
    }
    if path == "clusters" {
        return &ClustersService{}
    }
    if path.startswith('clusters/') {
        return ClustersService().service(path[9:])
    }
    if path == "cpuprofiles" {
        return &CpuProfilesService{}
    }
    if path.startswith('cpuprofiles/') {
        return CpuProfilesService().service(path[12:])
    }
    if path == "datacenters" {
        return &DataCentersService{}
    }
    if path.startswith('datacenters/') {
        return DataCentersService().service(path[12:])
    }
    if path == "diskprofiles" {
        return &DiskProfilesService{}
    }
    if path.startswith('diskprofiles/') {
        return DiskProfilesService().service(path[13:])
    }
    if path == "disks" {
        return &DisksService{}
    }
    if path.startswith('disks/') {
        return DisksService().service(path[6:])
    }
    if path == "domains" {
        return &DomainsService{}
    }
    if path.startswith('domains/') {
        return DomainsService().service(path[8:])
    }
    if path == "events" {
        return &EventsService{}
    }
    if path.startswith('events/') {
        return EventsService().service(path[7:])
    }
    if path == "externalhostproviders" {
        return &ExternalHostProvidersService{}
    }
    if path.startswith('externalhostproviders/') {
        return ExternalHostProvidersService().service(path[22:])
    }
    if path == "externalvmimports" {
        return &ExternalVmImportsService{}
    }
    if path.startswith('externalvmimports/') {
        return ExternalVmImportsService().service(path[18:])
    }
    if path == "groups" {
        return &GroupsService{}
    }
    if path.startswith('groups/') {
        return GroupsService().service(path[7:])
    }
    if path == "hosts" {
        return &HostsService{}
    }
    if path.startswith('hosts/') {
        return HostsService().service(path[6:])
    }
    if path == "icons" {
        return &IconsService{}
    }
    if path.startswith('icons/') {
        return IconsService().service(path[6:])
    }
    if path == "imagetransfers" {
        return &ImageTransfersService{}
    }
    if path.startswith('imagetransfers/') {
        return ImageTransfersService().service(path[15:])
    }
    if path == "instancetypes" {
        return &InstanceTypesService{}
    }
    if path.startswith('instancetypes/') {
        return InstanceTypesService().service(path[14:])
    }
    if path == "jobs" {
        return &JobsService{}
    }
    if path.startswith('jobs/') {
        return JobsService().service(path[5:])
    }
    if path == "katelloerrata" {
        return &KatelloErrataService{}
    }
    if path.startswith('katelloerrata/') {
        return KatelloErrataService().service(path[14:])
    }
    if path == "macpools" {
        return &MacPoolsService{}
    }
    if path.startswith('macpools/') {
        return MacPoolsService().service(path[9:])
    }
    if path == "networkfilters" {
        return &NetworkFiltersService{}
    }
    if path.startswith('networkfilters/') {
        return NetworkFiltersService().service(path[15:])
    }
    if path == "networks" {
        return &NetworksService{}
    }
    if path.startswith('networks/') {
        return NetworksService().service(path[9:])
    }
    if path == "openstackimageproviders" {
        return &OpenstackImageProvidersService{}
    }
    if path.startswith('openstackimageproviders/') {
        return OpenstackImageProvidersService().service(path[24:])
    }
    if path == "openstacknetworkproviders" {
        return &OpenstackNetworkProvidersService{}
    }
    if path.startswith('openstacknetworkproviders/') {
        return OpenstackNetworkProvidersService().service(path[26:])
    }
    if path == "openstackvolumeproviders" {
        return &OpenstackVolumeProvidersService{}
    }
    if path.startswith('openstackvolumeproviders/') {
        return OpenstackVolumeProvidersService().service(path[25:])
    }
    if path == "operatingsystems" {
        return &OperatingSystemsService{}
    }
    if path.startswith('operatingsystems/') {
        return OperatingSystemsService().service(path[17:])
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    if path == "roles" {
        return &RolesService{}
    }
    if path.startswith('roles/') {
        return RolesService().service(path[6:])
    }
    if path == "schedulingpolicies" {
        return &SchedulingPoliciesService{}
    }
    if path.startswith('schedulingpolicies/') {
        return SchedulingPoliciesService().service(path[19:])
    }
    if path == "schedulingpolicyunits" {
        return &SchedulingPolicyUnitsService{}
    }
    if path.startswith('schedulingpolicyunits/') {
        return SchedulingPolicyUnitsService().service(path[22:])
    }
    if path == "storageconnections" {
        return &StorageConnectionsService{}
    }
    if path.startswith('storageconnections/') {
        return StorageConnectionsService().service(path[19:])
    }
    if path == "storagedomains" {
        return &StorageDomainsService{}
    }
    if path.startswith('storagedomains/') {
        return StorageDomainsService().service(path[15:])
    }
    if path == "tags" {
        return &TagsService{}
    }
    if path.startswith('tags/') {
        return TagsService().service(path[5:])
    }
    if path == "templates" {
        return &TemplatesService{}
    }
    if path.startswith('templates/') {
        return TemplatesService().service(path[10:])
    }
    if path == "users" {
        return &UsersService{}
    }
    if path.startswith('users/') {
        return UsersService().service(path[6:])
    }
    if path == "vmpools" {
        return &VmPoolsService{}
    }
    if path.startswith('vmpools/') {
        return VmPoolsService().service(path[8:])
    }
    if path == "vms" {
        return &VmsService{}
    }
    if path.startswith('vms/') {
        return VmsService().service(path[4:])
    }
    if path == "vnicprofiles" {
        return &VnicProfilesService{}
    }
    if path.startswith('vnicprofiles/') {
        return VnicProfilesService().service(path[13:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *SystemService) String() string {
    return fmt.Sprintf("SystemService:%s", op.Path)
}


//
// This service doesn't add any new methods, it is just a placeholder for the annotation that specifies the path of the
// resource that manages the permissions assigned to the system object.
//
type SystemPermissionsService struct {
    AssignedPermissionsService

    PermissionService  *PermissionService
}

func NewSystemPermissionsService(connection *Connection, path string) *SystemPermissionsService {
    var result SystemPermissionsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Assign a new permission to a user or group for specific entity.
// For example, to assign the `UserVmManager` role to the virtual machine with id `123` to the user with id `456`
// send a request like this:
// ....
// POST /ovirt-engine/api/vms/123/permissions
// ....
// With a request body like this:
// [source,xml]
// ----
// <permission>
//   <role>
//     <name>UserVmManager</name>
//   </role>
//   <user id="456"/>
// </permission>
// ----
// To assign the `SuperUser` role to the system to the user with id `456` send a request like this:
// ....
// POST /ovirt-engine/api/permissions
// ....
// With a request body like this:
// [source,xml]
// ----
// <permission>
//   <role>
//     <name>SuperUser</name>
//   </role>
//   <user id="456"/>
// </permission>
// ----
// If you want to assign permission to the group instead of the user please replace the `user` element with the
// `group` element with proper `id` of the group. For example to assign the `UserRole` role to the cluster with
// id `123` to the group with id `789` send a request like this:
// ....
// POST /ovirt-engine/api/clusters/123/permissions
// ....
// With a request body like this:
// [source,xml]
// ----
// <permission>
//   <role>
//     <name>UserRole</name>
//   </role>
//   <group id="789"/>
// </permission>
// ----
// This method supports the following parameters:
// `Permission`:: The permission.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *SystemPermissionsService) Add (
    permission Permission,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(permission, headers, query, wait)
}

//
// List all the permissions of the specific entity.
// For example to list all the permissions of the cluster with id `123` send a request like this:
// ....
// GET /ovirt-engine/api/clusters/123/permissions
// ....
// [source,xml]
// ----
// <permissions>
//   <permission id="456">
//     <cluster id="123"/>
//     <role id="789"/>
//     <user id="451"/>
//   </permission>
//   <permission id="654">
//     <cluster id="123"/>
//     <role id="789"/>
//     <group id="127"/>
//   </permission>
// </permissions>
// ----
//
func (op *SystemPermissionsService) List (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *SystemPermissionsService) Service(path string) *SystemPermissionsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Permission_service(path)
    return self.Permission_service(path[:index]).Service(path[index + 1:])
}

func (op *SystemPermissionsService) String() string {
    return fmt.Sprintf("SystemPermissionsService:%s", op.Path)
}


//
// A service to manage a specific tag in the system.
//
type TagService struct {
    Service

}

func NewTagService(connection *Connection, path string) *TagService {
    var result TagService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Gets the information about the tag.
// For example to retrieve the information about the tag with the id `123` send a request like this:
// ....
// GET /ovirt-engine/api/tags/123
// ....
// [source,xml]
// ----
// <tag href="/ovirt-engine/api/tags/123" id="123">
//   <name>root</name>
//   <description>root</description>
// </tag>
// ----
//
func (op *TagService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes the tag from the system.
// For example to remove the tag with id `123` send a request like this:
// ....
// DELETE /ovirt-engine/api/tags/123
// ....
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TagService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates the tag entity.
// For example to update parent tag to tag with id `456` of the tag with id `123` send a request like this:
// ....
// PUT /ovirt-engine/api/tags/123
// ....
// With request body like:
// [source,xml]
// ----
// <tag>
//   <parent id="456"/>
// </tag>
// ----
// You may also specify a tag name instead of id. For example to update parent tag to tag with name `mytag`
// of the tag with id `123` send a request like this:
// [source,xml]
// ----
// <tag>
//   <parent>
//     <name>mytag</name>
//   </parent>
// </tag>
// ----
// This method supports the following parameters:
// `Tag`:: The updated tag.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TagService) Update (
    tag Tag,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(tag, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TagService) Service(path string) *TagService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *TagService) String() string {
    return fmt.Sprintf("TagService:%s", op.Path)
}


//
// Represents a service to manage collection of the tags in the system.
//
type TagsService struct {
    Service

    TagService  *TagService
}

func NewTagsService(connection *Connection, path string) *TagsService {
    var result TagsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add a new tag to the system.
// For example, to add new tag with name `mytag` to the system send a request like this:
// ....
// POST /ovirt-engine/api/tags
// ....
// With a request body like this:
// [source,xml]
// ----
// <tag>
//   <name>mytag</name>
// </tag>
// ----
// NOTE: The root tag is a special pseudo-tag assumed as the default parent tag if no parent tag is specified.
// The root tag cannot be deleted nor assigned a parent tag.
// To create new tag with specific parent tag send a request body like this:
// [source,xml]
// ----
// <tag>
//   <name>mytag</name>
//   <parent>
//     <name>myparenttag</name>
//   </parent>
// </tag>
// ----
// This method supports the following parameters:
// `Tag`:: The added tag.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TagsService) Add (
    tag Tag,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(tag, headers, query, wait)
}

//
// List the tags in the system.
// For example to list the full hierarchy of the tags in the system send a request like this:
// ....
// GET /ovirt-engine/api/tags
// ....
// [source,xml]
// ----
// <tags>
//   <tag href="/ovirt-engine/api/tags/222" id="222">
//     <name>root2</name>
//     <description>root2</description>
//     <parent href="/ovirt-engine/api/tags/111" id="111"/>
//   </tag>
//   <tag href="/ovirt-engine/api/tags/333" id="333">
//     <name>root3</name>
//     <description>root3</description>
//     <parent href="/ovirt-engine/api/tags/222" id="222"/>
//   </tag>
//   <tag href="/ovirt-engine/api/tags/111" id="111">
//     <name>root</name>
//     <description>root</description>
//   </tag>
// </tags>
// ----
// In the previous XML output you can see the following hierarchy of the tags:
// ....
// root:        (id: 111)
//   - root2    (id: 222)
//     - root3  (id: 333)
// ....
// This method supports the following parameters:
// `Max`:: Sets the maximum number of tags to return. If not specified all the tags are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TagsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TagsService) Service(path string) *TagsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Tag_service(path)
    return self.Tag_service(path[:index]).Service(path[index + 1:])
}

func (op *TagsService) String() string {
    return fmt.Sprintf("TagsService:%s", op.Path)
}


//
// Manages the virtual machine template and template versions.
//
type TemplateService struct {
    Service

    CdromsService  *CdromsService
    DiskAttachmentsService  *DiskAttachmentsService
    GraphicsConsolesService  *GraphicsConsolesService
    NicsService  *NicsService
    PermissionsService  *PermissionsService
    TagsService  *TagsService
    WatchdogsService  *WatchdogsService
}

func NewTemplateService(connection *Connection, path string) *TemplateService {
    var result TemplateService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Exports a template to the data center export domain.
// For example, the operation can be facilitated using the following request:
// [source]
// ----
// POST /ovirt-engine/api/templates/123/export
// ----
// With a request body like this:
// [source,xml]
// ----
// <action>
//   <storage_domain id="456"/>
//   <exclusive>true<exclusive/>
// </action>
// ----
// This method supports the following parameters:
// `Exclusive`:: Indicates if the existing templates with the same name should be overwritten.
// The export action reports a failed action if a template of the same name exists in the destination domain.
// Set this parameter to `true` to change this behavior and overwrite any existing template.
// `StorageDomain`:: Specifies the destination export storage domain.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplateService) Export (
    exclusive bool,
    storageDomain StorageDomain,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Exclusive: exclusive,
        StorageDomain: storageDomain,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'export', nil, headers, query, wait)
}

//
// Returns the information about this template or template version.
// This method supports the following parameters:
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplateService) Get (
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes a virtual machine template.
// [source]
// ----
// DELETE /ovirt-engine/api/templates/123
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplateService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Seal the template.
// Sealing erases all host-specific configuration from the filesystem:
// SSH keys, UDEV rules, MAC addresses, system ID, hostname etc.,
// thus making easy to use the template to create multiple virtual
// machines without manual intervention.
// Currently sealing is supported only for Linux OS.
//
func (op *TemplateService) Seal (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
    }

    // Send the request and wait for the response:
    return internalAction(action, 'seal', nil, headers, query, wait)
}

//
// Updates the template.
// The `name`, `description`, `type`, `memory`, `cpu`, `topology`, `os`, `high_availability`, `display`,
// `stateless`, `usb` and `timezone` elements can be updated after a template has been created.
// For example, to update a template to so that it has 1 GiB of memory send a request like this:
// [source]
// ----
// PUT /ovirt-engine/api/templates/123
// ----
// With the following request body:
// [source,xml]
// ----
// <template>
//   <memory>1073741824</memory>
// </template>
// ----
// The `version_name` name attribute is the only one that can be updated within the `version` attribute used for
// template versions:
// [source,xml]
// ----
// <template>
//   <version>
//     <version_name>mytemplate_2</version_name>
//   </version>
// </template>
// ----
//
func (op *TemplateService) Update (
    template Template,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(template, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TemplateService) Service(path string) *TemplateService {
    if path == nil {
        return op
    }
    if path == "cdroms" {
        return &CdromsService{}
    }
    if path.startswith('cdroms/') {
        return CdromsService().service(path[7:])
    }
    if path == "diskattachments" {
        return &DiskAttachmentsService{}
    }
    if path.startswith('diskattachments/') {
        return DiskAttachmentsService().service(path[16:])
    }
    if path == "graphicsconsoles" {
        return &GraphicsConsolesService{}
    }
    if path.startswith('graphicsconsoles/') {
        return GraphicsConsolesService().service(path[17:])
    }
    if path == "nics" {
        return &NicsService{}
    }
    if path.startswith('nics/') {
        return NicsService().service(path[5:])
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    if path == "tags" {
        return &TagsService{}
    }
    if path.startswith('tags/') {
        return TagsService().service(path[5:])
    }
    if path == "watchdogs" {
        return &WatchdogsService{}
    }
    if path.startswith('watchdogs/') {
        return WatchdogsService().service(path[10:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *TemplateService) String() string {
    return fmt.Sprintf("TemplateService:%s", op.Path)
}


//
// A service managing a CD-ROM device on templates.
//
type TemplateCdromService struct {
    Service

}

func NewTemplateCdromService(connection *Connection, path string) *TemplateCdromService {
    var result TemplateCdromService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Returns the information about this CD-ROM device.
// For example, to get information about the CD-ROM device of template `123` send a request like:
// [source]
// ----
// GET /ovirt-engine/api/templates/123/cdroms/
// ----
//
func (op *TemplateCdromService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TemplateCdromService) Service(path string) *TemplateCdromService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *TemplateCdromService) String() string {
    return fmt.Sprintf("TemplateCdromService:%s", op.Path)
}


//
// Lists the CD-ROM devices of a template.
//
type TemplateCdromsService struct {
    Service

    CdromService  *CdromService
}

func NewTemplateCdromsService(connection *Connection, path string) *TemplateCdromsService {
    var result TemplateCdromsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of CD-ROMs to return. If not specified all the CD-ROMs are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplateCdromsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TemplateCdromsService) Service(path string) *TemplateCdromsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Cdrom_service(path)
    return self.Cdrom_service(path[:index]).Service(path[index + 1:])
}

func (op *TemplateCdromsService) String() string {
    return fmt.Sprintf("TemplateCdromsService:%s", op.Path)
}


//
//
type TemplateDiskService struct {
    Service

}

func NewTemplateDiskService(connection *Connection, path string) *TemplateDiskService {
    var result TemplateDiskService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the copy should be performed asynchronously.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplateDiskService) Copy (
    async bool,
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Filter: filter,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'copy', nil, headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the export should be performed asynchronously.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplateDiskService) Export (
    async bool,
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Filter: filter,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'export', nil, headers, query, wait)
}

//
//
func (op *TemplateDiskService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplateDiskService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TemplateDiskService) Service(path string) *TemplateDiskService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *TemplateDiskService) String() string {
    return fmt.Sprintf("TemplateDiskService:%s", op.Path)
}


//
// This service manages the attachment of a disk to a template.
//
type TemplateDiskAttachmentService struct {
    Service

}

func NewTemplateDiskAttachmentService(connection *Connection, path string) *TemplateDiskAttachmentService {
    var result TemplateDiskAttachmentService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Returns the details of the attachment.
//
func (op *TemplateDiskAttachmentService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes the disk from the template. The disk will only be removed if there are other existing copies of the
// disk on other storage domains.
// A storage domain has to be specified to determine which of the copies should be removed (template disks can
// have copies on multiple storage domains).
// [source]
// ----
// DELETE /ovirt-engine/api/templates/{template:id}/diskattachments/{attachment:id}?storage_domain=072fbaa1-08f3-4a40-9f34-a5ca22dd1d74
// ----
// This method supports the following parameters:
// `StorageDomain`:: Specifies the identifier of the storage domain the image to be removed resides on.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplateDiskAttachmentService) Remove (
    storageDomain string,
    force bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if storageDomain != nil {
        query["storage_domain"] = storageDomain
    }
    if force != nil {
        query["force"] = force
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TemplateDiskAttachmentService) Service(path string) *TemplateDiskAttachmentService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *TemplateDiskAttachmentService) String() string {
    return fmt.Sprintf("TemplateDiskAttachmentService:%s", op.Path)
}


//
// This service manages the set of disks attached to a template. Each attached disk is represented by a
// <<types/disk_attachment,DiskAttachment>>.
//
type TemplateDiskAttachmentsService struct {
    Service

    AttachmentService  *AttachmentService
}

func NewTemplateDiskAttachmentsService(connection *Connection, path string) *TemplateDiskAttachmentsService {
    var result TemplateDiskAttachmentsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// List the disks that are attached to the template.
//
func (op *TemplateDiskAttachmentsService) List (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TemplateDiskAttachmentsService) Service(path string) *TemplateDiskAttachmentsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Attachment_service(path)
    return self.Attachment_service(path[:index]).Service(path[index + 1:])
}

func (op *TemplateDiskAttachmentsService) String() string {
    return fmt.Sprintf("TemplateDiskAttachmentsService:%s", op.Path)
}


//
//
type TemplateDisksService struct {
    Service

    DiskService  *DiskService
}

func NewTemplateDisksService(connection *Connection, path string) *TemplateDisksService {
    var result TemplateDisksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of disks to return. If not specified all the disks are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplateDisksService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TemplateDisksService) Service(path string) *TemplateDisksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Disk_service(path)
    return self.Disk_service(path[:index]).Service(path[index + 1:])
}

func (op *TemplateDisksService) String() string {
    return fmt.Sprintf("TemplateDisksService:%s", op.Path)
}


//
//
type TemplateGraphicsConsoleService struct {
    Service

}

func NewTemplateGraphicsConsoleService(connection *Connection, path string) *TemplateGraphicsConsoleService {
    var result TemplateGraphicsConsoleService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Gets graphics console configuration of the template.
//
func (op *TemplateGraphicsConsoleService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Remove the graphics console from the template.
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplateGraphicsConsoleService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TemplateGraphicsConsoleService) Service(path string) *TemplateGraphicsConsoleService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *TemplateGraphicsConsoleService) String() string {
    return fmt.Sprintf("TemplateGraphicsConsoleService:%s", op.Path)
}


//
//
type TemplateGraphicsConsolesService struct {
    Service

    ConsoleService  *ConsoleService
}

func NewTemplateGraphicsConsolesService(connection *Connection, path string) *TemplateGraphicsConsolesService {
    var result TemplateGraphicsConsolesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add new graphics console to the template.
//
func (op *TemplateGraphicsConsolesService) Add (
    console GraphicsConsole,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(console, headers, query, wait)
}

//
// Lists all the configured graphics consoles of the template.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of consoles to return. If not specified all the consoles are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplateGraphicsConsolesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TemplateGraphicsConsolesService) Service(path string) *TemplateGraphicsConsolesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Console_service(path)
    return self.Console_service(path[:index]).Service(path[index + 1:])
}

func (op *TemplateGraphicsConsolesService) String() string {
    return fmt.Sprintf("TemplateGraphicsConsolesService:%s", op.Path)
}


//
//
type TemplateNicService struct {
    Service

}

func NewTemplateNicService(connection *Connection, path string) *TemplateNicService {
    var result TemplateNicService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *TemplateNicService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplateNicService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
//
func (op *TemplateNicService) Update (
    nic Nic,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(nic, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TemplateNicService) Service(path string) *TemplateNicService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *TemplateNicService) String() string {
    return fmt.Sprintf("TemplateNicService:%s", op.Path)
}


//
//
type TemplateNicsService struct {
    Service

    NicService  *NicService
}

func NewTemplateNicsService(connection *Connection, path string) *TemplateNicsService {
    var result TemplateNicsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *TemplateNicsService) Add (
    nic Nic,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(nic, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of NICs to return. If not specified all the NICs are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplateNicsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TemplateNicsService) Service(path string) *TemplateNicsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Nic_service(path)
    return self.Nic_service(path[:index]).Service(path[index + 1:])
}

func (op *TemplateNicsService) String() string {
    return fmt.Sprintf("TemplateNicsService:%s", op.Path)
}


//
//
type TemplateWatchdogService struct {
    Service

}

func NewTemplateWatchdogService(connection *Connection, path string) *TemplateWatchdogService {
    var result TemplateWatchdogService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *TemplateWatchdogService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplateWatchdogService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
//
func (op *TemplateWatchdogService) Update (
    watchdog Watchdog,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(watchdog, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TemplateWatchdogService) Service(path string) *TemplateWatchdogService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *TemplateWatchdogService) String() string {
    return fmt.Sprintf("TemplateWatchdogService:%s", op.Path)
}


//
//
type TemplateWatchdogsService struct {
    Service

    WatchdogService  *WatchdogService
}

func NewTemplateWatchdogsService(connection *Connection, path string) *TemplateWatchdogsService {
    var result TemplateWatchdogsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *TemplateWatchdogsService) Add (
    watchdog Watchdog,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(watchdog, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of watchdogs to return. If not specified all the watchdogs are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplateWatchdogsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TemplateWatchdogsService) Service(path string) *TemplateWatchdogsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Watchdog_service(path)
    return self.Watchdog_service(path[:index]).Service(path[index + 1:])
}

func (op *TemplateWatchdogsService) String() string {
    return fmt.Sprintf("TemplateWatchdogsService:%s", op.Path)
}


//
// This service manages the virtual machine templates available in the system.
//
type TemplatesService struct {
    Service

    TemplateService  *TemplateService
}

func NewTemplatesService(connection *Connection, path string) *TemplatesService {
    var result TemplatesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a new template.
// This requires the `name` and `vm` elements. Identify the virtual machine with the `id` `name` attributes.
// [source]
// ----
// POST /ovirt-engine/api/templates
// ----
// With a request body like this:
// [source,xml]
// ----
// <template>
//   <name>mytemplate</name>
//   <vm id="123"/>
// </template>
// ----
// The template can be created as a sub version of an existing template.This requires the `name` and `vm` attributes
// for the new template, and the `base_template` and `version_name` attributes for the new template version. The
// `base_template` and `version_name` attributes must be specified within a `version` section enclosed in the
// `template` section. Identify the virtual machine with the `id` or `name` attributes.
// [source,xml]
// ----
// <template>
//   <name>mytemplate</name>
//   <vm id="123"/>
//   <version>
//     <base_template id="456"/>
//     <version_name>mytemplate_001</version_name>
//   </version>
// </template>
// ----
// This method supports the following parameters:
// `Template`:: The information about the template or template version.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplatesService) Add (
    template Template,
    clonePermissions bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if clonePermissions != nil {
        query["clone_permissions"] = clonePermissions
    }

    // Send the request
    return op.internalAdd(template, headers, query, wait)
}

//
// Returns the list of virtual machine templates.
// For example:
// [source]
// ----
// GET /ovirt-engine/api/templates
// ----
// Will return the list of virtual machines and virtual machine templates.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of templates to return. If not specified all the templates are returned.
// `Search`:: A query string used to restrict the returned templates.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *TemplatesService) List (
    caseSensitive bool,
    filter bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if filter != nil {
        query["filter"] = filter
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *TemplatesService) Service(path string) *TemplatesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Template_service(path)
    return self.Template_service(path[:index]).Service(path[index + 1:])
}

func (op *TemplatesService) String() string {
    return fmt.Sprintf("TemplatesService:%s", op.Path)
}


//
//
type UnmanagedNetworkService struct {
    Service

}

func NewUnmanagedNetworkService(connection *Connection, path string) *UnmanagedNetworkService {
    var result UnmanagedNetworkService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *UnmanagedNetworkService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *UnmanagedNetworkService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *UnmanagedNetworkService) Service(path string) *UnmanagedNetworkService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *UnmanagedNetworkService) String() string {
    return fmt.Sprintf("UnmanagedNetworkService:%s", op.Path)
}


//
//
type UnmanagedNetworksService struct {
    Service

    UnmanagedNetworkService  *UnmanagedNetworkService
}

func NewUnmanagedNetworksService(connection *Connection, path string) *UnmanagedNetworksService {
    var result UnmanagedNetworksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of networks to return. If not specified all the networks are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *UnmanagedNetworksService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *UnmanagedNetworksService) Service(path string) *UnmanagedNetworksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.UnmanagedNetwork_service(path)
    return self.UnmanagedNetwork_service(path[:index]).Service(path[index + 1:])
}

func (op *UnmanagedNetworksService) String() string {
    return fmt.Sprintf("UnmanagedNetworksService:%s", op.Path)
}


//
// A service to manage a user in the system.
// Use this service to either get users details or remove users.
// In order to add new users please use
// <<services/users>>.
//
type UserService struct {
    Service

    PermissionsService  *PermissionsService
    RolesService  *RolesService
    SshPublicKeysService  *SshPublicKeysService
    TagsService  *TagsService
}

func NewUserService(connection *Connection, path string) *UserService {
    var result UserService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Gets the system user information.
// Usage:
// ....
// GET /ovirt-engine/api/users/1234
// ....
// Will return the user information:
// [source,xml]
// ----
// <user href="/ovirt-engine/api/users/1234" id="1234">
//   <name>admin</name>
//   <link href="/ovirt-engine/api/users/1234/sshpublickeys" rel="sshpublickeys"/>
//   <link href="/ovirt-engine/api/users/1234/roles" rel="roles"/>
//   <link href="/ovirt-engine/api/users/1234/permissions" rel="permissions"/>
//   <link href="/ovirt-engine/api/users/1234/tags" rel="tags"/>
//   <department></department>
//   <domain_entry_id>23456</domain_entry_id>
//   <email>user1@domain.com</email>
//   <last_name>Lastname</last_name>
//   <namespace>*</namespace>
//   <principal>user1</principal>
//   <user_name>user1@domain-authz</user_name>
//   <domain href="/ovirt-engine/api/domains/45678" id="45678">
//     <name>domain-authz</name>
//   </domain>
// </user>
// ----
//
func (op *UserService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes the system user.
// Usage:
// ....
// DELETE /ovirt-engine/api/users/1234
// ....
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *UserService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *UserService) Service(path string) *UserService {
    if path == nil {
        return op
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    if path == "roles" {
        return &RolesService{}
    }
    if path.startswith('roles/') {
        return RolesService().service(path[6:])
    }
    if path == "sshpublickeys" {
        return &SshPublicKeysService{}
    }
    if path.startswith('sshpublickeys/') {
        return SshPublicKeysService().service(path[14:])
    }
    if path == "tags" {
        return &TagsService{}
    }
    if path.startswith('tags/') {
        return TagsService().service(path[5:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *UserService) String() string {
    return fmt.Sprintf("UserService:%s", op.Path)
}


//
// A service to manage the users in the system.
//
type UsersService struct {
    Service

    UserService  *UserService
}

func NewUsersService(connection *Connection, path string) *UsersService {
    var result UsersService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add user from a directory service.
// For example, to add the `myuser` user from the `myextension-authz` authorization provider send a request
// like this:
// [source]
// ----
// POST /ovirt-engine/api/users
// ----
// With a request body like this:
// [source,xml]
// ----
// <user>
//   <user_name>myuser@myextension-authz</user_name>
//   <domain>
//     <name>myextension-authz</name>
//   </domain>
// </user>
// ----
// In case you are working with Active Directory you have to pass user principal name (UPN) as `username`, followed
// by authorization provider name. Due to https://bugzilla.redhat.com/1147900[bug 1147900] you need to provide
// also `principal` parameter set to UPN of the user.
// For example, to add the user with UPN `myuser@mysubdomain.mydomain.com` from the `myextension-authz`
// authorization provider send a request body like this:
// [source,xml]
// ----
// <user>
//   <principal>myuser@mysubdomain.mydomain.com</principal>
//   <user_name>myuser@mysubdomain.mydomain.com@myextension-authz</user_name>
//   <domain>
//     <name>myextension-authz</name>
//   </domain>
// </user>
// ----
//
func (op *UsersService) Add (
    user User,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(user, headers, query, wait)
}

//
// List all the users in the system.
// Usage:
// ....
// GET /ovirt-engine/api/users
// ....
// Will return the list of users:
// [source,xml]
// ----
// <users>
//   <user href="/ovirt-engine/api/users/1234" id="1234">
//     <name>admin</name>
//     <link href="/ovirt-engine/api/users/1234/sshpublickeys" rel="sshpublickeys"/>
//     <link href="/ovirt-engine/api/users/1234/roles" rel="roles"/>
//     <link href="/ovirt-engine/api/users/1234/permissions" rel="permissions"/>
//     <link href="/ovirt-engine/api/users/1234/tags" rel="tags"/>
//     <domain_entry_id>23456</domain_entry_id>
//     <namespace>*</namespace>
//     <principal>user1</principal>
//     <user_name>user1@domain-authz</user_name>
//     <domain href="/ovirt-engine/api/domains/45678" id="45678">
//       <name>domain-authz</name>
//     </domain>
//   </user>
// </users>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of users to return. If not specified all the users are returned.
// `Search`:: A query string used to restrict the returned users.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *UsersService) List (
    caseSensitive bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *UsersService) Service(path string) *UsersService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.User_service(path)
    return self.User_service(path[:index]).Service(path[index + 1:])
}

func (op *UsersService) String() string {
    return fmt.Sprintf("UsersService:%s", op.Path)
}


//
//
type VirtualFunctionAllowedNetworkService struct {
    Service

}

func NewVirtualFunctionAllowedNetworkService(connection *Connection, path string) *VirtualFunctionAllowedNetworkService {
    var result VirtualFunctionAllowedNetworkService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *VirtualFunctionAllowedNetworkService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VirtualFunctionAllowedNetworkService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VirtualFunctionAllowedNetworkService) Service(path string) *VirtualFunctionAllowedNetworkService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *VirtualFunctionAllowedNetworkService) String() string {
    return fmt.Sprintf("VirtualFunctionAllowedNetworkService:%s", op.Path)
}


//
//
type VirtualFunctionAllowedNetworksService struct {
    Service

    NetworkService  *NetworkService
}

func NewVirtualFunctionAllowedNetworksService(connection *Connection, path string) *VirtualFunctionAllowedNetworksService {
    var result VirtualFunctionAllowedNetworksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *VirtualFunctionAllowedNetworksService) Add (
    network Network,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(network, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of networks to return. If not specified all the networks are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VirtualFunctionAllowedNetworksService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VirtualFunctionAllowedNetworksService) Service(path string) *VirtualFunctionAllowedNetworksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Network_service(path)
    return self.Network_service(path[:index]).Service(path[index + 1:])
}

func (op *VirtualFunctionAllowedNetworksService) String() string {
    return fmt.Sprintf("VirtualFunctionAllowedNetworksService:%s", op.Path)
}


//
//
type VmService struct {
    MeasurableService

    AffinityLabelsService  *AffinityLabelsService
    ApplicationsService  *ApplicationsService
    CdromsService  *CdromsService
    DiskAttachmentsService  *DiskAttachmentsService
    GraphicsConsolesService  *GraphicsConsolesService
    HostDevicesService  *HostDevicesService
    KatelloErrataService  *KatelloErrataService
    NicsService  *NicsService
    NumaNodesService  *NumaNodesService
    PermissionsService  *PermissionsService
    ReportedDevicesService  *ReportedDevicesService
    SessionsService  *SessionsService
    SnapshotsService  *SnapshotsService
    StatisticsService  *StatisticsService
    TagsService  *TagsService
    WatchdogsService  *WatchdogsService
}

func NewVmService(connection *Connection, path string) *VmService {
    var result VmService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This operation stops any migration of a virtual machine to another physical host.
// [source]
// ----
// POST /ovirt-engine/api/vms/123/cancelmigration
// ----
// The cancel migration action does not take any action specific parameters,
// so the request body should contain an empty `action`:
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the migration should cancelled asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) CancelMigration (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'cancelmigration', nil, headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the clone should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) Clone (
    async bool,
    vm Vm,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Vm: vm,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'clone', nil, headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the snapshots should be committed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) CommitSnapshot (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'commitsnapshot', nil, headers, query, wait)
}

//
// Detaches a virtual machine from a pool.
// [source]
// ----
// POST /ovirt-engine/api/vms/123/detach
// ----
// The detach action does not take any action specific parameters, so the request body should contain an
// empty `action`:
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the detach should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) Detach (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'detach', nil, headers, query, wait)
}

//
// Export a virtual machine to an export domain.
// For example to export virtual machine `123` to the export domain `myexport`, send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/vms/123/export
// ----
// With a request body like this:
// [source,xml]
// ----
// <action>
//   <storage_domain>
//     <name>myexport</name>
//   </storage_domain>
//   <exclusive>true</exclusive>
//   <discard_snapshots>true</discard_snapshots>
// </action>
// ----
// This method supports the following parameters:
// `DiscardSnapshots`:: The `discard_snapshots` parameter is to be used when the virtual machine should be exported with all its
// snapshots collapsed.
// `Exclusive`:: The `exclusive` parameter is to be used when the virtual machine should be exported even if another copy of
// it already exists in the export domain (override).
// `Async`:: Indicates if the export should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) Export (
    async bool,
    discardSnapshots bool,
    exclusive bool,
    storageDomain StorageDomain,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        DiscardSnapshots: discardSnapshots,
        Exclusive: exclusive,
        StorageDomain: storageDomain,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'export', nil, headers, query, wait)
}

//
// Freeze virtual machine file systems.
// This operation freezes a virtual machine's file systems using the QEMU guest agent when taking a live snapshot of
// a running virtual machine. Normally, this is done automatically by the manager, but this must be executed
// manually with the API for virtual machines using OpenStack Volume (Cinder) disks.
// Example:
// [source]
// ----
// POST /ovirt-engine/api/vms/123/freezefilesystems
// ----
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the freeze should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) FreezeFilesystems (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'freezefilesystems', nil, headers, query, wait)
}

//
// Retrieves the description of the virtual machine.
// This method supports the following parameters:
// `NextRun`:: Indicates if the returned result describes the virtual machine as it is currently running, or if describes
// it with the modifications that have already been performed but that will have effect only when it is
// restarted. By default the values is `false`.
// If the parameter is included in the request, but without a value, it is assumed that the value is `true`, so
// the following request:
// [source]
// ----
// GET /vms/{vm:id};next_run
// ----
// Is equivalent to using the value `true`:
// [source]
// ----
// GET /vms/{vm:id};next_run=true
// ----
// `AllContent`:: Indicates if all the attributes of the virtual machine should be included in the response.
// By default the following attributes are excluded:
// - `console`
// - `initialization.configuration.data` - The OVF document describing the virtual machine.
// - `rng_source`
// - `soundcard`
// - `virtio_scsi`
// For example, to retrieve the complete representation of the virtual machine '123' send a request like this:
// ....
// GET /ovirt-engine/api/vms/123?all_content=true
// ....
// NOTE: The reason for not including these attributes is performance: they are seldom used and they require
// additional queries to the database. So try to use the this parameter only when it is really needed.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) Get (
    allContent bool,
    filter bool,
    nextRun bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if allContent != nil {
        query["all_content"] = allContent
    }
    if filter != nil {
        query["filter"] = filter
    }
    if nextRun != nil {
        query["next_run"] = nextRun
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Initiates the automatic user logon to access a virtual machine from an external console.
// This action requires the `ovirt-guest-agent-gdm-plugin` and the `ovirt-guest-agent-pam-module` packages to be
// installed and the `ovirt-guest-agent` service to be running on the virtual machine.
// Users require the appropriate user permissions for the virtual machine in order to access the virtual machine
// from an external console.
// This is how an example request would look like:
// [source]
// ----
// POST /ovirt-engine/api/vms/123/logon
// ----
// Request body:
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the logon should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) Logon (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'logon', nil, headers, query, wait)
}

//
// Sets the global maintenance mode on the hosted engine virtual machine.
// This action has no effect on other virtual machines.
// Example:
// [source]
// ----
// POST /ovirt-engine/api/vms/123/maintenance
// ----
// [source,xml]
// ----
// <action>
//   <maintenance_enabled>true<maintenance_enabled/>
// </action>
// ----
// This method supports the following parameters:
// `MaintenanceEnabled`:: Indicates if global maintenance should be enabled or disabled.
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) Maintenance (
    async bool,
    maintenanceEnabled bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        MaintenanceEnabled: maintenanceEnabled,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'maintenance', nil, headers, query, wait)
}

//
// This operation migrates a virtual machine to another physical host.
// [source]
// ----
// POST /ovirt-engine/api/vms/123/migrate
// ----
// One can specify a specific host to migrate the virtual machine to:
// [source,xml]
// ----
// <action>
//   <host id="2ab5e1da-b726-4274-bbf7-0a42b16a0fc3"/>
// </action>
// ----
// This method supports the following parameters:
// `Cluster`:: Specifies the cluster the virtual machine should migrate to. This is an optional parameter. By default, the
// virtual machine is migrated to another host within the same cluster.
// `Force`:: Specifies the virtual machine should migrate although it might be defined as non migratable. This is an
// optional parameter. By default, it is set to `false`.
// `Host`:: Specifies a specific host the virtual machine should migrate to. This is an optional parameters. By default,
// the oVirt Engine automatically selects a default host for migration within the same cluster. If an API user
// requires a specific host, the user can specify the host with either an `id` or `name` parameter.
// `Async`:: Indicates if the migration should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) Migrate (
    async bool,
    cluster Cluster,
    force bool,
    host Host,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Cluster: cluster,
        Force: force,
        Host: host,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'migrate', nil, headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the preview should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) PreviewSnapshot (
    async bool,
    disks []*Disk,
    restoreMemory bool,
    snapshot Snapshot,
    vm Vm,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Disks: disks,
        RestoreMemory: restoreMemory,
        Snapshot: snapshot,
        Vm: vm,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'previewsnapshot', nil, headers, query, wait)
}

//
// This operation sends a reboot request to a virtual machine.
// [source]
// ----
// POST /ovirt-engine/api/vms/123/reboot
// ----
// The reboot action does not take any action specific parameters, so the request body should contain an
// empty `action`:
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the reboot should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) Reboot (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'reboot', nil, headers, query, wait)
}

//
// Removes the virtual machine, including the virtual disks attached to it.
// For example, to remove the virtual machine with identifier `123` send a request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/vms/123
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `DetachOnly`:: Indicates if the attached virtual disks should be detached first and preserved instead of being removed.
// `Force`:: Indicates if the virtual machine should be forcibly removed.
// Locked virtual machines and virtual machines with locked disk images
// cannot be removed without this flag set to true.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) Remove (
    async bool,
    detachOnly bool,
    force bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }
    if detachOnly != nil {
        query["detach_only"] = detachOnly
    }
    if force != nil {
        query["force"] = force
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) ReorderMacAddresses (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'reordermacaddresses', nil, headers, query, wait)
}

//
// This operation sends a shutdown request to a virtual machine.
// [source]
// ----
// POST /ovirt-engine/api/vms/123/shutdown
// ----
// The shutdown action does not take any action specific parameters,
// so the request body should contain an empty `action`:
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the shutdown should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) Shutdown (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'shutdown', nil, headers, query, wait)
}

//
// Starts the virtual machine.
// If the virtual environment is complete and the virtual machine contains all necessary components to function,
// it can be started.
// This example starts the virtual machine:
// [source]
// ----
// POST /ovirt-engine/api/vms/123/start
// ----
// With a request body:
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Pause`:: If set to `true`, start the virtual machine in paused mode. Default is `false`.
// `Vm`:: The definition of the virtual machine for this specific run.
// For example:
// [source,xml]
// ----
// <action>
//   <vm>
//     <os>
//       <boot>
//         <devices>
//           <device>cdrom</device>
//         </devices>
//       </boot>
//     </os>
//   </vm>
// </action>
// ----
// This will set the boot device to the CDROM only for this specific start. After the virtual machine will be
// powered off, this definition will be reverted.
// `UseCloudInit`:: If set to `true`, the initialization type is set to _cloud-init_. The default value is `false`.
// See https://cloudinit.readthedocs.io/en/latest[this] for details.
// `UseSysprep`:: If set to `true`, the initialization type is set to _Sysprep_. The default value is `false`.
// See https://en.wikipedia.org/wiki/Sysprep[this] for details.
// `Async`:: Indicates if the action should be performed asynchronously.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) Start (
    async bool,
    filter bool,
    pause bool,
    useCloudInit bool,
    useSysprep bool,
    vm Vm,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Filter: filter,
        Pause: pause,
        UseCloudInit: useCloudInit,
        UseSysprep: useSysprep,
        Vm: vm,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'start', nil, headers, query, wait)
}

//
// This operation forces a virtual machine to power-off.
// [source]
// ----
// POST /ovirt-engine/api/vms/123/stop
// ----
// The stop action does not take any action specific parameters,
// so the request body should contain an empty `action`:
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) Stop (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'stop', nil, headers, query, wait)
}

//
// This operation saves the virtual machine state to disk and stops it.
// Start a suspended virtual machine and restore the virtual machine state with the start action.
// [source]
// ----
// POST /ovirt-engine/api/vms/123/suspend
// ----
// The suspend action does not take any action specific parameters,
// so the request body should contain an empty `action`:
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) Suspend (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'suspend', nil, headers, query, wait)
}

//
// Thaw virtual machine file systems.
// This operation thaws a virtual machine's file systems using the QEMU guest agent when taking a live snapshot of a
// running virtual machine. Normally, this is done automatically by the manager, but this must be executed manually
// with the API for virtual machines using OpenStack Volume (Cinder) disks.
// Example:
// [source]
// ----
// POST /api/vms/123/thawfilesystems
// ----
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) ThawFilesystems (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'thawfilesystems', nil, headers, query, wait)
}

//
// Generates a time-sensitive authentication token for accessing a virtual machine's display.
// [source]
// ----
// POST /ovirt-engine/api/vms/123/ticket
// ----
// The client-provided action optionally includes a desired ticket value and/or an expiry time in seconds.
// In any case, the response specifies the actual ticket value and expiry used.
// [source,xml]
// ----
// <action>
//   <ticket>
//     <value>abcd12345</value>
//     <expiry>120</expiry>
//   </ticket>
// </action>
// ----
// [IMPORTANT]
// ====
// If the virtual machine is configured to support only one graphics protocol
// then the generated authentication token will be valid for that protocol.
// But if the virtual machine is configured to support multiple protocols,
// VNC and SPICE, then the authentication token will only be valid for
// the SPICE protocol.
// In order to obtain an authentication token for a specific protocol, for
// example for VNC, use the `ticket` method of the <<services/vm_graphics_console,
// service>> that manages the graphics consoles of the virtual machine, sending
// a request like this:
// [source]
// ----
// POST /ovirt-engine/api/vms/123/graphicsconsoles/456/ticket
// ----
// ====
// This method supports the following parameters:
// `Async`:: Indicates if the generation of the ticket should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) Ticket (
    async bool,
    ticket Ticket,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Ticket: ticket,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'ticket', 'ticket', headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmService) UndoSnapshot (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'undosnapshot', nil, headers, query, wait)
}

//
//
func (op *VmService) Update (
    vm Vm,
    async bool,
    nextRun bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }
    if nextRun != nil {
        query["next_run"] = nextRun
    }

    // Send the request
    return op.internalUpdate(vm, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmService) Service(path string) *VmService {
    if path == nil {
        return op
    }
    if path == "affinitylabels" {
        return &AffinityLabelsService{}
    }
    if path.startswith('affinitylabels/') {
        return AffinityLabelsService().service(path[15:])
    }
    if path == "applications" {
        return &ApplicationsService{}
    }
    if path.startswith('applications/') {
        return ApplicationsService().service(path[13:])
    }
    if path == "cdroms" {
        return &CdromsService{}
    }
    if path.startswith('cdroms/') {
        return CdromsService().service(path[7:])
    }
    if path == "diskattachments" {
        return &DiskAttachmentsService{}
    }
    if path.startswith('diskattachments/') {
        return DiskAttachmentsService().service(path[16:])
    }
    if path == "graphicsconsoles" {
        return &GraphicsConsolesService{}
    }
    if path.startswith('graphicsconsoles/') {
        return GraphicsConsolesService().service(path[17:])
    }
    if path == "hostdevices" {
        return &HostDevicesService{}
    }
    if path.startswith('hostdevices/') {
        return HostDevicesService().service(path[12:])
    }
    if path == "katelloerrata" {
        return &KatelloErrataService{}
    }
    if path.startswith('katelloerrata/') {
        return KatelloErrataService().service(path[14:])
    }
    if path == "nics" {
        return &NicsService{}
    }
    if path.startswith('nics/') {
        return NicsService().service(path[5:])
    }
    if path == "numanodes" {
        return &NumaNodesService{}
    }
    if path.startswith('numanodes/') {
        return NumaNodesService().service(path[10:])
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    if path == "reporteddevices" {
        return &ReportedDevicesService{}
    }
    if path.startswith('reporteddevices/') {
        return ReportedDevicesService().service(path[16:])
    }
    if path == "sessions" {
        return &SessionsService{}
    }
    if path.startswith('sessions/') {
        return SessionsService().service(path[9:])
    }
    if path == "snapshots" {
        return &SnapshotsService{}
    }
    if path.startswith('snapshots/') {
        return SnapshotsService().service(path[10:])
    }
    if path == "statistics" {
        return &StatisticsService{}
    }
    if path.startswith('statistics/') {
        return StatisticsService().service(path[11:])
    }
    if path == "tags" {
        return &TagsService{}
    }
    if path.startswith('tags/') {
        return TagsService().service(path[5:])
    }
    if path == "watchdogs" {
        return &WatchdogsService{}
    }
    if path.startswith('watchdogs/') {
        return WatchdogsService().service(path[10:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *VmService) String() string {
    return fmt.Sprintf("VmService:%s", op.Path)
}


//
// A service that provides information about an application installed in a virtual machine.
//
type VmApplicationService struct {
    Service

}

func NewVmApplicationService(connection *Connection, path string) *VmApplicationService {
    var result VmApplicationService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Returns the information about the application.
// This method supports the following parameters:
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmApplicationService) Get (
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmApplicationService) Service(path string) *VmApplicationService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *VmApplicationService) String() string {
    return fmt.Sprintf("VmApplicationService:%s", op.Path)
}


//
// A service that provides information about applications installed in a virtual machine.
//
type VmApplicationsService struct {
    Service

    ApplicationService  *ApplicationService
}

func NewVmApplicationsService(connection *Connection, path string) *VmApplicationsService {
    var result VmApplicationsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Returns a list of applications installed in the virtual machine.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of applications to return. If not specified all the applications are returned.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmApplicationsService) List (
    filter bool,
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmApplicationsService) Service(path string) *VmApplicationsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Application_service(path)
    return self.Application_service(path[:index]).Service(path[index + 1:])
}

func (op *VmApplicationsService) String() string {
    return fmt.Sprintf("VmApplicationsService:%s", op.Path)
}


//
// Manages a CDROM device of a virtual machine.
// Changing and ejecting the disk is done using always the `update` method, to change the value of the `file`
// attribute.
//
type VmCdromService struct {
    Service

}

func NewVmCdromService(connection *Connection, path string) *VmCdromService {
    var result VmCdromService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Returns the information about this CDROM device.
// The information consists of `cdrom` attribute containing reference to the CDROM device, the virtual machine,
// and optionally the inserted disk.
// If there is a disk inserted then the `file` attribute will contain a reference to the ISO image:
// [source,xml]
// ----
// <cdrom href="..." id="00000000-0000-0000-0000-000000000000">
//   <file id="mycd.iso"/>
//   <vm href="/ovirt-engine/api/vms/123" id="123"/>
// </cdrom>
// ----
// If there is no disk inserted then the `file` attribute won't be reported:
// [source,xml]
// ----
// <cdrom href="..." id="00000000-0000-0000-0000-000000000000">
//   <vm href="/ovirt-engine/api/vms/123" id="123"/>
// </cdrom>
// ----
// This method supports the following parameters:
// `Current`:: Indicates if the operation should return the information for the currently running virtual machine. This
// parameter is optional, and the default value is `false`.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmCdromService) Get (
    current bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if current != nil {
        query["current"] = current
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Updates the information about this CDROM device.
// It allows to change or eject the disk by changing the value of the `file` attribute.
// For example, to insert or change the disk send a request like this:
// [source]
// ----
// PUT /ovirt-engine/api/vms/123/cdroms/00000000-0000-0000-0000-000000000000
// ----
// The body should contain the new value for the `file` attribute:
// [source,xml]
// ----
// <cdrom>
//   <file id="mycd.iso"/>
// </cdrom>
// ----
// The value of the `id` attribute, `mycd.iso` in this example, should correspond to a file available in an
// attached ISO storage domain.
// To eject the disk use a `file` with an empty `id`:
// [source,xml]
// ----
// <cdrom>
//   <file id=""/>
// </cdrom>
// ----
// By default the above operations change permanently the disk that will be visible to the virtual machine
// after the next boot, but they don't have any effect on the currently running virtual machine. If you want
// to change the disk that is visible to the current running virtual machine, add the `current=true` parameter.
// For example, to eject the current disk send a request like this:
// [source]
// ----
// PUT /ovirt-engine/api/vms/123/cdroms/00000000-0000-0000-0000-000000000000?current=true
// ----
// With a request body like this:
// [source,xml]
// ----
// <cdrom>
//   <file id=""/>
// </cdrom>
// ----
// IMPORTANT: The changes made with the `current=true` parameter are never persisted, so they won't have any
// effect after the virtual machine is rebooted.
// This method supports the following parameters:
// `Cdrom`:: The information about the CDROM device.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmCdromService) Update (
    cdrom Cdrom,
    current bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if current != nil {
        query["current"] = current
    }

    // Send the request
    return op.internalUpdate(cdrom, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmCdromService) Service(path string) *VmCdromService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *VmCdromService) String() string {
    return fmt.Sprintf("VmCdromService:%s", op.Path)
}


//
// Manages the CDROM devices of a virtual machine.
// Currently virtual machines have exactly one CDROM device. No new devices can be added, and the existing one can't
// be removed, thus there are no `add` or `remove` methods. Changing and ejecting CDROM disks is done with the
// <<services/vm_cdrom/methods/update, update>> method of the <<services/vm_cdrom, service>> that manages the
// CDROM device.
//
type VmCdromsService struct {
    Service

    CdromService  *CdromService
}

func NewVmCdromsService(connection *Connection, path string) *VmCdromsService {
    var result VmCdromsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Returns the list of CDROM devices of the virtual machine.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of CDROMs to return. If not specified all the CDROMs are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmCdromsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmCdromsService) Service(path string) *VmCdromsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Cdrom_service(path)
    return self.Cdrom_service(path[:index]).Service(path[index + 1:])
}

func (op *VmCdromsService) String() string {
    return fmt.Sprintf("VmCdromsService:%s", op.Path)
}


//
//
type VmDiskService struct {
    MeasurableService

    PermissionsService  *PermissionsService
    StatisticsService  *StatisticsService
}

func NewVmDiskService(connection *Connection, path string) *VmDiskService {
    var result VmDiskService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the activation should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmDiskService) Activate (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'activate', nil, headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the deactivation should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmDiskService) Deactivate (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'deactivate', nil, headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the export should be performed asynchronously.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmDiskService) Export (
    async bool,
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Filter: filter,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'export', nil, headers, query, wait)
}

//
//
func (op *VmDiskService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the move should be performed asynchronously.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmDiskService) Move (
    async bool,
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Filter: filter,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'move', nil, headers, query, wait)
}

//
// Detach the disk from the virtual machine.
// NOTE: In version 3 of the API this used to also remove the disk completely from the system, but starting with
// version 4 it doesn't. If you need to remove it completely use the <<services/disk/methods/remove,remove
// method of the top level disk service>>.
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmDiskService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
//
func (op *VmDiskService) Update (
    disk Disk,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(disk, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmDiskService) Service(path string) *VmDiskService {
    if path == nil {
        return op
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    if path == "statistics" {
        return &StatisticsService{}
    }
    if path.startswith('statistics/') {
        return StatisticsService().service(path[11:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *VmDiskService) String() string {
    return fmt.Sprintf("VmDiskService:%s", op.Path)
}


//
//
type VmDisksService struct {
    Service

    DiskService  *DiskService
}

func NewVmDisksService(connection *Connection, path string) *VmDisksService {
    var result VmDisksService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *VmDisksService) Add (
    disk Disk,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(disk, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of disks to return. If not specified all the disks are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmDisksService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmDisksService) Service(path string) *VmDisksService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Disk_service(path)
    return self.Disk_service(path[:index]).Service(path[index + 1:])
}

func (op *VmDisksService) String() string {
    return fmt.Sprintf("VmDisksService:%s", op.Path)
}


//
//
type VmGraphicsConsoleService struct {
    Service

}

func NewVmGraphicsConsoleService(connection *Connection, path string) *VmGraphicsConsoleService {
    var result VmGraphicsConsoleService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Gets graphics console configuration of the virtual machine.
// This method supports the following parameters:
// `Current`:: Use the following query to obtain the current run-time configuration of the graphics console.
// [source]
// ----
// GET /ovit-engine/api/vms/123/graphicsconsoles/456?current=true
// ----
// The default value is `false`.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmGraphicsConsoleService) Get (
    current bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if current != nil {
        query["current"] = current
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the generation of the ticket should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmGraphicsConsoleService) ProxyTicket (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'proxyticket', 'proxyTicket', headers, query, wait)
}

//
// Generates the file which is compatible with `remote-viewer` client.
// Use the following request to generate remote viewer connection file of the graphics console.
// Note that this action generates the file only if virtual machine is running.
// [source]
// ----
// POST /ovirt-engine/api/vms/123/graphicsconsoles/456/remoteviewerconnectionfile
// ----
// The `remoteviewerconnectionfile` action does not take any action specific parameters,
// so the request body should contain an empty `action`:
// [source,xml]
// ----
// <action/>
// ----
// The response contains the file, which can be used with `remote-viewer` client.
// [source,xml]
// ----
// <action>
//   <remote_viewer_connection_file>
//     [virt-viewer]
//     type=spice
//     host=192.168.1.101
//     port=-1
//     password=123456789
//     delete-this-file=1
//     fullscreen=0
//     toggle-fullscreen=shift+f11
//     release-cursor=shift+f12
//     secure-attention=ctrl+alt+end
//     tls-port=5900
//     enable-smartcard=0
//     enable-usb-autoshare=0
//     usb-filter=null
//     tls-ciphers=DEFAULT
//     host-subject=O=local,CN=example.com
//     ca=...
//   </remote_viewer_connection_file>
// </action>
// ----
// E.g., to fetch the content of remote viewer connection file and save it into temporary file, user can use
// oVirt Python SDK as follows:
// [source,python]
// ----
// # Find the virtual machine:
// vm = vms_service.list(search='name=myvm')[0]
// # Locate the service that manages the virtual machine, as that is where
// # the locators are defined:
// vm_service = vms_service.vm_service(vm.id)
// # Find the graphic console of the virtual machine:
// graphics_consoles_service = vm_service.graphics_consoles_service()
// graphics_console = graphics_consoles_service.list()[0]
// # Generate the remote viewer connection file:
// console_service = graphics_consoles_service.console_service(graphics_console.id)
// remote_viewer_connection_file = console_service.remote_viewer_connection_file()
// # Write the content to file "/tmp/remote_viewer_connection_file.vv"
// path = "/tmp/remote_viewer_connection_file.vv"
// with open(path, "w") as f:
//     f.write(remote_viewer_connection_file)
// ----
// When you create the remote viewer connection file, then you can connect to virtual machine graphic console,
// as follows:
// [source,bash]
// ----
// #!/bin/sh -ex
// remote-viewer --ovirt-ca-file=/etc/pki/ovirt-engine/ca.pem /tmp/remote_viewer_connection_file.vv
// ----
//
func (op *VmGraphicsConsoleService) RemoteViewerConnectionFile (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
    }

    // Send the request and wait for the response:
    return internalAction(action, 'remoteviewerconnectionfile', 'remoteViewerConnectionFile', headers, query, wait)
}

//
// Remove the graphics console from the virtual machine.
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmGraphicsConsoleService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Generates a time-sensitive authentication token for accessing this virtual machine's console.
// [source]
// ----
// POST /ovirt-engine/api/vms/123/graphicsconsoles/456/ticket
// ----
// The client-provided action optionally includes a desired ticket value and/or an expiry time in seconds.
// In any case, the response specifies the actual ticket value and expiry used.
// [source,xml]
// ----
// <action>
//   <ticket>
//     <value>abcd12345</value>
//     <expiry>120</expiry>
//   </ticket>
// </action>
// ----
// This method supports the following parameters:
// `Ticket`:: The generated ticket that can be used to access this console.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmGraphicsConsoleService) Ticket (
    ticket Ticket,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Ticket: ticket,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'ticket', 'ticket', headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmGraphicsConsoleService) Service(path string) *VmGraphicsConsoleService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *VmGraphicsConsoleService) String() string {
    return fmt.Sprintf("VmGraphicsConsoleService:%s", op.Path)
}


//
//
type VmGraphicsConsolesService struct {
    Service

    ConsoleService  *ConsoleService
}

func NewVmGraphicsConsolesService(connection *Connection, path string) *VmGraphicsConsolesService {
    var result VmGraphicsConsolesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add new graphics console to the virtual machine.
//
func (op *VmGraphicsConsolesService) Add (
    console GraphicsConsole,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(console, headers, query, wait)
}

//
// Lists all the configured graphics consoles of the virtual machine.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of consoles to return. If not specified all the consoles are returned.
// `Current`:: Use the following query to obtain the current run-time configuration of the graphics consoles.
// [source]
// ----
// GET /ovirt-engine/api/vms/123/graphicsconsoles?current=true
// ----
// The default value is `false`.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmGraphicsConsolesService) List (
    current bool,
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if current != nil {
        query["current"] = current
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmGraphicsConsolesService) Service(path string) *VmGraphicsConsolesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Console_service(path)
    return self.Console_service(path[:index]).Service(path[index + 1:])
}

func (op *VmGraphicsConsolesService) String() string {
    return fmt.Sprintf("VmGraphicsConsolesService:%s", op.Path)
}


//
// A service to manage individual host device attached to a virtual machine.
//
type VmHostDeviceService struct {
    Service

}

func NewVmHostDeviceService(connection *Connection, path string) *VmHostDeviceService {
    var result VmHostDeviceService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieve information about particular host device attached to given virtual machine.
// Example:
// [source]
// ----
// GET /ovirt-engine/api/vms/123/hostdevices/456
// ----
// [source,xml]
// ----
// <host_device href="/ovirt-engine/api/hosts/543/devices/456" id="456">
//   <name>pci_0000_04_00_0</name>
//   <capability>pci</capability>
//   <iommu_group>30</iommu_group>
//   <placeholder>true</placeholder>
//   <product id="0x13ba">
//     <name>GM107GL [Quadro K2200]</name>
//   </product>
//   <vendor id="0x10de">
//     <name>NVIDIA Corporation</name>
//   </vendor>
//   <host href="/ovirt-engine/api/hosts/543" id="543"/>
//   <parent_device href="/ovirt-engine/api/hosts/543/devices/456" id="456">
//     <name>pci_0000_00_03_0</name>
//   </parent_device>
//   <vm href="/ovirt-engine/api/vms/123" id="123"/>
// </host_device>
// ----
//
func (op *VmHostDeviceService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Remove the attachment of this host device from given virtual machine.
// NOTE: In case this device serves as an IOMMU placeholder, it cannot be removed (remove will result only
// in setting its `placeholder` flag to `true`). Note that all IOMMU placeholder devices will be removed
// automatically as soon as there will be no more non-placeholder devices (all devices from given IOMMU
// group are detached).
// [source]
// ----
// DELETE /ovirt-engine/api/vms/123/hostdevices/456
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmHostDeviceService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmHostDeviceService) Service(path string) *VmHostDeviceService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *VmHostDeviceService) String() string {
    return fmt.Sprintf("VmHostDeviceService:%s", op.Path)
}


//
// A service to manage host devices attached to a virtual machine.
//
type VmHostDevicesService struct {
    Service

    DeviceService  *DeviceService
}

func NewVmHostDevicesService(connection *Connection, path string) *VmHostDevicesService {
    var result VmHostDevicesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Attach target device to given virtual machine.
// Example:
// [source]
// ----
// POST /ovirt-engine/api/vms/123/hostdevices
// ----
// With request body of type <<types/host_device,HostDevice>>, for example
// [source,xml]
// ----
// <host_device id="123" />
// ----
// NOTE: A necessary precondition for a successful host device attachment is that the virtual machine must be pinned
// to *exactly* one host. The device ID is then taken relative to this host.
// NOTE: Attachment of a PCI device that is part of a bigger IOMMU group will result in attachment of the remaining
// devices from that IOMMU group as "placeholders". These devices are then identified using the `placeholder`
// attribute of the <<types/host_device,HostDevice>> type set to `true`.
// In case you want attach a device that already serves as an IOMMU placeholder, simply issue an explicit Add operation
// for it, and its `placeholder` flag will be cleared, and the device will be accessible to the virtual machine.
// This method supports the following parameters:
// `Device`:: The host device to be attached to given virtual machine.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmHostDevicesService) Add (
    device HostDevice,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(device, headers, query, wait)
}

//
// List the host devices assigned to given virtual machine.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of devices to return. If not specified all the devices are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmHostDevicesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmHostDevicesService) Service(path string) *VmHostDevicesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Device_service(path)
    return self.Device_service(path[:index]).Service(path[index + 1:])
}

func (op *VmHostDevicesService) String() string {
    return fmt.Sprintf("VmHostDevicesService:%s", op.Path)
}


//
//
type VmNicService struct {
    MeasurableService

    NetworkFilterParametersService  *NetworkFilterParametersService
    ReportedDevicesService  *ReportedDevicesService
    StatisticsService  *StatisticsService
}

func NewVmNicService(connection *Connection, path string) *VmNicService {
    var result VmNicService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the activation should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmNicService) Activate (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'activate', nil, headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the deactivation should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmNicService) Deactivate (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'deactivate', nil, headers, query, wait)
}

//
//
func (op *VmNicService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes the NIC.
// For example, to remove the NIC with id `456` from the virtual machine with id `123` send a request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/vms/123/nics/456
// ----
// [IMPORTANT]
// ====
// The hotplugging feature only supports virtual machine operating systems with hotplugging operations.
// Example operating systems include:
// - Red Hat Enterprise Linux 6
// - Red Hat Enterprise Linux 5
// - Windows Server 2008 and
// - Windows Server 2003
// ====
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmNicService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates the NIC.
// For example, to update the NIC having with `456` belonging to virtual the machine with id `123` send a request
// like this:
// [source]
// ----
// PUT /ovirt-engine/api/vms/123/nics/456
// ----
// With a request body like this:
// [source,xml]
// ----
// <nic>
//   <name>mynic</name>
//   <interface>e1000</interface>
//   <vnic_profile id='789'/>
// </nic>
// ----
// [IMPORTANT]
// ====
// The hotplugging feature only supports virtual machine operating systems with hotplugging operations.
// Example operating systems include:
// - Red Hat Enterprise Linux 6
// - Red Hat Enterprise Linux 5
// - Windows Server 2008 and
// - Windows Server 2003
// ====
//
func (op *VmNicService) Update (
    nic Nic,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(nic, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmNicService) Service(path string) *VmNicService {
    if path == nil {
        return op
    }
    if path == "networkfilterparameters" {
        return &NetworkFilterParametersService{}
    }
    if path.startswith('networkfilterparameters/') {
        return NetworkFilterParametersService().service(path[24:])
    }
    if path == "reporteddevices" {
        return &ReportedDevicesService{}
    }
    if path.startswith('reporteddevices/') {
        return ReportedDevicesService().service(path[16:])
    }
    if path == "statistics" {
        return &StatisticsService{}
    }
    if path.startswith('statistics/') {
        return StatisticsService().service(path[11:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *VmNicService) String() string {
    return fmt.Sprintf("VmNicService:%s", op.Path)
}


//
//
type VmNicsService struct {
    Service

    NicService  *NicService
}

func NewVmNicsService(connection *Connection, path string) *VmNicsService {
    var result VmNicsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Adds a NIC to the virtual machine.
// The following example adds a network interface named `mynic` using `virtio` and the `ovirtmgmt` network to the
// virtual machine.
// [source]
// ----
// POST /ovirt-engine/api/vms/123/nics
// ----
// [source,xml]
// ----
// <nic>
//   <interface>virtio</interface>
//   <name>mynic</name>
//   <network>
//     <name>ovirtmgmt</name>
//   </network>
// </nic>
// ----
// The following example sends that request using `curl`:
// [source,bash]
// ----
// curl \
// --request POST \
// --header "Version: 4" \
// --header "Content-Type: application/xml" \
// --header "Accept: application/xml" \
// --user "admin@internal:mypassword" \
// --cacert /etc/pki/ovirt-engine/ca.pem \
// --data '
// <nic>
//   <name>mynic</name>
//   <network>
//     <name>ovirtmgmt</name>
//   </network>
// </nic>
// ' \
// https://myengine.example.com/ovirt-engine/api/vms/123/nics
// ----
// [IMPORTANT]
// ====
// The hotplugging feature only supports virtual machine operating systems with hotplugging operations.
// Example operating systems include:
// - Red Hat Enterprise Linux 6
// - Red Hat Enterprise Linux 5
// - Windows Server 2008 and
// - Windows Server 2003
// ====
//
func (op *VmNicsService) Add (
    nic Nic,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(nic, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of NICs to return. If not specified all the NICs are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmNicsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmNicsService) Service(path string) *VmNicsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Nic_service(path)
    return self.Nic_service(path[:index]).Service(path[index + 1:])
}

func (op *VmNicsService) String() string {
    return fmt.Sprintf("VmNicsService:%s", op.Path)
}


//
//
type VmNumaNodeService struct {
    Service

}

func NewVmNumaNodeService(connection *Connection, path string) *VmNumaNodeService {
    var result VmNumaNodeService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *VmNumaNodeService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes a virtual NUMA node.
// An example of removing a virtual NUMA node:
// [source]
// ----
// DELETE /ovirt-engine/api/vms/123/numanodes/456
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmNumaNodeService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates a virtual NUMA node.
// An example of pinning a virtual NUMA node to a physical NUMA node on the host:
// [source]
// ----
// PUT /ovirt-engine/api/vms/123/numanodes/456
// ----
// The request body should contain the following:
// [source,xml]
// ----
// <vm_numa_node>
//   <numa_node_pins>
//     <numa_node_pin>
//       <index>0</index>
//     </numa_node_pin>
//   </numa_node_pins>
// </vm_numa_node>
// ----
//
func (op *VmNumaNodeService) Update (
    node VirtualNumaNode,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(node, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmNumaNodeService) Service(path string) *VmNumaNodeService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *VmNumaNodeService) String() string {
    return fmt.Sprintf("VmNumaNodeService:%s", op.Path)
}


//
//
type VmNumaNodesService struct {
    Service

    NodeService  *NodeService
}

func NewVmNumaNodesService(connection *Connection, path string) *VmNumaNodesService {
    var result VmNumaNodesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a new virtual NUMA node for the virtual machine.
// An example of creating a NUMA node:
// [source]
// ----
// POST /ovirt-engine/api/vms/c7ecd2dc/numanodes
// Accept: application/xml
// Content-type: application/xml
// ----
// The request body can contain the following:
// [source,xml]
// ----
// <vm_numa_node>
//   <cpu>
//     <cores>
//       <core>
//         <index>0</index>
//       </core>
//     </cores>
//   </cpu>
//   <index>0</index>
//   <memory>1024</memory>
// </vm_numa_node>
// ----
//
func (op *VmNumaNodesService) Add (
    node VirtualNumaNode,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(node, headers, query, wait)
}

//
// Lists virtual NUMA nodes of a virtual machine.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of nodes to return. If not specified all the nodes are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmNumaNodesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmNumaNodesService) Service(path string) *VmNumaNodesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Node_service(path)
    return self.Node_service(path[:index]).Service(path[index + 1:])
}

func (op *VmNumaNodesService) String() string {
    return fmt.Sprintf("VmNumaNodesService:%s", op.Path)
}


//
// A service to manage a virtual machines pool.
//
type VmPoolService struct {
    Service

    PermissionsService  *PermissionsService
}

func NewVmPoolService(connection *Connection, path string) *VmPoolService {
    var result VmPoolService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This operation allocates a virtual machine in the virtual machine pool.
// [source]
// ----
// POST /ovirt-engine/api/vmpools/123/allocatevm
// ----
// The allocate virtual machine action does not take any action specific parameters, so the request body should
// contain an empty `action`:
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the allocation should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmPoolService) AllocateVm (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'allocatevm', nil, headers, query, wait)
}

//
// Get the virtual machine pool.
// [source]
// ----
// GET /ovirt-engine/api/vmpools/123
// ----
// You will get a XML response like that one:
// [source,xml]
// ----
// <vm_pool id="123">
//   <actions>...</actions>
//   <name>MyVmPool</name>
//   <description>MyVmPool description</description>
//   <link href="/ovirt-engine/api/vmpools/123/permissions" rel="permissions"/>
//   <max_user_vms>1</max_user_vms>
//   <prestarted_vms>0</prestarted_vms>
//   <size>100</size>
//   <stateful>false</stateful>
//   <type>automatic</type>
//   <use_latest_template_version>false</use_latest_template_version>
//   <cluster id="123"/>
//   <template id="123"/>
//   <vm id="123">...</vm>
//   ...
// </vm_pool>
// ----
// This method supports the following parameters:
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmPoolService) Get (
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes a virtual machine pool.
// [source]
// ----
// DELETE /ovirt-engine/api/vmpools/123
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmPoolService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Update the virtual machine pool.
// [source]
// ----
// PUT /ovirt-engine/api/vmpools/123
// ----
// The `name`, `description`, `size`, `prestarted_vms` and `max_user_vms`
// attributes can be updated after the virtual machine pool has been
// created.
// [source,xml]
// ----
// <vmpool>
//   <name>VM_Pool_B</name>
//   <description>Virtual Machine Pool B</description>
//   <size>3</size>
//   <prestarted_vms>1</size>
//   <max_user_vms>2</size>
// </vmpool>
// ----
// This method supports the following parameters:
// `Pool`:: The virtual machine pool that is being updated.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmPoolService) Update (
    pool VmPool,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(pool, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmPoolService) Service(path string) *VmPoolService {
    if path == nil {
        return op
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *VmPoolService) String() string {
    return fmt.Sprintf("VmPoolService:%s", op.Path)
}


//
// Provides read-write access to virtual machines pools.
//
type VmPoolsService struct {
    Service

    PoolService  *PoolService
}

func NewVmPoolsService(connection *Connection, path string) *VmPoolsService {
    var result VmPoolsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a new virtual machine pool.
// A new pool requires the `name`, `cluster` and `template` attributes. Identify the cluster and template with the
// `id` or `name` nested attributes:
// [source]
// ----
// POST /ovirt-engine/api/vmpools
// ----
// With the following body:
// [source,xml]
// ----
// <vmpool>
//   <name>mypool</name>
//   <cluster id="123"/>
//   <template id="456"/>
// </vmpool>
// ----
// This method supports the following parameters:
// `Pool`:: Pool to add.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmPoolsService) Add (
    pool VmPool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(pool, headers, query, wait)
}

//
// Get a list of available virtual machines pools.
// [source]
// ----
// GET /ovirt-engine/api/vmpools
// ----
// You will receive the following response:
// [source,xml]
// ----
// <vm_pools>
//   <vm_pool id="123">
//     ...
//   </vm_pool>
//   ...
// </vm_pools>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of pools to return. If this value is not specified, all of the pools are returned.
// `Search`:: A query string used to restrict the returned pools.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmPoolsService) List (
    caseSensitive bool,
    filter bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if filter != nil {
        query["filter"] = filter
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmPoolsService) Service(path string) *VmPoolsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Pool_service(path)
    return self.Pool_service(path[:index]).Service(path[index + 1:])
}

func (op *VmPoolsService) String() string {
    return fmt.Sprintf("VmPoolsService:%s", op.Path)
}


//
//
type VmReportedDeviceService struct {
    Service

}

func NewVmReportedDeviceService(connection *Connection, path string) *VmReportedDeviceService {
    var result VmReportedDeviceService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *VmReportedDeviceService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmReportedDeviceService) Service(path string) *VmReportedDeviceService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *VmReportedDeviceService) String() string {
    return fmt.Sprintf("VmReportedDeviceService:%s", op.Path)
}


//
//
type VmReportedDevicesService struct {
    Service

    ReportedDeviceService  *ReportedDeviceService
}

func NewVmReportedDevicesService(connection *Connection, path string) *VmReportedDevicesService {
    var result VmReportedDevicesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of devices to return. If not specified all the devices are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmReportedDevicesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmReportedDevicesService) Service(path string) *VmReportedDevicesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.ReportedDevice_service(path)
    return self.ReportedDevice_service(path[:index]).Service(path[index + 1:])
}

func (op *VmReportedDevicesService) String() string {
    return fmt.Sprintf("VmReportedDevicesService:%s", op.Path)
}


//
//
type VmSessionService struct {
    Service

}

func NewVmSessionService(connection *Connection, path string) *VmSessionService {
    var result VmSessionService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *VmSessionService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmSessionService) Service(path string) *VmSessionService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *VmSessionService) String() string {
    return fmt.Sprintf("VmSessionService:%s", op.Path)
}


//
// Provides information about virtual machine user sessions.
//
type VmSessionsService struct {
    Service

    SessionService  *SessionService
}

func NewVmSessionsService(connection *Connection, path string) *VmSessionsService {
    var result VmSessionsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Lists all user sessions for this virtual machine.
// For example, to retrieve the session information for virtual machine `123` send a request like this:
// [source]
// ----
// GET /ovirt-engine/api/vms/123/sessions
// ----
// The response body will contain something like this:
// [source,xml]
// ----
// <sessions>
//   <session href="/ovirt-engine/api/vms/123/sessions/456" id="456">
//     <console_user>true</console_user>
//     <ip>
//       <address>192.168.122.1</address>
//     </ip>
//     <user href="/ovirt-engine/api/users/789" id="789"/>
//     <vm href="/ovirt-engine/api/vms/123" id="123"/>
//   </session>
//   ...
// </sessions>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of sessions to return. If not specified all the sessions are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmSessionsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmSessionsService) Service(path string) *VmSessionsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Session_service(path)
    return self.Session_service(path[:index]).Service(path[index + 1:])
}

func (op *VmSessionsService) String() string {
    return fmt.Sprintf("VmSessionsService:%s", op.Path)
}


//
// A service managing a watchdog on virtual machines.
//
type VmWatchdogService struct {
    Service

}

func NewVmWatchdogService(connection *Connection, path string) *VmWatchdogService {
    var result VmWatchdogService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Returns the information about the watchdog.
//
func (op *VmWatchdogService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes the watchdog from the virtual machine.
// For example, to remove a watchdog from a virtual machine, send a request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/vms/123/watchdogs/00000000-0000-0000-0000-000000000000
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmWatchdogService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates the information about the watchdog.
// You can update the information using `action` and `model` elements.
// For example, to update a watchdog, send a request like this:
// [source]
// ----
// PUT /ovirt-engine/api/vms/123/watchdogs
// <watchdog>
//   <action>reset</action>
// </watchdog>
// ----
// with response body:
// [source,xml]
// ----
// <watchdog href="/ovirt-engine/api/vms/123/watchdogs/00000000-0000-0000-0000-000000000000" id="00000000-0000-0000-0000-000000000000">
//   <vm href="/ovirt-engine/api/vms/123" id="123"/>
//   <action>reset</action>
//   <model>i6300esb</model>
// </watchdog>
// ----
// This method supports the following parameters:
// `Watchdog`:: The information about the watchdog.
// The request data must contain at least one of `model` and `action`
// elements. The response data contains complete information about the
// updated watchdog.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmWatchdogService) Update (
    watchdog Watchdog,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(watchdog, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmWatchdogService) Service(path string) *VmWatchdogService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *VmWatchdogService) String() string {
    return fmt.Sprintf("VmWatchdogService:%s", op.Path)
}


//
// Lists the watchdogs of a virtual machine.
//
type VmWatchdogsService struct {
    Service

    WatchdogService  *WatchdogService
}

func NewVmWatchdogsService(connection *Connection, path string) *VmWatchdogsService {
    var result VmWatchdogsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Adds new watchdog to the virtual machine.
// For example, to add a watchdog to a virtual machine, send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/vms/123/watchdogs
// <watchdog>
//   <action>poweroff</action>
//   <model>i6300esb</model>
// </watchdog>
// ----
// with response body:
// [source,xml]
// ----
// <watchdog href="/ovirt-engine/api/vms/123/watchdogs/00000000-0000-0000-0000-000000000000" id="00000000-0000-0000-0000-000000000000">
//   <vm href="/ovirt-engine/api/vms/123" id="123"/>
//   <action>poweroff</action>
//   <model>i6300esb</model>
// </watchdog>
// ----
// This method supports the following parameters:
// `Watchdog`:: The information about the watchdog.
// The request data must contain `model` element (such as `i6300esb`) and `action` element
// (one of `none`, `reset`, `poweroff`, `dump`, `pause`). The response data additionally
// contains references to the added watchdog and to the virtual machine.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmWatchdogsService) Add (
    watchdog Watchdog,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(watchdog, headers, query, wait)
}

//
// The list of watchdogs of the virtual machine.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of watchdogs to return. If not specified all the watchdogs are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmWatchdogsService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmWatchdogsService) Service(path string) *VmWatchdogsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Watchdog_service(path)
    return self.Watchdog_service(path[:index]).Service(path[index + 1:])
}

func (op *VmWatchdogsService) String() string {
    return fmt.Sprintf("VmWatchdogsService:%s", op.Path)
}


//
//
type VmsService struct {
    Service

    VmService  *VmService
}

func NewVmsService(connection *Connection, path string) *VmsService {
    var result VmsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Creates a new virtual machine.
// The virtual machine can be created in different ways:
// - From a template. In this case the identifier or name of the template must be provided. For example, using a
//   plain shell script and XML:
// [source,bash]
// ----
// #!/bin/sh -ex
// url="https://engine.example.com/ovirt-engine/api"
// user="admin@internal"
// password="..."
// curl \
// --verbose \
// --cacert /etc/pki/ovirt-engine/ca.pem \
// --user "${user}:${password}" \
// --request POST \
// --header "Version: 4" \
// --header "Content-Type: application/xml" \
// --header "Accept: application/xml" \
// --data '
// <vm>
//   <name>myvm</name>
//   <template>
//     <name>Blank</name>
//   </template>
//   <cluster>
//     <name>mycluster</name>
//   </cluster>
// </vm>
// ' \
// "${url}/vms"
// ----
// - From a snapshot. In this case the identifier of the snapshot has to be provided. For example, using a plain
//   shel script and XML:
// [source,bash]
// ----
// #!/bin/sh -ex
// url="https://engine.example.com/ovirt-engine/api"
// user="admin@internal"
// password="..."
// curl \
// --verbose \
// --cacert /etc/pki/ovirt-engine/ca.pem \
// --user "${user}:${password}" \
// --request POST \
// --header "Content-Type: application/xml" \
// --header "Accept: application/xml" \
// --data '
// <vm>
//   <name>myvm</name>
//   <snapshots>
//     <snapshot id="266742a5-6a65-483c-816d-d2ce49746680"/>
//   </snapshots>
//   <cluster>
//     <name>mycluster</name>
//   </cluster>
// </vm>
// ' \
// "${url}/vms"
// ----
// When creating a virtual machine from a template or from a snapshot it is usually useful to explicitly indicate
// in what storage domain to create the disks for the virtual machine. If the virtual machine is created from
// a template then this is achieved passing a set of `disk_attachment` elements that indicate the mapping:
// [source,xml]
// ----
// <vm>
//   ...
//   <disk_attachments>
//     <disk_attachment>
//       <disk id="8d4bd566-6c86-4592-a4a7-912dbf93c298">
//         <storage_domains>
//           <storage_domain id="9cb6cb0a-cf1d-41c2-92ca-5a6d665649c9"/>
//         </storage_domains>
//       </disk>
//     <disk_attachment>
//   </disk_attachments>
// </vm>
// ----
// When the virtual machine is created from a snapshot this set of disks is slightly different, it uses the
// `image_id` attribute instead of `id`.
// [source,xml]
// ----
// <vm>
//   ...
//   <disk_attachments>
//     <disk_attachment>
//       <disk>
//         <image_id>8d4bd566-6c86-4592-a4a7-912dbf93c298</image_id>
//         <storage_domains>
//           <storage_domain id="9cb6cb0a-cf1d-41c2-92ca-5a6d665649c9"/>
//         </storage_domains>
//       </disk>
//     <disk_attachment>
//   </disk_attachments>
// </vm>
// ----
// It is possible to specify additional virtual machine parameters in the XML description, e.g. a virtual machine
// of `desktop` type, with 2 GiB of RAM and additional description can be added sending a request body like the
// following:
// [source,xml]
// ----
// <vm>
//   <name>myvm</name>
//   <description>My Desktop Virtual Machine</description>
//   <type>desktop</type>
//   <memory>2147483648</memory>
//   ...
// </vm>
// ----
// A bootable CDROM device can be set like this:
// [source,xml]
// ----
// <vm>
//   ...
//   <os>
//     <boot dev="cdrom"/>
//   </os>
// </vm>
// ----
// In order to boot from CDROM, you first need to insert a disk, as described in the
// <<services/vm_cdrom, CDROM service>>. Then booting from that CDROM can be specified using the `os.boot.devices`
// attribute:
// [source,xml]
// ----
// <vm>
//   ...
//   <os>
//     <boot>
//       <devices>
//         <device>cdrom</device>
//       </devices>
//     </boot>
//   </os>
// </vm>
// ----
// In all cases the name or identifier of the cluster where the virtual machine will be created is mandatory.
//
func (op *VmsService) Add (
    vm Vm,
    clone bool,
    clonePermissions bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if clone != nil {
        query["clone"] = clone
    }
    if clonePermissions != nil {
        query["clone_permissions"] = clonePermissions
    }

    // Send the request
    return op.internalAdd(vm, headers, query, wait)
}

//
// This method supports the following parameters:
// `Search`:: A query string used to restrict the returned virtual machines.
// `Max`:: The maximum number of results to return.
// `CaseSensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
// account. The default value is `true`, which means that case is taken into account. If you want to search
// ignoring case set it to `false`.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `AllContent`:: Indicates if all the attributes of the virtual machines should be included in the response.
// By default the following attributes are excluded:
// - `console`
// - `initialization.configuration.data` - The OVF document describing the virtual machine.
// - `rng_source`
// - `soundcard`
// - `virtio_scsi`
// For example, to retrieve the complete representation of the virtual machines send a request like this:
// ....
// GET /ovirt-engine/api/vms?all_content=true
// ....
// NOTE: The reason for not including these attributes is performance: they are seldom used and they require
// additional queries to the database. So try to use the this parameter only when it is really needed.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VmsService) List (
    allContent bool,
    caseSensitive bool,
    filter bool,
    max int64,
    search string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if allContent != nil {
        query["all_content"] = allContent
    }
    if caseSensitive != nil {
        query["case_sensitive"] = caseSensitive
    }
    if filter != nil {
        query["filter"] = filter
    }
    if max != nil {
        query["max"] = max
    }
    if search != nil {
        query["search"] = search
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VmsService) Service(path string) *VmsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Vm_service(path)
    return self.Vm_service(path[:index]).Service(path[index + 1:])
}

func (op *VmsService) String() string {
    return fmt.Sprintf("VmsService:%s", op.Path)
}


//
// This service manages a vNIC profile.
//
type VnicProfileService struct {
    Service

    PermissionsService  *PermissionsService
}

func NewVnicProfileService(connection *Connection, path string) *VnicProfileService {
    var result VnicProfileService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieves details about a vNIC profile.
//
func (op *VnicProfileService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes the vNIC profile.
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VnicProfileService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Updates details of a vNIC profile.
// This method supports the following parameters:
// `Profile`:: The vNIC profile that is being updated.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VnicProfileService) Update (
    profile VnicProfile,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(profile, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VnicProfileService) Service(path string) *VnicProfileService {
    if path == nil {
        return op
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *VnicProfileService) String() string {
    return fmt.Sprintf("VnicProfileService:%s", op.Path)
}


//
// This service manages the collection of all vNIC profiles.
//
type VnicProfilesService struct {
    Service

    ProfileService  *ProfileService
}

func NewVnicProfilesService(connection *Connection, path string) *VnicProfilesService {
    var result VnicProfilesService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Add a vNIC profile.
// For example to add vNIC profile `123` to network `456` send a request to:
// [source]
// ----
// POST /ovirt-engine/api/networks/456/vnicprofiles
// ----
// With the following body:
// [source,xml]
// ----
// <vnic_profile id="123">
//   <name>new_vNIC_name</name>
//   <pass_through>
//     <mode>disabled</mode>
//   </pass_through>
//   <port_mirroring>false</port_mirroring>
// </vnic_profile>
// ----
// Please note that there is a default network filter to each VNIC profile.
// For more details of how the default network filter is calculated please refer to
// the documentation in <<services/network_filters,NetworkFilters>>.
// The output of creating a new VNIC profile depends in the  body  arguments that were given.
// In case no network filter was given, the default network filter will be configured. For example:
// [source,xml]
// ----
// <vnic_profile href="/ovirt-engine/api/vnicprofiles/123" id="123">
//   <name>new_vNIC_name</name>
//   <link href="/ovirt-engine/api/vnicprofiles/123/permissions" rel="permissions"/>
//   <pass_through>
//     <mode>disabled</mode>
//   </pass_through>
//   <port_mirroring>false</port_mirroring>
//   <network href="/ovirt-engine/api/networks/456" id="456"/>
//   <network_filter href="/ovirt-engine/api/networkfilters/789" id="789"/>
// </vnic_profile>
// ----
// In case an empty network filter was given, no network filter will be configured for the specific VNIC profile
// regardless of the VNIC profile's default network filter. For example:
// [source,xml]
// ----
// <vnic_profile>
//   <name>no_network_filter</name>
//   <network_filter/>
// </vnic_profile>
// ----
// In case that a specific valid network filter id was given, the VNIC profile will be configured with the given
// network filter regardless of the VNIC profiles's default network filter. For example:
// [source,xml]
// ----
// <vnic_profile>
//   <name>user_choice_network_filter</name>
//   <network_filter id= "0000001b-001b-001b-001b-0000000001d5"/>
// </vnic_profile>
// ----
// This method supports the following parameters:
// `Profile`:: The vNIC profile that is being added.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VnicProfilesService) Add (
    profile VnicProfile,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(profile, headers, query, wait)
}

//
// List all vNIC profiles.
// This method supports the following parameters:
// `Max`:: Sets the maximum number of profiles to return. If not specified all the profiles are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *VnicProfilesService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *VnicProfilesService) Service(path string) *VnicProfilesService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Profile_service(path)
    return self.Profile_service(path[:index]).Service(path[index + 1:])
}

func (op *VnicProfilesService) String() string {
    return fmt.Sprintf("VnicProfilesService:%s", op.Path)
}


//
//
type WeightService struct {
    Service

}

func NewWeightService(connection *Connection, path string) *WeightService {
    var result WeightService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This method supports the following parameters:
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *WeightService) Get (
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *WeightService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *WeightService) Service(path string) *WeightService {
    if path == nil {
        return op
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *WeightService) String() string {
    return fmt.Sprintf("WeightService:%s", op.Path)
}


//
//
type WeightsService struct {
    Service

    WeightService  *WeightService
}

func NewWeightsService(connection *Connection, path string) *WeightsService {
    var result WeightsService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *WeightsService) Add (
    weight Weight,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalAdd(weight, headers, query, wait)
}

//
// This method supports the following parameters:
// `Max`:: Sets the maximum number of weights to return. If not specified all the weights are returned.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *WeightsService) List (
    filter bool,
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *WeightsService) Service(path string) *WeightsService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.Weight_service(path)
    return self.Weight_service(path[:index]).Service(path[index + 1:])
}

func (op *WeightsService) String() string {
    return fmt.Sprintf("WeightsService:%s", op.Path)
}


//
// Manages a single disk available in a storage domain attached to a data center.
// IMPORTANT: Since version 4.2 of the engine this service is intended only to list disks available in the storage
// domain, and to register unregistered disks. All the other operations, like copying a disk, moving a disk, etc, have
// been deprecated and will be removed in the future. To perform those operations use the <<services/disks, service
// that manages all the disks of the system>>, or the <<services/disk, service that manages an specific disk>>.
//
type AttachedStorageDomainDiskService struct {
    MeasurableService

    PermissionsService  *PermissionsService
    StatisticsService  *StatisticsService
}

func NewAttachedStorageDomainDiskService(connection *Connection, path string) *AttachedStorageDomainDiskService {
    var result AttachedStorageDomainDiskService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Copies a disk to the specified storage domain.
// IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
// compatibility. It will be removed in the future. To copy a disk use the <<services/disk/methods/copy, copy>>
// operation of the service that manages that disk.
// This method supports the following parameters:
// `Disk`:: Description of the resulting disk.
// `StorageDomain`:: The storage domain where the new disk will be created.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AttachedStorageDomainDiskService) Copy (
    disk Disk,
    storageDomain StorageDomain,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Disk: disk,
        StorageDomain: storageDomain,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'copy', nil, headers, query, wait)
}

//
// Exports a disk to an export storage domain.
// IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
// compatibility. It will be removed in the future. To export a disk use the <<services/disk/methods/export, export>>
// operation of the service that manages that disk.
// This method supports the following parameters:
// `StorageDomain`:: The export storage domain where the disk should be exported to.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AttachedStorageDomainDiskService) Export (
    storageDomain StorageDomain,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        StorageDomain: storageDomain,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'export', nil, headers, query, wait)
}

//
// Retrieves the description of the disk.
//
func (op *AttachedStorageDomainDiskService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Moves a disk to another storage domain.
// IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
// compatibility. It will be removed in the future. To move a disk use the <<services/disk/methods/move, move>>
// operation of the service that manages that disk.
// This method supports the following parameters:
// `StorageDomain`:: The storage domain where the disk will be moved to.
// `Async`:: Indicates if the move should be performed asynchronously.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AttachedStorageDomainDiskService) Move (
    async bool,
    filter bool,
    storageDomain StorageDomain,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Filter: filter,
        StorageDomain: storageDomain,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'move', nil, headers, query, wait)
}

//
// Registers an unregistered disk.
//
func (op *AttachedStorageDomainDiskService) Register (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
    }

    // Send the request and wait for the response:
    return internalAction(action, 'register', nil, headers, query, wait)
}

//
// Removes a disk.
// IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
// compatibility. It will be removed in the future. To remove a disk use the <<services/disk/methods/remove, remove>>
// operation of the service that manages that disk.
//
func (op *AttachedStorageDomainDiskService) Remove (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Sparsify the disk.
// IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
// compatibility. It will be removed in the future. To remove a disk use the <<services/disk/methods/remove, remove>>
// operation of the service that manages that disk.
//
func (op *AttachedStorageDomainDiskService) Sparsify (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
    }

    // Send the request and wait for the response:
    return internalAction(action, 'sparsify', nil, headers, query, wait)
}

//
// Updates the disk.
// IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
// compatibility. It will be removed in the future. To update a disk use the
// <<services/disk/methods/update, update>> operation of the service that manages that disk.
// This method supports the following parameters:
// `Disk`:: The update to apply to the disk.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *AttachedStorageDomainDiskService) Update (
    disk Disk,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalUpdate(disk, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *AttachedStorageDomainDiskService) Service(path string) *AttachedStorageDomainDiskService {
    if path == nil {
        return op
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    if path == "statistics" {
        return &StatisticsService{}
    }
    if path.startswith('statistics/') {
        return StatisticsService().service(path[11:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *AttachedStorageDomainDiskService) String() string {
    return fmt.Sprintf("AttachedStorageDomainDiskService:%s", op.Path)
}


//
// Manages a single disk.
//
type DiskService struct {
    MeasurableService

    PermissionsService  *PermissionsService
    StatisticsService  *StatisticsService
}

func NewDiskService(connection *Connection, path string) *DiskService {
    var result DiskService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// This operation copies a disk to the specified storage domain.
// For example, copy of a disk can be facilitated using the following request:
// [source]
// ----
// POST /ovirt-engine/api/disks/123/copy
// ----
// With a request body like this:
// [source,xml]
// ----
// <action>
//   <storage_domain id="456"/>
//   <disk>
//     <name>mydisk</name>
//   </disk>
// </action>
// ----
// This method supports the following parameters:
// `Disk`:: Description of the resulting disk. The only accepted value is the `name` attribute, which will be the name
// used for the new disk. For example, to copy disk `123` using `myname` as the name for the new disk, send
// a request like this:
// ....
// POST /ovirt-engine/disks/123
// ....
// With a request body like this:
// [source,xml]
// ----
// <action>
//   <disk>
//     <name>mydisk<name>
//   </disk>
//   <storage_domain id="456"/>
// </action>
// ----
// `StorageDomain`:: The storage domain where the new disk will be created. Can be specified using the `id` or `name`
// attributes. For example, to copy a disk to the storage domain named `mydata` send a request like this:
// ....
// POST /ovirt-engine/api/storagedomains/123/disks/789
// ....
// With a request body like this:
// [source,xml]
// ----
// <action>
//   <storage_domain>
//     <name>mydata</name>
//   </storage_domain>
// </action>
// ----
// `Async`:: Indicates if the copy should be performed asynchronously.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DiskService) Copy (
    async bool,
    disk Disk,
    filter bool,
    storageDomain StorageDomain,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Disk: disk,
        Filter: filter,
        StorageDomain: storageDomain,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'copy', nil, headers, query, wait)
}

//
// Exports a disk to an export storage domain.
// This method supports the following parameters:
// `StorageDomain`:: The export storage domain where the disk should be exported to.
// `Async`:: Indicates if the export should be performed asynchronously.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DiskService) Export (
    async bool,
    filter bool,
    storageDomain StorageDomain,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Filter: filter,
        StorageDomain: storageDomain,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'export', nil, headers, query, wait)
}

//
// Retrieves the description of the disk.
//
func (op *DiskService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Moves a disk to another storage domain.
// For example, to move the disk with identifier `123` to a storage domain with identifier `456` send the following
// request:
// [source]
// ----
// POST /ovirt-engine/api/disks/123/move
// ----
// With the following request body:
// [source,xml]
// ----
// <action>
//   <storage_domain id="456"/>
// </action>
// ----
// This method supports the following parameters:
// `StorageDomain`:: The storage domain where the disk will be moved to.
// `Async`:: Indicates if the move should be performed asynchronously.
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DiskService) Move (
    async bool,
    filter bool,
    storageDomain StorageDomain,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Filter: filter,
        StorageDomain: storageDomain,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'move', nil, headers, query, wait)
}

//
// Removes a disk.
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DiskService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Sparsify the disk.
// Sparsification frees space in the disk image that is not used by its
// filesystem. As a result, the image will occupy less space on the storage.
// Currently sparsification works only on disks without snapshots. Disks
// having derived disks are also not allowed.
//
func (op *DiskService) Sparsify (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
    }

    // Send the request and wait for the response:
    return internalAction(action, 'sparsify', nil, headers, query, wait)
}

//
// This operation updates the disk with the appropriate parameters.
// The only field that can be updated is `qcow_version`.
// For example, update disk can be facilitated using the following request:
// [source]
// ----
// PUT /ovirt-engine/api/disks/123
// ----
// With a request body like this:
// [source,xml]
// ----
// <disk>
//   <qcow_version>qcow2_v3</qcow_version>
// </disk>
// ----
// Since the backend operation is asynchronous the disk element which will be returned
// to the user might not be synced with the changed properties.
// This method supports the following parameters:
// `Disk`:: The update to apply to the disk.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *DiskService) Update (
    disk Disk,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request
    return op.internalUpdate(disk, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *DiskService) Service(path string) *DiskService {
    if path == nil {
        return op
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    if path == "statistics" {
        return &StatisticsService{}
    }
    if path.startswith('statistics/') {
        return StatisticsService().service(path[11:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *DiskService) String() string {
    return fmt.Sprintf("DiskService:%s", op.Path)
}


//
// A service to manage Katello errata assigned to the engine.
// The information is retrieved from Katello.
//
type EngineKatelloErrataService struct {
    KatelloErrataService

    KatelloErratumService  *KatelloErratumService
}

func NewEngineKatelloErrataService(connection *Connection, path string) *EngineKatelloErrataService {
    var result EngineKatelloErrataService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Retrieves the representation of the Katello errata.
// [source]
// ----
// GET /ovirt-engine/api/katelloerrata
// ----
// You will receive response in XML like this one:
// [source,xml]
// ----
// <katello_errata>
//   <katello_erratum href="/ovirt-engine/api/katelloerrata/123" id="123">
//     <name>RHBA-2013:XYZ</name>
//     <description>The description of the erratum</description>
//     <title>some bug fix update</title>
//     <type>bugfix</type>
//     <issued>2013-11-20T02:00:00.000+02:00</issued>
//     <solution>Few guidelines regarding the solution</solution>
//     <summary>Updated packages that fix one bug are now available for XYZ</summary>
//     <packages>
//       <package>
//         <name>libipa_hbac-1.9.2-82.11.el6_4.i686</name>
//       </package>
//       ...
//     </packages>
//   </katello_erratum>
//   ...
// </katello_errata>
// ----
// This method supports the following parameters:
// `Max`:: Sets the maximum number of errata to return. If not specified all the errata are returned.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *EngineKatelloErrataService) List (
    max int64,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if max != nil {
        query["max"] = max
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *EngineKatelloErrataService) Service(path string) *EngineKatelloErrataService {
    if path == nil {
        return op
    }
    index = path.find('/')
    if index == -1:
        return self.KatelloErratum_service(path)
    return self.KatelloErratum_service(path[:index]).Service(path[index + 1:])
}

func (op *EngineKatelloErrataService) String() string {
    return fmt.Sprintf("EngineKatelloErrataService:%s", op.Path)
}


//
//
type ExternalHostProviderService struct {
    ExternalProviderService

    CertificatesService  *CertificatesService
    ComputeResourcesService  *ComputeResourcesService
    DiscoveredHostsService  *DiscoveredHostsService
    HostGroupsService  *HostGroupsService
    HostsService  *HostsService
}

func NewExternalHostProviderService(connection *Connection, path string) *ExternalHostProviderService {
    var result ExternalHostProviderService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *ExternalHostProviderService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
//
func (op *ExternalHostProviderService) ImportCertificates (
    certificates []*Certificate,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Certificates: certificates,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'importcertificates', nil, headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *ExternalHostProviderService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the test should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *ExternalHostProviderService) TestConnectivity (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'testconnectivity', nil, headers, query, wait)
}

//
//
func (op *ExternalHostProviderService) Update (
    provider ExternalHostProvider,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(provider, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *ExternalHostProviderService) Service(path string) *ExternalHostProviderService {
    if path == nil {
        return op
    }
    if path == "certificates" {
        return &CertificatesService{}
    }
    if path.startswith('certificates/') {
        return CertificatesService().service(path[13:])
    }
    if path == "computeresources" {
        return &ComputeResourcesService{}
    }
    if path.startswith('computeresources/') {
        return ComputeResourcesService().service(path[17:])
    }
    if path == "discoveredhosts" {
        return &DiscoveredHostsService{}
    }
    if path.startswith('discoveredhosts/') {
        return DiscoveredHostsService().service(path[16:])
    }
    if path == "hostgroups" {
        return &HostGroupsService{}
    }
    if path.startswith('hostgroups/') {
        return HostGroupsService().service(path[11:])
    }
    if path == "hosts" {
        return &HostsService{}
    }
    if path.startswith('hosts/') {
        return HostsService().service(path[6:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *ExternalHostProviderService) String() string {
    return fmt.Sprintf("ExternalHostProviderService:%s", op.Path)
}


//
// This service manages a single gluster brick.
//
type GlusterBrickService struct {
    MeasurableService

    StatisticsService  *StatisticsService
}

func NewGlusterBrickService(connection *Connection, path string) *GlusterBrickService {
    var result GlusterBrickService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Get details of a brick.
// Retrieves status details of brick from underlying gluster volume with header `All-Content` set to `true`. This is
// the equivalent of running `gluster volume status <volumename> <brickname> detail`.
// For example, to get the details of brick `234` of gluster volume `123`, send a request like this:
// [source]
// ----
// GET /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks/234
// ----
// Which will return a response body like this:
// [source,xml]
// ----
// <brick id="234">
//   <name>host1:/rhgs/data/brick1</name>
//   <brick_dir>/rhgs/data/brick1</brick_dir>
//   <server_id>111</server_id>
//   <status>up</status>
//   <device>/dev/mapper/RHGS_vg1-lv_vmaddldisks</device>
//   <fs_name>xfs</fs_name>
//   <gluster_clients>
//     <gluster_client>
//       <bytes_read>2818417648</bytes_read>
//       <bytes_written>1384694844</bytes_written>
//       <client_port>1011</client_port>
//       <host_name>client2</host_name>
//     </gluster_client>
//   </gluster_clients>
//   <memory_pools>
//     <memory_pool>
//       <name>data-server:fd_t</name>
//       <alloc_count>1626348</alloc_count>
//       <cold_count>1020</cold_count>
//       <hot_count>4</hot_count>
//       <max_alloc>23</max_alloc>
//       <max_stdalloc>0</max_stdalloc>
//       <padded_size>140</padded_size>
//       <pool_misses>0</pool_misses>
//     </memory_pool>
//   </memory_pools>
//   <mnt_options>rw,seclabel,noatime,nodiratime,attr2,inode64,sunit=512,swidth=2048,noquota</mnt_options>
//   <pid>25589</pid>
//   <port>49155</port>
// </brick>
// ----
//
func (op *GlusterBrickService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Removes a brick.
// Removes a brick from the underlying gluster volume and deletes entries from database. This can be used only when
// removing a single brick without data migration. To remove multiple bricks and with data migration, use
// <<services/gluster_bricks/methods/migrate, migrate>> instead.
// For example, to delete brick `234` from gluster volume `123`, send a request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks/234
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterBrickService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Replaces this brick with a new one.
// IMPORTANT: This operation has been deprecated since version 3.5 of the engine and will be removed in the future.
// Use <<services/gluster_bricks/methods/add, add brick(s)>> and
// <<services/gluster_bricks/methods/migrate, migrate brick(s)>> instead.
// This method supports the following parameters:
// `Async`:: Indicates if the replacement should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterBrickService) Replace (
    async bool,
    force bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Force: force,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'replace', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *GlusterBrickService) Service(path string) *GlusterBrickService {
    if path == nil {
        return op
    }
    if path == "statistics" {
        return &StatisticsService{}
    }
    if path.startswith('statistics/') {
        return StatisticsService().service(path[11:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *GlusterBrickService) String() string {
    return fmt.Sprintf("GlusterBrickService:%s", op.Path)
}


//
// This service manages a single gluster volume.
//
type GlusterVolumeService struct {
    MeasurableService

    GlusterBricksService  *GlusterBricksService
    StatisticsService  *StatisticsService
}

func NewGlusterVolumeService(connection *Connection, path string) *GlusterVolumeService {
    var result GlusterVolumeService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Get the gluster volume details.
// For example, to get details of a gluster volume with identifier `123` in cluster `456`, send a request like this:
// [source]
// ----
// GET /ovirt-engine/api/clusters/456/glustervolumes/123
// ----
// This GET request will return the following output:
// [source,xml]
// ----
// <gluster_volume id="123">
//  <name>data</name>
//  <link href="/ovirt-engine/api/clusters/456/glustervolumes/123/glusterbricks" rel="glusterbricks"/>
//  <disperse_count>0</disperse_count>
//  <options>
//    <option>
//      <name>storage.owner-gid</name>
//      <value>36</value>
//    </option>
//    <option>
//      <name>performance.io-cache</name>
//      <value>off</value>
//    </option>
//    <option>
//      <name>cluster.data-self-heal-algorithm</name>
//      <value>full</value>
//    </option>
//  </options>
//  <redundancy_count>0</redundancy_count>
//  <replica_count>3</replica_count>
//  <status>up</status>
//  <stripe_count>0</stripe_count>
//  <transport_types>
//    <transport_type>tcp</transport_type>
//  </transport_types>
//  <volume_type>replicate</volume_type>
//  </gluster_volume>
// ----
//
func (op *GlusterVolumeService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Get gluster volume profile statistics.
// For example, to get profile statistics for a gluster volume with identifier `123` in cluster `456`, send a
// request like this:
// [source]
// ----
// POST /ovirt-engine/api/clusters/456/glustervolumes/123/getprofilestatistics
// ----
//
func (op *GlusterVolumeService) GetProfileStatistics (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
    }

    // Send the request and wait for the response:
    return internalAction(action, 'getprofilestatistics', 'details', headers, query, wait)
}

//
// Rebalance the gluster volume.
// Rebalancing a gluster volume helps to distribute the data evenly across all the bricks. After expanding or
// shrinking a gluster volume (without migrating data), we need to rebalance the data among the bricks. In a
// non-replicated volume, all bricks should be online to perform the rebalance operation. In a replicated volume, at
// least one of the bricks in the replica should be online.
// For example, to rebalance a gluster volume with identifier `123` in cluster `456`, send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/clusters/456/glustervolumes/123/rebalance
// ----
// This method supports the following parameters:
// `FixLayout`:: If set to true, rebalance will only fix the layout so that new data added to the volume is distributed
// across all the hosts. But it will not migrate/rebalance the existing data. Default is `false`.
// `Force`:: Indicates if the rebalance should be force started. The rebalance command can be executed with the force
// option even when the older clients are connected to the cluster. However, this could lead to a data loss
// situation. Default is `false`.
// `Async`:: Indicates if the rebalance should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterVolumeService) Rebalance (
    async bool,
    fixLayout bool,
    force bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        FixLayout: fixLayout,
        Force: force,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'rebalance', nil, headers, query, wait)
}

//
// Removes the gluster volume.
// For example, to remove a volume with identifier `123` in cluster `456`, send a request like this:
// [source]
// ----
// DELETE /ovirt-engine/api/clusters/456/glustervolumes/123
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterVolumeService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// Resets all the options set in the gluster volume.
// For example, to reset all options in a gluster volume with identifier `123` in cluster `456`, send a request like
// this:
// [source]
// ----
// POST /ovirt-engine/api/clusters/456/glustervolumes/123/resetalloptions
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the reset should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterVolumeService) ResetAllOptions (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'resetalloptions', nil, headers, query, wait)
}

//
// Resets a particular option in the gluster volume.
// For example, to reset a particular option `option1` in a gluster volume with identifier `123` in cluster `456`,
// send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/clusters/456/glustervolumes/123/resetoption
// ----
// With the following request body:
// [source,xml]
// ----
// <action>
//  <option name="option1"/>
// </action>
// ----
// This method supports the following parameters:
// `Option`:: Option to reset.
// `Async`:: Indicates if the reset should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterVolumeService) ResetOption (
    async bool,
    force bool,
    option Option,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Force: force,
        Option: option,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'resetoption', nil, headers, query, wait)
}

//
// Sets a particular option in the gluster volume.
// For example, to set `option1` with value `value1` in a gluster volume with identifier `123` in cluster `456`,
// send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/clusters/456/glustervolumes/123/setoption
// ----
// With the following request body:
// [source,xml]
// ----
// <action>
//  <option name="option1" value="value1"/>
// </action>
// ----
// This method supports the following parameters:
// `Option`:: Option to set.
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterVolumeService) SetOption (
    async bool,
    option Option,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Option: option,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'setoption', nil, headers, query, wait)
}

//
// Starts the gluster volume.
// A Gluster Volume should be started to read/write data. For example, to start a gluster volume with identifier
// `123` in cluster `456`, send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/clusters/456/glustervolumes/123/start
// ----
// This method supports the following parameters:
// `Force`:: Indicates if the volume should be force started. If a gluster volume is started already but few/all bricks
// are down then force start can be used to bring all the bricks up. Default is `false`.
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterVolumeService) Start (
    async bool,
    force bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Force: force,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'start', nil, headers, query, wait)
}

//
// Start profiling the gluster volume.
// For example, to start profiling a gluster volume with identifier `123` in cluster `456`, send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/clusters/456/glustervolumes/123/startprofile
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterVolumeService) StartProfile (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'startprofile', nil, headers, query, wait)
}

//
// Stops the gluster volume.
// Stopping a volume will make its data inaccessible.
// For example, to stop a gluster volume with identifier `123` in cluster `456`, send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/clusters/456/glustervolumes/123/stop
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterVolumeService) Stop (
    async bool,
    force bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Force: force,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'stop', nil, headers, query, wait)
}

//
// Stop profiling the gluster volume.
// For example, to stop profiling a gluster volume with identifier `123` in cluster `456`, send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/clusters/456/glustervolumes/123/stopprofile
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterVolumeService) StopProfile (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'stopprofile', nil, headers, query, wait)
}

//
// Stop rebalancing the gluster volume.
// For example, to stop rebalancing a gluster volume with identifier `123` in cluster `456`, send a request like
// this:
// [source]
// ----
// POST /ovirt-engine/api/clusters/456/glustervolumes/123/stoprebalance
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *GlusterVolumeService) StopRebalance (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'stoprebalance', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *GlusterVolumeService) Service(path string) *GlusterVolumeService {
    if path == nil {
        return op
    }
    if path == "glusterbricks" {
        return &GlusterBricksService{}
    }
    if path.startswith('glusterbricks/') {
        return GlusterBricksService().service(path[14:])
    }
    if path == "statistics" {
        return &StatisticsService{}
    }
    if path.startswith('statistics/') {
        return StatisticsService().service(path[11:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *GlusterVolumeService) String() string {
    return fmt.Sprintf("GlusterVolumeService:%s", op.Path)
}


//
// A service to manage a host.
//
type HostService struct {
    MeasurableService

    AffinityLabelsService  *AffinityLabelsService
    DevicesService  *DevicesService
    FenceAgentsService  *FenceAgentsService
    HooksService  *HooksService
    KatelloErrataService  *KatelloErrataService
    NetworkAttachmentsService  *NetworkAttachmentsService
    NicsService  *NicsService
    NumaNodesService  *NumaNodesService
    PermissionsService  *PermissionsService
    StatisticsService  *StatisticsService
    StorageService  *StorageService
    StorageConnectionExtensionsService  *StorageConnectionExtensionsService
    TagsService  *TagsService
    UnmanagedNetworksService  *UnmanagedNetworksService
}

func NewHostService(connection *Connection, path string) *HostService {
    var result HostService
    result.Connection = connection
    result.Path = path
    return &result
}

//
// Activate the host for use, such as running virtual machines.
// This method supports the following parameters:
// `Async`:: Indicates if the activation should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) Activate (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'activate', nil, headers, query, wait)
}

//
// Approve a pre-installed Hypervisor host for usage in the virtualization environment.
// This action also accepts an optional cluster element to define the target cluster for this host.
// This method supports the following parameters:
// `Async`:: Indicates if the approval should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) Approve (
    async bool,
    cluster Cluster,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Cluster: cluster,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'approve', nil, headers, query, wait)
}

//
// Marks the network configuration as good and persists it inside the host.
// An API user commits the network configuration to persist a host network interface attachment or detachment, or
// persist the creation and deletion of a bonded interface.
// IMPORTANT: Networking configuration is only committed after the engine has established that host connectivity is
// not lost as a result of the configuration changes. If host connectivity is lost, the host requires a reboot and
// automatically reverts to the previous networking configuration.
// For example, to commit the network configuration of host with id `123` send a request like this:
// [source]
// ----
// POST /ovirt-engine/api/hosts/123/commitnetconfig
// ----
// With a request body like this:
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) CommitNetConfig (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'commitnetconfig', nil, headers, query, wait)
}

//
// Deactivate the host to perform maintenance tasks.
// This method supports the following parameters:
// `Async`:: Indicates if the deactivation should be performed asynchronously.
// `StopGlusterService`:: Indicates if the gluster service should be stopped as part of deactivating the host. It can be used while
// performing maintenance operations on the gluster host. Default value for this variable is `false`.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) Deactivate (
    async bool,
    reason string,
    stopGlusterService bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Reason: reason,
        StopGlusterService: stopGlusterService,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'deactivate', nil, headers, query, wait)
}

//
// Enroll certificate of the host. Useful in case you get a warning that it is about to, or already expired.
// This method supports the following parameters:
// `Async`:: Indicates if the enrollment should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) EnrollCertificate (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'enrollcertificate', nil, headers, query, wait)
}

//
// Controls host's power management device.
// For example, let's assume you want to start the host. This can be done via:
// [source]
// ----
// #!/bin/sh -ex
// url="https://engine.example.com/ovirt-engine/api"
// user="admin@internal"
// password="..."
// curl \
// --verbose \
// --cacert /etc/pki/ovirt-engine/ca.pem \
// --user "${user}:${password}" \
// --request POST \
// --header "Version: 4" \
// --header "Content-Type: application/xml" \
// --header "Accept: application/xml" \
// --data '
// <action>
//   <fence_type>start</fence_type>
// </action>
// ' \
// "${url}/hosts/123/fence"
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the fencing should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) Fence (
    async bool,
    fenceType string,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        FenceType: fenceType,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'fence', 'powerManagement', headers, query, wait)
}

//
// Manually set a host as the storage pool manager (SPM).
// [source]
// ----
// POST /ovirt-engine/api/hosts/123/forceselectspm
// ----
// With a request body like this:
// [source,xml]
// ----
// <action/>
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) ForceSelectSpm (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'forceselectspm', nil, headers, query, wait)
}

//
// Get the host details.
// This method supports the following parameters:
// `Filter`:: Indicates if the results should be filtered according to the permissions of the user.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) Get (
    filter bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if filter != nil {
        query["filter"] = filter
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Install VDSM and related software on the host. The host type defines additional parameters for the action.
// Example of installing a host, using `curl` and JSON, plain:
// [source,bash]
// ----
// curl \
// --verbose \
// --cacert /etc/pki/ovirt-engine/ca.pem \
// --request PUT \
// --header "Content-Type: application/json" \
// --header "Accept: application/json" \
// --header "Version: 4" \
// --user "admin@internal:..." \
// --data '
// {
//   "root_password": "myrootpassword"
// }
// ' \
// "https://engine.example.com/ovirt-engine/api/hosts/123"
// ----
// Example of installing a host, using `curl` and JSON, with hosted engine components:
// [source,bash]
// ----
// curl \
// curl \
// --verbose \
// --cacert /etc/pki/ovirt-engine/ca.pem \
// --request PUT \
// --header "Content-Type: application/json" \
// --header "Accept: application/json" \
// --header "Version: 4" \
// --user "admin@internal:..." \
// --data '
// {
//   "root_password": "myrootpassword"
// }
// ' \
// "https://engine.example.com/ovirt-engine/api/hosts/123?deploy_hosted_engine=true"
// ----
// This method supports the following parameters:
// `RootPassword`:: The password of of the `root` user, used to connect to the host via SSH.
// `Ssh`:: The SSH details used to connect to the host.
// `Host`:: This `override_iptables` property is used to indicate if the firewall configuration should be
// replaced by the default one.
// `Image`:: When installing an oVirt node a image ISO file is needed.
// `Async`:: Indicates if the installation should be performed asynchronously.
// `DeployHostedEngine`:: When set to `true` it means this host should deploy also hosted
// engine components. Missing value is treated as `true` i.e deploy.
// Omitting this parameter means `false` and will perform no operation
// in hosted engine area.
// `UndeployHostedEngine`:: When set to `true` it means this host should un-deploy hosted engine
// components and this host will not function as part of the High
// Availability cluster. Missing value is treated as `true` i.e un-deploy
// Omitting this parameter means `false` and will perform no operation
// in hosted engine area.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) Install (
    async bool,
    deployHostedEngine bool,
    host Host,
    image string,
    rootPassword string,
    ssh Ssh,
    undeployHostedEngine bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        DeployHostedEngine: deployHostedEngine,
        Host: host,
        Image: image,
        RootPassword: rootPassword,
        Ssh: ssh,
        UndeployHostedEngine: undeployHostedEngine,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'install', nil, headers, query, wait)
}

//
// Discover iSCSI targets on the host, using the initiator details.
// This method supports the following parameters:
// `Iscsi`:: The target iSCSI device.
// `Async`:: Indicates if the discovery should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) IscsiDiscover (
    async bool,
    iscsi IscsiDetails,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Iscsi: iscsi,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'iscsidiscover', 'iscsiTargets', headers, query, wait)
}

//
// Login to iSCSI targets on the host, using the target details.
// This method supports the following parameters:
// `Iscsi`:: The target iSCSI device.
// `Async`:: Indicates if the login should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) IscsiLogin (
    async bool,
    iscsi IscsiDetails,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Iscsi: iscsi,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'iscsilogin', nil, headers, query, wait)
}

//
// Refresh the host devices and capabilities.
// This method supports the following parameters:
// `Async`:: Indicates if the refresh should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) Refresh (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'refresh', nil, headers, query, wait)
}

//
// Remove the host from the system.
// [source]
// ----
// #!/bin/sh -ex
// url="https://engine.example.com/ovirt-engine/api"
// user="admin@internal"
// password="..."
// curl \
// --verbose \
// --cacert /etc/pki/ovirt-engine/ca.pem \
// --user "${user}:${password}" \
// --request DELETE \
// --header "Version: 4" \
// "${url}/hosts/1ff7a191-2f3b-4eff-812b-9f91a30c3acc"
// ----
// This method supports the following parameters:
// `Async`:: Indicates if the remove should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) Remove (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request and wait for the response:
    op.internalRemove(headers, query, wait)
}

//
// This method is used to change the configuration of the network interfaces of a host.
// For example, lets assume that you have a host with three network interfaces `eth0`, `eth1` and `eth2` and that
// you want to configure a new bond using `eth0` and `eth1`, and put a VLAN on top of it. Using a simple shell
// script and the `curl` command line HTTP client that can be done as follows:
// [source]
// ----
// #!/bin/sh -ex
// url="https://engine.example.com/ovirt-engine/api"
// user="admin@internal"
// password="..."
// curl \
// --verbose \
// --cacert /etc/pki/ovirt-engine/ca.pem \
// --user "${user}:${password}" \
// --request POST \
// --header "Version: 4" \
// --header "Content-Type: application/xml" \
// --header "Accept: application/xml" \
// --data '
// <action>
//   <modified_bonds>
//     <host_nic>
//       <name>bond0</name>
//       <bonding>
//         <options>
//           <option>
//             <name>mode</name>
//             <value>4</value>
//           </option>
//           <option>
//             <name>miimon</name>
//             <value>100</value>
//           </option>
//         </options>
//         <slaves>
//           <host_nic>
//             <name>eth1</name>
//           </host_nic>
//           <host_nic>
//             <name>eth2</name>
//           </host_nic>
//         </slaves>
//       </bonding>
//     </host_nic>
//   </modified_bonds>
//   <modified_network_attachments>
//     <network_attachment>
//       <network>
//         <name>myvlan</name>
//       </network>
//       <host_nic>
//         <name>bond0</name>
//       </host_nic>
//       <ip_address_assignments>
//         <assignment_method>static</assignment_method>
//         <ip_address_assignment>
//           <ip>
//             <address>192.168.122.10</address>
//             <netmask>255.255.255.0</netmask>
//           </ip>
//         </ip_address_assignment>
//       </ip_address_assignments>
//       <dns_resolver_configuration>
//         <name_servers>
//           <name_server>1.1.1.1</name_server>
//           <name_server>2.2.2.2</name_server>
//         </name_servers>
//       </dns_resolver_configuration>
//     </network_attachment>
//   </modified_network_attachments>
//  </action>
// ' \
// "${url}/hosts/1ff7a191-2f3b-4eff-812b-9f91a30c3acc/setupnetworks"
// ----
// Note that this is valid for version 4 of the API. In previous versions some elements were represented as XML
// attributes instead of XML elements. In particular the `options` and `ip` elements were represented as follows:
// [source,xml]
// ----
// <options name="mode" value="4"/>
// <options name="miimon" value="100"/>
// <ip address="192.168.122.10" netmask="255.255.255.0"/>
// ----
// Using the Python SDK the same can be done with the following code:
// [source,python]
// ----
// # Find the service that manages the collection of hosts:
// hosts_service = connection.system_service().hosts_service()
// # Find the host:
// host = hosts_service.list(search='name=myhost')[0]
// # Find the service that manages the host:
// host_service = hosts_service.host_service(host.id)
// # Configure the network adding a bond with two slaves and attaching it to a
// # network with an static IP address:
// host_service.setup_networks(
//     modified_bonds=[
//         types.HostNic(
//             name='bond0',
//             bonding=types.Bonding(
//                 options=[
//                     types.Option(
//                         name='mode',
//                         value='4',
//                     ),
//                     types.Option(
//                         name='miimon',
//                         value='100',
//                     ),
//                 ],
//                 slaves=[
//                     types.HostNic(
//                         name='eth1',
//                     ),
//                     types.HostNic(
//                         name='eth2',
//                     ),
//                 ],
//             ),
//         ),
//     ],
//     modified_network_attachments=[
//         types.NetworkAttachment(
//             network=types.Network(
//                 name='myvlan',
//             ),
//             host_nic=types.HostNic(
//                 name='bond0',
//             ),
//             ip_address_assignments=[
//                 types.IpAddressAssignment(
//                     assignment_method=types.BootProtocol.STATIC,
//                     ip=types.Ip(
//                         address='192.168.122.10',
//                         netmask='255.255.255.0',
//                     ),
//                 ),
//             ],
//             dns_resolver_configuration=types.DnsResolverConfiguration(
//                 name_servers=[
//                     '1.1.1.1',
//                     '2.2.2.2',
//                 ],
//             ),
//         ),
//     ],
// )
// # After modifying the network configuration it is very important to make it
// # persistent:
// host_service.commit_net_config()
// ----
// IMPORTANT: To make sure that the network configuration has been saved in the host, and that it will be applied
// when the host is rebooted, remember to call <<services/host/methods/commit_net_config, commitnetconfig>>.
// This method supports the following parameters:
// `Async`:: Indicates if the action should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) SetupNetworks (
    async bool,
    checkConnectivity bool,
    connectivityTimeout int64,
    modifiedBonds []*HostNic,
    modifiedLabels []*NetworkLabel,
    modifiedNetworkAttachments []*NetworkAttachment,
    removedBonds []*HostNic,
    removedLabels []*NetworkLabel,
    removedNetworkAttachments []*NetworkAttachment,
    synchronizedNetworkAttachments []*NetworkAttachment,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        CheckConnectivity: checkConnectivity,
        ConnectivityTimeout: connectivityTimeout,
        ModifiedBonds: modifiedBonds,
        ModifiedLabels: modifiedLabels,
        ModifiedNetworkAttachments: modifiedNetworkAttachments,
        RemovedBonds: removedBonds,
        RemovedLabels: removedLabels,
        RemovedNetworkAttachments: removedNetworkAttachments,
        SynchronizedNetworkAttachments: synchronizedNetworkAttachments,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'setupnetworks', nil, headers, query, wait)
}

//
// This method supports the following parameters:
// `Async`:: Indicates if the discovery should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) UnregisteredStorageDomainsDiscover (
    async bool,
    iscsi IscsiDetails,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        Iscsi: iscsi,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'unregisteredstoragedomainsdiscover', 'storageDomains', headers, query, wait)
}

//
// Update the host properties.
// For example, to update a the kernel command line of a host send a request like this:
// [source]
// ----
// PUT /ovirt-engine/api/hosts/123
// ----
// With request body like this:
// [source, xml]
// ----
// <host>
//   <os>
//     <custom_kernel_cmdline>vfio_iommu_type1.allow_unsafe_interrupts=1</custom_kernel_cmdline>
//   </os>
// </host>
// ----
//
func (op *HostService) Update (
    host Host,
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }
    if async != nil {
        query["async"] = async
    }

    // Send the request
    return op.internalUpdate(host, headers, query, wait)
}

//
// Upgrade VDSM and selected software on the host.
// This method supports the following parameters:
// `Async`:: Indicates if the upgrade should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostService) Upgrade (
    async bool,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'upgrade', nil, headers, query, wait)
}

//
// Check if there are upgrades available for the host. If there are upgrades
// available an icon will be displayed next to host status icon in the webadmin.
// Audit log messages are also added to indicate the availability of upgrades.
// The upgrade can be started from the webadmin or by using the
// <<services/host/methods/upgrade, upgrade>> host action.
//
func (op *HostService) UpgradeCheck (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
    }

    // Send the request and wait for the response:
    return internalAction(action, 'upgradecheck', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *HostService) Service(path string) *HostService {
    if path == nil {
        return op
    }
    if path == "affinitylabels" {
        return &AffinityLabelsService{}
    }
    if path.startswith('affinitylabels/') {
        return AffinityLabelsService().service(path[15:])
    }
    if path == "devices" {
        return &DevicesService{}
    }
    if path.startswith('devices/') {
        return DevicesService().service(path[8:])
    }
    if path == "fenceagents" {
        return &FenceAgentsService{}
    }
    if path.startswith('fenceagents/') {
        return FenceAgentsService().service(path[12:])
    }
    if path == "hooks" {
        return &HooksService{}
    }
    if path.startswith('hooks/') {
        return HooksService().service(path[6:])
    }
    if path == "katelloerrata" {
        return &KatelloErrataService{}
    }
    if path.startswith('katelloerrata/') {
        return KatelloErrataService().service(path[14:])
    }
    if path == "networkattachments" {
        return &NetworkAttachmentsService{}
    }
    if path.startswith('networkattachments/') {
        return NetworkAttachmentsService().service(path[19:])
    }
    if path == "nics" {
        return &NicsService{}
    }
    if path.startswith('nics/') {
        return NicsService().service(path[5:])
    }
    if path == "numanodes" {
        return &NumaNodesService{}
    }
    if path.startswith('numanodes/') {
        return NumaNodesService().service(path[10:])
    }
    if path == "permissions" {
        return &PermissionsService{}
    }
    if path.startswith('permissions/') {
        return PermissionsService().service(path[12:])
    }
    if path == "statistics" {
        return &StatisticsService{}
    }
    if path.startswith('statistics/') {
        return StatisticsService().service(path[11:])
    }
    if path == "storage" {
        return &StorageService{}
    }
    if path.startswith('storage/') {
        return StorageService().service(path[8:])
    }
    if path == "storageconnectionextensions" {
        return &StorageConnectionExtensionsService{}
    }
    if path.startswith('storageconnectionextensions/') {
        return StorageConnectionExtensionsService().service(path[28:])
    }
    if path == "tags" {
        return &TagsService{}
    }
    if path.startswith('tags/') {
        return TagsService().service(path[5:])
    }
    if path == "unmanagednetworks" {
        return &UnmanagedNetworksService{}
    }
    if path.startswith('unmanagednetworks/') {
        return UnmanagedNetworksService().service(path[18:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *HostService) String() string {
    return fmt.Sprintf("HostService:%s", op.Path)
}


//
// A service to manage a network interface of a host.
//
type HostNicService struct {
    MeasurableService

    NetworkAttachmentsService  *NetworkAttachmentsService
    NetworkLabelsService  *NetworkLabelsService
    StatisticsService  *StatisticsService
    VirtualFunctionAllowedLabelsService  *VirtualFunctionAllowedLabelsService
    VirtualFunctionAllowedNetworksService  *VirtualFunctionAllowedNetworksService
}

func NewHostNicService(connection *Connection, path string) *HostNicService {
    var result HostNicService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *HostNicService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// The action updates virtual function configuration in case the current resource represents an SR-IOV enabled NIC.
// The input should be consisted of at least one of the following properties:
// - `allNetworksAllowed`
// - `numberOfVirtualFunctions`
// Please see the `HostNicVirtualFunctionsConfiguration` type for the meaning of the properties.
// This method supports the following parameters:
// `Async`:: Indicates if the update should be performed asynchronously.
// `headers`:: Additional HTTP headers.
// `query`:: Additional URL query parameters.
// `wait`:: If `True` wait for the response.
//
func (op *HostNicService) UpdateVirtualFunctionsConfiguration (
    async bool,
    virtualFunctionsConfiguration HostNicVirtualFunctionsConfiguration,
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Populate the action:
    action = Action{
        Async: async,
        VirtualFunctionsConfiguration: virtualFunctionsConfiguration,
    }

    // Send the request and wait for the response:
    return internalAction(action, 'updatevirtualfunctionsconfiguration', nil, headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *HostNicService) Service(path string) *HostNicService {
    if path == nil {
        return op
    }
    if path == "networkattachments" {
        return &NetworkAttachmentsService{}
    }
    if path.startswith('networkattachments/') {
        return NetworkAttachmentsService().service(path[19:])
    }
    if path == "networklabels" {
        return &NetworkLabelsService{}
    }
    if path.startswith('networklabels/') {
        return NetworkLabelsService().service(path[14:])
    }
    if path == "statistics" {
        return &StatisticsService{}
    }
    if path.startswith('statistics/') {
        return StatisticsService().service(path[11:])
    }
    if path == "virtualfunctionallowedlabels" {
        return &VirtualFunctionAllowedLabelsService{}
    }
    if path.startswith('virtualfunctionallowedlabels/') {
        return VirtualFunctionAllowedLabelsService().service(path[29:])
    }
    if path == "virtualfunctionallowednetworks" {
        return &VirtualFunctionAllowedNetworksService{}
    }
    if path.startswith('virtualfunctionallowednetworks/') {
        return VirtualFunctionAllowedNetworksService().service(path[31:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *HostNicService) String() string {
    return fmt.Sprintf("HostNicService:%s", op.Path)
}


//
//
type HostNumaNodeService struct {
    MeasurableService

    StatisticsService  *StatisticsService
}

func NewHostNumaNodeService(connection *Connection, path string) *HostNumaNodeService {
    var result HostNumaNodeService
    result.Connection = connection
    result.Path = path
    return &result
}

//
//
func (op *HostNumaNodeService) Get (
    headers map[string]string,
    query map[string]string,
    wait bool) {
    if wait == nil {
        wait = true
    }
    // Build the URL:
    if query == nil {
        query = make(map[string]string)
    }

    // Send the request and wait for the response:
    return op.internalGet(headers, query, wait)
}

//
// Service locator method, returns individual service on which the URI is dispatched.
//
func (op *HostNumaNodeService) Service(path string) *HostNumaNodeService {
    if path == nil {
        return op
    }
    if path == "statistics" {
        return &StatisticsService{}
    }
    if path.startswith('statistics/') {
        return StatisticsService().service(path[11:])
    }
    raise Error('The path \"%s\" doesn\'t correspond to any service' % path)
}

func (op *HostNumaNodeService) String() string {
    return fmt.Sprintf("HostNumaNodeService:%s", op.Path)
}

