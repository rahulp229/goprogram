package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type testHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}
type myForm struct {
	Colors []string `form:"colors[]"`
}

func Init() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("form.html")

	router.GET("/ping", func(c *gin.Context) {

		h := testHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(200, err)
		}

		fmt.Printf("%#v\n", h)
		c.JSON(200, gin.H{"Message": "Hello world!"})
	})

	router.GET("", func(c *gin.Context) {
		c.HTML(200, "form.html", nil)
	})

	router.POST("post-form", func(c *gin.Context) {
		var fakeForm myForm
		c.Bind(&fakeForm)
		//c.JSON(200, gin.H{"color": fakeForm})
		fmt.Sprintf()

		c.HTML(200, "form.html", gin.H{"color": fakeForm})
	})
	return router
}
