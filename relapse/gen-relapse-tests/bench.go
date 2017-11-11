//  Copyright 2016 Walter Schulze
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
	"math/rand"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/relapse/protonum"
)

type BenchValidator struct {
	Name       string
	CodecName  string
	Grammar    *ast.Grammar
	RandBytes  RandBytes
	SchemaName string
	Extension  string
}

var BenchValidators = []BenchValidator{}
var duplicatesBenches = map[string]struct{}{}

func checkDuplicateBenches(name, codecName string) {
	n := name + "." + codecName
	if _, ok := duplicatesBenches[n]; ok {
		panic("duplicate benchmark: " + n)
	}
	duplicatesBenches[n] = struct{}{}
}

func BenchValidateProtoNumEtc(name string, grammar combinator.G, randProto RandProto) {
	BenchValidateProtoNum(name, grammar, randProto)
	BenchValidateJson(name, grammar, randProto)
}

func BenchValidateProtoNum(name string, grammar combinator.G, randProto RandProto) {
	m := randProto(rand.New(rand.NewSource(1)))
	packageName := "main"
	messageName := reflect.TypeOf(m).Elem().Name()
	desc := m.(ProtoMessage).Description()
	g, err := protonum.FieldNamesToNumbers(packageName, messageName, desc, grammar.Grammar())
	if err != nil {
		panic(name + ": " + err.Error())
	}
	randBytes := func(r *rand.Rand) []byte {
		pb := randProto(r)
		return mustBytes(proto.Marshal(pb))
	}
	schemaName := registerProto(m)
	BenchValidators = append(BenchValidators, BenchValidator{
		Name:       name,
		CodecName:  "pbnum",
		Grammar:    g,
		RandBytes:  randBytes,
		SchemaName: schemaName,
		Extension:  schemaName + ".pbnum",
	})
	checkDuplicateBenches(name, "pbnum")
}

func BenchValidateJson(name string, grammar combinator.G, randProto RandProto) {
	randBytes := func(r *rand.Rand) []byte {
		pb := randProto(r)
		return mustBytes(json.Marshal(pb))
	}
	BenchValidators = append(BenchValidators, BenchValidator{
		Name:      name,
		CodecName: "json",
		Grammar:   grammar.Grammar(),
		RandBytes: randBytes,
		Extension: "json",
	})
	checkDuplicateBenches(name, "json")
}

type RandBytes func(r *rand.Rand) []byte

type RandProto func(r *rand.Rand) ProtoMessage

func RandomPerson(r *rand.Rand) ProtoMessage {
	return NewPopulatedPerson(r, true)
}

func RandomSrcTree(r *rand.Rand) ProtoMessage {
	pops := []*SrcTree{IoUtilSrcTree, PathSrcTree, RuntimeSrcTree, SyscallSrcTree}
	return pops[r.Intn(4)]
}

func RandomTypewriterPrison(r *rand.Rand) ProtoMessage {
	return NewPopulatedTypewriterPrison(r, true)
}

func RandomPuddingMilkshake(r *rand.Rand) ProtoMessage {
	return NewPopulatedPuddingMilkshake(r, true)
}
