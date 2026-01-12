package v1

func (in *WorkOrderTask) DeepCopyInto(out *WorkOrderTask) {
	*out = *in
	in.WorkOrderTaskTypedDetail.DeepCopyInto(&out.WorkOrderTaskTypedDetail)
}

func (in *WorkOrderTask) DeepCopy() (out *WorkOrderTask) {
	if in == nil {
		return
	}
	out = new(WorkOrderTask)
	in.DeepCopyInto(out)
	return
}

func (in *WorkOrderTaskTypedDetail) DeepCopyInto(out *WorkOrderTaskTypedDetail) {
	if in.DataAggregation != nil {
		in, out := &in.DataAggregation, &out.DataAggregation
		*out = make([]WorkOrderTaskDetailAggregationDetail, len(*in))
		copy(*out, *in)
	}

	// TODO: in.DataComprehension

	if in.DataFusion != nil {
		in, out := &in.DataFusion, &out.DataFusion
		*out = new(WorkOrderTaskDetailFusionDetail)
		**out = **in
	}

	// TODO: in.DataQuality

	if in.DataQualityAudit != nil {
		in, out := &in.DataQualityAudit, &out.DataQualityAudit
		*out = make([]*WorkOrderTaskDetailQualityAuditDetail, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto((*out)[i])
		}
	}
}

func (in *WorkOrderTaskTypedDetail) DeepCopy() (out *WorkOrderTaskTypedDetail) {
	if in == nil {
		return
	}
	out = new(WorkOrderTaskTypedDetail)
	in.DeepCopyInto(out)
	return
}

// 数据质量稽查工单的任务详情
func (in *WorkOrderTaskDetailQualityAuditDetail) DeepCopyInto(out *WorkOrderTaskDetailQualityAuditDetail) {
	*out = *in
}

func (in *WorkOrderTaskDetailQualityAuditDetail) DeepCopy() (out *WorkOrderTaskDetailQualityAuditDetail) {
	if in == nil {
		return
	}
	out = new(WorkOrderTaskDetailQualityAuditDetail)
	in.DeepCopyInto(out)
	return
}
