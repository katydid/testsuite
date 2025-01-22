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
	. "github.com/katydid/validator-go/validator/combinator"
	"google.golang.org/protobuf/proto"
)

var TopTop = &TopsyTurvy{
	Top: &Topsy{
		Top: proto.Int64(1),
	},
}

var TurfTurf = &TopsyTurvy{
	Turf: &Turvy{
		Turf: proto.Int64(1),
	},
}

var TopTop1 = G{"main": InAny(In("Top", Value(Eq(IntVar(), IntConst(1)))))}

func init() {
	ValidateProtoEtc("TopTop1", TopTop1, TopTop, true)
	ValidateProtoEtc("TurfTurf1", TopTop1, TurfTurf, false)
}
