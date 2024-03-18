package xctx

import (
	"context"
	"reflect"
	"testing"
)

func TestCtxGet(t *testing.T) {
	ctx := context.Background()

	k0 := "k0"
	k1 := "k1"
	wantInt := 1
	defaultInt := 0
	defaultStr := ""

	ctx = With(ctx, k1, wantInt)
	gotInt := Value(ctx, k1, defaultInt)
	if !reflect.DeepEqual(gotInt, wantInt) {
		t.Errorf("Ctx(int) Error: got=%v, want=%v", gotInt, wantInt)
	}
	gotZero := Value(ctx, k0, defaultInt)
	if !reflect.DeepEqual(gotZero, defaultInt) {
		t.Errorf("Ctx(int default) Error: got=%v, default=%v", gotZero, defaultInt)
	}
	gotString := Value(ctx, k1, defaultStr)
	if !reflect.DeepEqual(gotString, defaultStr) {
		t.Errorf("Ctx(str default) Error: got=%v, default=%v", gotZero, defaultInt)
	}
}
