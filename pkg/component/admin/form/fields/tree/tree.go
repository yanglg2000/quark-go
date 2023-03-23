package tree

import (
	"encoding/json"
	"strings"

	"github.com/quarkcms/quark-go/pkg/component/admin/component"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/when"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/rule"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
	"github.com/quarkcms/quark-go/pkg/untils"
)

type FieldNames struct {
	Title    string `json:"title"`
	Key      string `json:"key"`
	Children string `json:"children"`
}

type TreeData struct {
	Checkable       bool        `json:"checkable,omitempty"`       // 当树为 checkable 时，设置独立节点是否展示 Checkbox
	DisableCheckbox bool        `json:"disableCheckbox,omitempty"` // 禁掉 checkbox
	Disabled        bool        `json:"disabled,omitempty"`        // 禁掉响应
	Icon            interface{} `json:"icon,omitempty"`            // 自定义图标。可接收组件，props 为当前节点 props
	IsLeaf          bool        `json:"isLeaf,omitempty"`          // 设置为叶子节点 (设置了 loadData 时有效)。为 false 时会强制将其作为父节点
	Key             interface{} `json:"key"`                       // 被树的 (default)ExpandedKeys / (default)CheckedKeys / (default)SelectedKeys 属性所用。注意：整个树范围内的所有节点的 key 值不能重复！
	Selectable      bool        `json:"selectable,omitempty"`      // 设置节点是否可被选中
	Title           string      `json:"title"`                     // 标题
	Children        []*TreeData `json:"children,omitempty"`        // 子节点
}

type Tree struct {
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

	AutoExpandParent    bool                   `json:"autoExpandParent,omitempty"`    // 是否自动展开父节点
	BockNode            bool                   `json:"blockNode,omitempty"`           // 是否节点占据一行
	Checkable           bool                   `json:"checkable,omitempty"`           // 节点前添加 Checkbox 复选框
	CheckedKeys         []interface{}          `json:"checkedKeys,omitempty"`         // （受控）选中复选框的树节点（注意：父子节点有关联，如果传入父节点 key，则子节点自动选中；相应当子节点 key 都传入，父节点也自动选中。当设置 checkable 和 checkStrictly，它是一个有checked和halfChecked属性的对象，并且父子节点的选中与否不再关联
	CheckStrictly       bool                   `json:"checkStrictly,omitempty"`       // checkable 状态下节点选择完全受控（父子节点选中状态不再关联）
	DefaultCheckedKeys  []interface{}          `json:"defaultCheckedKeys,omitempty"`  // 默认选中复选框的树节点
	DefaultExpandAll    bool                   `json:"defaultExpandAll,omitempty"`    // 默认展开所有树节点
	DefaultExpandedKeys []interface{}          `json:"defaultExpandedKeys,omitempty"` // 默认展开指定的树节点
	DefaultExpandParent bool                   `json:"defaultExpandParent,omitempty"` // 默认展开父节点
	DefaultSelectedKeys []interface{}          `json:"defaultSelectedKeys,omitempty"` // 默认选中的树节点
	DefaultValue        interface{}            `json:"defaultValue,omitempty"`        // 默认选中的选项
	Disabled            bool                   `json:"disabled,omitempty"`            // 整组失效
	Draggable           bool                   `json:"draggable,omitempty"`           // 设置节点可拖拽，可以通过 icon: false 关闭拖拽提示图标
	ExpandedKeys        []interface{}          `json:"expandedKeys,omitempty"`        // （受控）展开指定的树节点
	FieldNames          *FieldNames            `json:"fieldNames,omitempty"`          // 自定义 options 中 label value children 的字段
	Height              int                    `json:"height,omitempty"`              // 设置虚拟滚动容器高度，设置后内部节点不再支持横向滚动
	Icon                interface{}            `json:"icon,omitempty"`                // 自定义树节点图标
	Multiple            bool                   `json:"multiple,omitempty"`            // 支持点选多个节点（节点本身）
	RootClassName       string                 `json:"rootClassName,omitempty"`       // 添加在 Tree 最外层的 className
	RootStyle           interface{}            `json:"rootStyle,omitempty"`           // 添加在 Tree 最外层的 style
	Selectable          bool                   `json:"selectable,omitempty"`          // 是否可选中
	SelectedKeys        []interface{}          `json:"selectedKeys,omitempty"`        // （受控）设置选中的树节点
	ShowIcon            bool                   `json:"showIcon,omitempty"`            // 是否展示 TreeNode title 前的图标，没有默认样式，如设置为 true，需要自行定义图标相关样式
	ShowLine            bool                   `json:"showLine,omitempty"`            // 是否展示连接线
	SwitcherIcon        interface{}            `json:"switcherIcon,omitempty"`        // 自定义树节点的展开/折叠图标
	TreeData            *TreeData              `json:"treeData,omitempty"`            // treeNodes 数据，如果设置则不需要手动构造 TreeNode 节点（value 在整个树范围内唯一）
	Value               interface{}            `json:"value,omitempty"`               // 指定当前选中的条目，多选时为一个数组。（value 数组引用未变化时，Select 不会更新）
	Virtual             bool                   `json:"virtual,omitempty"`             // 设置 false 时关闭虚拟滚动
	Style               map[string]interface{} `json:"style,omitempty"`               // 自定义样式
}

// 初始化组件
func New() *Tree {
	return (&Tree{}).Init()
}

// 初始化
func (p *Tree) Init() *Tree {
	p.Component = "treeField"
	p.Colon = true
	p.LabelAlign = "right"
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = true
	p.ShowOnImport = true
	p.Column = (&table.Column{}).Init()
	p.SetWidth(200)
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 设置Key
func (p *Tree) SetKey(key string, crypt bool) *Tree {
	p.ComponentKey = untils.MakeKey(key, crypt)

	return p
}

// 会在 label 旁增加一个 icon，悬浮后展示配置的信息
func (p *Tree) SetTooltip(tooltip string) *Tree {
	p.Tooltip = tooltip

	return p
}

// Field 的长度，我们归纳了常用的 Field 长度以及适合的场景，支持了一些枚举 "xs" , "s" , "m" , "l" , "x"
func (p *Tree) SetWidth(width interface{}) *Tree {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["width"] = width
	p.Style = style

	return p
}

// 配合 label 属性使用，表示是否显示 label 后面的冒号
func (p *Tree) SetColon(colon bool) *Tree {
	p.Colon = colon
	return p
}

// 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
func (p *Tree) SetExtra(extra string) *Tree {
	p.Extra = extra
	return p
}

// 配合 validateStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Tree) SetHasFeedback(hasFeedback bool) *Tree {
	p.HasFeedback = hasFeedback
	return p
}

// 配合 help 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Tree) SetHelp(help string) *Tree {
	p.Help = help
	return p
}

// 为 true 时不带样式，作为纯字段控件使用
func (p *Tree) SetNoStyle() *Tree {
	p.NoStyle = true
	return p
}

// label 标签的文本
func (p *Tree) SetLabel(label string) *Tree {
	p.Label = label

	return p
}

// 标签文本对齐方式
func (p *Tree) SetLabelAlign(align string) *Tree {
	p.LabelAlign = align
	return p
}

// label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。
// 你可以通过 Form 的 labelCol 进行统一设置。当和 Form 同时设置时，以 Item 为准
func (p *Tree) SetLabelCol(col interface{}) *Tree {
	p.LabelCol = col
	return p
}

// 字段名，支持数组
func (p *Tree) SetName(name string) *Tree {
	p.Name = name
	return p
}

// 是否必填，如不设置，则会根据校验规则自动生成
func (p *Tree) SetRequired() *Tree {
	p.Required = true
	return p
}

// 获取前端验证规则
func (p *Tree) GetFrontendRules(path string) *Tree {
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
func (p *Tree) SetRules(rules []*rule.Rule) *Tree {
	p.Rules = rules

	return p
}

// 校验规则，只在创建表单提交时生效
func (p *Tree) SetCreationRules(rules []*rule.Rule) *Tree {
	p.CreationRules = rules

	return p
}

// 校验规则，只在更新表单提交时生效
func (p *Tree) SetUpdateRules(rules []*rule.Rule) *Tree {
	p.UpdateRules = rules

	return p
}

// 子节点的值的属性，如 Switch 的是 "checked"
func (p *Tree) SetValuePropName(valuePropName string) *Tree {
	p.ValuePropName = valuePropName
	return p
}

// 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。
// 你可以通过 Form 的 wrapperCol 进行统一设置。当和 Form 同时设置时，以 Item 为准。
func (p *Tree) SetWrapperCol(col interface{}) *Tree {
	p.WrapperCol = col
	return p
}

// 指定当前选中的条目，多选时为一个数组。（value 数组引用未变化时，Select 不会更新）
func (p *Tree) SetValue(value interface{}) *Tree {
	p.Value = value
	return p
}

// 设置默认值。
func (p *Tree) SetDefault(value interface{}) *Tree {
	p.DefaultValue = value
	return p
}

// 是否禁用状态，默认为 false
func (p *Tree) SetDisabled(disabled bool) *Tree {
	p.Disabled = disabled
	return p
}

// 是否忽略保存到数据库，默认为 false
func (p *Tree) SetIgnore(ignore bool) *Tree {
	p.Ignore = ignore
	return p
}

// 表单联动
func (p *Tree) SetWhen(value ...any) *Tree {
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
func (p *Tree) HideFromIndex(callback bool) *Tree {
	p.ShowOnIndex = !callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Tree) HideFromDetail(callback bool) *Tree {
	p.ShowOnDetail = !callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Tree) HideWhenCreating(callback bool) *Tree {
	p.ShowOnCreation = !callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *Tree) HideWhenUpdating(callback bool) *Tree {
	p.ShowOnUpdate = !callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Tree) HideWhenExporting(callback bool) *Tree {
	p.ShowOnExport = !callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Tree) HideWhenImporting(callback bool) *Tree {
	p.ShowOnImport = !callback

	return p
}

// Specify that the element should be hidden from the index view.
func (p *Tree) OnIndexShowing(callback bool) *Tree {
	p.ShowOnIndex = callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Tree) OnDetailShowing(callback bool) *Tree {
	p.ShowOnDetail = callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Tree) ShowOnCreating(callback bool) *Tree {
	p.ShowOnCreation = callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *Tree) ShowOnUpdating(callback bool) *Tree {
	p.ShowOnUpdate = callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Tree) ShowOnExporting(callback bool) *Tree {
	p.ShowOnExport = callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Tree) ShowOnImporting(callback bool) *Tree {
	p.ShowOnImport = callback

	return p
}

// Specify that the element should only be shown on the index view.
func (p *Tree) OnlyOnIndex() *Tree {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on the detail view.
func (p *Tree) OnlyOnDetail() *Tree {
	p.ShowOnIndex = false
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on forms.
func (p *Tree) OnlyOnForms() *Tree {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on export file.
func (p *Tree) OnlyOnExport() *Tree {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on import file.
func (p *Tree) OnlyOnImport() *Tree {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = true

	return p
}

// Specify that the element should be hidden from forms.
func (p *Tree) ExceptOnForms() *Tree {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = true

	return p
}

// Check for showing when updating.
func (p *Tree) IsShownOnUpdate() bool {
	return p.ShowOnUpdate
}

// Check showing on index.
func (p *Tree) IsShownOnIndex() bool {
	return p.ShowOnIndex
}

// Check showing on detail.
func (p *Tree) IsShownOnDetail() bool {
	return p.ShowOnDetail
}

// Check for showing when creating.
func (p *Tree) IsShownOnCreation() bool {
	return p.ShowOnCreation
}

// Check for showing when exporting.
func (p *Tree) IsShownOnExport() bool {
	return p.ShowOnExport
}

// Check for showing when importing.
func (p *Tree) IsShownOnImport() bool {
	return p.ShowOnImport
}

// 设置为可编辑列
func (p *Tree) SetEditable(editable bool) *Tree {
	p.Editable = editable

	return p
}

// 闭包，透传表格列的属性
func (p *Tree) SetColumn(f func(column *table.Column) *table.Column) *Tree {
	p.Column = f(p.Column)

	return p
}

// 当前列值的枚举 valueEnum
func (p *Tree) GetValueEnum() map[interface{}]interface{} {
	data := map[interface{}]interface{}{}

	return data
}

// 设置回调函数
func (p *Tree) SetCallback(closure func() interface{}) *Tree {
	if closure != nil {
		p.Callback = closure
	}

	return p
}

// 获取回调函数
func (p *Tree) GetCallback() interface{} {
	return p.Callback
}

// 获取数据接口
func (p *Tree) SetApi(api string) *Tree {
	p.Api = api

	return p
}

// 是否自动展开父节点
func (p *Tree) SetAutoExpandParent(autoExpandParent bool) *Tree {
	p.AutoExpandParent = autoExpandParent

	return p
}

// 是否节点占据一行
func (p *Tree) SetBockNode(blockNode bool) *Tree {
	p.BockNode = blockNode

	return p
}

// 节点前添加 Checkbox 复选框
func (p *Tree) SetCheckable(checkable bool) *Tree {
	p.Checkable = checkable

	return p
}

// （受控）选中复选框的树节点（注意：父子节点有关联，如果传入父节点 key，则子节点自动选中；相应当子节点 key 都传入，父节点也自动选中。当设置 checkable 和 checkStrictly，它是一个有checked和halfChecked属性的对象，并且父子节点的选中与否不再关联
func (p *Tree) SetCheckedKeys(checkedKeys []interface{}) *Tree {
	p.CheckedKeys = checkedKeys

	return p
}

// checkable 状态下节点选择完全受控（父子节点选中状态不再关联）
func (p *Tree) SetCheckStrictly(checkStrictly bool) *Tree {
	p.CheckStrictly = checkStrictly

	return p
}

// 默认选中复选框的树节点
func (p *Tree) SetDefaultCheckedKeys(defaultCheckedKeys []interface{}) *Tree {
	p.DefaultCheckedKeys = defaultCheckedKeys

	return p
}

// 默认展开所有树节点
func (p *Tree) SetDefaultExpandAll(defaultExpandAll bool) *Tree {
	p.DefaultExpandAll = defaultExpandAll

	return p
}

// 默认展开指定的树节点
func (p *Tree) SetDefaultExpandedKeys(defaultExpandedKeys []interface{}) *Tree {
	p.DefaultExpandedKeys = defaultExpandedKeys

	return p
}

// 默认展开父节点
func (p *Tree) SetDefaultExpandParent(defaultExpandParent bool) *Tree {
	p.DefaultExpandParent = defaultExpandParent

	return p
}

// 默认选中的树节点
func (p *Tree) SetDefaultSelectedKeys(defaultSelectedKeys []interface{}) *Tree {
	p.DefaultSelectedKeys = defaultSelectedKeys

	return p
}

// 设置节点可拖拽，可以通过 icon: false 关闭拖拽提示图标
func (p *Tree) SetDraggable(draggable bool) *Tree {
	p.Draggable = draggable

	return p
}

// （受控）展开指定的树节点
func (p *Tree) SetExpandedKeys(expandedKeys []interface{}) *Tree {
	p.ExpandedKeys = expandedKeys

	return p
}

// 自定义 options 中 label value children 的字段
func (p *Tree) SetFieldNames(fieldNames *FieldNames) *Tree {
	p.FieldNames = fieldNames

	return p
}

// 设置虚拟滚动容器高度，设置后内部节点不再支持横向滚动
func (p *Tree) SetHeight(height int) *Tree {
	p.Height = height

	return p
}

// 自定义树节点图标
func (p *Tree) SetIcon(icon interface{}) *Tree {
	p.Icon = icon

	return p
}

// 支持点选多个节点（节点本身）
func (p *Tree) SetMultiple(multiple bool) *Tree {
	p.Multiple = multiple

	return p
}

// 添加在 Tree 最外层的 className
func (p *Tree) SetRootClassName(rootClassName string) *Tree {
	p.RootClassName = rootClassName

	return p
}

// 添加在 Tree 最外层的 style
func (p *Tree) SetRootStyle(rootStyle interface{}) *Tree {
	p.RootStyle = rootStyle

	return p
}

// 是否可选中
func (p *Tree) SetSelectable(selectable bool) *Tree {
	p.Selectable = selectable

	return p
}

// 设置选中的树节点
func (p *Tree) SetSelectedKeys(selectedKeys []interface{}) *Tree {
	p.SelectedKeys = selectedKeys

	return p
}

// 是否展示 TreeNode title 前的图标，没有默认样式，如设置为 true，需要自行定义图标相关样式
func (p *Tree) SetShowIcon(showIcon bool) *Tree {
	p.ShowIcon = showIcon

	return p
}

// 是否展示连接线
func (p *Tree) SetShowLine(showLine bool) *Tree {
	p.ShowLine = showLine

	return p
}

// 自定义树节点的展开/折叠图标
func (p *Tree) SetSwitcherIcon(switcherIcon interface{}) *Tree {
	p.SwitcherIcon = switcherIcon

	return p
}

// treeNodes 数据，如果设置则不需要手动构造 TreeNode 节点（value 在整个树范围内唯一）
func (p *Tree) SetTreeData(treeData *TreeData) *Tree {
	p.TreeData = treeData

	return p
}

// treeNodes 数据，如果设置则不需要手动构造 TreeNode 节点（value 在整个树范围内唯一）
func (p *Tree) SetData(treeData *TreeData) *Tree {
	p.TreeData = treeData

	return p
}

// 自定义样式
func (p *Tree) SetStyle(style map[string]interface{}) *Tree {
	p.Style = style

	return p
}