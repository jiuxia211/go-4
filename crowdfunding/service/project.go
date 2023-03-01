package service

import (
	"jiuxia/crowdfunding/model"
	"jiuxia/crowdfunding/serializer"
	"mime/multipart"
)

type ProjectService struct {
	Title     string `form:"title" json:"title" `
	Content   string `form:"content" json:"content" `
	Telephone string `form:"telephone" json:"telephone" `
}
type ShowProjectService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}
type SearchProjectService struct {
	Info     string `json:"info" form:"info"`
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
}
type DetailService struct {
}

func (service *ProjectService) Upload(uid string, file multipart.File, fileSize int64) serializer.Response {
	path, err := UploadToQiNiu(file, fileSize)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "上传失败",
		}

	}
	var user model.User
	if err := model.DB.Model(&model.User{}).First(&user, uid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到用户",
		}
	}
	// fmt.Println(path)
	project := model.Project{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		IsPass:    "unknown",
		Fund:      0,
		PicPath:   path,
		Telephone: service.Telephone,
	}
	if err := model.DB.Create(&project).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "创建项目失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建项目成功",
		Data: serializer.Project{
			Pid:     project.ID,
			Uid:     user.ID,
			Title:   service.Title,
			Content: service.Content,
			Fund:    project.Fund,
			IsPass:  project.IsPass,
			PicPath: path,
		},
	}

}
func (service *ShowProjectService) ShowPass() serializer.Response {
	var projects []model.Project
	count := 0
	if service.PageNum == 0 && service.PageSize == 0 {
		service.PageSize = 5
	}
	model.DB.Model(&model.Project{}).Where("is_pass=?", "pass").Count((&count)).Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).Find(&projects)
	return serializer.Response{
		Status: 200,
		Data: serializer.ProjectList{
			Item:  serializer.BuildProjects(projects),
			Total: uint(count),
		},
		Msg: "数据如上",
	}
}
func (service *ShowProjectService) ShowFail() serializer.Response {
	var projects []model.Project
	count := 0
	if service.PageNum == 0 && service.PageSize == 0 {
		service.PageSize = 5
	}
	model.DB.Model(&model.Project{}).Where("is_pass=?", "fail").Count((&count)).Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).Find(&projects)
	return serializer.Response{
		Status: 200,
		Data: serializer.ProjectList{
			Item:  serializer.BuildProjects(projects),
			Total: uint(count),
		},
		Msg: "数据如上",
	}
}
func (service *ShowProjectService) ShowUnknown() serializer.Response {
	var projects []model.Project
	count := 0
	if service.PageNum == 0 && service.PageSize == 0 {
		service.PageSize = 5
	}
	model.DB.Model(&model.Project{}).Where("is_pass=?", "unknown").Count((&count)).Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).Find(&projects)
	return serializer.Response{
		Status: 200,
		Data: serializer.ProjectList{
			Item:  serializer.BuildProjects(projects),
			Total: uint(count),
		},
		Msg: "数据如上",
	}
}
func (service *ShowProjectService) ShowMy(uid string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).First(&user, uid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到用户",
		}
	}
	var projects []model.Project
	count := 0
	if service.PageNum == 0 && service.PageSize == 0 {
		service.PageSize = 5
	}
	model.DB.Model(&model.Project{}).Where("uid =?", uid).Count((&count)).Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).Find(&projects)
	return serializer.Response{
		Status: 200,
		Data: serializer.ProjectList{
			Item:  serializer.BuildProjects(projects),
			Total: uint(count),
		},
		Msg: "数据如上",
	}
}
func (service *SearchProjectService) Search() serializer.Response {
	var projects []model.Project
	count := 0
	if service.PageNum == 0 && service.PageSize == 0 {
		service.PageSize = 5
	}
	model.DB.Model(&model.Project{}).Where("title LIKE ? OR content LIKE ?",
		"%"+service.Info+"%", "%"+service.Info+"%").Where("is_pass=?", "pass").Count((&count)).Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).Find(&projects)
	return serializer.Response{
		Status: 200,
		Data: serializer.ProjectList{
			Item:  serializer.BuildProjects(projects),
			Total: uint(count),
		},
		Msg: "搜索数据如上",
	}
}
func (service *DetailService) Detail(pid string) serializer.Response {
	var project model.Project
	if err := model.DB.Model(&model.Project{}).First(&project, pid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到项目",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildProject(project),
		Msg:    "项目数据如上",
	}
}
