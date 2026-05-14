package utils

func ConvertStringArrayPointer(inputs []*string) []string {
	if inputs == nil {
		return []string{}
	}

	res := make([]string, len(inputs))
	for i, input := range inputs {
		res[i] = *input
	}
	return res
}

func ToStringArrayPtr(input []string) []*string {
	if input == nil {
		return []*string{}
	}

	res := make([]*string, len(input))
	for i := range input {
		res[i] = &input[i]
	}
	return res
}
