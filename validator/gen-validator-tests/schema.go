//  Copyright 2017 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"reflect"

	"google.golang.org/protobuf/proto"
)

var Schemas = map[string]Schema{}

type Schema interface {
	Data() []byte
}

type ProtoDesc struct {
	PackageName string
	MessageName string
	Desc        []byte
}

func (this ProtoDesc) Data() []byte {
	return this.Desc
}

func registerProto(m ProtoMessage) string {
	packageName := "main"
	messageName := reflect.TypeOf(m).Elem().Name()
	schemaName := packageName + "." + messageName + ".desc"
	if _, ok := Schemas[schemaName]; ok {
		return schemaName
	}
	desc := m.(ProtoMessage).Description()
	descBytes := mustBytes(proto.Marshal(desc))
	Schemas[schemaName] = ProtoDesc{
		PackageName: packageName,
		MessageName: messageName,
		Desc:        descBytes,
	}
	return schemaName
}
