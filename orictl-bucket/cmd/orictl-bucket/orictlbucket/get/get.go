// Copyright 2022 IronCore authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package get

import (
	"github.com/ironcore-dev/ironcore/orictl-bucket/cmd/orictl-bucket/orictlbucket/common"
	"github.com/ironcore-dev/ironcore/orictl-bucket/cmd/orictl-bucket/orictlbucket/get/bucket"
	"github.com/ironcore-dev/ironcore/orictl-bucket/cmd/orictl-bucket/orictlbucket/get/bucketclass"
	orictlcmd "github.com/ironcore-dev/ironcore/orictl/cmd"
	"github.com/spf13/cobra"
)

func Command(streams orictlcmd.Streams, clientFactory common.ClientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "get",
	}

	cmd.AddCommand(
		bucket.Command(streams, clientFactory),
		bucketclass.Command(streams, clientFactory),
	)

	return cmd
}
