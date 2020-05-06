package repo

import (
	"testing"

	"github.com/google/uuid"
	"github.com/maximthomas/blazewall/auth-service/pkg/models"

	"github.com/stretchr/testify/assert"
)

var ur = UserLdapRepository{
	Address:        "localhost:50389",
	BindDN:         "cn=admin,dc=farawaygalaxy,dc=net",
	Password:       "passw0rd",
	BaseDN:         "ou=users,dc=farawaygalaxy,dc=net",
	ObjectClasses:  []string{"inetOrgPerson"},
	UserAttributes: []string{"sn", "cn"},
}

func TestLdapConnection(t *testing.T) {
	conn, err := ur.getConnection()
	assert.NoError(t, err)
	defer conn.Close()
}

func TestGetUser(t *testing.T) {
	user, exists := ur.GetUser("jerso")
	assert.True(t, exists)
	assert.Equal(t, "jerso", user.ID)

	_, exists2 := ur.GetUser("bad")
	assert.False(t, exists2)
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		user     string
		password string
		result   bool
	}{
		{"valid password", "jerso", "passw0rd", true},
		{"invalid password", "jerso", "bad", false},
		{"invalid user", "bad", "passw0rd", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ur.ValidatePassword(tt.user, tt.password)
			assert.Equal(t, tt.result, result)
		})
	}
}

func TestCreateUser(t *testing.T) {
	user := models.User{
		ID: "user1",
	}
	user, err := ur.CreateUser(user)
	assert.NoError(t, err)

	user, exists := ur.GetUser("jerso")
	assert.True(t, exists)
}

func TestSetPassword(t *testing.T) {
	var user = "jerso"
	newPassword := "newPassw0rd"

	err := ur.SetPassword(user, uuid.New().String())
	assert.NoError(t, err)

	result := ur.ValidatePassword(user, newPassword)
	assert.False(t, result)

	err = ur.SetPassword(user, newPassword)
	assert.NoError(t, err)

	result = ur.ValidatePassword(user, newPassword)
	assert.True(t, result)
}

func TestModifyUser(t *testing.T) {
	assert.Fail(t, "not implemented")
}