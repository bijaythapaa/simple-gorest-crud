module github.com/dbijay/simple-gorest-crud

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-ozzo/ozzo-dbx v1.5.0
	github.com/go-ozzo/ozzo-routing/v2 v2.3.0
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/google/uuid v1.2.0
	github.com/lib/pq v1.10.1
	github.com/qiangxue/go-env v1.0.1
	github.com/stretchr/testify v1.4.0
	go.uber.org/zap v1.16.0
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/dbijayagithub.com/dbijay/simple-gorest-crud => ../simple-gorest-crud
