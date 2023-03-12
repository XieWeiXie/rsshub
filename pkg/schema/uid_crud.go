package schema

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserOp struct {
	gorm.DB
}

type RequestUser struct {
	uid  string
	desc string
	page int
	size int
}

func (r RequestUser) WithRequestUserUID(uid string) func(r *RequestUser) {
	return func(r *RequestUser) {
		r.uid = uid
	}
}

func (r RequestUser) WithRequestUserDesc(desc string) func(r *RequestUser) {
	return func(r *RequestUser) {
		r.desc = desc
	}
}
func (r RequestUser) WithRequestUserPage(page, size int) func(r *RequestUser) {
	return func(r *RequestUser) {
		r.page, r.size = page, size
	}
}

func (u UserOp) ToUser(req RequestUser) error {
	return u.Model(Users{}).Create(&Users{
		Uid:         req.uid,
		Description: req.desc,
	}).Error
}

func (u UserOp) One(uid string) (Users, error) {
	var doc Users
	err := u.Model(Users{}).Where(fmt.Sprintf("%s = ?", Users{}.Fields()["Uid"]), uid).First(&doc).Error
	return doc, err
}

func (u UserOp) ForAll() *gorm.DB {
	return u.Model(Users{}).Order(clause.OrderByColumn{
		Column: clause.Column{
			Table: clause.CurrentTable, Name: clause.PrimaryKey,
		},
		Desc:    false,
		Reorder: false,
	}).Limit(-1)
}

func (u UserOp) List(req RequestUser) ([]Users, int, error) {
	var docs []Users
	query := u.Model(Users{}).Order(clause.OrderByColumn{Column: clause.Column{
		Table: clause.CurrentTable, Name: clause.PrimaryKey,
	}})
	if len(req.uid) > 0 {
		query = query.Where(fmt.Sprintf("%s = ?", Users{}.Fields()["Uid"]), req.uid)
	}
	var count int64 = 0
	query.Count(&count)
	err := query.Offset((req.page - 1) * req.size).Limit(req.size).Find(&docs).Error
	return docs, int(count), err
}
