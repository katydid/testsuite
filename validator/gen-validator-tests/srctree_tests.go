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
)

var IoUtilSrcTree = &SrcTree{
	PackageName: proto.String("io/ioutil"),
	Imports: []*SrcTree{
		{
			PackageName: proto.String("io"),
			Imports: []*SrcTree{
				{
					PackageName: proto.String("errors"),
				},
				{
					PackageName: proto.String("sync"),
				},
			},
		},
		{
			PackageName: proto.String("os"),
			Imports: []*SrcTree{
				{
					PackageName: proto.String("errors"),
				},
				{
					PackageName: proto.String("io"),
				},
				{
					PackageName: proto.String("runtime"),
				},
			},
		},
	},
}

var PathSrcTree = &SrcTree{
	PackageName: proto.String("path"),
	Imports: []*SrcTree{
		{
			PackageName: proto.String("errors"),
		},
		{
			PackageName: proto.String("strings"),
			Imports: []*SrcTree{
				{
					PackageName: proto.String("errors"),
				},
				{
					PackageName: proto.String("io"),
				},
				{
					PackageName: proto.String("uncode"),
				},
				{
					PackageName: proto.String("uncode/utf8"),
				},
			},
		},
		{
			PackageName: proto.String("unicode/utf8"),
		},
	},
}

var RuntimeSrcTree = &SrcTree{
	PackageName: proto.String("runtime"),
	Imports: []*SrcTree{
		{
			PackageName: proto.String("unsafe"),
		},
	},
}

var SyscallSrcTree = &SrcTree{
	PackageName: proto.String("syscall"),
	Imports: []*SrcTree{
		{
			PackageName: proto.String("errors"),
		},
		{
			PackageName: proto.String("runtime"),
		},
		{
			PackageName: proto.String("sync"),
			Imports: []*SrcTree{
				{
					PackageName: proto.String("sync/atomic"),
				},
				{
					PackageName: proto.String("unsafe"),
				},
			},
		},
		{
			PackageName: proto.String("unsafe"),
		},
	},
}

// Does this SrcTree depend on io or is its package name io
var RecursiveSrcTree = G{
	"main": AnyOf(
		InOrder(
			In("PackageName", Value(Eq(StringVar(), StringConst("io")))),
			Any(),
		),
		InOrder(
			Any(),
			In("Imports",
				Any(),
				InAny(Eval("main")),
				Any(),
			),
		),
	),
}

func init() {
	ValidateProtoEtc("RecursiveIoUtil", RecursiveSrcTree, IoUtilSrcTree, true)
	ValidateProtoEtc("RecursivePath", RecursiveSrcTree, PathSrcTree, true)
	ValidateProtoEtc("RecursiveRuntime", RecursiveSrcTree, RuntimeSrcTree, false)
	ValidateProtoEtc("RecursiveSyscall", RecursiveSrcTree, SyscallSrcTree, false)
}
