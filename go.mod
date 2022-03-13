module code.m-spezial.de/M-Spezial/go-mblazed

go 1.16

require (
	github.com/julienschmidt/httprouter v1.3.0
	github.com/matoous/go-nanoid v1.5.0
	github.com/pkg/errors v0.9.1 // indirect
	go.uber.org/zap v1.13.0
	gorm.io/driver/postgres v1.3.1
	gorm.io/gorm v1.23.2
)

replace (
	github.com/CloudyKit/jet/v6 v6.1.0 => code.m-spezial.de/GitHub-Mirrors/CloudyKit-jet/v6 v6.1.0
	github.com/jackc/pgx/v4 v4.13.0 => code.m-spezial.de/GitHub-Mirrors/jackc-pgx/v4 v4.13.0
)
