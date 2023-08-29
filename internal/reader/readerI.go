package reader

import "gen/internal/generator"

type Reader interface {
	ParseTreeGenerator(source string) generator.Generator
	ParseStructGenerator() generator.Generator
}
