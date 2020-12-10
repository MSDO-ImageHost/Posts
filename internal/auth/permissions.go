package auth

const (
	USER_RANK  float64 = 00
	MOD_RANK   float64 = 10
	ADMIN_RANK float64 = 20
)

func (requester *User) IsOwner(owner User) bool {
	return requester.UserID == owner.UserID
}

func (requester *User) IsMod() bool {
	return requester.Rank == MOD_RANK
}

func (requester *User) IsAdmin() bool {
	return requester.Rank == ADMIN_RANK
}

func (requester *User) CanModify(owner User) bool {
	if requester.IsOwner(owner) || requester.IsMod() || requester.IsAdmin() {
		return true
	}
	return false
}

func (requester *User) CanDelete(owner User) bool {
	if requester.IsOwner(owner) || requester.IsMod() || requester.IsAdmin() {
		return true
	}
	return false
}
