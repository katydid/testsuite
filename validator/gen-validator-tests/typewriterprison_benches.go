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
	"log"
	"math/rand"
	"strings"

	. "github.com/katydid/validator-go/validator/combinator"
	"google.golang.org/protobuf/proto"
)

func RandomValidTypewriterPrisonScarBusStop(r *rand.Rand) ProtoMessage {
	p := RandomTypewriterPrison(r).(*TypewriterPrison)
	for p.PocketRoses == nil || p.GetPocketRoses().ScarBusStop == nil || !strings.Contains(p.GetPocketRoses().GetScarBusStop(), "a") {
		log.Printf("random looped: TypewriterPrisonScarBusStop")
		p = RandomTypewriterPrison(r).(*TypewriterPrison)
	}
	log.Printf("random returned: TypewriterPrisonScarBusStop")
	return p
}

func RandomInvalidTypewriterPrisonScarBusStop(r *rand.Rand) ProtoMessage {
	p := RandomTypewriterPrison(r).(*TypewriterPrison)
	for strings.Contains(p.GetPocketRoses().GetScarBusStop(), "a") {
		log.Printf("random looped: TypewriterPrisonScarBusStop")
		p = RandomTypewriterPrison(r).(*TypewriterPrison)
	}
	log.Printf("random returned: TypewriterPrisonScarBusStop")
	return p
}

func RandomValidTypewriterPrisonDaisySled(r *rand.Rand) ProtoMessage {
	p := RandomTypewriterPrison(r).(*TypewriterPrison)
	for p.PocketRoses == nil || p.GetPocketRoses().DaisySled == nil {
		log.Printf("random looped: TypewriterPrisonDaisySled")
		p = RandomTypewriterPrison(r).(*TypewriterPrison)
	}
	p.GetPocketRoses().DaisySled = proto.Int64(1)
	log.Printf("random returned: TypewriterPrisonDaisySled")
	return p
}

func RandomInvalidTypewriterPrisonDaisySled(r *rand.Rand) ProtoMessage {
	p := RandomTypewriterPrison(r).(*TypewriterPrison)
	for p.GetPocketRoses().GetDaisySled() == 1 {
		log.Printf("random looped: TypewriterPrisonDaisySled")
		p = RandomTypewriterPrison(r).(*TypewriterPrison)
	}
	log.Printf("random returned: TypewriterPrisonDaisySled")
	return p
}

func RandomValidTypewriterPrisonSmileLetter(r *rand.Rand) ProtoMessage {
	p := RandomTypewriterPrison(r).(*TypewriterPrison)
	for p.PocketRoses == nil || p.GetPocketRoses().SmileLetter == nil || !p.GetPocketRoses().GetSmileLetter() {
		log.Printf("random looped: TypewriterPrisonSmileLetter")
		p = RandomTypewriterPrison(r).(*TypewriterPrison)
	}
	log.Printf("random returned: TypewriterPrisonSmileLetter")
	return p
}

func RandomInvalidTypewriterPrisonSmileLetter(r *rand.Rand) ProtoMessage {
	p := RandomTypewriterPrison(r).(*TypewriterPrison)
	for p.GetPocketRoses().GetSmileLetter() {
		log.Printf("random looped: TypewriterPrisonSmileLetter")
		p = RandomTypewriterPrison(r).(*TypewriterPrison)
	}
	log.Printf("random returned: TypewriterPrisonSmileLetter")
	return p
}

func RandomValidTypewriterPrisonMenuPaperclip(r *rand.Rand) ProtoMessage {
	p := RandomTypewriterPrison(r).(*TypewriterPrison)
	for p.PocketRoses == nil || p.GetPocketRoses().MenuPaperclip == nil {
		log.Printf("random looped: TypewriterPrisonMenuPaperclip")
		p = RandomTypewriterPrison(r).(*TypewriterPrison)
	}
	log.Printf("random returned: TypewriterPrisonMenuPaperclip")
	return p
}

func randStringWithoutA(r *rand.Rand) string {
	v25 := r.Intn(100)
	tmps := make([]rune, v25)
	for i := 0; i < v25; i++ {
		tmps[i] = randUTF8Rune(r)
		for tmps[i] == 'a' {
			tmps[i] = randUTF8Rune(r)
		}
	}
	return string(tmps)
}

func RandomInvalidTypewriterPrisonMenuPaperclip(r *rand.Rand) ProtoMessage {
	p := RandomTypewriterPrison(r).(*TypewriterPrison)
	for i := range p.GetPocketRoses().GetMenuPaperclip() {
		log.Printf("random looped: TypewriterPrisonMenuPaperclip")
		p.PocketRoses.MenuPaperclip[i] = randStringWithoutA(r)
	}
	log.Printf("random returned: TypewriterPrisonMenuPaperclip")
	return p
}

func RandomValidTypewriterPrisonMapShark(r *rand.Rand) ProtoMessage {
	p := RandomTypewriterPrison(r).(*TypewriterPrison)
	for p.PocketRoses == nil || p.GetPocketRoses().MapShark == nil || !strings.Contains(p.GetPocketRoses().GetMapShark(), "a") {
		log.Printf("random looped: TypewriterPrisonMapShark")
		p = RandomTypewriterPrison(r).(*TypewriterPrison)
	}
	log.Printf("random returned: TypewriterPrisonMapShark")
	return p
}

func RandomInvalidTypewriterPrisonMapShark(r *rand.Rand) ProtoMessage {
	p := RandomTypewriterPrison(r).(*TypewriterPrison)
	for strings.Contains(p.GetPocketRoses().GetMapShark(), "a") {
		log.Printf("random looped: TypewriterPrisonMapShark")
		p = RandomTypewriterPrison(r).(*TypewriterPrison)
	}
	log.Printf("random returned: TypewriterPrisonMapShark")
	return p
}

func init() {
	var scarBusStop = G{
		"main": InOrder(
			Any(),
			In("PocketRoses",
				Any(),
				In("ScarBusStop", Value(Contains(StringVar(), StringConst("a")))),
				Any(),
			),
			Any(),
		),
	}
	BenchValidateProtoJson("TypewriterPrisonScarBusStop", scarBusStop, RandomValidTypewriterPrisonScarBusStop, RandomInvalidTypewriterPrisonScarBusStop)
	ValidateProto("TypewriterPrisonScarBusStop", scarBusStop, &TypewriterPrison{PocketRoses: &PocketRoses{ScarBusStop: proto.String("a")}}, true)

	var daisySled = G{
		"main": InOrder(
			Any(),
			In("PocketRoses",
				Any(),
				In("DaisySled", Value(Eq(IntVar(), IntConst(1)))),
				Any(),
			),
			Any(),
		),
	}
	BenchValidateProtoJson("TypewriterPrisonDaisySled", daisySled, RandomValidTypewriterPrisonDaisySled, RandomInvalidTypewriterPrisonDaisySled)
	ValidateProto("TypewriterPrisonDaisySled", daisySled, &TypewriterPrison{PocketRoses: &PocketRoses{DaisySled: proto.Int64(1)}}, true)

	var smileLetter = G{
		"main": InOrder(
			Any(),
			In("PocketRoses",
				Any(),
				In("SmileLetter", Value(BoolVar())),
				Any(),
			),
			Any(),
		),
	}
	BenchValidateProtoJson("TypewriterPrisonSmileLetter", smileLetter, RandomValidTypewriterPrisonSmileLetter, RandomInvalidTypewriterPrisonSmileLetter)
	ValidateProto("TypewriterPrisonSmileLetter", smileLetter, &TypewriterPrison{PocketRoses: &PocketRoses{SmileLetter: proto.Bool(true)}}, true)

	var menuPaperClip = G{
		"main": InOrder(
			Any(),
			In("PocketRoses",
				Any(),
				In("MenuPaperclip",
					InAny(Value(Contains(StringVar(), StringConst("a")))),
					Any(),
				),
				Any(),
			),
			Any(),
		),
	}
	BenchValidateProtoJson("TypewriterPrisonMenuPaperclip", menuPaperClip, RandomValidTypewriterPrisonMenuPaperclip, RandomInvalidTypewriterPrisonMenuPaperclip)
	ValidateProto("TypewriterPrisonMenuPaperclip", menuPaperClip, &TypewriterPrison{PocketRoses: &PocketRoses{MenuPaperclip: []string{"a"}}}, true)

	var mapShark = G{
		"main": InOrder(
			Any(),
			In("PocketRoses",
				Any(),
				In("MapShark", Value(Contains(StringVar(), StringConst("a")))),
				Any(),
			),
			Any(),
		),
	}
	BenchValidateProtoJson("TypewriterPrisonMapShark", mapShark, RandomValidTypewriterPrisonMapShark, RandomInvalidTypewriterPrisonMapShark)
	ValidateProto("TypewriterPrisonMapShark", mapShark, &TypewriterPrison{PocketRoses: &PocketRoses{MapShark: proto.String("a")}}, true)

}
