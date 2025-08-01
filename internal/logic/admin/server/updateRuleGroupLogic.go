package server

import (
	"context"
	"strings"

	"github.com/Js41313/Futuer-2/pkg/tool"

	"github.com/Js41313/Futuer-2/internal/model/server"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type UpdateRuleGroupLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUpdateRuleGroupLogic Update rule group
func NewUpdateRuleGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRuleGroupLogic {
	return &UpdateRuleGroupLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRuleGroupLogic) UpdateRuleGroup(req *types.UpdateRuleGroupRequest) error {
	rs, err := parseAndValidateRules(req.Rules, req.Name)
	if err != nil {
		return err
	}
	err = l.svcCtx.ServerModel.UpdateRuleGroup(l.ctx, &server.RuleGroup{
		Id:      req.Id,
		Icon:    req.Icon,
		Type:    req.Type,
		Name:    req.Name,
		Tags:    tool.StringSliceToString(req.Tags),
		Rules:   strings.Join(rs, "\n"),
		Default: req.Default,
		Enable:  req.Enable,
	})
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), err.Error())
	}
	if req.Default {
		if err = l.svcCtx.ServerModel.SetDefaultRuleGroup(l.ctx, req.Id); err != nil {
			l.Errorf("SetDefaultRuleGroup error: %v", err.Error())
			return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), err.Error())
		}
	}
	return nil
}
