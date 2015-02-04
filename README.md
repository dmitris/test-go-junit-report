test-go-junit-report
====================

A minimal example to investigate failures in github.com/jstemmer/go-junit-report.

Tests are failing because of the test output when t.Parallel() is used.

Output with t.Parallel() in the test code:  
```
$ go test -v
=== RUN TestDoFoo
=== RUN TestDoFoo2
--- PASS: TestDoFoo (0.00s)
--- PASS: TestDoFoo2 (0.00s)
PASS
ok  	github.com/dmitris/test-go-junit-report	0.007s
```
Output with t.Parallel() in the test code:   
```
$ go test -v
=== RUN TestDoFoo
--- PASS: TestDoFoo (0.00s)
=== RUN TestDoFoo2
--- PASS: TestDoFoo2 (0.00s)
PASS
ok  	github.com/dmitris/test-go-junit-report	0.005s
```