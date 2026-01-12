package v1

func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.Fields != nil {
		in, out := &in.Fields, &out.Fields
		*out = make([]Field, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RowFilters != nil {
		in, out := &in.RowFilters, &out.RowFilters
		*out = new(RowFilters)
		(*in).DeepCopyInto(*out)
	}
}

func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

func (in *Field) DeepCopyInto(out *Field) {
	*out = *in
}

func (in *Field) DeepCopy() *Field {
	if in == nil {
		return nil
	}
	out := new(Field)
	in.DeepCopyInto(out)
	return out
}

func (in *RowFilters) DeepCopyInto(out *RowFilters) {
	*out = *in
	if in.Where != nil {
		in, out := &in.Where, &out.Where
		*out = make([]Where, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (in *RowFilters) DeepCopy() *RowFilters {
	if in == nil {
		return nil
	}
	out := new(RowFilters)
	in.DeepCopyInto(out)
	return out
}

func (in *Where) DeepCopyInto(out *Where) {
	*out = *in
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = make([]Member, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (in *Where) DeepCopy() *Where {
	if in == nil {
		return nil
	}
	out := new(Where)
	in.DeepCopyInto(out)
	return out
}

func (in *Member) DeepCopyInto(out *Member) {
	*out = *in
}

func (in *Member) DeepCopy() *Member {
	if in == nil {
		return nil
	}
	out := new(Member)
	in.DeepCopyInto(out)
	return out
}
