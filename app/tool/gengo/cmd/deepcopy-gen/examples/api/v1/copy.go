// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	model "github.com/namelessup/bilibili/app/tool/gengo/cmd/deepcopy-gen/examples/model"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseInfoReply) DeepCopyInto(out *BaseInfoReply) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseInfoReply.
func (in *BaseInfoReply) DeepCopy() *BaseInfoReply {
	if in == nil {
		return nil
	}
	out := new(BaseInfoReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyAsIntoMemberBase is an autogenerated deepcopy function, copying the receiver, writing into model.MemberBase.
func (in *BaseInfoReply) DeepCopyAsIntoMemberBase(out *model.MemberBase) {
	out.Mid = in.Mid
	out.Name = in.Name
	out.Sex = in.Sex
	out.Face = in.Face
	out.Sign = in.Sign
	out.Rank = in.Rank
	out.Birthday = in.Birthday
	return
}

// DeepCopyFromMemberBase is an autogenerated deepcopy function, copying the receiver, writing into model.MemberBase.
func (out *BaseInfoReply) DeepCopyFromMemberBase(in *model.MemberBase) {
	out.Mid = in.Mid
	out.Name = in.Name
	out.Sex = in.Sex
	out.Face = in.Face
	out.Sign = in.Sign
	out.Rank = in.Rank
	out.Birthday = in.Birthday
	return
}

// DeepCopyAsMemberBase is an autogenerated deepcopy function, copying the receiver, creating a new model.MemberBase.
func (in *BaseInfoReply) DeepCopyAsMemberBase() *model.MemberBase {
	if in == nil {
		return nil
	}
	out := new(model.MemberBase)
	in.DeepCopyAsIntoMemberBase(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MidsReply) DeepCopyInto(out *MidsReply) {
	*out = *in
	if in.Mids != nil {
		in, out := &in.Mids, &out.Mids
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MidsReply.
func (in *MidsReply) DeepCopy() *MidsReply {
	if in == nil {
		return nil
	}
	out := new(MidsReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyAsIntoMids is an autogenerated deepcopy function, copying the receiver, writing into model.Mids.
func (in *MidsReply) DeepCopyAsIntoMids(out *model.Mids) {
	if in.Mids != nil {
		in, out := &in.Mids, &out.Mids
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopyFromMids is an autogenerated deepcopy function, copying the receiver, writing into model.Mids.
func (out *MidsReply) DeepCopyFromMids(in *model.Mids) {
	if in.Mids != nil {
		in, out := &in.Mids, &out.Mids
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopyAsMids is an autogenerated deepcopy function, copying the receiver, creating a new model.Mids.
func (in *MidsReply) DeepCopyAsMids() *model.Mids {
	if in == nil {
		return nil
	}
	out := new(model.Mids)
	in.DeepCopyAsIntoMids(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamesReply) DeepCopyInto(out *NamesReply) {
	*out = *in
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make(map[int64]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamesReply.
func (in *NamesReply) DeepCopy() *NamesReply {
	if in == nil {
		return nil
	}
	out := new(NamesReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyAsIntoNames is an autogenerated deepcopy function, copying the receiver, writing into model.Names.
func (in *NamesReply) DeepCopyAsIntoNames(out *model.Names) {
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make(map[int64]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopyFromNames is an autogenerated deepcopy function, copying the receiver, writing into model.Names.
func (out *NamesReply) DeepCopyFromNames(in *model.Names) {
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make(map[int64]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopyAsNames is an autogenerated deepcopy function, copying the receiver, creating a new model.Names.
func (in *NamesReply) DeepCopyAsNames() *model.Names {
	if in == nil {
		return nil
	}
	out := new(model.Names)
	in.DeepCopyAsIntoNames(out)
	return out
}
