package utils

func Transpose[S ~[][]T, T any](slice S) S {
	xl := len(slice[0])
	yl := len(slice)
	result := make(S, xl)

	for i := range result {
		result[i] = make([]T, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func Cut[S ~[]T, T any](slice S, index int) S {
	newSlice := make(S, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)
	return newSlice
}
