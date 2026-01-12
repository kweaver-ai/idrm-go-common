package v1

func (in *Notification) DeepCopyInto(out *Notification) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

func (in *Notification) DeepCopy() *Notification {
	if in == nil {
		return nil
	}
	out := new(Notification)
	in.DeepCopyInto(out)
	return out
}

func (in *NotificationSpec) DeepCopyInto(out *NotificationSpec) {
	*out = *in
}

func (in *NotificationSpec) DeepCopy() *NotificationSpec {
	if in == nil {
		return nil
	}
	out := new(NotificationSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *NotificationStatus) DeepCopyInto(out *NotificationStatus) {
	*out = *in
}

func (in *NotificationStatus) DeepCopy() *NotificationStatus {
	if in == nil {
		return nil
	}
	out := new(NotificationStatus)
	in.DeepCopyInto(out)
	return out
}
