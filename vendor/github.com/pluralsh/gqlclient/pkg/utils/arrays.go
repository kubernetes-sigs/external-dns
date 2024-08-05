package utils

func ConvertStringArrayPointer(inputs []*string) []string {
<<<<<<< HEAD
	res := make([]string, 0)
	if inputs == nil {
		return res
	}

	for _, input := range inputs {
		res = append(res, *input)
	}

||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
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
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	return res
}
