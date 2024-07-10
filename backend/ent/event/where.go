// Code generated by ent, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kazuki1023/Next_Go_CRUD/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldTitle, v))
}

// Generation applies equality check predicate on the "generation" field. It's identical to GenerationEQ.
func Generation(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldGeneration, v))
}

// StartDate applies equality check predicate on the "start_date" field. It's identical to StartDateEQ.
func StartDate(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldStartDate, v))
}

// FinishDate applies equality check predicate on the "finish_date" field. It's identical to FinishDateEQ.
func FinishDate(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldFinishDate, v))
}

// Detail applies equality check predicate on the "detail" field. It's identical to DetailEQ.
func Detail(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDetail, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldTitle, v))
}

// GenerationEQ applies the EQ predicate on the "generation" field.
func GenerationEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldGeneration, v))
}

// GenerationNEQ applies the NEQ predicate on the "generation" field.
func GenerationNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldGeneration, v))
}

// GenerationIn applies the In predicate on the "generation" field.
func GenerationIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldGeneration, vs...))
}

// GenerationNotIn applies the NotIn predicate on the "generation" field.
func GenerationNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldGeneration, vs...))
}

// GenerationGT applies the GT predicate on the "generation" field.
func GenerationGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldGeneration, v))
}

// GenerationGTE applies the GTE predicate on the "generation" field.
func GenerationGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldGeneration, v))
}

// GenerationLT applies the LT predicate on the "generation" field.
func GenerationLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldGeneration, v))
}

// GenerationLTE applies the LTE predicate on the "generation" field.
func GenerationLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldGeneration, v))
}

// GenerationContains applies the Contains predicate on the "generation" field.
func GenerationContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldGeneration, v))
}

// GenerationHasPrefix applies the HasPrefix predicate on the "generation" field.
func GenerationHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldGeneration, v))
}

// GenerationHasSuffix applies the HasSuffix predicate on the "generation" field.
func GenerationHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldGeneration, v))
}

// GenerationEqualFold applies the EqualFold predicate on the "generation" field.
func GenerationEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldGeneration, v))
}

// GenerationContainsFold applies the ContainsFold predicate on the "generation" field.
func GenerationContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldGeneration, v))
}

// StartDateEQ applies the EQ predicate on the "start_date" field.
func StartDateEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldStartDate, v))
}

// StartDateNEQ applies the NEQ predicate on the "start_date" field.
func StartDateNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldStartDate, v))
}

// StartDateIn applies the In predicate on the "start_date" field.
func StartDateIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldStartDate, vs...))
}

// StartDateNotIn applies the NotIn predicate on the "start_date" field.
func StartDateNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldStartDate, vs...))
}

// StartDateGT applies the GT predicate on the "start_date" field.
func StartDateGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldStartDate, v))
}

// StartDateGTE applies the GTE predicate on the "start_date" field.
func StartDateGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldStartDate, v))
}

// StartDateLT applies the LT predicate on the "start_date" field.
func StartDateLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldStartDate, v))
}

// StartDateLTE applies the LTE predicate on the "start_date" field.
func StartDateLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldStartDate, v))
}

// FinishDateEQ applies the EQ predicate on the "finish_date" field.
func FinishDateEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldFinishDate, v))
}

// FinishDateNEQ applies the NEQ predicate on the "finish_date" field.
func FinishDateNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldFinishDate, v))
}

// FinishDateIn applies the In predicate on the "finish_date" field.
func FinishDateIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldFinishDate, vs...))
}

// FinishDateNotIn applies the NotIn predicate on the "finish_date" field.
func FinishDateNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldFinishDate, vs...))
}

// FinishDateGT applies the GT predicate on the "finish_date" field.
func FinishDateGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldFinishDate, v))
}

// FinishDateGTE applies the GTE predicate on the "finish_date" field.
func FinishDateGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldFinishDate, v))
}

// FinishDateLT applies the LT predicate on the "finish_date" field.
func FinishDateLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldFinishDate, v))
}

// FinishDateLTE applies the LTE predicate on the "finish_date" field.
func FinishDateLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldFinishDate, v))
}

// DetailEQ applies the EQ predicate on the "detail" field.
func DetailEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDetail, v))
}

// DetailNEQ applies the NEQ predicate on the "detail" field.
func DetailNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldDetail, v))
}

// DetailIn applies the In predicate on the "detail" field.
func DetailIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldDetail, vs...))
}

// DetailNotIn applies the NotIn predicate on the "detail" field.
func DetailNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldDetail, vs...))
}

// DetailGT applies the GT predicate on the "detail" field.
func DetailGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldDetail, v))
}

// DetailGTE applies the GTE predicate on the "detail" field.
func DetailGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldDetail, v))
}

// DetailLT applies the LT predicate on the "detail" field.
func DetailLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldDetail, v))
}

// DetailLTE applies the LTE predicate on the "detail" field.
func DetailLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldDetail, v))
}

// DetailContains applies the Contains predicate on the "detail" field.
func DetailContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldDetail, v))
}

// DetailHasPrefix applies the HasPrefix predicate on the "detail" field.
func DetailHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldDetail, v))
}

// DetailHasSuffix applies the HasSuffix predicate on the "detail" field.
func DetailHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldDetail, v))
}

// DetailEqualFold applies the EqualFold predicate on the "detail" field.
func DetailEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldDetail, v))
}

// DetailContainsFold applies the ContainsFold predicate on the "detail" field.
func DetailContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldDetail, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(sql.NotPredicates(p))
}