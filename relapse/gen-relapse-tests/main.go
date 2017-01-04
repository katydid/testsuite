//  Copyright 2017 Walter Schulze
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
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()
	path := "."
	if len(flag.Args()) == 1 {
		path = flag.Args()[0]
	}
	for name, codecs := range Validators {
		for codecName, v := range codecs {
			folder := filepath.Join(filepath.Join(path, codecName), name)
			if err := os.MkdirAll(folder, 0755); err != nil {
				log.Fatalf("error <%v> creating folder <%s>", err, folder)
			}

			grammarFilename := filepath.Join(folder, "relapse.grammar")
			grammarBytes := []byte(v.Grammar.String())
			if err := ioutil.WriteFile(grammarFilename, grammarBytes, 0644); err != nil {
				log.Fatalf("error <%v> writing file <%s>", err, grammarFilename)
			}

			bytesFilename := filepath.Join(folder, "invalid.dat")
			if v.Expected {
				bytesFilename = filepath.Join(folder, "valid.dat")
			}
			if err := ioutil.WriteFile(bytesFilename, v.Bytes, 0644); err != nil {
				log.Fatalf("error <%v> writing file <%s>", err, bytesFilename)
			}
		}
	}
}
