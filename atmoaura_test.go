package atmoaura

import (
	"testing"
	"os"
	"fmt"
)

const (
	LYON = "69123"
)

var TOKEN = os.Getenv("TOKEN")

func TestIndice(t *testing.T) {
   
	 client := NewClient(TOKEN)
	 current,err := client.GetCurrentIndex(LYON)
	 
	 if err != nil {
		t.Fatal(err)
	}
	fmt.Println(current)
}

func TestVigilance(t *testing.T) {
   
	client := NewClient(TOKEN)
	current,err := client.GetCurrentVigilance(LYON)
	
	if err != nil {
	   t.Fatal(err)
   }
   fmt.Println(current)
}


 