// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package servers

// [START cloud_game_servers_deployment_update]

import (
	"context"
	"fmt"
	"io"

	gaming "cloud.google.com/go/gaming/apiv1"
	gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// updateGameServerDeployment updates a game server deployment.
func updateGameServerDeployment(w io.Writer, projectID, deploymentID string) error {
	// projectID := "my-project"
	// deploymentID := "mydeployment"
	ctx := context.Background()
	client, err := gaming.NewGameServerDeploymentsClient(ctx)
	if err != nil {
		return fmt.Errorf("NewGameServerDeploymentsClient: %v", err)
	}
	defer client.Close()

	req := &gamingpb.UpdateGameServerDeploymentRequest{
		GameServerDeployment: &gamingpb.GameServerDeployment{
			Name:        fmt.Sprintf("projects/%s/locations/global/gameServerDeployments/%s", projectID, deploymentID),
			Description: "My Updated Game Server Deployment",
			Labels: map[string]string{
				"label-key-1": "label-value-1",
			},
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{
				"description", "labels",
			},
		},
	}

	op, err := client.UpdateGameServerDeployment(ctx, req)
	if err != nil {
		return fmt.Errorf("UpdateGameServerDeployment: %v", err)
	}
	resp, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Wait: %v", err)
	}

	fmt.Fprintf(w, "Deployment updated: %v", resp.Name)
	return nil
}

// [END cloud_game_servers_deployment_update]
