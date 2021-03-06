// x/blog/types/messages_post.go
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePost{}

func NewMsgCreatePost(creator string, title string, body string) *MsgCreatePost {
	return &MsgCreatePost{
		Creator: creator,
		Title:   title,
		Body:    body,
	}
}

// Route ...
func (msg MsgCreatePost) Route() string {
	return RouterKey
}

// Type ...
func (msg MsgCreatePost) Type() string {
	return "CreatePost"
}

// GetSigners ...
func (msg *MsgCreatePost) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes ...
func (msg *MsgCreatePost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg *MsgCreatePost) ValidateBasic() error {
	if msg.Body == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Body can't be empty!")
	}

	if msg.Title == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Title can't be empty!")
	}

	_, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePost{}

func NewMsgUpdatePost(creator string, id uint64, title string, body string) *MsgUpdatePost {
	return &MsgUpdatePost{
		Id:      id,
		Creator: creator,
		Title:   title,
		Body:    body,
	}
}

func (msg *MsgUpdatePost) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePost) Type() string {
	return "UpdatePost"
}

func (msg *MsgUpdatePost) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePost{}

func NewMsgDeletePost(creator string, id uint64) *MsgDeletePost {
	return &MsgDeletePost{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeletePost) Route() string {
	return RouterKey
}

func (msg *MsgDeletePost) Type() string {
	return "DeletePost"
}

func (msg *MsgDeletePost) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
