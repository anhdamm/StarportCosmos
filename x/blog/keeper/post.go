// x/blog/keeper/post.go
package keeper

import (
	"encoding/binary"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	blogTypes "github.com/example/blog/x/blog/types"
)

// GetPostCount get the total number of post
func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), blogTypes.KeyPrefix(blogTypes.PostCountKey))
	byteKey := blogTypes.KeyPrefix(blogTypes.PostCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

func (k Keeper) CreatePost(ctx sdk.Context, msg blogTypes.MsgCreatePost) uint64 {
	// Create the post
	count := k.GetPostCount(ctx)
	var post = blogTypes.Post{
		Creator: msg.Creator,
		Id:      count,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), blogTypes.KeyPrefix(blogTypes.PostKey))
	value := k.cdc.MustMarshalBinaryBare(&post)
	store.Set(GetPostIDBytes(post.Id), value)

	// Update post count
	k.SetPostCount(ctx, count+1)
	return count
}

// SetPostCount set the total number of post
func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), blogTypes.KeyPrefix(blogTypes.PostCountKey))
	byteKey := blogTypes.KeyPrefix(blogTypes.PostCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) AppendPost(
	ctx sdk.Context,
	creator string,
	title string,
	body string,
) uint64 {
	// Create the post
	count := k.GetPostCount(ctx)
	var post = blogTypes.Post{
		Creator: creator,
		Id:      count,
		Title:   title,
		Body:    body,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), blogTypes.KeyPrefix(blogTypes.PostKey))
	value := k.cdc.MustMarshalBinaryBare(&post)
	store.Set(GetPostIDBytes(post.Id), value)

	// Update post count
	k.SetPostCount(ctx, count+1)

	return count
}

func (k Keeper) GetPost(ctx sdk.Context, key uint64) blogTypes.Post {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), blogTypes.KeyPrefix(blogTypes.PostKey))
	var post blogTypes.Post
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetPostIDBytes(key)), &post)
	return post
}

func (k Keeper) SetPost(ctx sdk.Context, post blogTypes.Post) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), blogTypes.KeyPrefix(blogTypes.PostKey))
	b := k.cdc.MustMarshalBinaryBare(&post)
	store.Set(GetPostIDBytes(post.Id), b)
}

func (k Keeper) HasPost(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), blogTypes.KeyPrefix(blogTypes.PostKey))
	return store.Has(GetPostIDBytes(id))
}

func (k Keeper) RemovePost(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), blogTypes.KeyPrefix(blogTypes.PostKey))
	store.Delete(GetPostIDBytes(id))
}

func (k Keeper) GetPostOwner(ctx sdk.Context, key uint64) string {
	return k.GetPost(ctx, key).Creator
}

func (k Keeper) GetAllPost(ctx sdk.Context) (msgs []blogTypes.Post) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), blogTypes.KeyPrefix(blogTypes.PostKey))
	iterator := sdk.KVStorePrefixIterator(store, blogTypes.KeyPrefix(blogTypes.PostKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg blogTypes.Post
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}

func GetPostIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
