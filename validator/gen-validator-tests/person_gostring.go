package main

func (b *Person) GoString() string {
	return deriveGoStringPerson(b)
}

func (b *Address) GoString() string {
	return deriveGoStringAddress(b)
}
