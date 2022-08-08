package mathehelfer

func DieSummeVon(eingabeWerte ...int) int {

	summe := 0

	for _, v := range eingabeWerte {
		summe += v
	}

	return summe
}