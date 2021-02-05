package wallecontrollers

import (
	"github.com/wbaiyy/gopub/src/controllers"
	"github.com/wbaiyy/gopub/src/library/components"
	"github.com/wbaiyy/gopub/src/models"
)

type TagController struct {
	controllers.BaseController
}

func (c *TagController) Get() {
	if c.Project == nil || c.Project.Id == 0 {
		c.SetJson(1, nil, "Parameter error")
		return
	}
	s := components.BaseComponents{}
	s.SetProject(c.Project)
	s.SetTask(&models.Task{})
	g := components.BaseGit{}
	g.SetBaseComponents(s)
	res, err := g.GetTagList(200)
	if err != nil {
		c.SetJson(1, nil, "获取tag错误—"+err.Error())
		return
	} else {
		c.SetJson(0, res, "")
		return
	}

}
