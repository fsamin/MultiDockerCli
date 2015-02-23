package cli

import (
    "time"
    "github.com/crackcomm/go-clitable"
)

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

func PrintMDImagesList(slice []MDImage, showSize bool) {
    table := clitable.New([]string{"NODE", "HOST","IMAGE ID","TAGS","CREATED"})
    if showSize {
        table = clitable.New([]string{"NODE", "HOST","IMAGE ID","TAGS","CREATED", "VIRTUAL SIZE"})
    }

    for i := 0; i < len(slice);i++ {
        values := map[string]interface{}{
            "NODE" : slice[i].Node.Alias,
            "HOST" : slice[i].Node.Host,
            "IMAGE ID" : slice[i].Image.Id[:12],
            "TAGS" : slice[i].Image.RepoTags,
            "CREATED" : time.Unix(slice[i].Image.Created,0),
            "VIRTUAL SIZE" : slice[i].Image.VirtualSize,
        }
        table.AddRow(values)
    }
    table.Markdown = true
    table.Print()
}