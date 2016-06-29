// Copyright 2014-2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

// Package tcsclient wraps the generated aws-sdk-go client to provide marshalling
// and unmarshalling of data over a websocket connection in the format expected
// by TCS. It allows for bidirectional communication and acts as both a
// client-and-server in terms of requests, but only as a client in terms of
// connecting.

package tcsclient

import (
	"reflect"

	"github.com/adamrbennett/amazon-ecs-agent/agent/tcs/model/ecstcs"
)

var tcsTypeMappings map[string]reflect.Type

func init() {
	// This list is currently *manually updated* and assumes that the generated
	// struct type-names within the package *exactly match* the type sent by ACS/TCS
	// (true so far; careful with inflections)
	// TODO, this list should be autogenerated
	// I couldn't figure out how to get a list of all structs in a package via
	// reflection, but that would solve this. The alternative is to either parse
	// the .json model or the generated struct names.
	recognizedTypes := []interface{}{
		ecstcs.StopTelemetrySessionMessage{},
		ecstcs.AckPublishMetric{},
		ecstcs.HeartbeatMessage{},
		ecstcs.PublishMetricsRequest{},
		ecstcs.StartTelemetrySessionRequest{},
		ecstcs.ServerException{},
		ecstcs.BadRequestException{},
		ecstcs.ResourceValidationException{},
		ecstcs.InvalidParameterException{},
	}

	tcsTypeMappings = make(map[string]reflect.Type)
	// This produces a map of:
	// "MyMessage": TypeOf(ecstcs.MyMessage)
	for _, recognizedType := range recognizedTypes {
		tcsTypeMappings[reflect.TypeOf(recognizedType).Name()] = reflect.TypeOf(recognizedType)
	}
}

// TcsDecoder implments wsclient.TypeDecoder.
type TcsDecoder struct{}

func (dc *TcsDecoder) NewOfType(tcsType string) (interface{}, bool) {
	rtype, ok := tcsTypeMappings[tcsType]
	if !ok {
		return nil, false
	}
	return reflect.New(rtype).Interface(), true
}

func (dc *TcsDecoder) GetRecognizedTypes() map[string]reflect.Type {
	return tcsTypeMappings
}
