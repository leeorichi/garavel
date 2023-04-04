package generator

import (
	"github.com/leeorichi/getgo/api/internal/pkg/snowflake"
)

var (
	ProductSNF snowflake.SnowflakeGenerator
)

func InitSnowflakeGenerators() error {
	ProductSNF = snowflake.New()

	return nil
}
