// Code generated by entc, DO NOT EDIT.

package operative

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/mmildd_s/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OperativeName applies equality check predicate on the "operative_Name" field. It's identical to OperativeNameEQ.
func OperativeName(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOperativeName), v))
	})
}

// OperativeNameEQ applies the EQ predicate on the "operative_Name" field.
func OperativeNameEQ(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOperativeName), v))
	})
}

// OperativeNameNEQ applies the NEQ predicate on the "operative_Name" field.
func OperativeNameNEQ(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOperativeName), v))
	})
}

// OperativeNameIn applies the In predicate on the "operative_Name" field.
func OperativeNameIn(vs ...string) predicate.Operative {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Operative(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOperativeName), v...))
	})
}

// OperativeNameNotIn applies the NotIn predicate on the "operative_Name" field.
func OperativeNameNotIn(vs ...string) predicate.Operative {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Operative(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOperativeName), v...))
	})
}

// OperativeNameGT applies the GT predicate on the "operative_Name" field.
func OperativeNameGT(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOperativeName), v))
	})
}

// OperativeNameGTE applies the GTE predicate on the "operative_Name" field.
func OperativeNameGTE(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOperativeName), v))
	})
}

// OperativeNameLT applies the LT predicate on the "operative_Name" field.
func OperativeNameLT(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOperativeName), v))
	})
}

// OperativeNameLTE applies the LTE predicate on the "operative_Name" field.
func OperativeNameLTE(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOperativeName), v))
	})
}

// OperativeNameContains applies the Contains predicate on the "operative_Name" field.
func OperativeNameContains(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOperativeName), v))
	})
}

// OperativeNameHasPrefix applies the HasPrefix predicate on the "operative_Name" field.
func OperativeNameHasPrefix(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOperativeName), v))
	})
}

// OperativeNameHasSuffix applies the HasSuffix predicate on the "operative_Name" field.
func OperativeNameHasSuffix(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOperativeName), v))
	})
}

// OperativeNameEqualFold applies the EqualFold predicate on the "operative_Name" field.
func OperativeNameEqualFold(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOperativeName), v))
	})
}

// OperativeNameContainsFold applies the ContainsFold predicate on the "operative_Name" field.
func OperativeNameContainsFold(v string) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOperativeName), v))
	})
}

// HasOperativeOperativerecord applies the HasEdge predicate on the "Operative_Operativerecord" edge.
func HasOperativeOperativerecord() predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OperativeOperativerecordTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OperativeOperativerecordTable, OperativeOperativerecordColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOperativeOperativerecordWith applies the HasEdge predicate on the "Operative_Operativerecord" edge with a given conditions (other predicates).
func HasOperativeOperativerecordWith(preds ...predicate.Operativerecord) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OperativeOperativerecordInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OperativeOperativerecordTable, OperativeOperativerecordColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Operative) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Operative) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Operative) predicate.Operative {
	return predicate.Operative(func(s *sql.Selector) {
		p(s.Not())
	})
}
