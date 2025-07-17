package main
import(
	"net/http"

	"github.com/gin-gonic/gin"
    

	"github.com/jasimpn/trace/models"
	"github.com/jasimpn/trace/SQL"
	"github.com/jasimpn/trace/data"
)


// postAlbums adds an album from JSON received in the request body.
func postUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := SQL.CreateUser(data.DB, newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func main(){
	data.CreateSQLite()
	data.CreateStorage()
	
	router := gin.Default()
 
    router.POST("/login", postUser)

    router.Run("localhost:5000")
}