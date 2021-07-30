// Code generated by entc, DO NOT EDIT.

package lastupdated

import (
	"time"
	"urfu-abiturient-api/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
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
func IDGT(id int) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// LastUpdated applies equality check predicate on the "last_updated" field. It's identical to LastUpdatedEQ.
func LastUpdated(v time.Time) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastUpdated), v))
	})
}

// LastUpdatedEQ applies the EQ predicate on the "last_updated" field.
func LastUpdatedEQ(v time.Time) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastUpdated), v))
	})
}

// LastUpdatedNEQ applies the NEQ predicate on the "last_updated" field.
func LastUpdatedNEQ(v time.Time) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastUpdated), v))
	})
}

// LastUpdatedIn applies the In predicate on the "last_updated" field.
func LastUpdatedIn(vs ...time.Time) predicate.LastUpdated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LastUpdated(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastUpdated), v...))
	})
}

// LastUpdatedNotIn applies the NotIn predicate on the "last_updated" field.
func LastUpdatedNotIn(vs ...time.Time) predicate.LastUpdated {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LastUpdated(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastUpdated), v...))
	})
}

// LastUpdatedGT applies the GT predicate on the "last_updated" field.
func LastUpdatedGT(v time.Time) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastUpdated), v))
	})
}

// LastUpdatedGTE applies the GTE predicate on the "last_updated" field.
func LastUpdatedGTE(v time.Time) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastUpdated), v))
	})
}

// LastUpdatedLT applies the LT predicate on the "last_updated" field.
func LastUpdatedLT(v time.Time) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastUpdated), v))
	})
}

// LastUpdatedLTE applies the LTE predicate on the "last_updated" field.
func LastUpdatedLTE(v time.Time) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastUpdated), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LastUpdated) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LastUpdated) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
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
func Not(p predicate.LastUpdated) predicate.LastUpdated {
	return predicate.LastUpdated(func(s *sql.Selector) {
		p(s.Not())
	})
}