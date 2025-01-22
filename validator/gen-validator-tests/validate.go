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

	"github.com/katydid/validator-go/validator/ast"
	"github.com/katydid/validator-go/validator/combinator"
	"google.golang.org/protobuf/proto"
	descriptor "google.golang.org/protobuf/types/descriptorpb"
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

var duplicates = map[string]struct{}{}

func checkDuplicates(name, codecName string) {
	n := name + "." + codecName
	if _, ok := duplicates[n]; ok {
		panic("duplicate validator: " + n)
	}
	duplicates[n] = struct{}{}
}

type ProtoMessage interface {
	proto.Message
	Description() *descriptor.FileDescriptorSet
}

func ValidateProtoEtc(name string, grammar combinator.G, m ProtoMessage, expected bool) {
	ValidateReflect(name, grammar, m, expected)
	ValidateJson(name, grammar, m, expected)
	ValidateProto(name, grammar, m, expected)
}

func ValidateProto(name string, g combinator.G, m ProtoMessage, expected bool) {
	schemaName := registerProto(m)
	checkDuplicates(name, "pb")
	Validators = append(Validators, Validator{
		Name:       name,
		CodecName:  "pb",
		Grammar:    g.Grammar(),
		Expected:   expected,
		Bytes:      mustBytes(proto.Marshal(m)),
		SchemaName: schemaName,
		Extension:  schemaName + ".pb",
	})
}

func ValidateJsonString(name string, g combinator.G, s string, expected bool) {
	checkDuplicates(name, "json")
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
	checkDuplicates(name, "json")
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
	checkDuplicates(name, "goreflect")
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
	checkDuplicates(name, "xml")
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
	checkDuplicates(name, "xml")
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
