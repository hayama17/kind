/*
Copyright 2018 The Kubernetes Authors.

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

// Package create implements(実装) the `create` command
package create

import (
	"errors"

	"github.com/spf13/cobra"//CLIコマンド用のライブラリ

	"sigs.k8s.io/kind/pkg/cmd"
	createcluster "sigs.k8s.io/kind/pkg/cmd/kind/create/cluster"//エイリアスをつけてる
	"sigs.k8s.io/kind/pkg/log"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{//cmdにcreateの処理に上書き
		Args:  cobra.NoArgs,//引数があるとエラーを返す
		Use:   "create",//サブコマ
		Short: "Creates one of [cluster]",//説明
		Long:  "Creates one of local Kubernetes cluster (cluster)",//説明
		RunE: func(cmd *cobra.Command, args []string) error {//処理
			err := cmd.Help()//errにhelpの値を代入
			if err != nil {//errがnil出ないことを確認して
				return err//errの値を処理
			}
			return errors.New("Subcommand is required")
		},
	}
	cmd.AddCommand(createcluster.NewCommand(logger, streams))//cmdにcreateclusterコマンドを追加する
	return cmd//コマンドを返す
}
