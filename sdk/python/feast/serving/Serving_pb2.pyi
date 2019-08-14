# @generated by generate_proto_mypy_stubs.py.  Do not edit!
import sys
from feast.types.Value_pb2 import (
    Value as feast___types___Value_pb2___Value,
)

from google.protobuf.descriptor import (
    Descriptor as google___protobuf___descriptor___Descriptor,
)

from google.protobuf.internal.containers import (
    RepeatedScalarFieldContainer as google___protobuf___internal___containers___RepeatedScalarFieldContainer,
)

from google.protobuf.message import (
    Message as google___protobuf___message___Message,
)

from google.protobuf.timestamp_pb2 import (
    Timestamp as google___protobuf___timestamp_pb2___Timestamp,
)

from typing import (
    Iterable as typing___Iterable,
    Mapping as typing___Mapping,
    MutableMapping as typing___MutableMapping,
    Optional as typing___Optional,
    Text as typing___Text,
)

from typing_extensions import (
    Literal as typing_extensions___Literal,
)


class QueryFeaturesRequest(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    entityName = ... # type: typing___Text
    entityId = ... # type: google___protobuf___internal___containers___RepeatedScalarFieldContainer[typing___Text]
    featureId = ... # type: google___protobuf___internal___containers___RepeatedScalarFieldContainer[typing___Text]

    def __init__(self,
        *,
        entityName : typing___Optional[typing___Text] = None,
        entityId : typing___Optional[typing___Iterable[typing___Text]] = None,
        featureId : typing___Optional[typing___Iterable[typing___Text]] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> QueryFeaturesRequest: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def ClearField(self, field_name: typing_extensions___Literal[u"entityId",u"entityName",u"featureId"]) -> None: ...
    else:
        def ClearField(self, field_name: typing_extensions___Literal[u"entityId",b"entityId",u"entityName",b"entityName",u"featureId",b"featureId"]) -> None: ...

class QueryFeaturesResponse(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    class EntitiesEntry(google___protobuf___message___Message):
        DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
        key = ... # type: typing___Text

        @property
        def value(self) -> Entity: ...

        def __init__(self,
            *,
            key : typing___Optional[typing___Text] = None,
            value : typing___Optional[Entity] = None,
            ) -> None: ...
        @classmethod
        def FromString(cls, s: bytes) -> QueryFeaturesResponse.EntitiesEntry: ...
        def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
        def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
        if sys.version_info >= (3,):
            def HasField(self, field_name: typing_extensions___Literal[u"value"]) -> bool: ...
            def ClearField(self, field_name: typing_extensions___Literal[u"key",u"value"]) -> None: ...
        else:
            def HasField(self, field_name: typing_extensions___Literal[u"value",b"value"]) -> bool: ...
            def ClearField(self, field_name: typing_extensions___Literal[u"key",b"key",u"value",b"value"]) -> None: ...

    entityName = ... # type: typing___Text

    @property
    def entities(self) -> typing___MutableMapping[typing___Text, Entity]: ...

    def __init__(self,
        *,
        entityName : typing___Optional[typing___Text] = None,
        entities : typing___Optional[typing___Mapping[typing___Text, Entity]] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> QueryFeaturesResponse: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def ClearField(self, field_name: typing_extensions___Literal[u"entities",u"entityName"]) -> None: ...
    else:
        def ClearField(self, field_name: typing_extensions___Literal[u"entities",b"entities",u"entityName",b"entityName"]) -> None: ...

class Entity(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    class FeaturesEntry(google___protobuf___message___Message):
        DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
        key = ... # type: typing___Text

        @property
        def value(self) -> FeatureValue: ...

        def __init__(self,
            *,
            key : typing___Optional[typing___Text] = None,
            value : typing___Optional[FeatureValue] = None,
            ) -> None: ...
        @classmethod
        def FromString(cls, s: bytes) -> Entity.FeaturesEntry: ...
        def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
        def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
        if sys.version_info >= (3,):
            def HasField(self, field_name: typing_extensions___Literal[u"value"]) -> bool: ...
            def ClearField(self, field_name: typing_extensions___Literal[u"key",u"value"]) -> None: ...
        else:
            def HasField(self, field_name: typing_extensions___Literal[u"value",b"value"]) -> bool: ...
            def ClearField(self, field_name: typing_extensions___Literal[u"key",b"key",u"value",b"value"]) -> None: ...


    @property
    def features(self) -> typing___MutableMapping[typing___Text, FeatureValue]: ...

    def __init__(self,
        *,
        features : typing___Optional[typing___Mapping[typing___Text, FeatureValue]] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> Entity: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def ClearField(self, field_name: typing_extensions___Literal[u"features"]) -> None: ...
    else:
        def ClearField(self, field_name: typing_extensions___Literal[u"features",b"features"]) -> None: ...

class FeatureValue(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...

    @property
    def value(self) -> feast___types___Value_pb2___Value: ...

    @property
    def timestamp(self) -> google___protobuf___timestamp_pb2___Timestamp: ...

    def __init__(self,
        *,
        value : typing___Optional[feast___types___Value_pb2___Value] = None,
        timestamp : typing___Optional[google___protobuf___timestamp_pb2___Timestamp] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> FeatureValue: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def HasField(self, field_name: typing_extensions___Literal[u"timestamp",u"value"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"timestamp",u"value"]) -> None: ...
    else:
        def HasField(self, field_name: typing_extensions___Literal[u"timestamp",b"timestamp",u"value",b"value"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"timestamp",b"timestamp",u"value",b"value"]) -> None: ...