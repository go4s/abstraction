package usecase

type (
    SessionUpdater interface {
        NamespaceUpdater
        LoginStatusUpdater
        TimeStrategyUpdater
        /* Optional
           ResourceIdUpdater
        */
    }
    SessionUpdaterEx interface {
        SessionUpdater
        ResourceIdUpdater
    }
    VerifyRequest interface {
        PhysicalAddressGetter
        TokenGetter
        CollectorNameGetter
        LeaseGetter
        StyleGetter
    }
    VerifyResponse interface {
        Update(SessionUpdater) error
    }
)
type (
    CoordinateRequest interface {
        PhysicalAddressGetter
        TokenGetter
        CollectorNameGetter
        ClusterRoleIdGetter
    }
    CoordinateResponse interface {
        Update(ClusterRoleUpdater) error
    }
)

type (
    RefreshRequest interface {
        PhysicalAddressGetter
        TokenGetter
        CollectorNameGetter
    }
    RefreshResponse interface {
        LoginStatusUpdater
    }
    RefreshRequestEx interface {
        RefreshRequest
        ClusterRoleIdGetter
    }
    RefreshResponseEx interface {
        RefreshResponse
        ClusterRoleUpdater
    }
)

type SessionAction = uint8

const (
    List = SessionAction(iota)
    Delete
    Disable
)

type (
    SessionOperation interface {
        SetAction(SessionAction)
    }
    SessionOperationModifier func(SessionOperation)
)
