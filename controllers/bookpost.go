package controllers
import(
	"net/http"
	"github.com/gin-gonic/gin"
	"libraryBackend/models"
)

type  BookpostInput struct {
	ISBN string 	`json:"isbn"`
	Libid int `json:"libid"`
	Title string `json:"title`
	Authors string `json:"authors"`
	Publisher string `json:"publisher"`
	Version string `json:"version"`
	Totalcopies int `json:"totalcopies"`
	Avaiblecopies int `json:"avaiblecopies"`
}

func Bookpost(c *gin.Context){
	
	var input BookpostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Bookinventory{}

	u.ISBN = input.ISBN
	u.Libid = input.Libid
	u.Title=input.Title
	u.Authors=input.Authors
	u.Publisher=input.Publisher
	u.Version =input.Version
	u.Totalcopies=input.Totalcopies
	u.Avaiblecopies=input.Avaiblecopies

	_,err := u.SaveBook()

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"your book successfully add in your library "})

}

