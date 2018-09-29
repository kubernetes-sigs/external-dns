package egoscale

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	req := &CreateUser{}
	_ = req.response().(*User)
}

func TestRegisterUserKeys(t *testing.T) {
	req := &RegisterUserKeys{}
	_ = req.response().(*User)
}

func TestUpdateUser(t *testing.T) {
	req := &UpdateUser{}
	_ = req.response().(*User)
}

func TestListUsers(t *testing.T) {
	req := &ListUsers{}
	_ = req.response().(*ListUsersResponse)
}

func TestDeleteUser(t *testing.T) {
	req := &DeleteUser{}
	_ = req.response().(*booleanResponse)
}
