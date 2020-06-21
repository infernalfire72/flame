package constants

type AkatsukiPrivileges int

const (
	UserPublic AkatsukiPrivileges = 1 << iota
	UserNormal
	UserDonor
	AccessAdminPanel
	ManageUsers
	BanUsers
	SilenceUsers
	WipeUsers
	ManageBeatmaps
	ManageServers
	ManageSettings
	ManageBetaKeys
	ManageReports
	ManageDocs
	ManageBadges
	ViewAdminLogs
	ManagePrivileges
	SendAlerts
	ChatMod
	KickUsers
	UserPendingVerification
	TournamentStaff
	AdminCaker
	Premium
)

func (a AkatsukiPrivileges) Has(b AkatsukiPrivileges) bool {
	return b == 0 || (a & b) != 0
}

const FreeSupporter = true

func (a AkatsukiPrivileges) BanchoPrivileges() (result BanchoPrivileges) {
	if a.Has(UserNormal) {
		result |= Player
	}

	if a.Has(UserDonor) || a.Has(Premium) {
		result |= Supporter
	}

	if a.Has(AdminCaker) {
		result |= Supporter | Moderator
	}

	if a.Has(AdminCaker | ManagePrivileges) {
		result |= Developer
	}

	return
}

type BanchoPrivileges int

const (
	Player BanchoPrivileges = 1 << iota
	Moderator
	Supporter
	Owner
	Developer
	BanchoTournamentStaff
)

const None = 0