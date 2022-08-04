package helper

func GenerateQuestionsMark(length int) []string {
	s := make([]string, 0)
	for i := 0; i < length; i++ {
		s = append(s, "?")
	}
	return s
}
