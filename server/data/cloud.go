package data

import(
	"os"
	"log"

	// "github.com/jasimpn/trace/models"
)

func CreateStorage(){
	const FolderName = "storage001"
	err := os.Mkdir(FolderName, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}