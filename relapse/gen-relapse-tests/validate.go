//  Copyright 2015 Walter Schulze
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
	"encoding/json"
	"encoding/xml"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/relapse/protonum"
)

type Validator struct {
	Name       string
	CodecName  string
	Grammar    *ast.Grammar
	Expected   bool
	Bytes      []byte
	SchemaName string
	Extension  string
}

var Validators = []Validator{}

type ProtoMessage interface {
	proto.Message
	Description() *descriptor.FileDescriptorSet
}

func ValidateProtoEtc(name string, grammar combinator.G, m ProtoMessage, expected bool) {
	ValidateReflect(name, grammar, m, expected)
	ValidateJson(name, grammar, m, expected)
	ValidateProtoName(name, grammar, m, expected)
}

func ValidateProtoNumEtc(name string, grammar combinator.G, m ProtoMessage, expected bool) {
	ValidateReflect(name, grammar, m, expected)
	ValidateJson(name, grammar, m, expected)
	ValidateProtoName(name, grammar, m, expected)
	ValidateProtoNum(name, grammar, m, expected)
}

func ValidateProtoNum(name string, grammar combinator.G, m ProtoMessage, expected bool) {
	packageName := "main"
	messageName := reflect.TypeOf(m).Elem().Name()
	desc := m.(ProtoMessage).Description()
	g, err := protonum.FieldNamesToNumbers(packageName, messageName, desc, grammar.Grammar())
	if err != nil {
		panic(name + ": " + err.Error())
	}
	schemaName := registerProto(m)
	Validators = append(Validators, Validator{
		Name:       name,
		CodecName:  "protoNum",
		Grammar:    g,
		Expected:   expected,
		Bytes:      mustBytes(proto.Marshal(m)),
		SchemaName: schemaName,
		Extension:  schemaName + ".pbnum",
	})
}

func ValidateProtoName(name string, g combinator.G, m ProtoMessage, expected bool) {
	schemaName := registerProto(m)
	Validators = append(Validators, Validator{
		Name:       name,
		CodecName:  "protoName",
		Grammar:    g.Grammar(),
		Expected:   expected,
		Bytes:      mustBytes(proto.Marshal(m)),
		SchemaName: schemaName,
		Extension:  schemaName + ".pbname",
	})
}

func ValidateJsonString(name string, g combinator.G, s string, expected bool) {
	Validators = append(Validators, Validator{
		Name:      name,
		CodecName: "json",
		Grammar:   g.Grammar(),
		Expected:  expected,
		Bytes:     []byte(s),
		Extension: "json",
	})
}

func ValidateJson(name string, g combinator.G, m interface{}, expected bool) {
	Validators = append(Validators, Validator{
		Name:      name,
		CodecName: "json",
		Grammar:   g.Grammar(),
		Expected:  expected,
		Bytes:     mustBytes(json.MarshalIndent(m, "", "\t")),
		Extension: "json",
	})
}

func ValidateReflect(name string, g combinator.G, m interface{}, expected bool) {
	Validators = append(Validators, Validator{
		Name:      name,
		CodecName: "goreflect",
		Grammar:   g.Grammar(),
		Expected:  expected,
		Bytes:     mustBytes(json.MarshalIndent(m, "", "\t")),
		Extension: "goreflect",
	})
}

func ValidateXMLString(name string, g combinator.G, s string, expected bool) {
	Validators = append(Validators, Validator{
		Name:      name,
		CodecName: "xml",
		Grammar:   g.Grammar(),
		Expected:  expected,
		Bytes:     []byte(s),
		Extension: "xml",
	})
}

func ValidateXML(name string, g combinator.G, m interface{}, expected bool) {
	Validators = append(Validators, Validator{
		Name:      name,
		CodecName: "xml",
		Grammar:   g.Grammar(),
		Expected:  expected,
		Bytes:     mustBytes(xml.MarshalIndent(m, "", "\t")),
		Extension: "xml",
	})
}

func mustBytes(bs []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return bs
}
