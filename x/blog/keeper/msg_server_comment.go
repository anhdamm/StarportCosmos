package keeper

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/gen2brain/beeep"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/example/blog/x/blog/types"
)

func (k msgServer) CreateComment(goCtx context.Context, msg *types.MsgCreateComment) (*types.MsgCreateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get variable to get conditions
	creator_post := k.GetPost(ctx, msg.PostID).Creator
	current_time := time.Now().Unix()
	lastBlockTime := ctx.BlockTime().Unix()

	// The commenting time has to be between
	// ctx.BlockTime()- 5seconds and ctx.BlockTime() + 5seconds
	if math.Abs(float64(lastBlockTime-current_time)) <= 5 && creator_post != msg.Creator {
		id := k.AppendComment(
			ctx,
			msg.Id,
			msg.Creator,
			msg.Body,
			msg.PostID,
			msg.Time,
		)

		return &types.MsgCreateCommentResponse{
			Id: id,
		}, nil

	} else {
		if creator_post == msg.Creator {
			err := beeep.Alert("Alert!", "You cannot comment yourself", "warning.png")
			if err != nil {
				panic(err)
			}
		} else {
			err := beeep.Alert("Alert!", "Commenting time must be within 5 seconds", "warning.png")
			if err != nil {
				panic(err)
			}

		}
		return nil, nil

	}

}

func (k msgServer) UpdateComment(goCtx context.Context, msg *types.MsgUpdateComment) (*types.MsgUpdateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var comment = types.Comment{
		Creator: msg.Creator,
		Id:      msg.Id,
		Body:    msg.Body,
		PostID:  msg.PostID,
	}

	// Checks that the element exists
	if !k.HasComment(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetCommentOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetComment(ctx, comment)

	return &types.MsgUpdateCommentResponse{}, nil
}

func (k msgServer) DeleteComment(goCtx context.Context, msg *types.MsgDeleteComment) (*types.MsgDeleteCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasComment(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetCommentOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveComment(ctx, msg.Id)

	return &types.MsgDeleteCommentResponse{}, nil
}
