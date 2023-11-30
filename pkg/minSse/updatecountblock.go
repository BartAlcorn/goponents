package minsse

func UpdateCountBlock(name string) {

	t := Stats[name]
	t.Name = name
	t.Total = len(Assets)

	errors := 0
	failures := 0
	processings := 0
	pendings := 0
	completeds := 0
	others := 0

	for _, a := range Assets {
		switch a.Status {
		case "error":
			errors += 1
		case "failure":
			failures += 1
		case "processing":
			processings += 1
		case "pending":
			pendings += 1
		case "completed":
			completeds += 1
		default:
			others += 1
		}
	}

	t.Errors = errors
	t.Failures = failures
	t.Processings = processings
	t.Pendings = pendings
	t.Completeds = completeds
	t.Others = others
	Stats[name] = t
}
