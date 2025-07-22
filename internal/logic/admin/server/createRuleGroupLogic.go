package server

import (
	"context"
	"strings"

	"github.com/Js41313/Futuer-2/pkg/rules"

	"github.com/Js41313/Futuer-2/internal/model/server"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/internal/types"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/xerr"
	"github.com/pkg/errors"
)

type CreateRuleGroupLogic struct {
	logger.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Create rule group
func NewCreateRuleGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRuleGroupLogic {
	return &CreateRuleGroupLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func parseAndValidateRules(ruleText, ruleName string) ([]string, error) {
	var rs []string
	ruleArr := strings.Split(ruleText, "\n")
	if len(ruleArr) == 0 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.InvalidParams), "rules is empty")
	}

	for _, s := range ruleArr {
		r := rules.NewRule(s, ruleName)
		if r == nil {
			continue
		}
		if err := r.Validate(); err != nil {
			continue
		}
		rs = append(rs, r.String())
	}
	return rs, nil
}
func (l *CreateRuleGroupLogic) CreateRuleGroup(req *types.CreateRuleGroupRequest) error {
	rs, err := parseAndValidateRules(req.Rules, req.Name)
	if err != nil {
		return err
	}
	info := &server.RuleGroup{
		Name:    req.Name,
		Icon:    req.Icon,
		Type:    req.Type,
		Tags:    tool.StringSliceToString(req.Tags),
		Rules:   strings.Join(rs, "\n"),
		Default: req.Default,
		Enable:  req.Enable,
	}
	err = l.svcCtx.ServerModel.InsertRuleGroup(l.ctx, info)
	if err != nil {
		l.Errorw("[CreateRuleGroup] Insert Database Error: ", logger.Field("error", err.Error()))
		return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseInsertError), "create server rule group error: %v", err)
	}
	if req.Default {
		if err = l.svcCtx.ServerModel.SetDefaultRuleGroup(l.ctx, info.Id); err != nil {
			l.Errorw("[CreateRuleGroup] Set Default Rule Group Error: ", logger.Field("error", err.Error()))
			return errors.Wrapf(xerr.NewErrCode(xerr.DatabaseUpdateError), "set default rule group error: %v", err)
		}
	}

	return nil
}
