package iteration

func Repeat(ch string, times int) string {
	retstr := ""
	for t := 0; t < times; t++ {
		retstr = retstr + ch
	}
	return retstr
}
