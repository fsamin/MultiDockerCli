package cli

import (
    "time"
    "github.com/crackcomm/go-clitable"
)

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

func PrintMDContainersList(slice []MDContainer, showSize bool) {
    table := clitable.New([]string{"NODE", "HOST","CONTAINER ID","IMAGE","COMMAND","CREATED","STATUS"})
    if showSize {
        table = clitable.New([]string{"NODE", "HOST","CONTAINER ID","IMAGE","COMMAND","CREATED","STATUS", "SIZE"})
    }

    for i := 0; i < len(slice);i++ {
        values := map[string]interface{}{
            "NODE" : slice[i].Node.Alias,
            "HOST" : slice[i].Node.Host,
            "CONTAINER ID" : slice[i].Container.Id[:12],
            "IMAGE" : slice[i].Container.Image,
            "COMMAND" : slice[i].Container.Command,
            "CREATED" : time.Unix(slice[i].Container.Created,0),
            "STATUS" : slice[i].Container.Status,
            "SIZE" : slice[i].Container.SizeRootFs,
        }
        table.AddRow(values)
    }
    table.Markdown = true
    table.Print()
}
