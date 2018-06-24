package klarna

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// Set environment variables for subsequent tests.
	if err := godotenv.Load(".default.env"); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	act := New(TestEnvironment(), "username", "password")

	equals(t, "username", act.Credentials.Username)
	equals(t, "password", act.Credentials.Password)
}

func TestNewWithTimeout(t *testing.T) {
	const timeout = time.Second * 50

	act := New(Testing, "un", "pw", WithTimeout(timeout))
	equals(t, timeout, act.client.Timeout)
}

// getInstance - instanciate Klarna for tests
func getInstance(env Environment) *Klarna {
	instance := New(
		env,
		os.Getenv("ADYEN_USERNAME"),
		os.Getenv("ADYEN_PASSWORD"))

	return instance
}

func equals(tb *testing.T, exp interface{}, act interface{}) {
	_, fullPath, line, _ := runtime.Caller(1)
	file := filepath.Base(fullPath)

	if !reflect.DeepEqual(exp, act) {
		fmt.Printf("%s:%d:\n\texp: %[3]v (%[3]T)\n\tgot: %[4]v (%[4]T)\n", file, line, exp, act)
		tb.FailNow()
	}
}

func assert(tb *testing.T, cond bool, message string) {
	_, fullPath, line, _ := runtime.Caller(1)
	file := filepath.Base(fullPath)

	if !cond {
		fmt.Printf("%s:%d:\n\t%s\n", file, line, message)
		tb.FailNow()
	}
}

func TestCreateURL(t *testing.T) {
	cases := []struct {
		name     string
		env      Environment
		resource string
		expected string
	}{
		{
			"payment session",
			TestEnvironment(),
			"/payments/v1/sessions",
			"https://api.playground.klarna.com/payments/v1/sessions",
		},
		{
			"production environment",
			ProductionEnvironment(),
			"/payments/v1/sessions",
			"https://api.klarna.com/payments/v1/sessions",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			instance := getInstance(c.env)
			act := instance.createURL(c.resource)
			equals(t, c.expected, act)
		})
	}
}
