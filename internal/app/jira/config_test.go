package jira

import (
	"context"
	"errors"
	"testing"
)

func config() Configuration {
	return Configuration{
		Token:           "",
		Username:        "",
		Endpoint:        "",
		IssuePath:       "/rest/api/2/issue/",
		ProjectPath:     "/rest/api/2/project/",
		PriorityPath:    "/rest/api/2/priority/",
		UsersSearchPath: "/rest/api/2/users/search/",
	}
}

func assertErrorMessage(t testing.TB, error error) {
	t.Helper()

	if error != nil {
		t.Error(error)
		t.FailNow()
	}
}

func TestJira(t *testing.T) {
	var (
		projectKey = ""
		jiracfg    = config()
		ctx        = context.Background()

		isProjectKeyMissing = func(key string) {
			if key == "" {
				assertErrorMessage(t, errors.New("'ProjectKey' field is missing"))
			}
		}
	)

	assertErrorMessage(t, jiracfg.isRequiredFieldsEmpty())

	t.Run("GetProject", func(t *testing.T) {
		isProjectKeyMissing(projectKey)

		_, err := jiracfg.GetProject(ctx, projectKey)
		assertErrorMessage(t, err)
	})

	t.Run("GetUsers", func(t *testing.T) {
		_, err := jiracfg.GetUsers(ctx)
		assertErrorMessage(t, err)
	})

	t.Run("GetPriorities", func(t *testing.T) {
		_, err := jiracfg.GetPriorities(ctx)
		assertErrorMessage(t, err)
	})
}
