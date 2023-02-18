# goext_example

## What is this

This sample codes aims to confirm how gohan extension for golang works

## Reference documents

You need to confirm following documents

- [Golang (plugin) extensions](https://github.com/cloudwan/gohan/blob/master/docs/golang_extension.md)
- [Gohan extension](https://github.com/cloudwan/gohan/blob/master/docs/extension.md)

## Prerequisite

### Installation

(1) Build Gohan

    % git clone https://github.com/cloudwan/gohan.git
    % cd gohan
    % go install
    % cd ..

(2) Prepare goext_example environment

    % git clone https://github.com/ttsubo2000/goext_example.git
    % cd goext_example 
    % go mod tidy
    % make
    % cd tests
    % make
    % cd ..

### How to Run

    % gohan test_ex
    15:35:47.083: extest NOTICE  Found 0 test file(s) of type 'js' in path(s): .
    15:35:47.083: extest INFO  Run 0 test files.
    15:35:47.083: extest NOTICE  All tests have passed.
    15:35:47.083: extest NOTICE  Found 1 test file(s) of type 'so' in path(s): .
    15:35:47.084: gohan.extension.framework.goplugin.runner NOTICE  Running Go extensions test: tests/test_example.so
    15:35:47.088: gohan.schema INFO  Loading schema todo/entry.yaml ...
    15:35:47.089: gohan.schema INFO  Loading schema todo/link.yaml ...
    15:35:47.089: gohan.extension.goplugin DEBUG  Loading go extension: tests/../example.so
    15:35:47.092: gohan.extension.goplugin DEBUG  Loading go extension: /home/ttsubo/source/ttsubo2000/goext_example/example.so
    15:35:47.092: gohan.extension.goplugin DEBUG  Starting go environment: Go test environment
    15:35:47.092: gohan.extension.goplugin DEBUG  Calling before start for go environment: Go test environment

    <snip>

    15:35:47.096: gohan.extension.framework.goplugin.runner NOTICE  Go extension test finished: tests/test_example.so
    --------------------------------------------------------------------------------
    Failures:

    ••
    Report:

    +-----+-----------------------+-------+--------+--------+---------+---------+-----------+
    | NO  |         NAME          | TOTAL | PASSED | FAILED | SKIPPED | PENDING | TIME [MS] |
    +-----+-----------------------+-------+--------+--------+---------+---------+-----------+
    |   1 | tests/test_example.so |     2 |      2 |      0 |       0 |       0 |      2.17 |
    |     |               SUMMARY |     2 |      2 |      0 |       0 |       0 |      2.17 |
    +-----+-----------------------+-------+--------+--------+---------+---------+-----------+
