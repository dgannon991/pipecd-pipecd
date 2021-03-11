// Copyright 2021 The PipeCD Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ecs

import (
	"io/ioutil"

	"sigs.k8s.io/yaml"

	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

func loadTaskDefinition(path string) (types.TaskDefinition, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return types.TaskDefinition{}, err
	}
	return parseTaskDefinition(data)
}

func parseTaskDefinition(data []byte) (types.TaskDefinition, error) {
	var obj types.TaskDefinition
	// TODO: Support loading TaskDefinition file with JSON format
	if err := yaml.Unmarshal(data, &obj); err != nil {
		return types.TaskDefinition{}, err
	}
	return obj, nil
}