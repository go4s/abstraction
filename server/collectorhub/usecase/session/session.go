package session

import (
    "context"
    "github.com/go4s/abstraction/server/collectorhub/usecase"
)

type (
    Usecase interface {
        Verify(context.Context, usecase.VerifyRequest) (usecase.VerifyResponse, error)
        Coordinate(context.Context, usecase.CoordinateRequest) (usecase.CoordinateResponse, error)
        Refresh(context.Context, usecase.RefreshRequest) (usecase.RefreshResponse, error)
        Operation(context.Context, ...usecase.SessionOperationModifier)
        GetQualityRewriter(usecase.SessionRewriterGetter) usecase.PathRewriter
    }
)
