package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type ChangeStatus struct {
	actions.Action
}

// 初始化
func (p *ChangeStatus) Init() *ChangeStatus {
	// 初始化父结构
	p.ParentInit()

	// 行为名称，当行为在表格行展示时，支持js表达式
	p.Name = "<%= (status==1 ? '禁用' : '启用') %>"

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	//  执行成功后刷新的组件
	p.Reload = "table"

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	// 当行为在表格行展示时，支持js表达式
	p.WithConfirm("确定要<%= (status==1 ? '禁用' : '启用') %>数据吗？", "", "pop")

	return p
}

/**
 * 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
 *
 * @return array
 */
func (p *ChangeStatus) GetApiParams() []string {
	return []string{
		"id",
		"status",
	}
}

// 执行行为句柄
func (p *ChangeStatus) Handle(ctx *builder.Context, query *gorm.DB) error {
	status := ctx.Query("status")
	if status == "" {
		return ctx.JSON(200, message.Error("参数错误！"))
	}

	var fieldStatus int
	if status == "1" {
		fieldStatus = 0
	} else {
		fieldStatus = 1
	}

	err := query.Update("status", fieldStatus).Error
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return ctx.JSON(200, message.Success("操作成功"))
}