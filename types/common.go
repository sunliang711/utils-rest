package types

const (
	RoleUser  int = 1
	RoleAdmin int = 2
)

const (
	Contract721  = 1
	Contract1155 = 2
)

const (
	RopstenNetwork = 1
	MainNetwork    = 2
)

const (
	ContractStatusDraft    = 1
	ContractStatusComplete = 2
)

const (
	// royalties level
	ContractLevel = 1
	TokenLevel    = 2
)

const (
	JwtKeyAddress = "address"
	JwtKeyRole    = "role"
)

const DefaultCacheControl = "public; max-age=86400"
