package constants

type AkatsukiPrivileges int

const (
	UserPublic AkatsukiPrivileges = 1 << iota
	UserNormal
	UserDonor
	AdminAccessAdminPanel
	AdminManageUsers
	AdminBanUsers
	AdminSilenceUsers
	AdminWipeUsers
	AdminManageBeatmaps
	AdminManageServers
	AdminManageSettings
	AdminManageBetaKeys
	AdminManageReports
	AdminManageDocs
	AdminManageBadges
	AdminViewLogs
	AdminManagePrivileges
	AdminSendAlerts
	ChatMod
	AdminKickUsers
	UserPendingVerification
	TournamentStaff
	AdminCaker
	UserPremium
)

func (a AkatsukiPrivileges) Has(b AkatsukiPrivileges) bool {
	return b == 0 || (a&b) != 0
}

const FreeSupporter = true

func (a AkatsukiPrivileges) BanchoPrivileges() (result BanchoPrivileges) {
	if a.Has(UserNormal) {
		result |= Player
	}

	if a.Has(UserDonor) || a.Has(UserPremium) {
		result |= Supporter
	}

	if a.Has(AdminCaker) {
		result |= Supporter | Moderator
	}

	if a.Has(AdminCaker | AdminManagePrivileges) {
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
