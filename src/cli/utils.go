package cli


func ExtendMDContainersList(slice []MDContainer, element MDContainer) []MDContainer {
	n := len(slice)
	if n == cap(slice) {
		newSlice := make([]MDContainer, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]

	slice[n] = element
	return slice
}

func ExtendMDImageList(slice []MDImage, element MDImage) []MDImage {
    n := len(slice)
    if n == cap(slice) {
        newSlice := make([]MDImage, len(slice), 2*len(slice)+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0 : n+1]

    slice[n] = element
    return slice
}

