package services

type AppError struct {
	Error   error
	Message string
	Code    int
}

//
// Backup
//

// Create

type NewBackupException int

const (
	NewBackupOK NewBackupException = iota
	NewBackupForbidden
	NewBackupNotFound
	NewBackupMethodNotAllowed
	NewBackupInternalServerError
	NewBackupBadGateway
	NewBackupServiceUnavailable
	NewBackupGatewayTimeout
)

func (e NewBackupException) String() string {
	return [...]string{
		"NewBackupOK",
		"NewBackupForbidden",
		"NewBackupNotFound",
		"NewBackupMethodNotAllowed",
		"NewBackupInternalServerError",
		"NewBackupBadGateway",
		"NewBackupServiceUnavailable",
		"NewBackupGatewayTimeout",
	}[e]
}

func (e NewBackupException) EnumIndex() int {
	return int(e)
}

// Update

type UpdateBackupException int

const (
	UpdateBackupOK UpdateBackupException = iota
	UpdateBackupForbidden
	UpdateBackupNotFound
	UpdateBackupMethodNotAllowed
	UpdateBackupInternalServerError
	UpdateBackupBadGateway
	UpdateBackupServiceUnavailable
	UpdateBackupGatewayTimeout
)

func (e UpdateBackupException) String() string {
	return [...]string{
		"UpdateBackupOK",
		"UpdateBackupForbidden",
		"UpdateBackupNotFound",
		"UpdateBackupMethodNotAllowed",
		"UpdateBackupInternalServerError",
		"UpdateBackupBadGateway",
		"UpdateBackupServiceUnavailable",
		"UpdateBackupGatewayTimeout",
	}[e]
}

func (e UpdateBackupException) EnumIndex() int {
	return int(e)
}

// Retrieve

type GetBackupsException int

const (
	GetBackupsOK GetBackupsException = iota
	GetBackupsForbidden
	GetBackupsNotFound
	GetBackupsMethodNotAllowed
	GetBackupsInternalServerError
	GetBackupsServiceUnavailable
)

func (e GetBackupsException) String() string {
	return [...]string{
		"GetBackupsOK",
		"GetBackupsForbidden",
		"GetBackupsNotFound",
		"GetBackupsMethodNotAllowed",
		"GetBackupsInternalServerError",
		"GetBackupsServiceUnavailable",
	}[e]
}

func (e GetBackupsException) EnumIndex() int {
	return int(e)
}

// Retrieve by ID

type GetBackupByIdException int

const (
	GetBackupByIdOK UpdateBackupException = iota
	GetBackupByIdForbidden
	GetBackupByIdNotFound
	GetBackupByIdMethodNotAllowed
	GetBackupByIdInternalServerError
	GetBackupByIdServiceUnavailable
)

func (e GetBackupByIdException) String() string {
	return [...]string{
		"GetBackupByIdOK",
		"GetBackupByIdForbidden",
		"GetBackupByIdNotFound",
		"GetBackupByIdMethodNotAllowed",
		"GetBackupByIdInternalServerError",
		"GetBackupByIdServiceUnavailable",
	}[e]
}

func (e GetBackupByIdException) EnumIndex() int {
	return int(e)
}

//
// Restoration
//

// Create

type NewRestorationException int

const (
	NewRestorationOK NewRestorationException = iota
	NewRestorationForbidden
	NewRestorationNotFound
	NewRestorationMethodNotAllowed
	NewRestorationInternalServerError
	NewRestorationBadGateway
	NewRestorationServiceUnavailable
	NewRestorationGatewayTimeout
)

func (e NewRestorationException) String() string {
	return [...]string{
		"NewRestorationOK",
		"NewRestorationForbidden",
		"NewRestorationNotFound",
		"NewRestorationMethodNotAllowed",
		"NewRestorationInternalServerError",
		"NewRestorationBadGateway",
		"NewRestorationServiceUnavailable",
		"NewRestorationGatewayTimeout",
	}[e]
}

func (e NewRestorationException) EnumIndex() int {
	return int(e)
}

// Update

type UpdateRestorationException int

const (
	UpdateRestorationOK UpdateRestorationException = iota
	UpdateRestorationForbidden
	UpdateRestorationNotFound
	UpdateRestorationMethodNotAllowed
	UpdateRestorationInternalServerError
	UpdateRestorationBadGateway
	UpdateRestorationServiceUnavailable
	UpdateRestorationGatewayTimeout
)

func (e UpdateRestorationException) String() string {
	return [...]string{
		"UpdateRestorationOK",
		"UpdateRestorationForbidden",
		"UpdateRestorationNotFound",
		"UpdateRestorationMethodNotAllowed",
		"UpdateRestorationInternalServerError",
		"UpdateRestorationBadGateway",
		"UpdateRestorationServiceUnavailable",
		"UpdateRestorationGatewayTimeout",
	}[e]
}

func (e UpdateRestorationException) EnumIndex() int {
	return int(e)
}

// Retrieve

type GetRestorationsException int

const (
	GetRestorationsOK GetRestorationsException = iota
	GetRestorationsForbidden
	GetRestorationsNotFound
	GetRestorationsMethodNotAllowed
	GetRestorationsInternalServerError
	GetRestorationsServiceUnavailable
)

func (e GetRestorationsException) String() string {
	return [...]string{
		"GetRestorationsOK",
		"GetRestorationsForbidden",
		"GetRestorationsNotFound",
		"GetRestorationsMethodNotAllowed",
		"GetRestorationsInternalServerError",
		"GetRestorationsServiceUnavailable",
	}[e]
}

func (e GetRestorationsException) EnumIndex() int {
	return int(e)
}

// Retrieve by ID

type GetRestorationByIdException int

const (
	GetRestorationByIdOK GetRestorationByIdException = iota
	GetRestorationByIdForbidden
	GetRestorationByIdNotFound
	GetRestorationByIdMethodNotAllowed
	GetRestorationByIdInternalServerError
	GetRestorationByIdServiceUnavailable
)

func (e GetRestorationByIdException) String() string {
	return [...]string{
		"GetRestorationByIdOK",
		"GetRestorationByIdForbidden",
		"GetRestorationByIdNotFound",
		"GetRestorationByIdMethodNotAllowed",
		"GetRestorationByIdInternalServerError",
		"GetRestorationByIdServiceUnavailable",
	}[e]
}

func (e GetRestorationByIdException) EnumIndex() int {
	return int(e)
}

//
// Transfer
//

// Create

type NewTransferException int

const (
	NewTransferOK NewTransferException = iota
	NewTransferForbidden
	NewTransferNotFound
	NewTransferMethodNotAllowed
	NewTransferInternalServerError
	NewTransferBadGateway
	NewTransferServiceUnavailable
	NewTransferGatewayTimeout
)

func (e NewTransferException) String() string {
	return [...]string{
		"NewTransferOK",
		"NewTransferForbidden",
		"NewTransferNotFound",
		"NewTransferMethodNotAllowed",
		"NewTransferInternalServerError",
		"NewTransferBadGateway",
		"NewTransferServiceUnavailable",
		"NewTransferGatewayTimeout",
	}[e]
}

func (e NewTransferException) EnumIndex() int {
	return int(e)
}

// Update

type UpdateTransferException int

const (
	UpdateTransferOK UpdateTransferException = iota
	UpdateTransferForbidden
	UpdateTransferNotFound
	UpdateTransferMethodNotAllowed
	UpdateTransferInternalServerError
	UpdateTransferBadGateway
	UpdateTransferServiceUnavailable
	UpdateTransferGatewayTimeout
)

func (e UpdateTransferException) String() string {
	return [...]string{
		"UpdateTransferOK",
		"UpdateTransferForbidden",
		"UpdateTransferNotFound",
		"UpdateTransferMethodNotAllowed",
		"UpdateTransferInternalServerError",
		"UpdateTransferBadGateway",
		"UpdateTransferServiceUnavailable",
		"UpdateTransferGatewayTimeout",
	}[e]
}

func (e UpdateTransferException) EnumIndex() int {
	return int(e)
}

// Retrieve

type GetTransfersException int

const (
	GetTransfersOK GetTransfersException = iota
	GetTransfersForbidden
	GetTransfersNotFound
	GetTransfersMethodNotAllowed
	GetTransfersInternalServerError
	GetTransfersServiceUnavailable
)

func (e GetTransfersException) String() string {
	return [...]string{
		"GetTransfersOK",
		"GetTransfersForbidden",
		"GetTransfersNotFound",
		"GetTransfersMethodNotAllowed",
		"GetTransfersInternalServerError",
		"GetTransfesServiceUnavailable",
	}[e]
}

func (e GetTransfersException) EnumIndex() int {
	return int(e)
}

// Retrieve by ID

type GetTransferByIdException int

const (
	GetTransferByIdOK UpdateTransferException = iota
	GetTransferByIdForbidden
	GetTransferByIdNotFound
	GetTransferByIdMethodNotAllowed
	GetTransferByIdInternalServerError
	GetTransferByIdServiceUnavailable
)

func (e GetTransferByIdException) String() string {
	return [...]string{
		"GetTransferByIdOK",
		"GetTransferByIdForbidden",
		"GetTransferByIdNotFound",
		"GetTransferByIdMethodNotAllowed",
		"GetTransferByIdInternalServerError",
		"GetTransferByIdServiceUnavailable",
	}[e]
}

func (e GetTransferByIdException) EnumIndex() int {
	return int(e)
}

//
// Synchronization
//

// Create

type NewSynchronizationException int

const (
	NewSynchronizationOK NewSynchronizationException = iota
	NewSynchronizationForbidden
	NewSynchronizationNotFound
	NewSynchronizationMethodNotAllowed
	NewSynchronizationInternalServerError
	NewSynchronizationBadGateway
	NewSynchronizationServiceUnavailable
	NewSynchronizationGatewayTimeout
)

func (e NewSynchronizationException) String() string {
	return [...]string{
		"NewSynchronizationOK",
		"NewSynchronizationForbidden",
		"NewSynchronizationNotFound",
		"NewSynchronizationMethodNotAllowed",
		"NewSynchronizationInternalServerError",
		"NewSynchronizationBadGateway",
		"NewSynchronizationServiceUnavailable",
		"NewSynchronizationGatewayTimeout",
	}[e]
}

func (e NewSynchronizationException) EnumIndex() int {
	return int(e)
}

// Update

type UpdateSynchronizationException int

const (
	UpdateSynchronizationOK UpdateSynchronizationException = iota
	UpdateSynchronizationForbidden
	UpdateSynchronizationNotFound
	UpdateSynchronizationMethodNotAllowed
	UpdateSynchronizationInternalServerError
	UpdateSynchronizationBadGateway
	UpdateSynchronizationServiceUnavailable
	UpdateSynchronizationGatewayTimeout
)

func (e UpdateSynchronizationException) String() string {
	return [...]string{
		"UpdateSynchronizationOK",
		"UpdateSynchronizationForbidden",
		"UpdateSynchronizationNotFound",
		"UpdateSynchronizationMethodNotAllowed",
		"UpdateSynchronizationInternalServerError",
		"UpdateSynchronizationBadGateway",
		"UpdateSynchronizationServiceUnavailable",
		"UpdateSynchronizationGatewayTimeout",
	}[e]
}

func (e UpdateSynchronizationException) EnumIndex() int {
	return int(e)
}

// Retrieve

type GetSynchronizationsException int

const (
	GetSynchronizationsOK GetSynchronizationsException = iota
	GetSynchronizationsForbidden
	GetSynchronizationsNotFound
	GetSynchronizationsMethodNotAllowed
	GetSynchronizationsInternalServerError
	GetSynchronizationsServiceUnavailable
)

func (e GetSynchronizationsException) String() string {
	return [...]string{
		"GetSynchronizationsOK",
		"GetSynchronizationsForbidden",
		"GetSynchronizationsNotFound",
		"GetSynchronizationsMethodNotAllowed",
		"GetSynchronizationsInternalServerError",
		"GetTransfesServiceUnavailable",
	}[e]
}

func (e GetSynchronizationsException) EnumIndex() int {
	return int(e)
}

// Retrieve by ID

type GetSynchronizationByIdException int

const (
	GetSynchronizationByIdOK UpdateSynchronizationException = iota
	GetSynchronizationByIdForbidden
	GetSynchronizationByIdNotFound
	GetSynchronizationByIdMethodNotAllowed
	GetSynchronizationByIdInternalServerError
	GetSynchronizationByIdServiceUnavailable
)

func (e GetSynchronizationByIdException) String() string {
	return [...]string{
		"GetSynchronizationByIdOK",
		"GetSynchronizationByIdForbidden",
		"GetSynchronizationByIdNotFound",
		"GetSynchronizationByIdMethodNotAllowed",
		"GetSynchronizationByIdInternalServerError",
		"GetSynchronizationByIdServiceUnavailable",
	}[e]
}

func (e GetSynchronizationByIdException) EnumIndex() int {
	return int(e)
}

//
// Quota
//

// Create

type NewQuotaException int

const (
	NewQuotaOK NewQuotaException = iota
	NewQuotaForbidden
	NewQuotaNotFound
	NewQuotaMethodNotAllowed
	NewQuotaInternalServerError
	NewQuotaBadGateway
	NewQuotaServiceUnavailable
	NewQuotaGatewayTimeout
)

func (e NewQuotaException) String() string {
	return [...]string{
		"NewQuotaOK",
		"NewQuotaForbidden",
		"NewQuotaNotFound",
		"NewQuotaMethodNotAllowed",
		"NewQuotaInternalServerError",
		"NewQuotaBadGateway",
		"NewQuotaServiceUnavailable",
		"NewQuotaGatewayTimeout",
	}[e]
}

func (e NewQuotaException) EnumIndex() int {
	return int(e)
}

// Update

type UpdateQuotaException int

const (
	UpdateQuotaOK UpdateQuotaException = iota
	UpdateQuotaForbidden
	UpdateQuotaNotFound
	UpdateQuotaMethodNotAllowed
	UpdateQuotaInternalServerError
	UpdateQuotaBadGateway
	UpdateQuotaServiceUnavailable
	UpdateQuotaGatewayTimeout
)

func (e UpdateQuotaException) String() string {
	return [...]string{
		"UpdateQuotaOK",
		"UpdateQuotaForbidden",
		"UpdateQuotaNotFound",
		"UpdateQuotaMethodNotAllowed",
		"UpdateQuotaInternalServerError",
		"UpdateQuotaBadGateway",
		"UpdateQuotaServiceUnavailable",
		"UpdateQuotaGatewayTimeout",
	}[e]
}

func (e UpdateQuotaException) EnumIndex() int {
	return int(e)
}

// Retrieve

type GetQuotasException int

const (
	GetQuotasOK GetQuotasException = iota
	GetQuotasForbidden
	GetQuotasNotFound
	GetQuotasMethodNotAllowed
	GetQuotasInternalServerError
	GetQuotasServiceUnavailable
)

func (e GetQuotasException) String() string {
	return [...]string{
		"GetQuotasOK",
		"GetQuotasForbidden",
		"GetQuotasNotFound",
		"GetQuotasMethodNotAllowed",
		"GetQuotasInternalServerError",
		"GetTransfesServiceUnavailable",
	}[e]
}

func (e GetQuotasException) EnumIndex() int {
	return int(e)
}

// Retrieve by ID

type GetQuotaByIdException int

const (
	GetQuotaByIdOK UpdateQuotaException = iota
	GetQuotaByIdForbidden
	GetQuotaByIdNotFound
	GetQuotaByIdMethodNotAllowed
	GetQuotaByIdInternalServerError
	GetQuotaByIdServiceUnavailable
)

func (e GetQuotaByIdException) String() string {
	return [...]string{
		"GetQuotaByIdOK",
		"GetQuotaByIdForbidden",
		"GetQuotaByIdNotFound",
		"GetQuotaByIdMethodNotAllowed",
		"GetQuotaByIdInternalServerError",
		"GetQuotaByIdServiceUnavailable",
	}[e]
}

func (e GetQuotaByIdException) EnumIndex() int {
	return int(e)
}

//
// Log
//

// Create

type NewLogException int

const (
	NewLogOK NewLogException = iota
	NewLogForbidden
	NewLogNotFound
	NewLogMethodNotAllowed
	NewLogInternalServerError
	NewLogBadGateway
	NewLogServiceUnavailable
	NewLogGatewayTimeout
)

func (e NewLogException) String() string {
	return [...]string{
		"NewLogOK",
		"NewLogForbidden",
		"NewLogNotFound",
		"NewLogMethodNotAllowed",
		"NewLogInternalServerError",
		"NewLogBadGateway",
		"NewLogServiceUnavailable",
		"NewLogGatewayTimeout",
	}[e]
}

func (e NewLogException) EnumIndex() int {
	return int(e)
}

// Update

type UpdateLogException int

const (
	UpdateLogOK UpdateLogException = iota
	UpdateLogForbidden
	UpdateLogNotFound
	UpdateLogMethodNotAllowed
	UpdateLogInternalServerError
	UpdateLogBadGateway
	UpdateLogServiceUnavailable
	UpdateLogGatewayTimeout
)

func (e UpdateLogException) String() string {
	return [...]string{
		"UpdateLogOK",
		"UpdateLogForbidden",
		"UpdateLogNotFound",
		"UpdateLogMethodNotAllowed",
		"UpdateLogInternalServerError",
		"UpdateLogBadGateway",
		"UpdateLogServiceUnavailable",
		"UpdateLogGatewayTimeout",
	}[e]
}

func (e UpdateLogException) EnumIndex() int {
	return int(e)
}

// Retrieve

type GetLogsException int

const (
	GetLogsOK GetLogsException = iota
	GetLogsForbidden
	GetLogsNotFound
	GetLogsMethodNotAllowed
	GetLogsInternalServerError
	GetLogsServiceUnavailable
)

func (e GetLogsException) String() string {
	return [...]string{
		"GetLogsOK",
		"GetLogsForbidden",
		"GetLogsNotFound",
		"GetLogsMethodNotAllowed",
		"GetLogsInternalServerError",
		"GetTransfesServiceUnavailable",
	}[e]
}

func (e GetLogsException) EnumIndex() int {
	return int(e)
}

// Retrieve by ID

type GetLogByIdException int

const (
	GetLogByIdOK UpdateLogException = iota
	GetLogByIdForbidden
	GetLogByIdNotFound
	GetLogByIdMethodNotAllowed
	GetLogByIdInternalServerError
	GetLogByIdServiceUnavailable
)

func (e GetLogByIdException) String() string {
	return [...]string{
		"GetLogByIdOK",
		"GetLogByIdForbidden",
		"GetLogByIdNotFound",
		"GetLogByIdMethodNotAllowed",
		"GetLogByIdInternalServerError",
		"GetLogByIdServiceUnavailable",
	}[e]
}

func (e GetLogByIdException) EnumIndex() int {
	return int(e)
}

//
// Archive
//

// Create

type NewArchiveException int

const (
	NewArchiveOK NewArchiveException = iota
	NewArchiveForbidden
	NewArchiveNotFound
	NewArchiveMethodNotAllowed
	NewArchiveInternalServerError
	NewArchiveBadGateway
	NewArchiveServiceUnavailable
	NewArchiveGatewayTimeout
)

func (e NewArchiveException) String() string {
	return [...]string{
		"NewArchiveOK",
		"NewArchiveForbidden",
		"NewArchiveNotFound",
		"NewArchiveMethodNotAllowed",
		"NewArchiveInternalServerError",
		"NewArchiveBadGateway",
		"NewArchiveServiceUnavailable",
		"NewArchiveGatewayTimeout",
	}[e]
}

func (e NewArchiveException) EnumIndex() int {
	return int(e)
}

// Update

type UpdateArchiveException int

const (
	UpdateArchiveOK UpdateArchiveException = iota
	UpdateArchiveForbidden
	UpdateArchiveNotFound
	UpdateArchiveMethodNotAllowed
	UpdateArchiveInternalServerError
	UpdateArchiveBadGateway
	UpdateArchiveServiceUnavailable
	UpdateArchiveGatewayTimeout
)

func (e UpdateArchiveException) String() string {
	return [...]string{
		"UpdateArchiveOK",
		"UpdateArchiveForbidden",
		"UpdateArchiveNotFound",
		"UpdateArchiveMethodNotAllowed",
		"UpdateArchiveInternalServerError",
		"UpdateArchiveBadGateway",
		"UpdateArchiveServiceUnavailable",
		"UpdateArchiveGatewayTimeout",
	}[e]
}

func (e UpdateArchiveException) EnumIndex() int {
	return int(e)
}

// Retrieve

type GetArchivesException int

const (
	GetArchivesOK GetArchivesException = iota
	GetArchivesForbidden
	GetArchivesNotFound
	GetArchivesMethodNotAllowed
	GetArchivesInternalServerError
	GetArchivesServiceUnavailable
)

func (e GetArchivesException) String() string {
	return [...]string{
		"GetArchivesOK",
		"GetArchivesForbidden",
		"GetArchivesNotFound",
		"GetArchivesMethodNotAllowed",
		"GetArchivesInternalServerError",
		"GetTransfesServiceUnavailable",
	}[e]
}

func (e GetArchivesException) EnumIndex() int {
	return int(e)
}

// Retrieve by ID

type GetArchiveByIdException int

const (
	GetArchiveByIdOK UpdateArchiveException = iota
	GetArchiveByIdForbidden
	GetArchiveByIdNotFound
	GetArchiveByIdMethodNotAllowed
	GetArchiveByIdInternalServerError
	GetArchiveByIdServiceUnavailable
)

func (e GetArchiveByIdException) String() string {
	return [...]string{
		"GetArchiveByIdOK",
		"GetArchiveByIdForbidden",
		"GetArchiveByIdNotFound",
		"GetArchiveByIdMethodNotAllowed",
		"GetArchiveByIdInternalServerError",
		"GetArchiveByIdServiceUnavailable",
	}[e]
}

func (e GetArchiveByIdException) EnumIndex() int {
	return int(e)
}
