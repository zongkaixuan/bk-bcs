/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2 "github.com/Tencent/bk-bcs/bcs-mesos/pkg/client/internalclientset/typed/bkbcs/v2"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeBkbcsV2 struct {
	*testing.Fake
}

func (c *FakeBkbcsV2) AdmissionWebhookConfigurations(namespace string) v2.AdmissionWebhookConfigurationInterface {
	return &FakeAdmissionWebhookConfigurations{c, namespace}
}

func (c *FakeBkbcsV2) Agents(namespace string) v2.AgentInterface {
	return &FakeAgents{c, namespace}
}

func (c *FakeBkbcsV2) AgentSchedInfos(namespace string) v2.AgentSchedInfoInterface {
	return &FakeAgentSchedInfos{c, namespace}
}

func (c *FakeBkbcsV2) Applications(namespace string) v2.ApplicationInterface {
	return &FakeApplications{c, namespace}
}

func (c *FakeBkbcsV2) BcsClusterAgentSettings(namespace string) v2.BcsClusterAgentSettingInterface {
	return &FakeBcsClusterAgentSettings{c, namespace}
}

func (c *FakeBkbcsV2) BcsCommandInfos(namespace string) v2.BcsCommandInfoInterface {
	return &FakeBcsCommandInfos{c, namespace}
}

func (c *FakeBkbcsV2) BcsConfigMaps(namespace string) v2.BcsConfigMapInterface {
	return &FakeBcsConfigMaps{c, namespace}
}

func (c *FakeBkbcsV2) BcsDaemonsets(namespace string) v2.BcsDaemonsetInterface {
	return &FakeBcsDaemonsets{c, namespace}
}

func (c *FakeBkbcsV2) BcsEndpoints(namespace string) v2.BcsEndpointInterface {
	return &FakeBcsEndpoints{c, namespace}
}

func (c *FakeBkbcsV2) BcsSecrets(namespace string) v2.BcsSecretInterface {
	return &FakeBcsSecrets{c, namespace}
}

func (c *FakeBkbcsV2) BcsServices(namespace string) v2.BcsServiceInterface {
	return &FakeBcsServices{c, namespace}
}

func (c *FakeBkbcsV2) Crds(namespace string) v2.CrdInterface {
	return &FakeCrds{c, namespace}
}

func (c *FakeBkbcsV2) Crrs(namespace string) v2.CrrInterface {
	return &FakeCrrs{c, namespace}
}

func (c *FakeBkbcsV2) Deployments(namespace string) v2.DeploymentInterface {
	return &FakeDeployments{c, namespace}
}

func (c *FakeBkbcsV2) Frameworks(namespace string) v2.FrameworkInterface {
	return &FakeFrameworks{c, namespace}
}

func (c *FakeBkbcsV2) Tasks(namespace string) v2.TaskInterface {
	return &FakeTasks{c, namespace}
}

func (c *FakeBkbcsV2) TaskGroups(namespace string) v2.TaskGroupInterface {
	return &FakeTaskGroups{c, namespace}
}

func (c *FakeBkbcsV2) Versions(namespace string) v2.VersionInterface {
	return &FakeVersions{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeBkbcsV2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
