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

package providers

import (
	"sigs.k8s.io/kind/pkg/cluster/nodes"

	"sigs.k8s.io/kind/pkg/internal/apis/config"
	"sigs.k8s.io/kind/pkg/internal/cli"
)

// Provider represents a provider of cluster / node infrastructure(DockerやPodmanの事)
// This is an alpha-grade internal API(内部APIだよ)
type Provider interface {
	// Provision では、K8sを作成する前に必要なノードの起動をcluster configに基づいて行う
	Provision(status *cli.Status, cfg *config.Cluster) error
	// ListClustersはリソースからプロバイダ配下のクラスタの検出を行う？
	ListClusters() ([]string, error)
	// ListNodesはクラスタ名からプロバイダ配下にあるノードを返す
	ListNodes(cluster string) ([]nodes.Node, error)
	// DeleteNodes は指定されたノードのリストを削除する。
	// These should be from results previously returned by this provider
	// E.G. by ListNodes()
	DeleteNodes([]nodes.Node) error
	// GetAPIServerEndpointは、クラスターのAPIサーバーのホストエンドポイントを返します。
	GetAPIServerEndpoint(cluster string) (string, error)
	// GetAPIServerEndpointは、クラスターのAPIサーバーの内部ネットワークエンドポイントを返します。
	GetAPIServerInternalEndpoint(cluster string) (string, error)
	// CollectLogsは、クラスタログおよびその他のデバッグファイルをdirに格納します。
	CollectLogs(dir string, nodes []nodes.Node) error
	// Info returns the provider info
	Info() (*ProviderInfo, error)
}

// ProviderInfo is the info of the provider
type ProviderInfo struct {
	Rootless            bool
	Cgroup2             bool
	SupportsMemoryLimit bool
	SupportsPidsLimit   bool
	SupportsCPUShares   bool
}
