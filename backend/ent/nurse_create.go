// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/mmildd_s/app/ent/nurse"
	"github.com/mmildd_s/app/ent/operativerecord"
)

// NurseCreate is the builder for creating a Nurse entity.
type NurseCreate struct {
	config
	mutation *NurseMutation
	hooks    []Hook
}

// SetNurseName sets the nurse_Name field.
func (nc *NurseCreate) SetNurseName(s string) *NurseCreate {
	nc.mutation.SetNurseName(s)
	return nc
}

// SetNurseEmail sets the nurse_Email field.
func (nc *NurseCreate) SetNurseEmail(s string) *NurseCreate {
	nc.mutation.SetNurseEmail(s)
	return nc
}

// SetNursePassword sets the nurse_Password field.
func (nc *NurseCreate) SetNursePassword(s string) *NurseCreate {
	nc.mutation.SetNursePassword(s)
	return nc
}

// SetNurseTel sets the nurse_Tel field.
func (nc *NurseCreate) SetNurseTel(s string) *NurseCreate {
	nc.mutation.SetNurseTel(s)
	return nc
}

// AddNurseOperativerecordIDs adds the Nurse_Operativerecord edge to Operativerecord by ids.
func (nc *NurseCreate) AddNurseOperativerecordIDs(ids ...int) *NurseCreate {
	nc.mutation.AddNurseOperativerecordIDs(ids...)
	return nc
}

// AddNurseOperativerecord adds the Nurse_Operativerecord edges to Operativerecord.
func (nc *NurseCreate) AddNurseOperativerecord(o ...*Operativerecord) *NurseCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return nc.AddNurseOperativerecordIDs(ids...)
}

// Mutation returns the NurseMutation object of the builder.
func (nc *NurseCreate) Mutation() *NurseMutation {
	return nc.mutation
}

// Save creates the Nurse in the database.
func (nc *NurseCreate) Save(ctx context.Context) (*Nurse, error) {
	if _, ok := nc.mutation.NurseName(); !ok {
		return nil, &ValidationError{Name: "nurse_Name", err: errors.New("ent: missing required field \"nurse_Name\"")}
	}
	if v, ok := nc.mutation.NurseName(); ok {
		if err := nurse.NurseNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "nurse_Name", err: fmt.Errorf("ent: validator failed for field \"nurse_Name\": %w", err)}
		}
	}
	if _, ok := nc.mutation.NurseEmail(); !ok {
		return nil, &ValidationError{Name: "nurse_Email", err: errors.New("ent: missing required field \"nurse_Email\"")}
	}
	if v, ok := nc.mutation.NurseEmail(); ok {
		if err := nurse.NurseEmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "nurse_Email", err: fmt.Errorf("ent: validator failed for field \"nurse_Email\": %w", err)}
		}
	}
	if _, ok := nc.mutation.NursePassword(); !ok {
		return nil, &ValidationError{Name: "nurse_Password", err: errors.New("ent: missing required field \"nurse_Password\"")}
	}
	if v, ok := nc.mutation.NursePassword(); ok {
		if err := nurse.NursePasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "nurse_Password", err: fmt.Errorf("ent: validator failed for field \"nurse_Password\": %w", err)}
		}
	}
	if _, ok := nc.mutation.NurseTel(); !ok {
		return nil, &ValidationError{Name: "nurse_Tel", err: errors.New("ent: missing required field \"nurse_Tel\"")}
	}
	if v, ok := nc.mutation.NurseTel(); ok {
		if err := nurse.NurseTelValidator(v); err != nil {
			return nil, &ValidationError{Name: "nurse_Tel", err: fmt.Errorf("ent: validator failed for field \"nurse_Tel\": %w", err)}
		}
	}
	var (
		err  error
		node *Nurse
	)
	if len(nc.hooks) == 0 {
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NurseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nc.mutation = mutation
			node, err = nc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			mut = nc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NurseCreate) SaveX(ctx context.Context) *Nurse {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nc *NurseCreate) sqlSave(ctx context.Context) (*Nurse, error) {
	n, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	n.ID = int(id)
	return n, nil
}

func (nc *NurseCreate) createSpec() (*Nurse, *sqlgraph.CreateSpec) {
	var (
		n     = &Nurse{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: nurse.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nurse.FieldID,
			},
		}
	)
	if value, ok := nc.mutation.NurseName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseName,
		})
		n.NurseName = value
	}
	if value, ok := nc.mutation.NurseEmail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseEmail,
		})
		n.NurseEmail = value
	}
	if value, ok := nc.mutation.NursePassword(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNursePassword,
		})
		n.NursePassword = value
	}
	if value, ok := nc.mutation.NurseTel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseTel,
		})
		n.NurseTel = value
	}
	if nodes := nc.mutation.NurseOperativerecordIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return n, _spec
}
