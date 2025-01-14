// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (s SlurmConfigV0) SlotsPerNode() *int {
	return s.RawSlotsPerNode
}

func (s *SlurmConfigV0) SetSlotsPerNode(val *int) {
	s.RawSlotsPerNode = val
}

func (s SlurmConfigV0) SbatchArgs() []string {
	return s.RawSbatchArgs
}

func (s *SlurmConfigV0) SetSbatchArgs(val []string) {
	s.RawSbatchArgs = val
}

func (s SlurmConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedSlurmConfigV0()
}

func (s SlurmConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/hpc-cluster-slurm.json")
}

func (s SlurmConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/hpc-cluster-slurm.json")
}
