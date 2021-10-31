package main

import (
	"github.com/aerostatka/third-party-integrations/application"
	"os"
)

func main()  {
	args := os.Args[1:]

	application.Start(args)
}
