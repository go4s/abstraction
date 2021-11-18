package usecase

const (
    ServerTime = Strategy("2")
)
const (
    _ = ClusterRole(iota)
    Master
    Slave
)
const (
    _ = Style(iota)
    SupCollector
    Simulator
    GRPCGateway
    IoTGateway
    _
    OpcAE
)

type (
    Meta             struct{}
    Value            struct{}
    Style            = uint8
    ClusterRole      = uint8
    Strategy         = string
    NamespaceUpdater interface {
        SetNamespace(string)
    }
    LoginStatusUpdater interface {
        SetLoginStatus(uint64)
    }
    TimeStrategyUpdater interface {
        SetStrategy(Strategy)
    }
    ResourceIdUpdater interface {
        SetResourceId(string)
    }
    ClusterRoleUpdater interface {
        SetClusterRole(ClusterRole)
    }
    CollectorNameUpdater interface {
        SetCollectorName(string)
    }
    ClusterRoleGetter interface {
        GetClusterRole() ClusterRole
    }
    ClusterRoleIdGetter interface {
        GetClusterRoleId() string
    }
    NamespaceGetter interface {
        GetNamespace() string
    }
    CollectorNameGetter interface {
        GetCollectorName() string
    }
    TokenGetter interface {
        GetToken() string
    }
    PhysicalAddressGetter interface {
        GetMac() string
    }
    LeaseGetter interface {
        GetLease() uint64
    }
    StyleGetter interface {
        GetStyle() Style
    }
)
