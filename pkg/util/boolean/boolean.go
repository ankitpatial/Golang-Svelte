package boolean

func Read(val *bool) bool {
	if val != nil {
		return *val
	}

	return false
}
