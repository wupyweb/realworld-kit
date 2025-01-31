// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wupyweb/realworld-kit/ent/article"
	"github.com/wupyweb/realworld-kit/ent/favorite"
	"github.com/wupyweb/realworld-kit/ent/predicate"
	"github.com/wupyweb/realworld-kit/ent/user"
)

// FavoriteUpdate is the builder for updating Favorite entities.
type FavoriteUpdate struct {
	config
	hooks    []Hook
	mutation *FavoriteMutation
}

// Where appends a list predicates to the FavoriteUpdate builder.
func (fu *FavoriteUpdate) Where(ps ...predicate.Favorite) *FavoriteUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetFavoritedAt sets the "favorited_at" field.
func (fu *FavoriteUpdate) SetFavoritedAt(t time.Time) *FavoriteUpdate {
	fu.mutation.SetFavoritedAt(t)
	return fu
}

// SetNillableFavoritedAt sets the "favorited_at" field if the given value is not nil.
func (fu *FavoriteUpdate) SetNillableFavoritedAt(t *time.Time) *FavoriteUpdate {
	if t != nil {
		fu.SetFavoritedAt(*t)
	}
	return fu
}

// SetUserID sets the "user_id" field.
func (fu *FavoriteUpdate) SetUserID(i int) *FavoriteUpdate {
	fu.mutation.SetUserID(i)
	return fu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (fu *FavoriteUpdate) SetNillableUserID(i *int) *FavoriteUpdate {
	if i != nil {
		fu.SetUserID(*i)
	}
	return fu
}

// SetArticleID sets the "article_id" field.
func (fu *FavoriteUpdate) SetArticleID(i int) *FavoriteUpdate {
	fu.mutation.SetArticleID(i)
	return fu
}

// SetNillableArticleID sets the "article_id" field if the given value is not nil.
func (fu *FavoriteUpdate) SetNillableArticleID(i *int) *FavoriteUpdate {
	if i != nil {
		fu.SetArticleID(*i)
	}
	return fu
}

// SetUser sets the "user" edge to the User entity.
func (fu *FavoriteUpdate) SetUser(u *User) *FavoriteUpdate {
	return fu.SetUserID(u.ID)
}

// SetArticle sets the "article" edge to the Article entity.
func (fu *FavoriteUpdate) SetArticle(a *Article) *FavoriteUpdate {
	return fu.SetArticleID(a.ID)
}

// Mutation returns the FavoriteMutation object of the builder.
func (fu *FavoriteUpdate) Mutation() *FavoriteMutation {
	return fu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (fu *FavoriteUpdate) ClearUser() *FavoriteUpdate {
	fu.mutation.ClearUser()
	return fu
}

// ClearArticle clears the "article" edge to the Article entity.
func (fu *FavoriteUpdate) ClearArticle() *FavoriteUpdate {
	fu.mutation.ClearArticle()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FavoriteUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FavoriteUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FavoriteUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FavoriteUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FavoriteUpdate) check() error {
	if _, ok := fu.mutation.UserID(); fu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Favorite.user"`)
	}
	if _, ok := fu.mutation.ArticleID(); fu.mutation.ArticleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Favorite.article"`)
	}
	return nil
}

func (fu *FavoriteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(favorite.Table, favorite.Columns, sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.FavoritedAt(); ok {
		_spec.SetField(favorite.FieldFavoritedAt, field.TypeTime, value)
	}
	if fu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   favorite.UserTable,
			Columns: []string{favorite.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   favorite.UserTable,
			Columns: []string{favorite.UserColumn},
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
	if fu.mutation.ArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   favorite.ArticleTable,
			Columns: []string{favorite.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.ArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   favorite.ArticleTable,
			Columns: []string{favorite.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{favorite.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FavoriteUpdateOne is the builder for updating a single Favorite entity.
type FavoriteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FavoriteMutation
}

// SetFavoritedAt sets the "favorited_at" field.
func (fuo *FavoriteUpdateOne) SetFavoritedAt(t time.Time) *FavoriteUpdateOne {
	fuo.mutation.SetFavoritedAt(t)
	return fuo
}

// SetNillableFavoritedAt sets the "favorited_at" field if the given value is not nil.
func (fuo *FavoriteUpdateOne) SetNillableFavoritedAt(t *time.Time) *FavoriteUpdateOne {
	if t != nil {
		fuo.SetFavoritedAt(*t)
	}
	return fuo
}

// SetUserID sets the "user_id" field.
func (fuo *FavoriteUpdateOne) SetUserID(i int) *FavoriteUpdateOne {
	fuo.mutation.SetUserID(i)
	return fuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (fuo *FavoriteUpdateOne) SetNillableUserID(i *int) *FavoriteUpdateOne {
	if i != nil {
		fuo.SetUserID(*i)
	}
	return fuo
}

// SetArticleID sets the "article_id" field.
func (fuo *FavoriteUpdateOne) SetArticleID(i int) *FavoriteUpdateOne {
	fuo.mutation.SetArticleID(i)
	return fuo
}

// SetNillableArticleID sets the "article_id" field if the given value is not nil.
func (fuo *FavoriteUpdateOne) SetNillableArticleID(i *int) *FavoriteUpdateOne {
	if i != nil {
		fuo.SetArticleID(*i)
	}
	return fuo
}

// SetUser sets the "user" edge to the User entity.
func (fuo *FavoriteUpdateOne) SetUser(u *User) *FavoriteUpdateOne {
	return fuo.SetUserID(u.ID)
}

// SetArticle sets the "article" edge to the Article entity.
func (fuo *FavoriteUpdateOne) SetArticle(a *Article) *FavoriteUpdateOne {
	return fuo.SetArticleID(a.ID)
}

// Mutation returns the FavoriteMutation object of the builder.
func (fuo *FavoriteUpdateOne) Mutation() *FavoriteMutation {
	return fuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (fuo *FavoriteUpdateOne) ClearUser() *FavoriteUpdateOne {
	fuo.mutation.ClearUser()
	return fuo
}

// ClearArticle clears the "article" edge to the Article entity.
func (fuo *FavoriteUpdateOne) ClearArticle() *FavoriteUpdateOne {
	fuo.mutation.ClearArticle()
	return fuo
}

// Where appends a list predicates to the FavoriteUpdate builder.
func (fuo *FavoriteUpdateOne) Where(ps ...predicate.Favorite) *FavoriteUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FavoriteUpdateOne) Select(field string, fields ...string) *FavoriteUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Favorite entity.
func (fuo *FavoriteUpdateOne) Save(ctx context.Context) (*Favorite, error) {
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FavoriteUpdateOne) SaveX(ctx context.Context) *Favorite {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FavoriteUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FavoriteUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FavoriteUpdateOne) check() error {
	if _, ok := fuo.mutation.UserID(); fuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Favorite.user"`)
	}
	if _, ok := fuo.mutation.ArticleID(); fuo.mutation.ArticleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Favorite.article"`)
	}
	return nil
}

func (fuo *FavoriteUpdateOne) sqlSave(ctx context.Context) (_node *Favorite, err error) {
	if err := fuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(favorite.Table, favorite.Columns, sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Favorite.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, favorite.FieldID)
		for _, f := range fields {
			if !favorite.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != favorite.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.FavoritedAt(); ok {
		_spec.SetField(favorite.FieldFavoritedAt, field.TypeTime, value)
	}
	if fuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   favorite.UserTable,
			Columns: []string{favorite.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   favorite.UserTable,
			Columns: []string{favorite.UserColumn},
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
	if fuo.mutation.ArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   favorite.ArticleTable,
			Columns: []string{favorite.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.ArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   favorite.ArticleTable,
			Columns: []string{favorite.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Favorite{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{favorite.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
