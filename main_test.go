package main

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to Docker")
	//db, err := pool.RunWithOptions(&dockertest.RunOptions{
	//	Repository: "bitnami/mysql",
	//	Name: "mysql",
	//	Tag: "latest",
	//	Env: []string{
	//		"MYSQL_ROOT_PASSWORD=root123",
	//	},
	//	ExposedPorts: []string{"3306"},
	//	NetworkID: "app",
	//	Mounts:[]string{
	//		"mysql_data:/bitnami/mysql/data",
	//	},
	//})
	// _, err = pool.Run("bitnami/mysql", "latest", []string{"MYSQL_ROOT_PASSWORD=root123"})
	require.NoError(t, err, "could not start container")
	// db.GetPort("3306/tcp")
	//port := db.GetPort("3306/tcp")
	//fmt.Println(port)

	//s := fmt.Sprintf("DBPORT=%s", port)
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "docker-go-app",
		Tag:          "latest",
		Name:         "app",
		ExposedPorts: []string{"8080"},
	})
	// resource, err := pool.Run("goapp", "latest", []string{"DBUSER=root", "DBPASS=root123", s, "DBNAME=goapp", "DBHOST=127.0.0.1"})

	require.NoError(t, err, "could not start container")

	t.Cleanup(func() {
		require.NoError(t, pool.Purge(resource), "failed to remove container")
	})

	var resp *http.Response

	err = pool.Retry(func() error {
		resp, err = http.Get(fmt.Sprint("http://localhost:", resource.GetPort("8080/tcp"), "/ping"))
		if err != nil {
			t.Log("container not ready, waiting...")
			return err
		}
		return nil
	})
	require.NoError(t, err, "HTTP error")
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "failed to read HTTP body")

	// Finally, test the business requirement!
	require.JSONEq(t, `{"message":"pong"}`, string(body), "does not respond with valid JSON?")
}
