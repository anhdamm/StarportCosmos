// x/blog/handler_comment.go
package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/example/blog/x/blog/keeper"
	"github.com/example/blog/x/blog/types"
)

func handleMsgCreateComment(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateComment) (*sdk.Result, error) {
	k.AppendComment(ctx, msg.Id, msg.Creator, msg.Body, msg.PostID, msg.Time)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateComment(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateComment) (*sdk.Result, error) {
	var comment = types.Comment{
		Creator: msg.Creator,
		Id:      msg.Id,
		Body:    msg.Body,
		PostID:  msg.PostID,
	}

	if msg.Creator != k.GetCommentOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error
	}

	k.SetComment(ctx, comment)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteComment(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteComment) (*sdk.Result, error) {
	if msg.Creator != k.GetCommentOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")
	}

	k.RemoveComment(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
