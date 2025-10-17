module ProductCartService

go 1.24.3

require (
	github.com/gojuno/minimock/v3 v3.4.7
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace ProductCartService/internal/pkg/cart/service/mocks => ./internal/pkg/cart/service/mocks
