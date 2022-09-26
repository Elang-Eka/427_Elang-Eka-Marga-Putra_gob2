package main

func main() {

	//Empty Interface
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan Sudirman"

	randomValues = 20

	randomValues = true

	randomValues = []string{"Elang", "Eka"}

	// Empty interface (Type assertion)
	var v interface{}

	v = 20

	if value, ok := v.(int); ok == true {
		v = value * 9
	}

	// Empty interface (Map & Slice with Empty Interface)
	rs := []interface{}{1, "Elang", true, 2, "Eka", true}

	rm := map[string]interface{}{
		"name":   "Elang",
		"status": true,
		"Age":    23,
	}

	_, _ = rs, rm
}
