package wallecontrollers

import (
	"github.com/wbaiyy/gopub/src/controllers"
	"github.com/wbaiyy/gopub/src/library/components"
	"github.com/wbaiyy/gopub/src/models"
)

type DetectionsshController struct {
	controllers.BaseController
}

func (c *DetectionsshController) Get() {
	if c.Project == nil || c.Project.Id == 0 {
		c.SetJson(1, nil, "Parameter error")
		return
	}
	s := components.BaseComponents{}
	s.SetProject(c.Project)
	s.SetTask(&models.Task{Id: -1})
	err := s.TestSsh()
	if err != nil {
		c.SetJson(1, nil, "ssh目标机器错误"+err.Error())
		return
	}
	c.SetJson(0, nil, "")
	c.ServeJSON()

}
