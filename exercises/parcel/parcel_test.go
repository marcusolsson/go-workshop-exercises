package parcel

import "testing"

func TestDeliverParcel(t *testing.T) {
	p := parcel{
		delivered: false,
	}

	var res deliveryResult

	p.deliver(res)

	if got, want := res.status, "success"; got != want {
		t.Fatalf("expected delivery to be a success")
	}
	if got, want := p.delivered, true; got != want {
		t.Fatalf("expected parcel to be marked as delivered")
	}
}

func TestParcelAlreadyDelivered(t *testing.T) {
	p := parcel{
		delivered: true,
	}

	var res deliveryResult

	p.deliver(res)

	if got, want := res.status, "failed"; got != want {
		t.Fatalf("expected delivery to fail")
	}
	if got, want := p.delivered, true; got != want {
		t.Fatalf("expected parcel to be marked as delivered")
	}
}
