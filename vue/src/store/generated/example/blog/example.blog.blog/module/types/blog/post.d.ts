import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "example.blog.blog";
/** proto/blog/post.proto */
export interface CommentOnPost {
    creator: string;
    body: string;
    time: string;
}
export interface Post {
    creator: string;
    id: number;
    title: string;
    body: string;
    comments: CommentOnPost[];
}
export declare const CommentOnPost: {
    encode(message: CommentOnPost, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): CommentOnPost;
    fromJSON(object: any): CommentOnPost;
    toJSON(message: CommentOnPost): unknown;
    fromPartial(object: DeepPartial<CommentOnPost>): CommentOnPost;
};
export declare const Post: {
    encode(message: Post, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Post;
    fromJSON(object: any): Post;
    toJSON(message: Post): unknown;
    fromPartial(object: DeepPartial<Post>): Post;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
