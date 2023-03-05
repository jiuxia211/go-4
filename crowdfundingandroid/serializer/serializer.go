package serializer

import "jiuxia/crowdfundingandroid/model"

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
type User struct {
	UserName string `json:"username"`
	ID       string `json:"id"`
	Email    string `json:"email"`
	Money    int64  `json:"money"`
	Class    uint   `json:"class"`
}
type Project struct {
	Pid       uint   `json:"pid"`
	Uid       uint   `json:"uid"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Fund      int64  `json:"fund"`
	IsPass    string `json:"is_pass"`
	PicPath   string `json:"pic_path"`
	Telephone string `json:"telephone"`
}
type ProjectList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

func BuildProject(project model.Project) Project {
	return Project{
		Pid:       project.ID,
		Uid:       project.Uid,
		Title:     project.Title,
		Content:   project.Content,
		Fund:      project.Fund,
		IsPass:    project.IsPass,
		PicPath:   project.PicPath,
		Telephone: project.Telephone,
	}
}
func BuildProjects(items []model.Project) (projects []Project) {
	for _, item := range items {
		project := BuildProject(item)
		projects = append(projects, project)
	}
	return projects
}
