import "strconv"

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var encoded_str string

	for _, str := range strs {
		rune_len := len([]rune(str))
		encoded_str += strconv.Itoa(rune_len) + "#" + str

	}

	return encoded_str
}

func (s *Solution) Decode(encoded string) []string {

	encoded_rune := []rune(encoded) 
	res := []string{}
	
	for i := 0; i < len(encoded_rune) {

		var encoded_len_str string 
		for encoded_rune[i] != '#' {
		 encoded_len_str += string(encoded_rune[i])
			i++
		}

		length, _ := strconv.Atoi(encoded_len_str)
		i ++ 

		res = append(res, string(encoded_rune[i : i + length]))
		i += length

	}
	return res
}
