package main

import api "github.com/phasi/go-restapi"

// Define permissions
const (
	PermissionViewUsers   api.Permission = 1
	PermissionEditUsers   api.Permission = 2
	PermissionDeleteUsers api.Permission = 3
	PermissionAdmin       api.Permission = 10
)
