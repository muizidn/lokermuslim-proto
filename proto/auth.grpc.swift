//
// DO NOT EDIT.
//
// Generated by the protocol buffer compiler.
// Source: auth.proto
//

//
// Copyright 2018, gRPC Authors All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
import Foundation
import GRPC
import NIO
import NIOHTTP1
import SwiftProtobuf


/// Usage: instantiate Loker_AuthServiceClient, then call methods of this protocol to make API calls.
public protocol Loker_AuthService {
  func login(_ request: Loker_LoginRequest, callOptions: CallOptions?) -> UnaryCall<Loker_LoginRequest, Loker_LoginResponse>
  func registerPekerja(_ request: Loker_PekerjaRegisterRequest, callOptions: CallOptions?) -> UnaryCall<Loker_PekerjaRegisterRequest, Loker_RegisterResponse>
  func registerPerusahaan(_ request: Loker_PerusahaanRegisterRequest, callOptions: CallOptions?) -> UnaryCall<Loker_PerusahaanRegisterRequest, Loker_RegisterResponse>
}

public final class Loker_AuthServiceClient: GRPCClient, Loker_AuthService {
  public let connection: ClientConnection
  public var defaultCallOptions: CallOptions

  /// Creates a client for the loker.Auth service.
  ///
  /// - Parameters:
  ///   - connection: `ClientConnection` to the service host.
  ///   - defaultCallOptions: Options to use for each service call if the user doesn't provide them.
  public init(connection: ClientConnection, defaultCallOptions: CallOptions = CallOptions()) {
    self.connection = connection
    self.defaultCallOptions = defaultCallOptions
  }

  /// Asynchronous unary call to Login.
  ///
  /// - Parameters:
  ///   - request: Request to send to Login.
  ///   - callOptions: Call options; `self.defaultCallOptions` is used if `nil`.
  /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
  public func login(_ request: Loker_LoginRequest, callOptions: CallOptions? = nil) -> UnaryCall<Loker_LoginRequest, Loker_LoginResponse> {
    return self.makeUnaryCall(path: "/loker.Auth/Login",
                              request: request,
                              callOptions: callOptions ?? self.defaultCallOptions)
  }

  /// Asynchronous unary call to RegisterPekerja.
  ///
  /// - Parameters:
  ///   - request: Request to send to RegisterPekerja.
  ///   - callOptions: Call options; `self.defaultCallOptions` is used if `nil`.
  /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
  public func registerPekerja(_ request: Loker_PekerjaRegisterRequest, callOptions: CallOptions? = nil) -> UnaryCall<Loker_PekerjaRegisterRequest, Loker_RegisterResponse> {
    return self.makeUnaryCall(path: "/loker.Auth/RegisterPekerja",
                              request: request,
                              callOptions: callOptions ?? self.defaultCallOptions)
  }

  /// Asynchronous unary call to RegisterPerusahaan.
  ///
  /// - Parameters:
  ///   - request: Request to send to RegisterPerusahaan.
  ///   - callOptions: Call options; `self.defaultCallOptions` is used if `nil`.
  /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
  public func registerPerusahaan(_ request: Loker_PerusahaanRegisterRequest, callOptions: CallOptions? = nil) -> UnaryCall<Loker_PerusahaanRegisterRequest, Loker_RegisterResponse> {
    return self.makeUnaryCall(path: "/loker.Auth/RegisterPerusahaan",
                              request: request,
                              callOptions: callOptions ?? self.defaultCallOptions)
  }

}

/// To build a server, implement a class that conforms to this protocol.
public protocol Loker_AuthProvider: CallHandlerProvider {
  func login(request: Loker_LoginRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Loker_LoginResponse>
  func registerPekerja(request: Loker_PekerjaRegisterRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Loker_RegisterResponse>
  func registerPerusahaan(request: Loker_PerusahaanRegisterRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Loker_RegisterResponse>
}

extension Loker_AuthProvider {
  public var serviceName: String { return "loker.Auth" }

  /// Determines, calls and returns the appropriate request handler, depending on the request's method.
  /// Returns nil for methods not handled by this service.
  public func handleMethod(_ methodName: String, callHandlerContext: CallHandlerContext) -> GRPCCallHandler? {
    switch methodName {
    case "Login":
      return UnaryCallHandler(callHandlerContext: callHandlerContext) { context in
        return { request in
          self.login(request: request, context: context)
        }
      }

    case "RegisterPekerja":
      return UnaryCallHandler(callHandlerContext: callHandlerContext) { context in
        return { request in
          self.registerPekerja(request: request, context: context)
        }
      }

    case "RegisterPerusahaan":
      return UnaryCallHandler(callHandlerContext: callHandlerContext) { context in
        return { request in
          self.registerPerusahaan(request: request, context: context)
        }
      }

    default: return nil
    }
  }
}

