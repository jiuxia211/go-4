package service

import (
	"fmt"
	"jiuxia/crowdfundingandroid/model"
	"jiuxia/crowdfundingandroid/serializer"
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

func (service *ProjectService) Create(uid string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).First(&user, uid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到用户",
		}
	}
	if user.Class == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "你没有权限创建项目",
		}
	}
	project := model.Project{
		User:   user,
		Uid:    user.ID,
		IsPass: "unknown",
		Fund:   0,
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
			Pid:       project.ID,
			Uid:       user.ID,
			Title:     service.Title,
			Content:   service.Content,
			Fund:      project.Fund,
			IsPass:    project.IsPass,
			PicPath:   project.PicPath,
			Telephone: service.Telephone,
		},
	}
}
func (service *ProjectService) UploadFile(uid string, file multipart.File, fileSize int64, pid string) serializer.Response {
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
	var project model.Project
	if err := model.DB.Model(&model.Project{}).First(&project, pid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到项目",
		}
	}
	fmt.Println(path + "454564")
	project.PicPath = path
	model.DB.Save(&project)

	return serializer.Response{
		Status: 200,
		Msg:    "上传文件成功",
		Data: serializer.Project{
			Pid:       project.ID,
			Uid:       user.ID,
			Title:     service.Title,
			Content:   service.Content,
			Fund:      project.Fund,
			IsPass:    project.IsPass,
			PicPath:   path,
			Telephone: service.Telephone,
		},
	}

}
func (service *ProjectService) UploadInfo(uid string, pid string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).First(&user, uid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到用户",
		}
	}
	var project model.Project
	if err := model.DB.Model(&model.Project{}).First(&project, pid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到项目",
		}
	}
	fmt.Println(service.Title, service.Content, service.Telephone+"555")
	project.Title = service.Title
	project.Content = service.Content
	project.Telephone = service.Telephone
	model.DB.Save(&project)
	return serializer.Response{
		Status: 200,
		Msg:    "上传信息成功",
		Data: serializer.Project{
			Pid:       project.ID,
			Uid:       user.ID,
			Title:     service.Title,
			Content:   service.Content,
			Fund:      project.Fund,
			IsPass:    project.IsPass,
			PicPath:   project.PicPath,
			Telephone: service.Telephone,
		},
	}
}
func (service *ShowProjectService) ShowPass() serializer.Response {
	var projects []model.Project
	count := 0
	if service.PageNum == 0 && service.PageSize == 0 {
		service.PageSize = 5
	}
	model.DB.Model(&model.Project{}).Where("is_pass=?", "pass").Count((&count)).Find(&projects)
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
	model.DB.Model(&model.Project{}).Where("is_pass=?", "fail").Count((&count)).Find(&projects)
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
	model.DB.Model(&model.Project{}).Where("is_pass=?", "unknown").Count((&count)).Find(&projects)
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
	model.DB.Model(&model.Project{}).Where("uid =?", uid).Count((&count)).Find(&projects)
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
		"%"+service.Info+"%", "%"+service.Info+"%").Where("is_pass=?", "pass").Find(&projects)
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
