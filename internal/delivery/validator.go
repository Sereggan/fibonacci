package delivery

func ValidateInputParameters(from uint64, to uint64) bool {
	if from < 0 || to < 1 {
		return false
	}

	if !(to > from) {
		return false
	}

	return true
}
