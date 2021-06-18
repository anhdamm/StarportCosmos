import { Reader, Writer } from "protobufjs/minimal";
import { CommentOnPost } from "../blog/post";
export declare const protobufPackage = "example.blog.blog";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateComment {
    creator: string;
    body: string;
    id: number;
    postID: number;
    time: string;
}
export interface MsgCreateCommentResponse {
    id: number;
}
export interface MsgUpdateComment {
    creator: string;
    id: number;
    body: string;
    postID: number;
}
export interface MsgUpdateCommentResponse {
}
export interface MsgDeleteComment {
    creator: string;
    id: number;
}
export interface MsgDeleteCommentResponse {
}
export interface MsgCreatePost {
    creator: string;
    title: string;
    body: string;
    comment: CommentOnPost[];
}
export interface MsgCreatePostResponse {
    id: number;
}
export interface MsgUpdatePost {
    creator: string;
    id: number;
    title: string;
    body: string;
}
export interface MsgUpdatePostResponse {
}
export interface MsgDeletePost {
    creator: string;
    id: number;
}
export interface MsgDeletePostResponse {
}
export declare const MsgCreateComment: {
    encode(message: MsgCreateComment, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateComment;
    fromJSON(object: any): MsgCreateComment;
    toJSON(message: MsgCreateComment): unknown;
    fromPartial(object: DeepPartial<MsgCreateComment>): MsgCreateComment;
};
export declare const MsgCreateCommentResponse: {
    encode(message: MsgCreateCommentResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateCommentResponse;
    fromJSON(object: any): MsgCreateCommentResponse;
    toJSON(message: MsgCreateCommentResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateCommentResponse>): MsgCreateCommentResponse;
};
export declare const MsgUpdateComment: {
    encode(message: MsgUpdateComment, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateComment;
    fromJSON(object: any): MsgUpdateComment;
    toJSON(message: MsgUpdateComment): unknown;
    fromPartial(object: DeepPartial<MsgUpdateComment>): MsgUpdateComment;
};
export declare const MsgUpdateCommentResponse: {
    encode(_: MsgUpdateCommentResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateCommentResponse;
    fromJSON(_: any): MsgUpdateCommentResponse;
    toJSON(_: MsgUpdateCommentResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateCommentResponse>): MsgUpdateCommentResponse;
};
export declare const MsgDeleteComment: {
    encode(message: MsgDeleteComment, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteComment;
    fromJSON(object: any): MsgDeleteComment;
    toJSON(message: MsgDeleteComment): unknown;
    fromPartial(object: DeepPartial<MsgDeleteComment>): MsgDeleteComment;
};
export declare const MsgDeleteCommentResponse: {
    encode(_: MsgDeleteCommentResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteCommentResponse;
    fromJSON(_: any): MsgDeleteCommentResponse;
    toJSON(_: MsgDeleteCommentResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteCommentResponse>): MsgDeleteCommentResponse;
};
export declare const MsgCreatePost: {
    encode(message: MsgCreatePost, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreatePost;
    fromJSON(object: any): MsgCreatePost;
    toJSON(message: MsgCreatePost): unknown;
    fromPartial(object: DeepPartial<MsgCreatePost>): MsgCreatePost;
};
export declare const MsgCreatePostResponse: {
    encode(message: MsgCreatePostResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreatePostResponse;
    fromJSON(object: any): MsgCreatePostResponse;
    toJSON(message: MsgCreatePostResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreatePostResponse>): MsgCreatePostResponse;
};
export declare const MsgUpdatePost: {
    encode(message: MsgUpdatePost, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdatePost;
    fromJSON(object: any): MsgUpdatePost;
    toJSON(message: MsgUpdatePost): unknown;
    fromPartial(object: DeepPartial<MsgUpdatePost>): MsgUpdatePost;
};
export declare const MsgUpdatePostResponse: {
    encode(_: MsgUpdatePostResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdatePostResponse;
    fromJSON(_: any): MsgUpdatePostResponse;
    toJSON(_: MsgUpdatePostResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdatePostResponse>): MsgUpdatePostResponse;
};
export declare const MsgDeletePost: {
    encode(message: MsgDeletePost, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeletePost;
    fromJSON(object: any): MsgDeletePost;
    toJSON(message: MsgDeletePost): unknown;
    fromPartial(object: DeepPartial<MsgDeletePost>): MsgDeletePost;
};
export declare const MsgDeletePostResponse: {
    encode(_: MsgDeletePostResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeletePostResponse;
    fromJSON(_: any): MsgDeletePostResponse;
    toJSON(_: MsgDeletePostResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeletePostResponse>): MsgDeletePostResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreateComment(request: MsgCreateComment): Promise<MsgCreateCommentResponse>;
    UpdateComment(request: MsgUpdateComment): Promise<MsgUpdateCommentResponse>;
    DeleteComment(request: MsgDeleteComment): Promise<MsgDeleteCommentResponse>;
    CreatePost(request: MsgCreatePost): Promise<MsgCreatePostResponse>;
    UpdatePost(request: MsgUpdatePost): Promise<MsgUpdatePostResponse>;
    DeletePost(request: MsgDeletePost): Promise<MsgDeletePostResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateComment(request: MsgCreateComment): Promise<MsgCreateCommentResponse>;
    UpdateComment(request: MsgUpdateComment): Promise<MsgUpdateCommentResponse>;
    DeleteComment(request: MsgDeleteComment): Promise<MsgDeleteCommentResponse>;
    CreatePost(request: MsgCreatePost): Promise<MsgCreatePostResponse>;
    UpdatePost(request: MsgUpdatePost): Promise<MsgUpdatePostResponse>;
    DeletePost(request: MsgDeletePost): Promise<MsgDeletePostResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
