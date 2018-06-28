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
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var seed = flag.Int64("seed", time.Now().UnixNano(), "seed for generating benchmarks")
var benches = flag.Bool("benches", false, "generate benchmarks")

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()
	path := "."
	if len(flag.Args()) == 1 {
		path = flag.Args()[0]
	}
	for _, v := range Validators {
		codecFolder := filepath.Join(filepath.Join(path, "tests"), v.CodecName)
		folder := filepath.Join(codecFolder, v.Name)
		createFolder(folder)

		v.Grammar.Format()
		writeFile(
			filepath.Join(folder, "relapse.txt"),
			[]byte(v.Grammar.String()),
		)
		if !strings.Contains(v.Name, "Larger") {
			writeFile(
				filepath.Join(folder, "relapse.json"),
				mustBytes(json.MarshalIndent(v.Grammar, "", "\t")),
			)
			writeFile(
				filepath.Join(folder, "relapse.xml"),
				mustBytes(xml.MarshalIndent(v.Grammar, "", "\t")),
			)
		}

		bytesFilename := filepath.Join(folder, "invalid."+v.Extension)
		if v.Expected {
			bytesFilename = filepath.Join(folder, "valid."+v.Extension)
		}
		writeFile(bytesFilename, v.Bytes)

		schemaFilename := filepath.Join(codecFolder, v.SchemaName)
		if len(v.SchemaName) > 0 && notExists(schemaFilename) {
			s := Schemas[v.SchemaName]
			writeFile(
				schemaFilename,
				s.Data(),
			)
		}
	}

	if !*benches {
		return
	}

	for _, v := range BenchValidators {
		codecFolder := filepath.Join(filepath.Join(path, "benches"), v.CodecName)
		folder := filepath.Join(codecFolder, v.Name)
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

		r := rand.New(rand.NewSource(*seed))

		for i := 0; i < 500; i++ {
			bytes := v.ValidBytes(r)
			for !v.Validate(bytes) {
				log.Printf("generated invalid: %s - %s", v.Name, v.CodecName)
				bytes = v.ValidBytes(r)
			}
			bytesFilename := filepath.Join(folder, "valid_"+strconv.Itoa(i)+"."+v.Extension)
			writeFile(bytesFilename, bytes)
		}
		for i := 0; i < 500; i++ {
			bytes := v.InvalidBytes(r)
			for v.Validate(bytes) {
				log.Printf("generated valid, wanted invalid: %s - %s", v.Name, v.CodecName)
				bytes = v.InvalidBytes(r)
			}
			bytesFilename := filepath.Join(folder, "invalid_"+strconv.Itoa(i)+"."+v.Extension)
			writeFile(bytesFilename, bytes)
		}

		schemaFilename := filepath.Join(codecFolder, v.SchemaName)
		if len(v.SchemaName) > 0 && notExists(schemaFilename) {
			s := Schemas[v.SchemaName]
			writeFile(
				schemaFilename,
				s.Data(),
			)
		}
	}
}

func createFolder(folder string) {
	if err := os.MkdirAll(folder, 0755); err != nil {
		log.Fatalf("error <%v> creating folder <%s>", err, folder)
	}
}

func writeFile(filename string, data []byte) {
	if err := writeFileWithSync(filename, data, 0644); err != nil {
		log.Fatalf("error <%v> writing file <%s>", err, filename)
	}
	time.Sleep(10 * time.Millisecond)
}

// writeFileWithSync adds the os.O_SYNC flag, otherwise we get the error: too many open files in system.
func writeFileWithSync(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_SYNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		return io.ErrShortWrite
	}
	return f.Close()
}

func notExists(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsNotExist(err)
}
