package taskcontrollers

import (
	"github.com/astaxie/beego/orm"
	"github.com/wbaiyy/gopub/src/controllers"
	"github.com/wbaiyy/gopub/src/library/components"
	"github.com/wbaiyy/gopub/src/models"
)

type ChangesController struct {
	controllers.BaseController
}

func (c *ChangesController) Get() {
	taskId, _ := c.GetInt("taskId", 0)

	if taskId == 0 {
		c.SetJson(1, nil, "Parameter error")
		return
	}
	o := orm.NewOrm()

	var task models.Task
	o.Raw("SELECT * FROM task where task.id = ?", taskId).QueryRow(&task)

	project, _ := models.GetProjectById(task.ProjectId)

	if project.RepoType == "git" {
		var last_task models.Task
		o.Raw("SELECT * FROM task where project_id = ? AND status=3 AND id<? order by task.id DESC LIMIT 1",
			task.ProjectId, taskId).QueryRow(&last_task)

		s := components.BaseComponents{}
		s.SetProject(project)
		s.SetTask(&task)
		s.CmdNotRecord = 1 //不记录cmd record

		g := components.BaseGit{}
		g.SetBaseComponents(s)
		files, _ := g.DiffBetweenCommits(task.Branch, task.CommitId, last_task.CommitId)

		var fileinfos []map[string]string
		if len(files) < 10 && len(files) > 0 {
			for _, filepath := range files {
				fileinfo, _ := g.GetLastModifyInfo(task.Branch, filepath)
				fileinfo["path"] = filepath
				fileinfos = append(fileinfos, fileinfo)
			}
		} else {

		}
		c.SetJson(0, fileinfos, "")
		return
	} else {
		c.SetJson(1, nil, "Project is not git")
	}

}
