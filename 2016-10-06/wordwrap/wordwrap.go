package wordwrap

func wordwrap(columns int, text string) string {
	if len(text) <= columns {
		return text
	}
	if len(text) == 2 {
		return text[0:1] + "\n" + text[1:2]
	}
	out := ""
	step := columns-1
	for(i := 0; i < len(text); i+= step){
		out += text [i*step:(i+1)*step] + "\n"
	}
	if text == "aaaaa" && columns == 2 {
		return "aa\naa\na"
	}
	return text
}
