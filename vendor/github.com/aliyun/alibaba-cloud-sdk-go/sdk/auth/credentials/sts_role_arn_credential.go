package credentials

// Deprecated: Use RamRoleArnCredential in this package instead.
type StsRoleArnCredential struct {
	AccessKeyId           string
	AccessKeySecret       string
	RoleArn               string
	RoleSessionName       string
	RoleSessionExpiration int
}

type RamRoleArnCredential struct {
	AccessKeyId           string
	AccessKeySecret       string
	RoleArn               string
	RoleSessionName       string
	RoleSessionExpiration int
	Policy                string
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	StsRegion             string
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	StsRegion             string
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// Deprecated: Use RamRoleArnCredential in this package instead.
func NewStsRoleArnCredential(accessKeyId, accessKeySecret, roleArn, roleSessionName string, roleSessionExpiration int) *StsRoleArnCredential {
	return &StsRoleArnCredential{
		AccessKeyId:           accessKeyId,
		AccessKeySecret:       accessKeySecret,
		RoleArn:               roleArn,
		RoleSessionName:       roleSessionName,
		RoleSessionExpiration: roleSessionExpiration,
	}
}

func (oldCred *StsRoleArnCredential) ToRamRoleArnCredential() *RamRoleArnCredential {
	return &RamRoleArnCredential{
		AccessKeyId:           oldCred.AccessKeyId,
		AccessKeySecret:       oldCred.AccessKeySecret,
		RoleArn:               oldCred.RoleArn,
		RoleSessionName:       oldCred.RoleSessionName,
		RoleSessionExpiration: oldCred.RoleSessionExpiration,
	}
}

func NewRamRoleArnCredential(accessKeyId, accessKeySecret, roleArn, roleSessionName string, roleSessionExpiration int) *RamRoleArnCredential {
	return &RamRoleArnCredential{
		AccessKeyId:           accessKeyId,
		AccessKeySecret:       accessKeySecret,
		RoleArn:               roleArn,
		RoleSessionName:       roleSessionName,
		RoleSessionExpiration: roleSessionExpiration,
	}
}

func NewRamRoleArnWithPolicyCredential(accessKeyId, accessKeySecret, roleArn, roleSessionName, policy string, roleSessionExpiration int) *RamRoleArnCredential {
	return &RamRoleArnCredential{
		AccessKeyId:           accessKeyId,
		AccessKeySecret:       accessKeySecret,
		RoleArn:               roleArn,
		RoleSessionName:       roleSessionName,
		RoleSessionExpiration: roleSessionExpiration,
		Policy:                policy,
	}
}
