package resources

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/fields/radio"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/model"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/actions"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
	"github.com/quarkcloudio/quark-go/v3/pkg/dal/db"
)

type Account struct {
	resource.Template
}

// 初始化
func (p *Account) Init(ctx *builder.Context) interface{} {

	// 标题
	p.Title = "个人设置"

	// 模型
	p.Model = &model.User{}

	return p
}

// 表单接口
func (p *Account) FormApi(ctx *builder.Context) string {

	// 获取行为实例
	actionInstance := actions.ChangeAccount()

	// 获取行为接口接收的参数
	params := actionInstance.GetApiParams()

	// 获取行为key
	uriKey := actionInstance.GetUriKey(actionInstance)

	return p.BuildActionApi(ctx, params, uriKey)
}

// 字段
func (p *Account) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{

		field.Image("avatar", "头像"),

		field.Text("nickname", "昵称").
			SetRules([]*rule.Rule{
				rule.New().SetRequired().SetMessage("昵称必须填写"),
			}),

		field.Text("email", "邮箱").
			SetRules([]*rule.Rule{
				rule.New().SetRequired().SetMessage("邮箱必须填写"),
			}),

		field.Text("phone", "手机号").
			SetRules([]*rule.Rule{
				rule.New().SetRequired().SetMessage("手机号必须填写"),
			}),

		field.Radio("sex", "性别").
			SetOptions([]*radio.Option{
				{
					Value: 1,
					Label: "男",
				},
				{
					Value: 2,
					Label: "女",
				},
			}).
			SetDefault(1),

		field.Password("password", "密码"),
	}
}

// 行为
func (p *Account) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		actions.ChangeAccount(),
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}

// 创建页面显示前回调
func (p *Account) BeforeCreating(ctx *builder.Context) map[string]interface{} {
	data := map[string]interface{}{}
	adminInfo, _ := (&model.User{}).GetAuthUser(ctx.Engine.GetConfig().AppKey, ctx.Token())
	db.Client.
		Model(p.Model).
		Where("id = ?", adminInfo.Id).
		First(&data)

	delete(data, "password")

	return data
}
