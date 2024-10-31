package main

import "testing"

func BenchmarkSerialiseToJSON(b *testing.B) {
	for i := 0; i< b.N; i++ {
		serialiseToJSON(metadata)
	}
}
func BenchmarkSerialiseToXML(b *testing.B) {
	for i := 0; i< b.N; i++ {
		serialiseToXML(metadata)
	}
}
func BenchmarkSerialiseToProto(b *testing.B) {
	for i := 0; i< b.N; i++ {
		serialiseToProto(genMetadata)
	}
}