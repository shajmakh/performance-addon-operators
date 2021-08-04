package __latency_testing

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

const (
	EXPECTED_TO_PASS = true
	EXPECTED_TO_FAIL = false
	//invalid values error messages
	UNEXPECTED_ERROR = "unexpected error"
	//TODO need to handle atoi errors if needed, check with devs about UX issues
	//unexpected /incorrect values error messages
	INCORRECT_P1 = "the environment variable"
	INCORRECT_P2 = "has incorrect value"

	//success message
	SUCCESS = ""
	//failure messages
	FAIL = "The current latency * is bigger than the expected one *"
	//skip messages
	SKIP_TEST_RUN                         = "Skip the latency test, the LATENCY_TEST_RUN set to false"
	SKIP_DISCOVERY_ENABLED_PERF_NOT_FOUND = "Discovery mode enabled, performance profile not found"
	SKIP_CPU                              = "Skip the oslat test, the profile*has less than two isolated CPUs"
	SKIP_MAX_LATENCY                      = "no maximum latency value provided, skip buckets latency check"
)

var (
	INCORRECT_CPUS_NUMBER          = fmt.Sprintf("%s LATENCY_TEST_CPUS %s", INCORRECT_P1, INCORRECT_P2)
	INCORRECT_DELAY                = fmt.Sprintf("%s LATENCY_TEST_DELAY %s", INCORRECT_P1, INCORRECT_P2)
	INCORRECT_MAX_LATENCY          = fmt.Sprintf("%s MAXIMUM_LATENCY %s", INCORRECT_P1, INCORRECT_P2)
	INCORRECT_SPECIFIC_MAX_LATENCY = fmt.Sprintf("%s *_MAXIMUM_LATENCY %s", INCORRECT_P1, INCORRECT_P2)
)

func setEnv(lat_test_delay string, lat_test_run string, runtime string, max_latency string, cpus string) {
	os.Setenv("LATENCY_TEST_DELAY", lat_test_delay)
	os.Setenv("LATENCY_TEST_RUN", lat_test_run)
	os.Setenv("LATENCY_TEST_RUNTIME", runtime)
	os.Setenv("MAXIMUM_LATENCY", max_latency)
	os.Setenv("LATENCY_TEST_CPUS", cpus)
}

var _ = table.DescribeTable("[performance] Latency Test", func(lat_test_delay string, lat_test_run string, runtime string, max_latency string, cpus string, expectToPass string, msgs []string) {
	setEnv(lat_test_delay, lat_test_run, runtime, max_latency, cpus)
	command := exec.Command("../../build/_output/bin/latency-e2e.test")
	// set var to get the output
	var out bytes.Buffer

	// set the output to our variable
	command.Stdout = &out
	err := command.Run()
	if err != nil {
		log.Println(err)
	}
	for _, msg := range msgs {
		Expect(out.String()).To(ContainSubstring(msg))
	}
	Expect(out.String()).To(ContainSubstring("SUCCESS!"))
	//fmt.Println(out.String())
},
	//valid values tests
	table.Entry("test #1", "0", "true", "10", "10", "3", EXPECTED_TO_PASS, []string{SUCCESS}),
	/*table.Entry("test #2: minimum runtime", "0", "true", "1", "10", "2", EXPECTED_TO_PASS, []string{SUCCESS}),
	table.Entry("test #3: maximum runtime", "0", "true", "18446744073709551615", "10", "2", EXPECTED_TO_PASS, []string{SUCCESS}),
	table.Entry("test #3: runtime in hours", "0", "true", "1H", "10", "2", EXPECTED_TO_PASS, []string{SUCCESS}),
	table.Entry("test #3: runtime in hours", "0", "true", "1h", "10", "2", EXPECTED_TO_PASS, []string{SUCCESS}),
	//should force stop the past 3 tests , just check if it started to run
	table.Entry("test #4: minimum latency", "0", "true", "5", "1", "2", EXPECTED_TO_PASS, []string{SUCCESS}),
	table.Entry("test #5: maximum latency", "0", "true", "5", "1024", "2", EXPECTED_TO_PASS, []string{SUCCESS}),
	table.Entry("test #6: start latency test with defaults ", "", "true", "", "", "", EXPECTED_TO_PASS, []string{SUCCESS}),
	table.Entry("test #7: test latency with defaults ", "", "", "", "", "", EXPECTED_TO_PASS, []string{SKIP_TEST_RUN}),
	table.Entry("test #8: test cpus not enough ", "", "true", "10", "10", "1", EXPECTED_TO_PASS, []string{SKIP_CPU}),
	table.Entry("test #9: no maximum latency ", "", "true", "10", "", "2", EXPECTED_TO_PASS, []string{SKIP_MAX_LATENCY}),
	table.Entry("test #2: runtime in minutes", "0", "true", "1m", "10", "2", EXPECTED_TO_PASS, []string{SUCCESS}),
	table.Entry("test #2: runtime in minutes", "0", "true", "1M", "10", "2", EXPECTED_TO_PASS, []string{SUCCESS}),
	table.Entry("test #2: runtime in seconds", "0", "true", "10s", "10", "2", EXPECTED_TO_PASS, []string{SUCCESS}),
	table.Entry("test #2: runtime in seconds", "0", "true", "10S", "10", "2", EXPECTED_TO_PASS, []string{SUCCESS}),

	//negative checks
	table.Entry("test #8: runtime is zero", "", "true", "0", "10", "3", EXPECTED_TO_FAIL, []string{UNEXPECTED_ERROR}),
	table.Entry("test #8: runtime is less than allowed", "", "true", "-1", "10", "3", EXPECTED_TO_FAIL, []string{UNEXPECTED_ERROR}),
	table.Entry("test #8: runtime is greater than allowed", "", "true", "18446744073709551617", "10", "3", EXPECTED_TO_FAIL, []string{UNEXPECTED_ERROR}),
	table.Entry("test #8: latency is less than allowed", "", "true", "5", "-1", "3", EXPECTED_TO_FAIL, []string{UNEXPECTED_ERROR}),
	///*?table.Entry("test #8: latency is greater than allowed", "", "true", "5", "2000", "3", EXPECTED_TO_FAIL, []string{UNEXPECTED_ERROR}),
	table.Entry("test #8: invalid values", "#", "true", "L", "K*", "w", EXPECTED_TO_FAIL, []string{UNEXPECTED_ERROR}),
	*/
)
