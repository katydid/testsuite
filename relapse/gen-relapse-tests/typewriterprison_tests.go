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
	. "github.com/katydid/katydid/relapse/combinator"
)

var AbcPocketRoses = &PocketRoses{
	MenuPaperclip: []string{"a", "b", "c"},
}

var AStringArray = G{"main": In("MenuPaperclip", Elem(0, Value(Eq(StringVar(), StringConst("a")))), Any())}

var FinalCStringArray = G{"main": In("MenuPaperclip", Any(), Elem(2, Value(Eq(StringVar(), StringConst("c")))))}

var AbcStringArray = G{"main": In("MenuPaperclip",
	Elem(0, Value(Eq(StringVar(), StringConst("a")))),
	Elem(1, Value(Eq(StringVar(), StringConst("b")))),
	Elem(2, Value(Eq(StringVar(), StringConst("c")))),
)}

var NextAStringArray = G{"main": In("MenuPaperclip",
	Elem(1, Value(Eq(StringVar(), StringConst("a")))),
	Elem(0, Value(Eq(StringVar(), StringConst("b")))),
	Elem(2, Value(Eq(StringVar(), StringConst("c")))),
)}

var DStringArray = G{"main": In("MenuPaperclip", Elem(0, Value(Eq(StringVar(), StringConst("d")))), Any())}

func init() {
	ValidateProtoNumEtc("APocketRoses", AStringArray, AbcPocketRoses, true)
	ValidateProtoNumEtc("FinalCPocketRoses", FinalCStringArray, AbcPocketRoses, true)
	ValidateProtoNumEtc("AbcPocketRoses", AbcStringArray, AbcPocketRoses, true)
	ValidateProtoNumEtc("NextAPocketRoses", NextAStringArray, AbcPocketRoses, false)
	ValidateProtoNumEtc("DPocketRoses", DStringArray, AbcPocketRoses, false)
}

var LatentSimplificationTypewriterPrisonTrue = &TypewriterPrison{
	PocketRoses: &PocketRoses{
		MenuPaperclip: []string{"b"},
		BeetlePoker:   []string{"c"},
		BadgeShopping: proto.Int64(2),
		DaisySled:     proto.Int64(1),
	},
}

var LatentSimplificationTypewriterPrisonFalse = &TypewriterPrison{
	PocketRoses: &PocketRoses{
		MenuPaperclip: []string{"a"},
		BeetlePoker:   []string{"a"},
		BadgeShopping: proto.Int64(2),
		DaisySled:     proto.Int64(0),
	},
}

var LatentSimplificationOfExprs = G{"main": AllOf(
	AllOf(
		AnyOf(
			AnyOf(
				InPath("PocketRoses", InPath("MenuPaperclip", InAnyPath(Value(Contains(StringVar(), StringConst("a")))))),
				InPath("PocketRoses", InPath("BeetlePoker", InAnyPath(Value(Contains(StringVar(), StringConst("a")))))),
			),
			AnyOf(
				InPath("PocketRoses", InPath("MenuPaperclip", InAnyPath(Value(Contains(StringVar(), StringConst("b")))))),
				InPath("PocketRoses", InPath("BeetlePoker", InAnyPath(Value(Contains(StringVar(), StringConst("b")))))),
			),
		),
		AnyOf(
			InPath("PocketRoses", InPath("MenuPaperclip", InAnyPath(Value(Contains(StringVar(), StringConst("c")))))),
			InPath("PocketRoses", InPath("BeetlePoker", InAnyPath(Value(Contains(StringVar(), StringConst("c")))))),
		),
	),
	AllOf(
		InPath("PocketRoses", InPath("BadgeShopping", Value(LE(IntVar(), IntConst(2))))),
		AllOf(
			InPath("PocketRoses", InPath("BadgeShopping", Value(GE(IntVar(), IntConst(0))))),
			InPath("PocketRoses", InPath("DaisySled", Value(GE(IntVar(), IntConst(1))))),
		),
	),
)}

func init() {
	ValidateProtoNumEtc("LatentSimplificationTrue", LatentSimplificationOfExprs, LatentSimplificationTypewriterPrisonTrue, true)
	ValidateProtoNumEtc("LatentSimplificationFalse", LatentSimplificationOfExprs, LatentSimplificationTypewriterPrisonFalse, false)
}
