package usecase

import (
    "io"
)

type (
    QualityModifier interface {
        GetQuality() uint64
        SetQuality(uint64)
    }
    QualityRewriter interface {
        Update(QualityModifier) bool
    }
    WithQualityRewriter interface {
        SetQualityRewriter(...QualityRewriter)
    }
    PathRewriter interface {
        Update(PathModifier) bool
    }
    WithPathRewriter interface {
        SetPathRewriter(PathRewriter)
    }
    WithPublishOption interface {
        WithKey(string)
        WithHeader(string, []byte)
    }
    MetaFieldExtractor  func(*Meta) (string, interface{})
    ValueFieldExtractor func(*Value) (string, interface{})
    MetaReportOption    interface {
        WithPathRewriter
        WithPublishOption
        WithFieldsExtractor(...MetaFieldExtractor)
    }
    MetaReportOptionModifier func(MetaReportOption)
    ValueReportOption        interface {
        WithPathRewriter
        WithQualityRewriter
        WithPublishOption
        WithFieldsExtractor(...ValueFieldExtractor)
    }
    ValueReportOptionModifier func(ValueReportOption)
)
type SnapshotAction = uint8

const (
    ListCollector = SnapshotAction(iota)
    ListMeta
    ListValue
)

type (
    SnapshotOperation interface {
        SetAction(SnapshotAction)
        CollectorNameUpdater
        NamespaceUpdater
    }
    SnapshotOperationModifier func(SnapshotOperation)
)

type (
    PathRewriterGetter interface {
        NamespaceGetter
        CollectorNameGetter
    }
    SessionRewriterGetter interface {
        TokenGetter
        PhysicalAddressGetter
        CollectorNameGetter
    }
)
type (
    PathModifier interface {
        GetPath() string
        SetPath(string)
    }
    PathUnRewriter interface {
        Update(PathModifier) bool
    }
    PollRequest interface {
        NamespaceGetter
        CollectorNameGetter
    }
    PollRequestEx interface {
        PollRequest
        ClusterRoleGetter
    }
    PollResponse interface {
        PathRewrite(PathUnRewriter) error
        WriteTo(io.Writer) (int, error)
    }
)
