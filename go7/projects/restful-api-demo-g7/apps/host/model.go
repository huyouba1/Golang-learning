package host

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"time"
)

var (
	validate = validator.New()
)

func NewHostSet() *HostSet {
	return &HostSet{
		Items: []*Host{},
	}
}

type HostSet struct {
	Items []*Host `json:"items"`
	Total int     `json:"total"`
}

func (s *HostSet) Add(item *Host) {
	s.Items = append(s.Items, item)
}

func NewHost() *Host {
	return &Host{
		Resource: &Resource{},
		Describe: &Describe{},
	}
}

// Host 模型的定义
type Host struct {
	// 匿名结构体定义
	// 资源的公共属性部分
	*Resource
	// 资源独有属性部分
	*Describe
}

// 对象全量更新
func (h *Host) Put(obj *Host) error {
	if obj.Id != h.Id {
		return fmt.Errorf("id not equal")
	}
	*h.Describe = *obj.Describe
}

// 对象局部更新
func (h *Host) Patch(obj *Host) {
	if obj.
}

func (h *Host) Validate() error {
	validate.Struct(h)
	return nil
}

func (h *Host) InjectDefault() {
	if h.CreateAt == 0 {
		h.CreateAt = time.Now().UnixMilli()
	}
}

type Vendor int

const (
	// 枚举的默认值
	PRIVATE_IDC Vendor = iota
	// 阿里云
	ALIYUN
	// 腾讯云
	TXYUN
)

type Resource struct {
	Id          string            `json:"id"  validate:"required"`     // 全局唯一Id
	Vendor      Vendor            `json:"vendor"`                      // 厂商
	Region      string            `json:"region"  validate:"required"` // 地域
	CreateAt    int64             `json:"create_at"`                   // 创建时间
	ExpireAt    int64             `json:"expire_at"`                   // 过期时间
	Type        string            `json:"type"  validate:"required"`   // 规格
	Name        string            `json:"name"  validate:"required"`   // 名称
	Description string            `json:"description"`                 // 描述
	Status      string            `json:"status"`                      // 服务商中的状态
	Tags        map[string]string `json:"tags"`                        // 标签
	UpdateAt    int64             `json:"update_at"`                   // 更新时间
	SyncAt      int64             `json:"sync_at"`                     // 同步时间
	Account     string            `json:"accout"`                      // 资源的所属账号
	PublicIP    string            `json:"public_ip"`                   // 公网IP
	PrivateIP   string            `json:"private_ip"`                  // 内网IP
}

type Describe struct {
	CPU          int    `json:"cpu" validate:"required"`    // 核数
	Memory       int    `json:"memory" validate:"required"` // 内存
	GPUAmount    int    `json:"gpu_amount"`                 // GPU数量
	GPUSpec      string `json:"gpu_spec"`                   // GPU类型
	OSType       string `json:"os_type"`                    // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`                    // 操作系统名称
	SerialNumber string `json:"serial_number"`              // 序列号
}

func NewQueryHostFromHTTP(r *http.Request) *QueryHostRequst {
	req := NewQueryHostRequst()
	// query string
	qs := r.URL.Query()
	pss := qs.Get("page_size")
	if pss != "" {
		req.Pagesize, _ = strconv.Atoi(pss)
	}

	pns := qs.Get("page_number")
	if pns != "" {
		req.PageNumber, _ = strconv.Atoi(pns)
	}

	req.Keywords = qs.Get("kws")
	return req
}

func NewQueryHostRequst() *QueryHostRequst {
	return &QueryHostRequst{
		Pagesize:   20,
		PageNumber: 1,
	}
}

type QueryHostRequst struct {
	Pagesize   int    `json:"page_size"`
	PageNumber int    `json:"page_number"`
	Keywords   string `json:"kws"`
}

func (req *QueryHostRequst) OffSet() int64 {
	return int64((req.PageNumber - 1) * req.Pagesize)
}

func (req *QueryHostRequst) GetPagesize() uint {
	return uint(req.Pagesize)
}

func NewDescribeHostRequestWithId(id string) *DescribeHostRequest {
	return &DescribeHostRequest{
		Id: id,
	}
}

type DescribeHostRequest struct {
	Id string
}

type UPDATE_MODE string

const (
	// 全量更新
	UPDATE_MODE_PUT UPDATE_MODE = "put"
	// 局部更新
	UPDATE_MODE_PATCH UPDATE_MODE = "patch"
)

type UpdateHostRequest struct {
	Id         string      `json:"id"`
	UpdateMode UPDATE_MODE `json:"update_mode"`
	*Describe
}

type DeleteHostRequst struct {
	Id string
}
