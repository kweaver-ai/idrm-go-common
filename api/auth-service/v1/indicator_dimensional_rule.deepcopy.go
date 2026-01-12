package v1

// TODO: use code generator

func (in *IndicatorDimensionalRuleSpec) DeepCopyInto(out *IndicatorDimensionalRuleSpec) {
	out.IndicatorID = in.IndicatorID
	in.Rule.DeepCopyInto(&out.Rule)
}

func (in *IndicatorDimensionalRuleSpec) DeepCopy() *IndicatorDimensionalRuleSpec {
	if in == nil {
		return nil
	}
	out := new(IndicatorDimensionalRuleSpec)
	in.DeepCopyInto(out)
	return out
}
