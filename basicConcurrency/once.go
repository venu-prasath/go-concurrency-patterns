package basicConcurrency

import (
	"fmt"
	"sync"
)

/* This is an example code that uses sync.Once to implement a Singleton package.
 * The Singleton pattern is useful when you want to ensure that an application has
 * only one instance of a resource and provides a global point of access to it.
 */

type DBConnection struct {}

var (
	dbConnOnce sync.Once
	conn *DBConnection
)


func GetConnection() *DBConnection {
	dbConnOnce.Do(func() {
		fmt.Println("Initializing a database connection")
		conn = &DBConnection{}
	})
	return conn
}

func TestOnce() {
	var wg sync.WaitGroup
	wg.Add(5)

	//simulate 5 goroutines trying to create a database connection
	for i:=0; i<5; i++ {
		go func() {
			defer wg.Done()
			conn = GetConnection()
		}()
	}
	wg.Wait()
}