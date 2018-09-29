package linodego_test

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"

	. "github.com/linode/linodego"
	"github.com/dnaeon/go-vcr/recorder"
	"golang.org/x/oauth2"
)

var testingMode = recorder.ModeDisabled
var debugAPI = false
var validTestAPIKey = "NOTANAPIKEY"

var TestInstanceID int
var TestVolumeID int

func init() {
	if apiToken, ok := os.LookupEnv("LINODE_TOKEN"); ok {
		validTestAPIKey = apiToken
	}

	if apiDebug, ok := os.LookupEnv("LINODE_DEBUG"); ok {
		if parsed, err := strconv.ParseBool(apiDebug); err == nil {
			debugAPI = parsed
			log.Println("[INFO] LINODE_DEBUG being set to", debugAPI)
		} else {
			log.Println("[WARN] LINODE_DEBUG should be an integer, 0 or 1")
		}
	}

	if envFixtureMode, ok := os.LookupEnv("LINODE_FIXTURE_MODE"); ok {
		if envFixtureMode == "record" {
			log.Printf("[INFO] LINODE_FIXTURE_MODE %s will be used for tests", envFixtureMode)
			testingMode = recorder.ModeRecording
		} else if envFixtureMode == "play" {
			log.Printf("[INFO] LINODE_FIXTURE_MODE %s will be used for tests", envFixtureMode)
			testingMode = recorder.ModeReplaying
		}
	}

	if apiTestInstance, ok := os.LookupEnv("LINODE_TEST_INSTANCE"); ok {
		TestInstanceID, _ = strconv.Atoi(apiTestInstance)
		log.Printf("[INFO] LINODE_TEST_INSTANCE %d will be examined for tests", TestInstanceID)
	}

	if apiTestVolume, ok := os.LookupEnv("LINODE_TEST_VOLUME"); ok {
		TestVolumeID, _ = strconv.Atoi(apiTestVolume)
		log.Printf("[INFO] LINODE_TEST_VOLUME %d will be examined for tests", TestVolumeID)
	}

}

// testRecorder returns a go-vcr recorder and an associated function that the caller must defer
func testRecorder(t *testing.T, fixturesYaml string, testingMode recorder.Mode) (r *recorder.Recorder, recordStopper func()) {
	if t != nil {
		t.Helper()
	}

	r, err := recorder.NewAsMode(fixturesYaml, testingMode, nil)
	if err != nil {
		log.Fatalln(err)
	}

	recordStopper = func() {
		r.Stop()
	}
	return
}

// createTestClient is a testing helper to creates a linodego.Client initialized using
// environment variables and configured to record or playback testing fixtures.
// The returned function should be deferred by the caller to ensure the fixture
// recording is properly closed.
func createTestClient(t *testing.T, fixturesYaml string) (*Client, func()) {
	var (
		c      Client
		apiKey *string
	)
	if t != nil {
		t.Helper()
	}

	apiKey = &validTestAPIKey

	var recordStopper func()
	var r http.RoundTripper

	if testing.Short() {
		apiKey = nil
	}

	if len(fixturesYaml) > 0 {
		r, recordStopper = testRecorder(t, fixturesYaml, testingMode)
	} else {
		r = nil
		recordStopper = func() {}
	}

	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *apiKey})
	oc := &http.Client{
		Transport: &oauth2.Transport{
			Source: tokenSource,
			Base:   r,
		},
	}

	c = NewClient(oc)
	c.SetDebug(debugAPI)

	return &c, recordStopper
}

func TestClientAliases(t *testing.T) {
	client, _ := createTestClient(t, "")

	if client.Images == nil {
		t.Error("Expected alias for Images to return a *Resource")
	}
	if client.Instances == nil {
		t.Error("Expected alias for Instances to return a *Resource")
	}
	if client.InstanceSnapshots == nil {
		t.Error("Expected alias for Backups to return a *Resource")
	}
	if client.StackScripts == nil {
		t.Error("Expected alias for StackScripts to return a *Resource")
	}
	if client.Regions == nil {
		t.Error("Expected alias for Regions to return a *Resource")
	}
	if client.Volumes == nil {
		t.Error("Expected alias for Volumes to return a *Resource")
	}
}
