package taskcontrollers

import "C"
import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/wbaiyy/gopub/src/controllers"
	"github.com/wbaiyy/gopub/src/models"
	"time"
)

type RollBackController struct {
	controllers.BaseController
}

func (c *RollBackController) Get() {
	if c.Task == nil || c.Task.Id == 0 {
		c.SetJson(1, nil, "Parameter error")
		return
	}
	c.Project, _ = models.GetProjectById(c.Task.ProjectId)
	if c.Project == nil || c.Project.Id == 0 {
		c.SetJson(1, nil, "Parameter error")
		return
	}
	//上线成功 以及审核失败不允许上线
	if c.Task.Status != 3 {
		c.SetJson(1, nil, "此任务未完成")
		return
	}
	//正在上线的不允许上线
	if c.Task.Action == 1 {
		c.SetJson(1, nil, "此任务为回滚项目")
		return
	}
	//不允许回滚项目
	if c.Task.EnableRollback == 0 {
		c.SetJson(1, nil, "此任务为不允许回滚")
		return
	}
	if c.User == nil || c.User.Id == 0 {
		c.SetJson(2, nil, "not login")
		return
	}
	var task models.Task
	task.Action = 1
	task.EnableRollback = 0
	task.Branch = c.Task.Branch
	task.FileList = c.Task.FileList
	task.ExLinkId = c.Task.ExLinkId
	task.CommitId = c.Task.CommitId

	if c.GetString("this") == "this" {
		task.LinkId = c.Task.LinkId
		task.Title = fmt.Sprintf("%s【%d】-回滚当前版本", c.Task.Title, c.Task.Id)
	} else {
		task.LinkId = c.Task.ExLinkId
		task.Title = fmt.Sprintf("%s【%d】-回滚前个版本", c.Task.Title, c.Task.Id)

		o := orm.NewOrm()
		var lastTask models.Task
		err := o.Raw("select * from task where link_id=? AND project_id=? AND action = 0 limit 1",
			task.ExLinkId, c.Task.ProjectId).QueryRow(&lastTask)
		if err == nil && lastTask.CommitId != "" {
			task.CommitId = lastTask.CommitId
		}
	}
	task.ProjectId = c.Task.ProjectId
	task.Status = 0
	task.UserId = uint(c.User.Id)
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	task.Hosts = c.Task.Hosts
	_, err := models.AddTask(&task)
	if err != nil {
		c.SetJson(1, nil, "数据库更新错误")
	}
	c.SetJson(0, task, "创建成功")

	return
}
