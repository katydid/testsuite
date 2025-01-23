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
	"encoding/xml"
	"math/rand"
	"reflect"

	"google.golang.org/protobuf/proto"

	jsonparser "github.com/katydid/parser-go-json/json"
	protoparser "github.com/katydid/parser-go-proto/proto"
	xmlparser "github.com/katydid/parser-go-xml/xml"
	"github.com/katydid/validator-go/validator"
	"github.com/katydid/validator-go/validator/ast"
	"github.com/katydid/validator-go/validator/combinator"
)

type BenchValidator struct {
	Name         string
	CodecName    string
	Grammar      *ast.Grammar
	ValidBytes   RandBytes
	InvalidBytes RandBytes
	SchemaName   string
	Extension    string
	Validate     func([]byte) bool
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

func BenchValidateProtoJson(name string, grammar combinator.G, validProto, invalidProto RandProto) {
	BenchValidateProto(name, grammar, validProto, invalidProto)
	BenchValidateJson(name, grammar, validProto, invalidProto)
}

func BenchValidateProtoEtc(name string, grammar combinator.G, validProto, invalidProto RandProto) {
	BenchValidateProto(name, grammar, validProto, invalidProto)
	BenchValidateJson(name, grammar, validProto, invalidProto)
	BenchValidateXML(name, grammar, validProto, invalidProto)
}

func BenchValidateProto(name string, grammar combinator.G, validProto, invalidProto RandProto) {
	msg := validProto(rand.New(rand.NewSource(1)))
	packageName := "main"
	messageName := reflect.TypeOf(msg).Elem().Name()
	desc := msg.(ProtoMessage).Description()
	g := grammar.Grammar()
	validBytes := func(r *rand.Rand) []byte {
		pb := validProto(r)
		return mustBytes(proto.Marshal(pb))
	}
	invalidBytes := func(r *rand.Rand) []byte {
		pb := invalidProto(r)
		return mustBytes(proto.Marshal(pb))
	}
	schemaName := registerProto(msg)

	m, err := validator.Prepare(g)
	if err != nil {
		panic(err)
	}

	BenchValidators = append(BenchValidators, BenchValidator{
		Name:         name,
		CodecName:    "pb",
		Grammar:      g,
		ValidBytes:   validBytes,
		InvalidBytes: invalidBytes,
		SchemaName:   schemaName,
		Extension:    schemaName + ".pb",
		Validate: func(buf []byte) bool {
			p, err := protoparser.NewParserWithDesc(packageName, messageName, desc)
			if err != nil {
				panic(err)
			}
			if err := p.Init(buf); err != nil {
				panic(err)
			}
			v, err := validator.Validate(m, p)
			return v && err == nil
		},
	})
	checkDuplicateBenches(name, "pb")
}

func BenchValidateJson(name string, grammar combinator.G, validProto, invalidProto RandProto) {
	g := grammar.Grammar()
	validBytes := func(r *rand.Rand) []byte {
		pb := validProto(r)
		return mustBytes(json.Marshal(pb))
	}
	invalidBytes := func(r *rand.Rand) []byte {
		pb := invalidProto(r)
		return mustBytes(json.Marshal(pb))
	}

	m, err := validator.Prepare(g)
	if err != nil {
		panic(err)
	}

	BenchValidators = append(BenchValidators, BenchValidator{
		Name:         name,
		CodecName:    "json",
		Grammar:      g,
		ValidBytes:   validBytes,
		InvalidBytes: invalidBytes,
		Extension:    "json",
		Validate: func(buf []byte) bool {
			p := jsonparser.NewJsonParser()
			if err := p.Init(buf); err != nil {
				panic(err)
			}
			v, err := validator.Validate(m, p)
			return v && err == nil
		},
	})
	checkDuplicateBenches(name, "json")
}

func BenchValidateXML(name string, grammar combinator.G, validProto, invalidProto RandProto) {
	g := grammar.Grammar()
	validBytes := func(r *rand.Rand) []byte {
		pb := validProto(r)
		return mustBytes(xml.Marshal(pb))
	}
	invalidBytes := func(r *rand.Rand) []byte {
		pb := invalidProto(r)
		return mustBytes(xml.Marshal(pb))
	}

	m, err := validator.Prepare(g)
	if err != nil {
		panic(err)
	}

	BenchValidators = append(BenchValidators, BenchValidator{
		Name:         name,
		CodecName:    "xml",
		Grammar:      g,
		ValidBytes:   validBytes,
		InvalidBytes: invalidBytes,
		Extension:    "xml",
		Validate: func(buf []byte) bool {
			p := xmlparser.NewXMLParser()
			if err := p.Init(buf); err != nil {
				panic(err)
			}
			v, err := validator.Validate(m, p)
			return v && err == nil
		},
	})
	checkDuplicateBenches(name, "xml")
}

type RandBytes func(r *rand.Rand) []byte

type RandProto func(r *rand.Rand) ProtoMessage

func RandomPerson(r *rand.Rand) ProtoMessage {
	person := random(r, &Person{})
	return person.(ProtoMessage)
}

func RandomSrcTree(r *rand.Rand) ProtoMessage {
	pops := []*SrcTree{IoUtilSrcTree, PathSrcTree, RuntimeSrcTree, SyscallSrcTree}
	return pops[r.Intn(4)]
}

func RandomTypewriterPrison(r *rand.Rand) ProtoMessage {
	return random(r, &TypewriterPrison{}).(ProtoMessage)
}

func RandomPuddingMilkshake(r *rand.Rand) ProtoMessage {
	return random(r, &PuddingMilkshake{}).(ProtoMessage)
}
