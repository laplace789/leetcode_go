package stack

/*

Input: s = "()"
Output: true

Input: s = "()[]{}"
Output: true

Input: s = "{[]}"
Output: true


*/

func isValid(s string) bool {
	if s == "" {
		return true
	}
	arr := make([]rune, 0)

	for _, character := range s {
		if (character == '[') || (character == '(') || (character == '{') {
			arr = append(arr, character)
		} else {
			//stack empty
			if len(arr) == 0 {
				return false
			}
			top := arr[len(arr)-1]
			switch character {
			case '}':
				{
					if top != '{' {
						return false
					}
				}
			case ']':
				{
					if top != '[' {
						return false
					}
				}
			case ')':
				{
					if top != '(' {
						return false
					}
				}
			}
			arr = arr[:len(arr)-1]
		}
	}

	return len(arr) == 0
}
