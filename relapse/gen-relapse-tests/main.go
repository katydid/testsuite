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
	"encoding/json"
	"encoding/xml"
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
			createFolder(folder)

			v.Grammar.Format()
			writeFile(
				filepath.Join(folder, "relapse.txt"),
				[]byte(v.Grammar.String()),
			)

			writeFile(
				filepath.Join(folder, "relapse.json"),
				mustBytes(json.MarshalIndent(v.Grammar, "", "\t")),
			)

			writeFile(
				filepath.Join(folder, "relapse.xml"),
				mustBytes(xml.MarshalIndent(v.Grammar, "", "\t")),
			)

			bytesFilename := filepath.Join(folder, "invalid.dat")
			if v.Expected {
				bytesFilename = filepath.Join(folder, "valid.dat")
			}
			writeFile(bytesFilename, v.Bytes)
		}
	}
}

func createFolder(folder string) {
	if err := os.MkdirAll(folder, 0755); err != nil {
		log.Fatalf("error <%v> creating folder <%s>", err, folder)
	}
}

func writeFile(filename string, data []byte) {
	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		log.Fatalf("error <%v> writing file <%s>", err, filename)
	}
}
