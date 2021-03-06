// Code generated by entc, DO NOT EDIT.

package operativerecord

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/mmildd_s/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OperativeTime applies equality check predicate on the "OperativeTime" field. It's identical to OperativeTimeEQ.
func OperativeTime(v time.Time) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOperativeTime), v))
	})
}

// OperativeTimeEQ applies the EQ predicate on the "OperativeTime" field.
func OperativeTimeEQ(v time.Time) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOperativeTime), v))
	})
}

// OperativeTimeNEQ applies the NEQ predicate on the "OperativeTime" field.
func OperativeTimeNEQ(v time.Time) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOperativeTime), v))
	})
}

// OperativeTimeIn applies the In predicate on the "OperativeTime" field.
func OperativeTimeIn(vs ...time.Time) predicate.Operativerecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Operativerecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOperativeTime), v...))
	})
}

// OperativeTimeNotIn applies the NotIn predicate on the "OperativeTime" field.
func OperativeTimeNotIn(vs ...time.Time) predicate.Operativerecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Operativerecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOperativeTime), v...))
	})
}

// OperativeTimeGT applies the GT predicate on the "OperativeTime" field.
func OperativeTimeGT(v time.Time) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOperativeTime), v))
	})
}

// OperativeTimeGTE applies the GTE predicate on the "OperativeTime" field.
func OperativeTimeGTE(v time.Time) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOperativeTime), v))
	})
}

// OperativeTimeLT applies the LT predicate on the "OperativeTime" field.
func OperativeTimeLT(v time.Time) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOperativeTime), v))
	})
}

// OperativeTimeLTE applies the LTE predicate on the "OperativeTime" field.
func OperativeTimeLTE(v time.Time) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOperativeTime), v))
	})
}

// HasExaminationroom applies the HasEdge predicate on the "Examinationroom" edge.
func HasExaminationroom() predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ExaminationroomTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ExaminationroomTable, ExaminationroomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExaminationroomWith applies the HasEdge predicate on the "Examinationroom" edge with a given conditions (other predicates).
func HasExaminationroomWith(preds ...predicate.Examinationroom) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ExaminationroomInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ExaminationroomTable, ExaminationroomColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNurse applies the HasEdge predicate on the "Nurse" edge.
func HasNurse() predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NurseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NurseTable, NurseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNurseWith applies the HasEdge predicate on the "Nurse" edge with a given conditions (other predicates).
func HasNurseWith(preds ...predicate.Nurse) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NurseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NurseTable, NurseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOperative applies the HasEdge predicate on the "Operative" edge.
func HasOperative() predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OperativeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OperativeTable, OperativeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOperativeWith applies the HasEdge predicate on the "Operative" edge with a given conditions (other predicates).
func HasOperativeWith(preds ...predicate.Operative) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OperativeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OperativeTable, OperativeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTool applies the HasEdge predicate on the "Tool" edge.
func HasTool() predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ToolTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ToolTable, ToolColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasToolWith applies the HasEdge predicate on the "Tool" edge with a given conditions (other predicates).
func HasToolWith(preds ...predicate.Tool) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ToolInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ToolTable, ToolColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Operativerecord) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Operativerecord) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
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
func Not(p predicate.Operativerecord) predicate.Operativerecord {
	return predicate.Operativerecord(func(s *sql.Selector) {
		p(s.Not())
	})
}
