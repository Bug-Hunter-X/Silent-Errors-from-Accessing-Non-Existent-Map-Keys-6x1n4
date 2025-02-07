# Silent Errors from Accessing Non-Existent Map Keys in Go

This example demonstrates a common, yet subtle, error in Go when working with maps: accessing a key that doesn't exist.

In many languages, this would raise an exception. However, in Go, accessing a non-existent key in a map simply returns the zero value for the map's value type (in this case, 0 for an `int`). This can lead to hard-to-debug issues because there's no immediate error or panic.  The program continues execution, potentially with incorrect results.

The solution shows how to safely access map values using the comma ok idiom to explicitly check for the existence of a key.