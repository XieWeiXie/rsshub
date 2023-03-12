package weibo

import (
	"context"
	"fmt"
	"github.com/XieWeiXie/rsshub/api/v1"
	"github.com/XieWeiXie/rsshub/pkg/db"
	"github.com/XieWeiXie/rsshub/pkg/schema"
	"gorm.io/gorm"
	"log"
	"time"
)

type Service struct {
	v1.UnimplementedWeiboServer
}

const (
	WeiboType = "weibo"
)

func (s *Service) FistFull(uid string) error {
	first := NewFirst(uid, 0)
	op := schema.UniqueOp{DB: *db.DefaultMysql}
	ok, _ := op.CanFull(uid)
	if !ok {
		return nil
	}
	var req schema.RequestUniqueOp
	doFunc := func(url []string) {
		req := schema.ToRequestUniqueOp(
			req.WithRequestUniqueOpType(WeiboType),
			req.WithRequestUniqueOpUID(uid),
			req.WithRequestUniqueOpURL(url),
		)
		op.ToUnique(req)
	}
	return first.ToFirst(doFunc)
}

func (s *Service) ContentsFull(uid string) (err error) {
	op := schema.ContentsOp{DB: *db.DefaultMysql}
	if _, err = op.CanFull(uid); err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == nil {
		return
	}

	op2 := schema.UniqueOp{DB: *db.DefaultMysql}
	req := schema.RequestUniqueOp{}
	req.WithRequestUniqueOpUID(uid)(&req)
	var docs []schema.Unique
	content := NewContents()
	op2.ForAll(req).FindInBatches(&docs, 10, func(tx *gorm.DB, batch int) error {
		for _, one := range docs {
			fmt.Println(one.Url, "......")
			v, err := op.ForOne(one.Url)
			if err != nil && err != gorm.ErrRecordNotFound {
				continue
			}
			if v.ID > 0 {
				continue
			}
			_ = content.ToContent(one.Url)
			time.Sleep(2 * time.Second)
		}
		return nil
	})
	return
}

func (s *Service) NewOriginWeibo(ctx context.Context, req *v1.NewOriginWeiboReq) (*v1.NewOriginWeiboReply, error) {
	first := NewFirst(req.Uid, int(req.Size))
	op := schema.UniqueOp{}
	op.DB = gorm.DB{} // todo
	var request = new(schema.RequestUniqueOp)
	var newUrl []string
	doFunc := func(url []string) {
		req1 := schema.ToRequestUniqueOp(
			request.WithRequestUniqueOpType(WeiboType),
			request.WithRequestUniqueOpUID(req.Uid),
			request.WithRequestUniqueOpURL(url),
		)
		exists, err := op.Exists(req1)
		if err != nil {
			log.Fatalln("find url form db fail")
			return
		}
		for _, one := range url {
			if _, ok := exists[one]; ok {
				continue
			}
			newUrl = append(newUrl, one)
		}
		if len(newUrl) > 0 {
			req1.WithRequestUniqueOpURL(newUrl)(&req1)
			_ = op.ToUnique(req1)
		}
	}
	err := first.ToFirst(doFunc)
	var reply = new(v1.NewOriginWeiboReply)
	if len(newUrl) == 0 {
		return reply, nil
	}
	return reply, err

}

func (s *Service) NewAllWeibo(req *v1.NewAllWeiboReq, server v1.Weibo_NewAllWeiboServer) error {
	op := schema.ContentsOp{DB: *db.DefaultMysql}
	request := schema.RequestContentQuery{}
	request.WithRequestContentQuery(req.Uid)(&request)
	request.WithRequestContentSelector([]string{schema.Contents{}.Fields()["ID"], schema.Contents{}.Fields()["Uid"], schema.Contents{}.Fields()["Content"]})(&request)
	doc := op.Batch(request)
	var docs []schema.Contents
	doc.FindInBatches(&docs, 10, func(tx *gorm.DB, batch int) error {
		var reply = new(v1.NewAllWeiboReply)
		reply.Response = make(map[string][]byte)
		for _, doc := range docs {
			reply.Response[doc.Url] = []byte(doc.Content)
		}
		server.Send(reply)
		return nil
	})
	return nil
}

func (s *Service) mustEmbedUnimplementedWeiboServer() {
}

func (s *Service) Hello(ctx context.Context, req *v1.HelloReq) (*v1.HelloReply, error) {
	var (
		reply = new(v1.HelloReply)
	)
	reply.Data = req.Data
	return reply, nil
}

func (s *Service) NewAddWeibo(ctx context.Context, in *v1.NewAddWeiboReq) (*v1.NewAddWeiboReply, error) {
	var (
		reply = new(v1.NewAddWeiboReply)
		err   error
	)
	op := schema.UserOp{DB: *db.DefaultMysql}
	v, err := op.One(in.Uid)
	request := schema.RequestUser{}
	if len(in.Url) > 0 {
		in.Uid = toUrlUid(in.Url)
	}
	request.WithRequestUserUID(in.Uid)(&request)
	request.WithRequestUserDesc(in.Desc)(&request)
	if v.ID == 0 {
		err = op.ToUser(request)
	}
	return reply, err
}

func (s *Service) WeiboUsers(ctx context.Context, req *v1.WeiboUsersReq) (*v1.WeiboUsersReply, error) {
	var (
		reply = new(v1.WeiboUsersReply)
		err   error
	)
	op := schema.UserOp{DB: *db.DefaultMysql}
	request := schema.RequestUser{}
	request.WithRequestUserUID(req.Uid)(&request)
	request.WithRequestUserPage(int(req.Page), int(req.Size))(&request)
	v, c, err := op.List(request)
	if err != nil {
		return reply, err
	}
	reply.Total = int32(c)
	if reply.Total == 0 {
		return reply, nil
	}
	reply.List = make([]*v1.WeiboUsersReply_One, 0)
	for _, doc := range v {
		reply.List = append(reply.List, &v1.WeiboUsersReply_One{
			Uid:    doc.Uid,
			UrlPc:  fmt.Sprintf("%s%s", root, doc.Uid),
			UrlWeb: fmt.Sprintf("%s/u/%s", pc, doc.Uid),
			Desc:   doc.Description,
		})
	}
	return reply, err
}
