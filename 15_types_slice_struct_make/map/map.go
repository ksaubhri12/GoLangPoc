package _map

func MakeMap()  map[string]string {
	nameLastName := map[string]string{
		"Kalpesh" : "Saubhri",
		"Nitika" : "Jindal",
		"Mayank": "Grover",
		"Aditya" : "Ambikesh",
	}
	nameLastName["Moulik"] = "Adak"
	return nameLastName
}

func UpdateMap(nameLastName map[string]string)  {
	nameLastName["Aditya"] = "Ambik"
}
