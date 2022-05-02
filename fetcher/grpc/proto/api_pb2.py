# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/api.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/api.proto',
  package='billManagerAPI',
  syntax='proto3',
  serialized_options=b'Z%github.com/azuki774/bill-manager/grpc',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0fproto/api.proto\x12\x0e\x62illManagerAPI\x1a\x1bgoogle/protobuf/empty.proto\"G\n\x12onedayElectConsume\x12\x0f\n\x07\x64\x61ytime\x18\x01 \x01(\x02\x12\x11\n\tnighttime\x18\x02 \x01(\x02\x12\r\n\x05total\x18\x03 \x01(\x02\"K\n\x10\x65lectConsumeData\x12\x37\n\x0bpostRecords\x18\x01 \x03(\x0b\x32\".billManagerAPI.onedayElectConsume\"0\n\x04\x44\x61te\x12\x0c\n\x04year\x18\x01 \x01(\x05\x12\r\n\x05month\x18\x02 \x01(\x05\x12\x0b\n\x03\x64\x61y\x18\x03 \x01(\x05\x32\\\n\x10\x65lectConsumePost\x12H\n\x0c\x65lectConsume\x12 .billManagerAPI.electConsumeData\x1a\x16.google.protobuf.Empty2[\n\x0f\x65lectConsumeGet\x12H\n\x0c\x65lectConsume\x12\x14.billManagerAPI.Date\x1a\".billManagerAPI.onedayElectConsumeB\'Z%github.com/azuki774/bill-manager/grpcb\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,])




_ONEDAYELECTCONSUME = _descriptor.Descriptor(
  name='onedayElectConsume',
  full_name='billManagerAPI.onedayElectConsume',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='daytime', full_name='billManagerAPI.onedayElectConsume.daytime', index=0,
      number=1, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='nighttime', full_name='billManagerAPI.onedayElectConsume.nighttime', index=1,
      number=2, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='total', full_name='billManagerAPI.onedayElectConsume.total', index=2,
      number=3, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=64,
  serialized_end=135,
)


_ELECTCONSUMEDATA = _descriptor.Descriptor(
  name='electConsumeData',
  full_name='billManagerAPI.electConsumeData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='postRecords', full_name='billManagerAPI.electConsumeData.postRecords', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=137,
  serialized_end=212,
)


_DATE = _descriptor.Descriptor(
  name='Date',
  full_name='billManagerAPI.Date',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='year', full_name='billManagerAPI.Date.year', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='month', full_name='billManagerAPI.Date.month', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='day', full_name='billManagerAPI.Date.day', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=214,
  serialized_end=262,
)

_ELECTCONSUMEDATA.fields_by_name['postRecords'].message_type = _ONEDAYELECTCONSUME
DESCRIPTOR.message_types_by_name['onedayElectConsume'] = _ONEDAYELECTCONSUME
DESCRIPTOR.message_types_by_name['electConsumeData'] = _ELECTCONSUMEDATA
DESCRIPTOR.message_types_by_name['Date'] = _DATE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

onedayElectConsume = _reflection.GeneratedProtocolMessageType('onedayElectConsume', (_message.Message,), {
  'DESCRIPTOR' : _ONEDAYELECTCONSUME,
  '__module__' : 'proto.api_pb2'
  # @@protoc_insertion_point(class_scope:billManagerAPI.onedayElectConsume)
  })
_sym_db.RegisterMessage(onedayElectConsume)

electConsumeData = _reflection.GeneratedProtocolMessageType('electConsumeData', (_message.Message,), {
  'DESCRIPTOR' : _ELECTCONSUMEDATA,
  '__module__' : 'proto.api_pb2'
  # @@protoc_insertion_point(class_scope:billManagerAPI.electConsumeData)
  })
_sym_db.RegisterMessage(electConsumeData)

Date = _reflection.GeneratedProtocolMessageType('Date', (_message.Message,), {
  'DESCRIPTOR' : _DATE,
  '__module__' : 'proto.api_pb2'
  # @@protoc_insertion_point(class_scope:billManagerAPI.Date)
  })
_sym_db.RegisterMessage(Date)


DESCRIPTOR._options = None

_ELECTCONSUMEPOST = _descriptor.ServiceDescriptor(
  name='electConsumePost',
  full_name='billManagerAPI.electConsumePost',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=264,
  serialized_end=356,
  methods=[
  _descriptor.MethodDescriptor(
    name='electConsume',
    full_name='billManagerAPI.electConsumePost.electConsume',
    index=0,
    containing_service=None,
    input_type=_ELECTCONSUMEDATA,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_ELECTCONSUMEPOST)

DESCRIPTOR.services_by_name['electConsumePost'] = _ELECTCONSUMEPOST


_ELECTCONSUMEGET = _descriptor.ServiceDescriptor(
  name='electConsumeGet',
  full_name='billManagerAPI.electConsumeGet',
  file=DESCRIPTOR,
  index=1,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=358,
  serialized_end=449,
  methods=[
  _descriptor.MethodDescriptor(
    name='electConsume',
    full_name='billManagerAPI.electConsumeGet.electConsume',
    index=0,
    containing_service=None,
    input_type=_DATE,
    output_type=_ONEDAYELECTCONSUME,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_ELECTCONSUMEGET)

DESCRIPTOR.services_by_name['electConsumeGet'] = _ELECTCONSUMEGET

# @@protoc_insertion_point(module_scope)
