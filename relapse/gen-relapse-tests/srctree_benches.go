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
	"math/rand"
)

func RandomValidRecursiveSrcTree(r *rand.Rand) ProtoMessage {
	pops := []*SrcTree{IoUtilSrcTree, PathSrcTree}
	return pops[r.Intn(2)]
	// p := RandomSrcTree(r).(*SrcTree)
	// for {
	// 	index := r.Intn(len(p.GetImports()) + 1)
	// 	if index == 0 {
	// 		p.PackageName = proto.String("io")
	// 		break
	// 	}
	// 	p = p.GetImports()[index-1]
	// }
	// return p
}

func RandomInvalidRecursiveSrcTree(r *rand.Rand) ProtoMessage {
	pops := []*SrcTree{RuntimeSrcTree, SyscallSrcTree}
	return pops[r.Intn(2)]
}

func init() {
	BenchValidateProtoJson("RecursiveSrcTree", RecursiveSrcTree, RandomValidRecursiveSrcTree, RandomInvalidRecursiveSrcTree)
}
