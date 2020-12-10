package broker

const (
	_LOG_TAG string = "Broker:\t"

	// Events
	CreateOnePost            Intent = "CreateOnePost"
	CreateManyPosts          Intent = "CreateManyPosts"
	ConfirmOnePostCreation   Intent = "ConfirmOnePostCreation"
	ConfirmManyPostCreations Intent = "ConfirmManyPostCreations"

	RequestOnePost           Intent = "RequestOnePost"
	RequestManyPosts         Intent = "RequestManyPosts"
	RequestOnePostHistory    Intent = "RequestOnePostHistory"
	RequestManyPostHistories Intent = "RequestManyPostHistories"
	RequestUserPosts         Intent = "RequestUserPosts"
	ReturnOnePost            Intent = "ReturnOnePost"
	ReturnManyPosts          Intent = "ReturnManyPosts"
	ReturnOnePostHistory     Intent = "ReturnOnePostHistory"
	ReturnManyPostHistories  Intent = "ReturnManyPostHistories"
	ReturnUserPosts          Intent = "ReturnUserPosts"

	UpdateOnePost   Intent = "UpdateOnePost"
	UpdateManyPosts Intent = "UpdateManyPosts"

	ConfirmOnePostUpdate   Intent = "ConfirmOnePostUpdate"
	ConfirmManyPostUpdates Intent = "ConfirmManyPostUpdates"

	DeleteOnePost   Intent = "DeleteOnePost"
	DeleteManyPosts Intent = "DeleteManyPosts"

	ConfirmOnePostDeletion   Intent = "ConfirmOnePostDeletion"
	ConfirmManyPostDeletions Intent = "ConfirmManyPostDeletions"
)

var (
	AllReturnIntents = []Intent{
		ConfirmOnePostCreation,
		ConfirmManyPostCreations,
		ReturnOnePost,
		ReturnManyPosts,
		ReturnOnePostHistory,
		ReturnManyPostHistories,
		ReturnUserPosts,
		ConfirmOnePostUpdate,
		ConfirmManyPostUpdates,
		ConfirmOnePostDeletion,
		ConfirmManyPostDeletions,
	}
)
