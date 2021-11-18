package backend

import (
    "context"
    "github.com/go4s/abstraction/server/collectorhub/usecase"
)

type (
    Usecase interface {
        Poll(context.Context, usecase.PollRequest) (usecase.PollResponse, error)
        GetQualityRewriter(usecase.PathRewriterGetter) usecase.PathRewriter
        GetPathRewriter(usecase.PathRewriterGetter) usecase.PathRewriter
        GetPathUnRewriter(usecase.PathRewriterGetter) usecase.PathUnRewriter
        ReportMeta(context.Context, []*usecase.Meta, ...usecase.MetaReportOptionModifier) error // tbd;
        ReportValue(context.Context, []*usecase.Value, ...usecase.ValueReportOptionModifier)    // tbd;
        Operation(context.Context, ...usecase.SnapshotOperationModifier)
    }
)
