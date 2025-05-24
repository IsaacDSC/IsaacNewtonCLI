package entity

type CollaboratorPermissions string

func (p CollaboratorPermissions) String() string {
	return string(p)
}

const (
	CollaboratorPermissionsRead  CollaboratorPermissions = "read"
	CollaboratorPermissionsWrite CollaboratorPermissions = "write"
	CollaboratorPermissionsAdmin CollaboratorPermissions = "admin"
)
