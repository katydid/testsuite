package main

func (this *TreeRegister) GoString() string {
	return deriveGoStringTreeRegister(this)
}
func (this *Tree) GoString() string {
	return deriveGoStringTree(this)
}
func (this *GeneralInformation) GoString() string {
	return deriveGoStringGeneralInformation(this)
}
func (this *OtherInformation) GoString() string {
	return deriveGoStringOtherInformation(this)
}
func (this *Measurements) GoString() string {
	return deriveGoStringMeasurements(this)
}
func (this *HeightMeasurement) GoString() string {
	return deriveGoStringHeightMeasurement(this)
}
func (this *Measurement) GoString() string {
	return deriveGoStringMeasurement(this)
}
func (this *Condition) GoString() string {
	return deriveGoStringCondition(this)
}
func (this *AdditionalInformation) GoString() string {
	return deriveGoStringAdditionalInformation(this)
}
func (this *Photo) GoString() string {
	return deriveGoStringPhoto(this)
}
func (this *Sender) GoString() string {
	return deriveGoStringSender(this)
}
