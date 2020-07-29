package goco

/*
//headers: https://dmalcolm.fedorapeople.org/gcc/2015-08-31/rst-experiment/gnu-objective-c-runtime-api.html
*/

func BoolToBOOL(in bool) uint8 {
	if in {
		return 1
	}
	return 0
}

func IntToBool(in int) bool {
	if in == 1 {
		return true
	}
	return false
}
