PHPSTAN_VERSION ?= 0.11.16

lint:
	@test -f ${HOME}/.cache/composer/phpstan-${PHPSTAN_VERSION}.phar || (mkdir -p ${HOME}/.cache/composer/ && wget https://github.com/phpstan/phpstan/releases/download/${PHPSTAN_VERSION}/phpstan.phar -O ${HOME}/.cache/composer/phpstan-${PHPSTAN_VERSION}.phar)
	@php $$HOME/.cache/composer/phpstan-${PHPSTAN_VERSION}.phar analyze -l 7 -c phpstan.neon ./src

docker-lint:
	@docker run -v $$PWD:/app --rm phpstan/phpstan analyze -l 7 -c phpstan.neon ./src

test:
	@php -derror_reporting="E_ALL & ~E_DEPRECATED" vendor/bin/phpunit --configuration phpunit.xml

test-coverage:
	@php -derror_reporting="E_ALL & ~E_DEPRECATED" -dzend_extension=xdebug.so vendor/bin/phpunit --configuration phpunit.xml --coverage-text --coverage-clover=coverage.xml

test-go:
	@cd tests/resources/go && GO111MODULE=on go test ./...

lint-go:
	@cd tests/resources/go && golangci-lint run --enable-all --disable lll,maligned,gochecknoglobals,goimports,gofmt,funlen ./...

docs:
	@php ./vendor/bin/phpdoc-md generate src > API.md

