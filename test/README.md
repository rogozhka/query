# Testing
```bash

#
# 1) Assume you have installed package (via modules or go get)
#

#
# 2) Set required ENV variables
#
export TEST_QUERY_HOST="database host"
export TEST_QUERY_USER="database user"
export TEST_QUERY_PASSWORD="your test password"
export TEST_QUERY_DATABASE="database name"

#
# 3) Import required sql dump once. Currently i've made dump for MySQL
# and check number of rows, presence of columns and their values.
#
# If you prefer other rdbms, change driver and imports in main_test.go
#
mysql -h $TEST_QUERY_HOST -D $TEST_QUERY_DATABASE -u $TEST_QUERY_USER -p $TEST_QUERY_PASSWORD < ./src/github.com/rogozhka/query/test/dump.sql

#
# 4) Run the test
#
go test github.com/rogozhka/query

```



