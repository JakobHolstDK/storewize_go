package storwize_test

import (
	"context"
	"fmt"
	"log"

	"github.com/andswarian/storewize_go/storwize"
)

func Example() {
	client := storwize.NewClient(storwize.WithToken("token"))

	lun, _, err := client.lun.GetByID(context.Background(), 1)
	if err != nil {
		log.Fatalf("error retrieving lun: %s\n", err)
	}
	if lun != nil {
		fmt.Printf("lun 1 is called %q\n", lun.Name)
	} else {
		fmt.Println("lun 1 not found")
	}
}
