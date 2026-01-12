package v1

func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	in.CreatedAt.DeepCopyInto(&out.CreatedAt)
	in.UpdatedAt.DeepCopyInto(&out.UpdatedAt)
	if in.DeletedAt != nil {
		out.DeletedAt = in.DeletedAt.DeepCopy()
	}
}

func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}
