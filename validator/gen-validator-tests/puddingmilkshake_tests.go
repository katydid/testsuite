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
	"github.com/gogo/protobuf/proto"
	. "github.com/katydid/validator-go/validator/combinator"
	rparser "github.com/katydid/validator-go/validator/parser"
)

var Ab21FinanceJudo = &FinanceJudo{
	SaladWorry: &SaladWorry{
		MagazineFrame: []string{"a", "b"},
		XrayPilot: &XrayPilot{
			AnkleCoat: proto.Int64(2),
		},
	},
	RumourSpirit: proto.Int64(1),
}

var AnyFinanceJudo = G{"main": Any()}

func init() {
	ValidateProtoEtc("Ab21Any", AnyFinanceJudo, Ab21FinanceJudo, true)
}

var NoneFinanceJudo = G{"main": None()}

func init() {
	ValidateProtoEtc("Ab21None", NoneFinanceJudo, Ab21FinanceJudo, false)
}

var HasSpirit1FinanceJudo = G{
	"main":   InOrder(Any(), Eval("spirit"), Any()),
	"spirit": In("RumourSpirit", Value(Eq(IntVar(), IntConst(1)))),
}

func init() {
	ValidateProtoEtc("Ab21Spirit1", HasSpirit1FinanceJudo, Ab21FinanceJudo, true)
}

var HasSpirit2FinanceJudo = G{
	"main":   InOrder(Any(), Eval("spirit"), Any()),
	"spirit": In("RumourSpirit", Value(Eq(IntVar(), IntConst(2)))),
}

func init() {
	ValidateProtoEtc("Ab21Spirit2", HasSpirit2FinanceJudo, Ab21FinanceJudo, false)
}

var MagazineFrameAFinanceJudo = G{
	"main": InOrder(
		In("SaladWorry",
			In("MagazineFrame", InOrder(
				InAny(Value(Eq(StringVar(), StringConst("a")))),
				Any()),
			),
			In("XrayPilot", Any()),
			Any(),
		),
		Any(),
	),
}

func init() {
	ValidateProtoEtc("Ab21MagazineFrameA", MagazineFrameAFinanceJudo, Ab21FinanceJudo, true)
}

var MagazineFrameSingleAFinanceJudo = G{
	"main": InOrder(
		In("SaladWorry",
			In("MagazineFrame", Value(Eq(StringVar(), StringConst("a")))),
			In("XrayPilot", Any()),
			Any(),
		),
		Any(),
	),
}

func init() {
	ValidateProtoEtc("Ab21MagazineFrameSingleA", MagazineFrameSingleAFinanceJudo, Ab21FinanceJudo, false)
}

var InAnyExceptNotAFieldNameFinanceJudo = G{
	"main": InOrder(
		InAnyExcept("NotAFieldName",
			In("MagazineFrame",
				InOrder(
					InAny(Value(Eq(StringVar(), StringConst("a")))),
					Any()),
			),
			Any(),
		),
		Any(),
	),
}

func init() {
	ValidateProtoEtc("Ab21NotAFieldName", InAnyExceptNotAFieldNameFinanceJudo, Ab21FinanceJudo, true)
}

var InAnyExceptSaladWorryFinanceJudo = G{
	"main": InOrder(
		InAnyExcept("SaladWorry",
			In("MagazineFrame", Value(Eq(StringVar(), StringConst("a")))),
			Any(),
		),
		Any(),
	),
}

func init() {
	ValidateProtoEtc("Ab21InAnyExceptSaladWorry", InAnyExceptSaladWorryFinanceJudo, Ab21FinanceJudo, false)
}

func init() {
	g1 := `.FinanceJudo:.SaladWorry:.SpyCarpenter:
(
	(
		(
			(
					.BridgePepper:!(._ == "aaaaaaaa@mm~")
				&
					.FountainTarget:!(._ == "aaaaaaaa@mm~")
			)
			&
			.BridgePepper:!(._ == "bbbbbbb@~")
		)
		&
		.FountainTarget:!(._ == "bbbbbbb@~")
	)
	&
	(
			.BridgePepper:._ == "mmm.ddddddd~"
		|
			.FountainTarget:._ == "mmm.ddddddd~"
	)
)`

	grammar1, err := rparser.ParseGrammar(g1)
	if err != nil {
		panic(err)
	}
	msg := &PuddingMilkshake{FinanceJudo: &FinanceJudo{SaladWorry: &SaladWorry{SpyCarpenter: &SpyCarpenter{
		BridgePepper:   []string{"isHOrIGyoLbdXZ9a4t4abCuoFvDpXvxgscJQYRGZ6u"},
		FountainTarget: []string{"oqqST33HqlR5s30O61mPwPnXGrwM5AIRWwDQ1YDPZcr8iP56B7AFwemBq1MfsNojkOAPlkt58RuaNn7pTgV66TSpp"},
	}}}}

	ValidateProtoEtc("PuddingMilkShakeNotAny1", G{"main": grammar1.GetTopPattern()}, msg, false)
}

func init() {
	g2 := `.FinanceJudo:.SaladWorry:.SpyCarpenter:
(
	(
		(
			(
					.BridgePepper:!(._ == "aaaaaaaa@mm~")
				&
					.FountainTarget:!(._ == "aaaaaaaa@mm~")
			)
			&
			.BridgePepper:!(._ == "bbbbbbb@~")
		)
		&
		.FountainTarget:!(._ == "bbbbbbb@~")
	)
	&
	(
			.BridgePepper:._ == "mmm.dddddddaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasss~"
		|
			.FountainTarget:._ == "mmm.dddddddaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasss~"
	)
)`

	grammar2, err := rparser.ParseGrammar(g2)
	if err != nil {
		panic(err)
	}
	msg := &PuddingMilkshake{FinanceJudo: &FinanceJudo{SaladWorry: &SaladWorry{SpyCarpenter: &SpyCarpenter{
		BridgePepper:   []string{"isHOrIGyoLbdXZ9a4t4abCuoFvDpXvxgscJQYRGZ6u"},
		FountainTarget: []string{"oqqST33HqlR5s30O61mPwPnXGrwM5AIRWwDQ1YDPZcr8iP56B7AFwemBq1MfsNojkOAPlkt58RuaNn7pTgV66TSpp"},
	}}}}

	ValidateProtoEtc("PuddingMilkShakeNotAny2", G{"main": grammar2.GetTopPattern()}, msg, false)
}
