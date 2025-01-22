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

	. "github.com/katydid/validator-go/validator/combinator"
)

func RandomValidBridgePepper(r *rand.Rand) ProtoMessage {
	p := RandomPuddingMilkshake(r).(*PuddingMilkshake)
	for p.GetFinanceJudo().GetSaladWorry().GetSpyCarpenter().GetBridgePepper() == nil {
		log.Printf("random looped: BridgePepper")
		p = RandomPuddingMilkshake(r).(*PuddingMilkshake)
	}
	log.Printf("random returned: BridgePepper")
	return p
}

func RandomInvalidBridgePepper(r *rand.Rand) ProtoMessage {
	p := RandomPuddingMilkshake(r).(*PuddingMilkshake)
	if ss := p.GetFinanceJudo().GetSaladWorry().GetSpyCarpenter().GetBridgePepper(); ss != nil {
		for i := range ss {
			ss[i] = randStringWithoutA(r)
		}
	}
	return p
}

func RandomValidBridgePepperAndFountainTarget(r *rand.Rand) ProtoMessage {
	p := RandomPuddingMilkshake(r).(*PuddingMilkshake)
	for len(p.GetFinanceJudo().GetSaladWorry().GetSpyCarpenter().GetBridgePepper()) == 0 ||
		len(p.GetFinanceJudo().GetSaladWorry().GetSpyCarpenter().GetFountainTarget()) == 0 {
		log.Printf("random looped: BridgePepperAndFountainTarget")
		p = RandomPuddingMilkshake(r).(*PuddingMilkshake)
	}
	// bp := p.GetFinanceJudo().GetSaladWorry().GetSpyCarpenter().GetBridgePepper()
	// ft := p.GetFinanceJudo().GetSaladWorry().GetSpyCarpenter().GetFountainTarget()
	// bpi := rand.Intn(len(bp))
	// fti := rand.Intn(len(ft))
	p.FinanceJudo.SaladWorry.SpyCarpenter = &SpyCarpenter{
		BridgePepper:   []string{"a"},
		FountainTarget: []string{"a"},
	}
	// p.FinanceJudo.SaladWorry.SpyCarpenter.BridgePepper[bpi] = "a"
	// p.FinanceJudo.SaladWorry.SpyCarpenter.FountainTarget[fti] = "a"
	log.Printf("random returned: BridgePepperAndFountainTarget")
	return p
}

func RandomInvalidBridgePepperAndFountainTarget(r *rand.Rand) ProtoMessage {
	p := RandomPuddingMilkshake(r).(*PuddingMilkshake)
	if r.Intn(2) == 0 {
		if ss := p.GetFinanceJudo().GetSaladWorry().GetSpyCarpenter().GetBridgePepper(); ss != nil {
			for i := range ss {
				ss[i] = randStringWithoutA(r)
			}
		}
	} else {
		if ss := p.GetFinanceJudo().GetSaladWorry().GetSpyCarpenter().GetFountainTarget(); ss != nil {
			for i := range ss {
				ss[i] = randStringWithoutA(r)
			}
		}
	}
	return p
}

func init() {
	var bridgePepper = G{
		"main": InOrder(
			Any(),
			In("FinanceJudo",
				Any(),
				In("SaladWorry",
					Any(),
					In("SpyCarpenter",
						Any(),
						In("BridgePepper", InAny(Value(Contains(StringVar(), StringConst("a"))))),
						Any(),
					),
					Any(),
				),
				Any(),
			),
			Any(),
		),
	}
	BenchValidateProtoJson("BridgePepper", bridgePepper, RandomValidBridgePepper, RandomInvalidBridgePepper)
	ValidateProto("BridgePepper", bridgePepper, &PuddingMilkshake{FinanceJudo: &FinanceJudo{SaladWorry: &SaladWorry{SpyCarpenter: &SpyCarpenter{BridgePepper: []string{"a"}}}}}, true)

	var bridgePepperAndFountainTarget = G{
		"main": InOrder(
			Any(),
			In("FinanceJudo",
				Any(),
				In("SaladWorry",
					Any(),
					In("SpyCarpenter",
						AllOf(
							InPath("BridgePepper", InAny(Value(Contains(StringVar(), StringConst("a"))))),
							InPath("FountainTarget", InAny(Value(Contains(StringVar(), StringConst("a"))))),
						),
					),
					Any(),
				),
				Any(),
			),
			Any(),
		),
	}
	BenchValidateProtoJson("BridgePepperAndFountainTarget", bridgePepperAndFountainTarget, RandomValidBridgePepperAndFountainTarget, RandomInvalidBridgePepperAndFountainTarget)
	ValidateProto("BenchBridgePepperAndFountainTarget", bridgePepperAndFountainTarget, &PuddingMilkshake{FinanceJudo: &FinanceJudo{SaladWorry: &SaladWorry{SpyCarpenter: &SpyCarpenter{BridgePepper: []string{"a"}, FountainTarget: []string{"a"}}}}}, true)
}
