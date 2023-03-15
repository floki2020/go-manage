// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go-manage/pkg/ent/migrate"

	"go-manage/pkg/ent/company"
	"go-manage/pkg/ent/personal"
	"go-manage/pkg/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Company is the client for interacting with the Company builders.
	Company *CompanyClient
	// Personal is the client for interacting with the Personal builders.
	Personal *PersonalClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Company = NewCompanyClient(c.config)
	c.Personal = NewPersonalClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Company:  NewCompanyClient(cfg),
		Personal: NewPersonalClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Company:  NewCompanyClient(cfg),
		Personal: NewPersonalClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Company.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Company.Use(hooks...)
	c.Personal.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Company.Intercept(interceptors...)
	c.Personal.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CompanyMutation:
		return c.Company.mutate(ctx, m)
	case *PersonalMutation:
		return c.Personal.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CompanyClient is a client for the Company schema.
type CompanyClient struct {
	config
}

// NewCompanyClient returns a client for the Company from the given config.
func NewCompanyClient(c config) *CompanyClient {
	return &CompanyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `company.Hooks(f(g(h())))`.
func (c *CompanyClient) Use(hooks ...Hook) {
	c.hooks.Company = append(c.hooks.Company, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `company.Intercept(f(g(h())))`.
func (c *CompanyClient) Intercept(interceptors ...Interceptor) {
	c.inters.Company = append(c.inters.Company, interceptors...)
}

// Create returns a builder for creating a Company entity.
func (c *CompanyClient) Create() *CompanyCreate {
	mutation := newCompanyMutation(c.config, OpCreate)
	return &CompanyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Company entities.
func (c *CompanyClient) CreateBulk(builders ...*CompanyCreate) *CompanyCreateBulk {
	return &CompanyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Company.
func (c *CompanyClient) Update() *CompanyUpdate {
	mutation := newCompanyMutation(c.config, OpUpdate)
	return &CompanyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CompanyClient) UpdateOne(co *Company) *CompanyUpdateOne {
	mutation := newCompanyMutation(c.config, OpUpdateOne, withCompany(co))
	return &CompanyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CompanyClient) UpdateOneID(id int) *CompanyUpdateOne {
	mutation := newCompanyMutation(c.config, OpUpdateOne, withCompanyID(id))
	return &CompanyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Company.
func (c *CompanyClient) Delete() *CompanyDelete {
	mutation := newCompanyMutation(c.config, OpDelete)
	return &CompanyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CompanyClient) DeleteOne(co *Company) *CompanyDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CompanyClient) DeleteOneID(id int) *CompanyDeleteOne {
	builder := c.Delete().Where(company.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CompanyDeleteOne{builder}
}

// Query returns a query builder for Company.
func (c *CompanyClient) Query() *CompanyQuery {
	return &CompanyQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCompany},
		inters: c.Interceptors(),
	}
}

// Get returns a Company entity by its id.
func (c *CompanyClient) Get(ctx context.Context, id int) (*Company, error) {
	return c.Query().Where(company.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CompanyClient) GetX(ctx context.Context, id int) *Company {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Company.
func (c *CompanyClient) QueryUsers(co *Company) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(company.Table, company.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, company.UsersTable, company.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CompanyClient) Hooks() []Hook {
	return c.hooks.Company
}

// Interceptors returns the client interceptors.
func (c *CompanyClient) Interceptors() []Interceptor {
	return c.inters.Company
}

func (c *CompanyClient) mutate(ctx context.Context, m *CompanyMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CompanyCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CompanyUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CompanyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CompanyDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Company mutation op: %q", m.Op())
	}
}

// PersonalClient is a client for the Personal schema.
type PersonalClient struct {
	config
}

// NewPersonalClient returns a client for the Personal from the given config.
func NewPersonalClient(c config) *PersonalClient {
	return &PersonalClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `personal.Hooks(f(g(h())))`.
func (c *PersonalClient) Use(hooks ...Hook) {
	c.hooks.Personal = append(c.hooks.Personal, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `personal.Intercept(f(g(h())))`.
func (c *PersonalClient) Intercept(interceptors ...Interceptor) {
	c.inters.Personal = append(c.inters.Personal, interceptors...)
}

// Create returns a builder for creating a Personal entity.
func (c *PersonalClient) Create() *PersonalCreate {
	mutation := newPersonalMutation(c.config, OpCreate)
	return &PersonalCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Personal entities.
func (c *PersonalClient) CreateBulk(builders ...*PersonalCreate) *PersonalCreateBulk {
	return &PersonalCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Personal.
func (c *PersonalClient) Update() *PersonalUpdate {
	mutation := newPersonalMutation(c.config, OpUpdate)
	return &PersonalUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PersonalClient) UpdateOne(pe *Personal) *PersonalUpdateOne {
	mutation := newPersonalMutation(c.config, OpUpdateOne, withPersonal(pe))
	return &PersonalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PersonalClient) UpdateOneID(id int) *PersonalUpdateOne {
	mutation := newPersonalMutation(c.config, OpUpdateOne, withPersonalID(id))
	return &PersonalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Personal.
func (c *PersonalClient) Delete() *PersonalDelete {
	mutation := newPersonalMutation(c.config, OpDelete)
	return &PersonalDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PersonalClient) DeleteOne(pe *Personal) *PersonalDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PersonalClient) DeleteOneID(id int) *PersonalDeleteOne {
	builder := c.Delete().Where(personal.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PersonalDeleteOne{builder}
}

// Query returns a query builder for Personal.
func (c *PersonalClient) Query() *PersonalQuery {
	return &PersonalQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePersonal},
		inters: c.Interceptors(),
	}
}

// Get returns a Personal entity by its id.
func (c *PersonalClient) Get(ctx context.Context, id int) (*Personal, error) {
	return c.Query().Where(personal.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PersonalClient) GetX(ctx context.Context, id int) *Personal {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Personal.
func (c *PersonalClient) QueryUsers(pe *Personal) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(personal.Table, personal.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, personal.UsersTable, personal.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PersonalClient) Hooks() []Hook {
	return c.hooks.Personal
}

// Interceptors returns the client interceptors.
func (c *PersonalClient) Interceptors() []Interceptor {
	return c.inters.Personal
}

func (c *PersonalClient) mutate(ctx context.Context, m *PersonalMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PersonalCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PersonalUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PersonalUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PersonalDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Personal mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPersonals queries the personals edge of a User.
func (c *UserClient) QueryPersonals(u *User) *PersonalQuery {
	query := (&PersonalClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(personal.Table, personal.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PersonalsTable, user.PersonalsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCompanys queries the companys edge of a User.
func (c *UserClient) QueryCompanys(u *User) *CompanyQuery {
	query := (&CompanyClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CompanysTable, user.CompanysColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParent queries the parent edge of a User.
func (c *UserClient) QueryParent(u *User) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.ParentTable, user.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a User.
func (c *UserClient) QueryChildren(u *User) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ChildrenTable, user.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Company, Personal, User []ent.Hook
	}
	inters struct {
		Company, Personal, User []ent.Interceptor
	}
)
