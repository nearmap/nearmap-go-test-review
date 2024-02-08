# Nearmap Go Test (Code Review Version)

Instructions for the standard Nearmap Go test are provided below.

For this exercise, a full solution has been provided (there is no need to
provide any code yourself). However, we would like you to review the code so
that we can have an in-depth technical discussion around the solution and how
it can be improved.

The purpose of this assignment is to test your familiarity with Go, distributed
systems concepts and unit testing.

## Background

The source code that you are given is a very simple imitation of a key/value
store:

* `Database` represents a client to the central store that takes a long time
  (500ms) to store and retrieve data.
* `DistributedCache` represents a client to the distributed cache (Redis for
  example) that takes much less time to turn around (100ms to store or retrieve).

This scenario is a simplified example of a typical high performance server
cluster with a database, a distributed cache and multiple worker nodes.

## Assumptions and requirements


* Data in `Database` never changes and can be cached forever.
* If `Database.Value()` returns `nil` for a key, the requested data item does
  not exist and will never exist.
* `DistributedCache` is initially empty.
* For a frequently-requested item your `DataSource.Value()` implementation should
have a better response time than the distributed cache store (ie < 100ms).
* The user of the `DataSource` interface must not have to deal with thread
  synchronisation.
* Sufficient unit test coverage for the `DataSource` implementation.
* The solution should aim to minimise calls to the database.
* Use 10 goroutines (simulating separate threads on a single worker node) each
  making 50 consecutive requests for a random key in the range (key0-key9). I.e. there should be a total of 500 requests.
* For each request, print the requested key name, returned value, time to complete that request; similar to the following example:

      [1] Request 'key1', response 'value1', time: 50.05 ms
      [2] Request 'key2', response 'value2', time: 50.05 ms
