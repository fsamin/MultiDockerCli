package cli

func ExtendMDContainersList(slice []MDContainer, element MDContainer) []MDContainer {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]MDContainer, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func PrintMDContainersList(slice []MDContainer) {
	table.PrintHorizontal(map[string]string{
		"Name": "MongoLab",
		"Host": "mongolab.com",
		// ...
	})
}
