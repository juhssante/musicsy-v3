package утилиты

import "github.com/gin-gonic/gin"

// ХттнОшибка хттн ошибка обработчик
func ХттнОшибка(c *gin.Context, код int, ошиб error) {
	if ошиб != nil {
		c.AbortWithStatusJSON(код, gin.H{"ошибка": ошиб.Error()})
	}
}
