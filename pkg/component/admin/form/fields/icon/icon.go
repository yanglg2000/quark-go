package icon

import (
	"encoding/json"
	"strings"

	"github.com/quarkcms/quark-go/pkg/component/admin/component"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/when"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/rule"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
	"github.com/quarkcms/quark-go/pkg/untils"
)

type Icon struct {
	ComponentKey string `json:"componentkey"` // 组件标识
	Component    string `json:"component"`    // 组件名称

	Colon         bool        `json:"colon,omitempty"`        // 配合 label 属性使用，表示是否显示 label 后面的冒号
	Extra         string      `json:"extra,omitempty"`        // 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
	HasFeedback   bool        `json:"hasFeedback,omitempty"`  // 配合 validateStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
	Help          string      `json:"help,omitempty"`         // 提示信息，如不设置，则会根据校验规则自动生成
	Hidden        bool        `json:"hidden,omitempty"`       // 是否隐藏字段（依然会收集和校验字段）
	InitialValue  interface{} `json:"initialValue,omitempty"` // 设置子元素默认值，如果与 Form 的 initialValues 冲突则以 Form 为准
	Label         string      `json:"label,omitempty"`        // label 标签的文本
	LabelAlign    string      `json:"labelAlign,omitempty"`   // 标签文本对齐方式
	LabelCol      interface{} `json:"labelCol,omitempty"`     // label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。你可以通过 Form 的 labelCol 进行统一设置，不会作用于嵌套 Item。当和 Form 同时设置时，以 Item 为准
	Name          string      `json:"name,omitempty"`         // 字段名，支持数组
	NoStyle       bool        `json:"noStyle,omitempty"`      // 为 true 时不带样式，作为纯字段控件使用
	Required      bool        `json:"required,omitempty"`     // 必填样式设置。如不设置，则会根据校验规则自动生成
	Tooltip       string      `json:"tooltip,omitempty"`      // 会在 label 旁增加一个 icon，悬浮后展示配置的信息
	ValuePropName string      `json:"valuePropName"`          // 子节点的值的属性，如 Switch 的是 'checked'。该属性为 getValueProps 的封装，自定义 getValueProps 后会失效
	WrapperCol    interface{} `json:"wrapperCol"`             // 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。你可以通过 Form 的 wrapperCol 进行统一设置，不会作用于嵌套 Item。当和 Form 同时设置时，以 Item 为准

	Api            string        `json:"api,omitempty"` // 获取数据接口
	Ignore         bool          `json:"ignore"`        // 是否忽略保存到数据库，默认为 false
	Rules          []*rule.Rule  `json:"-"`             // 全局校验规则
	CreationRules  []*rule.Rule  `json:"-"`             // 创建页校验规则
	UpdateRules    []*rule.Rule  `json:"-"`             // 编辑页校验规则
	FrontendRules  []*rule.Rule  `json:"frontendRules"` // 前端校验规则，设置字段的校验逻辑
	When           *when.When    `json:"when"`          //
	WhenItem       []*when.Item  `json:"-"`             //
	ShowOnIndex    bool          `json:"-"`             // 在列表页展示
	ShowOnDetail   bool          `json:"-"`             // 在详情页展示
	ShowOnCreation bool          `json:"-"`             // 在创建页面展示
	ShowOnUpdate   bool          `json:"-"`             // 在编辑页面展示
	ShowOnExport   bool          `json:"-"`             // 在导出的Excel上展示
	ShowOnImport   bool          `json:"-"`             // 在导入Excel上展示
	Editable       bool          `json:"-"`             // 表格上是否可编辑
	Column         *table.Column `json:"-"`             // 表格列
	Callback       interface{}   `json:"-"`             // 回调函数

	DefaultValue interface{}            `json:"defaultValue,omitempty"` // 默认选中的选项
	Disabled     bool                   `json:"disabled,omitempty"`     // 整组失效
	Style        map[string]interface{} `json:"style,omitempty"`        // 自定义样式
	Value        interface{}            `json:"value,omitempty"`        // 指定选中项,string[] | number[]
	Size         string                 `json:"size"`                   // 大小，large | middle | small
	AllowClear   bool                   `json:"allowClear"`             // 是否支持清除，默认true
	Options      []string               `json:"options"`                // 可选项数据源
}

// 初始化组件
func New() *Icon {
	return (&Icon{}).Init()
}

// 初始化
func (p *Icon) Init() *Icon {
	p.Component = "iconField"
	p.Colon = true
	p.LabelAlign = "right"
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = true
	p.ShowOnImport = true
	p.Column = (&table.Column{}).Init()
	p.Options = []string{
		"icon-database", "icon-sever", "icon-mobile", "icon-tablet", "icon-redenvelope",
		"icon-book", "icon-filedone", "icon-reconciliation", "icon-file-exception",
		"icon-filesync", "icon-filesearch", "icon-solution", "icon-fileprotect",
		"icon-file-add", "icon-file-excel", "icon-file-exclamation", "icon-file-pdf",
		"icon-file-image", "icon-file-markdown", "icon-file-unknown", "icon-file-ppt",
		"icon-file-word", "icon-file", "icon-file-zip", "icon-file-text", "icon-file-copy",
		"icon-snippets", "icon-audit", "icon-diff", "icon-Batchfolding", "icon-securityscan",
		"icon-propertysafety", "icon-insurance", "icon-alert", "icon-delete", "icon-hourglass",
		"icon-bulb", "icon-experiment", "icon-bell", "icon-trophy", "icon-rest", "icon-USB",
		"icon-skin", "icon-home", "icon-bank", "icon-filter", "icon-funnelplot", "icon-like",
		"icon-unlike", "icon-unlock", "icon-lock", "icon-customerservice", "icon-flag",
		"icon-moneycollect", "icon-medicinebox", "icon-shop", "icon-rocket", "icon-shopping",
		"icon-folder", "icon-folder-open", "icon-folder-add", "icon-deploymentunit",
		"icon-accountbook", "icon-contacts", "icon-carryout", "icon-calendar-check",
		"icon-calendar", "icon-scan", "icon-select", "icon-boxplot", "icon-build", "icon-sliders",
		"icon-laptop", "icon-barcode", "icon-camera", "icon-cluster", "icon-gateway", "icon-car",
		"icon-printer", "icon-read", "icon-cloud-server", "icon-cloud-upload", "icon-cloud",
		"icon-cloud-download", "icon-cloud-sync", "icon-video", "icon-notification", "icon-sound",
		"icon-radarchart", "icon-qrcode", "icon-fund", "icon-image", "icon-mail", "icon-table",
		"icon-idcard", "icon-creditcard", "icon-heart", "icon-block", "icon-error", "icon-star",
		"icon-gold", "icon-heatmap", "icon-wifi", "icon-attachment", "icon-edit", "icon-key",
		"icon-api", "icon-disconnect", "icon-highlight", "icon-monitor", "icon-link", "icon-man",
		"icon-percentage", "icon-pushpin", "icon-phone", "icon-shake", "icon-tag", "icon-wrench",
		"icon-tags", "icon-scissor", "icon-mr", "icon-share", "icon-branches", "icon-fork", "icon-shrink",
		"icon-arrawsalt", "icon-verticalright", "icon-verticalleft", "icon-right", "icon-left",
		"icon-up", "icon-down", "icon-fullscreen", "icon-fullscreen-exit", "icon-doubleleft",
		"icon-doubleright", "icon-arrowright", "icon-arrowup", "icon-arrowleft", "icon-arrowdown",
		"icon-upload", "icon-colum-height", "icon-vertical-align-botto", "icon-vertical-align-middl",
		"icon-totop", "icon-vertical-align-top", "icon-download", "icon-sort-descending",
		"icon-sort-ascending", "icon-fall", "icon-swap", "icon-stock", "icon-rise", "icon-indent",
		"icon-outdent", "icon-menu", "icon-unorderedlist", "icon-orderedlist", "icon-align-right",
		"icon-align-center", "icon-align-left", "icon-pic-center", "icon-pic-right", "icon-pic-left",
		"icon-bold", "icon-font-colors", "icon-exclaimination", "icon-font-size", "icon-check-circle",
		"icon-infomation", "icon-CI", "icon-line-height", "icon-Dollar", "icon-strikethrough", "icon-compass",
		"icon-underline", "icon-close-circle", "icon-number", "icon-frown", "icon-italic", "icon-info-circle",
		"icon-code", "icon-left-circle", "icon-column-width", "icon-down-circle", "icon-check", "icon-EURO",
		"icon-ellipsis", "icon-copyright", "icon-dash", "icon-minus-circle", "icon-close", "icon-meh",
		"icon-enter", "icon-plus-circle", "icon-line", "icon-play-circle", "icon-minus", "icon-question-circle",
		"icon-question", "icon-Pound", "icon-rollback", "icon-right-circle", "icon-small-dash", "icon-smile",
		"icon-pause", "icon-trademark", "icon-bg-colors", "icon-time-circle", "icon-crown", "icon-timeout",
		"icon-drag", "icon-earth", "icon-desktop", "icon-YUAN", "icon-gift", "icon-up-circle", "icon-stop",
		"icon-warning-circle", "icon-fire", "icon-sync", "icon-thunderbolt", "icon-transaction",
		"icon-alipay", "icon-undo", "icon-taobao", "icon-redo", "icon-wechat-fill", "icon-reload",
		"icon-comment", "icon-reloadtime", "icon-login", "icon-message", "icon-clear", "icon-dashboard",
		"icon-issuesclose", "icon-poweroff", "icon-logout", "icon-piechart", "icon-setting",
		"icon-eye", "icon-location", "icon-edit-square", "icon-export", "icon-save", "icon-Import",
		"icon-appstore", "icon-close-square", "icon-down-square", "icon-layout", "icon-left-square",
		"icon-play-square", "icon-control", "icon-codelibrary", "icon-detail", "icon-minus-square",
		"icon-plus-square", "icon-right-square", "icon-project", "icon-wallet", "icon-up-square",
		"icon-calculator", "icon-interation", "icon-check-square", "icon-border", "icon-border-outer",
		"icon-border-top", "icon-border-bottom", "icon-border-left", "icon-border-right", "icon-border-inner",
		"icon-border-verticle", "icon-border-horizontal", "icon-radius-bottomleft", "icon-radius-bottomright",
		"icon-radius-upleft", "icon-radius-upright", "icon-radius-setting", "icon-adduser", "icon-deleteteam",
		"icon-deleteuser", "icon-addteam", "icon-user", "icon-team", "icon-areachart", "icon-linechart",
		"icon-barchart", "icon-pointmap", "icon-container", "icon-atom", "icon-zanwutupian", "icon-safetycertificate",
		"icon-password", "icon-article", "icon-page", "icon-plugin", "icon-admin", "icon-banner",
	}

	p.SetWidth(200)
	p.SetDefault("")
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 设置Key
func (p *Icon) SetKey(key string, crypt bool) *Icon {
	p.ComponentKey = untils.MakeKey(key, crypt)

	return p
}

// 会在 label 旁增加一个 icon，悬浮后展示配置的信息
func (p *Icon) SetTooltip(tooltip string) *Icon {
	p.Tooltip = tooltip

	return p
}

// 配合 label 属性使用，表示是否显示 label 后面的冒号
func (p *Icon) SetColon(colon bool) *Icon {
	p.Colon = colon
	return p
}

// 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
func (p *Icon) SetExtra(extra string) *Icon {
	p.Extra = extra
	return p
}

// 配合 validateStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Icon) SetHasFeedback(hasFeedback bool) *Icon {
	p.HasFeedback = hasFeedback
	return p
}

// 配合 help 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Icon) SetHelp(help string) *Icon {
	p.Help = help
	return p
}

// 为 true 时不带样式，作为纯字段控件使用
func (p *Icon) SetNoStyle() *Icon {
	p.NoStyle = true
	return p
}

// label 标签的文本
func (p *Icon) SetLabel(label string) *Icon {
	p.Label = label

	return p
}

// 标签文本对齐方式
func (p *Icon) SetLabelAlign(align string) *Icon {
	p.LabelAlign = align
	return p
}

// label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。
// 你可以通过 Form 的 labelCol 进行统一设置。当和 Form 同时设置时，以 Item 为准
func (p *Icon) SetLabelCol(col interface{}) *Icon {
	p.LabelCol = col
	return p
}

// 字段名，支持数组
func (p *Icon) SetName(name string) *Icon {
	p.Name = name
	return p
}

// 是否必填，如不设置，则会根据校验规则自动生成
func (p *Icon) SetRequired() *Icon {
	p.Required = true
	return p
}

// 获取前端验证规则
func (p *Icon) GetFrontendRules(path string) *Icon {
	var (
		frontendRules []*rule.Rule
		rules         []*rule.Rule
		creationRules []*rule.Rule
		updateRules   []*rule.Rule
	)

	uri := strings.Split(path, "/")
	isCreating := (uri[len(uri)-1] == "create") || (uri[len(uri)-1] == "store")
	isEditing := (uri[len(uri)-1] == "edit") || (uri[len(uri)-1] == "update")

	if len(p.Rules) > 0 {
		rules = rule.ConvertToFrontendRules(p.Rules)
	}
	if isCreating && len(p.CreationRules) > 0 {
		creationRules = rule.ConvertToFrontendRules(p.CreationRules)
	}
	if isEditing && len(p.UpdateRules) > 0 {
		updateRules = rule.ConvertToFrontendRules(p.UpdateRules)
	}
	if len(rules) > 0 {
		frontendRules = append(frontendRules, rules...)
	}
	if len(creationRules) > 0 {
		frontendRules = append(frontendRules, creationRules...)
	}
	if len(updateRules) > 0 {
		frontendRules = append(frontendRules, updateRules...)
	}

	p.FrontendRules = frontendRules

	return p
}

// 校验规则，设置字段的校验逻辑
func (p *Icon) SetRules(rules []*rule.Rule) *Icon {
	p.Rules = rules

	return p
}

// 校验规则，只在创建表单提交时生效
func (p *Icon) SetCreationRules(rules []*rule.Rule) *Icon {
	p.CreationRules = rules

	return p
}

// 校验规则，只在更新表单提交时生效
func (p *Icon) SetUpdateRules(rules []*rule.Rule) *Icon {
	p.UpdateRules = rules

	return p
}

// 子节点的值的属性，如 Switch 的是 "checked"
func (p *Icon) SetValuePropName(valuePropName string) *Icon {
	p.ValuePropName = valuePropName
	return p
}

// 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。
// 你可以通过 Form 的 wrapperCol 进行统一设置。当和 Form 同时设置时，以 Item 为准。
func (p *Icon) SetWrapperCol(col interface{}) *Icon {
	p.WrapperCol = col
	return p
}

// 设置保存值。
func (p *Icon) SetValue(value interface{}) *Icon {
	p.Value = value
	return p
}

// 设置默认值。
func (p *Icon) SetDefault(value interface{}) *Icon {
	p.DefaultValue = value
	return p
}

// 是否禁用状态，默认为 false
func (p *Icon) SetDisabled(disabled bool) *Icon {
	p.Disabled = disabled
	return p
}

// 是否忽略保存到数据库，默认为 false
func (p *Icon) SetIgnore(ignore bool) *Icon {
	p.Ignore = ignore
	return p
}

// 表单联动
func (p *Icon) SetWhen(value ...any) *Icon {
	w := when.New()
	i := when.NewItem()
	var operator string
	var option any

	if len(value) == 2 {
		operator = "="
		option = value[0]
		callback := value[1].(func() interface{})

		i.Body = callback()
	}

	if len(value) == 3 {
		operator = value[0].(string)
		option = value[1]
		callback := value[2].(func() interface{})

		i.Body = callback()
	}

	getOption := untils.InterfaceToString(option)

	switch operator {
	case "=":
		i.Condition = "<%=String(" + p.Name + ") === '" + getOption + "' %>"
		break
	case ">":
		i.Condition = "<%=String(" + p.Name + ") > '" + getOption + "' %>"
		break
	case "<":
		i.Condition = "<%=String(" + p.Name + ") < '" + getOption + "' %>"
		break
	case "<=":
		i.Condition = "<%=String(" + p.Name + ") <= '" + getOption + "' %>"
		break
	case ">=":
		i.Condition = "<%=String(" + p.Name + ") => '" + getOption + "' %>"
		break
	case "has":
		i.Condition = "<%=(String(" + p.Name + ").indexOf('" + getOption + "') !=-1) %>"
		break
	case "in":
		jsonStr, _ := json.Marshal(option)
		i.Condition = "<%=(" + string(jsonStr) + ".indexOf(" + p.Name + ") !=-1) %>"
		break
	default:
		i.Condition = "<%=String(" + p.Name + ") === '" + getOption + "' %>"
		break
	}

	i.ConditionName = p.Name
	i.ConditionOperator = operator
	i.Option = option
	p.WhenItem = append(p.WhenItem, i)
	p.When = w.SetItems(p.WhenItem)

	return p
}

// Specify that the element should be hidden from the index view.
func (p *Icon) HideFromIndex(callback bool) *Icon {
	p.ShowOnIndex = !callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Icon) HideFromDetail(callback bool) *Icon {
	p.ShowOnDetail = !callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Icon) HideWhenCreating(callback bool) *Icon {
	p.ShowOnCreation = !callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *Icon) HideWhenUpdating(callback bool) *Icon {
	p.ShowOnUpdate = !callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Icon) HideWhenExporting(callback bool) *Icon {
	p.ShowOnExport = !callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Icon) HideWhenImporting(callback bool) *Icon {
	p.ShowOnImport = !callback

	return p
}

// Specify that the element should be hidden from the index view.
func (p *Icon) OnIndexShowing(callback bool) *Icon {
	p.ShowOnIndex = callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Icon) OnDetailShowing(callback bool) *Icon {
	p.ShowOnDetail = callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Icon) ShowOnCreating(callback bool) *Icon {
	p.ShowOnCreation = callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *Icon) ShowOnUpdating(callback bool) *Icon {
	p.ShowOnUpdate = callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Icon) ShowOnExporting(callback bool) *Icon {
	p.ShowOnExport = callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Icon) ShowOnImporting(callback bool) *Icon {
	p.ShowOnImport = callback

	return p
}

// Specify that the element should only be shown on the index view.
func (p *Icon) OnlyOnIndex() *Icon {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on the detail view.
func (p *Icon) OnlyOnDetail() *Icon {
	p.ShowOnIndex = false
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on forms.
func (p *Icon) OnlyOnForms() *Icon {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on export file.
func (p *Icon) OnlyOnExport() *Icon {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on import file.
func (p *Icon) OnlyOnImport() *Icon {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = true

	return p
}

// Specify that the element should be hidden from forms.
func (p *Icon) ExceptOnForms() *Icon {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = true

	return p
}

// Check for showing when updating.
func (p *Icon) IsShownOnUpdate() bool {
	return p.ShowOnUpdate
}

// Check showing on index.
func (p *Icon) IsShownOnIndex() bool {
	return p.ShowOnIndex
}

// Check showing on detail.
func (p *Icon) IsShownOnDetail() bool {
	return p.ShowOnDetail
}

// Check for showing when creating.
func (p *Icon) IsShownOnCreation() bool {
	return p.ShowOnCreation
}

// Check for showing when exporting.
func (p *Icon) IsShownOnExport() bool {
	return p.ShowOnExport
}

// Check for showing when importing.
func (p *Icon) IsShownOnImport() bool {
	return p.ShowOnImport
}

// 设置为可编辑列
func (p *Icon) SetEditable(editable bool) *Icon {
	p.Editable = editable

	return p
}

// 闭包，透传表格列的属性
func (p *Icon) SetColumn(f func(column *table.Column) *table.Column) *Icon {
	p.Column = f(p.Column)

	return p
}

// 当前列值的枚举 valueEnum
func (p *Icon) GetValueEnum() map[interface{}]interface{} {
	data := map[interface{}]interface{}{}

	return data
}

// 设置回调函数
func (p *Icon) SetCallback(closure func() interface{}) *Icon {
	if closure != nil {
		p.Callback = closure
	}

	return p
}

// 获取回调函数
func (p *Icon) GetCallback() interface{} {
	return p.Callback
}

// 获取数据接口
func (p *Icon) SetApi(api string) *Icon {
	p.Api = api
	return p
}

// 长度
func (p *Icon) SetWidth(width interface{}) *Icon {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["width"] = width
	p.Style = style

	return p
}

// 设置属性
func (p *Icon) SetOptions(options []string) *Icon {
	p.Options = options

	return p
}