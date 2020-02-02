module http-health-checker

require (
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2 // indirect
	gopkg.in/h2non/baloo.v3 v3.0.2
	gopkg.in/h2non/gentleman.v2 v2.0.3 // indirect
	pkg/healthchecker v0.0.0
)

replace pkg/healthchecker => ./pkg/healthchecker

go 1.11
