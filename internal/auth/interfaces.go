package auth

type UserInterface interface {
	IsOwner(owner User) bool
	IsMod() bool
	isAdmin() bool

	CanUpdate(owner User)
	CanRead(owner User)
	CanDelete(owner User)
	CanMarkDelete(owner User)
}
