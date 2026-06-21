// Package bus wires the vllm inference engine onto the NATS observer bus via the shared
// cofiswarm-observer-sdk service component: it announces presence and serves the engine's
// .infer.vllm.* capability subjects (the bus equivalents of its HTTP /healthz and /v1/info).
package bus

import (
	"github.com/keepdevops/cofiswarm-observer-sdk/pkg/servicecomponent"
)

// Capability subjects (must match observer's bus/subjects.py).
const (
	SubjInfo   = servicecomponent.Prefix + ".infer.vllm.info"
	SubjHealth = servicecomponent.Prefix + ".infer.vllm.health"
)

// Routes wires the engine's capability subjects. Reply field names carry schema_version for
// the major-version gate, mirroring the resource components (kvpool/slot-manager/launcher).
func Routes(engine string) map[string]servicecomponent.Handler {
	return map[string]servicecomponent.Handler{
		SubjInfo:   infoHandler(engine),
		SubjHealth: healthHandler(),
	}
}

func infoHandler(engine string) servicecomponent.Handler {
	return func([]byte) (any, error) {
		return infoReply{
			SchemaVersion: servicecomponent.SchemaVersion, OK: true,
			Engine: engine, Stub: true,
			Note: "run deploy/Dockerfile for full vllm-metal",
		}, nil
	}
}

func healthHandler() servicecomponent.Handler {
	return func([]byte) (any, error) {
		return healthReply{SchemaVersion: servicecomponent.SchemaVersion, OK: true, Status: "ok"}, nil
	}
}

type infoReply struct {
	SchemaVersion string `json:"schema_version"`
	OK            bool   `json:"ok"`
	Error         string `json:"error,omitempty"`
	Engine        string `json:"engine"`
	Stub          bool   `json:"stub"`
	Note          string `json:"note,omitempty"`
}

type healthReply struct {
	SchemaVersion string `json:"schema_version"`
	OK            bool   `json:"ok"`
	Error         string `json:"error,omitempty"`
	Status        string `json:"status"`
}
