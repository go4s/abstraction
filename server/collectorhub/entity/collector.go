package entity

import "github.com/go4s/abstraction/server/collectorhub/usecase"

type (
    Basic struct {
        Lease    uint64
        code     uint64
        strategy usecase.Strategy
        Kind     usecase.Style
        Mac      string
        Uuid     string
        Name     string
        tenant   string
    }
)

func (b *Basic) SetLoginStatus(u uint64)               { b.code = u }
func (b *Basic) SetNamespace(s string)                 { b.tenant = s }
func (b *Basic) SetStrategy(strategy usecase.Strategy) { b.strategy = strategy }

func (b Basic) GetCollectorName() string { return b.Name }
func (b Basic) GetLease() uint64         { return b.Lease }
func (b Basic) GetLoginStatus() uint64   { return b.code }
func (b Basic) GetMac() string           { return b.Mac }
func (b Basic) GetNamespace() string     { return b.tenant }
func (b Basic) GetStyle() usecase.Style  { return b.Kind }
func (b Basic) GetToken() string         { return b.Uuid }

var _ usecase.PhysicalAddressGetter = (*Basic)(nil)
var _ usecase.TokenGetter = (*Basic)(nil)
var _ usecase.CollectorNameGetter = (*Basic)(nil)
var _ usecase.LeaseGetter = (*Basic)(nil)
var _ usecase.StyleGetter = (*Basic)(nil)
var _ usecase.LoginStatusUpdater = (*Basic)(nil)
var _ usecase.NamespaceUpdater = (*Basic)(nil)
var _ usecase.NamespaceGetter = (*Basic)(nil)
var _ usecase.TimeStrategyUpdater = (*Basic)(nil)
