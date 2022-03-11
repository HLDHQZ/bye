// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comment.ext.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HotCommentReq struct {
	MovieId int64 `protobuf:"varint,2,opt,name=movieId" json:"movieId,omitempty"`
}

func (m *HotCommentReq) Reset()                    { *m = HotCommentReq{} }
func (m *HotCommentReq) String() string            { return proto.CompactTextString(m) }
func (*HotCommentReq) ProtoMessage()               {}
func (*HotCommentReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *HotCommentReq) GetMovieId() int64 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

type HotCommentRsp struct {
	Data *CommentData `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *HotCommentRsp) Reset()                    { *m = HotCommentRsp{} }
func (m *HotCommentRsp) String() string            { return proto.CompactTextString(m) }
func (*HotCommentRsp) ProtoMessage()               {}
func (*HotCommentRsp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *HotCommentRsp) GetData() *CommentData {
	if m != nil {
		return m.Data
	}
	return nil
}

type CommentData struct {
	Mini *CommentMini `protobuf:"bytes,1,opt,name=mini" json:"mini,omitempty"`
	Plus *CommentPlus `protobuf:"bytes,2,opt,name=plus" json:"plus,omitempty"`
}

func (m *CommentData) Reset()                    { *m = CommentData{} }
func (m *CommentData) String() string            { return proto.CompactTextString(m) }
func (*CommentData) ProtoMessage()               {}
func (*CommentData) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *CommentData) GetMini() *CommentMini {
	if m != nil {
		return m.Mini
	}
	return nil
}

func (m *CommentData) GetPlus() *CommentPlus {
	if m != nil {
		return m.Plus
	}
	return nil
}

type CommentMini struct {
	List  []*CommentRecord `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
	Total int64            `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
}

func (m *CommentMini) Reset()                    { *m = CommentMini{} }
func (m *CommentMini) String() string            { return proto.CompactTextString(m) }
func (*CommentMini) ProtoMessage()               {}
func (*CommentMini) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *CommentMini) GetList() []*CommentRecord {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *CommentMini) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type CommentPlus struct {
	List  []*CommentRecord `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
	Total int64            `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
}

func (m *CommentPlus) Reset()                    { *m = CommentPlus{} }
func (m *CommentPlus) String() string            { return proto.CompactTextString(m) }
func (*CommentPlus) ProtoMessage()               {}
func (*CommentPlus) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *CommentPlus) GetList() []*CommentRecord {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *CommentPlus) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type CommentRecord struct {
	Title     string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Content   string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	HeadImg   string `protobuf:"bytes,3,opt,name=headImg" json:"headImg,omitempty"`
	Nickname  string `protobuf:"bytes,4,opt,name=nickname" json:"nickname,omitempty"`
	CreateAt  string `protobuf:"bytes,5,opt,name=createAt" json:"createAt,omitempty"`
	UpNum     int64  `protobuf:"varint,6,opt,name=upNum" json:"upNum,omitempty"`
	CommentID int64  `protobuf:"varint,7,opt,name=commentID" json:"commentID,omitempty"`
}

func (m *CommentRecord) Reset()                    { *m = CommentRecord{} }
func (m *CommentRecord) String() string            { return proto.CompactTextString(m) }
func (*CommentRecord) ProtoMessage()               {}
func (*CommentRecord) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *CommentRecord) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CommentRecord) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CommentRecord) GetHeadImg() string {
	if m != nil {
		return m.HeadImg
	}
	return ""
}

func (m *CommentRecord) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *CommentRecord) GetCreateAt() string {
	if m != nil {
		return m.CreateAt
	}
	return ""
}

func (m *CommentRecord) GetUpNum() int64 {
	if m != nil {
		return m.UpNum
	}
	return 0
}

func (m *CommentRecord) GetCommentID() int64 {
	if m != nil {
		return m.CommentID
	}
	return 0
}

type MakeCommentReq struct {
	MovieId  int64  `protobuf:"varint,1,opt,name=movieId" json:"movieId,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	HeadImg  string `protobuf:"bytes,3,opt,name=headImg" json:"headImg,omitempty"`
	Nickname string `protobuf:"bytes,4,opt,name=nickname" json:"nickname,omitempty"`
	UserId   int64  `protobuf:"varint,5,opt,name=userId" json:"userId,omitempty"`
	Content  string `protobuf:"bytes,6,opt,name=content" json:"content,omitempty"`
}

func (m *MakeCommentReq) Reset()                    { *m = MakeCommentReq{} }
func (m *MakeCommentReq) String() string            { return proto.CompactTextString(m) }
func (*MakeCommentReq) ProtoMessage()               {}
func (*MakeCommentReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *MakeCommentReq) GetMovieId() int64 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

func (m *MakeCommentReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MakeCommentReq) GetHeadImg() string {
	if m != nil {
		return m.HeadImg
	}
	return ""
}

func (m *MakeCommentReq) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *MakeCommentReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *MakeCommentReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type MakeCommentRsp struct {
}

func (m *MakeCommentRsp) Reset()                    { *m = MakeCommentRsp{} }
func (m *MakeCommentRsp) String() string            { return proto.CompactTextString(m) }
func (*MakeCommentRsp) ProtoMessage()               {}
func (*MakeCommentRsp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

type UpNumCommentReq struct {
	CommentID int64 `protobuf:"varint,1,opt,name=commentID" json:"commentID,omitempty"`
}

func (m *UpNumCommentReq) Reset()                    { *m = UpNumCommentReq{} }
func (m *UpNumCommentReq) String() string            { return proto.CompactTextString(m) }
func (*UpNumCommentReq) ProtoMessage()               {}
func (*UpNumCommentReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *UpNumCommentReq) GetCommentID() int64 {
	if m != nil {
		return m.CommentID
	}
	return 0
}

type UpNumCommentRsp struct {
	UpNum int64 `protobuf:"varint,1,opt,name=upNum" json:"upNum,omitempty"`
}

func (m *UpNumCommentRsp) Reset()                    { *m = UpNumCommentRsp{} }
func (m *UpNumCommentRsp) String() string            { return proto.CompactTextString(m) }
func (*UpNumCommentRsp) ProtoMessage()               {}
func (*UpNumCommentRsp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *UpNumCommentRsp) GetUpNum() int64 {
	if m != nil {
		return m.UpNum
	}
	return 0
}

type MyCommentsReq struct {
	UserId int64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *MyCommentsReq) Reset()                    { *m = MyCommentsReq{} }
func (m *MyCommentsReq) String() string            { return proto.CompactTextString(m) }
func (*MyCommentsReq) ProtoMessage()               {}
func (*MyCommentsReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *MyCommentsReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type MyCommentsRsp struct {
	MyComments []*MyComment `protobuf:"bytes,1,rep,name=myComments" json:"myComments,omitempty"`
}

func (m *MyCommentsRsp) Reset()                    { *m = MyCommentsRsp{} }
func (m *MyCommentsRsp) String() string            { return proto.CompactTextString(m) }
func (*MyCommentsRsp) ProtoMessage()               {}
func (*MyCommentsRsp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *MyCommentsRsp) GetMyComments() []*MyComment {
	if m != nil {
		return m.MyComments
	}
	return nil
}

type MyComment struct {
	FilmImage string `protobuf:"bytes,1,opt,name=filmImage" json:"filmImage,omitempty"`
	FilmName  string `protobuf:"bytes,2,opt,name=filmName" json:"filmName,omitempty"`
	Score     string `protobuf:"bytes,3,opt,name=score" json:"score,omitempty"`
	CommentID int64  `protobuf:"varint,4,opt,name=commentID" json:"commentID,omitempty"`
	Content   string `protobuf:"bytes,5,opt,name=content" json:"content,omitempty"`
	UpNum     int64  `protobuf:"varint,6,opt,name=upNum" json:"upNum,omitempty"`
}

func (m *MyComment) Reset()                    { *m = MyComment{} }
func (m *MyComment) String() string            { return proto.CompactTextString(m) }
func (*MyComment) ProtoMessage()               {}
func (*MyComment) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *MyComment) GetFilmImage() string {
	if m != nil {
		return m.FilmImage
	}
	return ""
}

func (m *MyComment) GetFilmName() string {
	if m != nil {
		return m.FilmName
	}
	return ""
}

func (m *MyComment) GetScore() string {
	if m != nil {
		return m.Score
	}
	return ""
}

func (m *MyComment) GetCommentID() int64 {
	if m != nil {
		return m.CommentID
	}
	return 0
}

func (m *MyComment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *MyComment) GetUpNum() int64 {
	if m != nil {
		return m.UpNum
	}
	return 0
}

type DeleteCommentReq struct {
	CommentID int64 `protobuf:"varint,1,opt,name=commentID" json:"commentID,omitempty"`
}

func (m *DeleteCommentReq) Reset()                    { *m = DeleteCommentReq{} }
func (m *DeleteCommentReq) String() string            { return proto.CompactTextString(m) }
func (*DeleteCommentReq) ProtoMessage()               {}
func (*DeleteCommentReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

func (m *DeleteCommentReq) GetCommentID() int64 {
	if m != nil {
		return m.CommentID
	}
	return 0
}

type DeleteCommentRsp struct {
}

func (m *DeleteCommentRsp) Reset()                    { *m = DeleteCommentRsp{} }
func (m *DeleteCommentRsp) String() string            { return proto.CompactTextString(m) }
func (*DeleteCommentRsp) ProtoMessage()               {}
func (*DeleteCommentRsp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{14} }

func init() {
	proto.RegisterType((*HotCommentReq)(nil), "pb.HotCommentReq")
	proto.RegisterType((*HotCommentRsp)(nil), "pb.HotCommentRsp")
	proto.RegisterType((*CommentData)(nil), "pb.CommentData")
	proto.RegisterType((*CommentMini)(nil), "pb.CommentMini")
	proto.RegisterType((*CommentPlus)(nil), "pb.CommentPlus")
	proto.RegisterType((*CommentRecord)(nil), "pb.CommentRecord")
	proto.RegisterType((*MakeCommentReq)(nil), "pb.MakeCommentReq")
	proto.RegisterType((*MakeCommentRsp)(nil), "pb.MakeCommentRsp")
	proto.RegisterType((*UpNumCommentReq)(nil), "pb.UpNumCommentReq")
	proto.RegisterType((*UpNumCommentRsp)(nil), "pb.UpNumCommentRsp")
	proto.RegisterType((*MyCommentsReq)(nil), "pb.MyCommentsReq")
	proto.RegisterType((*MyCommentsRsp)(nil), "pb.MyCommentsRsp")
	proto.RegisterType((*MyComment)(nil), "pb.MyComment")
	proto.RegisterType((*DeleteCommentReq)(nil), "pb.DeleteCommentReq")
	proto.RegisterType((*DeleteCommentRsp)(nil), "pb.DeleteCommentRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CommentServiceExt service

type CommentServiceExtClient interface {
	// 精彩影评
	HotComment(ctx context.Context, in *HotCommentReq, opts ...client.CallOption) (*HotCommentRsp, error)
	// 进行评论
	MakeComment(ctx context.Context, in *MakeCommentReq, opts ...client.CallOption) (*MakeCommentRsp, error)
	// 评论up
	UpNumComment(ctx context.Context, in *UpNumCommentReq, opts ...client.CallOption) (*UpNumCommentRsp, error)
	// 我的评论
	MyComments(ctx context.Context, in *MyCommentsReq, opts ...client.CallOption) (*MyCommentsRsp, error)
	// 删除评论
	DeleteComment(ctx context.Context, in *DeleteCommentReq, opts ...client.CallOption) (*DeleteCommentRsp, error)
}

type commentServiceExtClient struct {
	c           client.Client
	serviceName string
}

func NewCommentServiceExtClient(serviceName string, c client.Client) CommentServiceExtClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "pb"
	}
	return &commentServiceExtClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *commentServiceExtClient) HotComment(ctx context.Context, in *HotCommentReq, opts ...client.CallOption) (*HotCommentRsp, error) {
	req := c.c.NewRequest(c.serviceName, "CommentServiceExt.HotComment", in)
	out := new(HotCommentRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceExtClient) MakeComment(ctx context.Context, in *MakeCommentReq, opts ...client.CallOption) (*MakeCommentRsp, error) {
	req := c.c.NewRequest(c.serviceName, "CommentServiceExt.MakeComment", in)
	out := new(MakeCommentRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceExtClient) UpNumComment(ctx context.Context, in *UpNumCommentReq, opts ...client.CallOption) (*UpNumCommentRsp, error) {
	req := c.c.NewRequest(c.serviceName, "CommentServiceExt.UpNumComment", in)
	out := new(UpNumCommentRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceExtClient) MyComments(ctx context.Context, in *MyCommentsReq, opts ...client.CallOption) (*MyCommentsRsp, error) {
	req := c.c.NewRequest(c.serviceName, "CommentServiceExt.MyComments", in)
	out := new(MyCommentsRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceExtClient) DeleteComment(ctx context.Context, in *DeleteCommentReq, opts ...client.CallOption) (*DeleteCommentRsp, error) {
	req := c.c.NewRequest(c.serviceName, "CommentServiceExt.DeleteComment", in)
	out := new(DeleteCommentRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CommentServiceExt service

type CommentServiceExtHandler interface {
	// 精彩影评
	HotComment(context.Context, *HotCommentReq, *HotCommentRsp) error
	// 进行评论
	MakeComment(context.Context, *MakeCommentReq, *MakeCommentRsp) error
	// 评论up
	UpNumComment(context.Context, *UpNumCommentReq, *UpNumCommentRsp) error
	// 我的评论
	MyComments(context.Context, *MyCommentsReq, *MyCommentsRsp) error
	// 删除评论
	DeleteComment(context.Context, *DeleteCommentReq, *DeleteCommentRsp) error
}

func RegisterCommentServiceExtHandler(s server.Server, hdlr CommentServiceExtHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&CommentServiceExt{hdlr}, opts...))
}

type CommentServiceExt struct {
	CommentServiceExtHandler
}

func (h *CommentServiceExt) HotComment(ctx context.Context, in *HotCommentReq, out *HotCommentRsp) error {
	return h.CommentServiceExtHandler.HotComment(ctx, in, out)
}

func (h *CommentServiceExt) MakeComment(ctx context.Context, in *MakeCommentReq, out *MakeCommentRsp) error {
	return h.CommentServiceExtHandler.MakeComment(ctx, in, out)
}

func (h *CommentServiceExt) UpNumComment(ctx context.Context, in *UpNumCommentReq, out *UpNumCommentRsp) error {
	return h.CommentServiceExtHandler.UpNumComment(ctx, in, out)
}

func (h *CommentServiceExt) MyComments(ctx context.Context, in *MyCommentsReq, out *MyCommentsRsp) error {
	return h.CommentServiceExtHandler.MyComments(ctx, in, out)
}

func (h *CommentServiceExt) DeleteComment(ctx context.Context, in *DeleteCommentReq, out *DeleteCommentRsp) error {
	return h.CommentServiceExtHandler.DeleteComment(ctx, in, out)
}

func init() { proto.RegisterFile("comment.ext.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5f, 0x8b, 0x13, 0x3f,
	0x14, 0xfd, 0xa5, 0xff, 0xf6, 0xd7, 0x5b, 0xeb, 0x6e, 0xe3, 0x22, 0xc3, 0xe0, 0x43, 0x89, 0xc8,
	0xd6, 0x07, 0xab, 0xd4, 0x82, 0x20, 0x28, 0x88, 0x15, 0xac, 0xd0, 0x45, 0x22, 0xe2, 0xf3, 0x74,
	0x1a, 0xd7, 0xb0, 0x93, 0x99, 0x38, 0xc9, 0x2c, 0xeb, 0x37, 0xd2, 0x67, 0x3f, 0x83, 0xdf, 0x4b,
	0x92, 0x99, 0xce, 0x24, 0xb3, 0x7d, 0x10, 0x7d, 0x3c, 0xf7, 0xdc, 0x93, 0xb9, 0xf7, 0xe4, 0x4c,
	0x60, 0x12, 0x67, 0x42, 0xb0, 0x54, 0xcf, 0xd9, 0xb5, 0x9e, 0xcb, 0x3c, 0xd3, 0x19, 0xee, 0xc8,
	0x2d, 0x79, 0x08, 0xe3, 0xb7, 0x99, 0x7e, 0x5d, 0x72, 0x94, 0x7d, 0xc5, 0x01, 0x1c, 0x89, 0xec,
	0x8a, 0xb3, 0xf5, 0x2e, 0xe8, 0x4c, 0xd1, 0xac, 0x4b, 0xf7, 0x90, 0x2c, 0xbd, 0x56, 0x25, 0xf1,
	0x7d, 0xe8, 0xed, 0x22, 0x1d, 0x05, 0x68, 0x8a, 0x66, 0xa3, 0xc5, 0xf1, 0x5c, 0x6e, 0xe7, 0x15,
	0xbb, 0x8a, 0x74, 0x44, 0x2d, 0x49, 0x3e, 0xc1, 0xc8, 0x29, 0x1a, 0x8d, 0xe0, 0x29, 0x3f, 0xa0,
	0xd9, 0xf0, 0x94, 0x53, 0x4b, 0x9a, 0x26, 0x99, 0x14, 0xca, 0x0e, 0xe0, 0x37, 0xbd, 0x4f, 0x0a,
	0x45, 0x2d, 0x49, 0xde, 0xd5, 0x07, 0x1b, 0x25, 0x7e, 0x00, 0xbd, 0x84, 0x2b, 0x1d, 0xa0, 0x69,
	0x77, 0x36, 0x5a, 0x4c, 0x1c, 0x0d, 0x65, 0x71, 0x96, 0xef, 0xa8, 0xa5, 0xf1, 0x29, 0xf4, 0x75,
	0xa6, 0xa3, 0xa4, 0x5a, 0xae, 0x04, 0xce, 0x59, 0xe6, 0x03, 0xff, 0x76, 0xd6, 0x2f, 0x04, 0x63,
	0xaf, 0xdb, 0xf6, 0x71, 0x9d, 0x30, 0xbb, 0xf4, 0x90, 0x96, 0xc0, 0x18, 0x1d, 0x67, 0xa9, 0x66,
	0xa9, 0xb6, 0xfa, 0x21, 0xdd, 0x43, 0xc3, 0x7c, 0x61, 0xd1, 0x6e, 0x2d, 0x2e, 0x82, 0x6e, 0xc9,
	0x54, 0x10, 0x87, 0xf0, 0x7f, 0xca, 0xe3, 0xcb, 0x34, 0x12, 0x2c, 0xe8, 0x59, 0xaa, 0xc6, 0x86,
	0x8b, 0x73, 0x16, 0x69, 0xf6, 0x4a, 0x07, 0xfd, 0x92, 0xdb, 0x63, 0x33, 0x41, 0x21, 0xcf, 0x0b,
	0x11, 0x0c, 0xca, 0x49, 0x2d, 0xc0, 0xf7, 0x60, 0x58, 0x85, 0x62, 0xbd, 0x0a, 0x8e, 0x2c, 0xd3,
	0x14, 0xc8, 0x77, 0x04, 0xb7, 0x37, 0xd1, 0x25, 0x3b, 0x9c, 0x0d, 0xe4, 0x65, 0xa3, 0x59, 0xb1,
	0xd3, 0x5a, 0xf1, 0x2f, 0x16, 0xb9, 0x0b, 0x83, 0x42, 0xb1, 0x7c, 0xbd, 0xb3, 0x6b, 0x74, 0x69,
	0x85, 0x5c, 0xc3, 0x06, 0x9e, 0x61, 0xe4, 0xc4, 0x9f, 0x54, 0x49, 0xf2, 0x18, 0x8e, 0x3f, 0x9a,
	0x1d, 0x9d, 0xe1, 0xbd, 0x6d, 0x51, 0x7b, 0xdb, 0xb3, 0x96, 0x40, 0xc9, 0xc6, 0x34, 0xe4, 0x98,
	0x46, 0xce, 0x60, 0xbc, 0xf9, 0x56, 0x75, 0x29, 0x73, 0x6e, 0x33, 0x2e, 0x72, 0xc7, 0x25, 0x2f,
	0xbd, 0x46, 0x25, 0xf1, 0x23, 0x00, 0x51, 0x17, 0xaa, 0x6c, 0x8d, 0x4d, 0xb6, 0xea, 0x36, 0xea,
	0x34, 0x90, 0x1f, 0x08, 0x86, 0x35, 0x63, 0xa6, 0xff, 0xcc, 0x13, 0xb1, 0x16, 0xd1, 0xc5, 0x3e,
	0x47, 0x4d, 0xc1, 0xd8, 0x69, 0xc0, 0xb9, 0xb1, 0xb3, 0xbc, 0x81, 0x1a, 0x9b, 0x35, 0x54, 0x9c,
	0xe5, 0xac, 0xba, 0x82, 0x12, 0xf8, 0x6e, 0xf4, 0x5a, 0x6e, 0xb8, 0x56, 0xf7, 0xfd, 0x6c, 0x1e,
	0x4c, 0x12, 0x79, 0x02, 0x27, 0x2b, 0x96, 0x30, 0xcd, 0xfe, 0xd8, 0x6f, 0xdc, 0x56, 0x28, 0xb9,
	0xf8, 0xd9, 0x81, 0x49, 0x05, 0x3f, 0xb0, 0xfc, 0x8a, 0xc7, 0xec, 0xcd, 0xb5, 0xc6, 0x4b, 0x80,
	0xe6, 0xd9, 0xc1, 0xf6, 0x67, 0xf4, 0x5e, 0xac, 0xb0, 0x5d, 0x52, 0x92, 0xfc, 0x87, 0x9f, 0xc1,
	0xc8, 0x89, 0x04, 0xc6, 0xd6, 0x67, 0x2f, 0xcd, 0xe1, 0x8d, 0x9a, 0x15, 0x3e, 0x87, 0x5b, 0x6e,
	0x10, 0xf0, 0x1d, 0xd3, 0xd5, 0xca, 0x52, 0x78, 0xb3, 0x68, 0xb5, 0x4b, 0x80, 0xe6, 0xca, 0xcb,
	0x51, 0xbd, 0xac, 0x84, 0xed, 0x92, 0x55, 0xbd, 0x80, 0xb1, 0x67, 0x05, 0x3e, 0x35, 0x5d, 0x6d,
	0x3f, 0xc3, 0x03, 0x55, 0x23, 0xdf, 0x0e, 0xec, 0x63, 0xfe, 0xf4, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x79, 0xa3, 0x39, 0xbd, 0xe1, 0x05, 0x00, 0x00,
}
