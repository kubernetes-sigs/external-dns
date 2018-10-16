package egoscale

import (
	"strings"
	"testing"
)

func TestListAccounts(t *testing.T) {
	req := &ListAccounts{}
	_ = req.response().(*ListAccountsResponse)
}

func TestEnableAccount(t *testing.T) {
	req := &EnableAccount{}
	_ = req.response().(*Account)
}

func TestDisableAccount(t *testing.T) {
	req := &DisableAccount{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*Account)
}

func TestAccountTypeString(t *testing.T) {
	if UserAccount != AccountType(0) {
		t.Error("bad enum value", (int)(UserAccount), 0)
	}

	if UserAccount.String() != "UserAccount" {
		t.Error("mismatch", UserAccount, "UserAccount")
	}
	s := AccountType(45)

	if !strings.Contains(s.String(), "45") {
		t.Error("bad state", s.String())
	}
}
