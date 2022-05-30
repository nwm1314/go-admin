package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mojocn/base64Captcha"
	"go-admin/server/common"
	"go-admin/server/dao"
	"go-admin/server/model/database"
	"go-admin/server/model/request"
	"go-admin/server/model/response"
	"go-admin/server/utils"
	"time"
)

func GetAuthorityDataById(authorityId string) response.Authority {
	var authorityData response.Authority
	authority, _ := dao.GetAuthorityById(authorityId)
	authorityData.AuthorityId = authority.AuthorityId
	authorityData.AuthorityName = authority.AuthorityName
	authorityData.CreatedAt = authority.CreatedAt
	authorityData.DefaultRouter = authority.DefaultRouter
	authorityData.DeletedAt = authority.DeletedAt
	authorityData.ParentId = authority.ParentId
	authorityData.UpdatedAt = authority.UpdatedAt
	childAuthorityIds, _ := dao.GetChildAuthorityIdListByAuthorityId(authorityId)
	if childAuthorityIds != nil {
		var childAuthorities []response.Authority
		for _, childAuthorityId := range childAuthorityIds {
			childAuthority := GetAuthorityDataById(childAuthorityId)
			childAuthorities = append(childAuthorities, childAuthority)
		}
		authorityData.Children = childAuthorities
	}
	dataAuthorityIds, _ := dao.GetDataAuthorityIdListByAuthorityId(authorityId)
	if dataAuthorityIds != nil {
		var dataAuthorities []response.Authority
		for _, dataAuthorityId := range childAuthorityIds {
			dataAuthority := GetAuthorityDataById(dataAuthorityId)
			dataAuthorities = append(dataAuthorities, dataAuthority)
		}
		authorityData.DataAuthorityId = dataAuthorities
	}
	return authorityData
}

func GetUserData(id uint) response.UserInfo {
	user, _ := dao.GetUserById(id)
	var userInfo = response.UserInfo{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdateAt:    user.UpdatedAt,
		Uuid:        user.Uuid,
		Username:    user.Username,
		NickName:    user.NickName,
		SideMode:    user.SideMode,
		HeaderImg:   user.HeaderImg,
		BaseColor:   user.BaseColor,
		ActiveColor: user.ActiveColor,
		Phone:       user.Phone,
		Email:       user.Email,
		AuthorityId: user.AuthorityId,
	}
	authorityIds, _ := dao.GetAuthorityIdListByUserId(user.ID)
	if authorityIds != nil {
		var authorities []response.Authority
		for _, authorityId := range authorityIds {
			authority := GetAuthorityDataById(authorityId)
			authorities = append(authorities, authority)
		}
		userInfo.Authorities = authorities
	}
	userInfo.Authority = GetAuthorityDataById(user.AuthorityId)
	return userInfo
}

func GetUserInfo(id uint) common.Data {
	var userInfoResp response.UserInfoResp
	userInfo := GetUserData(id)
	userInfoResp.UserInfo = userInfo
	return common.NewData(0, userInfoResp, "获取成功")
}

func GetLoginResp(user *database.User) common.Data {
	var loginResp response.LoginResp
	var userInfo = response.UserInfo{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdateAt:    user.UpdatedAt,
		Uuid:        user.Uuid,
		Username:    user.Username,
		NickName:    user.NickName,
		SideMode:    user.SideMode,
		HeaderImg:   user.HeaderImg,
		BaseColor:   user.BaseColor,
		ActiveColor: user.ActiveColor,
		Phone:       user.Phone,
		Email:       user.Email,
		AuthorityId: user.AuthorityId,
	}
	authorityIds, _ := dao.GetAuthorityIdListByUserId(user.ID)
	if authorityIds != nil {
		var authorities []response.Authority
		for _, authorityId := range authorityIds {
			authority := GetAuthorityDataById(authorityId)
			authorities = append(authorities, authority)
		}
		userInfo.Authorities = authorities
	}
	userInfo.Authority = GetAuthorityDataById(user.AuthorityId)
	loginResp.User = userInfo
	expireTime := time.Now().Add(3600 * time.Second)
	var claims = request.BaseClaims{
		user.Uuid,
		user.ID,
		user.Username,
		user.NickName,
		user.AuthorityId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "qmPlus",
		},
	}
	token, err := utils.GenerateToken(claims)
	if err != nil {
		return common.NewData(0, loginResp, "toekn生成失败")
	}
	loginResp.Token = token
	loginResp.ExpiresAt = expireTime.Unix() * 1000
	return common.NewData(0, loginResp, "登录成功")
}

func VerifyUser(username, password, captchaContent, captchaId string, captcha *base64Captcha.Captcha) common.Data {
	if !captcha.Verify(captchaId, captchaContent, true) {
		return common.NewData(7, nil, "验证码错误!")
	}
	user, _ := dao.GetUserByUsername(username)
	if user == nil {
		return common.NewData(7, nil, "用户名不存在!")
	}
	if !utils.CheckPassword(password, user.Password) {
		return common.NewData(7, nil, "密码错误!")
	}
	return GetLoginResp(user)
}
