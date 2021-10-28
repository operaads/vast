package vast

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	extensionCustomTracking = []byte(`<Extension type="testCustomTracking"><CustomTracking><Tracking event="event.1"><![CDATA[http://event.1]]></Tracking><Tracking event="event.2"><![CDATA[http://event.2]]></Tracking></CustomTracking></Extension>`)
	extensionData           = []byte(`<Extension type="testCustomTracking"><SkippableAdType>Generic</SkippableAdType></Extension>`)
	extensionFallback0      = []byte(`<Extension type="waterfall" fallback_index="0" />`)
	extensionFallback1      = []byte(`<Extension type="waterfall" fallback_index="1" />`)
)

func TestExtensionCustomTrackingMarshal(t *testing.T) {
	e := Extension{
		Type: "testCustomTracking",
		CustomTracking: []Tracking{
			{
				Event: "event.1",
				URI:   "http://event.1",
			},
			{
				Event: "event.2",
				URI:   "http://event.2",
			},
		},
	}

	// marshal the extension
	xmlExtensionOutput, err := xml.Marshal(e)
	assert.NoError(t, err)

	// assert the resulting marshaled extension
	assert.Equal(t, string(extensionCustomTracking), string(xmlExtensionOutput))
}

func TestExtensionCustomTracking(t *testing.T) {
	// unmarshal the Extension
	var e Extension
	assert.NoError(t, xml.Unmarshal(extensionCustomTracking, &e))

	// assert the resulting extension
	assert.Equal(t, "testCustomTracking", e.Type)
	assert.Empty(t, string(e.Data))
	if assert.Len(t, e.CustomTracking, 2) {
		// first event
		assert.Equal(t, "event.1", e.CustomTracking[0].Event)
		assert.Equal(t, "http://event.1", e.CustomTracking[0].URI)
		// second event
		assert.Equal(t, "event.2", e.CustomTracking[1].Event)
		assert.Equal(t, "http://event.2", e.CustomTracking[1].URI)
	}

	// marshal the extension
	xmlExtensionOutput, err := xml.Marshal(e)
	assert.NoError(t, err)

	// assert the resulting marshaled extension
	assert.Equal(t, string(extensionCustomTracking), string(xmlExtensionOutput))
}

func TestExtensionGeneric(t *testing.T) {
	// unmarshal the Extension
	var e Extension
	assert.NoError(t, xml.Unmarshal(extensionData, &e))

	// assert the resulting extension
	assert.Equal(t, "testCustomTracking", e.Type)
	assert.Equal(t, "<SkippableAdType>Generic</SkippableAdType>", string(e.Data))
	assert.Empty(t, e.CustomTracking)

	// marshal the extension
	xmlExtensionOutput, err := xml.Marshal(e)
	assert.NoError(t, err)

	// assert the resulting marshaled extension
	assert.Equal(t, string(extensionData), string(xmlExtensionOutput))
}

func TestExtensionFallback(t *testing.T) {
	// unmarshal the Extension
	var e Extension
	assert.NoError(t, xml.Unmarshal(extensionFallback0, &e))
	assert.Equal(t, "waterfall", e.Type)
	assert.Equal(t, 0, e.FallbackIndex)

	// marshal the extension
	xmlExtensionOutput, err := xml.Marshal(e)
	assert.NoError(t, err)
	assert.Equal(t, `<Extension type="waterfall"></Extension>`, string(xmlExtensionOutput))

	var e1 Extension
	assert.NoError(t, xml.Unmarshal(extensionFallback1, &e1))
	assert.Equal(t, "waterfall", e1.Type)
	assert.Equal(t, 1, e1.FallbackIndex)

	// marshal the extension
	xmlExtensionOutput, err = xml.Marshal(e1)
	assert.NoError(t, err)
	assert.Equal(t, `<Extension type="waterfall" fallback_index="1"></Extension>`, string(xmlExtensionOutput))
}
