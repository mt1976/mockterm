package modes

import "fmt"

type Modality struct {
	is int
}

const (
	debugMode = 1
	liveMode  = 2
	trialMode = 1
)

var (
	DEBUG Modality = Modality{is: debugMode}
	LIVE  Modality = Modality{is: liveMode}
	TRIAL Modality = Modality{is: trialMode}
)

func (m *Modality) Is(q Modality) bool {
	return m.is == q.is
}

func (m *Modality) String() string {
	switch m.is {
	case debugMode:
		return "Debug"
	case liveMode:
		return "Live"
	}
	return fmt.Sprintf("Mode=(%d)", m.is)
}

func (m *Modality) IsDebug() bool {
	return m.is == debugMode
}

func (m *Modality) IsNotDebug() bool {
	return m.is != debugMode
}

func (m *Modality) IsLive() bool {
	return m.is == liveMode
}

func (m *Modality) IsNotLive() bool {
	return m.is != liveMode
}

func (m *Modality) IsTrialMode() bool {
	return m.IsDebug()
}

func (m *Modality) IsNotTrialMode() bool {
	return m.IsNotDebug()
}
