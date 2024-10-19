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
	"github.com/katydid/validator-go/validator/ast"
	. "github.com/katydid/validator-go/validator/combinator"
)

func init() {
	basicNone := G{"main": ast.NewNot(ast.NewZAny())}
	ValidateXMLString(
		"BasicNone_A",
		basicNone,
		"<A/>",
		false,
	)

	basicA := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewEmpty())}
	ValidateXMLString(
		"BasicA_A",
		basicA,
		"<A/>",
		true,
	)
	ValidateXMLString(
		"BasicA_B",
		basicA,
		"<B/>",
		false,
	)

	basicNotA := G{"main": ast.NewNot(ast.NewTreeNode(ast.NewStringName("A"), ast.NewEmpty()))}
	ValidateXMLString(
		"BasicNotA_A",
		basicNotA,
		"<A/>",
		false,
	)
	ValidateXMLString(
		"BasicNotA_B",
		basicNotA,
		"<B/>",
		true,
	)

	basicAB := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()))}
	ValidateXMLString(
		"BasicAB_AB",
		basicAB,
		"<A><B/></A>",
		true,
	)
	ValidateXMLString(
		"BasicAB_BB",
		basicAB,
		"<B><B/></B>",
		false,
	)

	basicALeafB := G{"main": ast.NewTreeNode(ast.NewStringName("A"), Value(Eq(StringConst("B"), StringVar())))}
	ValidateXMLString(
		"BasicALeafB_AB",
		basicALeafB,
		"<A>B</A>",
		true,
	)
	ValidateXMLString(
		"BasicALeafB_BB",
		basicALeafB,
		"<B>B</B>",
		false,
	)

	basicConcatBC := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewConcat(
		ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()),
		ast.NewTreeNode(ast.NewStringName("C"), ast.NewEmpty()),
	))}
	ValidateXMLString(
		"BasicConcatBC_BC",
		basicConcatBC,
		"<A><B/><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicConcatBC_BB",
		basicConcatBC,
		"<A><B/><B/></A>",
		false,
	)

	basicNotConcatBC := G{"main": ast.NewNot(ast.NewTreeNode(ast.NewStringName("A"), ast.NewConcat(
		ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()),
		ast.NewTreeNode(ast.NewStringName("C"), ast.NewEmpty()),
	)))}
	ValidateXMLString(
		"BasicNotConcatBC_BC",
		basicNotConcatBC,
		"<A><B/><C/></A>",
		false,
	)
	ValidateXMLString(
		"BasicNotConcatBC_BB",
		basicNotConcatBC,
		"<A><B/><B/></A>",
		true,
	)

	basicAorB := G{"main": ast.NewOr(
		ast.NewTreeNode(ast.NewStringName("A"), ast.NewEmpty()),
		ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()),
	)}
	ValidateXMLString(
		"BasicAorB_A",
		basicAorB,
		"<A/>",
		true,
	)

	ValidateXMLString(
		"BasicAorB_C",
		basicAorB,
		"<C/>",
		false,
	)

	basicTreeAAorBB := G{"main": ast.NewOr(
		ast.NewTreeNode(ast.NewStringName("A"), ast.NewTreeNode(ast.NewStringName("A"), ast.NewEmpty())),
		ast.NewTreeNode(ast.NewStringName("B"), ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty())),
	)}
	ValidateXMLString(
		"BasicTreeAAorBB_AA",
		basicTreeAAorBB,
		"<A><A/></A>",
		true,
	)
	ValidateXMLString(
		"BasicTreeAAorBB_AB",
		basicTreeAAorBB,
		"<A><B/></A>",
		false,
	)

	basicTreeBAorBB := G{"main": ast.NewOr(
		ast.NewTreeNode(ast.NewStringName("B"), ast.NewTreeNode(ast.NewStringName("A"), ast.NewEmpty())),
		ast.NewTreeNode(ast.NewStringName("B"), ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty())),
	)}
	ValidateXMLString(
		"BasicTreeBAorBB_BA",
		basicTreeBAorBB,
		"<B><A/></B>",
		true,
	)
	ValidateXMLString(
		"BasicTreeBAorBB_AA",
		basicTreeBAorBB,
		"<A><A/></A>",
		false,
	)

	basicTreeAOrOrC := G{"main": ast.NewOr(
		ast.NewTreeNode(ast.NewStringName("A"),
			ast.NewOr(
				ast.NewTreeNode(ast.NewStringName("A"), ast.NewEmpty()),
				ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()),
			),
		),
		ast.NewTreeNode(ast.NewStringName("C"),
			ast.NewOr(
				ast.NewTreeNode(ast.NewStringName("C"), ast.NewEmpty()),
				ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()),
			),
		),
	)}
	ValidateXMLString(
		"BasicTreeAOrOrC_AB",
		basicTreeAOrOrC,
		"<A><B/></A>",
		true,
	)
	ValidateXMLString(
		"BasicTreeAOrOrC_CA",
		basicTreeAOrOrC,
		"<C><A/></C>",
		false,
	)

	basicConcatZAnyC := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewConcat(
		ast.NewZAny(),
		ast.NewTreeNode(ast.NewStringName("C"), ast.NewEmpty()),
	))}
	ValidateXMLString(
		"BasicConcatZAnyC_AC",
		basicConcatZAnyC,
		"<A><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicConcatZAnyC_ABC",
		basicConcatZAnyC,
		"<A><B/><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicConcatZAnyC_BBBC",
		basicConcatZAnyC,
		"<A><B/><B/><B/><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicConcatZAnyC_BCCC",
		basicConcatZAnyC,
		"<A><B/><C/><C/><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicConcatZAnyC_BCBC",
		basicConcatZAnyC,
		"<A><B/><C/><B/><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicConcatZAnyC_AB",
		basicConcatZAnyC,
		"<A><B/></A>",
		false,
	)
	//typical fundemental flaw
	ValidateXMLString(
		"BasicConcatZAnyC_ACchildB_TypicalFundementalFlaw",
		basicConcatZAnyC,
		"<A><C>B</C></A>",
		false,
	)

	basicZeroOrMoreB0 := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewZeroOrMore(
		ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()),
	))}
	ValidateXMLString(
		"BasicZeroOrMoreB_0",
		basicZeroOrMoreB0,
		"<A></A>",
		true,
	)
	ValidateXMLString(
		"BasicZeroOrMoreB_1",
		basicZeroOrMoreB0,
		"<A><B/></A>",
		true,
	)
	ValidateXMLString(
		"BasicZeroOrMoreB_3",
		basicZeroOrMoreB0,
		"<A><B/><B/><B/></A>",
		true,
	)
	ValidateXMLString(
		"BasicZeroOrMoreB_C",
		basicZeroOrMoreB0,
		"<A><C/></A>",
		false,
	)
	ValidateXMLString(
		"BasicZeroOrMoreB_BC",
		basicZeroOrMoreB0,
		"<A><B/><C/><B/></A>",
		false,
	)

	basicZeroOrMoreEmpty := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewZeroOrMore(
		ast.NewEmpty(),
	))}
	ValidateXMLString(
		"BasicZeroOrMoreEmpty_Empty",
		basicZeroOrMoreEmpty,
		"<A></A>",
		true,
	)
	ValidateXMLString(
		"BasicZeroOrMoreEmpty_B",
		basicZeroOrMoreEmpty,
		"<A><B/></A>",
		false,
	)

	basicZeroOrMoreZeroOrMoreB := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewZeroOrMore(
		ast.NewZeroOrMore(ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty())),
	))}
	ValidateXMLString(
		"BasicZeroOrMoreZeroOrMoreB_BB",
		basicZeroOrMoreZeroOrMoreB,
		"<A><B/><B/></A>",
		true,
	)
	ValidateXMLString(
		"BasicZeroOrMoreZeroOrMoreB_C",
		basicZeroOrMoreZeroOrMoreB,
		"<A><C/></A>",
		false,
	)

	basicConcatOrEmpty := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewConcat(
		ast.NewOr(
			ast.NewEmpty(),
			ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()),
		),
		ast.NewTreeNode(ast.NewStringName("C"), ast.NewEmpty()),
	))}
	ValidateXMLString(
		"BasicConcatOrEmpty_BC",
		basicConcatOrEmpty,
		"<A><B/><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicConcatOrEmpty_C",
		basicConcatOrEmpty,
		"<A><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicConcatOrEmpty_BD",
		basicConcatOrEmpty,
		"<A><B/><D/></A>",
		false,
	)

	basicZeroOrMoreBOrEmpty := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewZeroOrMore(
		ast.NewOr(
			ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()),
			ast.NewEmpty(),
		),
	))}
	ValidateXMLString(
		"BasicZeroOrMoreBOrEmpty_BB",
		basicZeroOrMoreBOrEmpty,
		"<A><B/><B/></A>",
		true,
	)
	ValidateXMLString(
		"BasicZeroOrMoreBOrEmpty_BC",
		basicZeroOrMoreBOrEmpty,
		"<A><B/><C/></A>",
		false,
	)

	basicConcatCStar := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewConcat(
		ast.NewTreeNode(ast.NewStringName("C"), ast.NewEmpty()),
		ast.NewZeroOrMore(
			ast.NewTreeNode(ast.NewStringName("C"), ast.NewEmpty()),
		),
	))}
	ValidateXMLString(
		"BasicConcatCStar_0",
		basicConcatCStar,
		"<A></A>",
		false,
	)
	ValidateXMLString(
		"BasicConcatCStar_1",
		basicConcatCStar,
		"<A><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicConcatCStar_2",
		basicConcatCStar,
		"<A><C/><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicConcatCStar_3",
		basicConcatCStar,
		"<A><C/><C/><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicConcatCStar_B",
		basicConcatCStar,
		"<A><B/></A>",
		false,
	)
	ValidateXMLString(
		"BasicConcatCStar_CB",
		basicConcatCStar,
		"<A><C/><B/></A>",
		false,
	)

	basicTreeAandA := G{"main": ast.NewAnd(
		ast.NewTreeNode(ast.NewStringName("A"), ast.NewTreeNode(ast.NewStringName("A"), ast.NewEmpty())),
		ast.NewTreeNode(ast.NewStringName("A"), ast.NewTreeNode(ast.NewStringName("A"), ast.NewEmpty())),
	)}
	ValidateXMLString(
		"BasicTreeAandA_A",
		basicTreeAandA,
		"<A><A/></A>",
		true,
	)
	ValidateXMLString(
		"BasicTreeAandA_B",
		basicTreeAandA,
		"<A><B/></A>",
		false,
	)

	basicTreeAandB := G{"main": ast.NewAnd(
		ast.NewTreeNode(ast.NewStringName("A"), ast.NewTreeNode(ast.NewStringName("A"), ast.NewEmpty())),
		ast.NewTreeNode(ast.NewStringName("A"), ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty())),
	)}
	ValidateXMLString(
		"BasicTreeAandB_B",
		basicTreeAandB,
		"<A><B/></A>",
		false,
	)

	basicAndBAnyC := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewAnd(
		ast.NewConcat(
			ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()),
			ast.NewZAny(),
		),
		ast.NewConcat(
			ast.NewZAny(),
			ast.NewTreeNode(ast.NewStringName("C"), ast.NewEmpty()),
		),
	))}
	ValidateXMLString(
		"BasicAndBAnyC_BC",
		basicAndBAnyC,
		"<A><B/><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicAndBAnyC_CB",
		basicAndBAnyC,
		"<A><C/><B/></A>",
		false,
	)
	ValidateXMLString(
		"BasicAndBAnyC_B",
		basicAndBAnyC,
		"<A><B/></A>",
		false,
	)
	ValidateXMLString(
		"BasicAndBAnyC_C",
		basicAndBAnyC,
		"<A><C/></A>",
		false,
	)
	ValidateXMLString(
		"BasicAndBAnyC_BXXXC",
		basicAndBAnyC,
		"<A><B/><X/><X/><X/><C/></A>",
		true,
	)

	basicTreeAndBAnyC := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewAnd(
		ast.NewConcat(
			ast.NewTreeNode(ast.NewStringName("A"),
				ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()),
			),
			ast.NewZAny(),
		),
		ast.NewConcat(
			ast.NewZAny(),
			ast.NewTreeNode(ast.NewStringName("A"),
				ast.NewTreeNode(ast.NewStringName("C"), ast.NewEmpty()),
			),
		),
	))}
	ValidateXMLString(
		"BasicTreeAndBAnyC_BC",
		basicTreeAndBAnyC,
		"<A><A><B/></A><A><C/></A></A>",
		true,
	)
	ValidateXMLString(
		"BasicTreeAndBAnyC_CB",
		basicTreeAndBAnyC,
		"<A><A><C/></A><A><B/></A></A>",
		false,
	)
	ValidateXMLString(
		"BasicTreeAndBAnyC_BXXXC",
		basicTreeAndBAnyC,
		"<A><A><B/></A><X/><X/><X/><A><C/></A></A>",
		true,
	)
	ValidateXMLString(
		"BasicTreeAndBAnyC_CBC",
		basicTreeAndBAnyC,
		"<A><A><C/></A><A><B/></A><A><C/></A></A>",
		false,
	)

	basicAContainsB := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewContains(
		ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()),
	))}
	ValidateXMLString(
		"BasicAContainsB_B",
		basicAContainsB,
		"<A><B/></A>",
		true,
	)
	ValidateXMLString(
		"BasicAContainsB_CBC",
		basicAContainsB,
		"<A><C/><B/><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicAContainsB_CC",
		basicAContainsB,
		"<A><C/><C/></A>",
		false,
	)
	ValidateXMLString(
		"BasicAContainsB_0",
		basicAContainsB,
		"<A></A>",
		false,
	)

	basicOptionalB := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewOptional(
		ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()),
	))}
	ValidateXMLString(
		"BasicOptionalB_Empty",
		basicOptionalB,
		"<A></A>",
		true,
	)
	ValidateXMLString(
		"BasicOptionalB_B",
		basicOptionalB,
		"<A><B/></A>",
		true,
	)
	ValidateXMLString(
		"BasicOptionalB_BB",
		basicOptionalB,
		"<A><B/><B/></A>",
		false,
	)
	ValidateXMLString(
		"BasicOptionalB_C",
		basicOptionalB,
		"<A><C/></A>",
		false,
	)

	basicInterleaveBC := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewInterleave(
		ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()),
		ast.NewTreeNode(ast.NewStringName("C"), ast.NewEmpty()),
	))}
	ValidateXMLString(
		"BasicInterleaveBC_BC",
		basicInterleaveBC,
		"<A><B/><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicInterleaveBC_CB",
		basicInterleaveBC,
		"<A><C/><B/></A>",
		true,
	)
	ValidateXMLString(
		"BasicInterleaveBC_C",
		basicInterleaveBC,
		"<A><C/></A>",
		false,
	)

	basicInterleaveBAnyC := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewInterleave(
		ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()),
		ast.NewInterleave(
			ast.NewZAny(),
			ast.NewTreeNode(ast.NewStringName("C"), ast.NewEmpty()),
		),
	))}
	ValidateXMLString(
		"BasicInterleaveBAnyC_BC",
		basicInterleaveBAnyC,
		"<A><B/><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicInterleaveBAnyC_BAC",
		basicInterleaveBAnyC,
		"<A><B/><A/><C/></A>",
		true,
	)
	ValidateXMLString(
		"BasicInterleaveBAnyC_ABACA",
		basicInterleaveBAnyC,
		"<A><A/><B/><A/><C/><A/></A>",
		true,
	)
	ValidateXMLString(
		"BasicInterleaveBAnyC_ABAA",
		basicInterleaveBAnyC,
		"<A><A/><B/><A/><A/></A>",
		false,
	)
	ValidateXMLString(
		"BasicInterleaveBAnyC_ACCBA",
		basicInterleaveBAnyC,
		"<A><A/><C/><C/><B/><A/></A>",
		true,
	)
	ValidateXMLString(
		"BasicInterleaveBAnyC_ACCCA",
		basicInterleaveBAnyC,
		"<A><A/><C/><C/><C/><A/></A>",
		false,
	)

	basicRefLoop := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewOr(
		ast.NewEmpty(),
		ast.NewReference("main"),
	))}
	ValidateXMLString(
		"BasicRefLoop_A",
		basicRefLoop,
		"<A/>",
		true,
	)
	ValidateXMLString(
		"BasicRefLoop_AA",
		basicRefLoop,
		"<A><A/></A>",
		true,
	)
	ValidateXMLString(
		"BasicRefLoop_AB",
		basicRefLoop,
		"<A><B/></A>",
		false,
	)

	basicConcatBOptionalD := G{"main": ast.NewTreeNode(ast.NewStringName("A"),
		ast.NewConcat(
			ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()),
			ast.NewOptional(ast.NewTreeNode(ast.NewStringName("D"), ast.NewEmpty())),
		),
	)}
	ValidateXMLString(
		"BasicConcatBOptionalD_BD",
		basicConcatBOptionalD,
		"<A><B/><D/></A>",
		true,
	)
	ValidateXMLString(
		"BasicConcatBOptionalD_B",
		basicConcatBOptionalD,
		"<A><B/></A>",
		true,
	)
	ValidateXMLString(
		"BasicConcatBOptionalD_D",
		basicConcatBOptionalD,
		"<A><D/></A>",
		false,
	)

	//!(B:<empty>) can be <empty> and * can be B:<empty> which means B:<empty> can be accepted by [*, !(B:<empty>)]
	basicAnyNotB := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewConcat(
		ast.NewZAny(),
		ast.NewNot(ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty())),
	))}
	ValidateXMLString(
		"BasicAnyNotB_B",
		basicAnyNotB,
		"<A><B/></A>",
		true,
	)
	ValidateXMLString(
		"BasicAnyNotB_C",
		basicAnyNotB,
		"<A><C/></A>",
		true,
	)

	basicNotAndBStarC := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewNot(
		ast.NewAnd(
			ast.NewConcat(ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty()), ast.NewZAny()),
			ast.NewConcat(ast.NewZAny(), ast.NewTreeNode(ast.NewStringName("C"), ast.NewEmpty())),
		),
	))}
	ValidateXMLString(
		"BasicNotAndBStarC_BC",
		basicNotAndBStarC,
		"<A><B/><C/></A>",
		false,
	)
	ValidateXMLString(
		"BasicNotAndBStarC_CB",
		basicNotAndBStarC,
		"<A><C/><B/></A>",
		true,
	)

	basicAndNotAB := G{"main": ast.NewAnd(
		ast.NewNot(ast.NewTreeNode(ast.NewStringName("A"), ast.NewEmpty())),
		ast.NewNot(ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty())),
	)}
	ValidateXMLString(
		"BasicAndNotAB_A",
		basicAndNotAB,
		"<A/>",
		false,
	)
	ValidateXMLString(
		"BasicAndNotAB_B",
		basicAndNotAB,
		"<B/>",
		false,
	)
	ValidateXMLString(
		"BasicAndNotAB_C",
		basicAndNotAB,
		"<C/>",
		true,
	)

	basicOrNotAB := G{"main": ast.NewOr(
		ast.NewNot(ast.NewTreeNode(ast.NewStringName("A"), ast.NewEmpty())),
		ast.NewNot(ast.NewTreeNode(ast.NewStringName("B"), ast.NewEmpty())),
	)}
	ValidateXMLString(
		"BasicOrNotAB_A",
		basicOrNotAB,
		"<A/>",
		true,
	)
	ValidateXMLString(
		"BasicOrNotAB_C",
		basicOrNotAB,
		"<C/>",
		true,
	)

	//deeper fundamental flaw
	basicAEndsWithBContainsAnyD := G{"main": ast.NewTreeNode(ast.NewStringName("A"), ast.NewConcat(
		ast.NewZAny(),
		ast.NewTreeNode(ast.NewStringName("B"),
			ast.NewContains(ast.NewTreeNode(ast.NewAnyName(), ast.NewTreeNode(ast.NewStringName("D"), ast.NewEmpty()))),
		)),
	)}
	ValidateXMLString(
		"BasicAEndsWithBContainsAnyD_BCD_DeeperFundementalFlaw",
		basicAEndsWithBContainsAnyD,
		"<A><B><C><D/></C></B></A>",
		true,
	)
	ValidateXMLString(
		"BasicAEndsWithBContainsAnyD_BCA_DeeperFundementalFlaw",
		basicAEndsWithBContainsAnyD,
		"<A><B><C><A/></C></B></A>",
		false,
	)

	basicAndContainsTree := G{"main": In("A", AllOf(
		InPath("B", In("C", ast.NewEmpty())),
		InPath("B", In("D", ast.NewEmpty())),
	))}
	ValidateXMLString(
		"BasicAndContainsTree_BCBD",
		basicAndContainsTree,
		"<A><B><C/></B><B><D/></B></A>",
		true,
	)
	ValidateXMLString(
		"BasicAndContainsTree_BC",
		basicAndContainsTree,
		"<A><B><C/></B></A>",
		false,
	)
}
