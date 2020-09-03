// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/zhanghongnian/brand-index/ent/predicate"
	"github.com/zhanghongnian/brand-index/ent/qccenterprisedata"
)

// QccEnterpriseDataUpdate is the builder for updating QccEnterpriseData entities.
type QccEnterpriseDataUpdate struct {
	config
	hooks      []Hook
	mutation   *QccEnterpriseDataMutation
	predicates []predicate.QccEnterpriseData
}

// Where adds a new predicate for the builder.
func (qedu *QccEnterpriseDataUpdate) Where(ps ...predicate.QccEnterpriseData) *QccEnterpriseDataUpdate {
	qedu.predicates = append(qedu.predicates, ps...)
	return qedu
}

// SetName sets the name field.
func (qedu *QccEnterpriseDataUpdate) SetName(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetName(s)
	return qedu
}

// SetRegisterStatus sets the register_status field.
func (qedu *QccEnterpriseDataUpdate) SetRegisterStatus(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetRegisterStatus(s)
	return qedu
}

// SetLegalPerson sets the legal_person field.
func (qedu *QccEnterpriseDataUpdate) SetLegalPerson(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetLegalPerson(s)
	return qedu
}

// SetRegisteredCapital sets the registered_capital field.
func (qedu *QccEnterpriseDataUpdate) SetRegisteredCapital(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetRegisteredCapital(s)
	return qedu
}

// SetSetUpDate sets the set_up_date field.
func (qedu *QccEnterpriseDataUpdate) SetSetUpDate(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetSetUpDate(s)
	return qedu
}

// SetVerifyDate sets the verify_date field.
func (qedu *QccEnterpriseDataUpdate) SetVerifyDate(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetVerifyDate(s)
	return qedu
}

// SetProvince sets the province field.
func (qedu *QccEnterpriseDataUpdate) SetProvince(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetProvince(s)
	return qedu
}

// SetCity sets the city field.
func (qedu *QccEnterpriseDataUpdate) SetCity(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetCity(s)
	return qedu
}

// SetCounty sets the county field.
func (qedu *QccEnterpriseDataUpdate) SetCounty(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetCounty(s)
	return qedu
}

// SetPhone sets the phone field.
func (qedu *QccEnterpriseDataUpdate) SetPhone(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetPhone(s)
	return qedu
}

// SetOtherPhones sets the other_phones field.
func (qedu *QccEnterpriseDataUpdate) SetOtherPhones(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetOtherPhones(s)
	return qedu
}

// SetEmail sets the email field.
func (qedu *QccEnterpriseDataUpdate) SetEmail(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetEmail(s)
	return qedu
}

// SetOtherEmails sets the other_emails field.
func (qedu *QccEnterpriseDataUpdate) SetOtherEmails(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetOtherEmails(s)
	return qedu
}

// SetUnifiedSocialCreditCode sets the unified_social_credit_code field.
func (qedu *QccEnterpriseDataUpdate) SetUnifiedSocialCreditCode(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetUnifiedSocialCreditCode(s)
	return qedu
}

// SetTaxpayerIdentificationNumber sets the taxpayer_identification_number field.
func (qedu *QccEnterpriseDataUpdate) SetTaxpayerIdentificationNumber(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetTaxpayerIdentificationNumber(s)
	return qedu
}

// SetRegistrationNumber sets the registration_number field.
func (qedu *QccEnterpriseDataUpdate) SetRegistrationNumber(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetRegistrationNumber(s)
	return qedu
}

// SetOrganizationCode sets the organization_code field.
func (qedu *QccEnterpriseDataUpdate) SetOrganizationCode(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetOrganizationCode(s)
	return qedu
}

// SetInsurancePersonNums sets the insurance_person_nums field.
func (qedu *QccEnterpriseDataUpdate) SetInsurancePersonNums(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetInsurancePersonNums(s)
	return qedu
}

// SetEnterpriseType sets the enterprise_type field.
func (qedu *QccEnterpriseDataUpdate) SetEnterpriseType(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetEnterpriseType(s)
	return qedu
}

// SetIndustryInvolved sets the industry_involved field.
func (qedu *QccEnterpriseDataUpdate) SetIndustryInvolved(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetIndustryInvolved(s)
	return qedu
}

// SetFormerName sets the former_name field.
func (qedu *QccEnterpriseDataUpdate) SetFormerName(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetFormerName(s)
	return qedu
}

// SetWebsite sets the website field.
func (qedu *QccEnterpriseDataUpdate) SetWebsite(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetWebsite(s)
	return qedu
}

// SetAddress sets the address field.
func (qedu *QccEnterpriseDataUpdate) SetAddress(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetAddress(s)
	return qedu
}

// SetBusinessScope sets the business_scope field.
func (qedu *QccEnterpriseDataUpdate) SetBusinessScope(s string) *QccEnterpriseDataUpdate {
	qedu.mutation.SetBusinessScope(s)
	return qedu
}

// Mutation returns the QccEnterpriseDataMutation object of the builder.
func (qedu *QccEnterpriseDataUpdate) Mutation() *QccEnterpriseDataMutation {
	return qedu.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (qedu *QccEnterpriseDataUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(qedu.hooks) == 0 {
		affected, err = qedu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QccEnterpriseDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qedu.mutation = mutation
			affected, err = qedu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(qedu.hooks) - 1; i >= 0; i-- {
			mut = qedu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qedu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (qedu *QccEnterpriseDataUpdate) SaveX(ctx context.Context) int {
	affected, err := qedu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qedu *QccEnterpriseDataUpdate) Exec(ctx context.Context) error {
	_, err := qedu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qedu *QccEnterpriseDataUpdate) ExecX(ctx context.Context) {
	if err := qedu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (qedu *QccEnterpriseDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   qccenterprisedata.Table,
			Columns: qccenterprisedata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: qccenterprisedata.FieldID,
			},
		},
	}
	if ps := qedu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qedu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldName,
		})
	}
	if value, ok := qedu.mutation.RegisterStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldRegisterStatus,
		})
	}
	if value, ok := qedu.mutation.LegalPerson(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldLegalPerson,
		})
	}
	if value, ok := qedu.mutation.RegisteredCapital(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldRegisteredCapital,
		})
	}
	if value, ok := qedu.mutation.SetUpDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldSetUpDate,
		})
	}
	if value, ok := qedu.mutation.VerifyDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldVerifyDate,
		})
	}
	if value, ok := qedu.mutation.Province(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldProvince,
		})
	}
	if value, ok := qedu.mutation.City(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldCity,
		})
	}
	if value, ok := qedu.mutation.County(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldCounty,
		})
	}
	if value, ok := qedu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldPhone,
		})
	}
	if value, ok := qedu.mutation.OtherPhones(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldOtherPhones,
		})
	}
	if value, ok := qedu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldEmail,
		})
	}
	if value, ok := qedu.mutation.OtherEmails(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldOtherEmails,
		})
	}
	if value, ok := qedu.mutation.UnifiedSocialCreditCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldUnifiedSocialCreditCode,
		})
	}
	if value, ok := qedu.mutation.TaxpayerIdentificationNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldTaxpayerIdentificationNumber,
		})
	}
	if value, ok := qedu.mutation.RegistrationNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldRegistrationNumber,
		})
	}
	if value, ok := qedu.mutation.OrganizationCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldOrganizationCode,
		})
	}
	if value, ok := qedu.mutation.InsurancePersonNums(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldInsurancePersonNums,
		})
	}
	if value, ok := qedu.mutation.EnterpriseType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldEnterpriseType,
		})
	}
	if value, ok := qedu.mutation.IndustryInvolved(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldIndustryInvolved,
		})
	}
	if value, ok := qedu.mutation.FormerName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldFormerName,
		})
	}
	if value, ok := qedu.mutation.Website(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldWebsite,
		})
	}
	if value, ok := qedu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldAddress,
		})
	}
	if value, ok := qedu.mutation.BusinessScope(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldBusinessScope,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, qedu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{qccenterprisedata.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// QccEnterpriseDataUpdateOne is the builder for updating a single QccEnterpriseData entity.
type QccEnterpriseDataUpdateOne struct {
	config
	hooks    []Hook
	mutation *QccEnterpriseDataMutation
}

// SetName sets the name field.
func (qeduo *QccEnterpriseDataUpdateOne) SetName(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetName(s)
	return qeduo
}

// SetRegisterStatus sets the register_status field.
func (qeduo *QccEnterpriseDataUpdateOne) SetRegisterStatus(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetRegisterStatus(s)
	return qeduo
}

// SetLegalPerson sets the legal_person field.
func (qeduo *QccEnterpriseDataUpdateOne) SetLegalPerson(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetLegalPerson(s)
	return qeduo
}

// SetRegisteredCapital sets the registered_capital field.
func (qeduo *QccEnterpriseDataUpdateOne) SetRegisteredCapital(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetRegisteredCapital(s)
	return qeduo
}

// SetSetUpDate sets the set_up_date field.
func (qeduo *QccEnterpriseDataUpdateOne) SetSetUpDate(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetSetUpDate(s)
	return qeduo
}

// SetVerifyDate sets the verify_date field.
func (qeduo *QccEnterpriseDataUpdateOne) SetVerifyDate(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetVerifyDate(s)
	return qeduo
}

// SetProvince sets the province field.
func (qeduo *QccEnterpriseDataUpdateOne) SetProvince(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetProvince(s)
	return qeduo
}

// SetCity sets the city field.
func (qeduo *QccEnterpriseDataUpdateOne) SetCity(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetCity(s)
	return qeduo
}

// SetCounty sets the county field.
func (qeduo *QccEnterpriseDataUpdateOne) SetCounty(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetCounty(s)
	return qeduo
}

// SetPhone sets the phone field.
func (qeduo *QccEnterpriseDataUpdateOne) SetPhone(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetPhone(s)
	return qeduo
}

// SetOtherPhones sets the other_phones field.
func (qeduo *QccEnterpriseDataUpdateOne) SetOtherPhones(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetOtherPhones(s)
	return qeduo
}

// SetEmail sets the email field.
func (qeduo *QccEnterpriseDataUpdateOne) SetEmail(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetEmail(s)
	return qeduo
}

// SetOtherEmails sets the other_emails field.
func (qeduo *QccEnterpriseDataUpdateOne) SetOtherEmails(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetOtherEmails(s)
	return qeduo
}

// SetUnifiedSocialCreditCode sets the unified_social_credit_code field.
func (qeduo *QccEnterpriseDataUpdateOne) SetUnifiedSocialCreditCode(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetUnifiedSocialCreditCode(s)
	return qeduo
}

// SetTaxpayerIdentificationNumber sets the taxpayer_identification_number field.
func (qeduo *QccEnterpriseDataUpdateOne) SetTaxpayerIdentificationNumber(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetTaxpayerIdentificationNumber(s)
	return qeduo
}

// SetRegistrationNumber sets the registration_number field.
func (qeduo *QccEnterpriseDataUpdateOne) SetRegistrationNumber(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetRegistrationNumber(s)
	return qeduo
}

// SetOrganizationCode sets the organization_code field.
func (qeduo *QccEnterpriseDataUpdateOne) SetOrganizationCode(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetOrganizationCode(s)
	return qeduo
}

// SetInsurancePersonNums sets the insurance_person_nums field.
func (qeduo *QccEnterpriseDataUpdateOne) SetInsurancePersonNums(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetInsurancePersonNums(s)
	return qeduo
}

// SetEnterpriseType sets the enterprise_type field.
func (qeduo *QccEnterpriseDataUpdateOne) SetEnterpriseType(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetEnterpriseType(s)
	return qeduo
}

// SetIndustryInvolved sets the industry_involved field.
func (qeduo *QccEnterpriseDataUpdateOne) SetIndustryInvolved(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetIndustryInvolved(s)
	return qeduo
}

// SetFormerName sets the former_name field.
func (qeduo *QccEnterpriseDataUpdateOne) SetFormerName(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetFormerName(s)
	return qeduo
}

// SetWebsite sets the website field.
func (qeduo *QccEnterpriseDataUpdateOne) SetWebsite(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetWebsite(s)
	return qeduo
}

// SetAddress sets the address field.
func (qeduo *QccEnterpriseDataUpdateOne) SetAddress(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetAddress(s)
	return qeduo
}

// SetBusinessScope sets the business_scope field.
func (qeduo *QccEnterpriseDataUpdateOne) SetBusinessScope(s string) *QccEnterpriseDataUpdateOne {
	qeduo.mutation.SetBusinessScope(s)
	return qeduo
}

// Mutation returns the QccEnterpriseDataMutation object of the builder.
func (qeduo *QccEnterpriseDataUpdateOne) Mutation() *QccEnterpriseDataMutation {
	return qeduo.mutation
}

// Save executes the query and returns the updated entity.
func (qeduo *QccEnterpriseDataUpdateOne) Save(ctx context.Context) (*QccEnterpriseData, error) {
	var (
		err  error
		node *QccEnterpriseData
	)
	if len(qeduo.hooks) == 0 {
		node, err = qeduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QccEnterpriseDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qeduo.mutation = mutation
			node, err = qeduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(qeduo.hooks) - 1; i >= 0; i-- {
			mut = qeduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qeduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (qeduo *QccEnterpriseDataUpdateOne) SaveX(ctx context.Context) *QccEnterpriseData {
	qed, err := qeduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return qed
}

// Exec executes the query on the entity.
func (qeduo *QccEnterpriseDataUpdateOne) Exec(ctx context.Context) error {
	_, err := qeduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qeduo *QccEnterpriseDataUpdateOne) ExecX(ctx context.Context) {
	if err := qeduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (qeduo *QccEnterpriseDataUpdateOne) sqlSave(ctx context.Context) (qed *QccEnterpriseData, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   qccenterprisedata.Table,
			Columns: qccenterprisedata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: qccenterprisedata.FieldID,
			},
		},
	}
	id, ok := qeduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing QccEnterpriseData.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := qeduo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldName,
		})
	}
	if value, ok := qeduo.mutation.RegisterStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldRegisterStatus,
		})
	}
	if value, ok := qeduo.mutation.LegalPerson(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldLegalPerson,
		})
	}
	if value, ok := qeduo.mutation.RegisteredCapital(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldRegisteredCapital,
		})
	}
	if value, ok := qeduo.mutation.SetUpDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldSetUpDate,
		})
	}
	if value, ok := qeduo.mutation.VerifyDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldVerifyDate,
		})
	}
	if value, ok := qeduo.mutation.Province(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldProvince,
		})
	}
	if value, ok := qeduo.mutation.City(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldCity,
		})
	}
	if value, ok := qeduo.mutation.County(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldCounty,
		})
	}
	if value, ok := qeduo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldPhone,
		})
	}
	if value, ok := qeduo.mutation.OtherPhones(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldOtherPhones,
		})
	}
	if value, ok := qeduo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldEmail,
		})
	}
	if value, ok := qeduo.mutation.OtherEmails(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldOtherEmails,
		})
	}
	if value, ok := qeduo.mutation.UnifiedSocialCreditCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldUnifiedSocialCreditCode,
		})
	}
	if value, ok := qeduo.mutation.TaxpayerIdentificationNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldTaxpayerIdentificationNumber,
		})
	}
	if value, ok := qeduo.mutation.RegistrationNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldRegistrationNumber,
		})
	}
	if value, ok := qeduo.mutation.OrganizationCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldOrganizationCode,
		})
	}
	if value, ok := qeduo.mutation.InsurancePersonNums(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldInsurancePersonNums,
		})
	}
	if value, ok := qeduo.mutation.EnterpriseType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldEnterpriseType,
		})
	}
	if value, ok := qeduo.mutation.IndustryInvolved(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldIndustryInvolved,
		})
	}
	if value, ok := qeduo.mutation.FormerName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldFormerName,
		})
	}
	if value, ok := qeduo.mutation.Website(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldWebsite,
		})
	}
	if value, ok := qeduo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldAddress,
		})
	}
	if value, ok := qeduo.mutation.BusinessScope(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qccenterprisedata.FieldBusinessScope,
		})
	}
	qed = &QccEnterpriseData{config: qeduo.config}
	_spec.Assign = qed.assignValues
	_spec.ScanValues = qed.scanValues()
	if err = sqlgraph.UpdateNode(ctx, qeduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{qccenterprisedata.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return qed, nil
}
