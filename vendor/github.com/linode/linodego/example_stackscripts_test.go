package linodego_test

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/linode/linodego"
)

func ExampleClient_CreateStackscript() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleCreateStackscript")
	defer teardown()

	fmt.Println("## Stackscript create")

	var ss *linodego.Stackscript
	var err error
	for rev := 1; rev < 4; rev++ {
		fmt.Println("### Revision", rev)
		if rev == 1 {
			stackscript := linodego.Stackscript{}.GetCreateOptions()
			stackscript.Description = "description for example stackscript " + time.Now().String()
			// stackscript.Images = make([]string, 2, 2)
			stackscript.Images = []string{"linode/debian9", "linode/ubuntu18.04"}
			stackscript.IsPublic = false
			stackscript.Label = "example stackscript " + time.Now().String()
			stackscript.RevNote = "revision " + strconv.Itoa(rev)
			stackscript.Script = "#!/bin/bash\n"
			ss, err = linodeClient.CreateStackscript(context.Background(), stackscript)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			update := ss.GetUpdateOptions()
			update.RevNote = "revision " + strconv.Itoa(rev)
			update.Label = strconv.Itoa(rev) + " " + ss.Label
			update.Script += "echo " + strconv.Itoa(rev) + "\n"
			ss, err = linodeClient.UpdateStackscript(context.Background(), ss.ID, update)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println("### Get")
	ss, err = linodeClient.GetStackscript(context.Background(), ss.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("### Delete")
	err = linodeClient.DeleteStackscript(context.Background(), ss.ID)
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	// ## Stackscript create
	// ### Revision 1
	// ### Revision 2
	// ### Revision 3
	// ### Get
	// ### Delete
}
