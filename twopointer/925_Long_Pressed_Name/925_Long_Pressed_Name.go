package twopointer

/*
Input: name = "alex", typed = "aaleex"
Output: true
Explanation: 'a' and 'e' in 'alex' were long pressed.

Input: name = "saeed", typed = "ssaaedd"
Output: false
Explanation: 'e' must have been pressed twice, but it wasn't in the typed output.

*/

func isLongPressedName(name string, typed string) bool {
	if len(name) == 0 || len(typed) == 0 {
		return false
	}

	//Input: name = "saeed", typed = "ssaaedd"

	nameIndex, typedIndex := 0, 0
	for typedIndex < len(typed) && nameIndex < len(name) {
		if typed[typedIndex] != name[nameIndex] {
			return false
		}

		for nameIndex < len(name) && typedIndex < len(typed) && name[nameIndex] == typed[typedIndex] {
			nameIndex++
			typedIndex++
		}

		for typedIndex < len(typed) && typed[typedIndex] == typed[typedIndex-1] {
			//typed 字串重複往前走
			typedIndex++
		}

	}

	return typedIndex == len(typed) && len(name) == nameIndex
}
