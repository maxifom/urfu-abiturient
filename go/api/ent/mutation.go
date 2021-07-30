// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"
	"urfu-abiturient-api/ent/abituriententry"
	"urfu-abiturient-api/ent/lastupdated"
	"urfu-abiturient-api/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAbiturientEntry = "AbiturientEntry"
	TypeLastUpdated     = "LastUpdated"
)

// AbiturientEntryMutation represents an operation that mutates the AbiturientEntry nodes in the graph.
type AbiturientEntryMutation struct {
	config
	op              Op
	typ             string
	id              *int
	name            *string
	number          *int64
	addnumber       *int64
	status          *string
	_type           *string
	statement_given *bool
	original_given  *bool
	speciality      *string
	program         *string
	form            *abituriententry.Form
	basis           *string
	sum             *int64
	addsum          *int64
	clearedFields   map[string]struct{}
	done            bool
	oldValue        func(context.Context) (*AbiturientEntry, error)
	predicates      []predicate.AbiturientEntry
}

var _ ent.Mutation = (*AbiturientEntryMutation)(nil)

// abituriententryOption allows management of the mutation configuration using functional options.
type abituriententryOption func(*AbiturientEntryMutation)

// newAbiturientEntryMutation creates new mutation for the AbiturientEntry entity.
func newAbiturientEntryMutation(c config, op Op, opts ...abituriententryOption) *AbiturientEntryMutation {
	m := &AbiturientEntryMutation{
		config:        c,
		op:            op,
		typ:           TypeAbiturientEntry,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAbiturientEntryID sets the ID field of the mutation.
func withAbiturientEntryID(id int) abituriententryOption {
	return func(m *AbiturientEntryMutation) {
		var (
			err   error
			once  sync.Once
			value *AbiturientEntry
		)
		m.oldValue = func(ctx context.Context) (*AbiturientEntry, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().AbiturientEntry.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAbiturientEntry sets the old AbiturientEntry of the mutation.
func withAbiturientEntry(node *AbiturientEntry) abituriententryOption {
	return func(m *AbiturientEntryMutation) {
		m.oldValue = func(context.Context) (*AbiturientEntry, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AbiturientEntryMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AbiturientEntryMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *AbiturientEntryMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *AbiturientEntryMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *AbiturientEntryMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the AbiturientEntry entity.
// If the AbiturientEntry object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AbiturientEntryMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *AbiturientEntryMutation) ResetName() {
	m.name = nil
}

// SetNumber sets the "number" field.
func (m *AbiturientEntryMutation) SetNumber(i int64) {
	m.number = &i
	m.addnumber = nil
}

// Number returns the value of the "number" field in the mutation.
func (m *AbiturientEntryMutation) Number() (r int64, exists bool) {
	v := m.number
	if v == nil {
		return
	}
	return *v, true
}

// OldNumber returns the old "number" field's value of the AbiturientEntry entity.
// If the AbiturientEntry object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AbiturientEntryMutation) OldNumber(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldNumber is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldNumber requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNumber: %w", err)
	}
	return oldValue.Number, nil
}

// AddNumber adds i to the "number" field.
func (m *AbiturientEntryMutation) AddNumber(i int64) {
	if m.addnumber != nil {
		*m.addnumber += i
	} else {
		m.addnumber = &i
	}
}

// AddedNumber returns the value that was added to the "number" field in this mutation.
func (m *AbiturientEntryMutation) AddedNumber() (r int64, exists bool) {
	v := m.addnumber
	if v == nil {
		return
	}
	return *v, true
}

// ResetNumber resets all changes to the "number" field.
func (m *AbiturientEntryMutation) ResetNumber() {
	m.number = nil
	m.addnumber = nil
}

// SetStatus sets the "status" field.
func (m *AbiturientEntryMutation) SetStatus(s string) {
	m.status = &s
}

// Status returns the value of the "status" field in the mutation.
func (m *AbiturientEntryMutation) Status() (r string, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the AbiturientEntry entity.
// If the AbiturientEntry object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AbiturientEntryMutation) OldStatus(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *AbiturientEntryMutation) ResetStatus() {
	m.status = nil
}

// SetType sets the "type" field.
func (m *AbiturientEntryMutation) SetType(s string) {
	m._type = &s
}

// GetType returns the value of the "type" field in the mutation.
func (m *AbiturientEntryMutation) GetType() (r string, exists bool) {
	v := m._type
	if v == nil {
		return
	}
	return *v, true
}

// OldType returns the old "type" field's value of the AbiturientEntry entity.
// If the AbiturientEntry object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AbiturientEntryMutation) OldType(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldType: %w", err)
	}
	return oldValue.Type, nil
}

// ResetType resets all changes to the "type" field.
func (m *AbiturientEntryMutation) ResetType() {
	m._type = nil
}

// SetStatementGiven sets the "statement_given" field.
func (m *AbiturientEntryMutation) SetStatementGiven(b bool) {
	m.statement_given = &b
}

// StatementGiven returns the value of the "statement_given" field in the mutation.
func (m *AbiturientEntryMutation) StatementGiven() (r bool, exists bool) {
	v := m.statement_given
	if v == nil {
		return
	}
	return *v, true
}

// OldStatementGiven returns the old "statement_given" field's value of the AbiturientEntry entity.
// If the AbiturientEntry object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AbiturientEntryMutation) OldStatementGiven(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatementGiven is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStatementGiven requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatementGiven: %w", err)
	}
	return oldValue.StatementGiven, nil
}

// ResetStatementGiven resets all changes to the "statement_given" field.
func (m *AbiturientEntryMutation) ResetStatementGiven() {
	m.statement_given = nil
}

// SetOriginalGiven sets the "original_given" field.
func (m *AbiturientEntryMutation) SetOriginalGiven(b bool) {
	m.original_given = &b
}

// OriginalGiven returns the value of the "original_given" field in the mutation.
func (m *AbiturientEntryMutation) OriginalGiven() (r bool, exists bool) {
	v := m.original_given
	if v == nil {
		return
	}
	return *v, true
}

// OldOriginalGiven returns the old "original_given" field's value of the AbiturientEntry entity.
// If the AbiturientEntry object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AbiturientEntryMutation) OldOriginalGiven(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldOriginalGiven is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldOriginalGiven requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOriginalGiven: %w", err)
	}
	return oldValue.OriginalGiven, nil
}

// ResetOriginalGiven resets all changes to the "original_given" field.
func (m *AbiturientEntryMutation) ResetOriginalGiven() {
	m.original_given = nil
}

// SetSpeciality sets the "speciality" field.
func (m *AbiturientEntryMutation) SetSpeciality(s string) {
	m.speciality = &s
}

// Speciality returns the value of the "speciality" field in the mutation.
func (m *AbiturientEntryMutation) Speciality() (r string, exists bool) {
	v := m.speciality
	if v == nil {
		return
	}
	return *v, true
}

// OldSpeciality returns the old "speciality" field's value of the AbiturientEntry entity.
// If the AbiturientEntry object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AbiturientEntryMutation) OldSpeciality(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldSpeciality is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldSpeciality requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSpeciality: %w", err)
	}
	return oldValue.Speciality, nil
}

// ResetSpeciality resets all changes to the "speciality" field.
func (m *AbiturientEntryMutation) ResetSpeciality() {
	m.speciality = nil
}

// SetProgram sets the "program" field.
func (m *AbiturientEntryMutation) SetProgram(s string) {
	m.program = &s
}

// Program returns the value of the "program" field in the mutation.
func (m *AbiturientEntryMutation) Program() (r string, exists bool) {
	v := m.program
	if v == nil {
		return
	}
	return *v, true
}

// OldProgram returns the old "program" field's value of the AbiturientEntry entity.
// If the AbiturientEntry object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AbiturientEntryMutation) OldProgram(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldProgram is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldProgram requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldProgram: %w", err)
	}
	return oldValue.Program, nil
}

// ResetProgram resets all changes to the "program" field.
func (m *AbiturientEntryMutation) ResetProgram() {
	m.program = nil
}

// SetForm sets the "form" field.
func (m *AbiturientEntryMutation) SetForm(a abituriententry.Form) {
	m.form = &a
}

// Form returns the value of the "form" field in the mutation.
func (m *AbiturientEntryMutation) Form() (r abituriententry.Form, exists bool) {
	v := m.form
	if v == nil {
		return
	}
	return *v, true
}

// OldForm returns the old "form" field's value of the AbiturientEntry entity.
// If the AbiturientEntry object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AbiturientEntryMutation) OldForm(ctx context.Context) (v abituriententry.Form, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldForm is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldForm requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldForm: %w", err)
	}
	return oldValue.Form, nil
}

// ResetForm resets all changes to the "form" field.
func (m *AbiturientEntryMutation) ResetForm() {
	m.form = nil
}

// SetBasis sets the "basis" field.
func (m *AbiturientEntryMutation) SetBasis(s string) {
	m.basis = &s
}

// Basis returns the value of the "basis" field in the mutation.
func (m *AbiturientEntryMutation) Basis() (r string, exists bool) {
	v := m.basis
	if v == nil {
		return
	}
	return *v, true
}

// OldBasis returns the old "basis" field's value of the AbiturientEntry entity.
// If the AbiturientEntry object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AbiturientEntryMutation) OldBasis(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldBasis is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldBasis requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBasis: %w", err)
	}
	return oldValue.Basis, nil
}

// ResetBasis resets all changes to the "basis" field.
func (m *AbiturientEntryMutation) ResetBasis() {
	m.basis = nil
}

// SetSum sets the "sum" field.
func (m *AbiturientEntryMutation) SetSum(i int64) {
	m.sum = &i
	m.addsum = nil
}

// Sum returns the value of the "sum" field in the mutation.
func (m *AbiturientEntryMutation) Sum() (r int64, exists bool) {
	v := m.sum
	if v == nil {
		return
	}
	return *v, true
}

// OldSum returns the old "sum" field's value of the AbiturientEntry entity.
// If the AbiturientEntry object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AbiturientEntryMutation) OldSum(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldSum is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldSum requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSum: %w", err)
	}
	return oldValue.Sum, nil
}

// AddSum adds i to the "sum" field.
func (m *AbiturientEntryMutation) AddSum(i int64) {
	if m.addsum != nil {
		*m.addsum += i
	} else {
		m.addsum = &i
	}
}

// AddedSum returns the value that was added to the "sum" field in this mutation.
func (m *AbiturientEntryMutation) AddedSum() (r int64, exists bool) {
	v := m.addsum
	if v == nil {
		return
	}
	return *v, true
}

// ResetSum resets all changes to the "sum" field.
func (m *AbiturientEntryMutation) ResetSum() {
	m.sum = nil
	m.addsum = nil
}

// Op returns the operation name.
func (m *AbiturientEntryMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (AbiturientEntry).
func (m *AbiturientEntryMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AbiturientEntryMutation) Fields() []string {
	fields := make([]string, 0, 11)
	if m.name != nil {
		fields = append(fields, abituriententry.FieldName)
	}
	if m.number != nil {
		fields = append(fields, abituriententry.FieldNumber)
	}
	if m.status != nil {
		fields = append(fields, abituriententry.FieldStatus)
	}
	if m._type != nil {
		fields = append(fields, abituriententry.FieldType)
	}
	if m.statement_given != nil {
		fields = append(fields, abituriententry.FieldStatementGiven)
	}
	if m.original_given != nil {
		fields = append(fields, abituriententry.FieldOriginalGiven)
	}
	if m.speciality != nil {
		fields = append(fields, abituriententry.FieldSpeciality)
	}
	if m.program != nil {
		fields = append(fields, abituriententry.FieldProgram)
	}
	if m.form != nil {
		fields = append(fields, abituriententry.FieldForm)
	}
	if m.basis != nil {
		fields = append(fields, abituriententry.FieldBasis)
	}
	if m.sum != nil {
		fields = append(fields, abituriententry.FieldSum)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AbiturientEntryMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case abituriententry.FieldName:
		return m.Name()
	case abituriententry.FieldNumber:
		return m.Number()
	case abituriententry.FieldStatus:
		return m.Status()
	case abituriententry.FieldType:
		return m.GetType()
	case abituriententry.FieldStatementGiven:
		return m.StatementGiven()
	case abituriententry.FieldOriginalGiven:
		return m.OriginalGiven()
	case abituriententry.FieldSpeciality:
		return m.Speciality()
	case abituriententry.FieldProgram:
		return m.Program()
	case abituriententry.FieldForm:
		return m.Form()
	case abituriententry.FieldBasis:
		return m.Basis()
	case abituriententry.FieldSum:
		return m.Sum()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AbiturientEntryMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case abituriententry.FieldName:
		return m.OldName(ctx)
	case abituriententry.FieldNumber:
		return m.OldNumber(ctx)
	case abituriententry.FieldStatus:
		return m.OldStatus(ctx)
	case abituriententry.FieldType:
		return m.OldType(ctx)
	case abituriententry.FieldStatementGiven:
		return m.OldStatementGiven(ctx)
	case abituriententry.FieldOriginalGiven:
		return m.OldOriginalGiven(ctx)
	case abituriententry.FieldSpeciality:
		return m.OldSpeciality(ctx)
	case abituriententry.FieldProgram:
		return m.OldProgram(ctx)
	case abituriententry.FieldForm:
		return m.OldForm(ctx)
	case abituriententry.FieldBasis:
		return m.OldBasis(ctx)
	case abituriententry.FieldSum:
		return m.OldSum(ctx)
	}
	return nil, fmt.Errorf("unknown AbiturientEntry field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AbiturientEntryMutation) SetField(name string, value ent.Value) error {
	switch name {
	case abituriententry.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case abituriententry.FieldNumber:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNumber(v)
		return nil
	case abituriententry.FieldStatus:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case abituriententry.FieldType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetType(v)
		return nil
	case abituriententry.FieldStatementGiven:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatementGiven(v)
		return nil
	case abituriententry.FieldOriginalGiven:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOriginalGiven(v)
		return nil
	case abituriententry.FieldSpeciality:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSpeciality(v)
		return nil
	case abituriententry.FieldProgram:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProgram(v)
		return nil
	case abituriententry.FieldForm:
		v, ok := value.(abituriententry.Form)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetForm(v)
		return nil
	case abituriententry.FieldBasis:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBasis(v)
		return nil
	case abituriententry.FieldSum:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSum(v)
		return nil
	}
	return fmt.Errorf("unknown AbiturientEntry field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AbiturientEntryMutation) AddedFields() []string {
	var fields []string
	if m.addnumber != nil {
		fields = append(fields, abituriententry.FieldNumber)
	}
	if m.addsum != nil {
		fields = append(fields, abituriententry.FieldSum)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AbiturientEntryMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case abituriententry.FieldNumber:
		return m.AddedNumber()
	case abituriententry.FieldSum:
		return m.AddedSum()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AbiturientEntryMutation) AddField(name string, value ent.Value) error {
	switch name {
	case abituriententry.FieldNumber:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddNumber(v)
		return nil
	case abituriententry.FieldSum:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddSum(v)
		return nil
	}
	return fmt.Errorf("unknown AbiturientEntry numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AbiturientEntryMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AbiturientEntryMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AbiturientEntryMutation) ClearField(name string) error {
	return fmt.Errorf("unknown AbiturientEntry nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AbiturientEntryMutation) ResetField(name string) error {
	switch name {
	case abituriententry.FieldName:
		m.ResetName()
		return nil
	case abituriententry.FieldNumber:
		m.ResetNumber()
		return nil
	case abituriententry.FieldStatus:
		m.ResetStatus()
		return nil
	case abituriententry.FieldType:
		m.ResetType()
		return nil
	case abituriententry.FieldStatementGiven:
		m.ResetStatementGiven()
		return nil
	case abituriententry.FieldOriginalGiven:
		m.ResetOriginalGiven()
		return nil
	case abituriententry.FieldSpeciality:
		m.ResetSpeciality()
		return nil
	case abituriententry.FieldProgram:
		m.ResetProgram()
		return nil
	case abituriententry.FieldForm:
		m.ResetForm()
		return nil
	case abituriententry.FieldBasis:
		m.ResetBasis()
		return nil
	case abituriententry.FieldSum:
		m.ResetSum()
		return nil
	}
	return fmt.Errorf("unknown AbiturientEntry field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AbiturientEntryMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AbiturientEntryMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AbiturientEntryMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AbiturientEntryMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AbiturientEntryMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AbiturientEntryMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AbiturientEntryMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown AbiturientEntry unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AbiturientEntryMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown AbiturientEntry edge %s", name)
}

// LastUpdatedMutation represents an operation that mutates the LastUpdated nodes in the graph.
type LastUpdatedMutation struct {
	config
	op            Op
	typ           string
	id            *int
	last_updated  *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*LastUpdated, error)
	predicates    []predicate.LastUpdated
}

var _ ent.Mutation = (*LastUpdatedMutation)(nil)

// lastupdatedOption allows management of the mutation configuration using functional options.
type lastupdatedOption func(*LastUpdatedMutation)

// newLastUpdatedMutation creates new mutation for the LastUpdated entity.
func newLastUpdatedMutation(c config, op Op, opts ...lastupdatedOption) *LastUpdatedMutation {
	m := &LastUpdatedMutation{
		config:        c,
		op:            op,
		typ:           TypeLastUpdated,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withLastUpdatedID sets the ID field of the mutation.
func withLastUpdatedID(id int) lastupdatedOption {
	return func(m *LastUpdatedMutation) {
		var (
			err   error
			once  sync.Once
			value *LastUpdated
		)
		m.oldValue = func(ctx context.Context) (*LastUpdated, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().LastUpdated.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withLastUpdated sets the old LastUpdated of the mutation.
func withLastUpdated(node *LastUpdated) lastupdatedOption {
	return func(m *LastUpdatedMutation) {
		m.oldValue = func(context.Context) (*LastUpdated, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m LastUpdatedMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m LastUpdatedMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *LastUpdatedMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetLastUpdated sets the "last_updated" field.
func (m *LastUpdatedMutation) SetLastUpdated(t time.Time) {
	m.last_updated = &t
}

// LastUpdated returns the value of the "last_updated" field in the mutation.
func (m *LastUpdatedMutation) LastUpdated() (r time.Time, exists bool) {
	v := m.last_updated
	if v == nil {
		return
	}
	return *v, true
}

// OldLastUpdated returns the old "last_updated" field's value of the LastUpdated entity.
// If the LastUpdated object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LastUpdatedMutation) OldLastUpdated(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLastUpdated is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldLastUpdated requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastUpdated: %w", err)
	}
	return oldValue.LastUpdated, nil
}

// ResetLastUpdated resets all changes to the "last_updated" field.
func (m *LastUpdatedMutation) ResetLastUpdated() {
	m.last_updated = nil
}

// Op returns the operation name.
func (m *LastUpdatedMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (LastUpdated).
func (m *LastUpdatedMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *LastUpdatedMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.last_updated != nil {
		fields = append(fields, lastupdated.FieldLastUpdated)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *LastUpdatedMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case lastupdated.FieldLastUpdated:
		return m.LastUpdated()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *LastUpdatedMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case lastupdated.FieldLastUpdated:
		return m.OldLastUpdated(ctx)
	}
	return nil, fmt.Errorf("unknown LastUpdated field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LastUpdatedMutation) SetField(name string, value ent.Value) error {
	switch name {
	case lastupdated.FieldLastUpdated:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastUpdated(v)
		return nil
	}
	return fmt.Errorf("unknown LastUpdated field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *LastUpdatedMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *LastUpdatedMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LastUpdatedMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown LastUpdated numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *LastUpdatedMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *LastUpdatedMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *LastUpdatedMutation) ClearField(name string) error {
	return fmt.Errorf("unknown LastUpdated nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *LastUpdatedMutation) ResetField(name string) error {
	switch name {
	case lastupdated.FieldLastUpdated:
		m.ResetLastUpdated()
		return nil
	}
	return fmt.Errorf("unknown LastUpdated field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *LastUpdatedMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *LastUpdatedMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *LastUpdatedMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *LastUpdatedMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *LastUpdatedMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *LastUpdatedMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *LastUpdatedMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown LastUpdated unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *LastUpdatedMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown LastUpdated edge %s", name)
}