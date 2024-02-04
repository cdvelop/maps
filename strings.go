package maps

type strings struct{}

var letters = map[rune]rune{
	'A': 'a', 'B': 'b', 'C': 'c', 'D': 'd', 'E': 'e', 'F': 'f', 'G': 'g', 'H': 'h', 'I': 'i',
	'J': 'j', 'K': 'k', 'L': 'l', 'M': 'm', 'N': 'n', 'O': 'o', 'P': 'p', 'Q': 'q', 'R': 'r',
	'S': 's', 'T': 't', 'U': 'u', 'V': 'v', 'W': 'w', 'X': 'x', 'Y': 'y', 'Z': 'z',
}

func (strings) Join(values []string, separate string) (out string) {

	var space string
	for _, v := range values {

		out += space + v
		space = separate

	}
	return
}

// ej: A:a, B:b.. with_ñ = Ñ:ñ
func (strings) LettersUpperLowerCase(with_ñ ...bool) map[rune]rune {

	for _, ñ := range with_ñ {
		if ñ {
			letters['Ñ'] = 'ñ'
			return letters
		}
	}

	// Si no se incluye la letra "Ñ", se elimina del mapa
	delete(letters, 'Ñ')

	return letters
}

func (s strings) LowerCaseFirstLetter(name string) string {
	if newChar, ok := s.LettersUpperLowerCase()[rune(name[0])]; ok {
		return string(newChar) + name[1:]
	}
	return name
}
