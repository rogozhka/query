# Testing
```bash

#
# Assume you have a project
# and imported query by `go get -u github.com/rogozhka/query`
# and current dir is your GOPATH
#

#
# Set required ENV variables first
#
export TEST_QUERY_HOST="database host"
export TEST_QUERY_USER="database user"
export TEST_QUERY_PASSWORD="your test password"
export TEST_QUERY_DATABASE="database name"

#
# Import required sql dump once. Currently i've made dump for MySQL
# and check number of rows, presence of columns and their values.
#
# If you prefer other rdbms, change driver and imports in main_test.go
#
mysql -h $TEST_QUERY_HOST -D $TEST_QUERY_DATABASE -u $TEST_QUERY_USER -p $TEST_QUERY_PASSWORD < ./src/github.com/rogozhka/query/test/dump.sql

#
# Run the test
#
go test ./src/github.com/rogozhka/query/

```



