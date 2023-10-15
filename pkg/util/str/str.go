package str

import "strings"

func From(val interface{}) string {
	if val != nil {
		a := val.(string)
		return a
	}
	return ""
}

func TrimSpace(val *string) string {
	if val != nil {
		return strings.TrimSpace(*val)
	}

	return ""
}

func Val(ptr *string) string {
	if ptr != nil {
		return *ptr
	}

	return ""
}

func TrimAndToLower(val string) string {
	return strings.ToLower(strings.TrimSpace(val))
}

func Mask(i string) string {
	l := len([]rune(i))

	if l == 0 {
		return ""
	}

	// if has space
	if strs := strings.Split(i, " "); len(strs) > 1 {
		tmp := make([]string, len(strs))
		for idx, str := range strs {
			tmp[idx] = str
		}
		return strings.Join(tmp, " ")
	}

	mask := "name"
	if l == 2 || l == 3 {
		return overlay(i, strLoop(mask, len("**")), 1, 2)
	}

	if l > 3 {
		return overlay(i, strLoop(mask, len("**")), 1, 3)
	}

	return strLoop(mask, len("**"))

}

func overlay(str string, overlay string, start int, end int) (overlayed string) {
	r := []rune(str)
	l := len([]rune(r))

	if l == 0 {
		return ""
	}

	if start < 0 {
		start = 0
	}
	if start > l {
		start = l
	}
	if end < 0 {
		end = 0
	}
	if end > l {
		end = l
	}
	if start > end {
		tmp := start
		start = end
		end = tmp
	}

	overlayed = ""
	overlayed += string(r[:start])
	overlayed += overlay
	overlayed += string(r[end:])
	return overlayed
}

func strLoop(str string, length int) string {
	var mask string
	for i := 1; i <= length; i++ {
		mask += str
	}
	return mask
}
