package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// QccEnterpriseData holds the schema definition for the QccEnterpriseData entity.
type QccEnterpriseData struct {
	ent.Schema
}

// Fields of the QccEnterpriseData.
func (QccEnterpriseData) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),                           // 企业名称
		field.String("register_status"),                // 登记状态
		field.String("legal_person"),                   // 法人
		field.String("registered_capital"),             // 注册资本
		field.String("set_up_date"),                    // 成立日期
		field.String("verify_date"),                    // 核准日期
		field.String("province"),                       // 省
		field.String("city"),                           // 市
		field.String("county"),                         // 县
		field.String("phone"),                          // 电话
		field.String("other_phones"),                   // 更多电话
		field.String("email"),                          // 邮箱
		field.String("other_emails"),                   // 更多邮箱
		field.String("unified_social_credit_code"),     // 统一社会信用代码
		field.String("taxpayer_identification_number"), // 纳税人识别号
		field.String("registration_number"),            // 注册号
		field.String("organization_code"),              // 组织机构代码
		field.String("insurance_person_nums"),          // 参保人数
		field.String("enterprise_type"),                // 企业类型
		field.String("industry_involved"),              // 所属行业
		field.String("former_name"),                    // 曾用名
		field.String("website"),                        // 官网
		field.String("address"),                        // 企业地址
		field.String("business_scope"),                 // 经营范围
	}
}

// Edges of the QccEnterpriseData.
func (QccEnterpriseData) Edges() []ent.Edge {
	return nil
}
