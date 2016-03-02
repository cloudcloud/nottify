package cli

import (
	"fmt"

	"github.com/cloudcloud/nottify/nottify"
)

var cmdIngest = &Command{
	UsageLine: "ingest",
	Short:     "Complete an ingestion of Media",
	Long: `
Ingest will complete a file system trawl for media files.

For example:
	nottify ingest
`,
}

func init() {
	cmdIngest.Run = ingestRun
}

func ingestRun(args []string) {
	c := nottify.NewConfig()
	n := nottify.New(c)

	r, err := n.Ingest()
	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}
