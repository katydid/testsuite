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

	"github.com/gogo/protobuf/proto"
)

func RandomValidContextPerson(r *rand.Rand) ProtoMessage {
	p := RandomPerson(r).(*Person)
	for len(p.Addresses) == 0 {
		log.Printf("random looped: ContextPerson")
		p = RandomPerson(r).(*Person)
	}
	index := rand.Intn(len(p.Addresses))
	p.Addresses[index].Number = proto.Int64(456)
	p.Addresses[index].Street = proto.String("TheStreet")
	log.Printf("random returned: ContextPerson")
	return p
}

func RandomValidListIndexAddressPerson(r *rand.Rand) ProtoMessage {
	p := RandomPerson(r).(*Person)
	if len(p.Addresses) <= 1 {
		for i := 0; i < rand.Intn(10)+2; i++ {
			p.Addresses = append(p.Addresses, RandomAddress(r))
		}
	}
	p.Addresses[len(p.Addresses)-2].Number = proto.Int64(2)
	p.Addresses[len(p.Addresses)-1].Number = proto.Int64(1)
	return p
}

func RandomAddress(r *rand.Rand) *Address {
	return random(r, &Address{}).(*Address)
}

func RandomValidNilNamePerson(r *rand.Rand) ProtoMessage {
	p := RandomPerson(r).(*Person)
	p.Name = nil
	return p
}

func RandomValidLenNamePerson(r *rand.Rand) ProtoMessage {
	p := RandomPerson(r).(*Person)
	p.Name = proto.String("")
	return p
}

func RandomInvalidEmptyOrNilPerson(r *rand.Rand) ProtoMessage {
	p := RandomPerson(r).(*Person)
	for len(p.GetName()) == 0 {
		log.Printf("random looped: EmptyOrNilPerson")
		nonzero := randNonZeroString(r)
		p.Name = proto.String(nonzero)
	}
	log.Printf("random returned: EmptyOrNilPerson")
	return p
}

func RandomValidEmptyOrNilPerson(r *rand.Rand) ProtoMessage {
	p := RandomPerson(r).(*Person)
	if r.Intn(2) == 0 {
		p.Name = nil
	} else {
		p.Name = proto.String("")
	}
	return p
}

func RandomValidNaiveNotNamePerson(r *rand.Rand) ProtoMessage {
	p := RandomPerson(r).(*Person)
	if p.Name == nil {
		p.Name = proto.String(randString(r))
	}
	return p
}

func RandomInvalidNaiveNotNamePerson(r *rand.Rand) ProtoMessage {
	p := RandomPerson(r).(*Person)
	if r.Intn(2) == 0 {
		p.Name = nil
	} else {
		p.Name = proto.String("David")
	}
	return p
}

func RandomInvalidProperNotNamePerson(r *rand.Rand) ProtoMessage {
	p := RandomPerson(r).(*Person)
	p.Name = proto.String("David")
	return p
}

func RandomValidAndNameTelephonePerson(r *rand.Rand) ProtoMessage {
	p := RandomPerson(r).(*Person)
	p.Name = proto.String("David")
	p.Telephone = proto.String("0123456789")
	return p
}

func RandomValidOrNameTelephonePerson(r *rand.Rand) ProtoMessage {
	p := RandomPerson(r).(*Person)
	if r.Intn(2) == 0 {
		p.Name = proto.String("David")
	} else {
		p.Telephone = proto.String("0123456789")
	}
	return p
}

func init() {
	BenchValidateProtoJson("ContextPerson", ContextPerson, RandomValidContextPerson, RandomPerson)
	BenchValidateProtoJson("ListIndexAddress", ListIndexAddressPerson, RandomValidListIndexAddressPerson, RandomPerson)
	BenchValidateProtoJson("NilName", NilNamePerson, RandomValidNilNamePerson, RandomPerson)
	BenchValidateProtoJson("LenName", LenNamePerson, RandomValidLenNamePerson, RandomPerson)
	BenchValidateProtoJson("EmptyOrNil", EmptyOrNilPerson, RandomValidEmptyOrNilPerson, RandomInvalidEmptyOrNilPerson)
	BenchValidateProtoJson("IncorrectNotName", NaiveNotNamePerson, RandomValidNaiveNotNamePerson, RandomInvalidNaiveNotNamePerson)
	BenchValidateProtoJson("CorrectNotName", ProperNotNamePerson, RandomPerson, RandomInvalidProperNotNamePerson)
	BenchValidateProtoJson("AndNameTelephone", AndNameTelephonePerson, RandomValidAndNameTelephonePerson, RandomPerson)
	BenchValidateProtoJson("OrNameTelephone", OrNameTelephonePerson, RandomValidOrNameTelephonePerson, RandomPerson)
}
