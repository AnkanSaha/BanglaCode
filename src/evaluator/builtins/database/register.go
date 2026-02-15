package database

import (
	"BanglaCode/src/evaluator/builtins/database/mongodb"
	"BanglaCode/src/evaluator/builtins/database/mysql"
	"BanglaCode/src/evaluator/builtins/database/postgres"
	"BanglaCode/src/evaluator/builtins/database/redis"
	"BanglaCode/src/object"
)

// Builtins holds all database built-in functions merged from all connectors
var Builtins = make(map[string]*object.Builtin)

func init() {
	// Merge PostgreSQL built-ins
	for name, fn := range postgres.Builtins {
		Builtins[name] = fn
	}

	// Merge MySQL built-ins
	for name, fn := range mysql.Builtins {
		Builtins[name] = fn
	}

	// Merge MongoDB built-ins
	for name, fn := range mongodb.Builtins {
		Builtins[name] = fn
	}

	// Merge Redis built-ins
	for name, fn := range redis.Builtins {
		Builtins[name] = fn
	}

	// Register unified database functions (database-agnostic)
	registerUnifiedBuiltins()
}

// registerUnifiedBuiltins registers database-agnostic built-in functions
// These functions work with any database type
func registerUnifiedBuiltins() {
	// db_jukto - Universal database connection function
	// Automatically routes to the correct connector based on database type
	Builtins["db_jukto"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("db_jukto: wrong number of arguments. got=%d, want=2 (type, config)", len(args))
			}

			dbType, ok := args[0].(*object.String)
			if !ok {
				return newError("db_jukto: first argument must be STRING (database type), got %s", args[0].Type())
			}

			config, ok := args[1].(*object.Map)
			if !ok {
				return newError("db_jukto: second argument must be MAP (config), got %s", args[1].Type())
			}

			// Route to appropriate connector
			switch dbType.Value {
			case "postgres", "postgresql":
				return postgres.Builtins["db_jukto_postgres"].Fn(config)
			case "mysql":
				return mysql.Builtins["db_jukto_mysql"].Fn(config)
			case "mongodb", "mongo":
				return mongodb.Builtins["db_jukto_mongodb"].Fn(config)
			case "redis":
				return redis.Builtins["db_jukto_redis"].Fn(config)
			default:
				return newError("db_jukto: unsupported database type '%s'. Supported: postgres, mysql, mongodb, redis", dbType.Value)
			}
		},
	}

	// db_bandho - Universal database close function
	Builtins["db_bandho"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("db_bandho: wrong number of arguments. got=%d, want=1", len(args))
			}

			conn, ok := args[0].(*object.DBConnection)
			if !ok {
				return newError("db_bandho: argument must be DB_CONNECTION, got %s", args[0].Type())
			}

			// Route to appropriate connector based on connection type
			switch conn.DBType {
			case "postgres":
				return postgres.Builtins["db_bandho_postgres"].Fn(conn)
			case "mysql":
				return mysql.Builtins["db_bandho_mysql"].Fn(conn)
			case "mongodb":
				return mongodb.Builtins["db_bandho_mongodb"].Fn(conn)
			case "redis":
				return redis.Builtins["db_bandho_redis"].Fn(conn)
			default:
				return newError("db_bandho: unsupported connection type '%s'", conn.DBType)
			}
		},
	}

	// db_query - Universal SQL query function (for SQL databases only)
	Builtins["db_query"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("db_query: wrong number of arguments. got=%d, want=2", len(args))
			}

			conn, ok := args[0].(*object.DBConnection)
			if !ok {
				return newError("db_query: first argument must be DB_CONNECTION, got %s", args[0].Type())
			}

			query, ok := args[1].(*object.String)
			if !ok {
				return newError("db_query: second argument must be STRING (SQL query), got %s", args[1].Type())
			}

			// Route to appropriate connector
			switch conn.DBType {
			case "postgres":
				return postgres.Builtins["db_query_postgres"].Fn(conn, query)
			case "mysql":
				return mysql.Builtins["db_query_mysql"].Fn(conn, query)
			case "mongodb", "redis":
				return newError("db_query: %s does not support SQL queries. Use database-specific functions", conn.DBType)
			default:
				return newError("db_query: unsupported connection type '%s'", conn.DBType)
			}
		},
	}

	// db_exec - Universal SQL exec function (for SQL databases only)
	Builtins["db_exec"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("db_exec: wrong number of arguments. got=%d, want=2", len(args))
			}

			conn, ok := args[0].(*object.DBConnection)
			if !ok {
				return newError("db_exec: first argument must be DB_CONNECTION, got %s", args[0].Type())
			}

			query, ok := args[1].(*object.String)
			if !ok {
				return newError("db_exec: second argument must be STRING (SQL statement), got %s", args[1].Type())
			}

			// Route to appropriate connector
			switch conn.DBType {
			case "postgres":
				return postgres.Builtins["db_exec_postgres"].Fn(conn, query)
			case "mysql":
				return mysql.Builtins["db_exec_mysql"].Fn(conn, query)
			case "mongodb", "redis":
				return newError("db_exec: %s does not support SQL statements. Use database-specific functions", conn.DBType)
			default:
				return newError("db_exec: unsupported connection type '%s'", conn.DBType)
			}
		},
	}

	// db_proshno - Universal prepared query function (SQL injection safe)
	Builtins["db_proshno"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return newError("db_proshno: wrong number of arguments. got=%d, want=3", len(args))
			}

			conn, ok := args[0].(*object.DBConnection)
			if !ok {
				return newError("db_proshno: first argument must be DB_CONNECTION, got %s", args[0].Type())
			}

			query, ok := args[1].(*object.String)
			if !ok {
				return newError("db_proshno: second argument must be STRING (SQL query), got %s", args[1].Type())
			}

			params, ok := args[2].(*object.Array)
			if !ok {
				return newError("db_proshno: third argument must be ARRAY (parameters), got %s", args[2].Type())
			}

			// Route to appropriate connector
			switch conn.DBType {
			case "postgres":
				return postgres.Builtins["db_proshno_postgres"].Fn(conn, query, params)
			case "mysql":
				return mysql.Builtins["db_proshno_mysql"].Fn(conn, query, params)
			case "mongodb", "redis":
				return newError("db_proshno: %s does not support SQL prepared statements", conn.DBType)
			default:
				return newError("db_proshno: unsupported connection type '%s'", conn.DBType)
			}
		},
	}
}
