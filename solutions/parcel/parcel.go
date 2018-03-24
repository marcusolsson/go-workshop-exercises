package parcel

type deliveryResult struct {
	status string
}

type parcel struct {
	delivered bool
}

func (p *parcel) deliver(res *deliveryResult) {
	if p.delivered {
		res.status = "failed"
		return
	}

	res.status = "success"
	p.delivered = true
}
