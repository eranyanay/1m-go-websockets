This example shows how to implement an asynchronous I/O mechanism in order to reduce the number of running goroutines

This allows to use a single goroutine to detect when a connection has new data that is available to read/write
