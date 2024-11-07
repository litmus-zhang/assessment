package db

import (
	"context"
	"testing"

	"github.com/litmus-zhang/assessment/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	CreateTestUser(t)

}

func TestGetUser(t *testing.T) {
	user := CreateTestUser(t)

	u, err := testQueries.GetUserByID(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, u)
	require.Equal(t, user.ID, u.ID)
	require.Equal(t, user, u)

	u, err = testQueries.GetUserByEmail(context.Background(), user.Email)
	require.NoError(t, err)
	require.NotEmpty(t, u)
	require.Equal(t, user.Email, u.Email)
	require.Equal(t, user, u)

}

func TestUpdateUser(t *testing.T) {
	user := CreateTestUser(t)

	arg := UpdateUserParams{
		Email:     "new" + user.Email,
		ID:        user.ID,
		Password:  "new" + user.Password,
		FirstName: "new" + util.RandomUser(),
		LastName:  "new" + util.RandomUser(),
	}
	u, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, u)
	require.Equal(t, user.ID, u.ID)
	require.NotEqual(t, user, u)
}

func TestGetAllUser(t *testing.T) {
	n := 10
	for i := 0; i < n; i++ {
		CreateTestUser(t)
	}
	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}
	lists, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, lists, 5)

}

func TestDeleteUser(t *testing.T) {
	user := CreateTestUser(t)

	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)

	u, err := testQueries.GetUserByID(context.Background(), user.ID)
	require.Error(t, err)
	require.Empty(t, u)

}
