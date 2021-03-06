// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/mmildd_s/app/ent/nurse"
	"github.com/mmildd_s/app/ent/operativerecord"
	"github.com/mmildd_s/app/ent/predicate"
)

// NurseUpdate is the builder for updating Nurse entities.
type NurseUpdate struct {
	config
	hooks      []Hook
	mutation   *NurseMutation
	predicates []predicate.Nurse
}

// Where adds a new predicate for the builder.
func (nu *NurseUpdate) Where(ps ...predicate.Nurse) *NurseUpdate {
	nu.predicates = append(nu.predicates, ps...)
	return nu
}

// SetNurseName sets the nurse_Name field.
func (nu *NurseUpdate) SetNurseName(s string) *NurseUpdate {
	nu.mutation.SetNurseName(s)
	return nu
}

// SetNurseEmail sets the nurse_Email field.
func (nu *NurseUpdate) SetNurseEmail(s string) *NurseUpdate {
	nu.mutation.SetNurseEmail(s)
	return nu
}

// SetNursePassword sets the nurse_Password field.
func (nu *NurseUpdate) SetNursePassword(s string) *NurseUpdate {
	nu.mutation.SetNursePassword(s)
	return nu
}

// SetNurseTel sets the nurse_Tel field.
func (nu *NurseUpdate) SetNurseTel(s string) *NurseUpdate {
	nu.mutation.SetNurseTel(s)
	return nu
}

// AddNurseOperativerecordIDs adds the Nurse_Operativerecord edge to Operativerecord by ids.
func (nu *NurseUpdate) AddNurseOperativerecordIDs(ids ...int) *NurseUpdate {
	nu.mutation.AddNurseOperativerecordIDs(ids...)
	return nu
}

// AddNurseOperativerecord adds the Nurse_Operativerecord edges to Operativerecord.
func (nu *NurseUpdate) AddNurseOperativerecord(o ...*Operativerecord) *NurseUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return nu.AddNurseOperativerecordIDs(ids...)
}

// Mutation returns the NurseMutation object of the builder.
func (nu *NurseUpdate) Mutation() *NurseMutation {
	return nu.mutation
}

// RemoveNurseOperativerecordIDs removes the Nurse_Operativerecord edge to Operativerecord by ids.
func (nu *NurseUpdate) RemoveNurseOperativerecordIDs(ids ...int) *NurseUpdate {
	nu.mutation.RemoveNurseOperativerecordIDs(ids...)
	return nu
}

// RemoveNurseOperativerecord removes Nurse_Operativerecord edges to Operativerecord.
func (nu *NurseUpdate) RemoveNurseOperativerecord(o ...*Operativerecord) *NurseUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return nu.RemoveNurseOperativerecordIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (nu *NurseUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := nu.mutation.NurseName(); ok {
		if err := nurse.NurseNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "nurse_Name", err: fmt.Errorf("ent: validator failed for field \"nurse_Name\": %w", err)}
		}
	}
	if v, ok := nu.mutation.NurseEmail(); ok {
		if err := nurse.NurseEmailValidator(v); err != nil {
			return 0, &ValidationError{Name: "nurse_Email", err: fmt.Errorf("ent: validator failed for field \"nurse_Email\": %w", err)}
		}
	}
	if v, ok := nu.mutation.NursePassword(); ok {
		if err := nurse.NursePasswordValidator(v); err != nil {
			return 0, &ValidationError{Name: "nurse_Password", err: fmt.Errorf("ent: validator failed for field \"nurse_Password\": %w", err)}
		}
	}
	if v, ok := nu.mutation.NurseTel(); ok {
		if err := nurse.NurseTelValidator(v); err != nil {
			return 0, &ValidationError{Name: "nurse_Tel", err: fmt.Errorf("ent: validator failed for field \"nurse_Tel\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(nu.hooks) == 0 {
		affected, err = nu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NurseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nu.mutation = mutation
			affected, err = nu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nu.hooks) - 1; i >= 0; i-- {
			mut = nu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NurseUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NurseUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NurseUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NurseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nurse.Table,
			Columns: nurse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nurse.FieldID,
			},
		},
	}
	if ps := nu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.NurseName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseName,
		})
	}
	if value, ok := nu.mutation.NurseEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseEmail,
		})
	}
	if value, ok := nu.mutation.NursePassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNursePassword,
		})
	}
	if value, ok := nu.mutation.NurseTel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseTel,
		})
	}
	if nodes := nu.mutation.RemovedNurseOperativerecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.NurseOperativerecordTable,
			Columns: []string{nurse.NurseOperativerecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: operativerecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.NurseOperativerecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.NurseOperativerecordTable,
			Columns: []string{nurse.NurseOperativerecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: operativerecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nurse.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// NurseUpdateOne is the builder for updating a single Nurse entity.
type NurseUpdateOne struct {
	config
	hooks    []Hook
	mutation *NurseMutation
}

// SetNurseName sets the nurse_Name field.
func (nuo *NurseUpdateOne) SetNurseName(s string) *NurseUpdateOne {
	nuo.mutation.SetNurseName(s)
	return nuo
}

// SetNurseEmail sets the nurse_Email field.
func (nuo *NurseUpdateOne) SetNurseEmail(s string) *NurseUpdateOne {
	nuo.mutation.SetNurseEmail(s)
	return nuo
}

// SetNursePassword sets the nurse_Password field.
func (nuo *NurseUpdateOne) SetNursePassword(s string) *NurseUpdateOne {
	nuo.mutation.SetNursePassword(s)
	return nuo
}

// SetNurseTel sets the nurse_Tel field.
func (nuo *NurseUpdateOne) SetNurseTel(s string) *NurseUpdateOne {
	nuo.mutation.SetNurseTel(s)
	return nuo
}

// AddNurseOperativerecordIDs adds the Nurse_Operativerecord edge to Operativerecord by ids.
func (nuo *NurseUpdateOne) AddNurseOperativerecordIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.AddNurseOperativerecordIDs(ids...)
	return nuo
}

// AddNurseOperativerecord adds the Nurse_Operativerecord edges to Operativerecord.
func (nuo *NurseUpdateOne) AddNurseOperativerecord(o ...*Operativerecord) *NurseUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return nuo.AddNurseOperativerecordIDs(ids...)
}

// Mutation returns the NurseMutation object of the builder.
func (nuo *NurseUpdateOne) Mutation() *NurseMutation {
	return nuo.mutation
}

// RemoveNurseOperativerecordIDs removes the Nurse_Operativerecord edge to Operativerecord by ids.
func (nuo *NurseUpdateOne) RemoveNurseOperativerecordIDs(ids ...int) *NurseUpdateOne {
	nuo.mutation.RemoveNurseOperativerecordIDs(ids...)
	return nuo
}

// RemoveNurseOperativerecord removes Nurse_Operativerecord edges to Operativerecord.
func (nuo *NurseUpdateOne) RemoveNurseOperativerecord(o ...*Operativerecord) *NurseUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return nuo.RemoveNurseOperativerecordIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (nuo *NurseUpdateOne) Save(ctx context.Context) (*Nurse, error) {
	if v, ok := nuo.mutation.NurseName(); ok {
		if err := nurse.NurseNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "nurse_Name", err: fmt.Errorf("ent: validator failed for field \"nurse_Name\": %w", err)}
		}
	}
	if v, ok := nuo.mutation.NurseEmail(); ok {
		if err := nurse.NurseEmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "nurse_Email", err: fmt.Errorf("ent: validator failed for field \"nurse_Email\": %w", err)}
		}
	}
	if v, ok := nuo.mutation.NursePassword(); ok {
		if err := nurse.NursePasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "nurse_Password", err: fmt.Errorf("ent: validator failed for field \"nurse_Password\": %w", err)}
		}
	}
	if v, ok := nuo.mutation.NurseTel(); ok {
		if err := nurse.NurseTelValidator(v); err != nil {
			return nil, &ValidationError{Name: "nurse_Tel", err: fmt.Errorf("ent: validator failed for field \"nurse_Tel\": %w", err)}
		}
	}

	var (
		err  error
		node *Nurse
	)
	if len(nuo.hooks) == 0 {
		node, err = nuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NurseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nuo.mutation = mutation
			node, err = nuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nuo.hooks) - 1; i >= 0; i-- {
			mut = nuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NurseUpdateOne) SaveX(ctx context.Context) *Nurse {
	n, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

// Exec executes the query on the entity.
func (nuo *NurseUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NurseUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NurseUpdateOne) sqlSave(ctx context.Context) (n *Nurse, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nurse.Table,
			Columns: nurse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nurse.FieldID,
			},
		},
	}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Nurse.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := nuo.mutation.NurseName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseName,
		})
	}
	if value, ok := nuo.mutation.NurseEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseEmail,
		})
	}
	if value, ok := nuo.mutation.NursePassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNursePassword,
		})
	}
	if value, ok := nuo.mutation.NurseTel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseTel,
		})
	}
	if nodes := nuo.mutation.RemovedNurseOperativerecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.NurseOperativerecordTable,
			Columns: []string{nurse.NurseOperativerecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: operativerecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.NurseOperativerecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.NurseOperativerecordTable,
			Columns: []string{nurse.NurseOperativerecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: operativerecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	n = &Nurse{config: nuo.config}
	_spec.Assign = n.assignValues
	_spec.ScanValues = n.scanValues()
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nurse.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return n, nil
}
