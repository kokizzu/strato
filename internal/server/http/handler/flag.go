package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/purpledb/purple"
)

var falseValueJson = gin.H{
	"value": false,
}

func (h *Handler) FlagGet(c *gin.Context) {
	log := h.logger("flag/get")

	key := c.Param("key")

	val, err := h.b.FlagGet(key)
	if err != nil {
		if purple.IsNotFound(err) {
			c.JSON(http.StatusOK, falseValueJson)
			return
		}

		log.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	res := gin.H{
		"value": val,
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) FlagSet(c *gin.Context) {
	log := h.logger("flag/set")

	key, val := c.Param("key"), getFlagValue(c)

	if err := h.b.FlagSet(key, val); err != nil {
		log.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}
