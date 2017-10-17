package main

import(
	"fmt"

	"github.com/go-pg"
)

type Pet struct {
     Id    int64
     Name  string
     Owner *Person
}

func (p Pet) String () string {
     return fmt.Sprintf("Pet <%d %s %s>", p.Id, p.Name, p.Owner)
}

type Person struct {
     Id   int64
     Name string
     Pets []Pet
}

func (p Person) String () string {
     return fmt.Sprintf("Person <%d %s %v>", p.Id, p.Name, p.Pets)
}

func MakeSchema () {
     models := []interface{}{&Pet, &Person}
     
     for _, model := range models {
     	 err := db.CreateTable(model, nil)
	 if err != nil {
	    fmt.Println("couldn't create model %s", model)
	 }
     }
     
     return nil
}

func AddPets () {
     pets := map[int]string{
     	  1: "chuck",
	  2: "baxter",
	  3: "dolly",
	  4: "spike",
	  5: "molly",
	  6: "bozo",
	  7: "woof",
	  8: "felix",
	  9: "toot toot",
     }
     all_pets := []*Pet

     for k, v := range pets {
     	 p := Pet{
	     id: k,
	     
	 }
     }
}

func AddPeople () {}

func main () {

     fmt.Println("registering models")
     MakeSchema()

     fmt.Println("adding pets")
     AddPets()

     fmt.Println("adding people")
     AddPeople()

     
}