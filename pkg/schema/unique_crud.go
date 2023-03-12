package schema

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UniqueOp struct {
	gorm.DB
}

type RequestUniqueOp struct {
	url   []string
	uid   string
	_type string
}

func (r RequestUniqueOp) WithRequestUniqueOpURL(url []string) func(r *RequestUniqueOp) {
	return func(r *RequestUniqueOp) {
		r.url = url
	}
}

func (r RequestUniqueOp) WithRequestUniqueOpUID(uid string) func(r *RequestUniqueOp) {
	return func(r *RequestUniqueOp) {
		r.uid = uid
	}
}

func (r RequestUniqueOp) WithRequestUniqueOpType(_type string) func(r *RequestUniqueOp) {
	return func(r *RequestUniqueOp) {
		r._type = _type
	}
}

func ToRequestUniqueOp(funcs ...func(r *RequestUniqueOp)) RequestUniqueOp {
	var req = new(RequestUniqueOp)
	for _, f := range funcs {
		f(req)
	}
	return *req
}

func (u *UniqueOp) First(req RequestUniqueOp) (Unique, error) {
	fields := Unique{}.Fields()
	var target Unique
	err := u.Model(Unique{}).Where(fmt.Sprintf("%s = ? AND %s = ?", fields["Url"], fields["Uid"]), req.url[0], req.uid).First(&target).Error
	return target, err
}

func (u *UniqueOp) Exists(req RequestUniqueOp) (map[string]Unique, error) {
	fields := Unique{}.Fields()
	var target []Unique
	err := u.Model(Unique{}).Where(fmt.Sprintf("%s IN ? AND %s = ?", fields["Url"], fields["Uid"]), req.url[0], req.uid).Find(&target).Error
	var m = make(map[string]Unique)
	for _, v := range target {
		m[v.Url] = v
	}
	return m, err
}

func (u *UniqueOp) CanFull(req string) (bool, error) {
	var one Unique
	err := u.Model(Unique{}).Where(fmt.Sprintf("%s = ?", Unique{}.Fields()["Uid"]), req).First(&one).Error
	if one.ID > 0 && err == nil {
		return false, nil
	}
	return true, err
}

func (u *UniqueOp) ForAll(req RequestUniqueOp) *gorm.DB {
	fields := Unique{}.Fields()
	return u.Model(Unique{}).Where(fmt.Sprintf("%s = ?", fields["Uid"]), req.uid).Order(
		clause.OrderByColumn{
			Column: clause.Column{
				Table: clause.CurrentTable, Name: clause.PrimaryKey,
			},

			Reorder: false,
		},
	).Limit(-1)
}

func (u *UniqueOp) ToUnique(req RequestUniqueOp) error {
	var docs []*Unique
	for _, url := range req.url {
		docs = append(docs, &Unique{
			Uid:  req.uid,
			Url:  url,
			Type: req._type,
		})
	}
	return u.Model(Unique{}).CreateInBatches(&docs, 1000).Error
}
