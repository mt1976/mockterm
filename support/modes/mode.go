package modes

import "fmt"

type Modality struct {
	Is int
}

const (
	debugMode = 1
	liveMode  = 2
	trialMode = 1
)

var (
	DEBUG Modality = Modality{Is: debugMode}
	LIVE  Modality = Modality{Is: liveMode}
	TRIAL Modality = Modality{Is: trialMode}
)

func (m *Modality) String() string {
	return fmt.Sprintf("%d", m.Is)
}

func (m *Modality) IsDebug() bool {
	return m.Is == debugMode
}

func (m *Modality) IsNotDebug() bool {
	return m.Is != debugMode
}

func (m *Modality) IsLive() bool {
	return m.Is == liveMode
}

func (m *Modality) IsNotLive() bool {
	return m.Is != liveMode
}

func (m *Modality) IsTrialMode() bool {
	return m.IsDebug()
}

func (m *Modality) IsNotTrialMode() bool {
	return m.IsNotDebug()
}
