package hex

func ByteToHexByte(in byte) string {
	b1 := in & 0xf0
	b1 = b1 >> 4
	b2 := in & 0x0f
	
	if(b1 + '0' > '9'){
		b1 += 87
	}else{
		b1 += 48
	}

	if(b2 + 48 > 57){
		b2 += 87
	}else{
		b2 += 48
	}
	// fmt.Printf("%08b\n", b2)
	return string(b1) + string(b2)
}
func BytesToText(input string) string {
	out := ""
	for _, r := range input {
		out += ByteToHexByte(byte(r))
	}
	return out
}
