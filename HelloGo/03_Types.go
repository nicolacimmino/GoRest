package HelloGo

/******************
Types
*/
func GiveMeAnL() string {
	// Converting an int65 to an int8 truncates the higher bits
	var anInteger = int64(320)
	var an8BitInteger = int8(anInteger)
	intToSmallerTypeTruncates := an8BitInteger == 320&255

	// Converting a float32 to a int8 converts to int and truncates the higher bits
	var aFloat32 = float32(320.32)
	var another8BitInteger = int8(aFloat32)
	floatConversionToSmallerTypeTruncates := another8BitInteger == 320&255

	result := ""
	if intToSmallerTypeTruncates && floatConversionToSmallerTypeTruncates {
		result = "L"
	}
	return string(result)
}
