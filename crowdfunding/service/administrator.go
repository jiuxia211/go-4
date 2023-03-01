package service

import (
	"jiuxia/crowdfunding/model"
	"jiuxia/crowdfunding/serializer"
)

type AuditProjectService struct {
	Ispass string `json:"ispass" form:"ispass"`
}
type DeleteProjectService struct {
}

func (service *AuditProjectService) Audit(uid string, pid string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).First(&user, uid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到用户",
		}
	}
	if user.Class != 3 {
		return serializer.Response{
			Status: 400,
			Msg:    "你不是管理员",
		}
	}
	var project model.Project
	if err := model.DB.Model(&model.Project{}).First(&project, pid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到项目",
		}
	}
	if service.Ispass == "fail" {
		project.IsPass = "fail"
		model.DB.Save(&project)
		return serializer.Response{
			Status: 200,
			Msg:    "审核成功,项目未通过",
		}
	} else if service.Ispass == "pass" {
		project.IsPass = "pass"
		model.DB.Save(&project)
		return serializer.Response{
			Status: 200,
			Msg:    "审核成功,项目通过",
		}
	}
	return serializer.Response{
		Status: 500,
		Msg:    "审核失败",
	}
}
func (service *DeleteProjectService) Delete(uid string, pid string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).First(&user, uid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到用户",
		}
	}
	if user.Class != 3 {
		return serializer.Response{
			Status: 400,
			Msg:    "你不是管理员",
		}
	}
	var project model.Project
	if err := model.DB.Model(&model.Project{}).First(&project, pid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到项目",
		}
	}
	if err := model.DB.Delete(&model.Project{}, pid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "删除失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}
