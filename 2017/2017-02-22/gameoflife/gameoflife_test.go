package gameoflife

import (
	"reflect"
	"testing"
)

func TestNextGeneration(t *testing.T) {
	g := GenerationGrid{}
	want := GenerationGrid{}
	next := g.Next()
	if !reflect.DeepEqual(next, want) {
		t.Errorf("g.Next() = %v, want %v", next, want)
	}
}

func TestNextGeneration1x1AliveToDead(t *testing.T) {
	g := GenerationGrid{{true}}
	want := GenerationGrid{{false}}
	next := g.Next()
	if !reflect.DeepEqual(next, want) {
		t.Errorf("g.Next() = %v, want %v", next, want)
	}
}

func TestNextGeneration2x2LifeIsGood(t *testing.T) {
	g := GenerationGrid{
		{false, true},
		{true, true}}

	want := GenerationGrid{
		{true, true},
		{true, true}}
	next := g.Next()
	if !reflect.DeepEqual(next, want) {
		t.Errorf("g.Next() = %v, want %v", next, want)
	}
}

func TestNextGeneration2x2DeadStaysDead(t *testing.T) {
	g := GenerationGrid{
		{false, false},
		{false, false}}

	want := GenerationGrid{
		{false, false},
		{false, false}}
	next := g.Next()
	if !reflect.DeepEqual(next, want) {
		t.Errorf("g.Next() = %v, want %v", next, want)
	}
}

func TestNextGeneration2x2IsReallyGood(t *testing.T) {
	g := GenerationGrid{
		{true, false},
		{true, true}}

	want := GenerationGrid{
		{true, true},
		{true, true}}
	next := g.Next()
	if !reflect.DeepEqual(next, want) {
		t.Errorf("g.Next() = %v, want %v", next, want)
	}
}
