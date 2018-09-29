package linodego_test

/**
 * The tests in the examples directory demontrate use and test the library
 * in a real-use setting
 *
 * cd examples && go test -test.v
 */

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/linode/linodego"
)

func ExampleClient_ListTypes_all() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleListTypes_all")
	defer teardown()

	types, err := linodeClient.ListTypes(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID contains class:", strings.Index(types[0].ID, types[0].Class) > -1)
	fmt.Println("Plan has Ram:", types[0].Memory > 0)

	// Output:
	// ID contains class: true
	// Plan has Ram: true
}

// ExampleGetType_missing demonstrates the Error type, which allows inspecting
// the request and response.  Error codes will be the HTTP status code,
// or sub-100 for errors before the request was issued.
func ExampleClient_GetType_missing() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleGetType_missing")
	defer teardown()

	_, err := linodeClient.GetType(context.Background(), "missing-type")
	if err != nil {
		if v, ok := err.(*linodego.Error); ok {
			fmt.Println("Request was:", v.Response.Request.URL)
			fmt.Println("Response was:", v.Response.Status)
			fmt.Println("Error was:", v)
		}
	}

	// Output:
	// Request was: https://api.linode.com/v4/linode/types/missing-type
	// Response was: 404 NOT FOUND
	// Error was: [404] Not found
}

// ExampleListKernels_all Demonstrates how to list all Linode Kernels.  Paginated
// responses are automatically traversed and concatenated when the ListOptions are nil
func ExampleClient_ListKernels_all() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleListKernels_all")
	defer teardown()

	kernels, err := linodeClient.ListKernels(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// The Linode API default pagination size is 100.
	fmt.Println("Fetched > 100:", len(kernels) > 100)

	// Output:
	// Fetched > 100: true
}

func ExampleClient_ListKernels_allWithOpts() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleListKernels_allWithOpts")
	defer teardown()

	filterOpt := linodego.NewListOptions(0, "")
	kernels, err := linodeClient.ListKernels(context.Background(), filterOpt)
	if err != nil {
		log.Fatal(err)
	}

	// The Linode API default pagination size is 100.
	fmt.Println("Fetched > 100:", len(kernels) > 100)
	fmt.Println("Fetched Results/100 pages:", filterOpt.Pages > filterOpt.Results/100)
	fmt.Println("Fetched all results:", filterOpt.Results == len(kernels))

	// Output:
	// Fetched > 100: true
	// Fetched Results/100 pages: true
	// Fetched all results: true

}

func ExampleClient_ListKernels_filtered() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleListKernels_filtered")
	defer teardown()

	filterOpt := linodego.ListOptions{Filter: "{\"label\":\"Recovery - Finnix (kernel)\"}"}
	kernels, err := linodeClient.ListKernels(context.Background(), &filterOpt)
	if err != nil {
		log.Fatal(err)
	}
	for _, kern := range kernels {
		fmt.Println(kern.ID, kern.Label)
	}

	// Unordered output:
	// linode/finnix Recovery - Finnix (kernel)
	// linode/finnix-legacy Recovery - Finnix (kernel)
}

func ExampleClient_ListKernels_page1() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleListKernels_page1")
	defer teardown()

	filterOpt := linodego.NewListOptions(1, "")
	kernels, err := linodeClient.ListKernels(context.Background(), filterOpt)
	if err != nil {
		log.Fatal(err)
	}
	// The Linode API default pagination size is 100.
	fmt.Println("Fetched == 100:", len(kernels) == 100)
	fmt.Println("Results > 100:", filterOpt.Results > 100)
	fmt.Println("Pages > 1:", filterOpt.Pages > 1)
	k := kernels[len(kernels)-1]
	fmt.Println("Kernel Version in ID:", strings.Index(k.ID, k.Label) > -1)

	// Output:
	// Fetched == 100: true
	// Results > 100: true
	// Pages > 1: true
	// Kernel Version in ID: true
}

func ExampleClient_GetKernel_specific() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleGetKernel_specific")
	defer teardown()

	l32, err := linodeClient.GetKernel(context.Background(), "linode/latest-32bit")
	if err == nil {
		fmt.Println("Label starts:", l32.Label[0:9])
	} else {
		log.Fatalln(err)
	}

	l64, err := linodeClient.GetKernel(context.Background(), "linode/latest-64bit")
	if err == nil {
		fmt.Println("Label starts:", l64.Label[0:9])
	} else {
		log.Fatalln(err)
	}
	// Interference check
	fmt.Println("First Label still starts:", l32.Label[0:9])

	// Output:
	// Label starts: Latest 32
	// Label starts: Latest 64
	// First Label still starts: Latest 32
}

func ExampleClient_GetImage_missing() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleGetImage_missing")
	defer teardown()

	_, err := linodeClient.GetImage(context.Background(), "not-found")
	if err != nil {
		if v, ok := err.(*linodego.Error); ok {
			fmt.Println("Request was:", v.Response.Request.URL)
			fmt.Println("Response was:", v.Response.Status)
			fmt.Println("Error was:", v)
		}
	}

	// Output:
	// Request was: https://api.linode.com/v4/images/not-found
	// Response was: 404 NOT FOUND
	// Error was: [404] Not found
}
func ExampleClient_ListImages_all() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleListImages_all")
	defer teardown()

	filterOpt := linodego.NewListOptions(0, "")
	images, err := linodeClient.ListImages(context.Background(), filterOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fetched Results/100 pages:", filterOpt.Pages > filterOpt.Results/100)
	fmt.Println("Fetched all results:", filterOpt.Results == len(images))

	// Output:
	// Fetched Results/100 pages: true
	// Fetched all results: true

}

// ExampleListImages_notfound demonstrates that an empty slice is returned,
// not an error, when a filter matches no results.
func ExampleClient_ListImages_notfound() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleListImages_notfound")
	defer teardown()

	filterOpt := linodego.ListOptions{Filter: "{\"label\":\"not-found\"}"}
	images, err := linodeClient.ListImages(context.Background(), &filterOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Images with Label 'not-found':", len(images))

	// Output:
	// Images with Label 'not-found': 0
}

// ExampleListImages_notfound demonstrates that an error is returned by
// the API and linodego when an invalid filter is provided
func ExampleClient_ListImages_badfilter() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleListImages_badfilter")
	defer teardown()

	filterOpt := linodego.ListOptions{Filter: "{\"foo\":\"bar\"}"}
	images, err := linodeClient.ListImages(context.Background(), &filterOpt)
	if err == nil {
		log.Fatal(err)
	}
	fmt.Println("Error given on bad filter:", err)
	fmt.Println("Images on bad filter:", images) // TODO: nil would be better here

	// Output:
	// Error given on bad filter: [400] [X-Filter] Cannot filter on foo
	// Images on bad filter: []
}

func ExampleClient_ListLongviewSubscriptions_page1() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleListLongviewSubscriptions_page1")
	defer teardown()

	pageOpt := linodego.ListOptions{PageOptions: &linodego.PageOptions{Page: 1}}
	subscriptions, err := linodeClient.ListLongviewSubscriptions(context.Background(), &pageOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Longview Subscription Types:", len(subscriptions))

	// Output:
	// Longview Subscription Types: 4
}

func ExampleClient_ListStackscripts_page1() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleListStackscripts_page1")
	defer teardown()

	filterOpt := linodego.NewListOptions(1, "")
	scripts, err := linodeClient.ListStackscripts(context.Background(), filterOpt)
	if err != nil {
		log.Fatal(err)
	}
	// The Linode API default pagination size is 100.
	fmt.Println("Fetched == 100:", len(scripts) == 100)
	fmt.Println("Results > 100:", filterOpt.Results > 100)
	fmt.Println("Pages > 1:", filterOpt.Pages > 1)
	s := scripts[len(scripts)-1]
	fmt.Println("StackScript Script has shebang:", strings.Index(s.Script, "#!/") > -1)

	// Output:
	// Fetched == 100: true
	// Results > 100: true
	// Pages > 1: true
	// StackScript Script has shebang: true
}
