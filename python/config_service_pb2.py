# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: config-service.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='config-service.proto',
  package='config',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x14\x63onfig-service.proto\x12\x06\x63onfig\"\\\n\x11SetDesiredRequest\x12\x12\n\nidentifier\x18\x01 \x01(\t\x12\x11\n\tfieldName\x18\x02 \x01(\t\x12\x12\n\nfieldValue\x18\x03 \x01(\t\x12\x0c\n\x04slot\x18\x04 \x01(\x05\"b\n\x17\x43heckConsistencyRequest\x12\x11\n\tdeviceEUI\x18\x01 \x01(\t\x12\x0c\n\x04slot\x18\x02 \x01(\x05\x12\x12\n\nfieldIndex\x18\x03 \x01(\x05\x12\x12\n\nnumRetries\x18\x04 \x01(\x05\"`\n\x15UpdateReportedRequest\x12\x11\n\tdeviceEUI\x18\x01 \x01(\t\x12\x12\n\nfieldIndex\x18\x02 \x01(\x05\x12\x12\n\nfieldValue\x18\x03 \x01(\x0c\x12\x0c\n\x04slot\x18\x04 \x01(\x05\"\x19\n\x08Response\x12\r\n\x05reply\x18\x01 \x01(\t\"\xd0\x01\n\tConfigDoc\x12/\n\x07\x64\x65sired\x18\x01 \x03(\x0b\x32\x1e.config.ConfigDoc.DesiredEntry\x12\x31\n\x08reported\x18\x02 \x03(\x0b\x32\x1f.config.ConfigDoc.ReportedEntry\x1a.\n\x0c\x44\x65siredEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a/\n\rReportedEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"3\n\x0c\x43onfigFields\x12#\n\x06\x66ields\x18\x01 \x03(\x0b\x32\x13.config.ConfigField\"\xa0\x01\n\x0b\x43onfigField\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05index\x18\x02 \x01(\x05\x12\x0f\n\x07\x64\x65sired\x18\x03 \x01(\t\x12\x10\n\x08reported\x18\x04 \x01(\t\x12\x11\n\tfieldType\x18\x05 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x06 \x01(\t\x12\x0f\n\x07\x64\x65\x66\x61ult\x18\x07 \x01(\t\x12\x0b\n\x03min\x18\x08 \x01(\t\x12\x0b\n\x03max\x18\t \x01(\t\".\n\nIdentifier\x12\x12\n\nidentifier\x18\x01 \x01(\t\x12\x0c\n\x04slot\x18\x02 \x01(\x05\"M\n\x16GetConfigByNameRequest\x12\x12\n\nidentifier\x18\x01 \x01(\t\x12\x11\n\tfieldName\x18\x02 \x01(\t\x12\x0c\n\x04slot\x18\x03 \x01(\x05\"J\n\x17GetConfigByIndexRequest\x12\x12\n\nidentifier\x18\x01 \x01(\t\x12\r\n\x05index\x18\x02 \x01(\x05\x12\x0c\n\x04slot\x18\x03 \x01(\x05\")\n\x15UpdateFirmwareRequest\x12\x10\n\x08\x66irmware\x18\x01 \x01(\t2\xea\x04\n\rConfigService\x12;\n\nSetDesired\x12\x19.config.SetDesiredRequest\x1a\x10.config.Response\"\x00\x12G\n\x10\x43heckConsistency\x12\x1f.config.CheckConsistencyRequest\x1a\x10.config.Response\"\x00\x12\x43\n\x0eUpdateReported\x12\x1d.config.UpdateReportedRequest\x1a\x10.config.Response\"\x00\x12H\n\x0fGetConfigByName\x12\x1e.config.GetConfigByNameRequest\x1a\x13.config.ConfigField\"\x00\x12J\n\x10GetConfigByIndex\x12\x1f.config.GetConfigByIndexRequest\x1a\x13.config.ConfigField\"\x00\x12:\n\x0cGetAllConfig\x12\x12.config.Identifier\x1a\x14.config.ConfigFields\"\x00\x12:\n\x0fGetNewConfigDoc\x12\x12.config.Identifier\x1a\x11.config.ConfigDoc\"\x00\x12\x43\n\x0eUpdateFirmware\x12\x1d.config.UpdateFirmwareRequest\x1a\x10.config.Response\"\x00\x12;\n\x11\x41ssignRadioOffset\x12\x12.config.Identifier\x1a\x10.config.Response\"\x00\x62\x06proto3')
)




_SETDESIREDREQUEST = _descriptor.Descriptor(
  name='SetDesiredRequest',
  full_name='config.SetDesiredRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='identifier', full_name='config.SetDesiredRequest.identifier', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fieldName', full_name='config.SetDesiredRequest.fieldName', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fieldValue', full_name='config.SetDesiredRequest.fieldValue', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='slot', full_name='config.SetDesiredRequest.slot', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=32,
  serialized_end=124,
)


_CHECKCONSISTENCYREQUEST = _descriptor.Descriptor(
  name='CheckConsistencyRequest',
  full_name='config.CheckConsistencyRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='deviceEUI', full_name='config.CheckConsistencyRequest.deviceEUI', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='slot', full_name='config.CheckConsistencyRequest.slot', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fieldIndex', full_name='config.CheckConsistencyRequest.fieldIndex', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='numRetries', full_name='config.CheckConsistencyRequest.numRetries', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=126,
  serialized_end=224,
)


_UPDATEREPORTEDREQUEST = _descriptor.Descriptor(
  name='UpdateReportedRequest',
  full_name='config.UpdateReportedRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='deviceEUI', full_name='config.UpdateReportedRequest.deviceEUI', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fieldIndex', full_name='config.UpdateReportedRequest.fieldIndex', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fieldValue', full_name='config.UpdateReportedRequest.fieldValue', index=2,
      number=3, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='slot', full_name='config.UpdateReportedRequest.slot', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=226,
  serialized_end=322,
)


_RESPONSE = _descriptor.Descriptor(
  name='Response',
  full_name='config.Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='reply', full_name='config.Response.reply', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=324,
  serialized_end=349,
)


_CONFIGDOC_DESIREDENTRY = _descriptor.Descriptor(
  name='DesiredEntry',
  full_name='config.ConfigDoc.DesiredEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='config.ConfigDoc.DesiredEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='config.ConfigDoc.DesiredEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=465,
  serialized_end=511,
)

_CONFIGDOC_REPORTEDENTRY = _descriptor.Descriptor(
  name='ReportedEntry',
  full_name='config.ConfigDoc.ReportedEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='config.ConfigDoc.ReportedEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='config.ConfigDoc.ReportedEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=513,
  serialized_end=560,
)

_CONFIGDOC = _descriptor.Descriptor(
  name='ConfigDoc',
  full_name='config.ConfigDoc',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='desired', full_name='config.ConfigDoc.desired', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='reported', full_name='config.ConfigDoc.reported', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_CONFIGDOC_DESIREDENTRY, _CONFIGDOC_REPORTEDENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=352,
  serialized_end=560,
)


_CONFIGFIELDS = _descriptor.Descriptor(
  name='ConfigFields',
  full_name='config.ConfigFields',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='fields', full_name='config.ConfigFields.fields', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=562,
  serialized_end=613,
)


_CONFIGFIELD = _descriptor.Descriptor(
  name='ConfigField',
  full_name='config.ConfigField',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='config.ConfigField.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='index', full_name='config.ConfigField.index', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='desired', full_name='config.ConfigField.desired', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='reported', full_name='config.ConfigField.reported', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fieldType', full_name='config.ConfigField.fieldType', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='config.ConfigField.description', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default', full_name='config.ConfigField.default', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='min', full_name='config.ConfigField.min', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='max', full_name='config.ConfigField.max', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=616,
  serialized_end=776,
)


_IDENTIFIER = _descriptor.Descriptor(
  name='Identifier',
  full_name='config.Identifier',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='identifier', full_name='config.Identifier.identifier', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='slot', full_name='config.Identifier.slot', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=778,
  serialized_end=824,
)


_GETCONFIGBYNAMEREQUEST = _descriptor.Descriptor(
  name='GetConfigByNameRequest',
  full_name='config.GetConfigByNameRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='identifier', full_name='config.GetConfigByNameRequest.identifier', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fieldName', full_name='config.GetConfigByNameRequest.fieldName', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='slot', full_name='config.GetConfigByNameRequest.slot', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=826,
  serialized_end=903,
)


_GETCONFIGBYINDEXREQUEST = _descriptor.Descriptor(
  name='GetConfigByIndexRequest',
  full_name='config.GetConfigByIndexRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='identifier', full_name='config.GetConfigByIndexRequest.identifier', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='index', full_name='config.GetConfigByIndexRequest.index', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='slot', full_name='config.GetConfigByIndexRequest.slot', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=905,
  serialized_end=979,
)


_UPDATEFIRMWAREREQUEST = _descriptor.Descriptor(
  name='UpdateFirmwareRequest',
  full_name='config.UpdateFirmwareRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='firmware', full_name='config.UpdateFirmwareRequest.firmware', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=981,
  serialized_end=1022,
)

_CONFIGDOC_DESIREDENTRY.containing_type = _CONFIGDOC
_CONFIGDOC_REPORTEDENTRY.containing_type = _CONFIGDOC
_CONFIGDOC.fields_by_name['desired'].message_type = _CONFIGDOC_DESIREDENTRY
_CONFIGDOC.fields_by_name['reported'].message_type = _CONFIGDOC_REPORTEDENTRY
_CONFIGFIELDS.fields_by_name['fields'].message_type = _CONFIGFIELD
DESCRIPTOR.message_types_by_name['SetDesiredRequest'] = _SETDESIREDREQUEST
DESCRIPTOR.message_types_by_name['CheckConsistencyRequest'] = _CHECKCONSISTENCYREQUEST
DESCRIPTOR.message_types_by_name['UpdateReportedRequest'] = _UPDATEREPORTEDREQUEST
DESCRIPTOR.message_types_by_name['Response'] = _RESPONSE
DESCRIPTOR.message_types_by_name['ConfigDoc'] = _CONFIGDOC
DESCRIPTOR.message_types_by_name['ConfigFields'] = _CONFIGFIELDS
DESCRIPTOR.message_types_by_name['ConfigField'] = _CONFIGFIELD
DESCRIPTOR.message_types_by_name['Identifier'] = _IDENTIFIER
DESCRIPTOR.message_types_by_name['GetConfigByNameRequest'] = _GETCONFIGBYNAMEREQUEST
DESCRIPTOR.message_types_by_name['GetConfigByIndexRequest'] = _GETCONFIGBYINDEXREQUEST
DESCRIPTOR.message_types_by_name['UpdateFirmwareRequest'] = _UPDATEFIRMWAREREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SetDesiredRequest = _reflection.GeneratedProtocolMessageType('SetDesiredRequest', (_message.Message,), {
  'DESCRIPTOR' : _SETDESIREDREQUEST,
  '__module__' : 'config_service_pb2'
  # @@protoc_insertion_point(class_scope:config.SetDesiredRequest)
  })
_sym_db.RegisterMessage(SetDesiredRequest)

CheckConsistencyRequest = _reflection.GeneratedProtocolMessageType('CheckConsistencyRequest', (_message.Message,), {
  'DESCRIPTOR' : _CHECKCONSISTENCYREQUEST,
  '__module__' : 'config_service_pb2'
  # @@protoc_insertion_point(class_scope:config.CheckConsistencyRequest)
  })
_sym_db.RegisterMessage(CheckConsistencyRequest)

UpdateReportedRequest = _reflection.GeneratedProtocolMessageType('UpdateReportedRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEREPORTEDREQUEST,
  '__module__' : 'config_service_pb2'
  # @@protoc_insertion_point(class_scope:config.UpdateReportedRequest)
  })
_sym_db.RegisterMessage(UpdateReportedRequest)

Response = _reflection.GeneratedProtocolMessageType('Response', (_message.Message,), {
  'DESCRIPTOR' : _RESPONSE,
  '__module__' : 'config_service_pb2'
  # @@protoc_insertion_point(class_scope:config.Response)
  })
_sym_db.RegisterMessage(Response)

ConfigDoc = _reflection.GeneratedProtocolMessageType('ConfigDoc', (_message.Message,), {

  'DesiredEntry' : _reflection.GeneratedProtocolMessageType('DesiredEntry', (_message.Message,), {
    'DESCRIPTOR' : _CONFIGDOC_DESIREDENTRY,
    '__module__' : 'config_service_pb2'
    # @@protoc_insertion_point(class_scope:config.ConfigDoc.DesiredEntry)
    })
  ,

  'ReportedEntry' : _reflection.GeneratedProtocolMessageType('ReportedEntry', (_message.Message,), {
    'DESCRIPTOR' : _CONFIGDOC_REPORTEDENTRY,
    '__module__' : 'config_service_pb2'
    # @@protoc_insertion_point(class_scope:config.ConfigDoc.ReportedEntry)
    })
  ,
  'DESCRIPTOR' : _CONFIGDOC,
  '__module__' : 'config_service_pb2'
  # @@protoc_insertion_point(class_scope:config.ConfigDoc)
  })
_sym_db.RegisterMessage(ConfigDoc)
_sym_db.RegisterMessage(ConfigDoc.DesiredEntry)
_sym_db.RegisterMessage(ConfigDoc.ReportedEntry)

ConfigFields = _reflection.GeneratedProtocolMessageType('ConfigFields', (_message.Message,), {
  'DESCRIPTOR' : _CONFIGFIELDS,
  '__module__' : 'config_service_pb2'
  # @@protoc_insertion_point(class_scope:config.ConfigFields)
  })
_sym_db.RegisterMessage(ConfigFields)

ConfigField = _reflection.GeneratedProtocolMessageType('ConfigField', (_message.Message,), {
  'DESCRIPTOR' : _CONFIGFIELD,
  '__module__' : 'config_service_pb2'
  # @@protoc_insertion_point(class_scope:config.ConfigField)
  })
_sym_db.RegisterMessage(ConfigField)

Identifier = _reflection.GeneratedProtocolMessageType('Identifier', (_message.Message,), {
  'DESCRIPTOR' : _IDENTIFIER,
  '__module__' : 'config_service_pb2'
  # @@protoc_insertion_point(class_scope:config.Identifier)
  })
_sym_db.RegisterMessage(Identifier)

GetConfigByNameRequest = _reflection.GeneratedProtocolMessageType('GetConfigByNameRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETCONFIGBYNAMEREQUEST,
  '__module__' : 'config_service_pb2'
  # @@protoc_insertion_point(class_scope:config.GetConfigByNameRequest)
  })
_sym_db.RegisterMessage(GetConfigByNameRequest)

GetConfigByIndexRequest = _reflection.GeneratedProtocolMessageType('GetConfigByIndexRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETCONFIGBYINDEXREQUEST,
  '__module__' : 'config_service_pb2'
  # @@protoc_insertion_point(class_scope:config.GetConfigByIndexRequest)
  })
_sym_db.RegisterMessage(GetConfigByIndexRequest)

UpdateFirmwareRequest = _reflection.GeneratedProtocolMessageType('UpdateFirmwareRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEFIRMWAREREQUEST,
  '__module__' : 'config_service_pb2'
  # @@protoc_insertion_point(class_scope:config.UpdateFirmwareRequest)
  })
_sym_db.RegisterMessage(UpdateFirmwareRequest)


_CONFIGDOC_DESIREDENTRY._options = None
_CONFIGDOC_REPORTEDENTRY._options = None

_CONFIGSERVICE = _descriptor.ServiceDescriptor(
  name='ConfigService',
  full_name='config.ConfigService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1025,
  serialized_end=1643,
  methods=[
  _descriptor.MethodDescriptor(
    name='SetDesired',
    full_name='config.ConfigService.SetDesired',
    index=0,
    containing_service=None,
    input_type=_SETDESIREDREQUEST,
    output_type=_RESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='CheckConsistency',
    full_name='config.ConfigService.CheckConsistency',
    index=1,
    containing_service=None,
    input_type=_CHECKCONSISTENCYREQUEST,
    output_type=_RESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='UpdateReported',
    full_name='config.ConfigService.UpdateReported',
    index=2,
    containing_service=None,
    input_type=_UPDATEREPORTEDREQUEST,
    output_type=_RESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetConfigByName',
    full_name='config.ConfigService.GetConfigByName',
    index=3,
    containing_service=None,
    input_type=_GETCONFIGBYNAMEREQUEST,
    output_type=_CONFIGFIELD,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetConfigByIndex',
    full_name='config.ConfigService.GetConfigByIndex',
    index=4,
    containing_service=None,
    input_type=_GETCONFIGBYINDEXREQUEST,
    output_type=_CONFIGFIELD,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetAllConfig',
    full_name='config.ConfigService.GetAllConfig',
    index=5,
    containing_service=None,
    input_type=_IDENTIFIER,
    output_type=_CONFIGFIELDS,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetNewConfigDoc',
    full_name='config.ConfigService.GetNewConfigDoc',
    index=6,
    containing_service=None,
    input_type=_IDENTIFIER,
    output_type=_CONFIGDOC,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='UpdateFirmware',
    full_name='config.ConfigService.UpdateFirmware',
    index=7,
    containing_service=None,
    input_type=_UPDATEFIRMWAREREQUEST,
    output_type=_RESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='AssignRadioOffset',
    full_name='config.ConfigService.AssignRadioOffset',
    index=8,
    containing_service=None,
    input_type=_IDENTIFIER,
    output_type=_RESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_CONFIGSERVICE)

DESCRIPTOR.services_by_name['ConfigService'] = _CONFIGSERVICE

# @@protoc_insertion_point(module_scope)
