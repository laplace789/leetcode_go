package main

/*

Input: licensePlate = "1s3 PSt", words = ["step","steps","stripe","stepple"]
Output: "steps"

Input: licensePlate = "1s3 456", words = ["looks","pest","stew","show"]
Output: "pest"

Input: licensePlate = "Ah71752", words = ["suggest","letter","of","husband","easy","education","drug","prevent","writer","old"]
Output: "husband"

*/
func shortestCompletingWord(licensePlate string, words []string) string {
	//parser licensePlate
	//lowercase 97-122
	//uppercase 65-90

	//charArr := []rune(licensePlate)
	letters := make([]uint8, 0)
	//matchArr := make([]map[bool]int,0)
	for i := 0; i < len(licensePlate); i++ {
		//a-z
		if licensePlate[i] >= 65 && licensePlate[i] <= 90 {
			letters = append(letters, licensePlate[i]+32)
		} else if licensePlate[i] >= 97 && licensePlate[i] <= 122 {
			//A-Z
			letters = append(letters, licensePlate[i])
		}
	}

	for i := 0; i < len(words); i++ {

	}
	return ""

}

func removeIndex(s []uint8, index int) []uint8 {
	return append(s[:index], s[index+1:]...)
}

func main() {
	shortestCompletingWord("1s3 PSt", []string{})
}
