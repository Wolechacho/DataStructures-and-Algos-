package main

func restoreString(s string, indices []int) string {
	res := []byte(s)
	for i := 0; i < len(s); i++ {
		res[indices[i]] = s[i]
	}
	return string(res)
}
