// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/mmildd_s/app/ent/examinationroom"
)

// Examinationroom is the model entity for the Examinationroom schema.
type Examinationroom struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ExaminationroomName holds the value of the "examinationroom_name" field.
	ExaminationroomName string `json:"examinationroom_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExaminationroomQuery when eager-loading is set.
	Edges ExaminationroomEdges `json:"edges"`
}

// ExaminationroomEdges holds the relations/edges for other nodes in the graph.
type ExaminationroomEdges struct {
	// ExaminationroomOperativerecord holds the value of the Examinationroom_Operativerecord edge.
	ExaminationroomOperativerecord []*Operativerecord
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ExaminationroomOperativerecordOrErr returns the ExaminationroomOperativerecord value or an error if the edge
// was not loaded in eager-loading.
func (e ExaminationroomEdges) ExaminationroomOperativerecordOrErr() ([]*Operativerecord, error) {
	if e.loadedTypes[0] {
		return e.ExaminationroomOperativerecord, nil
	}
	return nil, &NotLoadedError{edge: "Examinationroom_Operativerecord"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Examinationroom) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // examinationroom_name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Examinationroom fields.
func (e *Examinationroom) assignValues(values ...interface{}) error {
	if m, n := len(values), len(examinationroom.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	e.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field examinationroom_name", values[0])
	} else if value.Valid {
		e.ExaminationroomName = value.String
	}
	return nil
}

// QueryExaminationroomOperativerecord queries the Examinationroom_Operativerecord edge of the Examinationroom.
func (e *Examinationroom) QueryExaminationroomOperativerecord() *OperativerecordQuery {
	return (&ExaminationroomClient{config: e.config}).QueryExaminationroomOperativerecord(e)
}

// Update returns a builder for updating this Examinationroom.
// Note that, you need to call Examinationroom.Unwrap() before calling this method, if this Examinationroom
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Examinationroom) Update() *ExaminationroomUpdateOne {
	return (&ExaminationroomClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (e *Examinationroom) Unwrap() *Examinationroom {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Examinationroom is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Examinationroom) String() string {
	var builder strings.Builder
	builder.WriteString("Examinationroom(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", examinationroom_name=")
	builder.WriteString(e.ExaminationroomName)
	builder.WriteByte(')')
	return builder.String()
}

// Examinationrooms is a parsable slice of Examinationroom.
type Examinationrooms []*Examinationroom

func (e Examinationrooms) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
