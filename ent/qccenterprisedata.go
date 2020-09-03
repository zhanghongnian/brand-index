// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/zhanghongnian/brand-index/ent/qccenterprisedata"
)

// QccEnterpriseData is the model entity for the QccEnterpriseData schema.
type QccEnterpriseData struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// RegisterStatus holds the value of the "register_status" field.
	RegisterStatus string `json:"register_status,omitempty"`
	// LegalPerson holds the value of the "legal_person" field.
	LegalPerson string `json:"legal_person,omitempty"`
	// RegisteredCapital holds the value of the "registered_capital" field.
	RegisteredCapital string `json:"registered_capital,omitempty"`
	// SetUpDate holds the value of the "set_up_date" field.
	SetUpDate string `json:"set_up_date,omitempty"`
	// VerifyDate holds the value of the "verify_date" field.
	VerifyDate string `json:"verify_date,omitempty"`
	// Province holds the value of the "province" field.
	Province string `json:"province,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// County holds the value of the "county" field.
	County string `json:"county,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// OtherPhones holds the value of the "other_phones" field.
	OtherPhones string `json:"other_phones,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// OtherEmails holds the value of the "other_emails" field.
	OtherEmails string `json:"other_emails,omitempty"`
	// UnifiedSocialCreditCode holds the value of the "unified_social_credit_code" field.
	UnifiedSocialCreditCode string `json:"unified_social_credit_code,omitempty"`
	// TaxpayerIdentificationNumber holds the value of the "taxpayer_identification_number" field.
	TaxpayerIdentificationNumber string `json:"taxpayer_identification_number,omitempty"`
	// RegistrationNumber holds the value of the "registration_number" field.
	RegistrationNumber string `json:"registration_number,omitempty"`
	// OrganizationCode holds the value of the "organization_code" field.
	OrganizationCode string `json:"organization_code,omitempty"`
	// InsurancePersonNums holds the value of the "insurance_person_nums" field.
	InsurancePersonNums string `json:"insurance_person_nums,omitempty"`
	// EnterpriseType holds the value of the "enterprise_type" field.
	EnterpriseType string `json:"enterprise_type,omitempty"`
	// IndustryInvolved holds the value of the "industry_involved" field.
	IndustryInvolved string `json:"industry_involved,omitempty"`
	// FormerName holds the value of the "former_name" field.
	FormerName string `json:"former_name,omitempty"`
	// Website holds the value of the "website" field.
	Website string `json:"website,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// BusinessScope holds the value of the "business_scope" field.
	BusinessScope string `json:"business_scope,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*QccEnterpriseData) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // register_status
		&sql.NullString{}, // legal_person
		&sql.NullString{}, // registered_capital
		&sql.NullString{}, // set_up_date
		&sql.NullString{}, // verify_date
		&sql.NullString{}, // province
		&sql.NullString{}, // city
		&sql.NullString{}, // county
		&sql.NullString{}, // phone
		&sql.NullString{}, // other_phones
		&sql.NullString{}, // email
		&sql.NullString{}, // other_emails
		&sql.NullString{}, // unified_social_credit_code
		&sql.NullString{}, // taxpayer_identification_number
		&sql.NullString{}, // registration_number
		&sql.NullString{}, // organization_code
		&sql.NullString{}, // insurance_person_nums
		&sql.NullString{}, // enterprise_type
		&sql.NullString{}, // industry_involved
		&sql.NullString{}, // former_name
		&sql.NullString{}, // website
		&sql.NullString{}, // address
		&sql.NullString{}, // business_scope
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the QccEnterpriseData fields.
func (qed *QccEnterpriseData) assignValues(values ...interface{}) error {
	if m, n := len(values), len(qccenterprisedata.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	qed.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		qed.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field register_status", values[1])
	} else if value.Valid {
		qed.RegisterStatus = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field legal_person", values[2])
	} else if value.Valid {
		qed.LegalPerson = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field registered_capital", values[3])
	} else if value.Valid {
		qed.RegisteredCapital = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field set_up_date", values[4])
	} else if value.Valid {
		qed.SetUpDate = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field verify_date", values[5])
	} else if value.Valid {
		qed.VerifyDate = value.String
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field province", values[6])
	} else if value.Valid {
		qed.Province = value.String
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field city", values[7])
	} else if value.Valid {
		qed.City = value.String
	}
	if value, ok := values[8].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field county", values[8])
	} else if value.Valid {
		qed.County = value.String
	}
	if value, ok := values[9].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field phone", values[9])
	} else if value.Valid {
		qed.Phone = value.String
	}
	if value, ok := values[10].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field other_phones", values[10])
	} else if value.Valid {
		qed.OtherPhones = value.String
	}
	if value, ok := values[11].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[11])
	} else if value.Valid {
		qed.Email = value.String
	}
	if value, ok := values[12].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field other_emails", values[12])
	} else if value.Valid {
		qed.OtherEmails = value.String
	}
	if value, ok := values[13].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field unified_social_credit_code", values[13])
	} else if value.Valid {
		qed.UnifiedSocialCreditCode = value.String
	}
	if value, ok := values[14].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field taxpayer_identification_number", values[14])
	} else if value.Valid {
		qed.TaxpayerIdentificationNumber = value.String
	}
	if value, ok := values[15].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field registration_number", values[15])
	} else if value.Valid {
		qed.RegistrationNumber = value.String
	}
	if value, ok := values[16].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field organization_code", values[16])
	} else if value.Valid {
		qed.OrganizationCode = value.String
	}
	if value, ok := values[17].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field insurance_person_nums", values[17])
	} else if value.Valid {
		qed.InsurancePersonNums = value.String
	}
	if value, ok := values[18].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field enterprise_type", values[18])
	} else if value.Valid {
		qed.EnterpriseType = value.String
	}
	if value, ok := values[19].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field industry_involved", values[19])
	} else if value.Valid {
		qed.IndustryInvolved = value.String
	}
	if value, ok := values[20].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field former_name", values[20])
	} else if value.Valid {
		qed.FormerName = value.String
	}
	if value, ok := values[21].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field website", values[21])
	} else if value.Valid {
		qed.Website = value.String
	}
	if value, ok := values[22].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field address", values[22])
	} else if value.Valid {
		qed.Address = value.String
	}
	if value, ok := values[23].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field business_scope", values[23])
	} else if value.Valid {
		qed.BusinessScope = value.String
	}
	return nil
}

// Update returns a builder for updating this QccEnterpriseData.
// Note that, you need to call QccEnterpriseData.Unwrap() before calling this method, if this QccEnterpriseData
// was returned from a transaction, and the transaction was committed or rolled back.
func (qed *QccEnterpriseData) Update() *QccEnterpriseDataUpdateOne {
	return (&QccEnterpriseDataClient{config: qed.config}).UpdateOne(qed)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (qed *QccEnterpriseData) Unwrap() *QccEnterpriseData {
	tx, ok := qed.config.driver.(*txDriver)
	if !ok {
		panic("ent: QccEnterpriseData is not a transactional entity")
	}
	qed.config.driver = tx.drv
	return qed
}

// String implements the fmt.Stringer.
func (qed *QccEnterpriseData) String() string {
	var builder strings.Builder
	builder.WriteString("QccEnterpriseData(")
	builder.WriteString(fmt.Sprintf("id=%v", qed.ID))
	builder.WriteString(", name=")
	builder.WriteString(qed.Name)
	builder.WriteString(", register_status=")
	builder.WriteString(qed.RegisterStatus)
	builder.WriteString(", legal_person=")
	builder.WriteString(qed.LegalPerson)
	builder.WriteString(", registered_capital=")
	builder.WriteString(qed.RegisteredCapital)
	builder.WriteString(", set_up_date=")
	builder.WriteString(qed.SetUpDate)
	builder.WriteString(", verify_date=")
	builder.WriteString(qed.VerifyDate)
	builder.WriteString(", province=")
	builder.WriteString(qed.Province)
	builder.WriteString(", city=")
	builder.WriteString(qed.City)
	builder.WriteString(", county=")
	builder.WriteString(qed.County)
	builder.WriteString(", phone=")
	builder.WriteString(qed.Phone)
	builder.WriteString(", other_phones=")
	builder.WriteString(qed.OtherPhones)
	builder.WriteString(", email=")
	builder.WriteString(qed.Email)
	builder.WriteString(", other_emails=")
	builder.WriteString(qed.OtherEmails)
	builder.WriteString(", unified_social_credit_code=")
	builder.WriteString(qed.UnifiedSocialCreditCode)
	builder.WriteString(", taxpayer_identification_number=")
	builder.WriteString(qed.TaxpayerIdentificationNumber)
	builder.WriteString(", registration_number=")
	builder.WriteString(qed.RegistrationNumber)
	builder.WriteString(", organization_code=")
	builder.WriteString(qed.OrganizationCode)
	builder.WriteString(", insurance_person_nums=")
	builder.WriteString(qed.InsurancePersonNums)
	builder.WriteString(", enterprise_type=")
	builder.WriteString(qed.EnterpriseType)
	builder.WriteString(", industry_involved=")
	builder.WriteString(qed.IndustryInvolved)
	builder.WriteString(", former_name=")
	builder.WriteString(qed.FormerName)
	builder.WriteString(", website=")
	builder.WriteString(qed.Website)
	builder.WriteString(", address=")
	builder.WriteString(qed.Address)
	builder.WriteString(", business_scope=")
	builder.WriteString(qed.BusinessScope)
	builder.WriteByte(')')
	return builder.String()
}

// QccEnterpriseDataSlice is a parsable slice of QccEnterpriseData.
type QccEnterpriseDataSlice []*QccEnterpriseData

func (qed QccEnterpriseDataSlice) config(cfg config) {
	for _i := range qed {
		qed[_i].config = cfg
	}
}
