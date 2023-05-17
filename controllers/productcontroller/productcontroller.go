package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/autumnleaf-ra/golang/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// index
func Index(c *gin.Context) {
	var product []models.Product
	models.DB.Find(&product)
	// response JSON
	c.JSON(http.StatusOK, gin.H{"products ": product})

}

// menampilkan produk berdasarkan ID
func Show(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"product ": product})
}

// membuat data
func Create(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})

}

// update data
func Update(c *gin.Context) {

	var product models.Product

	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message ": "Data berhasil diupdate"})

}

// delete data
func Delete(c *gin.Context) {

	var product models.Product

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message ": "Data berhasil dihapus"})

}
