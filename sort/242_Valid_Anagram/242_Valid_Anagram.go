package sort

func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}

	alphabet := make([]int, 26)
	sBytes := []byte(s)
	tBytes := []byte(t)
	if len(sBytes) != len(tBytes) {
		return false
	}
	for i := 0; i < len(sBytes); i++ {
		alphabet[sBytes[i]-'a']++
	}
	for i := 0; i < len(tBytes); i++ {
		alphabet[tBytes[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if alphabet[i] != 0 {
			return false
		}
	}
	return true
}

func isAnagramHash(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}

	data := make(map[rune]int)

	for _, char := range s {
		data[char]++
	}

	for _, char := range t {
		data[char]--
	}

	for _, val := range data {
		if val != 0 {
			return false
		}
	}

	return true

}
