// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/zhanghongnian/brand-index/ent/qccenterprisedata"
)

// QccEnterpriseDataCreate is the builder for creating a QccEnterpriseData entity.
type QccEnterpriseDataCreate struct {
	config
	mutation *QccEnterpriseDataMutation
	hooks    []Hook
}

// SetName sets the name field.
func (qedc *QccEnterpriseDataCreate) SetName(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetName(s)
	return qedc
}

// SetRegisterStatus sets the register_status field.
func (qedc *QccEnterpriseDataCreate) SetRegisterStatus(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetRegisterStatus(s)
	return qedc
}

// SetLegalPerson sets the legal_person field.
func (qedc *QccEnterpriseDataCreate) SetLegalPerson(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetLegalPerson(s)
	return qedc
}

// SetRegisteredCapital sets the registered_capital field.
func (qedc *QccEnterpriseDataCreate) SetRegisteredCapital(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetRegisteredCapital(s)
	return qedc
}

// SetSetUpDate sets the set_up_date field.
func (qedc *QccEnterpriseDataCreate) SetSetUpDate(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetSetUpDate(s)
	return qedc
}

// SetVerifyDate sets the verify_date field.
func (qedc *QccEnterpriseDataCreate) SetVerifyDate(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetVerifyDate(s)
	return qedc
}

// SetProvince sets the province field.
func (qedc *QccEnterpriseDataCreate) SetProvince(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetProvince(s)
	return qedc
}

// SetCity sets the city field.
func (qedc *QccEnterpriseDataCreate) SetCity(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetCity(s)
	return qedc
}

// SetCounty sets the county field.
func (qedc *QccEnterpriseDataCreate) SetCounty(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetCounty(s)
	return qedc
}

// SetPhone sets the phone field.
func (qedc *QccEnterpriseDataCreate) SetPhone(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetPhone(s)
	return qedc
}

// SetOtherPhones sets the other_phones field.
func (qedc *QccEnterpriseDataCreate) SetOtherPhones(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetOtherPhones(s)
	return qedc
}

// SetEmail sets the email field.
func (qedc *QccEnterpriseDataCreate) SetEmail(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetEmail(s)
	return qedc
}

// SetOtherEmails sets the other_emails field.
func (qedc *QccEnterpriseDataCreate) SetOtherEmails(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetOtherEmails(s)
	return qedc
}

// SetUnifiedSocialCreditCode sets the unified_social_credit_code field.
func (qedc *QccEnterpriseDataCreate) SetUnifiedSocialCreditCode(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetUnifiedSocialCreditCode(s)
	return qedc
}

// SetTaxpayerIdentificationNumber sets the taxpayer_identification_number field.
func (qedc *QccEnterpriseDataCreate) SetTaxpayerIdentificationNumber(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetTaxpayerIdentificationNumber(s)
	return qedc
}

// SetRegistrationNumber sets the registration_number field.
func (qedc *QccEnterpriseDataCreate) SetRegistrationNumber(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetRegistrationNumber(s)
	return qedc
}

// SetOrganizationCode sets the organization_code field.
func (qedc *QccEnterpriseDataCreate) SetOrganizationCode(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetOrganizationCode(s)
	return qedc
}

// SetInsurancePersonNums sets the insurance_person_nums field.
func (qedc *QccEnterpriseDataCreate) SetInsurancePersonNums(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetInsurancePersonNums(s)
	return qedc
}

// SetEnterpriseType sets the enterprise_type field.
func (qedc *QccEnterpriseDataCreate) SetEnterpriseType(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetEnterpriseType(s)
	return qedc
}

// SetIndustryInvolved sets the industry_involved field.
func (qedc *QccEnterpriseDataCreate) SetIndustryInvolved(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetIndustryInvolved(s)
	return qedc
}

// SetFormerName sets the former_name field.
func (qedc *QccEnterpriseDataCreate) SetFormerName(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetFormerName(s)
	return qedc
}

// SetWebsite sets the website field.
func (qedc *QccEnterpriseDataCreate) SetWebsite(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetWebsite(s)
	return qedc
}

// SetAddress sets the address field.
func (qedc *QccEnterpriseDataCreate) SetAddress(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetAddress(s)
	return qedc
}

// SetBusinessScope sets the business_scope field.
func (qedc *QccEnterpriseDataCreate) SetBusinessScope(s string) *QccEnterpriseDataCreate {
	qedc.mutation.SetBusinessScope(s)
	return qedc
}

// Mutation returns the QccEnterpriseDataMutation object of the builder.
func (qedc *QccEnterpriseDataCreate) Mutation() *QccEnterpriseDataMutation {
	return qedc.mutation
}

// Save creates the QccEnterpriseData in the database.
func (qedc *QccEnterpriseDataCreate) Save(ctx context.Context) (*QccEnterpriseData, error) {
	if err := qedc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *QccEnterpriseData
	)
	if len(qedc.hooks) == 0 {
		node, err = qedc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QccEnterpriseDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qedc.mutation = mutation
			node, err = qedc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(qedc.hooks) - 1; i >= 0; i-- {
			mut = qedc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qedc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (qedc *QccEnterpriseDataCreate) SaveX(ctx context.Context) *QccEnterpriseData {
	v, err := qedc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (qedc *QccEnterpriseDataCreate) preSave() error {
	if _, ok := qedc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := qedc.mutation.RegisterStatus(); !ok {
		return &ValidationError{Name: "register_status", err: errors.New("ent: missing required field \"register_status\"")}
	}
	if _, ok := qedc.mutation.LegalPerson(); !ok {
		return &ValidationError{Name: "legal_person", err: errors.New("ent: missing required field \"legal_person\"")}
	}
	if _, ok := qedc.mutation.RegisteredCapital(); !ok {
		return &ValidationError{Name: "registered_capital", err: errors.New("ent: missing required field \"registered_capital\"")}
	}
	if _, ok := qedc.mutation.SetUpDate(); !ok {
		return &ValidationError{Name: "set_up_date", err: errors.New("ent: missing required field \"set_up_date\"")}
	}
	if _, ok := qedc.mutation.VerifyDate(); !ok {
		return &ValidationError{Name: "verify_date", err: errors.New("ent: missing required field \"verify_date\"")}
	}
	if _, ok := qedc.mutation.Province(); !ok {
		return &ValidationError{Name: "province", err: errors.New("ent: missing required field \"province\"")}
	}
	if _, ok := qedc.mutation.City(); !ok {
		return &ValidationError{Name: "city", err: errors.New("ent: missing required field \"city\"")}
	}
	if _, ok := qedc.mutation.County(); !ok {
		return &ValidationError{Name: "county", err: errors.New("ent: missing required field \"county\"")}
	}
	if _, ok := qedc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New("ent: missing required field \"phone\"")}
	}
	if _, ok := qedc.mutation.OtherPhones(); !ok {
		return &ValidationError{Name: "other_phones", err: errors.New("ent: missing required field \"other_phones\"")}
	}
	if _, ok := qedc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New("ent: missing required field \"email\"")}
	}
	if _, ok := qedc.mutation.OtherEmails(); !ok {
		return &ValidationError{Name: "other_emails", err: errors.New("ent: missing required field \"other_emails\"")}
	}
	if _, ok := qedc.mutation.UnifiedSocialCreditCode(); !ok {
		return &ValidationError{Name: "unified_social_credit_code", err: errors.New("ent: missing required field \"unified_social_credit_code\"")}
	}
	if _, ok := qedc.mutation.TaxpayerIdentificationNumber(); !ok {
		return &ValidationError{Name: "taxpayer_identification_number", err: errors.New("ent: missing required field \"taxpayer_identification_number\"")}
	}
	if _, ok := qedc.mutation.RegistrationNumber(); !ok {
		return &ValidationError{Name: "registration_number", err: errors.New("ent: missing required field \"registration_number\"")}
	}
	if _, ok := qedc.mutation.OrganizationCode(); !ok {
		return &ValidationError{Name: "organization_code", err: errors.New("ent: missing required field \"organization_code\"")}
	}
	if _, ok := qedc.mutation.InsurancePersonNums(); !ok {
		return &ValidationError{Name: "insurance_person_nums", err: errors.New("ent: missing required field \"insurance_person_nums\"")}
	}
	if _, ok := qedc.mutation.EnterpriseType(); !ok {
		return &ValidationError{Name: "enterprise_type", err: errors.New("ent: missing required field \"enterprise_type\"")}
	}
	if _, ok := qedc.mutation.IndustryInvolved(); !ok {
		return &ValidationError{Name: "industry_involved", err: errors.New("ent: missing required field \"industry_involved\"")}
	}
	if _, ok := qedc.mutation.FormerName(); !ok {
		return &ValidationError{Name: "former_name", err: errors.New("ent: missing required field \"former_name\"")}
	}
	if _, ok := qedc.mutation.Website(); !ok {
		return &ValidationError{Name: "website", err: errors.New("ent: missing required field \"website\"")}
	}
	if _, ok := qedc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New("ent: missing required field \"address\"")}
	}
	if _, ok := qedc.mutation.BusinessScope(); !ok {
		return &ValidationError{Name: "business_scope", err: errors.New("ent: missing required field \"business_scope\"")}
	}
	return nil
}

func (qedc *QccEnterpriseDataCreate) sqlSave(ctx context.Context) (*QccEnterpriseData, error) {
	qed, _spec := qedc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qedc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	qed.ID = int(id)
	return qed, nil
}

func (qedc *QccEnterpriseDataCreate) createSpec() (*QccEnterpriseData, *sqlgraph.CreateSpec) {
	var (
		qed   = &QccEnterpriseData{config: qedc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: qccenterprisedata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: qccenterprisedata.FieldID,
			},
		}
	)
	if value, ok := qedc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldName,
		})
		qed.Name = value
	}
	if value, ok := qedc.mutation.RegisterStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldRegisterStatus,
		})
		qed.RegisterStatus = value
	}
	if value, ok := qedc.mutation.LegalPerson(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldLegalPerson,
		})
		qed.LegalPerson = value
	}
	if value, ok := qedc.mutation.RegisteredCapital(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldRegisteredCapital,
		})
		qed.RegisteredCapital = value
	}
	if value, ok := qedc.mutation.SetUpDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldSetUpDate,
		})
		qed.SetUpDate = value
	}
	if value, ok := qedc.mutation.VerifyDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldVerifyDate,
		})
		qed.VerifyDate = value
	}
	if value, ok := qedc.mutation.Province(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldProvince,
		})
		qed.Province = value
	}
	if value, ok := qedc.mutation.City(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldCity,
		})
		qed.City = value
	}
	if value, ok := qedc.mutation.County(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldCounty,
		})
		qed.County = value
	}
	if value, ok := qedc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldPhone,
		})
		qed.Phone = value
	}
	if value, ok := qedc.mutation.OtherPhones(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldOtherPhones,
		})
		qed.OtherPhones = value
	}
	if value, ok := qedc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldEmail,
		})
		qed.Email = value
	}
	if value, ok := qedc.mutation.OtherEmails(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldOtherEmails,
		})
		qed.OtherEmails = value
	}
	if value, ok := qedc.mutation.UnifiedSocialCreditCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldUnifiedSocialCreditCode,
		})
		qed.UnifiedSocialCreditCode = value
	}
	if value, ok := qedc.mutation.TaxpayerIdentificationNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldTaxpayerIdentificationNumber,
		})
		qed.TaxpayerIdentificationNumber = value
	}
	if value, ok := qedc.mutation.RegistrationNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldRegistrationNumber,
		})
		qed.RegistrationNumber = value
	}
	if value, ok := qedc.mutation.OrganizationCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldOrganizationCode,
		})
		qed.OrganizationCode = value
	}
	if value, ok := qedc.mutation.InsurancePersonNums(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldInsurancePersonNums,
		})
		qed.InsurancePersonNums = value
	}
	if value, ok := qedc.mutation.EnterpriseType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldEnterpriseType,
		})
		qed.EnterpriseType = value
	}
	if value, ok := qedc.mutation.IndustryInvolved(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldIndustryInvolved,
		})
		qed.IndustryInvolved = value
	}
	if value, ok := qedc.mutation.FormerName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldFormerName,
		})
		qed.FormerName = value
	}
	if value, ok := qedc.mutation.Website(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldWebsite,
		})
		qed.Website = value
	}
	if value, ok := qedc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldAddress,
		})
		qed.Address = value
	}
	if value, ok := qedc.mutation.BusinessScope(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldBusinessScope,
		})
		qed.BusinessScope = value
	}
	return qed, _spec
}

// QccEnterpriseDataCreateBulk is the builder for creating a bulk of QccEnterpriseData entities.
type QccEnterpriseDataCreateBulk struct {
	config
	builders []*QccEnterpriseDataCreate
}

// Save creates the QccEnterpriseData entities in the database.
func (qedcb *QccEnterpriseDataCreateBulk) Save(ctx context.Context) ([]*QccEnterpriseData, error) {
	specs := make([]*sqlgraph.CreateSpec, len(qedcb.builders))
	nodes := make([]*QccEnterpriseData, len(qedcb.builders))
	mutators := make([]Mutator, len(qedcb.builders))
	for i := range qedcb.builders {
		func(i int, root context.Context) {
			builder := qedcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*QccEnterpriseDataMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, qedcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qedcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, qedcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (qedcb *QccEnterpriseDataCreateBulk) SaveX(ctx context.Context) []*QccEnterpriseData {
	v, err := qedcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
