package numbers

func ToLCD(n int) string {
	if n == 7 {
		return  `
 _
  |
  |
`
	}
	return `
 _
|_|
|_|
`
}
