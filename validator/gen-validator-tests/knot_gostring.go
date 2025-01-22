package main

func (b *Knot) GoString() string {
	return deriveGoStringKnot(b)
}

func (b *BightKnot) GoString() string {
	return deriveGoStringBightKnot(b)
}
