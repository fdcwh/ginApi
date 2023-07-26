package article

import (
	"github.com/gin-gonic/gin"
	"goGIn/kernel/res"
	"goGIn/kernel/utils"
)

func GetTags(c *gin.Context) {
	/*name := c.Query("name")
	state := -1*/

}

type AddTagForm struct {
	Mobile   string `form:"mobile" binding:"required,mobile" required_err:"请输入手机号码" mobile_err:"手机号码格式"`
	Password string `form:"password" binding:"required"`
}

// Add article tag

func AddTag(c *gin.Context) {
	var formData AddTagForm
	if err := c.ShouldBind(&formData); err != nil {
		res.ErrParam.SetMsg(utils.ValidatorGetError(err, formData)).ToJson(c)
		return
	}

	res.Success.SetData(formData).ToJson(c)
}
