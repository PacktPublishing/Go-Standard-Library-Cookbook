package main

import (
	"log"
	"os"
)

func main() {

	key := "DB_CONN"
	// Set the environmental variable.
	os.Setenv(key, "postgres://as:as@example.com/pg?sslmode=verify-full")

	val := GetEnvDefault(key, "postgres://as:as@localhost/pg?sslmode=verify-full")

	log.Println("The value is :" + val)

	os.Unsetenv(key)

	val = GetEnvDefault(key, "postgres://as:as@127.0.0.1/pg?sslmode=verify-full")

	log.Println("The default value is :" + val)

}

func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
