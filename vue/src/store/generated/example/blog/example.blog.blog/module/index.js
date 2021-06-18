// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdatePost } from "./types/blog/tx";
import { MsgCreateComment } from "./types/blog/tx";
import { MsgUpdateComment } from "./types/blog/tx";
import { MsgDeleteComment } from "./types/blog/tx";
import { MsgCreatePost } from "./types/blog/tx";
import { MsgDeletePost } from "./types/blog/tx";
const types = [
    ["/example.blog.blog.MsgUpdatePost", MsgUpdatePost],
    ["/example.blog.blog.MsgCreateComment", MsgCreateComment],
    ["/example.blog.blog.MsgUpdateComment", MsgUpdateComment],
    ["/example.blog.blog.MsgDeleteComment", MsgDeleteComment],
    ["/example.blog.blog.MsgCreatePost", MsgCreatePost],
    ["/example.blog.blog.MsgDeletePost", MsgDeletePost],
];
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw new Error("wallet is required");
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee = defaultFee, memo = null }) => memo ? client.signAndBroadcast(address, msgs, fee, memo) : client.signAndBroadcast(address, msgs, fee),
        msgUpdatePost: (data) => ({ typeUrl: "/example.blog.blog.MsgUpdatePost", value: data }),
        msgCreateComment: (data) => ({ typeUrl: "/example.blog.blog.MsgCreateComment", value: data }),
        msgUpdateComment: (data) => ({ typeUrl: "/example.blog.blog.MsgUpdateComment", value: data }),
        msgDeleteComment: (data) => ({ typeUrl: "/example.blog.blog.MsgDeleteComment", value: data }),
        msgCreatePost: (data) => ({ typeUrl: "/example.blog.blog.MsgCreatePost", value: data }),
        msgDeletePost: (data) => ({ typeUrl: "/example.blog.blog.MsgDeletePost", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
