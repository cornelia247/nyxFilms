package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/cornelia247/nyxfilms/gen"
	"github.com/cornelia247/nyxfilms/metadata/pkg/model"
	"google.golang.org/protobuf/proto"
)

var metadata = &model.Metadata{
	ID: "123",
	Title: "The sounds of music",
	Description: "Just a musical",
	Director: "Cornel Stars",

}

var genMetadata = &gen.Metadata{
	Id: "123",
	Title: "The sounds of music",
	Description: "Just a musical" ,
	Director: "Cornel Stars" ,
}

func main() {
	jsonBytes, err := serialiseToJSON(metadata)
	if err != nil {
		panic(err)
	}

	xmlBytes, err := serialiseToXML(metadata)
	if err != nil {
		panic(err)
	}

	protoBytes, err := serialiseToProto(genMetadata)
	if err != nil {
		panic(err)
	}

	fmt.Printf("JSON size: \t%dB\n", len(jsonBytes))
	fmt.Printf("XML size: \t%dB\n", len(xmlBytes))
	fmt.Printf("Proto size: \t%dB\n", len(protoBytes))
}

func serialiseToJSON(m *model.Metadata)([]byte, error){
	return json.Marshal(m)
	
}
func serialiseToXML(m *model.Metadata)([]byte, error){
	return xml.Marshal(m)

}
func serialiseToProto(m *gen.Metadata)([]byte, error){
	return proto.Marshal(m)

}