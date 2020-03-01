// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: tensorflow_serving/apis/classification.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Tensorflow.Serving {

  /// <summary>Holder for reflection information generated from tensorflow_serving/apis/classification.proto</summary>
  public static partial class ClassificationReflection {

    #region Descriptor
    /// <summary>File descriptor for tensorflow_serving/apis/classification.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static ClassificationReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Cix0ZW5zb3JmbG93X3NlcnZpbmcvYXBpcy9jbGFzc2lmaWNhdGlvbi5wcm90",
            "bxISdGVuc29yZmxvdy5zZXJ2aW5nGiN0ZW5zb3JmbG93X3NlcnZpbmcvYXBp",
            "cy9pbnB1dC5wcm90bxojdGVuc29yZmxvd19zZXJ2aW5nL2FwaXMvbW9kZWwu",
            "cHJvdG8iJQoFQ2xhc3MSDQoFbGFiZWwYASABKAkSDQoFc2NvcmUYAiABKAIi",
            "PQoPQ2xhc3NpZmljYXRpb25zEioKB2NsYXNzZXMYASADKAsyGS50ZW5zb3Jm",
            "bG93LnNlcnZpbmcuQ2xhc3MiVAoUQ2xhc3NpZmljYXRpb25SZXN1bHQSPAoP",
            "Y2xhc3NpZmljYXRpb25zGAEgAygLMiMudGVuc29yZmxvdy5zZXJ2aW5nLkNs",
            "YXNzaWZpY2F0aW9ucyJ0ChVDbGFzc2lmaWNhdGlvblJlcXVlc3QSMQoKbW9k",
            "ZWxfc3BlYxgBIAEoCzIdLnRlbnNvcmZsb3cuc2VydmluZy5Nb2RlbFNwZWMS",
            "KAoFaW5wdXQYAiABKAsyGS50ZW5zb3JmbG93LnNlcnZpbmcuSW5wdXQihQEK",
            "FkNsYXNzaWZpY2F0aW9uUmVzcG9uc2USMQoKbW9kZWxfc3BlYxgCIAEoCzId",
            "LnRlbnNvcmZsb3cuc2VydmluZy5Nb2RlbFNwZWMSOAoGcmVzdWx0GAEgASgL",
            "MigudGVuc29yZmxvdy5zZXJ2aW5nLkNsYXNzaWZpY2F0aW9uUmVzdWx0QgP4",
            "AQFiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Tensorflow.Serving.InputReflection.Descriptor, global::Tensorflow.Serving.ModelReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Tensorflow.Serving.Class), global::Tensorflow.Serving.Class.Parser, new[]{ "Label", "Score" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Tensorflow.Serving.Classifications), global::Tensorflow.Serving.Classifications.Parser, new[]{ "Classes" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Tensorflow.Serving.ClassificationResult), global::Tensorflow.Serving.ClassificationResult.Parser, new[]{ "Classifications" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Tensorflow.Serving.ClassificationRequest), global::Tensorflow.Serving.ClassificationRequest.Parser, new[]{ "ModelSpec", "Input" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Tensorflow.Serving.ClassificationResponse), global::Tensorflow.Serving.ClassificationResponse.Parser, new[]{ "ModelSpec", "Result" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  /// <summary>
  /// A single class.
  /// </summary>
  public sealed partial class Class : pb::IMessage<Class> {
    private static readonly pb::MessageParser<Class> _parser = new pb::MessageParser<Class>(() => new Class());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<Class> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Tensorflow.Serving.ClassificationReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Class() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Class(Class other) : this() {
      label_ = other.label_;
      score_ = other.score_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Class Clone() {
      return new Class(this);
    }

    /// <summary>Field number for the "label" field.</summary>
    public const int LabelFieldNumber = 1;
    private string label_ = "";
    /// <summary>
    /// Label or name of the class.
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Label {
      get { return label_; }
      set {
        label_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "score" field.</summary>
    public const int ScoreFieldNumber = 2;
    private float score_;
    /// <summary>
    /// Score for this class (e.g., the probability the item belongs to this
    /// class). As per the proto3 default-value semantics, if the score is missing,
    /// it should be treated as 0.
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public float Score {
      get { return score_; }
      set {
        score_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as Class);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(Class other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Label != other.Label) return false;
      if (!pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.Equals(Score, other.Score)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Label.Length != 0) hash ^= Label.GetHashCode();
      if (Score != 0F) hash ^= pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.GetHashCode(Score);
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Label.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Label);
      }
      if (Score != 0F) {
        output.WriteRawTag(21);
        output.WriteFloat(Score);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Label.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Label);
      }
      if (Score != 0F) {
        size += 1 + 4;
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(Class other) {
      if (other == null) {
        return;
      }
      if (other.Label.Length != 0) {
        Label = other.Label;
      }
      if (other.Score != 0F) {
        Score = other.Score;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            Label = input.ReadString();
            break;
          }
          case 21: {
            Score = input.ReadFloat();
            break;
          }
        }
      }
    }

  }

  /// <summary>
  /// List of classes for a single item (tensorflow.Example).
  /// </summary>
  public sealed partial class Classifications : pb::IMessage<Classifications> {
    private static readonly pb::MessageParser<Classifications> _parser = new pb::MessageParser<Classifications>(() => new Classifications());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<Classifications> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Tensorflow.Serving.ClassificationReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Classifications() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Classifications(Classifications other) : this() {
      classes_ = other.classes_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Classifications Clone() {
      return new Classifications(this);
    }

    /// <summary>Field number for the "classes" field.</summary>
    public const int ClassesFieldNumber = 1;
    private static readonly pb::FieldCodec<global::Tensorflow.Serving.Class> _repeated_classes_codec
        = pb::FieldCodec.ForMessage(10, global::Tensorflow.Serving.Class.Parser);
    private readonly pbc::RepeatedField<global::Tensorflow.Serving.Class> classes_ = new pbc::RepeatedField<global::Tensorflow.Serving.Class>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::Tensorflow.Serving.Class> Classes {
      get { return classes_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as Classifications);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(Classifications other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if(!classes_.Equals(other.classes_)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= classes_.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      classes_.WriteTo(output, _repeated_classes_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      size += classes_.CalculateSize(_repeated_classes_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(Classifications other) {
      if (other == null) {
        return;
      }
      classes_.Add(other.classes_);
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            classes_.AddEntriesFrom(input, _repeated_classes_codec);
            break;
          }
        }
      }
    }

  }

  /// <summary>
  /// Contains one result per input example, in the same order as the input in
  /// ClassificationRequest.
  /// </summary>
  public sealed partial class ClassificationResult : pb::IMessage<ClassificationResult> {
    private static readonly pb::MessageParser<ClassificationResult> _parser = new pb::MessageParser<ClassificationResult>(() => new ClassificationResult());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<ClassificationResult> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Tensorflow.Serving.ClassificationReflection.Descriptor.MessageTypes[2]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ClassificationResult() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ClassificationResult(ClassificationResult other) : this() {
      classifications_ = other.classifications_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ClassificationResult Clone() {
      return new ClassificationResult(this);
    }

    /// <summary>Field number for the "classifications" field.</summary>
    public const int ClassificationsFieldNumber = 1;
    private static readonly pb::FieldCodec<global::Tensorflow.Serving.Classifications> _repeated_classifications_codec
        = pb::FieldCodec.ForMessage(10, global::Tensorflow.Serving.Classifications.Parser);
    private readonly pbc::RepeatedField<global::Tensorflow.Serving.Classifications> classifications_ = new pbc::RepeatedField<global::Tensorflow.Serving.Classifications>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::Tensorflow.Serving.Classifications> Classifications {
      get { return classifications_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as ClassificationResult);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(ClassificationResult other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if(!classifications_.Equals(other.classifications_)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= classifications_.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      classifications_.WriteTo(output, _repeated_classifications_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      size += classifications_.CalculateSize(_repeated_classifications_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(ClassificationResult other) {
      if (other == null) {
        return;
      }
      classifications_.Add(other.classifications_);
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            classifications_.AddEntriesFrom(input, _repeated_classifications_codec);
            break;
          }
        }
      }
    }

  }

  public sealed partial class ClassificationRequest : pb::IMessage<ClassificationRequest> {
    private static readonly pb::MessageParser<ClassificationRequest> _parser = new pb::MessageParser<ClassificationRequest>(() => new ClassificationRequest());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<ClassificationRequest> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Tensorflow.Serving.ClassificationReflection.Descriptor.MessageTypes[3]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ClassificationRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ClassificationRequest(ClassificationRequest other) : this() {
      modelSpec_ = other.modelSpec_ != null ? other.modelSpec_.Clone() : null;
      input_ = other.input_ != null ? other.input_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ClassificationRequest Clone() {
      return new ClassificationRequest(this);
    }

    /// <summary>Field number for the "model_spec" field.</summary>
    public const int ModelSpecFieldNumber = 1;
    private global::Tensorflow.Serving.ModelSpec modelSpec_;
    /// <summary>
    /// Model Specification. If version is not specified, will use the latest
    /// (numerical) version.
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Tensorflow.Serving.ModelSpec ModelSpec {
      get { return modelSpec_; }
      set {
        modelSpec_ = value;
      }
    }

    /// <summary>Field number for the "input" field.</summary>
    public const int InputFieldNumber = 2;
    private global::Tensorflow.Serving.Input input_;
    /// <summary>
    /// Input data.
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Tensorflow.Serving.Input Input {
      get { return input_; }
      set {
        input_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as ClassificationRequest);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(ClassificationRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(ModelSpec, other.ModelSpec)) return false;
      if (!object.Equals(Input, other.Input)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (modelSpec_ != null) hash ^= ModelSpec.GetHashCode();
      if (input_ != null) hash ^= Input.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (modelSpec_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(ModelSpec);
      }
      if (input_ != null) {
        output.WriteRawTag(18);
        output.WriteMessage(Input);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (modelSpec_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(ModelSpec);
      }
      if (input_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Input);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(ClassificationRequest other) {
      if (other == null) {
        return;
      }
      if (other.modelSpec_ != null) {
        if (modelSpec_ == null) {
          ModelSpec = new global::Tensorflow.Serving.ModelSpec();
        }
        ModelSpec.MergeFrom(other.ModelSpec);
      }
      if (other.input_ != null) {
        if (input_ == null) {
          Input = new global::Tensorflow.Serving.Input();
        }
        Input.MergeFrom(other.Input);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            if (modelSpec_ == null) {
              ModelSpec = new global::Tensorflow.Serving.ModelSpec();
            }
            input.ReadMessage(ModelSpec);
            break;
          }
          case 18: {
            if (input_ == null) {
              Input = new global::Tensorflow.Serving.Input();
            }
            input.ReadMessage(Input);
            break;
          }
        }
      }
    }

  }

  public sealed partial class ClassificationResponse : pb::IMessage<ClassificationResponse> {
    private static readonly pb::MessageParser<ClassificationResponse> _parser = new pb::MessageParser<ClassificationResponse>(() => new ClassificationResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<ClassificationResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Tensorflow.Serving.ClassificationReflection.Descriptor.MessageTypes[4]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ClassificationResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ClassificationResponse(ClassificationResponse other) : this() {
      modelSpec_ = other.modelSpec_ != null ? other.modelSpec_.Clone() : null;
      result_ = other.result_ != null ? other.result_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ClassificationResponse Clone() {
      return new ClassificationResponse(this);
    }

    /// <summary>Field number for the "model_spec" field.</summary>
    public const int ModelSpecFieldNumber = 2;
    private global::Tensorflow.Serving.ModelSpec modelSpec_;
    /// <summary>
    /// Effective Model Specification used for classification.
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Tensorflow.Serving.ModelSpec ModelSpec {
      get { return modelSpec_; }
      set {
        modelSpec_ = value;
      }
    }

    /// <summary>Field number for the "result" field.</summary>
    public const int ResultFieldNumber = 1;
    private global::Tensorflow.Serving.ClassificationResult result_;
    /// <summary>
    /// Result of the classification.
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Tensorflow.Serving.ClassificationResult Result {
      get { return result_; }
      set {
        result_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as ClassificationResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(ClassificationResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(ModelSpec, other.ModelSpec)) return false;
      if (!object.Equals(Result, other.Result)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (modelSpec_ != null) hash ^= ModelSpec.GetHashCode();
      if (result_ != null) hash ^= Result.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (result_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Result);
      }
      if (modelSpec_ != null) {
        output.WriteRawTag(18);
        output.WriteMessage(ModelSpec);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (modelSpec_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(ModelSpec);
      }
      if (result_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Result);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(ClassificationResponse other) {
      if (other == null) {
        return;
      }
      if (other.modelSpec_ != null) {
        if (modelSpec_ == null) {
          ModelSpec = new global::Tensorflow.Serving.ModelSpec();
        }
        ModelSpec.MergeFrom(other.ModelSpec);
      }
      if (other.result_ != null) {
        if (result_ == null) {
          Result = new global::Tensorflow.Serving.ClassificationResult();
        }
        Result.MergeFrom(other.Result);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            if (result_ == null) {
              Result = new global::Tensorflow.Serving.ClassificationResult();
            }
            input.ReadMessage(Result);
            break;
          }
          case 18: {
            if (modelSpec_ == null) {
              ModelSpec = new global::Tensorflow.Serving.ModelSpec();
            }
            input.ReadMessage(ModelSpec);
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code