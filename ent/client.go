// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/zhanghongnian/brand-index/ent/migrate"

	"github.com/zhanghongnian/brand-index/ent/qccenterprisedata"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// QccEnterpriseData is the client for interacting with the QccEnterpriseData builders.
	QccEnterpriseData *QccEnterpriseDataClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.QccEnterpriseData = NewQccEnterpriseDataClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:               ctx,
		config:            cfg,
		QccEnterpriseData: NewQccEnterpriseDataClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:            cfg,
		QccEnterpriseData: NewQccEnterpriseDataClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		QccEnterpriseData.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.QccEnterpriseData.Use(hooks...)
}

// QccEnterpriseDataClient is a client for the QccEnterpriseData schema.
type QccEnterpriseDataClient struct {
	config
}

// NewQccEnterpriseDataClient returns a client for the QccEnterpriseData from the given config.
func NewQccEnterpriseDataClient(c config) *QccEnterpriseDataClient {
	return &QccEnterpriseDataClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `qccenterprisedata.Hooks(f(g(h())))`.
func (c *QccEnterpriseDataClient) Use(hooks ...Hook) {
	c.hooks.QccEnterpriseData = append(c.hooks.QccEnterpriseData, hooks...)
}

// Create returns a create builder for QccEnterpriseData.
func (c *QccEnterpriseDataClient) Create() *QccEnterpriseDataCreate {
	mutation := newQccEnterpriseDataMutation(c.config, OpCreate)
	return &QccEnterpriseDataCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of QccEnterpriseData entities.
func (c *QccEnterpriseDataClient) CreateBulk(builders ...*QccEnterpriseDataCreate) *QccEnterpriseDataCreateBulk {
	return &QccEnterpriseDataCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for QccEnterpriseData.
func (c *QccEnterpriseDataClient) Update() *QccEnterpriseDataUpdate {
	mutation := newQccEnterpriseDataMutation(c.config, OpUpdate)
	return &QccEnterpriseDataUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *QccEnterpriseDataClient) UpdateOne(qed *QccEnterpriseData) *QccEnterpriseDataUpdateOne {
	mutation := newQccEnterpriseDataMutation(c.config, OpUpdateOne, withQccEnterpriseData(qed))
	return &QccEnterpriseDataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *QccEnterpriseDataClient) UpdateOneID(id int) *QccEnterpriseDataUpdateOne {
	mutation := newQccEnterpriseDataMutation(c.config, OpUpdateOne, withQccEnterpriseDataID(id))
	return &QccEnterpriseDataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for QccEnterpriseData.
func (c *QccEnterpriseDataClient) Delete() *QccEnterpriseDataDelete {
	mutation := newQccEnterpriseDataMutation(c.config, OpDelete)
	return &QccEnterpriseDataDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *QccEnterpriseDataClient) DeleteOne(qed *QccEnterpriseData) *QccEnterpriseDataDeleteOne {
	return c.DeleteOneID(qed.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *QccEnterpriseDataClient) DeleteOneID(id int) *QccEnterpriseDataDeleteOne {
	builder := c.Delete().Where(qccenterprisedata.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &QccEnterpriseDataDeleteOne{builder}
}

// Query returns a query builder for QccEnterpriseData.
func (c *QccEnterpriseDataClient) Query() *QccEnterpriseDataQuery {
	return &QccEnterpriseDataQuery{config: c.config}
}

// Get returns a QccEnterpriseData entity by its id.
func (c *QccEnterpriseDataClient) Get(ctx context.Context, id int) (*QccEnterpriseData, error) {
	return c.Query().Where(qccenterprisedata.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *QccEnterpriseDataClient) GetX(ctx context.Context, id int) *QccEnterpriseData {
	qed, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return qed
}

// Hooks returns the client hooks.
func (c *QccEnterpriseDataClient) Hooks() []Hook {
	return c.hooks.QccEnterpriseData
}
