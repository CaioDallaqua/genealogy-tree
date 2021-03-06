package env

import "os"

const (
	GENEALOGY_TREE_ADDR = "GENEALOGY_TREE_ADDR"

	// -------- CODE ENVIRONMENT ----------

	// dev  -> logs to console
	CODE_ENV             = "CODE_ENV"
	CODE_ENV_DEFAULT_VAL = "dev"

	// ----------- DATABASE -------------

	DB_ADDR             = "DB_ADDR"
	DB_USER             = "DB_USER"
	DB_USER_DEFAULT_VAL = "neo4j"
	DB_PWD              = "DB_PWD"
	DB_PWD_DEFAULT_VAL  = "test"
)

// Service ports
// Could be defined by user for applications like port masking, but YAGNI
var svcPorts = map[string]string{
	GENEALOGY_TREE_ADDR: "8998",
	DB_ADDR:             "7687",
}

// Service protocols
var svcProtocols = map[string]string{
	GENEALOGY_TREE_ADDR: "",
	DB_ADDR:             "bolt://",
}

// Application setup to be injected from main
type AppConfig struct {
	CodeEnv string
	Addr    string
	DB      dbConfig
}

type dbConfig struct {
	Addr string
	User string
	Pwd  string
}

func NewConfig() *AppConfig {
	var (
		codeEnv string = getEnv(CODE_ENV, CODE_ENV_DEFAULT_VAL)
		apiAddr string = getEnvAddr(GENEALOGY_TREE_ADDR)

		dbAddr string = getEnvAddr(DB_ADDR)
		dbUser string = getEnv(DB_USER, DB_USER_DEFAULT_VAL)
		dbPwd  string = getEnv(DB_PWD, DB_PWD_DEFAULT_VAL)
	)

	return &AppConfig{
		CodeEnv: codeEnv,
		Addr:    apiAddr,
		DB: dbConfig{
			Addr: dbAddr,
			User: dbUser,
			Pwd:  dbPwd,
		},
	}
}

func getEnv(key, defaultValue string) string {
	var value string = os.Getenv(key)
	if len(value) != 0 {
		return value
	}

	return defaultValue
}

func getEnvAddr(key string) string {
	var addr string = os.Getenv(key)
	if len(addr) == 0 {
		addr = "127.0.0.1"
	}

	return svcProtocols[key] + addr + ":" + svcPorts[key]
}
