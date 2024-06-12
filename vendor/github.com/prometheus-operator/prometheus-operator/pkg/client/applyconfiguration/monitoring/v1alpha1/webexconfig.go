// Copyright The prometheus-operator Authors
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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1alpha1"
)

// WebexConfigApplyConfiguration represents an declarative configuration of the WebexConfig type for use
// with apply.
type WebexConfigApplyConfiguration struct {
	SendResolved *bool                         `json:"sendResolved,omitempty"`
	APIURL       *v1alpha1.URL                 `json:"apiURL,omitempty"`
	HTTPConfig   *HTTPConfigApplyConfiguration `json:"httpConfig,omitempty"`
	Message      *string                       `json:"message,omitempty"`
	RoomID       *string                       `json:"roomID,omitempty"`
}

// WebexConfigApplyConfiguration constructs an declarative configuration of the WebexConfig type for use with
// apply.
func WebexConfig() *WebexConfigApplyConfiguration {
	return &WebexConfigApplyConfiguration{}
}

// WithSendResolved sets the SendResolved field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SendResolved field is set to the value of the last call.
func (b *WebexConfigApplyConfiguration) WithSendResolved(value bool) *WebexConfigApplyConfiguration {
	b.SendResolved = &value
	return b
}

// WithAPIURL sets the APIURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIURL field is set to the value of the last call.
func (b *WebexConfigApplyConfiguration) WithAPIURL(value v1alpha1.URL) *WebexConfigApplyConfiguration {
	b.APIURL = &value
	return b
}

// WithHTTPConfig sets the HTTPConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPConfig field is set to the value of the last call.
func (b *WebexConfigApplyConfiguration) WithHTTPConfig(value *HTTPConfigApplyConfiguration) *WebexConfigApplyConfiguration {
	b.HTTPConfig = value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *WebexConfigApplyConfiguration) WithMessage(value string) *WebexConfigApplyConfiguration {
	b.Message = &value
	return b
}

// WithRoomID sets the RoomID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RoomID field is set to the value of the last call.
func (b *WebexConfigApplyConfiguration) WithRoomID(value string) *WebexConfigApplyConfiguration {
	b.RoomID = &value
	return b
}
