/*
 * Copyright 2022-2023 Chaos Meta Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package experiment

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

var (
	JmeterConfigStr string
	JobYamlStr      string
)

type LoadTestSpec struct {
	FlowType    string     `json:"flowType"`
	Duration    string     `json:"duration"`
	Parallelism int        `json:"parallelism"`
	Source      int        `json:"source"`
	Stopped     bool       `json:"stopped"`
	Args        []FlowArgs `json:"args"`
}

type FlowArgs struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type LoadTestStatus struct {
	Status       StatusType `json:"status"`
	Message      string     `json:"message"`
	SuccessCount int        `json:"successCount"`
	TotalCount   int        `json:"totalCount"`
	AvgRPS       int        `json:"avgRPS"`
	CreateTime   string     `json:"createTime"`
	UpdateTime   string     `json:"updateTime"`
}

type LoadTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoadTestSpec   `json:"spec,omitempty"`
	Status LoadTestStatus `json:"status,omitempty"`
}

type LoadTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadTest `json:"items"`
}
