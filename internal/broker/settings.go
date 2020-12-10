package broker

const (
	_LOG_TAG string = "Broker:\t"

	// Events
	CreateOnePost            Intent = "CreateOnePost"
	CreateManyPosts          Intent = "CreateManyPosts"
	ConfirmOnePostCreation   Intent = "ConfirmOnePostCreation"
	ConfirmManyPostCreations Intent = "ConfirmManyPostCreations"

	ReadOnePost           Intent = "ReadOnePost"
	ReadManyPosts         Intent = "ReadManyPosts"
	ReadOnePostHistory    Intent = "ReadOnePostHistory"
	ReadManyPostHistories Intent = "ReadManyPostHistories"
	ReadUserPosts         Intent = "ReadUserPosts"

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
	AllIntents = []Intent{
		//CreateOnePost,
		//CreateManyPosts,
		ConfirmOnePostCreation,
		ConfirmManyPostCreations,
		ReadOnePost,
		ReadManyPosts,
		ReadOnePostHistory,
		ReadManyPostHistories,
		ReadUserPosts,
		//UpdateOnePost,
		//UpdateManyPosts,
		ConfirmOnePostUpdate,
		ConfirmManyPostUpdates,
		//DeleteOnePost,
		//DeleteManyPosts,
		ConfirmOnePostDeletion,
		ConfirmManyPostDeletions,
	}
)
