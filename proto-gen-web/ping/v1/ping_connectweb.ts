// @generated by protoc-gen-connect-web v0.1.0 with parameter "target=ts"
// @generated from file ping/v1/ping.proto (package ping.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {PingRequest, PingResponse} from "./ping_pb.js";
import {MethodKind} from "@bufbuild/protobuf";

/**
 * @generated from service ping.v1.PingService
 */
export const PingService = {
  typeName: "ping.v1.PingService",
  methods: {
    /**
     * Ping sends a ping to the server to determine if it's reachable.
     *
     * @generated from rpc ping.v1.PingService.Ping
     */
    ping: {
      name: "Ping",
      I: PingRequest,
      O: PingResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

