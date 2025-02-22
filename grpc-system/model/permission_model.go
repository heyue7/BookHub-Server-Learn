package model

import (
	pb_sys_v1 "github.com/heyue7/BookHub-Server-Learn/proto-pb/system/v1"
	"strings"
	"time"
)

type PermissionModel struct {
	Code       string
	GroupCode  string
	Name       string
	CreateTime time.Time
	UpdateTime time.Time
}

func (*PermissionModel) TableName() string {
	return "sys_permission"
}

func (m *PermissionModel) BeforeInsert() {
	m.Code = strings.TrimSpace(m.Code)
	m.GroupCode = strings.TrimSpace(m.GroupCode)
	m.Name = strings.TrimSpace(m.Name)

	if m.CreateTime.IsZero() {
		m.CreateTime = time.Now()
	}
	if m.UpdateTime.IsZero() {
		m.UpdateTime = time.Now()
	}
}

func (m *PermissionModel) BeforeUpdate() {
	m.Code = strings.TrimSpace(m.Code)
	m.GroupCode = strings.TrimSpace(m.GroupCode)
	m.Name = strings.TrimSpace(m.Name)

	if m.UpdateTime.IsZero() {
		m.UpdateTime = time.Now()
	}
}

func (m *PermissionModel) Proto() pb_sys_v1.Permission {
	return pb_sys_v1.Permission{
		Code:       m.Code,
		GroupCode:  m.GroupCode,
		Name:       m.Name,
		CreateTime: m.CreateTime.Unix(),
		UpdateTime: m.UpdateTime.Unix(),
	}
}
