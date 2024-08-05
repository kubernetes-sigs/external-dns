package common

import (
	"log"
	"time"
)

type RoleArnCredential struct {
	roleArn         string
	roleSessionName string
	durationSeconds int64
	expiredTime     int64
	token           string
	tmpSecretId     string
	tmpSecretKey    string
<<<<<<< HEAD
	source          *RoleArnProvider
}

func (c *RoleArnCredential) GetSecretId() string {
	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretId
}

func (c *RoleArnCredential) GetSecretKey() string {
	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretKey

}

func (c *RoleArnCredential) GetToken() string {
	if c.needRefresh() {
		c.refresh()
	}
	return c.token
}

func (c *RoleArnCredential) needRefresh() bool {
	if c.tmpSecretKey == "" || c.tmpSecretId == "" || c.token == "" || c.expiredTime <= time.Now().Unix() {
		return true
	}
	return false
}

func (c *RoleArnCredential) refresh() {
	newCre, err := c.source.GetCredential()
	if err != nil {
		log.Println(err)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	source          Provider
}

func (c *RoleArnCredential) GetSecretId() string {
	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretId
}

func (c *RoleArnCredential) GetSecretKey() string {
	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretKey

}

func (c *RoleArnCredential) GetToken() string {
	if c.needRefresh() {
		c.refresh()
	}
	return c.token
}

func (c *RoleArnCredential) GetCredential() (string, string, string) {
	if c.needRefresh() {
		c.refresh()
	}
	return c.tmpSecretId, c.tmpSecretKey, c.token
}

func (c *RoleArnCredential) needRefresh() bool {
	if c.tmpSecretKey == "" || c.tmpSecretId == "" || c.token == "" || c.expiredTime <= time.Now().Unix() {
		return true
	}
	return false
}

func (c *RoleArnCredential) refresh() {
	newCre, err := c.source.GetCredential()
	if err != nil {
		log.Println(err)
		return
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}
	*c = *newCre.(*RoleArnCredential)
}
