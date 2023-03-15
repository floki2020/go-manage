// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-manage/pkg/ent/company"
	"go-manage/pkg/ent/personal"
	"go-manage/pkg/ent/predicate"
	"go-manage/pkg/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUserName sets the "user_name" field.
func (uu *UserUpdate) SetUserName(s []string) *UserUpdate {
	uu.mutation.SetUserName(s)
	return uu
}

// AppendUserName appends s to the "user_name" field.
func (uu *UserUpdate) AppendUserName(s []string) *UserUpdate {
	uu.mutation.AppendUserName(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s []string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// AppendPassword appends s to the "password" field.
func (uu *UserUpdate) AppendPassword(s []string) *UserUpdate {
	uu.mutation.AppendPassword(s)
	return uu
}

// SetParentID sets the "parent_id" field.
func (uu *UserUpdate) SetParentID(i int) *UserUpdate {
	uu.mutation.SetParentID(i)
	return uu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableParentID(i *int) *UserUpdate {
	if i != nil {
		uu.SetParentID(*i)
	}
	return uu
}

// ClearParentID clears the value of the "parent_id" field.
func (uu *UserUpdate) ClearParentID() *UserUpdate {
	uu.mutation.ClearParentID()
	return uu
}

// AddPersonalIDs adds the "personals" edge to the Personal entity by IDs.
func (uu *UserUpdate) AddPersonalIDs(ids ...int) *UserUpdate {
	uu.mutation.AddPersonalIDs(ids...)
	return uu
}

// AddPersonals adds the "personals" edges to the Personal entity.
func (uu *UserUpdate) AddPersonals(p ...*Personal) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPersonalIDs(ids...)
}

// AddCompanyIDs adds the "companys" edge to the Company entity by IDs.
func (uu *UserUpdate) AddCompanyIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCompanyIDs(ids...)
	return uu
}

// AddCompanys adds the "companys" edges to the Company entity.
func (uu *UserUpdate) AddCompanys(c ...*Company) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCompanyIDs(ids...)
}

// SetParent sets the "parent" edge to the User entity.
func (uu *UserUpdate) SetParent(u *User) *UserUpdate {
	return uu.SetParentID(u.ID)
}

// AddChildIDs adds the "children" edge to the User entity by IDs.
func (uu *UserUpdate) AddChildIDs(ids ...int) *UserUpdate {
	uu.mutation.AddChildIDs(ids...)
	return uu
}

// AddChildren adds the "children" edges to the User entity.
func (uu *UserUpdate) AddChildren(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddChildIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearPersonals clears all "personals" edges to the Personal entity.
func (uu *UserUpdate) ClearPersonals() *UserUpdate {
	uu.mutation.ClearPersonals()
	return uu
}

// RemovePersonalIDs removes the "personals" edge to Personal entities by IDs.
func (uu *UserUpdate) RemovePersonalIDs(ids ...int) *UserUpdate {
	uu.mutation.RemovePersonalIDs(ids...)
	return uu
}

// RemovePersonals removes "personals" edges to Personal entities.
func (uu *UserUpdate) RemovePersonals(p ...*Personal) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePersonalIDs(ids...)
}

// ClearCompanys clears all "companys" edges to the Company entity.
func (uu *UserUpdate) ClearCompanys() *UserUpdate {
	uu.mutation.ClearCompanys()
	return uu
}

// RemoveCompanyIDs removes the "companys" edge to Company entities by IDs.
func (uu *UserUpdate) RemoveCompanyIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCompanyIDs(ids...)
	return uu
}

// RemoveCompanys removes "companys" edges to Company entities.
func (uu *UserUpdate) RemoveCompanys(c ...*Company) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCompanyIDs(ids...)
}

// ClearParent clears the "parent" edge to the User entity.
func (uu *UserUpdate) ClearParent() *UserUpdate {
	uu.mutation.ClearParent()
	return uu
}

// ClearChildren clears all "children" edges to the User entity.
func (uu *UserUpdate) ClearChildren() *UserUpdate {
	uu.mutation.ClearChildren()
	return uu
}

// RemoveChildIDs removes the "children" edge to User entities by IDs.
func (uu *UserUpdate) RemoveChildIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveChildIDs(ids...)
	return uu
}

// RemoveChildren removes "children" edges to User entities.
func (uu *UserUpdate) RemoveChildren(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedUserName(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldUserName, value)
		})
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedPassword(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldPassword, value)
		})
	}
	if uu.mutation.PersonalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PersonalsTable,
			Columns: []string{user.PersonalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personal.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPersonalsIDs(); len(nodes) > 0 && !uu.mutation.PersonalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PersonalsTable,
			Columns: []string{user.PersonalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PersonalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PersonalsTable,
			Columns: []string{user.PersonalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.CompanysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CompanysTable,
			Columns: []string{user.CompanysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCompanysIDs(); len(nodes) > 0 && !uu.mutation.CompanysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CompanysTable,
			Columns: []string{user.CompanysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CompanysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CompanysTable,
			Columns: []string{user.CompanysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !uu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUserName sets the "user_name" field.
func (uuo *UserUpdateOne) SetUserName(s []string) *UserUpdateOne {
	uuo.mutation.SetUserName(s)
	return uuo
}

// AppendUserName appends s to the "user_name" field.
func (uuo *UserUpdateOne) AppendUserName(s []string) *UserUpdateOne {
	uuo.mutation.AppendUserName(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s []string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// AppendPassword appends s to the "password" field.
func (uuo *UserUpdateOne) AppendPassword(s []string) *UserUpdateOne {
	uuo.mutation.AppendPassword(s)
	return uuo
}

// SetParentID sets the "parent_id" field.
func (uuo *UserUpdateOne) SetParentID(i int) *UserUpdateOne {
	uuo.mutation.SetParentID(i)
	return uuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableParentID(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetParentID(*i)
	}
	return uuo
}

// ClearParentID clears the value of the "parent_id" field.
func (uuo *UserUpdateOne) ClearParentID() *UserUpdateOne {
	uuo.mutation.ClearParentID()
	return uuo
}

// AddPersonalIDs adds the "personals" edge to the Personal entity by IDs.
func (uuo *UserUpdateOne) AddPersonalIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddPersonalIDs(ids...)
	return uuo
}

// AddPersonals adds the "personals" edges to the Personal entity.
func (uuo *UserUpdateOne) AddPersonals(p ...*Personal) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPersonalIDs(ids...)
}

// AddCompanyIDs adds the "companys" edge to the Company entity by IDs.
func (uuo *UserUpdateOne) AddCompanyIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCompanyIDs(ids...)
	return uuo
}

// AddCompanys adds the "companys" edges to the Company entity.
func (uuo *UserUpdateOne) AddCompanys(c ...*Company) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCompanyIDs(ids...)
}

// SetParent sets the "parent" edge to the User entity.
func (uuo *UserUpdateOne) SetParent(u *User) *UserUpdateOne {
	return uuo.SetParentID(u.ID)
}

// AddChildIDs adds the "children" edge to the User entity by IDs.
func (uuo *UserUpdateOne) AddChildIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddChildIDs(ids...)
	return uuo
}

// AddChildren adds the "children" edges to the User entity.
func (uuo *UserUpdateOne) AddChildren(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddChildIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearPersonals clears all "personals" edges to the Personal entity.
func (uuo *UserUpdateOne) ClearPersonals() *UserUpdateOne {
	uuo.mutation.ClearPersonals()
	return uuo
}

// RemovePersonalIDs removes the "personals" edge to Personal entities by IDs.
func (uuo *UserUpdateOne) RemovePersonalIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemovePersonalIDs(ids...)
	return uuo
}

// RemovePersonals removes "personals" edges to Personal entities.
func (uuo *UserUpdateOne) RemovePersonals(p ...*Personal) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePersonalIDs(ids...)
}

// ClearCompanys clears all "companys" edges to the Company entity.
func (uuo *UserUpdateOne) ClearCompanys() *UserUpdateOne {
	uuo.mutation.ClearCompanys()
	return uuo
}

// RemoveCompanyIDs removes the "companys" edge to Company entities by IDs.
func (uuo *UserUpdateOne) RemoveCompanyIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCompanyIDs(ids...)
	return uuo
}

// RemoveCompanys removes "companys" edges to Company entities.
func (uuo *UserUpdateOne) RemoveCompanys(c ...*Company) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCompanyIDs(ids...)
}

// ClearParent clears the "parent" edge to the User entity.
func (uuo *UserUpdateOne) ClearParent() *UserUpdateOne {
	uuo.mutation.ClearParent()
	return uuo
}

// ClearChildren clears all "children" edges to the User entity.
func (uuo *UserUpdateOne) ClearChildren() *UserUpdateOne {
	uuo.mutation.ClearChildren()
	return uuo
}

// RemoveChildIDs removes the "children" edge to User entities by IDs.
func (uuo *UserUpdateOne) RemoveChildIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveChildIDs(ids...)
	return uuo
}

// RemoveChildren removes "children" edges to User entities.
func (uuo *UserUpdateOne) RemoveChildren(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveChildIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedUserName(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldUserName, value)
		})
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedPassword(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldPassword, value)
		})
	}
	if uuo.mutation.PersonalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PersonalsTable,
			Columns: []string{user.PersonalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personal.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPersonalsIDs(); len(nodes) > 0 && !uuo.mutation.PersonalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PersonalsTable,
			Columns: []string{user.PersonalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PersonalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PersonalsTable,
			Columns: []string{user.PersonalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.CompanysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CompanysTable,
			Columns: []string{user.CompanysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCompanysIDs(); len(nodes) > 0 && !uuo.mutation.CompanysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CompanysTable,
			Columns: []string{user.CompanysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CompanysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CompanysTable,
			Columns: []string{user.CompanysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !uuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
