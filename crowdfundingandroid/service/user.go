package service

import (
	"fmt"
	"jiuxia/crowdfundingandroid/conf"
	"jiuxia/crowdfundingandroid/model"
	"jiuxia/crowdfundingandroid/serializer"
	"jiuxia/crowdfundingandroid/utils"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/mail.v2"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" `
	Email    string `form:"email" json:"email" `
	Password string `form:"password" json:"password" `
	Class    uint   `form:"class" json:"class" `
}
type SendEmailService struct {
}
type MakeContributionService struct {
	Fund int64 `form:"fund" json:"fund"`
}
type UpdatePasswordService struct {
	Original_password string `form:"original_password" json:"original_password"`
	New_password      string `form:"new_password" json:"new_password"`
	Token             string `form:"token" json:"token"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).
		First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "该用户已注册！",
		}
	}
	user.UserName = service.UserName
	user.Email = service.Email
	user.Class = service.Class
	user.Money = 1000
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库创建用户失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建成功",
	}
}
func (service *UserService) Login() serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).
		First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在",
			}
		} else {
			return serializer.Response{
				Status: 500,
				Msg:    "数据库错误",
			}
		}

	}
	if !user.CheckPassword(service.Password) {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	token, err := utils.GenerateToken(service.UserName, user.ID, service.Password)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "登录成功",
		Data: serializer.TokenData{
			User: serializer.User{
				UserName: service.UserName,
				ID:       strconv.Itoa(int(user.ID)),
				Email:    user.Email,
				Money:    user.Money,
				Class:    user.Class,
			},
			Token: token,
		},
	}
}

func (service *UserService) UpdateEmail(uid string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).First(&user, uid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到用户",
		}
	}
	if service.Email == "" {
		return serializer.Response{
			Status: 400,
			Msg:    "请输入邮箱",
		}
	}
	user.Email = service.Email
	model.DB.Save(&user)
	return serializer.Response{
		Status: 200,
		Msg:    "修改成功",
	}

}
func (service UpdatePasswordService) UpdatePassword(uid string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).First(&user, uid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到用户",
		}
	}
	if !user.CheckPassword(service.Original_password) {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	claims, err := utils.ParseEmailToken(service.Token)
	if time.Now().Unix() > claims.ExpiresAt {
		return serializer.Response{
			Status: 400,
			Msg:    "邮箱token过期",
		}
	}
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "邮箱token解析失败",
		}
	}

	if err := user.SetPassword(service.New_password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}

	model.DB.Save(&user)
	return serializer.Response{
		Status: 200,
		Msg:    "修改成功",
	}

}
func (service *SendEmailService) Send(uid string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).First(&user, uid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到用户",
		}
	}
	token, err := utils.GenerateEmailToken(uid, user.Email)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "token签发错误",
		}
	}
	mailText := conf.ValidEmail + token
	// mailText := conf.ValidEmail + token
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "zzz")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	fmt.Println(user.Email + "zzz")
	if err := d.DialAndSend(m); err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "发送邮件失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "发送成功",
	}
}
func (service *MakeContributionService) Contribution(uid string, pid string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).First(&user, uid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到用户",
		}
	}
	if service.Fund < 0 {
		return serializer.Response{
			Status: 400,
			Msg:    "不能反向出资哦",
		}
	}
	if (user.Money - service.Fund) >= 0 {
		user.Money -= service.Fund
	} else {
		return serializer.Response{
			Status: 400,
			Msg:    "余额不足",
		}
	}
	model.DB.Save(&user)
	var project model.Project
	if err := model.DB.Model(&model.Project{}).First(&project, pid).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到项目",
		}
	}
	project.Fund += service.Fund
	model.DB.Save(&project)
	return serializer.Response{
		Status: 200,
		Msg:    "出资成功",
	}
}
