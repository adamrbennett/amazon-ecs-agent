// Copyright 2014-2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// The 'engine' package contains code for interacting with container-running backends and handling events from them.
// It supports Docker as the sole task engine type.
package engine

import (
	"github.com/adamrbennett/amazon-ecs-agent/agent/config"
	"github.com/adamrbennett/amazon-ecs-agent/agent/logger"
)

var log = logger.ForModule("TaskEngine")

// NewTaskEngine returns a default TaskEngine
func NewTaskEngine(cfg *config.Config, client DockerClient) TaskEngine {
	return NewDockerTaskEngine(cfg, client)
}
