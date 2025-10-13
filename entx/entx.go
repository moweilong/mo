package entx

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/XSAM/otelsql"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"

	entSql "entgo.io/ent/dialect/sql"
)

// CreateDriver 创建数据库驱动
func CreateDriver(driverName, dsn string, enableTrace, enableMetrics bool) (*entSql.Driver, error) {
	var db *sql.DB
	var drv *entSql.Driver
	var err error

	if enableTrace {
		if db, err = otelsql.Open(driverName, dsn, otelsql.WithAttributes(
			driverNameToSemConvKeyValue(driverName)),
		); err != nil {
			return nil, fmt.Errorf("failed opening connection to db: %w", err)
		}
		drv = entSql.OpenDB(driverName, db)
	} else {
		if drv, err = entSql.Open(driverName, dsn); err != nil {
			return nil, fmt.Errorf("failed opening connection to db: %w", err)
		}

		db = drv.DB()
	}

	if enableMetrics {
		if err = otelsql.RegisterDBStatsMetrics(db, otelsql.WithAttributes(
			driverNameToSemConvKeyValue(driverName)),
		); err != nil {
			return nil, fmt.Errorf("failed registering db stats metrics: %w", err)
		}
	}

	return drv, nil
}

func driverNameToSemConvKeyValue(driverName string) attribute.KeyValue {
	switch driverName {
	case "mariadb":
		return semconv.DBSystemNameMariaDB
	case "mysql":
		return semconv.DBSystemNameMySQL
	case "postgresql":
		return semconv.DBSystemNamePostgreSQL
	case "sqlite":
		return semconv.DBSystemNameSqlite
	default:
		return semconv.DBSystemNameKey.String(driverName)
	}
}

func NewEntClient[T IEntClient](db T, driver *entSql.Driver) *EntClient[T] {
	return &EntClient[T]{
		db:     db,
		driver: driver,
	}
}

type EntClient[T IEntClient] struct {
	db     T
	driver *entSql.Driver
}

type IEntClient interface {
	Close() error
}

func (e *EntClient[T]) Client() T {
	return e.db
}

func (e *EntClient[T]) Driver() *entSql.Driver {
	return e.driver
}

func (e *EntClient[T]) DB() *sql.DB {
	return e.driver.DB()
}

func (e *EntClient[T]) Close() error {
	return e.db.Close()
}

// Query 查询数据
func (e *EntClient[T]) Query(ctx context.Context, query string, args, v any) error {
	return e.driver.Query(ctx, query, args, v)
}

// Exec 执行语句
func (e *EntClient[T]) Exec(ctx context.Context, query string, args, v any) error {
	return e.driver.Exec(ctx, query, args, v)
}

// SetConnectionOption 设置连接配置
func (e *EntClient[T]) SetConnectionOption(maxIdleConnections, maxOpenConnections int, connMaxLifetime time.Duration) {
	// 连接池中最多保留的空闲连接数量
	e.DB().SetMaxIdleConns(maxIdleConnections)
	// 连接池在同一时间打开连接的最大数量
	e.DB().SetMaxOpenConns(maxOpenConnections)
	// 连接可重用的最大时间长度
	e.DB().SetConnMaxLifetime(connMaxLifetime)
}
