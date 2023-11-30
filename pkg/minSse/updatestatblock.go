package minsse

func UpdateStatBlock(name string, status string) {

	t := Stats[name]
	t.Name = name
	t.Total += 1
	switch status {
	case "error":
		t.Errors += 1
	case "failure":
		t.Failures += 1
	case "processing":
		t.Processings += 1
	case "pending":
		t.Pendings += 1
	case "completed":
		t.Completeds += 1
	default:
		t.Others += 1
	}

	t.Compute()
	Stats[name] = t

	// fmt.Println(names["overall"])
}
