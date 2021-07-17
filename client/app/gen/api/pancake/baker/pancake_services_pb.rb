# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: pancake.proto for package 'pancake.baker'

require 'grpc'
require 'pancake_pb'

module Pancake
  module Baker
    module PancakeBakerService
      class Service

        include ::GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'pancake.baker.pancakeBakerService'

        # Bake は指定されたメニューのパンケーキを焼くメソッドです。
        # 焼かれたパンケーキをレスポンスとして返します。
        rpc :Bake, ::Pancake::Baker::BakeRequest, ::Pancake::Baker::BakeResponse
        # Report はメニューごとに焼いたパンケーキの数を返します。
        rpc :Report, ::Pancake::Baker::ReportRequest, ::Pancake::Baker::ReportResponse
      end

      Stub = Service.rpc_stub_class
    end
  end
end
