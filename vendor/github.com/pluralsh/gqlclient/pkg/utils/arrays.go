package utils

func ConvertStringArrayPointer(inputs []*string) []string {
	res := make([]string, 0)
	if inputs == nil {
		return res
	}

	for _, input := range inputs {
		res = append(res, *input)
	}

	return res
}
