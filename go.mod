module github.com/jdxj/share

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/jdxj/logger v0.1.0
	github.com/micro/go-micro/v3 v3.0.0-beta.2.0.20200911124113-3bb76868d194
	github.com/micro/micro/v3 v3.0.0-beta.4
	github.com/satori/go.uuid v1.2.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.3.0
)
